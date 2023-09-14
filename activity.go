package getDwellStatus

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	dwellStatus := RestCallGetDwellTime(input.IP, input.CustomerId, input.Username, input.Password, input.StaffItem, input.ZoneItem)

	output := &Output{DwellStatus: dwellStatus}

	// fmt.Println("Output: ", output.DwellStatus)
	// ctx.Logger().Info("Output: ", output)

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

//http://52.45.17.177:802/XpertRestApi/api/MetaData/GetGroups?CustomerId=1
func RestCallGetDwellTime(IP string, CustomerId string, username string, password string, StaffId string, ZoneItem string) bool {

	// Create an HTTP client
	client := &http.Client{}

	var staffObj Staff // check if Staff Object
	objectCheck := json.Unmarshal([]byte(StaffId), &staffObj) 
	if (objectCheck == nil){ // StaffId is object
		StaffId = strconv.Itoa(staffObj.ID)
	}// else already Integer string
	
	//Get Times and url escape them
	StartDateTime, EndDateTime := getCurrentAndPastTime()
		// //TESTING TIMES
		// StartDateTime = "2023-07-27 16:47:48.807"
		// EndDateTime = "2023-07-27 16:57:48.807"
		// //TESTING TIMES
	StartDateTime = url.QueryEscape(StartDateTime)
	EndDateTime = url.QueryEscape(EndDateTime)

	// Create the request
	// url := "http://"+IP+"/XpertRestApi/api/Event/GetGroups?CustomerId="+customerId
	url := "http://" + IP + "/XpertRestApi/api/DeviceLogs/GetByStaffId?StartDateTime="+StartDateTime+"&EndDateTime="+EndDateTime+"&StaffId="+StaffId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	// Add basic authentication to the request header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return false
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	//Declare response struct object
	var response Response
	// Unmarshal the config JSON into Response struct
	errUnmarshal := json.Unmarshal([]byte(body), &response)
	if errUnmarshal != nil {
	 	fmt.Println("response error: ", errUnmarshal)
		return false
	}

	dwellStatus := checkLogsForZoneTarget(ZoneItem, response.List)

	return dwellStatus
}

func getCurrentAndPastTime() (string, string) {
    currentTime := time.Now()
    pastTime := currentTime.Add(-10 * time.Minute)

    EndDateTime := currentTime.Format("2006-01-02 15:04:05.000")
    StartDateTime := pastTime.Format("2006-01-02 15:04:05.000")

    return StartDateTime, EndDateTime
}

func checkLogsForZoneTarget(ZoneItem string, Logs []Log) bool {
	ZoneId, _ := strconv.Atoi(ZoneItem) // check if zone Int Id, if not, set to -1
	if (ZoneId == 0){ZoneId = -1}
	var zoneObj Zone // check if Zone Object
	zoneCheck := json.Unmarshal([]byte(ZoneItem), &zoneObj) 
	if (zoneCheck == nil){ // Zone is object
		ZoneId = (zoneObj.ZoneID)
	}

	for _, element := range Logs {
		// Access ZoneID data
		var zones []Zone
		if err := json.Unmarshal([]byte(element.ZoneID), &zones); err != nil {
			fmt.Println("Zone Error:", err)
		}

		// check ZoneId OR ZoneName
		if zones[0].ZoneID == ZoneId || zones[0].ZoneName == ZoneItem {
			return true
		}
	}

	return false
}