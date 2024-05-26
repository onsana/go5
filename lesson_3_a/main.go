package main

import (
	"fmt"
)

type Player struct {
	Health int
	Items  []string
}

type Location struct {
	Description string
	Options     map[string]string
	Effects     map[string]func(*Player)
}

func (p *Player) ChangeHealth(amount int) {
	p.Health += amount
	if p.Health < 0 {
		p.Health = 0
	}
}

func (p *Player) AddItem(item string) {
	p.Items = append(p.Items, item)
}

func NewLocation(description string, options map[string]string, effects map[string]func(*Player)) *Location {
	return &Location{Description: description, Options: options, Effects: effects}
}

func initializeGame() (*Player, map[string]*Location) {
	player := &Player{
		Health: 100,
		Items:  []string{},
	}

	locations := make(map[string]*Location)

	// Початкова локація
	locations["start"] = NewLocation(
		"Ви прокинулись у невідомому місці з кількома речами. Ви нічого не пам'ятаєте.",
		map[string]string{
			"1": "Піти ліворуч",
			"2": "Піти праворуч",
		},
		map[string]func(*Player){
			"1": func(p *Player) {
				fmt.Println("Ви йдете ліворуч і знаходите яблуко. Здоров'я +10.")
				p.ChangeHealth(10)
				p.AddItem("яблуко")
			},
			"2": func(p *Player) {
				fmt.Println("Ви йдете праворуч і наступаєте на капкан. Здоров'я -20.")
				p.ChangeHealth(-20)
			},
		},
	)

	// Лівий шлях
	locations["left"] = NewLocation(
		"Ви йдете ліворуч. Перед вами ліс.",
		map[string]string{
			"1": "Увійти до лісу",
			"2": "Повернутися назад",
		},
		map[string]func(*Player){
			"1": func(p *Player) {
				fmt.Println("Ви входите до лісу і бачите ведмедя. Здоров'я -30.")
				p.ChangeHealth(-30)
			},
			"2": nil,
		},
	)

	// Правий шлях
	locations["right"] = NewLocation(
		"Ви йдете праворуч. Перед вами річка.",
		map[string]string{
			"1": "Переплисти річку",
			"2": "Повернутися назад",
		},
		map[string]func(*Player){
			"1": func(p *Player) {
				fmt.Println("Ви перепливаєте річку, але течія сильна. Здоров'я -10.")
				p.ChangeHealth(-10)
			},
			"2": nil,
		},
	)

	return player, locations
}

func main() {
	player, locations := initializeGame()

	currentLocation := "start"

	for {
		loc := locations[currentLocation]
		fmt.Println(loc.Description)
		for option, description := range loc.Options {
			fmt.Printf("%s: %s\n", option, description)
		}

		var choice string
		fmt.Print("Ваш вибір: ")
		fmt.Scan(&choice)

		if effect, ok := loc.Effects[choice]; ok && effect != nil {
			effect(player)
		}

		if _, ok := loc.Options[choice]; ok {
			switch currentLocation {
			case "start":
				if choice == "1" {
					currentLocation = "left"
				} else if choice == "2" {
					currentLocation = "right"
				}
			case "left", "right":
				if choice == "2" {
					currentLocation = "start"
				}
			}
		} else {
			fmt.Println("Невірний вибір, спробуйте ще раз.")
		}

		fmt.Printf("Ваше здоров'я: %d\n", player.Health)
		if player.Health <= 0 {
			fmt.Println("Ви загинули. Гра завершена.")
			break
		}
	}
}
