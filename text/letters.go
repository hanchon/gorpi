package text

const (
	LetterWidth  = 5
	LetterHeight = 5
)

var LetterPixelArrays = map[rune][][]bool{
	'a': {
		{false, true, true, true, false},
		{true, false, false, false, true},
		{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
	},
	'b': {
		{true, true, true, true, false},
		{true, false, false, false, true},
		{true, true, true, true, false},
		{true, false, false, false, true},
		{true, true, true, true, false},
	},
	'c': {
		{false, true, true, true, false},
		{true, false, false, false, true},
		{true, false, false, false, false},
		{true, false, false, false, true},
		{false, true, true, true, false},
	},
	'd': {
		{true, true, true, true, false},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, false},
	},
	'e': {
		{true, true, true, true, true},
		{true, false, false, false, false},
		{true, true, true, true, false},
		{true, false, false, false, false},
		{true, true, true, true, true},
	},
	'f': {
		{true, true, true, true, true},
		{true, false, false, false, false},
		{true, true, true, true, false},
		{true, false, false, false, false},
		{true, false, false, false, false},
	},
	'g': {
		{false, true, true, true, true},
		{true, false, false, false, false},
		{true, false, true, true, true},
		{true, false, false, false, true},
		{false, true, true, true, true},
	},
	'h': {
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
	},
	'i': {
		{true, true, true, true, true},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{true, true, true, true, true},
	},
	'j': {
		{true, true, true, true, true},
		{false, false, false, true, false},
		{false, false, false, true, false},
		{true, false, false, true, false},
		{false, true, true, true, false},
	},
	'k': {
		{true, false, false, false, true},
		{true, false, false, true, false},
		{true, true, true, false, false},
		{true, false, false, true, false},
		{true, false, false, false, true},
	},
	'l': {
		{true, false, false, false, false},
		{true, false, false, false, false},
		{true, false, false, false, false},
		{true, false, false, false, false},
		{true, true, true, true, true},
	},
	'm': {
		{true, false, false, false, true},
		{true, true, false, true, true},
		{true, false, true, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
	},
	'n': {
		{true, false, false, false, true},
		{true, true, false, false, true},
		{true, false, true, false, true},
		{true, false, false, true, true},
		{true, false, false, false, true},
	},
	'o': {
		{false, true, true, true, false},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{false, true, true, true, false},
	},
	'p': {
		{true, true, true, true, false},
		{true, false, false, false, true},
		{true, true, true, true, false},
		{true, false, false, false, false},
		{true, false, false, false, false},
	},
	'q': {
		{false, true, true, true, false},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{false, true, true, false, true},
		{false, false, false, true, true},
	},
	'r': {
		{true, true, true, true, false},
		{true, false, false, false, true},
		{true, true, true, true, false},
		{true, false, false, true, false},
		{true, false, false, false, true},
	},
	's': {
		{false, true, true, true, true},
		{true, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, true},
		{true, true, true, true, false},
	},
	't': {
		{true, true, true, true, true},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
	},
	'u': {
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{false, true, true, true, false},
	},
	'v': {
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{false, true, false, true, false},
		{false, false, true, false, false},
	},
	'w': {
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, true, false, true},
		{true, true, false, true, true},
		{true, false, false, false, true},
	},
	'x': {
		{true, false, false, false, true},
		{false, true, false, true, false},
		{false, false, true, false, false},
		{false, true, false, true, false},
		{true, false, false, false, true},
	},
	'y': {
		{true, false, false, false, true},
		{false, true, false, true, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
	},
	'z': {
		{true, true, true, true, true},
		{false, false, false, true, false},
		{false, false, true, false, false},
		{false, true, false, false, false},
		{true, true, true, true, true},
	},
}
