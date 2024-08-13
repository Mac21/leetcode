package main

import (
	"fmt"
)

/*
1105. Filling Bookcase Shelves Medium
You are given an array books where books[i] = [thicknessi, heighti] indicates the thickness and height of the ith book. You are also given an integer shelfWidth.

We want to place these books in order onto bookcase shelves that have a total width shelfWidth.

We choose some of the books to place on this shelf such that the sum of their thickness is less than or equal to shelfWidth, then build another level of the shelf of the bookcase so that the total height of the bookcase has increased by the maximum height of the books we just put down. We repeat this process until there are no more books to place.

Note that at each step of the above process, the order of the books we place is the same order as the given sequence of books.

For example, if we have an ordered list of 5 books, we might place the first and second book onto the first shelf, the third book on the second shelf, and the fourth and fifth book on the last shelf.
Return the minimum possible height that the total bookshelf can be after placing shelves in this manner.

Constraints:

1 <= books.length <= 1000
1 <= thicknessi <= shelfWidth <= 1000
1 <= heighti <= 1000

Hint: Use dynamic programming: dp(i) will be the answer to the problem for books[i:].
*/
func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	indexHeights := make([]int, n)
	for i := 1; i < n; i++ {
		width := books[i][0]
		currentHeight := books[i][1]
		// Height of book on next level
		indexHeights[i] = currentHeight + indexHeights[i-1]
		for j := 0; j < i; j++ {
			width += books[j][0]
			if width > shelfWidth {
				break
			}

			currentHeight = max(currentHeight, books[j][1])

			indexHeights[i] = min(indexHeights[i-j], currentHeight+indexHeights[i-j])
		}
	}

	fmt.Println(indexHeights)

	return indexHeights[n-1]
}
