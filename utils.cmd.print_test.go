package utils

import "testing"

func TestPrintColor(t *testing.T) {
	ii := &Utils{}
	for i := 0; i < 15; i++ {
		ii.Print.Println(i, "嗨喽,Seaii！", "Color:", i)
	}
}

