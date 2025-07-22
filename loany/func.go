package loany

import (
	"fmt"
	"github.com/samber/lo/loslice"
	"reflect"
)

type argSource string

const (
	fromArgs   argSource = "args"
	fromKwargs argSource = "kwargs"
)

func errCall(index int, value any, source argSource, fnType reflect.Type, kwType reflect.Type) error {
	typeName := fnType.In(index).Name()
	if kwType == nil {
		return fmt.Errorf("expected argument #%d to be of type `%s`, got `%#v`", index+1, typeName, value)
	}

	argName := kwType.Field(index).Name
	return fmt.Errorf("expected argument #%d (%s) to be of type `%s`, got from %s `%#v`", index+1, argName, typeName, source, value)
}

// Call calls a function with the provided args and kwargs.
// Does not support variadic functions.
// Does not panic on invalid input, but returns an error instead.
func Call(fn any, args []any, kwargs any) ([]any, error) {
	rvfn := reflect.ValueOf(fn)
	rvkw := reflect.ValueOf(kwargs)

	tfn := rvfn.Type()
	if tfn.Kind() != reflect.Func {
		return nil, fmt.Errorf("expected function, got `%#v`", fn)
	}

	if tfn.IsVariadic() {
		return nil, fmt.Errorf("variadic functions are not supported, got `%#v`", fn)
	}

	numIn := tfn.NumIn()

	if kwargs != nil {
		if rvkw.Kind() == reflect.Pointer {
			if rvkw.IsNil() {
				return nil, fmt.Errorf("expected kwargs to be non-nil pointer, got %#v", kwargs)
			}
			rvkw = rvkw.Elem()
		}

		if rvkw.Kind() != reflect.Struct {
			return nil, fmt.Errorf("expected kwargs to be struct, got `%#v`", kwargs)
		}

		if rvkw.NumField() != numIn {
			return nil, fmt.Errorf("expected kwargs (`%T`) to have %d fields, got %d: %#v",
				kwargs, numIn, rvkw.NumField(), kwargs)
		}
	} else if len(args) != numIn {
		return nil, fmt.Errorf("no kwargs were provided, expected exactly %d args, got %d: %#v", numIn, len(args), args)
	}

	var (
		val    any
		source argSource
	)

	rvargs := make([]reflect.Value, numIn)

	for i := range numIn {
		if len(args) > 0 {
			source = fromArgs
			val = args[0]
		} else if kwargs != nil {
			source = fromKwargs
			val = rvkw.Field(0).Interface()
		} else {
			panic("impossible: neither args nor kwargs were provided")
		}

		rv := reflect.ValueOf(val)
		tin := tfn.In(i)

		if !rv.CanConvert(tin) {
			if kwargs == nil {
				return nil, errCall(i, val, source, tfn, nil)
			} else {
				return nil, errCall(i, val, source, tfn, rvkw.Type())
			}
		}

		rvargs[i] = rv.Convert(tin)
	}

	return loslice.Map(rvfn.Call(rvargs), reflect.Value.Interface), nil
}
