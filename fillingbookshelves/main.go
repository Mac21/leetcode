package main

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

// 1) put book on this shelf
// 2) put book on next shelf
func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		remainingWidth := shelfWidth
		maxHeight := 0
		dp[i] = 10000
		for j := i; j > 0; j-- {
			w, h := books[j-1][0], books[j-1][1]
			if remainingWidth < w {
				break
			}
			remainingWidth -= w

			maxHeight = max(maxHeight, h)
			// current shelf height is the minimum of the current shelf height and the current book on the previous shelf
			dp[i] = min(dp[i], maxHeight+dp[j-1])
		}
	}

	return dp[n]
}
