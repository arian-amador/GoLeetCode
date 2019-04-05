package uniqueEmailAddresses

import (
	"strings"
)

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
