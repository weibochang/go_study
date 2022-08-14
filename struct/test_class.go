package main

import "fmt"

// 如果类名称首字母大写，表示其他包可以访问
type Hero struct {
	// 如果类属性首字母大写，表示该属性对外能够访问，否则的话只能够类的内部访问
	Name  string
	Ad    int
	Level int
}

/*
func (this Hero) Show() {
	fmt.Println("Name= ", this.Name)
	fmt.Println("Ad= ", this.Ad)
	fmt.Println("Level= ", this.Level)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	// 当前this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}
*/

// 方法命大写，表示可以从其他包可以访问
func (this *Hero) Show() {
	fmt.Println("Name= ", this.Name)
	fmt.Println("Ad= ", this.Ad)
	fmt.Println("Level= ", this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	// 当前this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}

func main() {
	// 创建一个对象
	hero := Hero{
		Name:  "zhang3",
		Ad:    100,
		Level: 1,
	}
	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
