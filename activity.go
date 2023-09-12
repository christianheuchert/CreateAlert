package getStaffByZone

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

	Staff := RestCallGetStaffByZone(input.IP, input.CustomerId, input.Username, input.Password, input.Zone)

	output := &Output{Staff: Staff}

	// fmt.Println("Output: ", output.Staff)
	// ctx.Logger().Info("Output: ", output)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

// RestCallGetStaffByZoneGroup("52.45.17.177:802", 2047, "afadmin", "admin", 12091)
func RestCallGetStaffByZone(IP string, customerId string, uname string, pword string, zone string) []string {
	username := uname
	password := pword
	ListOfObjsWithZone := []Staff{}

	// Create an HTTP client
	client := &http.Client{}

	// Create the request ADDED &NumberOfRecords=10000. NOT GOOD
	url := "http://" + IP + "/XpertRestApi/api/Staff/GetAll?CustomerId=" + (customerId)+"&NumberOfRecords=10000"
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

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []string{}
	}

	translatedData := &GetAllStaffReponse{}
	json.Unmarshal(body, &translatedData)

	var zoneObj Zone
	objectCheck := json.Unmarshal([]byte(zone), &zoneObj) // check if Zone Object
	if (objectCheck == nil){ // Zone is object
		zone = strconv.Itoa(zoneObj.ZoneID)
	}

	ZoneId, _ := strconv.Atoi(zone) // var checking if building is int Id, if not, set to -1
	if (ZoneId == 0){ZoneId = -1} 
	for _, obj := range translatedData.List {

		var zones []Zone
		err := json.Unmarshal([]byte(obj.CurrentZones), &zones)
		if err != nil {
			continue
		} else if len(zones) == 0 {
			continue
		}
		
		if zones[0].ZoneID == ZoneId || zones[0].ZoneName == zone {
			ListOfObjsWithZone = append(ListOfObjsWithZone, obj)
		}
	}

	var jsonStrings []string // return data as string array
	for _, asset := range ListOfObjsWithZone {
		jsonData, err2 := json.Marshal(asset)
		if err2 != nil {
			fmt.Println("Error:", err)
			return []string{}
		}
		jsonStrings = append(jsonStrings, string(jsonData))
	}


	return jsonStrings
}