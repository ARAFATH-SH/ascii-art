package ascii

type CharacterPattern struct {
	Character byte
	Pattern   [3][3]float64
}

var CharacterPatterns = []CharacterPattern{
	{
		Character: ' ',
		Pattern: [3][3]float64{
			{0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
		},
	},
	{
		Character: '.',
		Pattern: [3][3]float64{
			{0.0, 0.0, 0.0},
			{0.0, 1.0, 0.0},
			{0.0, 0.0, 0.0},
		},
	},
	{
		Character: ':',
		Pattern: [3][3]float64{
			{0.0, 1.0, 0.0},
			{0.0, 0.0, 0.0},
			{0.0, 1.0, 0.0},
		},
	},
	{
		Character: '-',
		Pattern: [3][3]float64{
			{0.0, 0.0, 0.0},
			{1.0, 1.0, 1.0},
			{0.0, 0.0, 0.0},
		},
	},
	{
		Character: '|',
		Pattern: [3][3]float64{
			{0.0, 1.0, 0.0},
			{0.0, 1.0, 0.0},
			{0.0, 1.0, 0.0},
		},
	},
	{
		Character: '/',
		Pattern: [3][3]float64{
			{0.0, 0.0, 1.0},
			{0.0, 1.0, 0.0},
			{1.0, 0.0, 0.0},
		},
	},
	{
		Character: '\\',
		Pattern: [3][3]float64{
			{1.0, 0.0, 0.0},
			{0.0, 1.0, 0.0},
			{0.0, 0.0, 1.0},
		},
	},
	{
		Character: '+',
		Pattern: [3][3]float64{
			{0.0, 1.0, 0.0},
			{1.0, 1.0, 1.0},
			{0.0, 1.0, 0.0},
		},
	},
	{
		Character: '#',
		Pattern: [3][3]float64{
			{1.0, 1.0, 1.0},
			{1.0, 0.5, 1.0},
			{1.0, 1.0, 1.0},
		},
	},
	{
		Character: '@',
		Pattern: [3][3]float64{
			{1.0, 1.0, 1.0},
			{1.0, 1.0, 1.0},
			{1.0, 1.0, 1.0},
		},
	},
}
