# GetUsersByGroup

This activity retrieves users in a specified Group.

## Configuration

### Input:

| Name       | Type   | Description                                                 |
| :--------- | :----- | :---------------------------------------------------------- |
| IP         | string | IP address followed by port. e.g. "32.41.13.112:322"        |
| CustomerId | string | Customer unique identifier. e.g. "43"                       |
| Username   | string | Username used to log into Sofia. e.g. "adminUser"           |
| Password   | string | Password used to log into Sofia. e.g. "adminPassword"       |
| Group      | string | Group identificaiton number or name. e.g. "1063" OR "Sales" |

### Output:

| Name  | Type  | Description                           |
| :---- | :---- | :------------------------------------ |
| Users | array | Returns an array of stringified users |
