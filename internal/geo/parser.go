package geo

import "strings"

func ParseRegion(s string) Region {
	s = strings.TrimSpace(strings.ToUpper(s))
	parts := strings.Split(s, "-")

	switch len(parts) {
	case 3:
		return Region{City: parts[0], State: parts[1], Country: parts[2]}
	case 2:
		return Region{State: parts[0], Country: parts[1]}
	case 1:
		return Region{Country: parts[0]}
	default:
		return Region{}
	}
}

func Matches(target, rule Region) bool {
	if rule.Country != "" && rule.Country != target.Country {
		return false
	}

	if rule.State != "" && rule.State != target.State {
		return false
	}
	
	if rule.City != "" && rule.City != target.City {
		return false
	}
	return true
}
