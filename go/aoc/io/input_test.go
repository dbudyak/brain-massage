package io

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(".."); err != nil {
		t.Fatal(err)
	}

	testFile := filepath.Join("input", "input1.txt")
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Skip("input1.txt does not exist, skipping test")
	}

	content, err := ReadInput(1)
	if err != nil {
		t.Errorf("ReadInput(1) returned error: %v", err)
	}
	if content == "" {
		t.Error("ReadInput(1) returned empty content")
	}
	if !strings.Contains(content, "L68") {
		t.Errorf("ReadInput(1) returned unexpected content: %s", content)
	}
}

func TestReadInputNonExistent(t *testing.T) {
	_, err := ReadInput(9999)
	if err == nil {
		t.Error("ReadInput(9999) should return error for non-existent file")
	}
}

func TestMustReadInputPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustReadInput(9999) should panic for non-existent file")
		}
	}()

	MustReadInput(9999)
}
