package blackjack

func ParseCard(card string) int {
	value := 0
	switch card {
	case "ace":
		value = 11
	case "two":
		return 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		fallthrough
	case "jack":
		fallthrough
	case "queen":
		fallthrough
	case "king":
		value = 10
	case "other":
		value = 0
	}
	return value
}

func dealerValidator(dealerCard string) bool {
	switch dealerCard {
	case "ace":
		fallthrough
	case "figure":
		fallthrough
	case "ten":
		fallthrough
	case "jack":
		fallthrough
	case "queen":
		fallthrough
	case "king":
		return false
	default:
		return true
	}
}
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	score := card1Value + card2Value

	result := ""
	switch {
	case score == 22:
		result = "P"
	case score == 21:
		if dealerValidator(dealerCard) {
			result = "W"
		} else {
			result = "S"
		}
	case score >= 17 && score <= 20:
		result = "S"
	case score >= 12 && score <= 16:
		if ParseCard(dealerCard) >= 7 {
			result = "H"
		} else {
			result = "S"
		}
	case score <= 11:
		result = "H"

	}
	return result
}