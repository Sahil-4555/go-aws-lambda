package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sqlc/configs"
	"go-sqlc/sqlc"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type CustomerWithOrders struct {
	CustomerID    int32   `json:"customer_id"`
	CustomerName  string  `json:"customer_name"`
	ContactName   string  `json:"contact_name"`
	Address       string  `json:"address"`
	City          string  `json:"city"`
	PostalCode    string  `json:"postal_code"`
	Country       string  `json:"country"`
	OrderID       int32   `json:"order_id"`
	OrderDate     string  `json:"order_date"`
	OrderDetailID int32   `json:"order_detail_id"`
	ProductName   string  `json:"product_name"`
	Quantity      int32   `json:"quantity"`
	Price         float64 `json:"price"`
	TotalAmount   float64 `json:"total_amount"`
}

func HelloWorld() (*events.APIGatewayProxyResponse, error) {
	body := "Hello World"
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

func GetCustomerOrders(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	Id := request.PathParameters["id"]
	id, _ := strconv.Atoi(Id)

	db := configs.NewConnection().GetDB()
	queries := sqlc.New(db)

	fetchedAuthor, err := queries.OrdersAssocaitedWithCustomer(ctx, int32(id))
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	fmt.Println("----------> id", Id)
	var orders []CustomerWithOrders

	for _, entry := range fetchedAuthor {
		data := CustomerWithOrders{
			CustomerID:    entry.CustomerID,
			CustomerName:  entry.CustomerName.String,
			ContactName:   entry.ContactName.String,
			Address:       entry.Address.String,
			City:          entry.City.String,
			PostalCode:    entry.PostalCode.String,
			Country:       entry.Country.String,
			OrderID:       entry.OrderID.Int32,
			OrderDate:     entry.OrderDate.String,
			OrderDetailID: entry.OrderDetailID.Int32,
			ProductName:   entry.ProductName.String,
			Quantity:      entry.Quantity.Int32,
			Price:         entry.Price.Float64,
			TotalAmount:   entry.TotalAmount,
		}

		orders = append(orders, data)
	}

	body, _ := json.Marshal(orders)
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil
}

type Customer struct {
	Id           int32  `json:"id"`
	CustomerName string `json:"customer_name"`
	ContactName  string `json:"contact_name"`
	Address      string `json:"address"`
	City         string `json:"city"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
}

func CreateCustomer(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var req Customer
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	db := configs.NewConnection().GetDB()
	queries := sqlc.New(db)

	// create a customer
	result, err := queries.CreateCustomer(ctx, sqlc.CreateCustomerParams{
		CustomerName: sql.NullString{String: req.CustomerName, Valid: true},
		ContactName:  sql.NullString{String: req.ContactName, Valid: true},
		Address:      sql.NullString{String: req.Address, Valid: true},
		City:         sql.NullString{String: req.City, Valid: true},
		PostalCode:   sql.NullString{String: req.PostalCode, Valid: true},
		Country:      sql.NullString{String: req.Country, Valid: true},
	})

	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	id := result.LastInsertId

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       fmt.Sprintf("Customer Created Successfully With ID %v", id),
	}, nil
}

func UpdateCustomer(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var req Customer
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	db := configs.NewConnection().GetDB()
	queries := sqlc.New(db)

	_, err := queries.UpdateCustomer(ctx, sqlc.UpdateCustomerParams{
		ID:           req.Id,
		CustomerName: sql.NullString{String: req.CustomerName, Valid: true},
		ContactName:  sql.NullString{String: req.ContactName, Valid: true},
		Address:      sql.NullString{String: req.Address, Valid: true},
		City:         sql.NullString{String: req.City, Valid: true},
		PostalCode:   sql.NullString{String: req.PostalCode, Valid: true},
		Country:      sql.NullString{String: req.Country, Valid: true},
	})

	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       fmt.Sprintf("Customer Updated Successfully With ID %v", req.Id),
	}, nil
}

func HardDeleteCustomer(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	Id := request.PathParameters["id"]
	id, _ := strconv.Atoi(Id)

	db := configs.NewConnection().GetDB()
	queries := sqlc.New(db)

	err := queries.DeleteCustomer(ctx, int32(id))

	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       fmt.Sprintf("Customer Deleted Successfully With ID %v", id),
	}, nil
}

func SoftDeleteCustomer(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	Id := request.PathParameters["id"]
	id, _ := strconv.Atoi(Id)

	db := configs.NewConnection().GetDB()
	queries := sqlc.New(db)

	err := queries.SoftDeleteCustomer(ctx, sqlc.SoftDeleteCustomerParams{
		ID:        int32(id),
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       fmt.Sprintf("Customer Deleted Successfully With ID %v", id),
	}, nil
}
