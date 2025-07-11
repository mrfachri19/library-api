package entity

import "time"

type Lending struct {
	ID         int        `json:"id"`
	BookID     int        `json:"book_id"`
	Borrower   string     `json:"borrower"`
	BorrowDate time.Time  `json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"`
}
