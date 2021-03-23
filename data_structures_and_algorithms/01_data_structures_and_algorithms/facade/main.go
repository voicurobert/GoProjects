package main

import "fmt"

/*
Facade is used to abstract subsystem interfaces with a helper. The facade design pattern is used in scenarios when the number of interfaces
increases and the system gets complicated. Facade is an entry point to different subsystems, and it simplifies the dependencies between the systems.
The facade pattern provides an interface that hides the implementation details of the hidden code.

A loosely coupled principle can be realized with a facade pattern. You can use a facade to improve poorly designed APIs.
In SOA, a service facade can be used to incorporate changes to the contract and implementation.
The facade pattern is made up of the facade class, module classes, and a client:
	The facade delegates the requests from the client to the module classes.
	The facade class hides the complexities of the subsystem logic and rules. Module classes implement the behaviors and functionalities of the module subsystem.

The client invokes the facade method. The facade class functionality can be spread across multiple packages and assemblies.
For example, account, customer, and transaction are the classes that have account, customer, and transaction creation methods.
BranchManagerFacade can be used by the client to create an account, customer, and transaction:
*/

// Account struct
type Account struct {
	id          string
	accountType string
}

func (account *Account) create(accountType string) *Account {
	fmt.Println("create account with type")
	account.accountType = accountType
	return account
}

func (account *Account) getByID(id string) *Account {
	fmt.Println("get account by id")
	return account
}

func (account *Account) deleteByID(id string) {
	fmt.Println("delete account by id")
}

// Customer struct
type Customer struct {
	name string
	id   int
}

func (customer *Customer) create(name string) *Customer {
	fmt.Println("Creating customer with name")
	customer.name = name
	return customer
}

// Transaction struct
type Transaction struct {
	id            string
	amount        float32
	srcAccountID  string
	destAccountID string
}

func (t *Transaction) create(srcAccountID, destAccountID string, amount float32) *Transaction {
	fmt.Println("create transaction")
	t.srcAccountID = srcAccountID
	t.destAccountID = destAccountID
	t.amount = amount
	return t
}

// BranchManagerFacade struct
type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

// NewBranchManagerFacade new NewBranchMasterFacade instance
func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.create(accountType)
	return customer, account
}

func (facade *BranchManagerFacade) createTransaction(srcAccountID, destAccountID string, amount float32) *Transaction {
	var transaction = facade.transaction.create(srcAccountID, destAccountID, amount)
	return transaction
}

func main() {
	var facade = NewBranchManagerFacade()
	var customer *Customer
	var account *Account
	customer, account = facade.createCustomerAccount("Thomas Smith", "Savings")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	var transaction = facade.createTransaction("21456", "87345", 1000)
	fmt.Println(transaction.amount)
}
