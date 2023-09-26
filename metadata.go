package getDwellStatus

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	IP string  `md:"IP,required"`
	CustomerId string`md:"CustomerId,required"`
	Username string `md:"Username,required"`
	Password string `md:"Password,required"`
	MAC string `md:"MAC,required"`
	GracePeriod string `md:"GracePeriod,required"`
	ZoneItem string `md:"ZoneItem,required"`
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

	strVal, _ = coerce.ToString(values["MAC"])
	i.MAC = strVal

	strVal, _ = coerce.ToString(values["GracePeriod"])
	i.GracePeriod = strVal

	strVal, _ = coerce.ToString(values["ZoneItem"])
	i.ZoneItem = strVal

	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"IP": i.IP,
		"CustomerId": i.CustomerId,
		"Username": i.Username,
		"Password": i.Password,
		"MAC": i.MAC,
		"GracePeriod":i.GracePeriod,
		"ZoneItem": i.ZoneItem,
	}
}

type Output struct {
	DwellStatus bool `md:"DwellStatus"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	boolVal, _ := coerce.ToBool(values["DwellStatus"])
	o.DwellStatus = boolVal

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"DwellStatus":    o.DwellStatus,
	}
}

type Response struct {
	List 					 []Log `json:"List"`
}

type Log struct {
		BuildingID                int    `json:"BuildingId"`
		BuildingName              string `json:"BuildingName"`
		EventDisplayName          string `json:"EventDisplayName"`
		DataMessageID             int    `json:"DataMessageId"`
		DeviceID                  int    `json:"DeviceId"`
		EventID                   int    `json:"EventID"`
		DeviceType                string `json:"DeviceType"`
		DeviceUniqueID            string `json:"DeviceUniqueId"`
		DeviceUniqueIDDisplayName string `json:"DeviceUniqueId_DisplayName"`
		ItemID                    int    `json:"ItemId"`
		MapID                     int    `json:"MapId"`
		MapName                   string `json:"MapName"`
		Message                   string `json:"Message"`
		ModelID                   int    `json:"ModelId"`
		SiteID                    string `json:"SiteId"`
		SiteName                  string `json:"SiteName"`
		X                         float64    `json:"X"`
		Y                         float64    `json:"Y"`
		ZoneID                    string `json:"ZoneId"`
		ZoneName                  string `json:"ZoneName"`
		CustomerID                int    `json:"CustomerId"`
		DateCreated               string `json:"DateCreated"`
		DateUpdated               string `json:"DateUpdated"`
		Description               string `json:"Description"`
		EnableTenancy             bool   `json:"EnableTenancy"`
		Name                      string `json:"Name"`
		TenantID                  string `json:"TenantId"`
		ElapsedTimeInMillseconds  float64    `json:"ElapsedTimeInMillseconds"`
		ErrorMessage              string `json:"ErrorMessage"`
		SuccessMessage            string `json:"SuccessMessage"`
		HasError                  bool   `json:"HasError"`
		ID                        int    `json:"Id"`
}

type Staff struct {
	ID                       int    `json:"Id"`
}

type Zone struct {
	ZoneID   int    `json:"ZoneID"`
	ZoneName string `json:"ZoneName"`
	ZoneType string `json:"ZoneType"`
}
type Device struct {
	ItemID                   int       `json:"ItemId"`
}