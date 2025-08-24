package internals

func ShuntingYard(tokens []Token) []Token {
	stack := NewStack()
	output := NewStack()
	for _, token := range tokens {
		switch token.TokenType {
		case "rightParenthesis":
			for stack.Peek().TokenType != "leftParenthesis" {
				output.Push(stack.Pop())
				if stack.Size() == 0 {
					panic("Mismatched parentheses")
				}
			}
			stack.Pop()
		case "leftParenthesis":
			stack.Push(token)
		case "operator":
			for stack.Size() > 0 && stack.Peek().TokenType == "operator" && token.Precedence <= stack.Peek().Precedence {
				output.Push(stack.Pop())
			}
			stack.Push(token)
		default:
			output.Push(token)
		}
	}
	for stack.Size() > 0 {
		output.Push(stack.Pop())
	}
	return output.items
}
