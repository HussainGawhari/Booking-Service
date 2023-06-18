package models

type Users struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id" gorm:"unique"`
	UserName  string `json:"user_name"`
	UserPhone string `json:"user_phone"`
	UserEmail string `json:"user_email" gorm:"unique"`
	SeatID    int    `json:"seat_id"`
	SeatClass string `json:"seat_class"`
}
