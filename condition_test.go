package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	is := assert.New(t)

	result1 := Ternary[string](true, "a", "b")
	result2 := Ternary[string](false, "a", "b")

	is.Equal(result1, "a")
	is.Equal(result2, "b")
}

func TestIfElse(t *testing.T) {
	is := assert.New(t)

	result1 := If[int](true, 1).ElseIf(false, 2).Else(3)
	result2 := If[int](true, 1).ElseIf(true, 2).Else(3)
	result3 := If[int](false, 1).ElseIf(true, 2).Else(3)
	result4 := If[int](false, 1).ElseIf(false, 2).Else(3)

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 2)
	is.Equal(result4, 3)
}

func TestIfFElseF(t *testing.T) {
	is := assert.New(t)

	result1 := IfF[int](true, func() int { return 1 }).ElseIfF(false, func() int { return 2 }).ElseF(func() int { return 3 })
	result2 := IfF[int](true, func() int { return 1 }).ElseIfF(true, func() int { return 2 }).ElseF(func() int { return 3 })
	result3 := IfF[int](false, func() int { return 1 }).ElseIfF(true, func() int { return 2 }).ElseF(func() int { return 3 })
	result4 := IfF[int](false, func() int { return 1 }).ElseIfF(false, func() int { return 2 }).ElseF(func() int { return 3 })

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 2)
	is.Equal(result4, 3)
}

func TestSwitchCase(t *testing.T) {
	is := assert.New(t)

	result1 := Switch[int, int](42).Case(42, 1).Case(1, 2).Default(3)
	result2 := Switch[int, int](42).Case(42, 1).Case(42, 2).Default(3)
	result3 := Switch[int, int](42).Case(1, 1).Case(42, 2).Default(3)
	result4 := Switch[int, int](42).Case(1, 1).Case(1, 2).Default(3)

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 2)
	is.Equal(result4, 3)
}

func TestSwitchCaseF(t *testing.T) {
	is := assert.New(t)

	result1 := Switch[int, int](42).CaseF(42, func() int { return 1 }).CaseF(1, func() int { return 2 }).DefaultF(func() int { return 3 })
	result2 := Switch[int, int](42).CaseF(42, func() int { return 1 }).CaseF(42, func() int { return 2 }).DefaultF(func() int { return 3 })
	result3 := Switch[int, int](42).CaseF(1, func() int { return 1 }).CaseF(42, func() int { return 2 }).DefaultF(func() int { return 3 })
	result4 := Switch[int, int](42).CaseF(1, func() int { return 1 }).CaseF(1, func() int { return 2 }).DefaultF(func() int { return 3 })

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 2)
	is.Equal(result4, 3)
}
