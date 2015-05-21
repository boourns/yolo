package yolo

import (
	"os"
	"testing"
)

type Thing struct {
	Name   string
	Number int
}

var Collection = []Thing{
	Thing{Name: "Alice", Number: 1},
	Thing{Name: "Bob", Number: 2},
	Thing{Name: "Chris", Number: 3},
}

func TestYoloWrite(t *testing.T) {
	file := randomName()
	err := Save(file, Collection)

	defer func() {
		os.Remove(file)
	}()

	if err != nil {
		t.Errorf("Error saving: %s", err)
	}

	var input []*Thing
	err = Load(file, &input)

	if err != nil {
		t.Errorf("Error loading: %s", err)
	}

	if len(input) != 3 {
		t.Errorf("Loaded %d items, expected 3", len(input))
	}

	if input[2].Name != "Chris" {
		t.Errorf("Expected last user name to be Chris, received %s", input[2].Name)
	}

}
