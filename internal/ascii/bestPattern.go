package ascii

func FindBestCharacter(pattern [3][3]float64) byte {
	bestCharacter := CharacterPatterns[0].Character
	bestDifference := PatternDifference(pattern, CharacterPatterns[0].Pattern)

	for _, characterPattern := range CharacterPatterns[1:] {
		difference := PatternDifference(pattern, characterPattern.Pattern)

		if difference < bestDifference {
			bestDifference = difference
			bestCharacter = characterPattern.Character
		}
	}
	return bestCharacter
}
