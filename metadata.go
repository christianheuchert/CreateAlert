package getDepartments

import (
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
	Users []string `md:"Users"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	userList, ok:= (values["Users"]).([]string) // type assertion
	if ok {
		o.Users = userList
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Users": o.Users,
	}
}

type GetAllUsersResponse struct {
	List []User `json:"List"`
}

type User struct {
	AccountExpirationDate    string `json:"AccountExpirationDate"`
	AccountLockedOut         bool   `json:"AccountLockedOut"`
	AccountLockedOutDateTime string `json:"AccountLockedOutDateTime"`
	Address1                 string `json:"Address1"`
	Address2                 string `json:"Address2"`
	AssociatedDepts          []struct {
		CustomerID               int    `json:"CustomerId"`
		DateCreated              string `json:"DateCreated"`
		DateUpdated              string `json:"DateUpdated"`
		Description              string `json:"Description"`
		EnableTenancy            bool   `json:"EnableTenancy"`
		Name                     string `json:"Name"`
		TenantID                 string `json:"TenantId"`
		ElapsedTimeInMillseconds float64    `json:"ElapsedTimeInMillseconds"`
		ErrorMessage             string `json:"ErrorMessage"`
		SuccessMessage           string `json:"SuccessMessage"`
		HasError                 bool   `json:"HasError"`
		ID                       int    `json:"Id"`
	} `json:"AssociatedDepts"`
	AssociatedGroups []struct {
		Icon                     string `json:"Icon"`
		MultiAssign              bool   `json:"MultiAssign"`
		CustomerID               int    `json:"CustomerId"`
		DateCreated              string `json:"DateCreated"`
		DateUpdated              string `json:"DateUpdated"`
		Description              string `json:"Description"`
		EnableTenancy            bool   `json:"EnableTenancy"`
		Name                     string `json:"Name"`
		TenantID                 string `json:"TenantId"`
		ElapsedTimeInMillseconds float64    `json:"ElapsedTimeInMillseconds"`
		ErrorMessage             string `json:"ErrorMessage"`
		SuccessMessage           string `json:"SuccessMessage"`
		HasError                 bool   `json:"HasError"`
		ID                       int    `json:"Id"`
	} `json:"AssociatedGroups"`
	AuditEnabled               bool   `json:"AuditEnabled"`
	AuthorizationCode          string `json:"AuthorizationCode"`
	City                       string `json:"City"`
	CurrentPassword            string `json:"CurrentPassword"`
	Domain                     string `json:"Domain"`
	Email                      string `json:"Email"`
	EnableAccess               bool   `json:"EnableAccess"`
	EnableAccountExpiration    bool   `json:"EnableAccountExpiration"`
	EnableNotifications        bool   `json:"EnableNotifications"`
	Fax                        string `json:"Fax"`
	FirstLoginDateTime         string `json:"FirstLoginDateTime"`
	FirstName                  string `json:"FirstName"`
	FullName                   string `json:"FullName"`
	FromLDAP                   bool   `json:"FromLDAP"`
	IsDeleted                  bool   `json:"IsDeleted"`
	IsProtectedUser            bool   `json:"IsProtectedUser"`
	IsSuperUser                bool   `json:"IsSuperUser"`
	LanguageID                 int    `json:"LanguageId"`
	LastLoginDateTime          string `json:"LastLoginDateTime"`
	LastAgreedToEULADateTime   string `json:"LastAgreedToEULADateTime"`
	LastName                   string `json:"LastName"`
	Password                   string `json:"Password"`
	PasswordExpirationDateTime string `json:"PasswordExpirationDateTime"`
	PasswordHash               string `json:"PasswordHash"`
	Phone                      string `json:"Phone"`
	PreferencesXML             string `json:"PreferencesXML"`
	ResetPasswordOnLogin       bool   `json:"ResetPasswordOnLogin"`
	State                      string `json:"State"`
	Status                     string `json:"Status"`
	TimeZoneID                 int    `json:"TimeZoneId"`
	UserDefine1                string `json:"UserDefine1"`
	UserDefine2                string `json:"UserDefine2"`
	UserDefine4                string `json:"UserDefine4"`
	UserDefine5                string `json:"UserDefine5"`
	UserName                   string `json:"UserName"`
	UserSettings               string `json:"UserSettings"`
	UserXML                    string `json:"UserXML"`
	ZipCode                    string `json:"ZipCode"`
	RoleName                   string `json:"RoleName"`
	RolePermissions            string `json:"RolePermissions"`
	RoleID                     int    `json:"RoleId"`
	Permissions                []interface{}  `json:"Permissions"`
	CustomerID                 int    `json:"CustomerId"`
	DateCreated                string `json:"DateCreated"`
	DateUpdated                string `json:"DateUpdated"`
	Description                string `json:"Description"`
	EnableTenancy              bool   `json:"EnableTenancy"`
	Name                       string `json:"Name"`
	TenantID                   string `json:"TenantId"`
	ElapsedTimeInMillseconds   float64    `json:"ElapsedTimeInMillseconds"`
	ErrorMessage               string `json:"ErrorMessage"`
	SuccessMessage             string `json:"SuccessMessage"`
	HasError                   bool   `json:"HasError"`
	ID                         int    `json:"Id"`
}