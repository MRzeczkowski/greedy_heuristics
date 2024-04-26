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

// Generuje nowe rozwiązania przez perturbację obecnego punktu
func generateNewSolutions(x0 []float64, numberOfSolutions int, stdDev float64) [][]float64 {
	n := len(x0)
	solutions := make([][]float64, numberOfSolutions)
	for i := range solutions {
		solutions[i] = make([]float64, n)
		for j := range solutions[i] {
			shift := rand.NormFloat64() * stdDev
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

func drawInitialSolution(dimensions int) []float64 {
	initialSolution := make([]float64, dimensions)
	for i := range initialSolution {
		initialSolution[i] = rand.Float64()*10 - 5
	}

	return initialSolution
}

func basicGreedy(dimensions int, maxIterations int, numberOfSolutions int) []float64 {

	x0 := drawInitialSolution(dimensions)

	currentValue := rastrigin(x0)
	currentSolution := x0

	for k := 0; k < maxIterations; k++ {
		newSolutions := generateNewSolutions(currentSolution, numberOfSolutions, 1.0)
		bestNewSolution := findBestSolution(newSolutions)
		bestNewValue := rastrigin(bestNewSolution)

		if bestNewValue < currentValue {
			currentSolution = bestNewSolution
			currentValue = bestNewValue
		}
	}
	return currentSolution
}

func multiStartGreedy(dimensions int, maxStarts int, numberOfSolutions int) []float64 {
	bestGlobalSolution := make([]float64, dimensions)
	bestGlobalValue := math.Inf(1) // Nieskończoność, szukamy minimum

	// Powtarzamy algorytm z różnych początkowych punktów
	for t := 0; t < maxStarts; t++ {

		// Wybierz losowe rozwiązanie początkowe
		currentSolution := drawInitialSolution(dimensions)
		currentValue := rastrigin(currentSolution)
		local := false

		for !local {
			newSolutions := generateNewSolutions(currentSolution, numberOfSolutions, 1.0)
			bestNewSolution := findBestSolution(newSolutions)
			bestNewValue := rastrigin(bestNewSolution)

			if bestNewValue < currentValue {
				currentSolution = bestNewSolution
				currentValue = bestNewValue
			} else {
				local = true // Nie znaleziono lepszego lokalnego rozwiązania
			}
		}

		// Sprawdzenie, czy obecne rozwiązanie jest najlepsze globalnie
		if currentValue < bestGlobalValue {
			bestGlobalValue = currentValue
			bestGlobalSolution = currentSolution
		}
	}

	return bestGlobalSolution
}

func variableNeighborhoodGreedy(dimensions int, maxIterations int, stdDevChangeRate float64) []float64 {

	x0 := drawInitialSolution(dimensions)
	currentValue := rastrigin(x0)
	currentSolution := x0
	k := 1

	for k < maxIterations {
		newSolution := generateNewSolutions(currentSolution, 1, float64(k)*stdDevChangeRate)[0]
		newValue := rastrigin(newSolution)

		if newValue < currentValue {
			currentSolution = newSolution
			currentValue = newValue
			k = 1
		} else {
			k++
		}
	}
	return currentSolution
}

func main() {

	dimensions := 3
	maxIterations := 1000
	numberOfSolutions := 100
	stdDevChangeRate := 0.1

	// Test podstawowego algorytmu
	solution1 := basicGreedy(dimensions, maxIterations, numberOfSolutions)
	fmt.Println("Basic Greedy Result:", solution1, "=", rastrigin(solution1))

	// Test wielostartowego algorytmu
	solution2 := multiStartGreedy(dimensions, maxIterations, numberOfSolutions)
	fmt.Println("Multi-Start Greedy Result:", solution2, "=", rastrigin(solution2))

	// Test algorytmu ze zmiennym sąsiedztwem
	solution3 := variableNeighborhoodGreedy(dimensions, maxIterations, stdDevChangeRate)
	fmt.Println("Variable Neighborhood Greedy Result:", solution3, "=", rastrigin(solution3))
}
