package getDepartments

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

	departments := RestCallGetDepartments(input.IP, input.CustomerId, input.Username, input.Password)

	output := &Output{Departments: departments}

	// fmt.Println("Output: ", output.Departments[0])
	// ctx.Logger().Info("Output: ", output)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

//http://52.45.17.177:802/XpertRestApi/api/MetaData/GetGroups?CustomerId=1
func RestCallGetDepartments(IP string, customerId string, username string, password string )[]string {

	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "http://"+IP+"/XpertRestApi/api/MetaData/GetDepartments?CustomerId="+customerId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return []string{}
	}

	// Add basic authentication to the request header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return []string{}
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// Unmarshal the respnse JSON into an object
	var departments Response
	errUnmarshal := json.Unmarshal(body, &departments)
	if errUnmarshal != nil {
	 	fmt.Println(errUnmarshal)
		return []string{}
	}

	var jsonStrings []string // return data as string array
	for _, asset := range departments.List {
		jsonData, err2 := json.Marshal(asset)
		if err2 != nil {
			fmt.Println("Error:", err)
			return []string{}
		}
		jsonStrings = append(jsonStrings, string(jsonData))
	}

	return jsonStrings
}