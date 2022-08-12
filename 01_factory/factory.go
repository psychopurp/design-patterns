package factory

import "fmt"

// Gun describe the gun's method
type Gun interface {
	Shoot()
}

type AK47 struct {
}

func (g *AK47) Shoot() {
	fmt.Println("AK47 is shooting")
}

type Musket struct {
}

func (g *Musket) Shoot() {
	fmt.Println("Musket is shooting")
}

// NewGun
func NewGun(gunType string) Gun {
	switch gunType {
	case "ak47":
		return &AK47{}
	case "musket":
		return &Musket{}
	default:
		panic("wrong gun type")
	}
}
