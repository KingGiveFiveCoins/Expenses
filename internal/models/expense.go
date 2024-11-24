package models

type Expense struct {
	ID          int            `json:"id"`
	UserID      int            `json:"userId"`
	GroupID     int            `json:"groupId"`
	Amount      int            `json:"amount"`
	Currency    string         `json:"currency"`
	Name        string         `json:"name"`
	Category    string         `json:"category"`
	Description string         `json:"description"`
	Date        string         `json:"date"`
	Shares      []ExpenseShare `json:"shares"`
}
