package service

import (
	"pogusanqian.com/project01/customerManage/model"
)

//声明service结构体
type CustomerService struct {
	customers   []model.Customer // 定义了切片数组, 相当于是数据库
	customerNum int              // 这个相当于是数据库表的自增主键ID记录
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//添加客户信息
func (this *CustomerService) Add(customer model.Customer) bool {
	//我们确定一个分配id的规则,就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//查看客户信息
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 修改客户信息
func (this *CustomerService) Update(customer model.Customer) bool {
	index := this.FindById(customer.Id)
	this.customers = append(append(this.customers[:index], customer), this.customers[index+1:]...)
	return true
}

//删除客户信息
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	//如果index == -1, 说明没有这个客户
	if index == -1 {
		return false
	}
	//从切片中删除信息(两个切片合成了一个新的切片, 忽略了删除的那个字段)
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

//根据id查找index下标, 查找不到就返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历this.customers 切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

//根据id查找用户信息
func (this *CustomerService) FindCustomerById(id int) string {
	res := "用户信息不存在, 请重新输入ID"
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			res = this.customers[i].GetInfo()
		}
	}
	return res
}
