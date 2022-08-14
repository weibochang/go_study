package main

import "fmt"

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func ShowAnimal(animal AnimalIF) {
	animal.Sleep() // 多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {
	fmt.Println()
	/*
		var animal AnimalIF // 接口的数据类型，父类指针
		animal = &Cat{"Green"}

		animal.Sleep() // 调用Cat的Sleep()方法，多态的现象

		animal = &Dog{"Green"}
		animal.Sleep() // 调用Dog的Sleep()方法，多态的现象
	*/

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	ShowAnimal(&cat)
	ShowAnimal(&dog)

}
