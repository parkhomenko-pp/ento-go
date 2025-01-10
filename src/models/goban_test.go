package models

import (
	"testing"
)

func TestNewGoban(t *testing.T) {
	goban := NewGoban7()
	if goban.size != 7 {
		t.Errorf("expected size 7, got %d", goban.size)
	}
	if len(goban.dots) != 7 {
		t.Errorf("expected 7 rows, got %d", len(goban.dots))
	}
	for _, row := range goban.dots {
		if len(row) != 7 {
			t.Errorf("expected 7 columns, got %d", len(row))
		}
	}
}

func TestPlaceBlack(t *testing.T) {
	goban := NewGoban7()

	err := goban.PlaceBlack('D', 4)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	err = goban.PlaceBlack('D', 4)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	err = goban.PlaceWhite('D', 4)
	err = goban.PlaceBlack('A', 8)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	err = goban.PlaceBlack('A', 8)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	err = goban.PlaceBlack('Z', 8)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestPlaceWhite(t *testing.T) {
	goban := NewGoban7()

	err := goban.PlaceWhite('D', 4)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	err = goban.PlaceBlack('D', 4)
	err = goban.PlaceWhite('D', 4)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	err = goban.PlaceWhite('H', 8)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	err = goban.PlaceWhite('H', 8)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	err = goban.PlaceBlack('A', 1)
	err = goban.PlaceWhite('Z', 8)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestChangeTheme(t *testing.T) {
	goban := NewGoban7()
	newTheme := NewDarkGobanTheme()
	goban.ChangeTheme(newTheme)

	if goban.theme != *newTheme {
		t.Errorf("expected theme %v, got %v", newTheme, goban.theme)
	}
}

func TestLetterToNumber(t *testing.T) {
	goban := NewGoban7()

	tests := []struct {
		letter   rune
		expected uint8
		hasError bool
	}{
		{'A', 0, false},
		{'B', 1, false},
		{'G', 6, false},
		{'H', 0, true}, // Out of range for 7x7 goban
		{'a', 0, false},
		{'1', 0, true}, // Not a letter
	}

	for _, test := range tests {
		result, err := goban.letterToNumber(test.letter)
		if (err != nil) != test.hasError {
			t.Errorf("letterToNumber(%c) error = %v, expected error = %v", test.letter, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("letterToNumber(%c) = %d, expected %d", test.letter, result, test.expected)
		}
	}
}

func TestNewGoban9(t *testing.T) {
	goban := NewGoban9()
	if goban.size != 9 {
		t.Errorf("expected size 9, got %d", goban.size)
	}
	if len(goban.dots) != 9 {
		t.Errorf("expected 9 rows, got %d", len(goban.dots))
	}
	for _, row := range goban.dots {
		if len(row) != 9 {
			t.Errorf("expected 9 columns, got %d", len(row))
		}
	}
}

func TestNewGoban11(t *testing.T) {
	goban := NewGoban11()
	if goban.size != 11 {
		t.Errorf("expected size 11, got %d", goban.size)
	}
	if len(goban.dots) != 11 {
		t.Errorf("expected 11 rows, got %d", len(goban.dots))
	}
	for _, row := range goban.dots {
		if len(row) != 11 {
			t.Errorf("expected 11 columns, got %d", len(row))
		}
	}
}

func TestNewGoban13(t *testing.T) {
	goban := NewGoban13()
	if goban.size != 13 {
		t.Errorf("expected size 13, got %d", goban.size)
	}
	if len(goban.dots) != 13 {
		t.Errorf("expected 13 rows, got %d", len(goban.dots))
	}
	for _, row := range goban.dots {
		if len(row) != 13 {
			t.Errorf("expected 13 columns, got %d", len(row))
		}
	}
}

func TestNewGoban19(t *testing.T) {
	goban := NewGoban19()
	if goban.size != 19 {
		t.Errorf("expected size 19, got %d", goban.size)
	}
	if len(goban.dots) != 19 {
		t.Errorf("expected 19 rows, got %d", len(goban.dots))
	}
	for _, row := range goban.dots {
		if len(row) != 19 {
			t.Errorf("expected 19 columns, got %d", len(row))
		}
	}
}
