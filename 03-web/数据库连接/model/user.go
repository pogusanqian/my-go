package model

import (
	"fmt"
	"pogusanqian.com/db/utils"
)

//User 结构体
type Student struct {
	Id   int
	Name string
	Age  int
	Sex  string
}

func (stu *Student) AddStudent() error {
	sqlStr := "insert into t_student(f_name,f_age,f_sex) values(?,?,?)"
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return err
	}
	_, err2 := inStmt.Exec(stu.Name, stu.Age, stu.Sex)
	if err2 != nil {
		fmt.Println("执行出现异常：", err2)
		return err
	}
	return nil
}

func (stu *Student) AddStudent2() error {
	sqlStr := "insert into t_student(f_name,f_age,f_sex) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, stu.Name, stu.Age, stu.Sex)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

func (this *Student) GetStudentByID() (*Student, error) {
	sqlStr := "select f_id, f_name, f_age, f_sex from t_student where f_id = ?"
	// QuertRow只查询出一条数据, 就算由多条也只去1条
	row := utils.Db.QueryRow(sqlStr, this.Id)

	// 将取出的row值添加的对象上, 这里的步骤好麻烦, 应该会由框架来自动处理的
	var id int
	var name string
	var age int
	var sex string
	err := row.Scan(&id, &name, &age, &sex)
	if err != nil {
		return nil, err
	}
	stu := &Student{
		Id:   id,
		Name: name,
		Age:  age,
		Sex:  sex,
	}
	return stu, nil
}

func (this *Student) GetStudents() ([]*Student, error) {
	//写sql语句
	sqlStr := "select f_id, f_name, f_age, f_sex from t_student"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建Student切片
	var stus []*Student
	for rows.Next() {
		var id int
		var name string
		var age int
		var sex string
		err := rows.Scan(&id, &name, &age, &sex)
		if err != nil {
			return nil, err
		}
		stu := &Student{
			Id:   id,
			Name: name,
			Age:  age,
			Sex:  sex,
		}
		stus = append(stus, stu)
	}
	return stus, nil
}
