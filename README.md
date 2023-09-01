# CreateAlert

This activity creates an alert on Sofia

## Configuration

### Input:

| Name             | Type   | Description                                           |
| :--------------- | :----- | :---------------------------------------------------- |
| IP               | string | IP address followed by port. e.g. "32.41.13.112:322"  |
| CustomerId       | string | Customer unique identifier. e.g. "43"                 |
| Username         | string | Username used to log into Sofia. e.g. "adminUser"     |
| Password         | string | Password used to log into Sofia. e.g. "adminPassword" |
| AlertDisplayName | string | Alert name to appear in Sofia                         |

### Output:

| Name         | Type    | Description                        |
| :----------- | :------ | :--------------------------------- |
| AlertBoolean | boolean | True if alert is successfully sent |
