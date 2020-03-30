package service

import (
	"go_code/chapter12项目/Customer/model"
)
type CustomerService struct {
	customers []model.Customer
	//表示当前切片含有多少个客户，该字段后面还可以作为新用户的id+1
	CustomerNum int
}
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.CustomerNum = 1
	customers := model.NewCustomer(1, "张三", "男", 20, "112", "aaa@163.com")
	customerService.customers = append(customerService.customers, customers)
	return customerService

}
func (this *CustomerService) List() []model.Customer {
	return this.customers
}
func (this *CustomerService) Add(customer model.Customer) bool {
	//我们确定一个分配id的规则，就是添加的顺序
	this.CustomerNum++
	customer.ID = this.CustomerNum
	this.customers = append(this.customers, customer)
	return true
}
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
func (this *CustomerService) Update(customer model.Customer, id int) bool {
	index1 := this.FindById(id)
	if index1 == -1 {
		return false
	}
	this.customers[id-1] = customer
	return true
}

//根据id查找客户在切片中对应的下标,没有该客户返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].ID == id {
			index = i
		}
	}
	return index
}
