package main

func main() {}

func lengthOfLongestSubstring(s string) int {
	var head, max int
	n := len(s)
	visited := make(map[string]int)
	for tail := 0; tail < n; tail++ {
		curr := s[tail : tail+1]
		if visited[curr] > 0 && visited[curr] > head {
			head = visited[curr]
		}
		if max < tail-head+1 {
			max = tail - head + 1
		}
		visited[curr] = tail + 1
	}
	return max
}
