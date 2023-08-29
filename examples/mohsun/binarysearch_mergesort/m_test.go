package main

import (
	"testing"
)

// nece print edek testde
func TestSimpleCheck(t *testing.T) {
	//eslinde yuoxlamaliyiq

	yourslc := []int{3, 4, 2, 15, 788, 92, 4, 1, 5, 6, 16}
	mylice := []int{3, 4, 2, 888, 4545, 666, 44444, 1, 5, 6}
	t.Log(selectionsort(mylice))
	wordList := []string{"alma", "almaci", "almadi", "almali", "almasi", "almayin", "banan", "barama"}
	letters := []string{"alma", "almaci", "almadi", "almali", "almasi", "almayin", "banan", "barama"}
	t.Log(binarysearch("almayin", wordList))
	f := iterBinary(wordList, "barama", false)
	if f != 7 {
		t.Fatalf("Sehv ishledi", f, 7)
	}
	t.Log(iterBinary(wordList, "barama", true))
	t.Log(mergeSort(yourslc))
	t.Log(mylice)
	fst, end := findWordsPrefix(letters, "alm", 10)

	if fst != 0 || end != 6 {
		//log demeli shev ishleyende olur
		t.Fatalf("Sehv ishledi")
	}
}
