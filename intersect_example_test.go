package lo

import (
	"fmt"
)

func ExampleWithoutBy() {
	type User struct {
		ID   int
		Name string
	}
	// original users
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	// exclude users with IDs 2 and 3
	excludedIDs := []int{2, 3}

	// extract function to get the user ID
	extractID := func(user User) int {
		return user.ID
	}

	// filtering users
	filteredUsers := WithoutBy(users, extractID, excludedIDs...)

	// output the filtered users
	fmt.Printf("%v", filteredUsers)
	// Output:
	// [{1 Alice}]
}
