package main

import "fmt"

type Man struct {
	Name string
}

func (m Man) Print() {
	fmt.Println("man")
}

func (m Man) change() {
	m.Name = "changed"
}

func (m *Man) changeP() {
	m.Name = "point-changed"
}

func main() {
	m := Man{}
	mp := new(Man)
	m.Name = "m"
	mp.Name = "mp"
	m.change()
	mp.change()
	fmt.Println(m, mp)
	m.changeP()
	mp.changeP()
	fmt.Println(m, mp)
}
