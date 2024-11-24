package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/hareem7bilal/go-microservice/internal/dberrors"
	"github.com/hareem7bilal/go-microservice/internal/models"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).Where(models.Customer{Email: emailAddress}).Find(&customers)
	return customers, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return customer, nil

}

func (c Client) GetCustomerByID(ctx context.Context, ID string) (*models.Customer, error) {
	customer := &models.Customer{}
	result := c.DB.WithContext(ctx).Where(&models.Customer{CustomerID: ID}).First(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "customer", ID: ID}
		}
		return nil, result.Error
	}
	return customer, nil

}

func (c Client) UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	// Create a customer variable to store the updated record
	var updatedCustomer models.Customer

	result := c.DB.WithContext(ctx).
		Model(&models.Customer{}). // Specify the model
		Clauses(clause.Returning{}).
		Where(&models.Customer{CustomerID: customer.CustomerID}).
		Updates(models.Customer{
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Phone:     customer.Phone,
			Address:   customer.Address,
		}).
		First(&updatedCustomer) // Retrieve the updated record

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "customer", ID: customer.CustomerID}
		}
		return nil, result.Error
	}

	return &updatedCustomer, nil
}

func (c Client) DeleteCustomer(ctx context.Context, ID string) error {
    result := c.DB.WithContext(ctx).
        Where(&models.Customer{CustomerID: ID}).
        Delete(&models.Customer{})
    
    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return &dberrors.NotFoundError{
            Entity: "customer",
            ID: ID,
        }
    }
    
    return nil
}
