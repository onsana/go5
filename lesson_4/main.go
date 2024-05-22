package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

type Model struct {
	Id int
}

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
	var nums = []Model{}
	for i := 0; i < 10; i++ {
		nums = append(nums, Model{
			Id: rand.Intn(10),
		})
	}
	unique := removeDuplicates(nums)
	sort.Slice(unique, func(i, j int) bool {
		return unique[i].Id < unique[j].Id
	})
	fmt.Println(unique)
}

func removeDuplicates(slice []Model) []Model {
	result := []Model{}
	for i := 0; i < len(slice); i++ {
		duplicate := false
		for j := 0; j < len(result); j++ {
			if slice[i].Id == result[j].Id {
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
