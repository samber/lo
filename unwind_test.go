package lo

import (
	"reflect"
	"testing"
)

func TestUnwind(t *testing.T) {
	type args struct {
		order      []int
		collection []string
	}
	tests := []struct {
		name                 string
		args                 args
		wantSortedCollection []string
		wantStochasticTenet  []int
	}{
		{
			name:                 "accepts zero items",
			args:                 args{order: []int{}, collection: []string{}},
			wantSortedCollection: []string{},
			wantStochasticTenet:  []int{},
		},
		{
			name:                 "accepts 1 item",
			args:                 args{order: []int{0}, collection: []string{"a"}},
			wantSortedCollection: []string{"a"},
			wantStochasticTenet:  []int{0},
		},
		{
			name:                 "accepts 2 items",
			args:                 args{order: []int{1, 0}, collection: []string{"b", "a"}},
			wantSortedCollection: []string{"a", "b"},
			wantStochasticTenet:  []int{1, 0},
		},
		{
			name:                 "accepts 3 items",
			args:                 args{order: []int{1, 2, 0}, collection: []string{"b", "c", "a"}},
			wantSortedCollection: []string{"a", "b", "c"},
			wantStochasticTenet:  []int{2, 0, 1},
		},
		{
			name:                 "accepts 4 items",
			args:                 args{order: []int{1, 3, 2, 0}, collection: []string{"b", "d", "c", "a"}},
			wantSortedCollection: []string{"a", "b", "c", "d"},
			wantStochasticTenet:  []int{3, 0, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSortedCollection, gotStochasticTenet := Unwind(tt.args.order, tt.args.collection)
			if !reflect.DeepEqual(gotSortedCollection, tt.wantSortedCollection) {
				t.Errorf("Unwind() gotSortedCollection = %v, want %v", gotSortedCollection, tt.wantSortedCollection)
			}
			if !reflect.DeepEqual(gotStochasticTenet, tt.wantStochasticTenet) {
				t.Errorf("Unwind() gotStochasticTenet = %v, want %v", gotStochasticTenet, tt.wantStochasticTenet)
			}
		})
	}
	t.Run("is a reversible operation", func(t *testing.T) {
		baseOrder, baseCollection := []int{1, 3, 2, 0}, []string{"b", "d", "c", "a"}
		midSortedCollection, midStochasticTenet := Unwind(baseOrder, baseCollection)
		gotSortedCollection, gotStochasticTenet := Unwind(midStochasticTenet, midSortedCollection)

		if !reflect.DeepEqual(gotSortedCollection, baseCollection) {
			t.Errorf("Unwind() gotSortedCollection = %v, want %v", gotSortedCollection, baseCollection)
		}
		if !reflect.DeepEqual(gotStochasticTenet, baseOrder) {
			t.Errorf("Unwind() gotStochasticTenet = %v, want %v", gotStochasticTenet, baseOrder)
		}
	})
}
