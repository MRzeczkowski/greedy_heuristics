package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func rastrigin(x []float64) float64 {
	n := len(x)
	sum := 10.0 * float64(n)
	for _, xi := range x {
		sum += (xi*xi - 10.0*math.Cos(2*math.Pi*xi))
	}
	return sum
}

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
	bestGlobalValue := math.Inf(1)

	for t := 0; t < maxStarts; t++ {
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
				local = true
			}
		}

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
	maxStdDev := 1.0

	for k <= maxIterations {
		stdDev := float64(k) * stdDevChangeRate
		if stdDev > maxStdDev {
			stdDev = maxStdDev
		}

		newSolution := generateNewSolutions(currentSolution, 1, stdDev)[0]
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

	dimensions := []int{3, 5}
	maxIterations := 1000
	numberOfSolutions := 10
	stdDevChangeRate := 0.1
	numberOfTests := 10

	fmt.Println("Running simulations...")
	fmt.Println("| Dimensions | Algorithm | Average Result | Average Time (ms) |")
	fmt.Println("|------------|-----------|----------------|-------------------|")

	// Iteracja przez różne wymiary zadania
	for _, dim := range dimensions {
		basicGreedySum, multiStartGreedySum, variableNeighborhoodGreedySum := 0.0, 0.0, 0.0
		basicGreedyTime, multiStartGreedyTime, variableNeighborhoodGreedyTime := time.Duration(0), time.Duration(0), time.Duration(0)

		for i := 0; i < numberOfTests; i++ {
			start := time.Now()
			solution1 := basicGreedy(dim, maxIterations, numberOfSolutions)
			result1 := rastrigin(solution1)
			basicGreedySum += result1
			basicGreedyTime += time.Since(start)

			start = time.Now()
			solution2 := multiStartGreedy(dim, maxIterations, numberOfSolutions)
			result2 := rastrigin(solution2)
			multiStartGreedySum += result2
			multiStartGreedyTime += time.Since(start)

			start = time.Now()
			solution3 := variableNeighborhoodGreedy(dim, maxIterations, stdDevChangeRate)
			result3 := rastrigin(solution3)
			variableNeighborhoodGreedySum += result3
			variableNeighborhoodGreedyTime += time.Since(start)
		}

		fmt.Printf("| %d          | Basic Greedy | %.4f         | %v                |\n", dim, basicGreedySum/float64(numberOfTests), basicGreedyTime.Milliseconds()/int64(numberOfTests))
		fmt.Printf("| %d          | Multi-Start Greedy | %.4f    | %v                |\n", dim, multiStartGreedySum/float64(numberOfTests), multiStartGreedyTime.Milliseconds()/int64(numberOfTests))
		fmt.Printf("| %d          | Variable Neighborhood Greedy | %.4f | %v           |\n", dim, variableNeighborhoodGreedySum/float64(numberOfTests), variableNeighborhoodGreedyTime.Milliseconds()/int64(numberOfTests))
	}
}
