package lo

import (
	"time"

	"github.com/kr/pretty"
)

// This example demonstrates a relatively straight forward way of running a function asynchronously
func ExampleAsync() {
	ch := Async(func() Tuple2[int, error] {
		time.Sleep(10 * time.Second)
		return Tuple2[int, error]{42, nil}
	})
	result := <-ch
	pretty.Println(result)
	// Output:
	// lo.Tuple2[int,error]{
	//     A:  42,
	//     B:  nil,
	// }
}
