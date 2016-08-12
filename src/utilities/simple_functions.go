package utilities

import (
	"regexp"
	"strconv"
)

func SplitHost (host string, default_port int) (string, int) {
	// Splits a URL Host into a pair of (Hostname, Port)
	re := regexp.MustCompile("([^:]+)(?::([1-9][0-9]+))?")
	submatches := re.FindStringSubmatch(host)
	if len(submatches) == 3 {
		matched_port, _ := strconv.Atoi(submatches[2])
		return submatches[1], matched_port
	} else {
		return submatches[1], default_port
	}
}
