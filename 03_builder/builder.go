package builder

type IBuilder interface {
	buildBoard(board string)
	buildDisplay(display string)
	buildOS()
	build() Computer
}

type Computer struct {
	board   string
	display string
	os      string
}

// macbook builder
type MacBookBuilder struct {
	computer Computer
}

// build implements IBuilder
func (b *MacBookBuilder) build() Computer {
	return b.computer
}

// buildBoard implements IBuilder
func (b *MacBookBuilder) buildBoard(board string) {
	b.computer.board = board
}

// buildDisplay implements IBuilder
func (b *MacBookBuilder) buildDisplay(display string) {
	b.computer.display = display
}

// buildOS implements IBuilder
func (b *MacBookBuilder) buildOS() {
	b.computer.os = "MAC OS"
}

//HP builder
type HPBuilder struct {
	computer Computer
}

// build implements IBuilder
func (b *HPBuilder) build() Computer {
	return b.computer
}

// buildBoard implements IBuilder
func (b *HPBuilder) buildBoard(board string) {
	b.computer.board = board
}

// buildDisplay implements IBuilder
func (b *HPBuilder) buildDisplay(display string) {
	b.computer.display = display
}

// buildOS implements IBuilder
func (b *HPBuilder) buildOS() {
	b.computer.os = "Windows OS"
}

type Director struct {
	builder IBuilder
}

func NewDirector(builder IBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct(board string, display string) Computer {
	d.builder.buildBoard(board)
	d.builder.buildDisplay(display)
	d.builder.buildOS()
	return d.builder.build()
}
