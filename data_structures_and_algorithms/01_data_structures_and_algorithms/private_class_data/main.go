package main

import (
	"encoding/json"
	"fmt"
)

/*
The private class data pattern secures the data within a class. This pattern encapsulates the initialization of the class data.
The write privileges of properties within the private class are protected, and properties are set during construction.
The private class pattern prints the exposure of information by securing it in a class that retains the state.
The encapsulation of class data initialization is a scenario where this pattern is applicable.

Account is a class with account details and a customer name.
AccountDetails is the private attribute of Account , and CustomerName is the public attribute.
JSON marshaling of Account has CustomerName as a public property.
AccountDetails is the package property in Go (modeled as private class data):
*/

// AccountDetails struct
type AccountDetails struct {
	id          string
	accountType string
}

// Account struct
type Account struct {
	details      *AccountDetails
	CustomerName string
}

// Account class method setDetails
func (account *Account) setDetails(id string, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

//Account class method getId
func (account *Account) getID() string {
	return account.details.id
}

//Account class method getAccountType
func (account *Account) getAccountType() string {
	return account.details.accountType
}

func main() {
	var account *Account = &Account{CustomerName: "John Smith"}
	account.setDetails("4532", "current")
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private Class hidden", string(jsonAccount))
	fmt.Println("Account Id", account.getID())
	fmt.Println("Account Type", account.getAccountType())
}
