package SendEmail

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

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

	SendEmailResponse := RestCallSendEmailAdvanced(input.IP, input.CustomerId, input.Username, input.Password, input.UserEmailAddress, input.EmailSubject, input.EmailMessage)


	output := &Output{}
	if (SendEmailResponse.ElapsedTimeInMillseconds!=0){
		output.SentBoolean = true
	}
	// fmt.Println("Output: ", output)
	// ctx.Logger().Info("Output: ", output)

	
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

func RestCallSendEmailAdvanced(IP string, customerId string, username string, password string, userEmail string, emailSubject string, emailMessage string )SendEmailResponse{

	//Declare response struct object
	var response SendEmailResponse
	// query escape text items
	cleanUserEmail := url.QueryEscape(userEmail)
	cleanEmailSubject := url.QueryEscape(emailSubject)
	cleanEmailMessage := url.QueryEscape(emailMessage)

	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "http://"+IP+"/XpertRestApi/api/Users/SendEmailAdvanced?userEmailAddress=" + cleanUserEmail + "&subject=" + cleanEmailSubject + 
	"&emailMessage=" + cleanEmailMessage + "&CustomerId=" + customerId
	
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return response
	}

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
	body, _ := io.ReadAll(resp.Body)

	// Unmarshal the config JSON into an object
	errUnmarshal := json.Unmarshal(body, &response)
	if errUnmarshal != nil {
	 	fmt.Println(errUnmarshal)
		return response
	}

	return response
}