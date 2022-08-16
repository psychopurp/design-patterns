package composite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackage_price(t *testing.T) {

	product1 := &Product{val: 5}
	product2 := &Product{val: 4}
	product3 := &Product{val: 2}

	package1 := &Package{packageFees: 1}

	package2 := &Package{packageFees: 2}

	package1.AddComponent(product1)
	package1.AddComponent(product2)

	package2.AddComponent(package1)
	package2.AddComponent(product3)

	assert.Equal(t, 14, package2.price())

}
