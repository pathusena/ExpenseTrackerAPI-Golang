package sqlserver

import (
	"fmt"
	"expense-tracker-api/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	connectionstring := "sqlserver://azure:@Sql2023321@interviewtrackersqlserver.database.windows.net?database=ExpenseTrackerDB"

	var err error
	Database, err = gorm.Open(sqlserver.Open(connectionstring), &gorm.Config{})

	Database.AutoMigrate(&models.Expense{})
	if err != nil {
		fmt.Println(err.Error())
	}
}