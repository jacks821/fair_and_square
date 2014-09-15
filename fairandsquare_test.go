package main

import (
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

func TestSolve(t *testing.T) {
	assert := assert.New(t)
	for _, pair := range tests {
		v := Solve(pair.value)
		assert.Equal(v, pair.returnval,
			"For %v expected %v got %v.", pair.value, pair.returnval, v)
	}
}
