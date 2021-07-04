package algorithm

func Calculate(calc string, operator1 float64, operator2 float64) float64 {
	switch calc {
	case "mul":
		return operator1 * operator2
	case "sub":
		return operator1 - operator2
	case "add":
		return operator1 + operator2
	case "div":
		return operator1 / operator2
	default:
		return 0
	}
}