package main

import (
	"reflect"
	"testing"
)

func Test_searchAnagram(t *testing.T) {
	itogMap := make(map[string][]string)
	a1 := []string{"кот", "кто", "ток"}
	a2 := []string{"листок", "слиток", "столик"}
	a3 := []string{"пятак", "пятка", "тяпка"}
	itogMap["ток"] = a1
	itogMap["листок"] = a2
	itogMap["тяпка"] = a3

	t.Run("test", func(t *testing.T) {
		massiv := []string{"Тяпка", "пятка", "пятак", "ток", "коТ", "Листок", "слитОк", "столик", "Яблоко", "КТО"}
		gotMap := searchAnagram(massiv)
		for o, oo := range *gotMap {
			if !reflect.DeepEqual(oo, itogMap[o]) {
				t.Errorf("anagramma() = %v, want %v", gotMap, itogMap)
			}
		}
	})
}
