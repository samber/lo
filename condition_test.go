package lo_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Ternary(true, "a", "b")
	result2 := lo.Ternary(false, "a", "b")

	is.Equal(result1, "a")
	is.Equal(result2, "b")
}

func TestTernaryF(t *testing.T) {
	is := assert.New(t)

	result1 := lo.TernaryF(true, func() string { return "a" }, func() string { return "b" })
	result2 := lo.TernaryF(false, func() string { return "a" }, func() string { return "b" })

	is.Equal(result1, "a")
	is.Equal(result2, "b")
}

func TestIfElse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.If(true, 1).ElseIf(false, 2).Else(3)
	result2 := lo.If(true, 1).ElseIf(true, 2).Else(3)
	result3 := lo.If(false, 1).ElseIf(true, 2).Else(3)
	result4 := lo.If(false, 1).ElseIf(false, 2).Else(3)

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 2)
	is.Equal(result4, 3)
}

func TestIfFElseF(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.IfF(true, func() int { return 1 }).ElseIfF(false, func() int { return 2 }).ElseF(func() int { return 3 })
	result2 := lo.IfF(true, func() int { return 1 }).ElseIfF(true, func() int { return 2 }).ElseF(func() int { return 3 })
	result3 := lo.IfF(false, func() int { return 1 }).ElseIfF(true, func() int { return 2 }).ElseF(func() int { return 3 })
	result4 := lo.IfF(false, func() int { return 1 }).ElseIfF(false, func() int { return 2 }).ElseF(func() int { return 3 })

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 2)
	is.Equal(result4, 3)
}

func TestSwitchCase(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Switch[int, int](42).Case(42, 1).Case(1, 2).Default(3)
	result2 := lo.Switch[int, int](42).Case(42, 1).Case(42, 2).Default(3)
	result3 := lo.Switch[int, int](42).Case(1, 1).Case(42, 2).Default(3)
	result4 := lo.Switch[int, int](42).Case(1, 1).Case(1, 2).Default(3)

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 2)
	is.Equal(result4, 3)
}

func TestSwitchCaseF(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := lo.Switch[int, int](42).CaseF(42, func() int { return 1 }).CaseF(1, func() int { return 2 }).DefaultF(func() int { return 3 })
	result2 := lo.Switch[int, int](42).CaseF(42, func() int { return 1 }).CaseF(42, func() int { return 2 }).DefaultF(func() int { return 3 })
	result3 := lo.Switch[int, int](42).CaseF(1, func() int { return 1 }).CaseF(42, func() int { return 2 }).DefaultF(func() int { return 3 })
	result4 := lo.Switch[int, int](42).CaseF(1, func() int { return 1 }).CaseF(1, func() int { return 2 }).DefaultF(func() int { return 3 })

	is.Equal(result1, 1)
	is.Equal(result2, 1)
	is.Equal(result3, 2)
	is.Equal(result4, 3)
}
