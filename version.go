package buildver

import (
	"strconv"
	"strings"
)

// Version is a build version
type Version struct {
	vers []int
}

// New creates a new Version.
// It expects a period-seperated string of integers.
func New(ver string) (Version, error) {
	split := strings.Split(ver, ".")
	vers := make([]int, len(split))
	for i, part := range split {
		n, err := strconv.Atoi(part)
		if err != nil {
			return Version{}, err
		}
		vers[i] = n
	}

	// remove meaningless zeroes!
	clip := 0
	for i := len(vers) - 1; i >= 0; i-- {
		if vers[i] == 0 {
			clip++
		} else {
			break
		}
	}
	vers = vers[:len(vers)-clip]

	return Version{vers}, nil
}

// Less compares two Versions, returning true if this version is less than the other.
func (v Version) Less(other Version) bool {
	size := min(len(v.vers), len(other.vers))
	for i := 0; i < size; i++ {
		if v.vers[i] > other.vers[i] {
			return false
		} else if v.vers[i] < other.vers[i] {
			return true
		}
	}

	if len(v.vers) > len(other.vers) {
		return false
	}

	return true
}

// Equals comapres two Versions, returning true if both are equal.
func (v Version) Equals(other Version) bool {
	if len(v.vers) != len(other.vers) {
		return false
	}
	for i, ver := range v.vers {
		if ver != other.vers[i] {
			return false
		}
	}
	return true
}

// String returns a string representation of this Version.
func (v Version) String() string {
	if v.vers == nil {
		return "0"
	}
	strs := make([]string, len(v.vers))
	for i, n := range v.vers {
		strs[i] = strconv.Itoa(n)
	}
	return strings.Join(strs, ".")
}
