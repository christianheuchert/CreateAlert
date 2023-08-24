package getDepartments

import (
	"encoding/json"

	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	IP         string `md:"IP,required"`
	CustomerId string `md:"CustomerId,required"`
	Username   string `md:"Username,required"`
	Password   string `md:"Password,required"`
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
		"IP":         i.IP,
		"CustomerId": i.CustomerId,
		"Username":   i.Username,
		"Password":   i.Password,
	}
}

type Output struct {
	Users string `md:"Users"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["Users"])
	o.Users = strVal

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Users": o.Users,
	}
}

type GetAllUsersResponse struct {
	List json.RawMessage `json:"List"`
}

