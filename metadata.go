package SendEmail

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	IP string  `md:"IP,required"`
	CustomerId string`md:"CustomerId,required"`
	Username string `md:"Username,required"`
	Password string `md:"Password,required"`
	UserEmailAddress string `md:"UserEmailAddress,required"`
	EmailSubject string `md:"EmailSubject,required"`
	EmailMessage string `md:"EmailMessage,required"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["IP"])
	i.IP = strVal

	strVal, _ = coerce.ToString(values["CustomerId"])
	i.CustomerId = strVal

	strVal, _ = coerce.ToString(values["Username"])
	i.Username = strVal

	strVal, _ = coerce.ToString(values["Password"])
	i.Password = strVal

	strVal, _ = coerce.ToString(values["UserEmailAddress"])
	i.UserEmailAddress = strVal

	strVal, _ = coerce.ToString(values["EmailSubject"])
	i.EmailSubject = strVal

	strVal, _ = coerce.ToString(values["EmailMessage"])
	i.EmailMessage = strVal
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"IP": i.IP,
		"CustomerId": i.CustomerId,
		"Username": i.Username,
		"Password": i.Password,
		"UserEmailAddress": i.UserEmailAddress,
		"EmailSubject": i.EmailSubject,
		"EmailMessage": i.EmailMessage,
	}
}
type Output struct {
	SentBoolean bool `md:"SentBoolean"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	boolVal, _ := coerce.ToBool(values["SentBoolean"])
	o.SentBoolean = boolVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"SentBoolean": o.SentBoolean,
	}
}

	
type SendEmailResponse struct {
	ElapsedTimeInMillseconds float64 `json:"ElapsedTimeInMillseconds"`
	ErrorMessage             string  `json:"ErrorMessage"`
	SuccessMessage           string  `json:"SuccessMessage"`
	HasError                 bool    `json:"HasError"`
	ID                       int     `json:"Id"`
}