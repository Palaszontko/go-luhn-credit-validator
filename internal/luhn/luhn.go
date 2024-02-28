package luhn

func ValidateCard(number string) bool {

	var checkSum int = 0
	var second bool = false

	for i := len(number) - 1; i >= 0; i-- {
		var value int = int(number[i] - '0')

		if second {
			value *= 2
			// 2 digit number -> sum of digits
			if value > 9 {
				value = value/10 + value%10
			}
		}
		checkSum += value
		second = !second
	}
	return checkSum%10 == 0
}

func CardNetwork(number string) string {
	switch number[0] - '0' {
	case 2, 5:
		return "Mastercard"
	case 3:
		return "American Express"
	case 4:
		return "Visa"
	case 6:
		return "Discover"
	default:
		return "Unknown network"
	}
}
