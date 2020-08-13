package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	builder := &MacBookBuilder{}
	dir := &Director{}
	fmt.Println(dir)
	pcDirector := NewDirector(builder)
	pcDirector.Construct("英特尔主板", "Retina显示器")

	pcDirector.builder = &HPBuilder{}
	pcDirector.Construct("AMD主办单位", "SONY 显示器")

	computer := builder.build()
	computer.Print()

	hp := pcDirector.builder.build()
	hp.Print()

}
