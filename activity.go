package ParseUnknownXpertMessage

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata( &Input{}, &Output{})

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

	// Unmarshal the JSON string into a map[string]interface{}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(input.XpertMessageJSON), &data); err != nil {
		fmt.Println("Error converting JSON to Map")
	}

	output := &Output{
		DeviceMAC: extractValues(data, input.DeviceMACTarget),
		Timestamp: extractValues(data, input.TimestampTarget),
		DeviceLogId: extractValues(data, input.DeviceLogIdTarget),
		StatusReportReason: extractValues(data, input.StatusReportReasonTarget),
		BatteryLevel: extractValues(data, input.BatteryLevelTarget),
		Temperature: "",
		Humidity: "",
		MapId: extractValues(data, input.MapIdTarget),
		X: extractValues(data, input.XTarget),
		Y: extractValues(data, input.YTarget),
		Zone: extractValues(data, input.ZoneTarget),
		GeoLattitude: "",
		GeoLongitude: "",
		ItemId: extractValues(data, input.ItemIdTarget), 
		DisplayName: extractValues(data, input.DisplayNameTarget),
	}

	// ctx.Logger().Info("Output: ", output)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

func extractValues (JSON map[string]interface{}, target string) string{

	// Split the target into dot-separated keys
	keys := strings.Split(target, ".")

	// Traverse the map using the keys
	var value interface{}
	current := JSON
	for _, key := range keys {
		val, ok := current[key]
		if ok {
			// Update 'current' to the value of the current key
			current, _ = val.(map[string]interface{})
			value = val
		}
	}

	if value == nil {
		return ""
	}

	return fmt.Sprintf("%v", value)
}