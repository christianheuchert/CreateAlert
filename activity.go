package getAllByGroup

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

	staffInGroup := RestCallGetAllByGroup(input.IP, input.CustomerId, input.Username, input.Password, input.GroupItem)

	output := &Output{Staff: staffInGroup}

	// fmt.Println("Output: ", output.Staff)
	// ctx.Logger().Info("Output: ", output)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

//RestCallGetAllByGroup ---------------------------------------------------- RestCallGetAllByGroup
//http://52.45.17.177:802/XpertRestApi/api/MetaData/GetAllByGroup?CustomerId=1
func RestCallGetAllByGroup(IP string, customerId string, username string, password string, groupItem string)[]string{
	groupId := ""

	var group Group
	objectCheck := json.Unmarshal([]byte(groupItem), &group) // check if Group
	_, intCheck := strconv.Atoi(groupItem) // check if Int ID

	if (objectCheck == nil){ // obj
		groupId = strconv.Itoa(group.ID)
	}else if(intCheck == nil){// int
		groupId = groupItem
	}else{ // name
		Response := RestCallGetGroups(IP, customerId, username, password)
		for _, element := range Response.List{
			if (element.Name == groupItem){// find Group with Name and use Id
				groupId = strconv.Itoa(element.ID)
				break
			}
		}
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "http://"+IP+"/XpertRestApi/api/Staff/GetAllByGroup?CustomerId="+customerId+"&GroupId="+groupId+"&NumberOfRecords=10000"
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

	// Unmarshal request into Staff struct
	var staffResponse GetAllByGroupResponse
	errUnmarshal := json.Unmarshal(body, &staffResponse)
	if errUnmarshal != nil {
	 	fmt.Println(errUnmarshal)
	}

	var jsonStrings []string // return data as string array
	for _, asset := range staffResponse.List {
		jsonData, err2 := json.Marshal(asset)
		if err2 != nil {
			fmt.Println("Error:", err)
			return []string{}
		}
		jsonStrings = append(jsonStrings, string(jsonData))
	}

	return jsonStrings
}

//RestCallGetGroups ---------------------------------------------------- RestCallGetGroups
//http://52.45.17.177:802/XpertRestApi/api/MetaData/GetGroups?CustomerId=1
func RestCallGetGroups(IP string, customerId string, username string, password string )GetGroupsResponse{
	fmt.Println("rest call")
	var responseArr GetGroupsResponse
	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "http://"+IP+"/XpertRestApi/api/MetaData/GetGroups?CustomerId="+customerId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return responseArr
	}

	// Add basic authentication to the request header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return responseArr
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// Unmarshal the config JSON into an object	
	errUnmarshal := json.Unmarshal([]byte(body), &responseArr)
	if errUnmarshal != nil {
	 	fmt.Println(errUnmarshal)
		return responseArr
	}

	return responseArr
}