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

func (c Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).Where(models.Product{VendorID: vendorId}).Find(&products)
	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return product, nil

}

func (c Client) GetProductByID(ctx context.Context, ID string) (*models.Product, error) {
	product := &models.Product{}
	result := c.DB.WithContext(ctx).Where(&models.Product{ProductID: ID}).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "Product", ID: ID}
		}
		return nil, result.Error
	}
	return product, nil

}

func (c Client) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	var updatedProduct models.Product

	result := c.DB.WithContext(ctx).
		Model(&models.Product{}).
		Clauses(clause.Returning{}).
		Where(&models.Product{ProductID: product.ProductID}).
		Updates(models.Product{
			VendorID: product.VendorID,
			Name:     product.Name,
			Price:    product.Price,
		}).
		First(&updatedProduct)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "product", ID: product.ProductID}
		}
		return nil, result.Error
	}

	return &updatedProduct, nil
}

func (c Client) DeleteProduct(ctx context.Context, ID string) error {
	result := c.DB.WithContext(ctx).
        Where(&models.Product{ProductID: ID}).
        Delete(&models.Product{})
    
    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return &dberrors.NotFoundError{
            Entity: "product",
            ID: ID,
        }
    }
    
    return nil
}
