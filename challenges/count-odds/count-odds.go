package countodds

func CountOdds(low int, high int) int {
	firstBorder := low&1 != 0
	secondBorder := high&1 != 0

	between := (high - low) / 2
	if firstBorder || secondBorder {
		between += 1
	}

	return between

}
