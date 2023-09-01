package CreateAlert

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})


// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	InsertAlertResponse := PostCallInsertEvent(input.IP, input.CustomerId, input.Username, input.Password, input.AlertDisplayName)

	output := &Output{}
	if (InsertAlertResponse.ElapsedTimeInMillseconds!=0){
		output.AlertBoolean = true
	}
	
	// fmt.Println("Response: ", InsertAlertResponse)
	// ctx.Logger().Info("Output: ", output)


	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

func PostCallInsertEvent(IP string, customerId string, uname string, pword string, dispName string) InsertEventResponse {
	username := uname
	password := pword

	//Declare response struct object
	var response InsertEventResponse
	currentTime := time.Now().UTC()
	timeFormat := "2006-01-02T15:04:05.999Z"
	formattedTime := currentTime.Format(timeFormat)

	// Create an HTTP client
	client := &http.Client{}

	event := EventModel{
		AllowedValueRange: "Flogo Allowed Value Range",
		Details: "Flogo Details",
		DeviceID: 0,
		DeviceLogID: 0,
		DisplayName:  dispName,
		EndDateTime: formattedTime,
		IsAcknowledgementRequired: true,
		ItemID: 0,
		MaxValue: 0,
		MinValue: 0,
		PlanID: 0,
		RuleName: "Flogo Rule Name",
		RuleSetMajorVersion: 0,
		RuleSetMinorVersion: 0,
		RuleSetName: "Flogo Rule Set Name",
		SeverityColor: "Flogo Severity Color",
		SeverityIcon: "Flogo Severity Icon",
		SiteID: "Flogo SiteId",
		StartDateTime: formattedTime,
		SystemName: "AREA_ALERT",
		Type: "Flogo Type", 
		UseCase: 0,
		ViolationValue: "Flogo Violation Value",
		DateUpdated: formattedTime,
		Description: "Flogo Description",
	}
	payload, err := json.Marshal(event) // make event JSON bytes
	if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return response
    }

	// Create the request
	url := "http://" + IP + "/XpertRestApi/api/Events/InsertEvent?CustomerId=" + customerId
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return response
	}
	req.Header.Set("Content-Type", "application/json") //Set the Content-Type header

	// Add basic authentication to the request header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return response
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return response
	}
	
	// Unmarshal the config JSON into an object
	errUnmarshal := json.Unmarshal(body, &response)
	if errUnmarshal != nil {
	 	fmt.Println(errUnmarshal)
		return response
	}

	return response
}
