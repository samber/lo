package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		cond     bool
		expected string
	}{
		{name: "true", cond: true, expected: "a"},
		{name: "false", cond: false, expected: "b"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Ternary(tt.cond, "a", "b")
			is.Equal(tt.expected, result)
		})
	}
}

func TestTernaryF(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		cond     bool
		expected string
	}{
		{name: "true", cond: true, expected: "a"},
		{name: "false", cond: false, expected: "b"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := TernaryF(tt.cond, func() string { return "a" }, func() string { return "b" })
			is.Equal(tt.expected, result)
		})
	}
}

func TestIfElse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		cond1    bool
		cond2    bool
		expected int
	}{
		{name: "first condition true", cond1: true, cond2: false, expected: 1},
		{name: "both conditions true", cond1: true, cond2: true, expected: 1},
		{name: "second condition true", cond1: false, cond2: true, expected: 2},
		{name: "no condition true", cond1: false, cond2: false, expected: 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := If(tt.cond1, 1).ElseIf(tt.cond2, 2).Else(3)
			is.Equal(tt.expected, result)
		})
	}
}

func TestIfFElseF(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		cond1    bool
		cond2    bool
		expected int
	}{
		{name: "first condition true", cond1: true, cond2: false, expected: 1},
		{name: "both conditions true", cond1: true, cond2: true, expected: 1},
		{name: "second condition true", cond1: false, cond2: true, expected: 2},
		{name: "no condition true", cond1: false, cond2: false, expected: 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := IfF(tt.cond1, func() int { return 1 }).ElseIfF(tt.cond2, func() int { return 2 }).ElseF(func() int { return 3 })
			is.Equal(tt.expected, result)
		})
	}
}

func TestSwitchCase(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		case1    int
		case2    int
		expected int
	}{
		{name: "first case matches", case1: 42, case2: 1, expected: 1},
		{name: "both cases match", case1: 42, case2: 42, expected: 1},
		{name: "second case matches", case1: 1, case2: 42, expected: 2},
		{name: "no case matches", case1: 1, case2: 1, expected: 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Switch[int, int](42).Case(tt.case1, 1).Case(tt.case2, 2).Default(3)
			is.Equal(tt.expected, result)
		})
	}
}

func TestSwitchCaseF(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		case1    int
		case2    int
		expected int
	}{
		{name: "first case matches", case1: 42, case2: 1, expected: 1},
		{name: "both cases match", case1: 42, case2: 42, expected: 1},
		{name: "second case matches", case1: 1, case2: 42, expected: 2},
		{name: "no case matches", case1: 1, case2: 1, expected: 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Switch[int, int](42).CaseF(tt.case1, func() int { return 1 }).CaseF(tt.case2, func() int { return 2 }).DefaultF(func() int { return 3 })
			is.Equal(tt.expected, result)
		})
	}
}
