package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func isMatch(b []byte, pattern string) (bool, error) {
	return regexp.Match(pattern, b)
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
func grepA(arr []string, n int, pattern string, sliceFlags [4]bool) []string {
	var Match bool
	var result []string
	for i := 0; i < len(arr); i++ {
		if sliceFlags[0] { //invert
			Match, _ = isMatch([]byte(strings.ToLower(arr[i])), strings.ToLower(pattern))
		} else {
			Match, _ = isMatch([]byte(arr[i]), pattern)
		}
		if !sliceFlags[1] {
			if Match {
				fmt.Printf("Найдено совпадения на строке:%v\nПечать %v строк(и) после\n", arr[i], n)
				for j, stringsToPrint := i, n; j < len(arr) && stringsToPrint > 0; {
					stringsToPrint--
					j++
					if !sliceFlags[3] {
						result = append(result, arr[j])
						fmt.Println(arr[j])
					} else {
						result = append(result, strconv.Itoa(j))
						fmt.Println(j)
					}

				}
				break
			}
		} else {
			if Match {
			} else {
				if !sliceFlags[3] {
					result = append(result, arr[i])
					fmt.Println(arr[i])
				} else {
					result = append(result, strconv.Itoa(i))
					fmt.Println(i)
				}
			}
		}

	}
	return result
}
func grepB(arr []string, n int, pattern string, sliceFlags [4]bool) []string {
	var Match bool
	var result []string
	for i := 0; i < len(arr); i++ {
		if sliceFlags[0] { //invert
			Match, _ = isMatch([]byte(strings.ToLower(arr[i])), strings.ToLower(pattern))
		} else {
			Match, _ = isMatch([]byte(arr[i]), pattern)
		}
		if !sliceFlags[1] {
			if Match {
				fmt.Printf("Найдено совпадения на строке:%v\nПечать %v строк(и) до\n", arr[i], n)
				for j, stringsToPrint := i, n; j > 0 && stringsToPrint > 0; {
					stringsToPrint--
					j--
					if !sliceFlags[3] {
						result = append(result, arr[j])
						fmt.Println(arr[j])
					} else {
						result = append(result, strconv.Itoa(j))
						fmt.Println(j)
					}

				}
				break
			}
		} else {
			if Match {
			} else {
				if !sliceFlags[3] {
					result = append(result, arr[i])
					fmt.Println(arr[i])
				} else {
					result = append(result, strconv.Itoa(i))
					fmt.Println(i)
				}
			}
		}

	}
	return result
}
func grepC(arr []string, n int, pattern string, sliceFlags [4]bool) []string {
	var result []string
	var Match bool
	for i := 0; i < len(arr); i++ {
		if sliceFlags[0] { //invert
			Match, _ = isMatch([]byte(strings.ToLower(arr[i])), strings.ToLower(pattern))
		} else {
			Match, _ = isMatch([]byte(arr[i]), pattern)
		}
		if !sliceFlags[1] {
			if Match {
				fmt.Printf("Найдено совпадения на строке:%v\nПечать %v строк(и) после\n", arr[i], n)
				for j, stringsToPrint := i, n; j < len(arr) && stringsToPrint > 0; {
					stringsToPrint--
					j++
					if !sliceFlags[3] {
						result = append(result, arr[j])
						fmt.Println(arr[j])
					} else {
						result = append(result, strconv.Itoa(j))
						fmt.Println(j)
					}

				}
				break
			}
		} else {
			if Match {
			} else {
				if !sliceFlags[3] {
					result = append(result, arr[i])
					fmt.Println(arr[i])
				} else {
					result = append(result, strconv.Itoa(i))
					fmt.Println(i)
				}
			}

		}

	}

	for i := 0; i < len(arr); i++ {
		if sliceFlags[0] { //ignore Case
			Match, _ = isMatch([]byte(strings.ToLower(arr[i])), strings.ToLower(pattern))
		} else {
			Match, _ = isMatch([]byte(arr[i]), pattern)
		}
		if !sliceFlags[1] {
			if Match {
				fmt.Printf("Печать %v строк(и) до\n", n)
				for j, stringsToPrint := i, n; j > 0 && stringsToPrint > 0; {
					stringsToPrint--
					j--
					if !sliceFlags[3] {
						result = append(result, arr[j])
						fmt.Println(arr[j])
					} else {
						result = append(result, strconv.Itoa(j))
						fmt.Println(j)
					}

				}
				break
			}
		}

	}
	return result
}
func grepCount(arr []string, pattern string) {
	var count int
	for i := 0; i < len(arr); i++ {
		Match, _ := isMatch([]byte(arr[i]), pattern)

		if Match {
			count++
		}

	}
	fmt.Println("Количество строк ", count)
}
// go run task.go -A 2 Go test.txt
func main() {
	A := flag.Int("A", 0, "\"after\" печатать +N строк после совпадения")
	B := flag.Int("B", 0, "\"before\" печатать +N строк до совпадения")
	C := flag.Int("C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	Count := flag.Bool("Count", false, "\"count\" (количество строк)")
	i := flag.Bool("i", false, "\"ignore-case\" (игнорировать регистр)")
	v := flag.Bool("v", false, "\"invert\" (вместо совпадения, исключать)")
	f := flag.Bool("f", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "\"line num\", печатать номер строки")

	var unsortedStrings []string
	var err error
	var sliceFlags [4]bool
	flag.Parse()

	regExpr := flag.Arg(0)
	if regExpr == "" {
		fmt.Println("No regExpr")
		os.Exit(1)
	}
	if filename := flag.Arg(1); filename != "" {
		unsortedStrings, err = openFile(filename)
		if err != nil {
			fmt.Println("Error opening file: err:", err)
			os.Exit(1)
		}
	}

	if *i != false {
		sliceFlags[0] = *i
	}
	if *v != false {
		sliceFlags[1] = *v
	}
	if *f != false {
		sliceFlags[2] = *f
	}
	if *n != false {
		sliceFlags[3] = *n
	}
	if *A != 0 {

		grepA(unsortedStrings, *A, regExpr, sliceFlags)
	}

	if *B != 0 {

		grepB(unsortedStrings, *B, regExpr, sliceFlags)
	}
	if *C != 0 {

		grepC(unsortedStrings, *C, regExpr, sliceFlags)
	}
	if *Count != false {

		grepCount(unsortedStrings, regExpr)
	}

}
