package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var sun int
	var hp int
	var result int
	sun = 200
	hp = rand.Intn(1000)
	result = hp - sun
	{
		{
			// Инициализация генератора случайных чисел с использованием текущего времени в качестве источника
			rand.Seed(time.Now().UnixNano())

			// Список доступных действий
			actions := []string{"влево", "вправо", "вверх", "вниз"}

			// Генерация случайного индекса для выбора действия
			randomIndex := rand.Intn(len(actions))

			// Выбор случайного действия
			selectedAction := actions[randomIndex]

			fmt.Printf("Враг пошел %s\n", selectedAction)
			{
				if randomIndex == 0 {
					fmt.Println("Вы кинули санстрайк и попали во врага!")
					if result <= 0 {
						fmt.Printf("%d First Blood, у него было %d хп", result, hp)
					}
					if result > 0 {
						fmt.Printf("%d жизни у него осталось, Враг выжил, у врага было %d хп", result, hp)
					}
				} else {
					fmt.Println("Вы промахнулись санстрайком")
				}
			}
		}

	}
}
