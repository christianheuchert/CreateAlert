package ParseXpertMessage

import (
	"encoding/json"
	"fmt"
	"strconv"

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

	// Unmarshal input into XpertMsg struct
	var XpertMessage XpertMsg
	errUnmarshal := json.Unmarshal([]byte(input.XpertMessage), &XpertMessage)
	if errUnmarshal != nil {
		fmt.Println(errUnmarshal)
		return true, err
	}

	// Unmarshal zones into an array of strings 
	var zones []Zone
	var zoneJsonStrings []string // return data as string array
	zonesCheck := json.Unmarshal([]byte(XpertMessage.DeviceReports[0].RTLSModel2D.PosZoneIDs), &zones)
	if (zonesCheck == nil){ // Zone array exists
		for _, zoneObj := range zones {
			jsonData, err2 := json.Marshal(zoneObj)
			if err2 != nil {
				fmt.Println("Error:", err)
			}
			zoneJsonStrings = append(zoneJsonStrings, string(jsonData))
		}
	}
	fmt.Println("test: ", XpertMessage.DeviceReports[0].RTLSModel2D.PosX)

	output := &Output{
		DeviceMAC: XpertMessage.DeviceReports[0].DeviceUniqueID,
		Timestamp: XpertMessage.DeviceReports[0].DataTimestamp,
		DeviceLogId: strconv.Itoa(XpertMessage.DeviceReports[0].DeviceLogID),
		StatusReportReason: strconv.Itoa(XpertMessage.DeviceReports[0].Status.DeviceReportReason),
		BatteryLevel: strconv.FormatFloat(XpertMessage.DeviceReports[0].Status.BatteryLevel1, 'f', -1, 64),
		Temperature: "",
		Humidity: "",
		MapId: strconv.Itoa(XpertMessage.DeviceReports[0].RTLSModel2D.PosMapID),
		X: strconv.FormatFloat(XpertMessage.DeviceReports[0].RTLSModel2D.PosX, 'f', -1, 64),
		Y: strconv.FormatFloat(XpertMessage.DeviceReports[0].RTLSModel2D.PosY, 'f', -1, 64),
		Zone: zoneJsonStrings,
		GeoLattitude: "",
		GeoLongitude: "",
		ItemId: strconv.Itoa(XpertMessage.DeviceReports[0].Item.ItemID), 
		DisplayName: XpertMessage.DeviceReports[0].RTLSModel2D.PosDisplayName,
	}

	// ctx.Logger().Info("Output: ", output.Zone)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

// {
// 	"AllowedValueRange": "Flogo Alert AllowedValueRange",
// 	"Details": "Flogo Alert Details",
// 	"DeviceId": 0,
// 	"DeviceLogId": 0,
// 	"DisplayName": "Flogo Alert Display Name",
// 	"EndDateTime": "2023-08-30T18:54:22.578Z",
// 	"IsAcknowledgementRequired": true,
// 	"ItemId": 0,
// 	"MaxValue": 0,
// 	"MinValue": 0,
// 	"PlanId": 0,
// 	"RuleName": "Flogo Alert Rule Name",
// 	"RuleSetMajorVersion": 0,
// 	"RuleSetMinorVersion": 0,
// 	"RuleSetName": "Flogo Alert Rule Set Name",
// 	"SeverityColor": "Flogo Blue",
// 	"SeverityIcon": "Flogo Icon",
// 	"SiteId": "Flogo Alert SiteId",
// 	"StartDateTime": "2023-08-30T18:54:22.578Z",
// 	"SystemName": "AREA_ALERT",
// 	"Type": "Flogo Alert Type",
// 	"UseCase": 5,
// 	"ViolationValue": "Flogo Alert ViolationValue",
// 	"DateUpdated": "2023-08-30T18:54:22.578Z",
// 	"Description": "Flogo Alert Description"
// }