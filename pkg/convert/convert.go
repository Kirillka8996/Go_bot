package convert

import "math"

func KelToCel(kelvin float64) float64 {
	return math.Round(kelvin - 273.15)
}
