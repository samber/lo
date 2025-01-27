package lo

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type panicError struct{}

func (e panicError) Error() string {
	return "panic test error"
}

type recoverFunc interface{}
type recoverFuncName string
type testFuncName string

var (
	panicErr    = panicError{}
	callbackErr = errors.New("callback function test error")
	noErr       = errors.New("no error")
)

type testParams struct {
	x         uint8
	typed     bool
	withError bool
}

var functionMapping = map[recoverFuncName]struct {
	params      testParams
	recoverFunc recoverFunc
}{
	"Recover0":           {params: testParams{0, false, false}, recoverFunc: Recover0},
	"Recover1":           {params: testParams{1, false, false}, recoverFunc: Recover1[uint8]},
	"Recover2":           {params: testParams{2, false, false}, recoverFunc: Recover2[uint8, uint8]},
	"Recover3":           {params: testParams{3, false, false}, recoverFunc: Recover3[uint8, uint8, uint8]},
	"Recover4":           {params: testParams{4, false, false}, recoverFunc: Recover4[uint8, uint8, uint8, uint8]},
	"Recover5":           {params: testParams{5, false, false}, recoverFunc: Recover5[uint8, uint8, uint8, uint8, uint8]},
	"Recover6":           {params: testParams{6, false, false}, recoverFunc: Recover6[uint8, uint8, uint8, uint8, uint8, uint8]},
	"Recover0Error":      {params: testParams{0, false, true}, recoverFunc: Recover0Error},
	"Recover1Error":      {params: testParams{1, false, true}, recoverFunc: Recover1Error[uint8]},
	"Recover2Error":      {params: testParams{2, false, true}, recoverFunc: Recover2Error[uint8, uint8]},
	"Recover3Error":      {params: testParams{3, false, true}, recoverFunc: Recover3Error[uint8, uint8, uint8]},
	"Recover4Error":      {params: testParams{4, false, true}, recoverFunc: Recover4Error[uint8, uint8, uint8, uint8]},
	"Recover5Error":      {params: testParams{5, false, true}, recoverFunc: Recover5Error[uint8, uint8, uint8, uint8, uint8]},
	"Recover6Error":      {params: testParams{6, false, true}, recoverFunc: Recover6Error[uint8, uint8, uint8, uint8, uint8, uint8]},
	"Recover0Typed":      {params: testParams{0, true, false}, recoverFunc: Recover0Typed[error]},
	"Recover1Typed":      {params: testParams{1, true, false}, recoverFunc: Recover1Typed[uint8, error]},
	"Recover2Typed":      {params: testParams{2, true, false}, recoverFunc: Recover2Typed[uint8, uint8, error]},
	"Recover3Typed":      {params: testParams{3, true, false}, recoverFunc: Recover3Typed[uint8, uint8, uint8, error]},
	"Recover4Typed":      {params: testParams{4, true, false}, recoverFunc: Recover4Typed[uint8, uint8, uint8, uint8, error]},
	"Recover5Typed":      {params: testParams{5, true, false}, recoverFunc: Recover5Typed[uint8, uint8, uint8, uint8, uint8, error]},
	"Recover6Typed":      {params: testParams{6, true, false}, recoverFunc: Recover6Typed[uint8, uint8, uint8, uint8, uint8, uint8, error]},
	"Recover0ErrorTyped": {params: testParams{0, true, true}, recoverFunc: Recover0ErrorTyped[error]},
	"Recover1ErrorTyped": {params: testParams{1, true, true}, recoverFunc: Recover1ErrorTyped[uint8, error]},
	"Recover2ErrorTyped": {params: testParams{2, true, true}, recoverFunc: Recover2ErrorTyped[uint8, uint8, error]},
	"Recover3ErrorTyped": {params: testParams{3, true, true}, recoverFunc: Recover3ErrorTyped[uint8, uint8, uint8, error]},
	"Recover4ErrorTyped": {params: testParams{4, true, true}, recoverFunc: Recover4ErrorTyped[uint8, uint8, uint8, uint8, error]},
	"Recover5ErrorTyped": {params: testParams{5, true, true}, recoverFunc: Recover5ErrorTyped[uint8, uint8, uint8, uint8, uint8, error]},
	"Recover6ErrorTyped": {params: testParams{6, true, true}, recoverFunc: Recover6ErrorTyped[uint8, uint8, uint8, uint8, uint8, uint8, error]},
}

