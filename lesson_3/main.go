package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Name     string
	Location string
	Backpack []string
	Health   int
}

type Location struct {
	Description string
	Choices     map[string]string
}

func main() {
	player := Player{
		Name:     "Стівен",
		Location: "печера",
		Backpack: []string{"сірники", "ліхтарик", "ніж"},
		Health:   100,
	}

	locations := map[string]Location{
		"печера": {
			Description: "Стівен прокинувся біля входу в печеру. Він лише памʼятає своє імʼя. Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж. У печері темно. ",
			Choices: map[string]string{
				"йти до лісу": "ліс",
			},
		},
		"ліс": {
			Description: "Стівен іде стежкою, яка веде від печери в ліс. У лісі Стівен натикається на мертве тіло дивної тварини.",
			Choices: map[string]string{
				"нічого не робити":  "табір",
				"дослідити тварину": "обід",
			},
		},
		"обід": {
			Description: "Стівен готує обід з мʼяса тварини",
			Choices: map[string]string{
				"відпочити": "табір",
			},
		},
		"табір": {
			Description: "Через деякий час Стівен приходить до безлюдного табору. Він вже втомлений і вирішує відпочити, а не йти далі.",
			Choices: map[string]string{
				"відпочити": "сейф",
			},
		},
		"сейф": {
			Description: "У найближчому наметі він знаходить сейф з кодовим замком з двох чисел.",
			Choices: map[string]string{
				"спробувати код": "кінець",
			},
		},
		"кінець": {
			Description: "Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає. Стівен непритомніє. Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає. Стівен непритомніє.",
			Choices:     map[string]string{},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		currentLocation := locations[player.Location]
		fmt.Println(currentLocation.Description)
		if len(currentLocation.Choices) == 0 {
			fmt.Println("Гра закінчена.")
			break
		}

		fmt.Println("Ваші варіанти:")
		for choice := range currentLocation.Choices {
			fmt.Println("- " + choice)
		}

		fmt.Print("Введіть ваш вибір: ")
		scanner.Scan()
		choice := strings.ToLower(scanner.Text())

		if nextLocation, exists := currentLocation.Choices[choice]; exists {
			player.Location = nextLocation
		} else {
			fmt.Println("Невірний вибір, спробуйте ще раз.")
		}
	}
}
