# GetStaffByBuilding

This activity retrieves staff in a given building.

## Configuration

### Input:

| Name       | Type   | Description                                                  |
| :--------- | :----- | :----------------------------------------------------------- |
| IP         | string | IP address followed by port. e.g. "32.41.13.112:322"         |
| CustomerId | string | Customer unique identifier. e.g. "43"                        |
| Username   | string | Username used to log into Sofia. e.g. "adminUser"            |
| Password   | string | Password used to log into Sofia. e.g. "adminPassword"        |
| IP         | string | IP address followed by port. e.g. "32.41.13.112:322"         |
| Building   | string | Building to check for Staff. e.g. "Library Room 2", OR "416" |

### Output:

| Name  | Type  | Description                           |
| :---- | :---- | :------------------------------------ |
| Staff | array | Returns an array of stringified staff |
