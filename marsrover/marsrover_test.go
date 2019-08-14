package marsrover_test

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/dengzhi007/tdd/marsrover"

func TestRover(t *testing.T) {
	s := marsrover.Rover(5, 5, 1, 2, "N", "LMLMLMLMM")
	assert.Equal(t, s, "1 3 N")

	s = marsrover.Rover(5, 5, 3, 3, "E", "MMM")
	assert.Equal(t, s, "5 3 E RIP")
}
