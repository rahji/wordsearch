package vector

// Vector is a private type that represents the 2 axes of a cardinal direction
type Vector struct {
	X int
	Y int
}

// CardinalToVector returns an xy vector for a given one- or two-letter abbreviation
// for a cardinal direction
func CardinalToVector(cardinal string) Vector {
	switch cardinal {
	case "N":
		return Vector{X: 0, Y: -1}
	case "NE":
		return Vector{X: 0, Y: 0}
	case "E":
		return Vector{X: 1, Y: 0}
	case "SE":
		return Vector{X: 1, Y: 1}
	case "S":
		return Vector{X: 0, Y: 1}
	case "SW":
		return Vector{X: -1, Y: 1}
	case "W":
		return Vector{X: -1, Y: 0}
	case "NW":
		return Vector{X: -1, Y: -1}
	default:
		panic("unrecognized cardinal direction")
	}
}
