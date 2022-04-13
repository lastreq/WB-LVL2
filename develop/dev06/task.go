package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	fields    string
	delimiter string
	separated bool
}

func newDelimiter(s []string, d string) [][]string {
	var result [][]string
	for _, val := range s {
		result = append(result, strings.Split(val, d))
	}
	return result
}

func separatedStrings(s [][]string) [][]string {
	var result [][]string
	for _, val := range s {
		if len(val) > 1 {
			result = append(result, val)
		}
	}
	return result
}

func printStrings(s [][]string, arrayFields []int, delim string) {

	m := make(map[int]bool, 0)
	for i := range arrayFields {
		m[arrayFields[i]] = true
	}

	for i := range s {
		f := true //первый столбик
		for j := range s[i] {
			if m[j] || (arrayFields[0]) == -1 { //если столбцы не указаны - печатаем все
				if f {
					f = false
				} else {

					fmt.Print(delim) //разделитель перед непервым столбцов
				}

				fmt.Print(s[i][j])
			}
		}
		if !f { //если что-то было напечатано в строке
			fmt.Println(" newString") //новая строка
		}
	}
}

func main() {
	f := flag.String("f", "-1", `"fields" - выбрать поля (колонки)`)
	d := flag.String("d", `\t`, `"delimiter" - использовать другой разделитель`)
	s := flag.Bool("s", false, `"separated" - только строки с разделителем`)
	flag.Parse()
	flagsStruct := Flags{fields: *f, delimiter: *d, separated: *s}
	var arrayFields []int
	var arrayTemp []string
	arrayTemp = strings.Split(flagsStruct.fields, ",")
	for _, val := range arrayTemp {
		v, _ := strconv.Atoi(val)
		arrayFields = append(arrayFields, v)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {

		ok := scanner.Scan()
		if !ok && scanner.Err() == nil {
			break
		}

		text := scanner.Text()
		lines := strings.Split(text, "\n")

		result := newDelimiter(lines, flagsStruct.delimiter)
		if flagsStruct.separated {
			result = separatedStrings(result)
		}

		printStrings(result, arrayFields, flagsStruct.delimiter)
	}
}
