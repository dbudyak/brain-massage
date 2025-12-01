package structs

import (
	"reflect"
	"testing"
)

func TestCyclicList_Size(t *testing.T) {
	list := CreateList(10)
	if list.Length() != 10 {
		t.Errorf("CyclicList size is incorrect, expected %d, got %d", 10, list.Length())
	}
}

func TestCyclicList_Add(t *testing.T) {
	list := CreateList(2)
	list.Add(1)
	if list.Length() != 3 {
		t.Errorf("CyclicList size is incorrect, expected %d, got %d", 3, list.Length())
	}
}

func TestCyclicList_contents(t *testing.T) {
	want := []int{0, 1, 2, 3, 4}
	got := CreateList(5).ToArray()

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("CyclicList contents is incorrect, expected %d, got %d", want, got)
	}
}

func TestCyclicList_Get(t *testing.T) {
	list := CreateList(5)

	want := 1
	got := list.Get(6).value
	if got != want {
		t.Errorf("CyclicList get is incorrect, expected %d, got %d", want, got)
	}
}

func TestCyclicList_GetNegativeIndex(t *testing.T) {
	list := CreateList(5)

	tests := []struct {
		name  string
		index int
		want  int
	}{
		{"last element", -1, 4},
		{"second to last", -2, 3},
		{"third to last", -3, 2},
		{"fourth to last", -4, 1},
		{"fifth to last", -5, 0},
		{"wraparound negative", -6, 4}, // Should wrap around to last element
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := list.Get(tt.index).value
			if got != tt.want {
				t.Errorf("Get(%d) = %d, want %d", tt.index, got, tt.want)
			}
		})
	}
}

func TestCyclicList_GetDoesNotModifyList(t *testing.T) {
	list := CreateList(5)
	original := list.ToArray()

	list.Get(0)
	list.Get(3)
	list.Get(-1)
	list.Get(10)

	current := list.ToArray()
	if !reflect.DeepEqual(original, current) {
		t.Errorf("Get modified the list. Original: %v, After Get: %v", original, current)
	}
}
