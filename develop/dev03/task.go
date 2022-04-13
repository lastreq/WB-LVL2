package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortMain(s []string) []string {
	sort.Strings(s)
	return s
}
func sortK(s []string, column int) ([]string, error) {
	map1 := make(map[string]string)
	var arr, arr1 []string
	for _, v := range s {
		a := strings.Split(v, " ")
		if column > len(a) || column < 1 {
			return nil, fmt.Errorf("index out of range [%d] with length %d", column, len(a))
		}
		map1[a[column-1]] = v
		arr = append(arr, a[column-1])
	}
	sort.Strings(arr)
	for _, v := range arr {
		arr1 = append(arr1, map1[v])
	}
	return arr1, nil
}
func sortN(s []string) ([]string, error) {
	var resultStringSlice []string
	var err error
	for _, val := range s {
		var sliceByte []string
		var tempSlice []int
		var tempInt int
		sliceByte = strings.Split(val, " ")
		for _, val2 := range sliceByte {
			tempInt, err = strconv.Atoi(val2)
			tempSlice = append(tempSlice, tempInt)
		}
		sort.Ints(tempSlice)
		resultStringSlice = append(resultStringSlice, strings.Trim(strings.Join(strings.Split(fmt.Sprint(tempSlice), " "), ","), "[]"))
	}
	return resultStringSlice, err
}
func sortR(s []string) ([]string, error) {
	sort.Strings(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s, nil
}
func sortU(s []string) ([]string, error) {
	map1 := make(map[string]int)
	var slice []string
	for _, v := range s {
		map1[v]++
		if map1[v] == 1 {
			slice = append(slice, v)
		}
	}
	sort.Strings(slice)

	return slice, nil
}
func printSortedStrings(s []string) {
	for _, a := range s {
		fmt.Println(a)
	}
}
func openFile(s string) ([]string, error) {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}
	aa := strings.ReplaceAll(string(data), "\r", "")
	a := strings.Split(aa, "\n")
	for i, v := range a {
		runes := []rune(v)
		if runes[0] == 65279 {
			a[i] = string(runes[1:])
		}
	}
	return a, nil
}

func main() {
	k := flag.Int("k", 0, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	var unsortedStrings, sortedStrings []string
	var err error

	flag.Parse()
	if filename := flag.Arg(0); filename != "" {
		unsortedStrings, err = openFile(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
	}

	if *k != 0 {
		sortedStrings, err = sortK(unsortedStrings, *k)
	}
	if *n {
		sortedStrings, err = sortN(unsortedStrings)
	}
	if *r {
		sortedStrings, err = sortR(unsortedStrings)
	}
	if *u {
		sortedStrings, err = sortU(unsortedStrings)
	}

	printSortedStrings(sortedStrings)
}
