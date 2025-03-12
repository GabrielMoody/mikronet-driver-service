package model

import (
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey;type:varchar(36)"`
	Email     string `gorm:"unique;type:varchar(255)"`
	Password  string
	Role      string `gorm:"type:enum('admin','user','driver')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DriverDetails struct {
	ID             string `gorm:"column:id;primaryKey;type:varchar(255)" json:"id"`
	User           User   `gorm:"foreignKey:ID;references:ID;"`
	Name           string `gorm:"column:name;type:varchar(255)" json:"name"`
	PhoneNumber    string `gorm:"column:phone_number;type:varchar(255)" json:"phone_number"`
	RouteID        *uint  `gorm:"column:route_id" json:"route_id"`
	Route          Route  `gorm:"foreignKey:RouteID;references:ID;"`
	LicenseNumber  string `gorm:"column:license_number;type:varchar(255)" json:"license_number" form:"license_number"`
	SIM            string `gorm:"column:sim;type:varchar(255)" json:"sim" form:"sim"`
	Status         string `gorm:"column:status;type:varchar(255)" json:"status" form:"status"`
	Verified       bool   `gorm:"column:verified;default:false" json:"verified" form:"verified"`
	AvailableSeats int    `gorm:"column:available_seats" json:"available_seats" form:"available_seats"`
	ProfilePicture string `gorm:"column:profile_picture" json:"profile_picture" form:"photo"`
}
type PassengerDetails struct {
	ID   string `gorm:"column:id;primaryKey;type:varchar(255)" json:"id"`
	User User   `gorm:"foreignKey:ID;references:ID;"`
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
}
type Admin struct {
	ID   string `gorm:"column:id;primaryKey;type:varchar(255)" json:"id"`
	User User   `gorm:"foreignKey:ID;references:ID;"`
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
}

type ResetPassword struct {
	ID     int    `gorm:"primaryKey"`
	UserID string `gorm:"unique;type:varchar(36)"`
	User   User   `gorm:"foreignKey:UserID"`
	Code   string `gorm:"type:varchar(255)"`
}

type BlockedAccount struct {
	ID        int    `gorm:"column:id;primaryKey" json:"id"`
	AccountID string `gorm:"column:account_id;unique;type:varchar(255)" json:"account_id"`
	User      User   `gorm:"foreignKey:AccountID;references:ID;"`
	Role      string `gorm:"column:role;type:varchar(255)" json:"role"`
}

type Route struct {
	ID        int    `gorm:"column:id;primaryKey" json:"id"`
	RouteName string `gorm:"column:route_name;type:varchar(255)" json:"route_name"`
}

type Review struct {
	ID        int              `gorm:"column:id;primaryKey" json:"id"`
	UserID    string           `gorm:"column:user_id;type:varchar(255)" json:"user_id"`
	DriverID  string           `gorm:"column:driver_id;type:varchar(255)" json:"driver_id"`
	Passenger PassengerDetails `gorm:"foreignKey:UserID;references:ID;"`
	Driver    DriverDetails    `gorm:"foreignKey:DriverID;references:ID;"`
	Comment   string           `gorm:"column:comment;type:varchar(255)" json:"comment"`
	Star      int              `gorm:"column:star" json:"star"`
}
