package main

import (
	"fmt"

	"github.com/mac21/gods"
)

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
	neighbors []*Neighbor
}

func (c *City) String() string {
	if c == nil {
		return "nil"
	}

	return fmt.Sprintf("City %d -> (%d)%v\n", c.Number, c.NumNeighbors(), c.Neighbors())
}

func (c *City) NumNeighbors() int {
	if c == nil {
		return 0
	}

	return len(c.neighbors)
}

func (c *City) NumNeighborsInThreshold(threshold int) int {
	if c == nil {
		return 0
	}

	res := 0
	for _, nb := range c.neighbors {
		if nb.Weight <= threshold {
			res++
		}
	}

	return res
}

func (c *City) HasNeighbor(cn, weight int) bool {
	if c == nil {
		return true
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

func (c *City) AddNewNeighbor(n *City, weight int) {
	if c.HasNeighbor(n.Number, weight) {
		return
	}

	c.neighbors = append(c.neighbors, &Neighbor{
		City:   n,
		Weight: weight,
	})
}

func (c *City) AddNeighbor(nb *Neighbor) {
	if c.HasNeighbor(nb.City.Number, nb.Weight) {
		return
	}

	c.neighbors = append(c.neighbors, nb)
}

func (c *City) Neighbors() []*Neighbor {
	if c == nil {
		return nil
	}

	return c.neighbors
}

func (c *City) NeighborsInThreshold(threshold int) []*Neighbor {
	if c == nil {
		return nil
	}

	res := make([]*Neighbor, 0)
	for _, nb := range c.Neighbors() {
		if nb.Weight <= threshold {
			res = append(res, nb)
		}
	}

	return res
}

func NewCity(number int) *City {
	return &City{
		Number:    number,
		neighbors: make([]*Neighbor, 0),
	}
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	cities := make([]*City, n)
	for i := 0; i < n; i++ {
		cities[i] = NewCity(i)
	}

	for i := 0; i < len(edges); i++ {
		from, to, weight := edges[i][0], edges[i][1], edges[i][2]
		fromCity := cities[from]
		toCity := cities[to]
		if weight <= distanceThreshold {
			fromCity.AddNewNeighbor(toCity, weight)
			toCity.AddNewNeighbor(fromCity, weight)
		}
	}

	seen := make(map[*Neighbor]bool)
	queue := gods.NewPriorityQueue[*Neighbor]()
	for i := 0; i < n; i++ {
		city := cities[i]
		for _, nb := range city.Neighbors() {
			queue.Push(nb, nb.Weight)
		}

		for queue.Len() > 0 {
			nb, err := queue.Pop()
			if err != nil {
				break
			}

            if seen[nb] || nb.Weight >= distanceThreshold {
                continue
            }

			for _, fnb := range nb.City.Neighbors() {
				if seen[fnb] || fnb.Weight >= distanceThreshold {
					continue
				}

				seen[fnb] = true

				if (nb.Weight + fnb.Weight) <= distanceThreshold {
					cnb := &Neighbor{
						City:   fnb.City,
						Weight: nb.Weight + fnb.Weight,
					}
					city.AddNeighbor(cnb)
                    fnb.City.AddNewNeighbor(city, nb.Weight+fnb.Weight)
                    queue.Push(cnb, cnb.Weight)
				}
			}
		}
        clear(seen)
	}

	answer := cities[0]
	for i := 1; i < n; i++ {
		cityNumNbrs := cities[i].NumNeighborsInThreshold(distanceThreshold)
		answerNumNbrs := answer.NumNeighborsInThreshold(distanceThreshold)
		if cityNumNbrs <= answerNumNbrs {
			answer = cities[i]
		}
	}

	return answer.Number
}
