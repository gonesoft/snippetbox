package main

import (
	"fmt"
	"math"
)

func AverageCalculator(filterFunc func(float64) bool, transformFunc func(float64) float64) func([]float64) float64 {
	return func(numbers []float64) float64 {
		filteredNumbers := make([]float64, 0, len(numbers))
		for _, number := range numbers {
			if filterFunc(number) {
				filteredNumbers = append(filteredNumbers, transformFunc(number))
			}
		}
		return CalculateAverage(filteredNumbers)
	}
}

func CalculateAverage(numbers []float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	evenFilter := func(number float64) bool {
		return math.Mod(number, 2) == 0
	}

	evenAverageCalculator := AverageCalculator(evenFilter, func(number float64) float64 {
		return number
	})

	numbers := []float64{1, 2, 3, 4, 5, 6}
	evenAverage := evenAverageCalculator(numbers)
	fmt.Println(evenAverage) // Output: 3
}
