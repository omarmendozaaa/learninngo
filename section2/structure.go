package main

import "fmt"

//structure encapsulation
type Car struct {
	Name  string
	Age   int
	Model int
}

func (c Car) Print() {
	fmt.Println(c)
}
func (c Car) Drive() {
	fmt.Println("Driving . . . ")
}
func (c Car) GetName() {
	fmt.Print(c.Name + "-->")
}
func main() {
	c := Car{
		Name:  "chevy",
		Age:   1,
		Model: 2,
	}
	c.GetName()
	c.Drive()
}
