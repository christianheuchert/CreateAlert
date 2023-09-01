# GetStaffByDepartment

This activity retrieves staff for a specified department.

## Configuration

### Input:

| Name           | Type   | Description                                                                                                        |
| :------------- | :----- | :----------------------------------------------------------------------------------------------------------------- |
| IP             | string | IP address followed by port. e.g. "32.41.13.112:322"                                                               |
| CustomerId     | string | Customer unique identifier. e.g. "43"                                                                              |
| Username       | string | Username used to log into Sofia. e.g. "adminUser"                                                                  |
| Password       | string | Password used to log into Sofia. e.g. "adminPassword"                                                              |
| DepartmentItem | string | Accepts department identificaiton number, department name, or a department item. e.g. "3091", "Administration", or |

    `{
     	"CustomerId": 123,
     	"DateCreated": "2023-08-09T12:34:56Z",
     	"DateUpdated": "2023-08-09T13:45:21Z",
     	"Description": "Example department",
     	"EnableTenancy": true,
     	"Name": "Engineering",
     	"TenantId": "987654",
     	"ElapsedTimeInMillseconds": 1234.567,
     	"ErrorMessage": "Error occurred",
     	"SuccessMessage": "Department created successfully",
     	"HasError": false,
     	"Id": 3098
     }`                                                                                                      |

### Output:

| Name  | Type  | Description                        |
| :---- | :---- | :--------------------------------- |
| Staff | array | Returns array of stringified staff |
