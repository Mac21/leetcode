package main

/*
1395. Count Number of Teams Medium
There are n soldiers standing in a line. Each soldier is assigned a unique rating value.

You have to form a team of 3 soldiers amongst them under the following rules:

Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rating[k]).
A team is valid if: (rating[i] < rating[j] < rating[k]) or (rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n).
Return the number of teams you can form given the conditions. (soldiers can be part of multiple teams).

Example 1:

Input: rating = [2,5,3,4,1]
Output: 3
Explanation: We can form three teams given the conditions. (2,3,4), (5,4,1), (5,3,1).

Example 2:

Input: rating = [2,1,3]
Output: 0
Explanation: We can't form any team given the conditions.

Example 3:

Input: rating = [1,2,3,4]
Output: 4

Constraints:

n == rating.length
3 <= n <= 1000
1 <= rating[i] <= 105
All the integers in rating are unique.
*/
// O(n-1 * n)
func numTeamsMostOptimal(rating []int) int {
	n := len(rating)

	teamCount := 0
	for i := 0; i < n - 1; i++ {
        leftmore, leftless , rightmore, rightless := 0, 0, 0, 0
		for j := i + 1; j < n; j++ {
            if rating[j] < rating[i] {
                rightless++
            } else if rating[j] > rating[i] {
                rightmore++
            }
		}

        for k := 0; k < i; k++ {
            if rating[k] < rating[i] {
                leftless++
            }
            if rating[k] > rating[i] {
                leftmore++
            }
        }
        teamCount += rightless * leftmore + leftless * rightmore
	}
	return teamCount
}

// O(n^3)
func numTeams(rating []int) int {
	n := len(rating)

	teamCount := 0
	for i := 0; i < n - 2; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if ((rating[i] < rating[j]) && (rating[j] < rating[k])) || ((rating[i] > rating[j]) && (rating[j] > rating[k])) {
					teamCount += 1
				}
			}
		}
	}
	return teamCount
}
