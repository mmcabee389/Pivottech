package main

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	IPAddress string `json:"ip_address"`
}
