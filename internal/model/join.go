package model

type Drivers struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	LicenseNumber  string `json:"license_number"`
	SIM            string `json:"sim"`
	Verified       bool   `json:"verified"`
	ProfilePicture string `json:"profile_picture"`
}
