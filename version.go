// Package buildver helps compare period-separated version strings of arbitrary length.
package buildver

import (
	"strconv"
	"strings"
)

// Version is a build version of arbitrary length.
type Version struct {
	vers []int
}

// New creates a new Version.
// It expects a period-seperated string of integers.
func New(ver string) (Version, error) {
	split := strings.Split(ver, ".")
	vers := make([]int, len(split))
	for i, part := range split {
		part = strings.TrimFunc(part, notDigitOrDot)
		n, err := strconv.Atoi(part)
		if err != nil {
			return Version{}, err
		}
		vers[i] = n
	}

	vers = normalize(vers)
	return Version{vers}, nil
}

// FromInts creates a new Version.
// New("1.2.3") is equivalent to FromInts(1, 2, 3).
func FromInts(versions ...int) Version {
	versions = normalize(versions)
	return Version{versions}
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

	switch {
	case len(v.vers) > len(other.vers):
		// we're longer, and thus more than the other one
		return false
	case len(v.vers) < len(other.vers):
		// we're shorter, and thus less than the other one
		return true
	}

	// we're equal
	return false
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

// Contains returns true if this Version has the same components as other, ignoring extra components.
// For example, version 1.2.3 would contain versions 1, 1.2, and 1.2.3.
func (v Version) Contains(other Version) bool {
	if len(other.vers) > len(v.vers) || (other.vers == nil && v.vers != nil) {
		return false
	}
	for i, ver := range other.vers {
		if v.vers[i] != ver {
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

// normalize removes meaningless zeroes
func normalize(vers []int) []int {
	clip := 0
	for i := len(vers) - 1; i >= 0; i-- {
		if vers[i] == 0 {
			clip++
		} else {
			break
		}
	}
	return vers[:len(vers)-clip]
}
