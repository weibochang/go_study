// 面向对象的继承
package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

// 子类（继承）
type SuperMan struct {
	Human // SuperMan 类继承了Human类的方法
	levle int
}

// 重定义父类的方法Eat（）
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()....")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.levle)
}

func main() {
	fmt.Println()

	h := Human{"zhang3", "female"}
	h.Eat()
	h.Walk()

	// 定义一个子类对象
	// 定义方式一
	//s := SuperMan{Human{"li4", "female"}, 1}
	// 定义方式二
	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.levle = 1

	s.Walk() // 父类的方法
	s.Eat()  // 子类的方法
	s.Fly()  // 子类的方法

	s.Print()
}
