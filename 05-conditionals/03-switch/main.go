package main

func Direction(d string) string {
	switch d {
	case "N": return "north"
	case "S": return "south"
	case "W": return "west"
	case "E": return "east"
	default: return "unknown"

	}
}
