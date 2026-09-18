package ascii

type Glyph struct {
	Character byte
	Pattern   [3][3]float64
}

func NewGlyph(character byte, pattern [3][3]float64) Glyph {
	return Glyph{
		Character: character,
		Pattern:   pattern,
	}
}
