# GetGroups

This activity retrieves all Groups for a given customer.

## Configuration

### Input:

| Name       | Type   | Description                                           |
| :--------- | :----- | :---------------------------------------------------- |
| IP         | string | IP address followed by port. e.g. "32.41.13.112:322"  |
| CustomerId | string | Customer unique identifier. e.g. "43"                 |
| Username   | string | Username used to log into Sofia. e.g. "adminUser"     |
| Password   | string | Password used to log into Sofia. e.g. "adminPassword" |

### Output:

| Name   | Type  | Description                               |
| :----- | :---- | :---------------------------------------- |
| Groups | array | Returns an array of the customer's groups |
