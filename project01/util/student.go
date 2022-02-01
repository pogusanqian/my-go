package util

type student struct {
	Name string
	Age  int
}

func NewStudent(name string, age int) *student {
	return &student{
		Name: name,
		Age:  age,
	}
}
