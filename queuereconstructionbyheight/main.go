package queuereconstructionbyheight

import "sort"

/*
https://leetcode.com/problems/queue-reconstruction-by-height/


Suppose you have a random list of people standing in a queue. Each person is described by a pair of integers (h, k), where h is the height of the person and k is the number of people in front of this person who have a height greater than or equal to h. Write an algorithm to reconstruct the queue.

Note:
The number of people is less than 1,100.


Example

Input:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

Output:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

Approach:

* Sort people from tallest to shortest
* If two people have the same height, sort by how many people in front of them are >= their height (i.e. the k value)
* Iterating over this sorted array, insert the person at the index specified by their k value
* This approach works because people only care about people just as tall or taller
* So correctly placing the tall people relative to each other is important
* They don't care about shorter people
* So can place the shorter people at the index at the time of insertion, even if there was a previous taller person at that index
* As long as they get "pushed" back to the next index, their relative rule (i.e. their k value) should be valid
*/

func reconstructQueue(people [][]int) [][]int {
	sort.SliceStable(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[j][1] > people[i][1]
		}

		return people[j][0] < people[i][0]
	})

	queue := make([][]int, len(people))

	for i := 0; i < len(people); i++ {
		person := people[i]
		index := person[1]
		copy(queue[index+1:], queue[index:])
		queue[index] = person
	}

	return queue
}
