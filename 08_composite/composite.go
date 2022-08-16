package composite

// define the component
type Component interface {
	price() int
}

// Package contains small packages and products
type Package struct {
	components  []Component
	packageFees int
}

type Product struct {
	val int
}

func (p *Package) AddComponent(c Component) {
	p.components = append(p.components, c)
}

func (p *Package) price() int {
	var price int
	for _, c := range p.components {
		price += c.price()
	}
	return price + p.packageFees
}

func (p *Product) price() int {
	return p.val
}
