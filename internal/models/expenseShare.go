package models

type ExpenseShare struct {
	ExpenseID   int     `json:"expenseId"`
	UserID      int     `json:"userId"`
	ShareAmount float64 `json:"shareAmount"`
}
