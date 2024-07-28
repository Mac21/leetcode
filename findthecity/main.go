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
	Number int
	Weight int
}

func (n Neighbor) String() string {
	return fmt.Sprintf("City %d:%d", n.Number, n.Weight)
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
		if nb != nil && nb.Weight <= threshold {
			res++
		}
	}

	return res
}

func getRealIndex(n int) int {
    if n == 0 {
        return 0
    }
    return n - 1
}

func (c *City) UpdateNeighbor(cn, weight int) {
	if c.neighbors[getRealIndex(cn)] == nil {
        c.neighbors[getRealIndex(cn)] = &Neighbor{
            Number: cn,
            Weight: weight,
        }
		return
	}

	cnb := c.neighbors[getRealIndex(cn)]
	if weight < cnb.Weight {
		cnb.Weight = weight
        c.neighbors[getRealIndex(cn)] = cnb
	}
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
		if nb != nil && nb.Weight <= threshold {
			res = append(res, nb)
		}
	}

	return res
}

func NewCity(number, n int) *City {
	return &City{
		Number:    number,
		neighbors: make([]*Neighbor, n-1),
	}
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	cities := make([]*City, n)
	for i := 0; i < n; i++ {
		cities[i] = NewCity(i, n)
	}

	for i := 0; i < len(edges); i++ {
		from, to, weight := edges[i][0], edges[i][1], edges[i][2]
		fromCity := cities[from]
		toCity := cities[to]
		fromCity.UpdateNeighbor(to, weight)
		toCity.UpdateNeighbor(from, weight)
	}

	queue := gods.NewMinPriorityQueue[*Neighbor, int]()
	for i := 0; i < n; i++ {
		city := cities[i]
		for _, nb := range city.Neighbors() {
			if nb != nil {
				queue.Push(nb, nb.Weight)
			}
		}

		for queue.Len() > 0 {
			nb, prio, ok := queue.Pop()
			if !ok {
				break
			}

			for _, fnb := range cities[nb.Number].Neighbors() {
                if fnb == nil {
                    continue
                }

				newWeight := prio + fnb.Weight
                if newWeight <= distanceThreshold {
					city.UpdateNeighbor(fnb.Number, newWeight)
                    cities[fnb.Number].UpdateNeighbor(city.Number, newWeight)
                }
			}
		}
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
