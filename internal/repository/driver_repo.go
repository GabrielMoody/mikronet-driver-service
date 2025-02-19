package repository

import (
	"context"
	"time"

	"github.com/GabrielMoody/mikronet-driver-service/internal/helper"
	"github.com/GabrielMoody/mikronet-driver-service/internal/model"
	"github.com/GabrielMoody/mikronet-driver-service/internal/pb"
	"gorm.io/gorm"
)

type DriverRepo interface {
	CreateDriver(c context.Context, data model.DriverDetails) (model.DriverDetails, error)
	GetAllDrivers(c context.Context, verified *pb.ReqDrivers) ([]model.DriverDetails, error)
	GetDriverDetails(c context.Context, id string) (model.DriverDetails, error)
	EditDriverDetails(c context.Context, user model.DriverDetails) (model.DriverDetails, error)
	DeleteDriver(c context.Context, id string) (model.DriverDetails, error)
	GetStatus(c context.Context, id string) (res interface{}, err error)
	SetStatus(c context.Context, status string, id string) (res interface{}, err error)
	SetVerified(c context.Context, data model.DriverDetails) (res model.DriverDetails, err error)
	GetTripHistories(c context.Context, id string) (res interface{}, err error)
	GetAllDriverLastSeen(c context.Context) (res []model.DriverDetails, err error)
	SetLastSeen(c context.Context, id string) (res *time.Time, err error)
}

type DriverRepoImpl struct {
	db *gorm.DB
}

func (a *DriverRepoImpl) GetAllDriverLastSeen(c context.Context) (res []model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Find(&res).Where("last_seen >= ?", time.Now().Add(-5*time.Minute)).Scan(&res).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return res, nil
}

func (a *DriverRepoImpl) SetLastSeen(c context.Context, id string) (res *time.Time, err error) {
	if err := a.db.WithContext(c).Model(&model.DriverDetails{}).Where("id = ?", id).Update("last_seen", time.Now()).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DriverRepoImpl) DeleteDriver(c context.Context, id string) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Delete(&res, "id = ?", id).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DriverRepoImpl) SetVerified(c context.Context, data model.DriverDetails) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Updates(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *DriverRepoImpl) CreateDriver(c context.Context, data model.DriverDetails) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *DriverRepoImpl) GetAllDrivers(c context.Context, verified *pb.ReqDrivers) (res []model.DriverDetails, err error) {
	switch v := verified.Verified.(type) {
	case *pb.ReqDrivers_IsVerified:
		if err := a.db.WithContext(c).Find(&res, "verified = ?", v.IsVerified).Error; err != nil {
			return res, helper.ErrDatabase
		}
	case *pb.ReqDrivers_NotVerified:
		if err := a.db.WithContext(c).Find(&res, "verified = ?", v.NotVerified).Error; err != nil {
			return res, helper.ErrDatabase
		}
	default:
		if err := a.db.WithContext(c).Find(&res).Error; err != nil {
			return res, helper.ErrDatabase
		}
	}

	return res, nil
}

func (a *DriverRepoImpl) GetDriverDetails(c context.Context, id string) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).First(&res, "id = ?", id).Error; err != nil {
		return res, helper.ErrNotFound
	}

	return res, nil
}

func (a *DriverRepoImpl) EditDriverDetails(c context.Context, driver model.DriverDetails) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Updates(&driver).Error; err != nil {
		return driver, helper.ErrDatabase
	}

	return driver, nil
}

func (a *DriverRepoImpl) GetTripHistories(c context.Context, id string) (res interface{}, err error) {
	row, err := a.db.WithContext(c).Table("trips").
		Select("trips.location, trips.destination, trips.trip_date, reviews.review, reviews.star").
		Joins("JOIN reviews ON trips.id = reviews.id").
		Where("trips.driver_id = ?", id).
		Rows()

	if err != nil {
		return nil, helper.ErrDatabase
	}

	defer row.Close()

	type data struct {
		Location    string
		Destination string
		TripDate    time.Time
		Review      string
		Star        int64
	}

	var d data
	var trip []data

	for row.Next() {
		_ = row.Scan(&d.Location, &d.Destination, &d.TripDate, &d.Review, &d.Star)
		trip = append(trip, d)
	}

	return trip, nil
}

func (a *DriverRepoImpl) GetStatus(c context.Context, id string) (res interface{}, err error) {
	var driver model.DriverDetails

	if err := a.db.WithContext(c).Select("status").First(&driver, "id = ?", id).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return driver.Status, nil
}

func (a *DriverRepoImpl) SetStatus(c context.Context, status string, id string) (res interface{}, err error) {
	if err := a.db.WithContext(c).Where("id = ?", id).Updates(model.DriverDetails{Status: status}).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return status, nil
}

func NewDriverRepo(db *gorm.DB) DriverRepo {
	return &DriverRepoImpl{
		db: db,
	}
}
