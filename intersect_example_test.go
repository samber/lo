package lo

import (
	"fmt"
)

func ExampleWithoutBy() {
	type user struct {
		id   int
		name string
	}
	// original users
	users := []user{
		{id: 1, name: "Alice"},
		{id: 2, name: "Bob"},
		{id: 3, name: "Charlie"},
	}

	// exclude users with IDs 2 and 3
	excludedIDs := []int{2, 3}

	// extract function to get the user ID
	extractID := func(user user) int {
		return user.id
	}

	// filtering users
	filteredUsers := WithoutBy(users, extractID, excludedIDs...)

	// output the filtered users
	fmt.Printf("%v\n", filteredUsers)
	// Output:
	// [{1 Alice}]
}
