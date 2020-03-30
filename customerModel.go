package model

import (
	"fmt"
)
type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
func NewCustomer2(name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
func (this Customer) GetInfo() string {
	Info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		this.ID, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return Info
}
