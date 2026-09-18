package ascii

func PatternDifference(a [3][3]float64, b [3][3]float64) float64 {
	var total float64

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			difference := a[y][x] - b[y][x]
			total += difference * difference
		}
	}
	return total / 9.0
}
