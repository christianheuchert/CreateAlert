package getAllUsersByGroup

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

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

	Response := RestCallGetUsersByGroup(input.IP, input.CustomerId, input.Username, input.Password, input.Group)

	output := &Output{Users: Response}

	// fmt.Println("Output: ", output.Users)
	// ctx.Logger().Info("Output: ", output)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

func RestCallGetUsersByGroup(IP string, customerId string, uname string, pword string, group string) []string {

	// Create an HTTP client
	client := &http.Client{}

	var groupValue interface{}
	if intValue, err := strconv.Atoi(group); err == nil {
		groupValue = intValue
	} else {
		groupValue = group
	}

	// Create the request
	url := "http://" + IP + "/XpertRestApi/api/Users/GetAll?CustomerId=" + customerId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return []string{}
	}

	// Add basic authentication to the request header
	auth := uname + ":" + pword
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return []string{}
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []string{}
	}

	translatedData := &GetAllUsersResponse{}
	json.Unmarshal([]byte(string(body)), &translatedData)

	listOfUsers := []User{}

	for _, obj := range translatedData.List {
		if len(obj.AssociatedGroups) > 0 {
			for _, grp := range obj.AssociatedGroups {
				if grp.ID == groupValue || grp.Name == groupValue {
					listOfUsers = append(listOfUsers, obj)
				}
			}
		}
	}

	var jsonStrings []string // return data as string array
	for _, asset := range listOfUsers {
		jsonData, err2 := json.Marshal(asset)
		if err2 != nil {
			fmt.Println("Error:", err)
			return []string{}
		}
		jsonStrings = append(jsonStrings, string(jsonData))
	}
	return jsonStrings
}
