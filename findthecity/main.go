package main

/*
Problem description:
    There are n cities numbered from 0 to n - 1. 
    Given the array edges where edges[i] = [fromi, toi, weighti] represents a bidirectional and weighted edge between cities fromi and toi, and given the integer distanceThreshold.

    Return the city with the smallest number of cities that are reachable through some path and whose distance is at most distanceThreshold.
    If there are multiple such cities, return the city with the greatest number.

    Notice that the distance of a path connecting cities i and j is equal to the sum of the eges' weights along that path.
Constraints: 
    2 <= n <= 100
    1 <= edges.length <= n * (n - 1) / 2
    edges[i].length == 3
    0 <= fromi < toi < n
    1 <= weighti, distanceThreshold <= 10^4
    All pairs (fromi, toi) are distinct.
*/
type City struct {
    number int
    neighbors []*City
    weight int
}

func (c City) AddNeighbor(n *City) {
    c.neighbors = append(c.neighbors, n)
}

func NewCity(number int) *City {
    return &City{
        number: number,
        neighbors: make([]*City, 0),
    }
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
    cities := make([]*City, 0)
    for i := 0; i < n; i++ {
        cities[i] = NewCity(i)
    }

    for i := 0; i < n; i++ {
        from, to, weight := edges[i][0], edges[i][1], edges[i][2]
        fromCity := cities[from]
        toCity := cities[to]
        fromCity.AddNeighbor(toCity)
        toCity.AddNeighbor(fromCity)
    }

    return n
}
