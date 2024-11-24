package database

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/hareem7bilal/go-microservice/internal/dberrors"
	"github.com/hareem7bilal/go-microservice/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).Find(&vendors)
	return vendors, result.Error
}

func (c Client) AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	vendor.VendorID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return vendor, nil

}

func (c Client) GetVendorByID(ctx context.Context, ID string) (*models.Vendor, error) {
	vendor := &models.Vendor{}
	result := c.DB.WithContext(ctx).Where(&models.Vendor{VendorID: ID}).First(&vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "Vendor", ID: ID}
		}
		return nil, result.Error
	}
	return vendor, nil

}

func (c Client) UpdateVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	var updatedVendor models.Vendor

	result := c.DB.WithContext(ctx).
		Model(&models.Vendor{}).
		Clauses(clause.Returning{}).
		Where(&models.Vendor{VendorID: vendor.VendorID}).
		Updates(models.Vendor{
			Name:    vendor.Name,
			Contact: vendor.Contact,
			Email:   vendor.Email,
			Phone:   vendor.Phone,
			Address: vendor.Address,
		}).
		First(&updatedVendor)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "vendor", ID: vendor.VendorID}
		}
		return nil, result.Error
	}

	return &updatedVendor, nil
}

func (c Client) DeleteVendor(ctx context.Context, ID string) error {
	result := c.DB.WithContext(ctx).
		Where(&models.Vendor{VendorID: ID}).
		Delete(&models.Vendor{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &dberrors.NotFoundError{
			Entity: "vendor",
			ID:     ID,
		}
	}

	return nil
}
