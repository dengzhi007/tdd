package fizzbuzz_test

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/dengzhi007/tdd/fizzbuzz"

func TestSay(t *testing.T) {
	s := fizzbuzz.Say(3, 5, 1)
	assert.Equal(t, s, "1")

	s = fizzbuzz.Say(3, 5, 2)
	assert.Equal(t, s, "2")

	s = fizzbuzz.Say(3, 5, 3)
	assert.Equal(t, s, "Fizz")
}
