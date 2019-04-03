# MEDIUM - Longest Substring Without Repeating Characters [![Build Status](https://api.travis-ci.org/arian-amador/GoLeetCode.svg)](https://travis-ci.org/arian-amador/GoLeetCode)

### Problem

Given a string, find the length of the _longest substring_ without repeating characters.

##### Example
```Go
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

### Solutions

There are a few ideas to solving this exercise. A naive solution would be to start with nested loops and build all substring combinations. This would be very slow not only building the lookup hash but then finding the longest string afterwards.

##### Lookup Visited Characters

Ideally we only have to loop over the given string _once_ keeping track of what characters we've visited.
Along with tracking the position in the string we'll update the lookup hash with the most recently visited position.

For example given the string `abaab` we would update [a:0] at first, but the entry for `a` would be updated 2 more times to [a:2] and then finally [a:3].

As we loop over each of the characters we are also keeping track and updated a head pointer that indicates where our substring starts and a tail pointer indicating where our substring ends. If a character is found to have been visited we update the characters position in the lookup hash and set the head to this new position.

The above makes sure we are only tracking substrings without repeating characters. The final piece is to keep a running `max` of the longest substring we've found.

```Go
func lengthOfLongestSubstring(s string) int {
    var head, max int
    n := len(s)
    visited := make(map[string]int)

    for tail := 0; tail < n; tail++ {
        // The current character
        curr := s[tail : tail+1]

        // Have we seen this char and
        // is it's position higher than the current head?
        if visited[curr] > 0 && visited[curr] > head {
            head = visited[curr] // Update head pointer
        }

        // Is the current substring length longer than
        // previous tracked substring length/
        if max < tail-head+1 {
            max = tail - head + 1 // Update max substring found
        }

        // Update the lookup table for the current char with the current position
        visited[curr] = tail + 1
    }

    return max
}
```

A further optimization to the above solution is to replace the lookup hash of strings as keys to a hash of integers for keys. Each with an initial value of `-1` to indicate the character has not been visited.

---

#### Hope you find this useful!
