package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func addByOne(input int) int {
	return input + 1
}

func multiplyByTwo(input int) int {
	return input * 2
}

func TestChain(t *testing.T) {
	is := assert.New(t)

	chain := Chain(addByOne, multiplyByTwo)
	output := chain(5)
	is.Equal(12, output)
}

func TestChainNoMiddlewares(t *testing.T) {
	is := assert.New(t)

	chain := Chain(addByOne)
	output := chain(5)
	is.Equal(6, output)
}

func TestChainComposed(t *testing.T) {
	is := assert.New(t)

	chained := Chain(addByOne, multiplyByTwo)
	chainedWithSquare := Chain(chained, func(input int) int {
		return input * input
	})

	output := chainedWithSquare(5)
	is.Equal(144, output)
}
