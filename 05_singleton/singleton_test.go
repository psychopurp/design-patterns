package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetInstance(), GetInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			assert.Equal(b, GetInstance(), GetInstance())
		}
	})
}
