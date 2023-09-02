package getStaffByZone

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	IP string  `md:"IP,required"`
	CustomerId string`md:"CustomerId,required"`
	Username string `md:"Username,required"`
	Password string `md:"Password,required"`
	Zone string `md:"Zone,required"`
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

	strVal, _ = coerce.ToString(values["Zone"])
	i.Zone = strVal
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"IP": i.IP,
		"CustomerId": i.CustomerId,
		"Username": i.Username,
		"Password": i.Password,
		"Zone": i.Zone,
	}
}

type Output struct {
	Staff []string `md:"Staff"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	assetList, ok:= (values["Staff"]).([]string) // type assertion
	if ok {
		o.Staff = assetList
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Staff":    o.Staff,
	}
}

type Staff struct {
	Address                      string    `json:"Address"`
	AlarisStatus                 string    `json:"AlarisStatus"`
	AlertStatus                  string `json:"AlertStatus"`
	AssociatedDevices            []interface{}  `json:"AssociatedDevices"`
	BatteryLevel                 float64    `json:"BatteryLevel"`
	BedStatus                    string `json:"BedStatus"`
	CurrentBuildingID            int    `json:"CurrentBuildingID"`
	CurrentBuildingName          string `json:"CurrentBuildingName"`
	CurrentSiteName              string `json:"CurrentSiteName"`
	CurrentFloorName             string `json:"CurrentFloorName"`
	CurrentMapID                 int    `json:"CurrentMapId"`
	CurrentModelID               int    `json:"CurrentModelId"`
	CurrentSiteID                int    `json:"CurrentSiteID"`
	CurrentTimestamp             string `json:"CurrentTimestamp"`
	CurrentX                     float64    `json:"CurrentX"`
	CurrentY                     float64    `json:"CurrentY"`
	CurrentZones                 string `json:"CurrentZones"`
	DepartmentID                 int    `json:"DepartmentID"`
	DepartmentName               string `json:"DepartmentName"`
	DeviceID                     int    `json:"DeviceID"`
	DeviceLogID                  int    `json:"DeviceLogID"`
	DeviceName                   string `json:"DeviceName"`
	Email                        string `json:"Email"`
	EnableAlerts                 bool   `json:"EnableAlerts"`
	EnableHygiene                bool   `json:"EnableHygiene"`
	EnableSDCT                   bool   `json:"EnableSDCT"`
	EventCountAcknowledged       int    `json:"EventCountAcknowledged"`
	EventCountClosed             int    `json:"EventCountClosed"`
	EventCountNew                int    `json:"EventCountNew"`
	EventCountOpen               int    `json:"EventCountOpen"`
	FromLDAP                     bool   `json:"FromLDAP"`
	GroupID                      int    `json:"GroupID"`
	GroupName                    string `json:"GroupName"`
	HealthStatus                 string `json:"HealthStatus"`
	Icon                         string `json:"Icon"`
	ImageData                    string `json:"ImageData"`
	ImageType                    string `json:"ImageType"`
	IsTestMode                   bool   `json:"IsTestMode"`
	OldTamper                    bool   `json:"OldTamper"`
	OldMotion                    bool   `json:"OldMotion"`
	Latitude                     float64    `json:"Latitude"`
	Longitude                    float64    `json:"Longitude"`
	LocationUpdated              string `json:"LocationUpdated"`
	ModelName                    string `json:"ModelName"`
	OldBuildingID                int    `json:"OldBuildingID"`
	OldMapID                     int    `json:"OldMapId"`
	OldModelID                   int    `json:"OldModelId"`
	OldSiteID                    int    `json:"OldSiteID"`
	OldLocationUpdated           string `json:"OldLocationUpdated"`
	OldX                         float64    `json:"OldX"`
	OldY                         float64    `json:"OldY"`
	OldZones                     string `json:"OldZones"`
	PendingDepartmentDateUpdated string `json:"PendingDepartmentDateUpdated"`
	PendingDepartmentID          int    `json:"PendingDepartmentId"`
	PhoneNumber                  string `json:"PhoneNumber"`
	Portrait                     string `json:"Portrait"`
	StaffID                      string `json:"StaffID"`
	StaffSettings                string    `json:"StaffSettings"`
	Temperature                  string `json:"Temperature"`
	UseCases                     []string  `json:"UseCases"`
	MultiAssign                  bool   `json:"MultiAssign"`
	AssocItemID                  int    `json:"AssocItemID"`
	PendingDepartmentName        string    `json:"PendingDepartmentName"`
	AssocItemName                string    `json:"AssocItemName"`
	CustomerID                   int    `json:"CustomerId"`
	DateCreated                  string `json:"DateCreated"`
	DateUpdated                  string `json:"DateUpdated"`
	Description                  string `json:"Description"`
	EnableTenancy                bool   `json:"EnableTenancy"`
	Name                         string `json:"Name"`
	TenantID                     string `json:"TenantId"`
	ElapsedTimeInMillseconds     float64    `json:"ElapsedTimeInMillseconds"`
	ErrorMessage                 string `json:"ErrorMessage"`
	SuccessMessage               string `json:"SuccessMessage"`
	HasError                     bool   `json:"HasError"`
	ID                           int    `json:"Id"`
} 

type GetAllStaffReponse struct {
	List                     []Staff      `json:"List"` //
}

type Zones struct {
	ZoneID   int    `json:"ZoneID"`
	ZoneName string `json:"ZoneName"`
	ZoneType string `json:"ZoneType"`
}