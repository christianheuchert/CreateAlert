# GetStaffByZone

This activity retrieves staff in a specific zone.

## Configuration

### Input:

| Name       | Type   | Description                                                                                                                            |
| :--------- | :----- | :------------------------------------------------------------------------------------------------------------------------------------- |
| IP         | string | IP address followed by port. e.g. "32.41.13.112:322"                                                                                   |
| CustomerId | string | Customer unique identifier. e.g. "43"                                                                                                  |
| Username   | string | Username used to log into Sofia. e.g. "adminUser"                                                                                      |
| Password   | string | Password used to log into Sofia. e.g. "adminPassword"                                                                                  |
| Zone       | string | Accepts either Zone name, Zone Id, Or Zone Obj. e.g. "Hallway 3", "234", OR {"ZoneID":5215,"ZoneName":"Audiotorium","ZoneType":"Open"} |

### Output:

| Name  | Type  | Description                           |
| :---- | :---- | :------------------------------------ |
| Staff | array | Returns an array of stringified staff |
