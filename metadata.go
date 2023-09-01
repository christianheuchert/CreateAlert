package getGroups

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	IP string  `md:"IP,required"`
	CustomerId string`md:"CustomerId,required"`
	Username string `md:"Username,required"`
	Password string `md:"Password,required"`
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
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"IP": i.IP,
		"CustomerId": i.CustomerId,
		"Username": i.Username,
		"Password": i.Password,
	}
}

type Output struct {
	Groups []string `md:"Groups"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	groupList, ok:= (values["Groups"]).([]string) // type assertion
	if ok {
		o.Groups = groupList
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Groups":    o.Groups,
	}
}

type Response struct {
	List 					 []Group `json:"List"`
}

type Group struct {
	    Icon                     string `json:"Icon"`
		MultiAssign              bool   `json:"MultiAssign"`
		CustomerID               int    `json:"CustomerId"`
		DateCreated              string `json:"DateCreated"`
		DateUpdated              string `json:"DateUpdated"`
		Description              string `json:"Description"`
		EnableTenancy            bool   `json:"EnableTenancy"`
		Name                     string `json:"Name"`
		TenantID                 string `json:"TenantId"`
		ElapsedTimeInMillseconds float64 `json:"ElapsedTimeInMillseconds"`
		ErrorMessage             string `json:"ErrorMessage"`
		SuccessMessage           string `json:"SuccessMessage"`
		HasError                 bool   `json:"HasError"`
		ID                       int    `json:"Id"`
}