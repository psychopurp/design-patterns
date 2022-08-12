package factory

import (
	"testing"
)

func TestNewGun(t *testing.T) {

	gun1 := NewGun("ak47")
	gun1.Shoot()

	gun2 := NewGun("musket")
	gun2.Shoot()
}
