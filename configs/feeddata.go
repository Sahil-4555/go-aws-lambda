package configs

import (
	"context"
	"database/sql"
	"encoding/json"
	"go-sqlc/sqlc"
	"io/ioutil"
	"log"
	"time"
)

func MigrateData() {
	FeedCategoriesData()
	FeedCustomerData()
	FeedEmployeeData()
	FeedShippersData()
	FeedSuplliersData()
	FeedProductsData()
	FeedOrdersData()
	FeedOrderDetailsData()
}

// Customer struct to map the JSON data
type Customer struct {
	CustomerName string    `json:"customer_name"`
	ContactName  string    `json:"contact_name"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	PostalCode   string    `json:"postal_code"`
	Country      string    `json:"country"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type Categorie struct {
	CategoryName string    `json:"category_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type Employee struct {
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	BirthDate string    `json:"birth_date"`
	Photo     string    `json:"photo"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type OrderDetail struct {
	OrderID   int32     `json:"order_id"`
	ProductID int32     `json:"product_id"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Order struct {
	OrderID    uint      `json:"id"` // Primary Key
	CustomerID uint      `json:"customer_id"`
	EmployeeID uint      `json:"employee_id"`
	ShipperID  uint      `json:"shipper_id"`
	OrderDate  string    `json:"order_date"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type Product struct {
	ProductName string    `json:"product_name"`
	SupplierID  uint      `json:"supplier_id"`
	CategoryID  uint      `json:"category_id"`
	Unit        string    `json:"unit"`
	UnitInt     int       `json:"unit_int"`
	Price       float32   `json:"price"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type Shipper struct {
	ShipperName string    `json:"shipper_name"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type Supplier struct {
	SupplierName string    `json:"supplier_name"`
	ContactName  string    `json:"contact_name"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	PostalCode   string    `json:"postal_code"`
	Country      string    `json:"country"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

func FeedCategoriesData() {

	jsonData, err := ioutil.ReadFile("data/categories.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var categories []Categorie
	if err := json.Unmarshal(jsonData, &categories); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	ctx := context.Background()
	queries := sqlc.New(db)

	for _, categorie := range categories {
		_, err := queries.CreateCategorie(ctx, sqlc.CreateCategorieParams{
			CategoryName: sql.NullString{String: categorie.CategoryName, Valid: true},
			Description:  sql.NullString{String: categorie.Description, Valid: true},
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func FeedCustomerData() {

	jsonData, err := ioutil.ReadFile("data/customers.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var customers []Customer
	if err := json.Unmarshal(jsonData, &customers); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	ctx := context.Background()
	queries := sqlc.New(db)

	for _, customer := range customers {
		_, err := queries.CreateCustomer(ctx, sqlc.CreateCustomerParams{
			CustomerName: sql.NullString{String: customer.CustomerName, Valid: true},
			ContactName:  sql.NullString{String: customer.ContactName, Valid: true},
			Address:      sql.NullString{String: customer.Address, Valid: true},
			City:         sql.NullString{String: customer.City, Valid: true},
			PostalCode:   sql.NullString{String: customer.PostalCode, Valid: true},
			Country:      sql.NullString{String: customer.Country, Valid: true},
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func FeedEmployeeData() {

	jsonData, err := ioutil.ReadFile("data/employee.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var employees []Employee
	if err := json.Unmarshal(jsonData, &employees); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	ctx := context.Background()
	queries := sqlc.New(db)
	for _, employee := range employees {
		_, err := queries.CreateEmployee(ctx, sqlc.CreateEmployeeParams{
			LastName:  sql.NullString{String: employee.LastName, Valid: true},
			FirstName: sql.NullString{String: employee.LastName, Valid: true},
			BirthDate: sql.NullString{String: employee.BirthDate, Valid: true},
			Photo:     sql.NullString{String: employee.Photo, Valid: true},
			Notes:     sql.NullString{String: employee.Notes, Valid: true},
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func FeedOrderDetailsData() {

	jsonData, err := ioutil.ReadFile("data/orderdetails.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var orderdetails []OrderDetail
	if err := json.Unmarshal(jsonData, &orderdetails); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	ctx := context.Background()
	queries := sqlc.New(db)

	for _, orderdetail := range orderdetails {
		_, err := queries.CreateOrderDetail(ctx, sqlc.CreateOrderDetailParams{
			OrderID:   sql.NullInt32{Int32: orderdetail.OrderID, Valid: true},
			ProductID: sql.NullInt32{Int32: orderdetail.ProductID, Valid: true},
			Quantity:  sql.NullInt32{Int32: orderdetail.Quantity, Valid: true},
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func FeedOrdersData() {

	jsonData, err := ioutil.ReadFile("data/orders.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var orders []Order
	if err := json.Unmarshal(jsonData, &orders); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	ctx := context.Background()
	queries := sqlc.New(db)

	for _, order := range orders {
		_, err := queries.CreateOrder(ctx, sqlc.CreateOrderParams{
			ID:         int32(order.OrderID),
			CustomerID: sql.NullInt32{Int32: int32(order.CustomerID), Valid: true},
			EmployeeID: sql.NullInt32{Int32: int32(order.EmployeeID), Valid: true},
			ShipperID:  sql.NullInt32{Int32: int32(order.ShipperID), Valid: true},
			OrderDate:  sql.NullString{String: order.OrderDate, Valid: true},
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func FeedProductsData() {

	jsonData, err := ioutil.ReadFile("data/products.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var products []Product
	if err := json.Unmarshal(jsonData, &products); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	ctx := context.Background()
	queries := sqlc.New(db)

	for _, product := range products {
		_, err := queries.CreateProduct(ctx, sqlc.CreateProductParams{
			ProductName: sql.NullString{String: product.ProductName, Valid: true},
			SupplierID:  sql.NullInt32{Int32: int32(product.SupplierID), Valid: true},
			CategoryID:  sql.NullInt32{Int32: int32(product.CategoryID), Valid: true},
			Unit:        sql.NullString{String: product.Unit, Valid: true},
			Price:       sql.NullFloat64{Float64: float64(product.Price), Valid: true},
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func FeedShippersData() {

	jsonData, err := ioutil.ReadFile("data/shippers.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var shippers []Shipper
	if err := json.Unmarshal(jsonData, &shippers); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	ctx := context.Background()
	queries := sqlc.New(db)

	for _, shipper := range shippers {
		_, err := queries.CreateShipper(ctx, sqlc.CreateShipperParams{
			ShipperName: sql.NullString{String: shipper.ShipperName, Valid: true},
			Phone:       sql.NullString{String: shipper.Phone, Valid: true},
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func FeedSuplliersData() {

	jsonData, err := ioutil.ReadFile("data/suppliers.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var suppliers []Supplier
	if err := json.Unmarshal(jsonData, &suppliers); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	ctx := context.Background()
	queries := sqlc.New(db)

	for _, supplier := range suppliers {
		_, err := queries.CreateSupplier(ctx, sqlc.CreateSupplierParams{
			SupplierName: sql.NullString{String: supplier.SupplierName, Valid: true},
			ContactName:  sql.NullString{String: supplier.ContactName, Valid: true},
			Address:      sql.NullString{String: supplier.Address, Valid: true},
			City:         sql.NullString{String: supplier.City, Valid: true},
			PostalCode:   sql.NullString{String: supplier.PostalCode, Valid: true},
			Country:      sql.NullString{String: supplier.Country, Valid: true},
			Phone:        sql.NullString{String: supplier.Phone, Valid: true},
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}
