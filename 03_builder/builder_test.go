package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirector_Construct(t *testing.T) {

	computer := NewDirector(&MacBookBuilder{}).Construct("英特尔主板", "Retina显示器")

	assert.Equal(t, Computer{
		board:   "英特尔主板",
		display: "Retina显示器",
		os:      "MAC OS",
	}, computer)

	computer = NewDirector(&HPBuilder{}).Construct("AMD主办单位", "SONY 显示器")

	assert.Equal(t, Computer{
		board:   "AMD主办单位",
		display: "SONY 显示器",
		os:      "Windows OS",
	}, computer)

}
