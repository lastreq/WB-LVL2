package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
func searchAnagram(s []string) *map[string][]string {
	var lowStrings []string
	for i := 0; i < len(s); i++ {
		lowStrings = append(lowStrings, strings.ToLower(s[i]))
	}

	tempMap := make(map[string][]string)
	result := make(map[string][]string)

	for _, val := range lowStrings {
		sorted := sortString(val)
		tempMap[sorted] = append(tempMap[sorted], val)
	}

	for _, v := range tempMap {
		if len(v) > 1 { //Множества из одного элемента не должны попасть в результат.
			result[v[0]] = v
			sort.Strings(v)
		}
	}
	return &result
}

func main() {
	massiv := []string{"Тяпка", "пятка", "пятак", "ток", "коТ", "Листок", "слитОк", "столик", "Яблоко", "КТО"}

	fmt.Println(searchAnagram(massiv))
}
