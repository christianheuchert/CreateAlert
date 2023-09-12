package ParseXpertMessage

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	XpertMessage string `md:"XpertMessage,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["XpertMessage"])
	r.XpertMessage = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"XpertMessage": r.XpertMessage,
	}
}

type Output struct {
	DeviceMAC string `md:"DeviceMAC"`
	Timestamp string `md:"Timestamp"`
	DeviceLogId string `md:"DeviceLogId"`
	StatusReportReason string `md:"StatusReportReason"`
	BatteryLevel string `md:"BatteryLevel"`
	Temperature string `md:"Temperature"`
	Humidity string `md:"Humidity"`
	MapId string `md:"MapId"`
	X string `md:"X"`
	Y string `md:"Y"`
	Zone []string `md:"Zone"`
	GeoLattitude string `md:"GeoLattitude"`
	GeoLongitude string `md:"GeoLongitude"`
	ItemId string `md:"ItemId"`
	DisplayName string `md:"DisplayName"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["DeviceMAC"])
	o.DeviceMAC = strVal
	strVal, _ = coerce.ToString(values["Timestamp"])
	o.Timestamp = strVal
	strVal, _ = coerce.ToString(values["DeviceLogId"])
	o.DeviceLogId = strVal
	strVal, _ = coerce.ToString(values["StatusReportReason"])
	o.StatusReportReason = strVal
	strVal, _ = coerce.ToString(values["BatteryLevel"])
	o.BatteryLevel = strVal
	strVal, _ = coerce.ToString(values["Temperature"])
	o.Temperature = strVal
	strVal, _ = coerce.ToString(values["Humidity"])
	o.Humidity = strVal
	strVal, _ = coerce.ToString(values["MapId"])
	o.MapId = strVal
	strVal, _ = coerce.ToString(values["X"])
	o.X = strVal
	strVal, _ = coerce.ToString(values["Y"])
	o.Y = strVal
	zoneList, ok:= (values["Zone"]).([]string) // type assertion
	if ok {
		o.Zone = zoneList
	}
	strVal, _ = coerce.ToString(values["GeoLattitude"])
	o.GeoLattitude = strVal
	strVal, _ = coerce.ToString(values["GeoLongitude"])
	o.GeoLongitude = strVal
	strVal, _ = coerce.ToString(values["ItemId"])
	o.ItemId = strVal
	strVal, _ = coerce.ToString(values["DisplayName"])
	o.DisplayName = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"DeviceMAC": o.DeviceMAC,
		"Timestamp": o.Timestamp,
		"DeviceLogId": o.DeviceLogId,
		"StatusReportReason": o.StatusReportReason,
		"BatteryLevel": o.BatteryLevel,
		"Temperature": o.Temperature,
		"Humidity": o.Humidity,
		"MapId": o.MapId,
		"X": o.X,
		"Y": o.Y,
		"Zone": o.Zone,
		"GeoLattitude": o.GeoLattitude,
		"GeoLongitude": o.GeoLongitude,
		"ItemId": o.ItemId,
		"DisplayName": o.DisplayName,
	}
}

type XpertMsg struct {
	DeviceReports    []struct {
		Attributes                []interface{}  `json:"Attributes"`
		AutoidEpc                 string `json:"AUTOID_EPC"`
		DataTimestamp             string `json:"DataTimestamp"`
		DeviceActionBit           bool   `json:"DeviceActionBit"`
		DeviceLogID               int    `json:"DeviceLogID"`
		DeviceModel               string `json:"DeviceModel"`
		DeviceSerialNumber        string `json:"DeviceSerialNumber"`
		DeviceType                string `json:"DeviceType"`
		DeviceUniqueID            string `json:"DeviceUniqueID"`
		DeviceUniqueIDDisplayName string `json:"DeviceUniqueID_DisplayName"`
		Events                    []interface{}  `json:"Events"`
		Item                      struct {
			ItemID      int    `json:"ItemId"`
			DateCreated string `json:"DateCreated"`
			DateUpdated string `json:"DateUpdated"`
		} `json:"Item"`
		LastEvent struct {
			StartDateTime string `json:"StartDateTime"`
			SystemName    string `json:"SystemName"`
		} `json:"LastEvent"`
		LastEventIndex    int       `json:"LastEventIndex"`
		LOCMessageContent string    `json:"LOCMessageContent"`
		MessageID         string    `json:"MessageId"`
		MessageType       string    `json:"MessageType"`
		ReceivedTimestamp string `json:"ReceivedTimestamp"`
		RTLSAddress struct {
		} `json:"RTLSAddress"`
		RTLSContact struct {
		} `json:"RTLSContact"`
		RTLSGeo struct {
		} `json:"RTLSGeo"`
		Rtlsgps struct {
		} `json:"RTLSGPS"`
		RTLSModel2D struct {
			IsValid        bool      `json:"IsValid"`
			PosDisplayName string    `json:"PosDisplayName"`
			PosMapID       int       `json:"PosMapID"`
			PosModelID     int       `json:"PosModelID"`
			PosX           float64   `json:"PosX"`
			PosY           float64   `json:"PosY"`
			PosZoneIDs     string    `json:"PosZoneIDs"`
			Timestamp      string `json:"Timestamp"`
		} `json:"RTLSModel2D"`
		Sensor struct {
		} `json:"Sensor"`
		SequenceNumber int `json:"SequenceNumber"`
		Status         struct {
			BatteryLevel1      float64   `json:"BatteryLevel1"`
			ChargerConnected   bool      `json:"ChargerConnected"`
			Data2              string    `json:"Data2"`
			DeviceReportReason int       `json:"DeviceReportReason"`
			IsValid            bool      `json:"IsValid"`
			Timestamp          string `json:"Timestamp"`
		} `json:"Status"`
		DateCreated string `json:"DateCreated"`
		DateUpdated string `json:"DateUpdated"`
	} `json:"DeviceReports"`
	ItemInfo struct {
		ItemID      int    `json:"ItemId"`
		DateCreated string `json:"DateCreated"`
		DateUpdated string `json:"DateUpdated"`
	} `json:"ItemInfo"`
	ProximityReports  []interface{}     `json:"ProximityReports"`
	ReceivedTimestamp string `json:"ReceivedTimestamp"`
	SchemaName        string    `json:"SchemaName"`
	SchemaVersion     string    `json:"SchemaVersion"`
}

type Zone struct {
	ZoneID   int    `json:"ZoneID"`
	ZoneName string `json:"ZoneName"`
	ZoneType string `json:"ZoneType"`
}