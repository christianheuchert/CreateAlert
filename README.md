# SendEmail

This activity sends an email to a specified email address.

## Configuration

### Input:

| Name             | Type   | Description                                                 |
| :--------------- | :----- | :---------------------------------------------------------- |
| IP               | string | IP address followed by port. e.g. "32.41.13.112:322"        |
| CustomerId       | string | Customer unique identifier. e.g. "43"                       |
| Username         | string | Username used to log into Sofia. e.g. "adminUser"           |
| UserEmailAddress | string | Email of email recipient. e.g. "email@airista.com"          |
| EmailSubject     | string | Subject for sent email. e.g. "Email Subject"                |
| EmailMessage     | string | Content of sent message. e.g. "Hello You, Here is an email" |

### Output:

| Name        | Type    | Description                     |
| :---------- | :------ | :------------------------------ |
| SentBoolean | boolean | True if email sent successfully |
