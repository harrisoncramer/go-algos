package doubleletters

func DoubleLetters(input string) string {
	var output string
	for _, r := range input {
		output += (string(r) + string(r))
	}

	return output
}

func DoubleLettersTwo(original string) string {
	endString := make([]rune, 0, len(original)*2)
	for _, c := range original {
		endString = append(endString, c, c)
	}
	return string(endString)
}
