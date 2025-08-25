package calculator

import "strconv"

func EvaluateRPN(tokens []Token) (string, error) {
	stack := NewStack()

	for _, token := range tokens {
		if token.TokenType == "number" {
			stack.Push(token)
		}
		if token.TokenType == "operator" {
			switch token.OperandCount {
			case 1:
				f, err := HandleUnary(stack, token)
				if err != nil {
					return f, err
				}
			case 2:
				if stack.Size() >= 2 {
					if token.Value == "%" {
						operand2, err1 := strconv.ParseInt(stack.Pop().Value, 10, 64)
						if err1 != nil {
							return "", err1
						}
						operand1, err2 := strconv.ParseInt(stack.Pop().Value, 10, 64)
						if err2 != nil {
							return "", err2
						}
						result, err := token.Operation.(func(int64, int64) (int64, error))(operand1, operand2)
						if err != nil {
							return "", err
						}
						stack.Push(Token{TokenType: "number", Value: strconv.FormatInt(result, 10)})
						continue
					}
					operand2, err1 := strconv.ParseFloat(stack.Pop().Value, 64)
					if err1 != nil {
						return "", err1
					}
					operand1, err2 := strconv.ParseFloat(stack.Pop().Value, 64)
					if err2 != nil {
						return "", err2
					}
					result, err := token.Operation.(func(float64, float64) (float64, error))(operand1, operand2)
					if err != nil {
						return "", err
					}
					stack.Push(Token{TokenType: "number", Value: strconv.FormatFloat(result, 'f', -1, 64)})
				}
			}
		}
	}
	return stack.Pop().Value, nil
}

func HandleUnary(stack *Stack, token Token) (string, error) {
	if stack.Size() >= 1 {
		operand, err := strconv.ParseFloat(stack.Pop().Value, 64)
		if err != nil {
			return "", err
		}
		result, err := token.Operation.(func(float64) (float64, error))(operand)
		if err != nil {
			return "", err
		}
		stack.Push(Token{TokenType: "number", Value: strconv.FormatFloat(result, 'f', -1, 64)})
	}
	return "", nil
}
