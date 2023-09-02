package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type CommmentData struct {
	Timestamp    int64  `json:"timestamp"`
	WebhookEvent string `json:"webhookEvent"`
	Comment      struct {
		Self   string `json:"self"`
		ID     string `json:"id"`
		Author struct {
			Self       string `json:"self"`
			AccountID  string `json:"accountId"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
			AccountType string `json:"accountType"`
		} `json:"author"`
		Body         string `json:"body"`
		UpdateAuthor struct {
			Self       string `json:"self"`
			AccountID  string `json:"accountId"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
			AccountType string `json:"accountType"`
		} `json:"updateAuthor"`
		Created   string `json:"created"`
		Updated   string `json:"updated"`
		JsdPublic bool   `json:"jsdPublic"`
	} `json:"comment"`
	Issue struct {
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
		Fields struct {
			Summary   string `json:"summary"`
			Issuetype struct {
				Self           string `json:"self"`
				ID             string `json:"id"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				Subtask        bool   `json:"subtask"`
				AvatarID       int    `json:"avatarId"`
				EntityID       string `json:"entityId"`
				HierarchyLevel int    `json:"hierarchyLevel"`
			} `json:"issuetype"`
			Project struct {
				Self           string `json:"self"`
				ID             string `json:"id"`
				Key            string `json:"key"`
				Name           string `json:"name"`
				ProjectTypeKey string `json:"projectTypeKey"`
				Simplified     bool   `json:"simplified"`
				AvatarUrls     struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
			} `json:"project"`
			Assignee any `json:"assignee"`
			Priority struct {
				Self    string `json:"self"`
				IconURL string `json:"iconUrl"`
				Name    string `json:"name"`
				ID      string `json:"id"`
			} `json:"priority"`
			Status struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
		} `json:"fields"`
	} `json:"issue"`
	EventType string `json:"eventType"`
}

func LambdaHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Check if the JSON data is empty
	if event.Body == "" {
		log.Println("Empty JSON data")
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}
	// Check if the query parameter is eqal to the env
	// Get lamda cred from env
	lambdaCred := os.Getenv("LAMBDA_CRED")
	if lambdaCred == "" {
		panic("jira_URL environment variable is not set")
	}
	customParam, paramExists := event.QueryStringParameters["lamda-auth"]
	if !paramExists || customParam != lambdaCred {
		// Return a response indicating that the parameter is missing or has an invalid value
		return events.APIGatewayProxyResponse{
			StatusCode: 400, // Bad Request
			Body:       "The 'Authenticaion' query parameter is missing or has an invalid value.",
		}, nil
	}

	var eventData CommmentData

	// Unmarshal the JSON data
	if err := json.Unmarshal([]byte(event.Body), &eventData); err != nil {
		log.Printf("Error unmarshaling JSON: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	// Extract the required fields
	issueKey := eventData.Issue.Key
	issueSummary := eventData.Issue.Fields.Summary

	// Extract the project name
	projectName := eventData.Issue.Fields.Project.Name

	// Construct the output
	output := fmt.Sprintf("Issue Key: %s\nSummary: %s\nAssignee Display Name: %s", issueKey, issueSummary, projectName)

	// Log the output (you can remove this in production)
	SendZohoMessge(issueKey, issueSummary, projectName)

	// Return a successful response with the extracted data
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       output,
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}

func SendZohoMessge(issueKey string, issueSummary string, projectName string) {
	// URL for sending messages to the channel
	apiToken := os.Getenv("ZOHO_CLIQ_API_TOKEN")
	if apiToken == "" {
		panic("ZOHO_CLIQ_API_TOKEN environment variable is not set")
	}

	// Jira Url
	jiraUrl := os.Getenv("JIRA_URL")
	if apiToken == "" {
		panic("jira_URL environment variable is not set")
	}
	channelId := os.Getenv("CHANNEL_ENDPOINT")
	if apiToken == "" {
		panic("jira_URL environment variable is not set")
	}
	url := channelId + "?zapikey=" + apiToken
	issueLink := jiraUrl + "/browse/" + issueKey
	message := map[string]interface{}{
		"text": "Jira Updates \n" + "A new comment added in the Issue " + issueKey + "\n Project Name:   " + projectName + "\n Issue ID:   " + issueKey + "\n Issue Summary:   " + issueSummary,
		"card": map[string]interface{}{
			"theme":     "prompt",
			"thumbnail": "https://www.zoho.com/cliq/help/restapi/images/announce_icon.png",
		},
		"buttons": []map[string]interface{}{
			{
				"label": "View Issue",
				"type":  "+",
				"action": map[string]interface{}{
					"type": "open.url",
					"data": map[string]interface{}{
						"web": issueLink,
					},
				},
			},
		},
	}

	// Convert the message to JSON
	payload, _ := json.Marshal(message)

	// Create an HTTP client
	client := &http.Client{}

	// Create a request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	// Set the API token as an HTTP header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the response status code and body
	println("Response Status Code:", resp.Status)
}
