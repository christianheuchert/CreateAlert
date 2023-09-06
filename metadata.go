package ParseUnknownXpertMessage

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	XpertMessageJSON string `md:"XpertMessageJSON,required"`
	DeviceMACTarget string `md:"DeviceMACTarget"`
	TimestampTarget string `md:"TimestampTarget"`
	DeviceLogIdTarget string `md:"DeviceLogIdTarget"`
	StatusReportReasonTarget string `md:"StatusReportReasonTarget"`
	BatteryLevelTarget string `md:"BatteryLevelTarget"`
	TemperatureTarget string `md:"TemperatureTarget"`
	HumidityTarget string `md:"HumidityTarget"`
	MapIdTarget string `md:"MapIdTarget"`
	XTarget string `md:"XTarget"`
	YTarget string `md:"YTarget"`
	ZoneTarget string `md:"ZoneTarget"`
	GeoLattitudeTarget string `md:"GeoLattitudeTarget"`
	GeoLongitudeTarget string `md:"GeoLongitudeTarget"`
	ItemIdTarget string `md:"ItemIdTarget"`
	DisplayNameTarget string `md:"DisplayNameTarget"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["XpertMessageJSON"])
	r.XpertMessageJSON = strVal
	strVal, _ = coerce.ToString(values["DeviceMACTarget"])
	r.DeviceMACTarget = strVal
	strVal, _ = coerce.ToString(values["TimestampTarget"])
	r.TimestampTarget = strVal
	strVal, _ = coerce.ToString(values["DeviceLogIdTarget"])
	r.DeviceLogIdTarget = strVal
	strVal, _ = coerce.ToString(values["StatusReportReasonTarget"])
	r.StatusReportReasonTarget = strVal
	strVal, _ = coerce.ToString(values["BatteryLevelTarget"])
	r.BatteryLevelTarget = strVal
	strVal, _ = coerce.ToString(values["TemperatureTarget"])
	r.TemperatureTarget = strVal
	strVal, _ = coerce.ToString(values["HumidityTarget"])
	r.HumidityTarget = strVal
	strVal, _ = coerce.ToString(values["MapIdTarget"])
	r.MapIdTarget = strVal
	strVal, _ = coerce.ToString(values["XTarget"])
	r.XTarget = strVal
	strVal, _ = coerce.ToString(values["YTarget"])
	r.YTarget = strVal
	strVal, _ = coerce.ToString(values["ZoneTarget"])
	r.ZoneTarget = strVal
	strVal, _ = coerce.ToString(values["GeoLattitudeTarget"])
	r.GeoLattitudeTarget = strVal
	strVal, _ = coerce.ToString(values["GeoLongitudeTarget"])
	r.GeoLongitudeTarget = strVal
	strVal, _ = coerce.ToString(values["ItemIdTarget"])
	r.ItemIdTarget = strVal
	strVal, _ = coerce.ToString(values["DisplayNameTarget"])
	r.DisplayNameTarget = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"XpertMessageJSON": r.XpertMessageJSON,
		"DeviceMACTarget": r.DeviceMACTarget,
		"TimestampTarget": r.TimestampTarget,
		"DeviceLogIdTarget": r.DeviceLogIdTarget,
		"StatusReportReasonTarget": r.StatusReportReasonTarget,
		"BatteryLevelTarget": r.BatteryLevelTarget,
		"TemperatureTarget": r.TemperatureTarget,
		"HumidityTarget": r.HumidityTarget,
		"MapIdTarget": r.MapIdTarget,
		"XTarget": r.XTarget,
		"YTarget": r.YTarget,
		"ZoneTarget": r.ZoneTarget,
		"GeoLattitudeTarget": r.GeoLattitudeTarget,
		"GeoLongitudeTarget": r.GeoLongitudeTarget,
		"ItemIdTarget": r.ItemIdTarget,
		"DisplayNameTarget": r.DisplayNameTarget,
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
	Zone string `md:"Zone"`
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
	strVal, _ = coerce.ToString(values["Zone"])
	o.Zone = strVal
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
