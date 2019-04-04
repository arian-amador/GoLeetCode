# EASY - Unique Email Addresses [![Build Status](https://api.travis-ci.org/arian-amador/GoLeetCode.svg)](https://travis-ci.org/arian-amador/GoLeetCode)

### Problem

Every email consists of a local name and a domain name, separated by the @ sign.

For example, in alice@leetcode.com, alice is the local name, and leetcode.com is the domain name.

Besides lowercase letters, these emails may contain `.`s or `+`s.

If you add periods (`.`) between some characters in the local name part of an email address, mail sent there will be forwarded to the same address without dots in the local name. For example, `"alice.z@leetcode.com"` and `"alicez@leetcode.com"` forward to the same email address. (Note that this rule does not apply for domain names.)

If you add a plus (`+`) in the local name, everything after the first plus sign will be ignored. This allows certain emails to be filtered, for example m.y+name@email.com will be forwarded to my@email.com. (Again, this rule does not apply for domain names.)

It is possible to use both of these rules at the same time.

Given a list of emails, we send one email to each address in the list. How many different addresses actually receive mails?

###### Example

```Go
Input: ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]

Output: 2
Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails
```

### Solutions

One caveat when working with this was the use of Replace vs. ReplaceAll. ReplaceAll was added in Go 1.12 and not yet available on leetcode.com. ReplaceAll is actually a helper to Replace. Here we use the original Go way of strings.Replace which uses a `-1` to indicate all instances throughout the string.

Another was using an empty struct for the value of the email listing map. Using a `struct{}` vs a `bool` or `int` allows us to save some additiona memory for every entry added to the map.

```Go
func numUniqueEmails(emails []string) int {
	t := make(map[string]struct{})

	for _, e := range emails {
		s := strings.Split(e, "@")
		local := strings.Replace(s[0], ".", "", -1)
		local = strings.Split(local, "+")[0]
		t[strings.Join([]string{local, s[1]}, "@")] = struct{}{}
	}

	return len(t)
}
```

#####

- [ReplaceAll](https://golang.org/src/strings/strings.go?s=22820:22862#L877)

---

#### Hope you find this useful!
