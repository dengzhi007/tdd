package fizzbuzz_test

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/dengzhi007/tdd/fizzbuzz"

func TestFizzBuzz(t *testing.T) {
	s := fizzbuzz.FizzBuzz(3, 5, 1)
	assert.Equal(t, s, "1")
	s = fizzbuzz.FizzBuzz(3, 5, 3)
	assert.Equal(t, s, "Fizz")
	s = fizzbuzz.FizzBuzz(3, 5, 5)
	assert.Equal(t, s, "Buzz")
	s = fizzbuzz.FizzBuzz(3, 5, 6)
	assert.Equal(t, s, "Fizz")
	s = fizzbuzz.FizzBuzz(3, 5, 9)
	assert.Equal(t, s, "Fizz")
	s = fizzbuzz.FizzBuzz(3, 5, 10)
	assert.Equal(t, s, "Buzz")
	s = fizzbuzz.FizzBuzz(3, 5, 13)
	assert.Equal(t, s, "Fizz")
	s = fizzbuzz.FizzBuzz(3, 5, 15)
	assert.Equal(t, s, "FizzBuzz")
	s = fizzbuzz.FizzBuzz(3, 5, 23)
	assert.Equal(t, s, "Fizz")
	s = fizzbuzz.FizzBuzz(3, 5, 25)
	assert.Equal(t, s, "Buzz")
	s = fizzbuzz.FizzBuzz(3, 5, 52)
	assert.Equal(t, s, "Buzz")
}