var testFuncMapping = map[testFuncName]func(t *testing.T, x uint8, inputErr error, rf recoverFunc, typed bool){
	"Success":          testSuccess,
	"Panic":            testPanic,
	"Error":            testError,
	"Unconfined":       testUnconfined,
	"StringPanic":      testStringPanic,
	"StringUnconfined": testStringUnconfined,
}

func TestTypedRecover(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Test case where no panic occurs but with a different type
	var err error
	is.PanicsWithError("test error", func() {
		defer TypedRecover[myError](&err)
		panic(errors.New("test error"))
	}, "expected unconfined panic but got nothing")
	is.Nil(err, "expected no caught error")

	// Test case where a string panic occurs and catchString is false
	is.PanicsWithValue("test string panic", func() {
		defer TypedRecover[error](&err, false)
		panic("test string panic")
	})
	is.Nil(err, "expected no redirected panic string")

	// Test case where no panic occurs
	is.NotPanics(func() {
		defer TypedRecover[error](&err)
	})
	is.NoError(err)

	// Test case where an error panic occurs
	is.NotPanics(func() {
		defer TypedRecover[error](&err)
		panic(errors.New("test error"))
	})
	is.EqualError(err, "test error")

	// Test case where a custom error panic occurs
	is.NotPanics(func() {
		defer TypedRecover[myError](&err)
		panic(myError{})
	})
	is.EqualError(err, myError{}.Error())

	// Test case where a string panic occurs and catchString is true
	is.NotPanics(func() {
		defer TypedRecover[error](&err, true)
		panic("test string panic")
	})
	is.EqualError(err, "test string panic")
}

func callRecover(rf recoverFunc, x uint8, f func(), inputErr error, catchString bool, typed bool) ([]uint8, error) {
	if x > 6 {
		panic("Unsupported X value")
	}

	results := make([]uint8, x)
	var err error

	nilErrorType := reflect.TypeOf((*error)(nil)).Elem()

	callbackOutTypes := make([]reflect.Type, x)
	for i := range callbackOutTypes {
		callbackOutTypes[i] = reflect.TypeOf(uint8(0))
	}
	if inputErr != nil {
		callbackOutTypes = append(callbackOutTypes, nilErrorType)
	}
	inTypes := make([]reflect.Type, 1, 2)
	inTypes[0] = reflect.FuncOf(nil, callbackOutTypes, false)

	if typed {
		inTypes = append(inTypes, reflect.SliceOf(reflect.TypeOf(bool(false))))
	}

	outTypes := make([]reflect.Type, x+1)
	for i := uint8(0); i < x; i++ {
		outTypes[i] = reflect.TypeOf(uint8(0))
	}
	outTypes[x] = nilErrorType

	dynamicFuncType := reflect.FuncOf(inTypes, outTypes, typed)

	callArgs := make([]reflect.Value, 1, 2)
	callArgs[0] = reflect.MakeFunc(inTypes[0], func([]reflect.Value) []reflect.Value {
		f()
		returnValues := make([]reflect.Value, x, x+1)
		for i := uint8(0); i < x; i++ {
			returnValues[i] = reflect.ValueOf(i + 1)
		}
		if inputErr == noErr {
			returnValues = append(returnValues, reflect.Zero(nilErrorType))
		} else if inputErr != nil {
			returnValues = append(returnValues, reflect.ValueOf(inputErr))
		}
		return returnValues
	})

	if typed {
		callArgs = append(callArgs, reflect.ValueOf(catchString))
	}

	fn := reflect.ValueOf(rf).Convert(dynamicFuncType)
	result := fn.Call(callArgs)

	// Extract results and error
	for i, val := range result[:x] {
		results[i] = uint8(val.Uint())
	}
	if !result[x].IsNil() {
		err = result[x].Interface().(error)
	}

	return results, err
}

func testSuccess(t *testing.T, x uint8, inputErr error, rf recoverFunc, typed bool) {
	t.Helper()
	expected := createTestValues(x)
	is := assert.New(t)

	is.NotPanics(func() {
		results, err := callRecover(rf, x, func() {}, inputErr, false, typed)
		is.NoError(err)
		is.Equal(expected, results)
	}, "expected no panic but got one")
}

func testPanic(t *testing.T, x uint8, inputErr error, rf recoverFunc, typed bool) {
	t.Helper()
	is := assert.New(t)

	assert.NotPanics(t, func() {
		results, err := callRecover(rf, x, func() { panic(panicErr) }, inputErr, false, typed)
		is.ErrorIs(err, panicErr)
		is.Empty(results)
	}, "panic was not caught")
}

