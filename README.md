# ExpenseTrackerGoLanguage

This is a GoLang API that provides functionality for managing expenses.

## Getting Started

To get started with this API, follow the instructions below.

### Prerequisites

- Go (Install from [https://golang.org/doc/install](https://golang.org/doc/install))
- Dependencies (Install using `go get`)

### Installing

Clone the repository:

```bash
git clone https://github.com/pathusena/ExpenseTrackerGoLanguage.git
cd your-api

Install the required dependencies:

go get -u github.com/gorilla/mux
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlserver

Database Configuration

This API uses gorm for database interactions and the sqlserver driver for SQL Server. You need to configure the database connection in your project. Edit the sqlserverDB.go file and replace "sqlserver://<User>:<Password>@<Server>?database=<Database>" with your actual SQL Server connection string:

package sqlserver

import (
	"fmt"
	"expense-tracker-api/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	connectionstring := "sqlserver://<User>:<Password>@<Server>?database=<Database>"

	var err error
	Database, err = gorm.Open(sqlserver.Open(connectionstring), &gorm.Config{})

	Database.AutoMigrate(&models.Expense{})
	if err != nil {
		fmt.Println(err.Error())
	}
}

Running the API

Run the API using the following command:

go run main.go

The API will start and be available at http://localhost:8080.

Expense Struct

The API uses the following Expense struct for managing expenses:

type Expense struct {
    Id          uint
    Description string
    Amount      float64
    Date        string
}

API Endpoints

Get All Expenses
URL: /expenses
Method: GET
Description: Retrieve a list of all expenses.

Get an Expense
URL: /expense/{id}
Method: GET
Description: Retrieve details of a specific expense by ID.

Delete an Expense
URL: /expense/{id}
Method: DELETE
Description: Delete a specific expense by ID.

Create an Expense
URL: /expense
Method: POST
Description: Create a new expense.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments

Gorilla Mux - Router and dispatcher for Go.
gorm - The Go programming language ORM.
sqlserver - GORM SQL Server driver.