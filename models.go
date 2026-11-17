package main

type Driver struct {
	DriverID    string `json:"driverid"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Nationality string `json:"nationality"`
	Number      string `json:"number"`
	Code        string `json:"code"`
}

type APIResponse struct {
	Season       string   `json:"season"`
	TotalDrivers int      `json:"total_drivers"`
	Drivers      []Driver `json:"drivers"`
}
