package lo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttempt(t *testing.T) {
	is := assert.New(t)

	err := fmt.Errorf("failed")

	iter1, err1 := Attempt(42, func(i int) error {
		return nil
	})
	iter2, err2 := Attempt(42, func(i int) error {
		if i == 5 {
			return nil
		}

		return err
	})
	iter3, err3 := Attempt(2, func(i int) error {
		if i == 5 {
			return nil
		}

		return err
	})
	iter4, err4 := Attempt(0, func(i int) error {
		if i < 42 {
			return fmt.Errorf("failed")
		}

		return nil
	})

	is.Equal(iter1, 1)
	is.Equal(err1, nil)
	is.Equal(iter2, 6)
	is.Equal(err2, nil)
	is.Equal(iter3, 2)
	is.Equal(err3, err)
	is.Equal(iter4, 43)
	is.Equal(err4, nil)
}
