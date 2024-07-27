package main

import "fmt"

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

type Neighbor struct {
	City   *City
	Weight int
}

func (n Neighbor) String() string {
    return fmt.Sprintf("City %d:%d", n.City.Number, n.Weight)
}

type City struct {
	Number    int
	neighbors []Neighbor
}

func (c *City) String() string {
    if c == nil {
        return "nil"
    }

    return fmt.Sprintf("City %d -> %v\n", c.Number, c.neighbors)
}

func (c *City) NumNeighbors() int {
    if c == nil {
        return 0
    }

    return len(c.neighbors)
}

func (c *City) AddNeighbor(n *City, weight int) {
	c.neighbors = append(c.neighbors, Neighbor{
		City:   n,
		Weight: weight,
	})
}

func (c *City) Neighbors() []Neighbor {
    if c == nil {
        return nil
    }

    return c.neighbors
}

func (c *City) HasNeighbor(cn int) bool {
    if c == nil {
        return false
    }

    if c.Number == cn {
        return true
    }

    for _, nb := range c.neighbors {
       if nb.City.Number == cn {
           return true
       }
    }
    return false
}


func NewCity(number int) *City {
	return &City{
		Number:    number,
		neighbors: make([]Neighbor, 0),
	}
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	cities := make([]*City, n)
	for i := 0; i < n; i++ {
		cities[i] = NewCity(i)
	}

	for i := 0; i < n; i++ {
		fromi, toi, weighti := edges[i][0], edges[i][1], edges[i][2]
        ifromCity := cities[fromi]
        itoCity := cities[toi]
		if weighti <= distanceThreshold {
			ifromCity.AddNeighbor(itoCity, weighti)
			itoCity.AddNeighbor(ifromCity, weighti)
		}
	}

    for i := 0; i < n; i++ {
        city := cities[i]
        for _, nb := range city.Neighbors() {
            for _, pnb := range nb.City.Neighbors() {
                if (nb.Weight + pnb.Weight) <= distanceThreshold {
                    if !city.HasNeighbor(pnb.City.Number) {
                        city.AddNeighbor(pnb.City, nb.Weight + pnb.Weight)
                    }

                    if !pnb.City.HasNeighbor(city.Number) {
                        pnb.City.AddNeighbor(city, nb.Weight + pnb.Weight)
                    }
                }
            }
        }
    }

    fmt.Println(cities)

    answer := cities[0]
    for i := 1; i < n; i++ {
        if answer.NumNeighbors() > cities[i].NumNeighbors() {
            answer = cities[i]
        } else if answer.NumNeighbors() == cities[i].NumNeighbors() {
            answer = cities[i]
        }
    }

	return answer.Number
}
