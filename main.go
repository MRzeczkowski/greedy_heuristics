package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Funkcja Rastrigina dla wektora x
func rastrigin(x []float64) float64 {
	n := len(x)
	sum := 10.0 * float64(n)
	for _, xi := range x {
		sum += (xi*xi - 10.0*math.Cos(2*math.Pi*xi))
	}
	return sum
}

// Generuje nowe rozwiązania wokół obecnego punktu
func generateNewSolutions(x0 []float64, numberOfSolutions int, stepSize float64) [][]float64 {
	n := len(x0)
	solutions := make([][]float64, numberOfSolutions)
	for i := range solutions {
		solutions[i] = make([]float64, n)
		for j := range solutions[i] {
			shift := rand.Float64()*2*stepSize - stepSize
			solutions[i][j] = x0[j] + shift
		}
	}
	return solutions
}

// Znajduje najlepsze rozwiązanie z listy
func findBestSolution(solutions [][]float64) []float64 {
	bestSolution := solutions[0]
	bestValue := rastrigin(bestSolution)
	for _, solution := range solutions {
		currentValue := rastrigin(solution)
		if currentValue < bestValue {
			bestValue = currentValue
			bestSolution = solution
		}
	}
	return bestSolution
}

func main() {

	// Parametry algorytmu
	dimensions := 5
	x0 := make([]float64, dimensions)
	for i := range x0 {
		x0[i] = rand.Float64()*10 - 5 // Losowe rozwiązanie początkowe w każdym wymiarze
	}
	maxIterations := 1000
	numberOfSolutions := 10
	stepSize := 0.1

	currentSolution := x0
	currentValue := rastrigin(currentSolution)
	fmt.Printf("Początkowe rozwiązanie: %v, wartość: %f\n", currentSolution, currentValue)

	// Główna pętla algorytmu
	for k := 0; k < maxIterations; k++ {
		newSolutions := generateNewSolutions(currentSolution, numberOfSolutions, stepSize)
		bestNewSolution := findBestSolution(newSolutions)
		bestNewValue := rastrigin(bestNewSolution)

		if bestNewValue < currentValue {
			currentSolution = bestNewSolution
			currentValue = bestNewValue
		}

		// Zmniejszanie rozmiaru kroku może pomóc w precyzyjniejszym poszukiwaniu
		stepSize *= 0.95
	}

	fmt.Printf("Znalezione najlepsze rozwiązanie: %v, wartość: %f\n", currentSolution, currentValue)
}
