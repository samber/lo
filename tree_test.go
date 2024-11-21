package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTreeRowType struct {
	id       string
	parentId string
	value    string
}

type testTreeLeafType struct {
	value string
	leafs []testTreeLeafType
}

func testTreeRowToLeaf(row *testTreeRowType, children []testTreeLeafType) testTreeLeafType {
	return testTreeLeafType{
		value: row.value,
		leafs: children,
	}
}

func testTreeGetKeys(row *testTreeRowType) (string, string, bool) {
	return row.id, row.parentId, row.parentId == "<ROOT>"
}

func TestTree(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	table := []testTreeRowType{
		{"a", "<ROOT>", "A"},
		{"aa", "a", "AA"},
		{"ab", "a", "AB"},
		{"aba", "ab", "ABA"},
		{"ac", "a", "AC"},
		{"b", "<ROOT>", "B"},
	}

	tree, err := TableToTree(table, testTreeRowToLeaf, testTreeGetKeys)

	is.NoError(err)
	expected := []testTreeLeafType{
		{value: "A", leafs: []testTreeLeafType{
			{value: "AA"},
			{value: "AB", leafs: []testTreeLeafType{
				{value: "ABA"},
			}},
			{value: "AC"},
		}},
		{value: "B"},
	}
	is.Equal(expected, tree)
}

func TestEmptyTree(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	{
		table := []testTreeRowType{}
		tree, err := TableToTree(table, testTreeRowToLeaf, testTreeGetKeys)
		is.NoError(err)
		is.Nil(tree)
	}

	{
		tree, err := TableToTree(nil, testTreeRowToLeaf, testTreeGetKeys)
		is.NoError(err)
		is.Nil(tree)
	}
}

func TestDuplicatePrimaryId(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	table := []testTreeRowType{
		{"a", "", "A"},
		{"a", "", "A2"},
	}

	_, err := TableToTree(table, testTreeRowToLeaf, testTreeGetKeys)
	is.ErrorContains(err, "lo.TableToTree: duplicate primary id")
}

func TestBadParentId(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	table := []testTreeRowType{
		{"a", "", "A"},
		{"b", "x", "X"},
	}

	_, err := TableToTree(table, testTreeRowToLeaf, testTreeGetKeys)
	is.ErrorContains(err, "lo.TableToTree: bad parent id")
}

func TestTreeFromDoc(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type rowType struct {
		id       int
		parentId int
		value    string
	}

	type leafType struct {
		value string
		leafs []leafType
	}

	table := []rowType{
		{1, 0, "A"},
		{2, 1, "AA"},
		{3, 1, "AB"},
		{4, 3, "ABA"},
		{5, 1, "AC"},
		{6, 0, "B"},
	}

	tree, err := TableToTree(
		table,
		func(row *rowType, children []leafType) leafType {
			return leafType{
				value: row.value,
				leafs: children,
			}
		},
		func(row *rowType) (int, int, bool) {
			return row.id, row.parentId, row.parentId == 0
		},
	)

	is.NoError(err)
	expected := []leafType{
		{value: "A", leafs: []leafType{
			{value: "AA"},
			{value: "AB", leafs: []leafType{
				{value: "ABA"},
			}},
			{value: "AC"},
		}},
		{value: "B"},
	}
	is.Equal(expected, tree)
}
