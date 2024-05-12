package main

import (
	"fmt"
	"go-sqlc/configs"
	"go-sqlc/routes"
)

func main() {

	configs.InitDB()

	// Uncomment the below function to feed the data in database
	// configs.MigrateData()

	fmt.Printf("Connected To Database Sucessfully on Port\n")
	routes.SetupLambda()
}
