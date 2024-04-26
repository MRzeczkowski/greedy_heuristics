# Greedy heuristics

This program is a simple demonstration of using greedy heuristic algorithms to find the minimum of the Rastrigin function, a non-convex function commonly used as a performance test problem for optimization algorithms. The program compares three different approaches: Basic Greedy, Multi-Start Greedy, and Variable Neighborhood Greedy, exploring their effectiveness in different dimensions.

Initial solutions for the algorithm are taken from the uniform distribution in range [-5.12, 5.12] since for these arguments the Rastrigin function has values for. Changes to solutions are made using the Cauchy distribution with hope that its heavy-tailed nature will allow for escaping local minima and find the global one.

## Algorithm Descriptions
### Basic
The Basic Greedy algorithm starts from a randomly selected initial point and iteratively searches for better solutions within a defined neighborhood. At each step, it evaluates multiple neighboring solutions and moves to the best one found, continuing this process until no further improvements are found. This method is straightforward but can easily become trapped in local minima.

### Multi-Start
The Multi-Start approach extends the basic greedy algorithm by performing multiple runs from different starting points. Each run operates independently, using the basic greedy approach, which helps to mitigate the risk of getting stuck in local minima by increasing the chances of exploring more of the solution space. The best result from all runs is selected as the final solution.

### Variable Neighborhood Greedy
The Variable Neighborhood algorithm introduces a dynamic element into the neighborhood search. It starts similar to the basic greedy but expands the search radius if no improvements are found within the current neighborhood. This expansion continues until a better solution is found or a maximum predetermined limit is reached. The ability to change the search radius helps to escape local minima and promotes a more global search of the problem space.

## Simulation results
Below is the output of the program, which includes simulation parameters and the results of each algorithm under various conditions:

Simulation Parameters:
- Number of dimensions: [3 5]
- Max iterations per test: 1000
- Number of solutions per iteration: 10
- Gamma change rate for Variable Neighborhood: 0.0001
- Number of tests per algorithm: 100

| Dimensions | Algorithm | Average Result | Average Time (ms) |
|-|-|-|-|
| 3 | Basic | 1.5294 | 3 |
| 3 | Multi-Start | 2.4637 | 21 |
| 3 | Variable Neighborhood | 1.2935 | 3 |
| 5 | Basic | 10.6019 | 4 |
| 5 | Multi-Start | 15.2964 | 36 |
| 5 | Variable Neighborhood | 4.7868 | 7 |

Unfortunately the global minima have not been found and the algorithm got stuck in a local minimum. The Variable Neighborhood faired best and this is probably because it was able to extend its search space and find a better solution. What is interesting is that it seems to react better to a smaller Gamma change rate - the bigger it was the worse the results and the value of 0.0001 was the best. Notably it was also very fast and much faster than the Multi-Start method.
