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

type DeletedData struct {
	Timestamp    int64  `json:"timestamp"`
	WebhookEvent string `json:"webhookEvent"`
	User         struct {
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
	} `json:"user"`
	Issue struct {
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
		Fields struct {
			Statuscategorychangedate string `json:"statuscategorychangedate"`
			Issuetype                struct {
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
			Timespent        any `json:"timespent"`
			Customfield10030 any `json:"customfield_10030"`
			Project          struct {
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
			FixVersions        []any `json:"fixVersions"`
			Aggregatetimespent any   `json:"aggregatetimespent"`
			Resolution         struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"resolution"`
			Customfield10027 any    `json:"customfield_10027"`
			Customfield10028 any    `json:"customfield_10028"`
			Customfield10029 any    `json:"customfield_10029"`
			Resolutiondate   string `json:"resolutiondate"`
			Workratio        int    `json:"workratio"`
			Watches          struct {
				Self       string `json:"self"`
				WatchCount int    `json:"watchCount"`
				IsWatching bool   `json:"isWatching"`
			} `json:"watches"`
			LastViewed       string `json:"lastViewed"`
			Issuerestriction struct {
				Issuerestrictions struct {
				} `json:"issuerestrictions"`
				ShouldDisplay bool `json:"shouldDisplay"`
			} `json:"issuerestriction"`
			Created          string `json:"created"`
			Customfield10020 any    `json:"customfield_10020"`
			Customfield10021 any    `json:"customfield_10021"`
			Customfield10022 any    `json:"customfield_10022"`
			Priority         struct {
				Self    string `json:"self"`
				IconURL string `json:"iconUrl"`
				Name    string `json:"name"`
				ID      string `json:"id"`
			} `json:"priority"`
			Customfield10023 any    `json:"customfield_10023"`
			Customfield10024 any    `json:"customfield_10024"`
			Customfield10025 string `json:"customfield_10025"`
			Labels           []any  `json:"labels"`
			Customfield10026 any    `json:"customfield_10026"`
			Customfield10016 any    `json:"customfield_10016"`
			Customfield10017 any    `json:"customfield_10017"`
			Customfield10018 struct {
				HasEpicLinkFieldDependency bool `json:"hasEpicLinkFieldDependency"`
				ShowField                  bool `json:"showField"`
				NonEditableReason          struct {
					Reason  string `json:"reason"`
					Message string `json:"message"`
				} `json:"nonEditableReason"`
			} `json:"customfield_10018"`
			Customfield10019              string `json:"customfield_10019"`
			Aggregatetimeoriginalestimate any    `json:"aggregatetimeoriginalestimate"`
			Timeestimate                  any    `json:"timeestimate"`
			Versions                      []any  `json:"versions"`
			Issuelinks                    []any  `json:"issuelinks"`
			Assignee                      struct {
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
			} `json:"assignee"`
			Updated string `json:"updated"`
			Status  struct {
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
			Components           []any `json:"components"`
			Timeoriginalestimate any   `json:"timeoriginalestimate"`
			Description          any   `json:"description"`
			Customfield10010     any   `json:"customfield_10010"`
			Customfield10014     any   `json:"customfield_10014"`
			Timetracking         struct {
			} `json:"timetracking"`
			Customfield10015      any    `json:"customfield_10015"`
			Customfield10005      any    `json:"customfield_10005"`
			Customfield10006      any    `json:"customfield_10006"`
			Customfield10007      any    `json:"customfield_10007"`
			Security              any    `json:"security"`
			Customfield10008      any    `json:"customfield_10008"`
			Attachment            []any  `json:"attachment"`
			Customfield10009      any    `json:"customfield_10009"`
			Aggregatetimeestimate any    `json:"aggregatetimeestimate"`
			Summary               string `json:"summary"`
			Creator               struct {
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
			} `json:"creator"`
			Subtasks []any `json:"subtasks"`
			Reporter struct {
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
			} `json:"reporter"`
			Aggregateprogress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"aggregateprogress"`
			Customfield10001 any `json:"customfield_10001"`
			Customfield10002 any `json:"customfield_10002"`
			Customfield10003 any `json:"customfield_10003"`
			Customfield10004 any `json:"customfield_10004"`
			Environment      any `json:"environment"`
			Duedate          any `json:"duedate"`
			Progress         struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"progress"`
			Votes struct {
				Self     string `json:"self"`
				Votes    int    `json:"votes"`
				HasVoted bool   `json:"hasVoted"`
			} `json:"votes"`
			Comment struct {
				Comments   []any  `json:"comments"`
				Self       string `json:"self"`
				MaxResults int    `json:"maxResults"`
				Total      int    `json:"total"`
				StartAt    int    `json:"startAt"`
			} `json:"comment"`
			Worklog struct {
				StartAt    int   `json:"startAt"`
				MaxResults int   `json:"maxResults"`
				Total      int   `json:"total"`
				Worklogs   []any `json:"worklogs"`
			} `json:"worklog"`
		} `json:"fields"`
	} `json:"issue"`
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

	var eventData DeletedData

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
	url := "https://cliq.zoho.in/company/60023389010/api/v2/channelsbyname/testalerts/message?zapikey=" + apiToken
	issueLink := jiraUrl + "/browse/" + issueKey
	message := map[string]interface{}{
		"text": "Jira Updates \n" + "The Issue " + issueKey + " has been Deletd in Jira" + "\n Project Name:   " + projectName + "\n Issue ID:   " + issueKey + "\n Issue Summary:   " + issueSummary,
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
