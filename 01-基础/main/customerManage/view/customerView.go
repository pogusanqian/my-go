package main

import (
	"fmt"
	"pogusanqian.com/project01/customerManage/model"
	"pogusanqian.com/project01/customerManage/service"
)

// 注意这里key和loop是直接封装到了结构体的属性中的, 并没有写在主页面方法中
// 这是一种抽象的思想, 我们的程序一致会记录用户的选项值的, 这个loop值是可以不要的, 因为5就是表示退出; 可能是为了以后好添加对应的扩展
type customerView struct {
	key             string //接收用户输入...
	loop            bool   //表示是否循环的显示主菜单
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

//得到用户的输入，信息构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例, id号没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

// 修改用户
func (this *customerView) update() {
	fmt.Println("---------------------修改客户---------------------")
	fmt.Println("请输入修改客户ID:")
	// 输出原本信息
	var id int
	fmt.Scanln(&id)
	// TODO 这里还要在做一个异常处理的, 因为用户输入的id不一定都能找到值
	str := this.customerService.FindCustomerById(id)
	fmt.Println("原本客户信息: ", str)
	// 输入新信息
	fmt.Println("请输入新用户信息")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer(id, name, gender, age, phone, email)
	// 进行更改
	this.customerService.Update(customer)
}

//得到用户的输入id，删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("---------------------删除失败，输入的id号不存在----")
		}
	}
}

//退出软件
func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

func (this *customerView) mainMenu() {
	for {
		// 显示首页
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")

		// 进行增删改查选择
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		//  判断是否要跳出循环
		if !this.loop {
			break
		}

	}

	fmt.Println("你退出了客户关系管理系统...")
}

func main() {
	// 实例化view对象
	customerView := customerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}

	//显示主菜单..
	customerView.mainMenu()
}
