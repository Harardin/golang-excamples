// removeCharacters Removes certain chars from a string
func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}
	return strings.Map(filter, input)
}
