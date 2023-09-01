package CreateAlert

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	IP         string `md:"IP,required"`
	CustomerId string `md:"CustomerId,required"`
	Username   string `md:"Username,required"`
	Password   string `md:"Password,required"`
	AlertDisplayName string `md:"AlertDisplayName,required"`
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

	strVal, _ = coerce.ToString(values["AlertDisplayName"])
	i.AlertDisplayName = strVal
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"IP":         i.IP,
		"CustomerId": i.CustomerId,
		"Username":   i.Username,
		"Password":   i.Password,
		"AlertDisplayName": i.AlertDisplayName,
	}
}


type Output struct {
	AlertBoolean bool `md:"AlertBoolean"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToBool(values["AlertBoolean"])
	o.AlertBoolean = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"AlertBoolean": o.AlertBoolean,
	}
}

	
type EventModel struct {
	AllowedValueRange         string    `json:"AllowedValueRange"`
	Details                   string    `json:"Details"`
	DeviceID                  int       `json:"DeviceId"`
	DeviceLogID               int       `json:"DeviceLogId"`
	DisplayName               string    `json:"DisplayName"`
	EndDateTime               string `json:"EndDateTime"`
	IsAcknowledgementRequired bool      `json:"IsAcknowledgementRequired"`
	ItemID                    int       `json:"ItemId"`
	MaxValue                  int       `json:"MaxValue"`
	MinValue                  int       `json:"MinValue"`
	PlanID                    int       `json:"PlanId"`
	RuleName                  string    `json:"RuleName"`
	RuleSetMajorVersion       int       `json:"RuleSetMajorVersion"`
	RuleSetMinorVersion       int       `json:"RuleSetMinorVersion"`
	RuleSetName               string    `json:"RuleSetName"`
	SeverityColor             string    `json:"SeverityColor"`
	SeverityIcon              string    `json:"SeverityIcon"`
	SiteID                    string    `json:"SiteId"`
	StartDateTime             string `json:"StartDateTime"`
	SystemName                string    `json:"SystemName"`
	Type                      string    `json:"Type"`
	UseCase                   int       `json:"UseCase"`
	ViolationValue            string    `json:"ViolationValue"`
	DateUpdated               string `json:"DateUpdated"`
	Description               string    `json:"Description"`
}

type InsertEventResponse struct {
	ElapsedTimeInMillseconds float64 `json:"ElapsedTimeInMillseconds"`
	ErrorMessage             string  `json:"ErrorMessage"`
	SuccessMessage           string  `json:"SuccessMessage"`
	HasError                 bool    `json:"HasError"`
	ID                       int     `json:"Id"`
}