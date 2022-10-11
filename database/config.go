package database

import "fmt"

var (
	dbUsername = "postgres"
	dbPassword = "postgres"
	dbHost     = "localhost"
	dbDB       = "postgres"
	dbPort     = "5432"
	pgConnStr  = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbUsername, dbPassword, dbHost, dbPort, dbDB)
	// pgConnStr  = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	// 	dbHost, dbPort, dbUsername, dbTable, dbPassword)
)

// pgConnStr  = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
// dbUsername, dbPassword, dbHost, dbPort, dbDB)
