package utilities

import (
	"regexp"
	"strconv"
)

func SplitHost (host string, defaultPort int) (string, int) {
	// Splits a URL Host into a pair of (Hostname, Port)
	re := regexp.MustCompile("([^:]+)(?::([1-9][0-9]+))?")
	println("Matching host", host)
	submatches := re.FindStringSubmatch(host)
	println("Submatches", submatches)
	if len(submatches) == 0 {
		println("How curious!")
		return "", defaultPort
	} else if len(submatches) == 3 {
		matched_port, _ := strconv.Atoi(submatches[2])
		return submatches[1], matched_port
	} else {
		return submatches[1], defaultPort
	}
}
