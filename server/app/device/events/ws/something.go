package ws

type Something struct { // TODO: init dynamic data
	Id               string `json:"id"`
	Username         string `json:"username"`
	DeviceName       string `json:"deviceName"`
	SyncDeviceJoined string `json:"syncDeviceJoined"`
	CurrentDevice    Device `json:"currentDevice"`
	LastModifiedAt   int    `json:"lastModifiedAt"`
	AppUsername      string `json:"appUsername"`
	AppEmailAddress  string `json:"appEmailAddress"`
	AppAlignedCb     bool   `json:"appAlignedCb"`
	AppBillingPeriod int    `json:"appBillingPeriod"`
	AppSalary        int    `json:"appSalary"`
}

type Device struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
