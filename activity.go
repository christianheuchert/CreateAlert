package getAllByDept

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

	staffInDept := RestCallGetAllByDepartment(input.IP, input.CustomerId, input.Username, input.Password, input.DepartmentItem)

	output := &Output{Staff: staffInDept}

	// fmt.Println("Output: ", output.Staff)
	// ctx.Logger().Info("Output: ", output)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

//RestCallGetStaffByDepartment ---------------------------------------------------- RestCallGetStaffByDepartment
//http://52.45.17.177:802/XpertRestApi/api/MetaData/GetDepartments?CustomerId=1
func RestCallGetAllByDepartment(IP string, customerId string, username string, password string, departmentItem string)[]string {
	departmentId := ""

	var dept Department
	objectCheck := json.Unmarshal([]byte(departmentItem), &dept) // check if Department
	_, intCheck := strconv.Atoi(departmentItem) // check if Int ID

	if (objectCheck == nil){ // obj
		departmentId = strconv.Itoa(dept.ID)
	}else if(intCheck == nil){// int
		departmentId = departmentItem
	}else{ // name
		Response := RestCallGetDepartments(IP, customerId, username, password)
		for _, element := range Response.List{
			if (element.Name == departmentItem){// find Department with Name and use Id
				departmentId = strconv.Itoa(element.ID)
				break
			}
		}
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "http://"+IP+"/XpertRestApi/api/Staff/GetAllByDepartment?CustomerId="+customerId+"&DepartmentId="+departmentId+"&NumberOfRecords=10000"
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

	// Unmarshal request into Asset json
	var staffResponse GetAllByDepartmentResponse
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

// Helper function to get ID of Department with Name Input
func RestCallGetDepartments(IP string, customerId string, username string, password string )GetDepartmentsResponse{

	//Declare response struct object
	var response GetDepartmentsResponse

	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "http://"+IP+"/XpertRestApi/api/MetaData/GetDepartments?CustomerId="+customerId
	req, err := http.NewRequest("GET", url, nil)
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
	errUnmarshal := json.Unmarshal([]byte(body), &response)
	if errUnmarshal != nil {
	 	fmt.Println(errUnmarshal)
		return response
	}

	return response
}