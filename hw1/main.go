package main

import (
	"fmt"
	"math"
)

//У Саши было 10 в степени 2 яблок, каждое из которых пронумеровано и лежит в мешке.
//Саша решил разложить все яблоки в ящики по номерам в порядке убывания по 10 яблок в каждом ящике.
//Полученные ящики Саша разделил на 2 машины. Ящики с нечётными номерами, Саша загрузил в первую машину, а с чётными во вторую.
//
//
//Напишите код используя функции или методы.
//Используя пакет `fmt`, вам необхоидмо вывести номера яблок в следующем формате:
//`Машина: <номер>, Ящик: <номер>, Яблоки: [<номер>,<номер>,<номер>,...]`

func printFormatted(crates [][]int) {
	carIdx := 1
	for crateIdx := 0; crateIdx < len(crates); crateIdx++ {
		apples := crates[crateIdx]
		fmt.Printf("Машина: %d, Ящик: %d, Яблоки: %v\n", carIdx, crateIdx+1, apples)
		if carIdx == 1 {
			carIdx = 2
		} else {
			carIdx = 1
		}
	}

}

func solve(base, degree int) [][]int {
	apples := int(math.Pow(float64(base), float64(degree)))

	var crates [][]int
	size := min(10, apples)
	crates = append(crates, make([]int, size))

	currentCrateCapacity := 10
	cratesIdx := 0

	for appleIdx := apples; appleIdx > 0; appleIdx-- {
		applesCrateIdx := 10 - currentCrateCapacity

		crates[cratesIdx][applesCrateIdx] = appleIdx
		currentCrateCapacity--

		if currentCrateCapacity == 0 {
			cratesIdx++
			currentCrateCapacity = 10

			nextCrateSize := min(10, appleIdx-1)
			crates = append(crates, make([]int, nextCrateSize))
		}
	}

	return crates
}

func main() {
	printFormatted(solve(2, 10))
}
