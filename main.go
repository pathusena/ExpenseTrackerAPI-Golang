package main

import (
	sqlserver "expense-tracker-api/database"
	models "expense-tracker-api/models"
	utils "expense-tracker-api/utils"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getExpensesFromDB() []models.Expense {
	var expenses []models.Expense
	sqlserver.Database.Find(&expenses)
	return expenses
}

func allExpenses(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, getExpensesFromDB())
}

func getExpense(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var expense models.Expense

	result := sqlserver.Database.Where("Id = ?", id).First(&expense)
	if result.Error == nil {
		utils.JsonResponse(w, expense)
	}
}

func createExpense(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var newExpense models.Expense

	utils.JsonDeserialize(reqBody, &newExpense)
	fmt.Println(reqBody)
	result := sqlserver.Database.Create(&newExpense)
	fmt.Println(result.Error)

	utils.JsonResponse(w, models.BaseResult{
		Result:  true,
		Message: "Expense has been created",
	})
}

func deleteExpense(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var deletedExpense models.Expense

	result := sqlserver.Database.Where("Id = ?", id).Delete(deletedExpense)
	fmt.Println(result.Error)

	utils.JsonResponse(w, models.BaseResult{
		Result:  true,
		Message: "Expense has been deleted",
	})
}

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(false)
	myrouter.HandleFunc("/expenses", allExpenses).Methods("GET")
	myrouter.HandleFunc("/expense{id}", getExpense).Methods("GET")
	myrouter.HandleFunc("/expense{id}", deleteExpense).Methods("DELETE")
	myrouter.HandleFunc("/expense", createExpense).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myrouter))
}

func main() {
	sqlserver.Init()
	handleRequests()
}
