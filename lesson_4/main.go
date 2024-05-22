package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// first
	// Ой у лузі червона калина похилилася,
	// Чогось наша славна Україна зажурилася.
	// А ми тую червону калину підіймемо,
	// А ми нашу славну Україну, гей, гей, розвеселимо!
	scanner := bufio.NewScanner(os.Stdin)
	redactor := []string{}
	for i := 0; i < 4; i++ {
		fmt.Print("Введіть новий рядок: ")
		scanner.Scan()
		redactor = append(redactor, scanner.Text())
	}
	text := "А ми тую червону калину підіймемо,"
	for _, u := range redactor {
		fmt.Println("Текст користувача ", u)
		if u == text {
			fmt.Println("Ураааа, нарешті знайшли")
		} else {
			fmt.Println("На жаль, знову не то")
		}
	}
	// second
	nums := []int{1, 5, 3, 2, 4, 3, 5, 4, 2}
	unique := removeDuplicates(nums)
	sort.Ints(unique)
	fmt.Println(unique)
}

func removeDuplicates(slice []int) []int {
	result := []int{}
	for i := 0; i < len(slice); i++ {
		duplicate := false
		for j := 0; j < len(result); j++ {
			if slice[i] == result[j] {
				duplicate = true
				break
			}
		}
		if !duplicate {
			result = append(result, slice[i])
		}
	}
	return result
}
