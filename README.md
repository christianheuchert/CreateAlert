# GetStaff

This activity retrieves all staff for a given customer

## Configuration

### Input:

| Name       | Type   | Description                                           |
| :--------- | :----- | :---------------------------------------------------- |
| IP         | string | IP address followed by port. e.g. "32.41.13.112:322"  |
| CustomerId | string | Customer unique identifier. e.g. "43"                 |
| Username   | string | Username used to log into Sofia. e.g. "adminUser"     |
| Password   | string | Password used to log into Sofia. e.g. "adminPassword" |

### Output:

| Name  | Type  | Description                                  |
| :---- | :---- | :------------------------------------------- |
| Staff | array | Returns an array of all staff for a customer |
