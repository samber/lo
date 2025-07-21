package loslice

type AllocateMode int

const (
	AllocateZero     AllocateMode = iota // grows while elements were added (few results, expensive precount)
	AllocateAll                          // preallocate capacity for all elements (lots results, expensive precount)
	AllocatePrecount                     // counts the number of elements and allocates exactly that amount of memory
)
