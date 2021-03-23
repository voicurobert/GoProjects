package main

import (
	"database/sql"
)

// Customer struct
type Customer struct {
	CustomerID   int
	CustomerName string
	SSN          string
}

// GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "newuser"
	databaseName := "crm"
	database, err := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	return database
}

// GetCustomers method returns Customer Array
func GetCustomers() []Customer {
	var database *sql.DB = GetConnection()
	defer database.Close()
	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}
	var customer Customer = Customer{}
	var customers []Customer
	for rows.Next() {
		var customerID int
		var customerName string
		var ssn string
		error = rows.Scan(&customerID, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerID = customerID
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	return customers
}

// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	var database *sql.DB = GetConnection()
	defer database.Close()
	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO CUSTOMER(CustomerName,SSN) VALUES(?,?)")
	if error != nil {
		panic(error.Error())
	}
	defer insert.Close()
	insert.Exec(customer.CustomerName, customer.SSN)
}

// UpdateCustomer func
func UpdateCustomer(customer Customer) {
	var database *sql.DB = GetConnection()
	defer database.Close()
	update, error := database.Prepare("UPDATE CUSTOMER SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	defer update.Close()
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerID)
}

// Delete Customer method with parameter customer
func deleteCustomer(customer Customer) {
	var database *sql.DB = GetConnection()
	defer database.Close()
	delete, error := database.Prepare("DELETE FROM Customer WHERE Customerid=?")
	if error != nil {
		panic(error.Error())
	}
	defer delete.Close()
	delete.Exec(customer.CustomerID)
}

func main() {
	//var customers []Customer = GetCustomers()
	//fmt.Println("Customers", customers)
}
