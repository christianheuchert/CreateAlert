package getStaffByBuilding

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	Staff := RestCallGetStaffByBuilding(input.IP, input.CustomerId, input.Username, input.Password, input.Building)

	output := &Output{Staff: Staff}

	// fmt.Println("Output: ", output.Staff)
	// ctx.Logger().Info("Output: ", output)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

//RestCallGetStaffByBuilding ---------------------------------------------------- RestCallGetStaffByBuilding
func RestCallGetStaffByBuilding(IP string, customerId string, uname string, pword string, building string) []string {
	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "http://" + IP + "/XpertRestApi/api/Staff/GetAll?CustomerId=" + customerId+"&NumberOfRecords=10000"
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []string{}
	}

	translatedData := &RestCallGetAllStaffReponse{}
	json.Unmarshal(body, &translatedData)

	
	buildingId, _ := strconv.Atoi(building) // var checking if building is int Id
	if (buildingId == 0){buildingId = -1}
	listOfStaff := []Staff{}
	for _, obj := range translatedData.List {
		// Check if the targetValue is equal to either "CurrentBuildingName" or "CurrentBuildingID"
		if obj.CurrentBuildingName == building || obj.CurrentBuildingID == buildingId{
			// fmt.Println("found object with corresponding building: object id ", obj.ID)
			listOfStaff = append(listOfStaff, obj)
		}
	}

	var jsonStrings []string // return data as string array
	for _, asset := range listOfStaff {
		jsonData, err2 := json.Marshal(asset)
		if err2 != nil {
			fmt.Println("Error:", err)
			return []string{}
		}
		jsonStrings = append(jsonStrings, string(jsonData))
	}

	return jsonStrings
}