# GetStaffByGroup

This activity retrieves staff in the specified group

## Configuration

### Input:

| Name       | Type   | Description                                                                                  |
| :--------- | :----- | :------------------------------------------------------------------------------------------- |
| IP         | string | IP address followed by port. e.g. "32.41.13.112:322"                                         |
| CustomerId | string | Customer unique identifier. e.g. "43"                                                        |
| Username   | string | Username used to log into Sofia. e.g. "adminUser"                                            |
| Password   | string | Password used to log into Sofia. e.g. "adminPassword"                                        |
| GroupItem  | string | Accepts group identificaiton number, group name, or a group item. e.g. "3091", "Surgery", or |

    `{
      "Icon": "assets\\icons\\792794ff-17db-46bc-b1fd-8110cf81f061.png",
      "MultiAssign": false,
      "CustomerId": 0,
      "DateCreated": "2023-06-22T16:40:25.43",
      "DateUpdated": "2023-06-22T16:40:25.43",
      "Description": "",
      "EnableTenancy": false,
      "Name": "Acute Care",
      "TenantId": "",
      "ElapsedTimeInMillseconds": 0,
      "ErrorMessage": "",
      "SuccessMessage": "",
      "HasError": false,
      "Id": 5644
    }`

|

### Output:

| Name  | Type  | Description                        |
| :---- | :---- | :--------------------------------- |
| Staff | array | Returns array of stringified staff |
