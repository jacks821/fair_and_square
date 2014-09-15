package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testpair struct {
	value     []string
	returnval int
}

var tests = []testpair{
	{[]string{"1", "4"}, 2},
	{[]string{"10", "120"}, 0},
	{[]string{"100", "1000"}, 2},
	{[]string{"9", "10000200002"}, 20},
}

func generateTestPalins(init int64, ending int64) []int64 {
	var palins []int64
	for _, palin := range GeneratePalindromes(init, ending) {
		if isSquare(palin) {
			root := int64(math.Sqrt(float64(palin)))
			if isPalindrome(root) {
				palins = append(palins, palin)
			}
		}
	}
	return palins
}

func TestSolve(t *testing.T) {
	palins := generateTestPalins(1, 10000200002)
	assert := assert.New(t)
	for _, pair := range tests {
		v := Solve(pair.value, palins)
		assert.Equal(v, pair.returnval,
			"For %v expected %v got %v.", pair.value, pair.returnval, v)
	}
}
