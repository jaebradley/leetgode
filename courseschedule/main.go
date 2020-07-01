package courseschedule

/*
https://leetcode.com/problems/course-schedule/

There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?



Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0, and to take course 0 you should
             also have finished course 1. So it is impossible.


Constraints:

The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
You may assume that there are no duplicate edges in the input prerequisites.
1 <= numCourses <= 10^5

Approach:

* Build an adjacency matrix of course ids to an array of dependent courses
* Build two hashes of course ids to booleans - these booleans indicate whether or not the node has been processed
	* First hash indicates if the course has been processed as a requirement of the current course
	* Second hash indicates if the course and it's requirements have been processed before
	* If a course has been processed before, we don't need to process it's requirements again
* For each course, if it hasn't been processed before, check to see if there's a cycle in it's dependency tree
	* This means checking if a previous requirement that needs to be processed (i.e. a node to do in the future) is required again
	* If a requirement has already been checked (i.e. it has already been fully processed), can early exit
	* After processing the current node and it's dependencies (recursively), unset that the current node needs to be processed (i.e. a node to do in the future)
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacencyMatrix := make(map[int][]int, numCourses+1)
	nodesTodo := make(map[int]bool, numCourses+1)
	nodesDone := make(map[int]bool, numCourses+1)

	for _, prerequisite := range prerequisites {
		adjacencyMatrix[prerequisite[0]] = append(adjacencyMatrix[prerequisite[0]], prerequisite[1])
	}

	for i := 0; i < numCourses; i++ {
		if nodesDone[i] == false {
			if isAcyclic(adjacencyMatrix, nodesTodo, nodesDone, i) == false {
				return false
			}
		}
	}

	return true
}

func isAcyclic(adjacencyMatrix map[int][]int, nodesTodo map[int]bool, nodesDone map[int]bool, currentNode int) bool {
	if nodesTodo[currentNode] == true {
		return false
	}

	if nodesDone[currentNode] == true {
		return true
	}

	nodesTodo[currentNode] = true
	nodesDone[currentNode] = true

	prerequisites := adjacencyMatrix[currentNode]

	for _, prerequisite := range prerequisites {
		if isAcyclic(adjacencyMatrix, nodesTodo, nodesDone, prerequisite) == false {
			return false
		}
	}

	nodesTodo[currentNode] = false

	return true
}
