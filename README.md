# Store Xpert Message

This activity receives and Xpert Message and stores it for later use.
WIP: This activity is not complete. GeoLattitude, GeoLongitude, Temperature, and Humidity are always returned as empty strings.
An xpertMessage with those fields included is required to accruately include them into the parser.

## Configuration

### Input:

| Name         | Type   | Description                                       |
| :----------- | :----- | :------------------------------------------------ |
| XpertMessage | string | JSON object holding device realted status details |

### Output:

| Name               | Type   | Description                                                                                                |
| :----------------- | :----- | :--------------------------------------------------------------------------------------------------------- |
| DeviceMAC          | string | Unique identifier of the device sending the message. e.g. "C4:CB:6B:63:3F:6F"                              |
| Timestamp          | string | Date/time the message was established. e.g. "2023-08-11 20:00:18.323"                                      |
| DeviceLogId        | string | Unique identifier of the log of the device message. e.g. "123"                                             |
| StatusReportReason | string | Identification number associated with report cause. e.g. "5"                                               |
| BatteryLevel       | string | The current battery life of the device expressed as a percentage. e.g. "24"                                |
| Temperature        | string | The current temperature. e.g. "89"                                                                         |
| Humidity           | string | The current level of humidity. e.g. "17"                                                                   |
| MapId              | string | Unique identifier of the map. e.g. "234"                                                                   |
| X                  | string | X position on the map. e.g. "1123"                                                                         |
| Y                  | string | Y position on the map. e.g. "1234"                                                                         |
| Zone               | array  | Zone(s) associated with Xpert Message. e.g. "[{"ZoneID":5215,"ZoneName":"Audiotorium","ZoneType":"Open"}]" |
| GeoLattitude       | string | Coordinates of the latitude. e.g. "234.01"                                                                 |
| GeoLongitude       | string | Coordinates of the longitude. e.g. "234.34"                                                                |
| ItemId             | string | Unique identifier of the item. e.g. "35443"                                                                |
| DisplayName        | string | The name assigned to the device. e.g. "F Asset"                                                            |
