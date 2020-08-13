package builder

import "fmt"

/*
场景：构造一台电脑
*/

type (
	//builder接口
	Builder interface {
		buildBoard(board string)
		buildDisplay(display string)
		buildOS()
		build() Computer
	}

	Computer struct {
		board   string
		display string
		os      string
	}

	//macbook
	MacBookBuilder struct {
		computer Computer
	}

	//HP 惠普电脑
	HPBuilder struct {
		computer Computer
	}
)

func (c Computer) Print() {
	fmt.Printf("board: %v , display: %v , os: %v \n", c.board, c.display, c.os)
}

// MacBook builder 实现builder接口
func (builder *MacBookBuilder) buildBoard(board string) {
	builder.computer.board = board
}
func (builder *MacBookBuilder) buildDisplay(display string) {
	builder.computer.display = display
}
func (builder *MacBookBuilder) buildOS() {
	builder.computer.os = "MAC OS"
}
func (builder *MacBookBuilder) build() Computer {
	return builder.computer
}

// HP builder 实现builder接口
func (builder *HPBuilder) buildBoard(board string) {
	builder.computer.board = board
}
func (builder *HPBuilder) buildDisplay(display string) {
	builder.computer.display = display
}
func (builder *HPBuilder) buildOS() {
	builder.computer.os = "MAC OS"
}
func (builder *HPBuilder) build() Computer {
	return builder.computer
}

//监工
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct(board string, display string) {
	d.builder.buildBoard(board)
	d.builder.buildDisplay(display)
	d.builder.buildOS()
}
