package expronaut

import "testing"

func TestNewLexer(t *testing.T) {
	input := `(3 + 4) * 2 `
	lexer := NewLexer(input)

	exp := []TokenType{
		TokenTypeParenLeft,
		TokenTypeInt,
		TokenTypePlus,
		TokenTypeInt,
		TokenTypeParenRight,
		TokenTypeMultiply,
		TokenTypeInt,
	}

	i := 0
	for tok := lexer.NextToken(); tok.Type != TokenTypeEOF; tok = lexer.NextToken() {
		if tok.Type != exp[i] {
			t.Fatalf("expected %v, got %v", exp[i], tok.Type)
		}
		i++
	}
}

func TestNewLexerEq(t *testing.T) {
	input := `(2+5) == 7`

	lexer := NewLexer(input)

	exp := []TokenType{
		TokenTypeParenLeft,
		TokenTypeInt,
		TokenTypePlus,
		TokenTypeInt,
		TokenTypeParenRight,
		TokenTypeEqual,
		TokenTypeInt,
	}

	i := 0
	for tok := lexer.NextToken(); tok.Type != TokenTypeEOF; tok = lexer.NextToken() {
		if tok.Type != exp[i] {
			t.Fatalf("expected %v, got %v", exp[i], tok.Type)
		}
		i++
	}
}

func TestNewLexerEqFloat(t *testing.T) {
	input := `(2.5+4.5) == 7`

	lexer := NewLexer(input)

	exp := []TokenType{
		TokenTypeParenLeft,
		TokenTypeFloat,
		TokenTypePlus,
		TokenTypeFloat,
		TokenTypeParenRight,
		TokenTypeEqual,
		TokenTypeInt,
	}

	i := 0
	for tok := lexer.NextToken(); tok.Type != TokenTypeEOF; tok = lexer.NextToken() {
		if tok.Type != exp[i] {
			t.Fatalf("expected %v, got %v", exp[i], tok.Type)
		}
		i++
	}
}

func TestNewLexerNe(t *testing.T) {
	input := `(2+5) != 4`

	lexer := NewLexer(input)

	exp := []TokenType{
		TokenTypeParenLeft,
		TokenTypeInt,
		TokenTypePlus,
		TokenTypeInt,
		TokenTypeParenRight,
		TokenTypeNotEqual,
		TokenTypeInt,
	}

	i := 0
	for tok := lexer.NextToken(); tok.Type != TokenTypeEOF; tok = lexer.NextToken() {
		if tok.Type != exp[i] {
			t.Fatalf("expected %v, got %v", exp[i], tok.Type)
		}
		i++
	}
}

func TestNewLexerString(t *testing.T) {
	input := `a == b`

	lexer := NewLexer(input)

	exp := []TokenType{
		TokenTypeVariable,
		TokenTypeEqual,
		TokenTypeVariable,
	}

	i := 0
	for tok := lexer.NextToken(); tok.Type != TokenTypeEOF; tok = lexer.NextToken() {
		if tok.Type != exp[i] {
			t.Fatalf("expected %v, got %v", exp[i], tok.Type)
		}
		i++
	}
}

func TestNewLexerStringLong(t *testing.T) {
	input := `abc.def == "abc"`

	lexer := NewLexer(input)

	exp := []TokenType{
		TokenTypeVariable,
		TokenTypeEqual,
		TokenTypeString,
	}

	i := 0
	for tok := lexer.NextToken(); tok.Type != TokenTypeEOF; tok = lexer.NextToken() {
		if tok.Type != exp[i] {
			t.Fatalf("expected %v, got %v", exp[i], tok.Type)
		}
		i++
	}
}

func TestNewLexerVariable(t *testing.T) {
	input := `foo == ( bar >= baz )`

	lexer := NewLexer(input)

	exp := []TokenType{
		TokenTypeVariable,
		TokenTypeEqual,
		TokenTypeParenLeft,
		TokenTypeVariable,
		TokenTypeGreaterThanOrEqual,
		TokenTypeVariable,
		TokenTypeParenRight,
	}

	i := 0
	for tok := lexer.NextToken(); tok.Type != TokenTypeEOF; tok = lexer.NextToken() {
		if tok.Type != exp[i] {
			t.Fatalf("expected %v, got %v", exp[i], tok.Type)
		}
		i++
	}
}

func TestNewLexerVariableBool(t *testing.T) {
	input := `false == foo || ( bar >= baz )`

	lexer := NewLexer(input)

	exp := []TokenType{
		TokenTypeBool,
		TokenTypeEqual,
		TokenTypeVariable,
		TokenTypeOr,
		TokenTypeParenLeft,
		TokenTypeVariable,
		TokenTypeGreaterThanOrEqual,
		TokenTypeVariable,
		TokenTypeParenRight,
	}

	i := 0
	for tok := lexer.NextToken(); tok.Type != TokenTypeEOF; tok = lexer.NextToken() {
		if tok.Type != exp[i] {
			t.Fatalf("expected %v, got %v", exp[i], tok.Type)
		}
		i++
	}
}
