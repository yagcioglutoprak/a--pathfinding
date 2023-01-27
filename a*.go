package main

import (
    "fmt"
    "math"
)

type Node struct {
    x int
    y int
    g int
    h int
    parent *Node
}

func (n *Node) fValue() int {
    return n.g + n.h
}

func euclidean(start, goal *Node) int {
    x := (start.x - goal.x) * (start.x - goal.x)
    y := (start.y - goal.y) * (start.y - goal.y)
    return int(math.Sqrt(float64(x + y)))
}

func aStar(start, goal *Node, grid [][]int) []*Node {
    open := []*Node{start}
    closed := []*Node{}

    for len(open) > 0 {
        current := open[0]
        open = open[1:]

        closed = append(closed, current)

        if current.x == goal.x && current.y == goal.y {
            path := []*Node{}
            for current != nil {
                path = append([]*Node{current}, path...)
                current = current.parent
            }
            return path
        }

        for x := -1; x <= 1; x++ {
            for y := -1; y <= 1; y++ {
                if x == 0 && y == 0 {
                    continue
                }

                neighbor := &Node{x: current.x + x, y: current.y + y}

                if neighbor.x < 0 || neighbor.x >= len(grid) || neighbor.y < 0 || neighbor.y >= len(grid[0]) {
                    continue
                }

                if grid[neighbor.x][neighbor.y] == 1 {
                    continue
                }

                if containsNode(closed, neighbor) {
                    continue
                }

                tentativeG := current.g + 1

                if !containsNode(open, neighbor) {
                    neighbor.g = tentativeG
                    neighbor.h = euclidean(neighbor, goal)
                    neighbor.parent = current
                    open = append(open, neighbor)
                } else if tentativeG < neighbor.g {
                    neighbor.g = tentativeG
                    neighbor.parent = current
                }
            }
        }

        open = sortNodesByFValue(open)
    }

    return []*Node{}
}

func containsNode(nodes []*Node, node *Node) bool {
    for _, n := range nodes {
        if n.x == node.x && n.y == node.y {
            return true
        }
    }
    return false
}

func sortNodesByFValue(nodes []*Node) []*Node {
    if len(nodes) == 0 {
        return []*Node{}
    }
    pivot := nodes[0]
    less := []*Node{}
    greater := []*Node{}

    for i := 1; i < len(nodes); i++ {
        if nodes[i].fValue() < pivot.fValue() {
            less = append(less, nodes[i])
        } else {
            greater = append(greater, nodes[i])
        }
    }

    return append(append(sortNodesByFValue(less), pivot), sortNodesByFValue(greater)...)
}

func main() {
    start := &Node{x: 0, y: 0}
    goal := &Node{x: 5, y: 5}

    grid := [][]int{
        {0, 1, 0, 0, 0, 0},
        {0, 1, 0, 1, 0, 0},
        {0, 1, 0, 1, 0, 0},
        {0, 0, 0, 1, 0, 0},
        {0, 0, 0, 0, 0, 0},
        {0, 0, 0, 0, 0, 0},
    }

    path := aStar(start, goal, grid)

    fmt.Println("Shortest Path:")
    for _, node := range path {
        fmt.Printf("(%d, %d) ", node.x, node.y)
    }
}
