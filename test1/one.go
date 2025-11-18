package main

import "fmt"

// 函数的返回方法
func fu1() int {
	x := 1
	return x
}

func fu2() (x int) {
	x = 2
	return
}

// 方法总是绑定对象实例fu
type User struct {
	id   int
	name string
}

// 方法的定义
func (u User) test() {
	fmt.Println(u)
}

// 嵌入字段
type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	tp   Person
	id   int
	addr string
}

type Student2 struct {
	tp   *Person
	id   int
	addr string
}

func main() {
	//x := fu1()
	//y := fu2()
	//one := User{1, "ZhangSan"}
	//one.test()
	//println(x)
	//println(y)

	s1 := Student{Person{"yang", "male", 18}, 1, "chongqing"}
	s2 := Student2{&Person{"yang", "male", 18}, 1, "chongqing"}
	fmt.Println(s1)
	fmt.Println(s2)
}
