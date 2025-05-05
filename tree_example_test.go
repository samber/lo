package lo

import (
	"fmt"
	"strings"
)

type exampleTreeRowType struct {
	id       int
	parentId int
	slug     string
	title    string
}

type exampleTreeLeafType struct {
	slug     string
	title    string
	children []exampleTreeLeafType
}

func ExampleTree() {
	table := []exampleTreeRowType{
		{1, 0, "el", "Electronics"},
		{2, 1, "ph", "Phones"},
		{3, 2, "ios", "iOS"},
		{4, 2, "and", "Android"},
		{5, 1, "tab", "Tablets"},
		{6, 1, "lap", "Laptops"},
		{7, 0, "fnt", "Furniture"},
		{8, 7, "kt", "Kitchen"},
		{9, 7, "bd", "Bedroom"},
		{10, 7, "lng", "Lounge"},
		{11, 0, "car", "Car"},
		{12, 11, "gps", "GPS"},
		{13, 11, "au", "Audio"},
		{14, 11, "al", "Alarm"},
		{15, 0, "oth", "Other"},
	}

	tree, _ := TableToTree(
		table,
		func(row *exampleTreeRowType, children []exampleTreeLeafType) exampleTreeLeafType {
			return exampleTreeLeafType{
				slug:  row.slug,
				title: row.title,

				children: children,
			}
		},
		func(row *exampleTreeRowType) (int, int, bool) {
			return row.id, row.parentId, row.parentId == 0
		},
	)

	printExmpleTree(tree, 0)
	// Output:
	// Electronics
	// .Phones
	// ..iOS
	// ..Android
	// .Tablets
	// .Laptops
	// Furniture
	// .Kitchen
	// .Bedroom
	// .Lounge
	// Car
	// .GPS
	// .Audio
	// .Alarm
	// Other
}

func printExmpleTree(leafs []exampleTreeLeafType, depth int) {
	for _, leaf := range leafs {
		fmt.Println(strings.Repeat(".", depth) + leaf.title)
		printExmpleTree(leaf.children, depth+1)
	}
}
