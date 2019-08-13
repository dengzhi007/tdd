package fizzbuzz_test

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/dengzhi007/tdd/fizzbuzz"

func TestSay(t *testing.T) {
	s1 := fizzbuzz.Say(3, 5, 1)
	assert.Equal(t, s1, "1")
}