func testError(t *testing.T, x uint8, inputErr error, rf recoverFunc, typed bool) {
	t.Helper()
	if inputErr == nil {
		t.Skip("Error test only applicable to -Error functions")
	}
	is := assert.New(t)

	results, err := callRecover(rf, x, func() {}, inputErr, false, typed)

	is.Equal(callbackErr, err)
	is.Empty(results)
}

func testUnconfined(t *testing.T, x uint8, inputErr error, rf recoverFunc, typed bool) {
	t.Helper()

	assert.PanicsWithValue(t, 1, func() {
		_, _ = callRecover(rf, x, func() { panic(1) }, inputErr, false, typed)
	}, "expected unconfined error panic but got nothing or wrong panic value")
}

func testStringPanic(t *testing.T, x uint8, inputErr error, rf recoverFunc, typed bool) {
	t.Helper()
	is := assert.New(t)

	is.NotPanics(func() {
		results, err := callRecover(rf, x, func() { panic("string panic") }, inputErr, true, typed)
		is.EqualError(err, "string panic")
		is.Empty(results)
	}, "expected string panic to be redirected to error return")
}

func testStringUnconfined(t *testing.T, x uint8, inputErr error, rf recoverFunc, typed bool) {
	t.Helper()
	is := assert.New(t)

	is.PanicsWithValue("string panic", func() {
		results, err := callRecover(rf, x, func() { panic("string panic") }, inputErr, false, typed)
		is.NoError(err)
		is.Empty(results)
	}, "expected unconfined string panic but got nothing or wrong panic value")
}

func createTestValues(x uint8) []uint8 {
	values := make([]uint8, x)
	for i := uint8(0); i < x; i++ {
		values[i] = i + 1
	}
	return values
}

func TestRecoverXTyped(t *testing.T) {
	t.Parallel()

	testMapping := map[testParams][]struct {
		inputErr error
		testFunc testFuncName
	}{
		{
			typed:     false,
			withError: false,
		}: {
			{
				inputErr: nil,
				testFunc: "Success",
			},
			{
				inputErr: nil,
				testFunc: "Panic",
			},
			{
				inputErr: nil,
				testFunc: "Unconfined",
			},
			{
				inputErr: nil,
				testFunc: "StringPanic",
			},
		},
		{
			typed:     false,
			withError: true,
		}: {
			{
				inputErr: noErr,
				testFunc: "Success",
			},
			{
				inputErr: noErr,
				testFunc: "Panic",
			},
			{
				inputErr: callbackErr,
				testFunc: "Error",
			},
			{
				inputErr: noErr,
				testFunc: "Unconfined",
			},
			{
				inputErr: noErr,
				testFunc: "StringPanic",
			},
		},
		{
			typed:     true,
			withError: false,
		}: {
			{
				inputErr: nil,
				testFunc: "Success",
			},
			{
				inputErr: nil,
				testFunc: "Panic",
			},
			{
				inputErr: nil,
				testFunc: "Unconfined",
			},
			{
				inputErr: nil,
				testFunc: "StringPanic",
			},
			{
				inputErr: nil,
				testFunc: "StringUnconfined",
			},
		},
		{
			typed:     true,
			withError: true,
		}: {
			{
				inputErr: noErr,
				testFunc: "Success",
			},
			{
				inputErr: noErr,
				testFunc: "Panic",
			},
			{
				inputErr: callbackErr,
				testFunc: "Error",
			},
			{
				inputErr: noErr,
				testFunc: "Unconfined",
			},
			{
				inputErr: noErr,
				testFunc: "StringPanic",
			},
			{
				inputErr: noErr,
				testFunc: "StringUnconfined",
			},
		},
	}

	for targetName, testTarget := range functionMapping {
		t.Run(string(targetName), func(t *testing.T) {
			t.Parallel()
			testFuncs := testMapping[testTarget.params]
			for _, testFuncParams := range testFuncs {
				runName := fmt.Sprintf("%s/%s", targetName, testFuncParams.testFunc)
				testFunc := testFuncMapping[testFuncParams.testFunc]
				t.Run(runName, func(t *testing.T) {
					t.Parallel()
					testFunc(t, testTarget.params.x, testFuncParams.inputErr, testTarget.recoverFunc, testTarget.params.typed)
				})
			}
		})
	}
}
