package hw

import (
	"math"
)

// distance метод вычисляет расстояния между двумя точками
func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
