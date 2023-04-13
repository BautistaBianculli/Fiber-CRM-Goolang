package models



type User struct {
	ID 			string 
	User 		string 	`json:"user"`
	Password 	string	`json:"password"`
	FirstName 	string	`json:"firstName,omitempty"`
	LastName 	string	`json:"lastName,omitempty"`
}