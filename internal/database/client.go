package database

import (
	"context"
	"fmt"
	"github.com/hareem7bilal/go-microservice/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DatabaseClient interface {
	Ready() bool
	GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error)
	GetAllProducts(ctx context.Context, vendorId string) ([]models.Product, error)
	GetAllServices(ctx context.Context) ([]models.Service, error)
	GetAllVendors(ctx context.Context) ([]models.Vendor, error)
	AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	AddService(ctx context.Context, service *models.Service) (*models.Service, error)
	AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error)
	GetCustomerByID(ctx context.Context, ID string) (*models.Customer, error)
	GetProductByID(ctx context.Context, ID string) (*models.Product, error)
	GetServiceByID(ctx context.Context, ID string) (*models.Service, error)
	GetVendorByID(ctx context.Context, ID string) (*models.Vendor, error)
	UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	UpdateVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error)
	UpdateService(ctx context.Context, service *models.Service) (*models.Service, error)
	UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	DeleteCustomer(ctx context.Context, ID string) error
	DeleteProduct(ctx context.Context, ID string) error
	DeleteVendor(ctx context.Context, ID string) error
	DeleteService(ctx context.Context, ID string) error
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", "localhost", "postgres", "postgres", "postgres", 5435, "disable")
	// Open a connection to the PostgreSQL database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Configure naming strategy for the database tables
		NamingStrategy: schema.NamingStrategy{
			// Set a prefix for all table names, in this case, "wisdom."
			// This means all table names will start with "wisdom." (e.g., "wisdom.users")
			TablePrefix: "wisdom.",
		},

		// Specify a custom function to get the current time
		NowFunc: func() time.Time {
			// Use UTC time instead of the local time for consistency
			return time.Now().UTC()
		},

		// Enable querying all fields explicitly
		// This ensures that GORM includes all fields in SQL queries, even if they are not used in the struct
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		DB: db,
	}
	return client, nil

}

func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}
