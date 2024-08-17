package lo

import (
	"fmt"
)

func ExampleWithoutBy() {
	type user struct {
		id   int
		name string
	}
	// Example usage
	users := []user{
		{id: 1, name: "Alice"},
		{id: 2, name: "Bob"},
		{id: 3, name: "Charlie"},
	}

	// Exclude users with IDs 2 and 3
	excludedIDs := []int{2, 3}

	// Extract function to get the user ID
	extractID := func(user user) int {
		return user.id
	}

	// Filtering users
	filteredUsers := WithoutBy(users, extractID, excludedIDs...)

	// Output the filtered users
	fmt.Printf("%v\n", filteredUsers)
	// Output:
	// [{1 Alice}]
}
