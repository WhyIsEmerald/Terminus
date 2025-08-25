package calculator

import (
	"errors"
	"fmt"
	"math"
)

// Operator struct to hold precedence, associativity, operation, and operand count
type Operator struct {
	Precedence    int
	Associativity string
	Operation     interface{}
	OperandCount  int
}

var operators = map[string]Operator{ // Use Operator struct as the value type
	"+": {Precedence: 1, Associativity: "left", Operation: func(x, y float64) (float64, error) { return x + y, nil }, OperandCount: 2},
	"-": {Precedence: 1, Associativity: "left", Operation: func(x, y float64) (float64, error) { return x - y, nil }, OperandCount: 2},
	"*": {Precedence: 2, Associativity: "left", Operation: func(x, y float64) (float64, error) { return x * y, nil }, OperandCount: 2},
	"/": {Precedence: 2, Associativity: "left", Operation: func(x, y float64) (float64, error) {
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	}, OperandCount: 2},
	"%": {Precedence: 2, Associativity: "left", Operation: func(x, y int64) (int64, error) {
		if y == 0 {
			return 0, errors.New("modulo by zero")
		}
		return x % y, nil
	}, OperandCount: 2},
	"sin":   {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Sin(x), nil }, OperandCount: 1},
	"cos":   {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Cos(x), nil }, OperandCount: 1},
	"tan":   {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Tan(x), nil }, OperandCount: 1},
	"csc":   {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return 1 / math.Sin(x), nil }, OperandCount: 1},
	"sec":   {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return 1 / math.Cos(x), nil }, OperandCount: 1},
	"cot":   {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return 1 / math.Tan(x), nil }, OperandCount: 1},
	"asin":  {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Asin(x), nil }, OperandCount: 1},
	"acos":  {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Acos(x), nil }, OperandCount: 1},
	"atan":  {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Atan(x), nil }, OperandCount: 1},
	"atan2": {Precedence: 4, Associativity: "left", Operation: func(y, x float64) (float64, error) { return math.Atan2(y, x), nil }, OperandCount: 2},
	"log":   {Precedence: 4, Associativity: "left", Operation: func(base, x float64) (float64, error) { return math.Log(x) / math.Log(base), nil }, OperandCount: 2},
	"pow":   {Precedence: 3, Associativity: "right", Operation: func(x, y float64) (float64, error) { return math.Pow(x, y), nil }, OperandCount: 2},
	"sqrt":  {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Sqrt(x), nil }, OperandCount: 1},
	"abs":   {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Abs(x), nil }, OperandCount: 1},
	"exp":   {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Exp(x), nil }, OperandCount: 1},
	"ln":    {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Log(x), nil }, OperandCount: 1},
	"log10": {Precedence: 4, Associativity: "left", Operation: func(x float64) (float64, error) { return math.Log10(x), nil }, OperandCount: 1},
}

type Token struct {
	TokenType     string
	Value         string
	Precedence    int
	Associativity string
	Operation     interface{}
	OperandCount  int
}

func getSpecialToken(tokenValue string) (Operator, error) {
	op, ok := operators[tokenValue]
	if !ok {
		return Operator{}, fmt.Errorf("unknown operator: %s", tokenValue)
	}
	return op, nil
}

func generateToken(token string, tokenType string) Token {
	var op Operator
	if tokenType == "operator" {
		op, _ = getSpecialToken(token)
	}
	if tokenType == "leftParenthesis" {
		return Token{
			Value:         token,
			TokenType:     tokenType,
			Precedence:    0,
			Associativity: "left",
			Operation:     nil,
			OperandCount:  0,
		}
	}
	if tokenType == "rightParenthesis" {
		return Token{
			Value:         token,
			TokenType:     tokenType,
			Precedence:    0,
			Associativity: "right",
			Operation:     nil,
			OperandCount:  0,
		}
	}
	return Token{
		Value:         token,
		TokenType:     tokenType,
		Precedence:    op.Precedence,
		Associativity: op.Associativity,
		Operation:     op.Operation,
		OperandCount:  op.OperandCount,
	}
}
