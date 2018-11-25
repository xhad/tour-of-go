package one

// Strings is a bunch of string methods
type Strings struct{}

// Swap returns the swaps strings
func (d Strings) Swap(x, y string) (string, string) {
	return y, x
}
