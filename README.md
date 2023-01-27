## A* algorithm in Go

This is an implementation of the A* algorithm in Go, which is used to find the shortest path in a warehouse environment. The algorithm considers the X and Y coordinates of each grid cell, and uses Euclidean distance as the heuristic function.

### Usage

1. Clone this repository
```git clone https://github.com/YOUR_USERNAME/A-Star-Algorithm-in-Go.git```
2. Go to the directory
```cd A-Star-Algorithm-in-Go```
3. Build and run
```go build```
```./A-Star-Algorithm-in-Go```



### Inputs
- The starting point  coordinates 
- The goal point  coordinates 
- The grid of obstacles, represented by a 2-dimensional array of integers where 0 represents a valid point and 1 represents an obstacle.

### Output
- A slice of nodes containing the shortest path from start to goal.

### Customization
- You can modify the grid and heuristic function to fit your specific use case.
- You can also write test functions to check the functionality of the algorithm

### Contributing

If you wish to contribute to this project, please fork the repository, make your changes and submit a pull request.

### License

This project is licensed under the MIT license. Feel free to use and modify the code as you wish.

### References
- [A* Search Algorithm](https://en.wikipedia.org/wiki/A*_search_algorithm)
- [Motion Planning](https://en.wikipedia.org/wiki/Motion_planning)
