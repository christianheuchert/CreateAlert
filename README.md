# Store Xpert Message

This activity takes a user provided JSON object and dot notation keys for said object. Then, the Xpert Message items are returned.

- NOTE: This activity may not always find values. In the example below, many of the outputs will be empty.

```
    XpertMessageJSON = {"MAC":"C4:CB:6B:22:1B:DA","pos":{"mapID":"14741","x":1168,"y":565},"status":{"rpt":0,"sfsw":0,"chg":0,"batt":40,"sqn":0}}
    DeviceMACTarget = "XpertMessage.MAC"
    MapIdTarget = "XpertMessage.pos.mapID"
    XTarget = "XpertMessage.pos.x"
    YTarget = "XpertMessage.pos.y"
    BatteryLevelTarget = "XpertMessage.status.batt"

    Output:
    DeviceMAC = C4:CB:6B:22:1B:DA
    Timestamp = ""
    DeviceLogId = ""
    StatusReportReason = ""
    BatteryLevel =   "40"
    Temperature = ""
    Humidity = ""
    MapId = 14741
    X = 1168
    Y = 565
    Zone = ""
    GeoLattitude = ""
    GeoLongitude = ""
    ItemId = ""
    DisplayName = ""
```

## Configuration

### Input:

| Name                     | Type   | Description                                                                                                         |
| :----------------------- | :----- | :------------------------------------------------------------------------------------------------------------------ |
| XpertMessageJSON         | string | JSON object holding device related status details. e.g. "{}"                                                        |
| DeviceMACTarget          | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.DeviceReports[0].DeviceUniqueID"            |
| TimestampTarget          | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.DeviceReports[0].TimeStamp"                 |
| DeviceLogIdTarget        | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.DeviceReports[0].DeviceLogId"               |
| StatusReportReasonTarget | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.DeviceReports[0].Status.DeviceReportReason" |
| BatteryLevelTarget       | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.DeviceReports[0].Status.BatteryLevel"       |
| TemperatureTarget        | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.DeviceReports[0].Therm.Temperature"         |
| HumidityTarget           | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.DeviceReports[0].Them.Humidity"             |
| MapIdTarget              | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.Data.MapId"                                 |
| XTarget                  | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.Location.X"                                 |
| YTarget                  | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.Location.Y"                                 |
| ZoneTarget               | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.Location.Zone"                              |
| GeoLattitudeTarget       | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.Location.Lattitude"                         |
| GeoLongitudeTarget       | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.Location.Longitude"                         |
| ItemIdTarget             | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.DeviceReports[0].Itme.ItemId"               |
| DisplayNameTarget        | string | Dot notation indicating where to search JSON object. e.g. "XpertMessage.DeviceReports[0].RTLSModel2D.DisplayName"   |

### Output:

| Name               | Type   | Description                                                                   |
| :----------------- | :----- | :---------------------------------------------------------------------------- |
| DeviceMAC          | string | Unique identifier of the device sending the message. e.g. "C4:CB:6B:63:3F:6F" |
| Timestamp          | string | Date/time the message was established. e.g. "2023-08-11 20:00:18.323"         |
| DeviceLogId        | string | Unique identifier of the log of the device message. e.g. "123"                |
| StatusReportReason | string | Identification number associated with report cuase. e.g. "5"                  |
| BatteryLevel       | string | The current battery life of the device expressed as a percentage. e.g. "24"   |
| Temperature        | string | The current temperature. e.g. "89"                                            |
| Humidity           | string | The current level of humidity. e.g. "17"                                      |
| MapId              | string | Unique identifier of the map. e.g. "234"                                      |
| X                  | string | X position on the map. e.g. "1123"                                            |
| Y                  | string | Y position on the map. e.g. "1234"                                            |
| Zone               | string | Zone(s) associated with Xpert Message. e.g. "5215" "Audiotorium"              |
| GeoLattitude       | string | Coordinates of the latitude. e.g. "234.01"                                    |
| GeoLongitude       | string | Coordinates of the longitude. e.g. "234.34"                                   |
| ItemId             | string | Unique identifier of the item. e.g. "35443"                                   |
| DisplayName        | string | The name assigned to the device. e.g. "F Asset"                               |
