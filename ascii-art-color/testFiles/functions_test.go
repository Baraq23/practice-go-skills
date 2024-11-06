package test

import (
	"asciiart/functions"
	"reflect"
	"testing"
)

// Func getSubstringPositions() returns a slice of intergers containing indexes where substring starts.
func TestGetSubStringPositions(t *testing.T) {

	tests  := []struct {
		MainStr string
		SubStr string
		Positions []int
	}{
		{"Hello, world", "Hello", []int{0}},
		{"Hey guys", "guys", []int{4}},
		{"Processed", "e", []int{4,7}},
	}


	for _, test := range tests {
		results := asciiart.GetSubStringPositions(test.MainStr, test.SubStr)
		if !reflect.DeepEqual(results, test.Positions) {
			t.Errorf("GetSubStringPositions(%q,%q) = %v, Expected = %v", test.MainStr, test.SubStr, results, test.Positions)
		}
	}
}



// func getAnsi() takes in the RGB color code provide the color name, then returns a formated ansi format
func TestGetansi(t *testing.T) {

	tests := []struct{
		color string
		ansi string
		ansi2 string
	}{
		{"red", "\033[38;2;255;0;0m", ""},
		{"orange","\033[38;2;255;165;0m", ""},
		{"blue,green", "\033[38;2;0;0;255m", "\033[38;2;0;255;0m"},
		//rgb
		{"rgb(255,0,0)", "\033[38;2;255;0;0m", ""},
		{"rgb(0,0,255)", "\033[38;2;0;0;255m", ""},
	}

	for _, test := range tests {
		ansi1, ansi2, _ := asciiart.GetAnsi(test.color)
		if ansi1 != test.ansi && ansi2 != test.ansi2 {
			t.Errorf("Getansi(%q) = %q %q. Expected = %q %q", test.color, ansi1, ansi2, test.ansi, test.ansi2)
		}
	}
}




// func TestStart(t *testing.T) {
// 	type test struct {
// 		name   string
// 		input  string
// 		output string
// 	}
// 	tests := []test{
// 		{"Testing for letter 'A'", "A", "           \n    /\\     \n   /  \\    \n  / /\\ \\   \n / ____ \\  \n/_/    \\_\\ \n           \n           "},
// 		{"Testing for number '7'", "7", "        \n _____  \n|___  | \n   / /  \n  / /   \n /_/    \n        \n        "},
// 		{"Testing for other characters '#'", "#", "   _  _    \n _| || |_  \n|_  __  _| \n _| || |_  \n|_  __  _| \n  |_||_|   \n           \n           "},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			if got := WordDistributer(tc.input, filepath.Join("..", "banners", "standard.txt")); got != tc.output {
// 				t.Errorf("Expected:\n %v \n Got:\n %v \n", tc.output, got)
// 			}
// 		})
// 	}
// }
