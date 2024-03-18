package expronaut

import (
	"context"
	"math"
	"reflect"
	"testing"
	"time"
)

func TestResult50(t *testing.T) {
	input := `(5 + 5) * 5 `
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 50
	if !equalNumber(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestResult30(t *testing.T) {
	input := `5 + 5 * 5 `
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 30
	if !equalNumber(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestEqual(t *testing.T) {
	input := `"abc" == "abc" && 5 == 5.00`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	expected := `and ( eq "abc" "abc" ) ( eq 5 5.000000 )`
	if tree.GoTemplate() != expected {
		t.Errorf("expected %s, got %s", expected, tree.GoTemplate())
	}
}

func TestModulo(t *testing.T) {
	input := `5 % 2 == 1`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := true
	if !equalBool(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestVariablesNested(t *testing.T) {
	input := `foo == ( bar.baz + bar.qux.quux )`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	vars := map[string]any{
		"foo": 10,
		"bar": map[string]any{
			"baz": 5,
			"qux": map[string]any{
				"quux": 5,
			},
		},
	}

	ctx := context.TODO()
	ctx = SetVariables(ctx, vars)

	result, err := tree.Evaluate(ctx)
	if err != nil {
		t.Error(err)
	}

	expected := true
	if !equalBool(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	expectedGoTemplate := `eq .foo (add .bar.baz .bar.qux.quux)`
	if tree.GoTemplate() != expectedGoTemplate {
		t.Errorf("expected %s, got %s", expectedGoTemplate, tree.GoTemplate())
	}
}

func TestVariableBool(t *testing.T) {
	input := `false == foo || ( bar >= baz )`

	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	expected := `or ( eq false .foo ) ( ge .bar .baz )`
	if tree.GoTemplate() != expected {
		t.Errorf("expected %s, got %s", expected, tree.GoTemplate())
	}
}

func TestSQRT(t *testing.T) {
	input := `sqrt( 5.25 )`

	out, err := Evaluate(context.TODO(), input)
	if err != nil {
		t.Error(err)
	}

	expected := 2.29128784747792
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestSQRTAdvanced(t *testing.T) {
	input := `sqrt( 5 * 5 ) == 5 && sqrt( 6 * 6  ) == 6`

	out, err := Evaluate(context.TODO(), input)
	if err != nil {
		t.Error(err)
	}

	expected := true
	if !equalBool(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestExponentiation(t *testing.T) {
	input := "2 ** 3"

	out, err := Evaluate(context.TODO(), input)
	if err != nil {
		t.Error(err)
	}

	expected := 8.000000
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestNegative(t *testing.T) {
	input := "-5 * 4"

	out, err := Evaluate(context.TODO(), input)
	if err != nil {
		t.Error(err)
	}

	expected := -20
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}
func TestNegativeParenthesis(t *testing.T) {
	input := "-(5 * 4)"

	out, err := Evaluate(context.TODO(), input)
	if err != nil {
		t.Error(err)
	}

	expected := -20
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestLeftShift(t *testing.T) {
	input := "1 << 2"

	out, err := Evaluate(context.TODO(), input)
	if err != nil {
		t.Error(err)
	}

	expected := 4
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestRightShift(t *testing.T) {
	input := "4 >> 2"

	out, err := Evaluate(context.TODO(), input)
	if err != nil {
		t.Error(err)
	}

	expected := 1
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestOrderOfOperations(t *testing.T) {
	input := "-2 + 3 * 4 - 5 // 2 ^ 2 << 1 >> 2 % 3"

	out, err := Evaluate(context.TODO(), input)
	if err != nil {
		t.Error(err)
	}

	expected := 4
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestTimeDiff(t *testing.T) {
	input := " date1 == date2 "

	now := time.Now().UTC()
	past := now.Add(-time.Hour * 24)

	vars := map[string]any{
		"date1": now,
		"date2": past,
	}

	ctx := context.TODO()
	ctx = SetVariables(ctx, vars)

	out, err := Evaluate(ctx, input)
	if err != nil {
		t.Error(err)
	}

	expected := false
	if !equalBool(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestMaxNested(t *testing.T) {
	input := `max( max(10, 20), max(5.25, 25) )`

	out, err := Evaluate(context.TODO(), input)
	if err != nil {
		t.Error(err)
	}

	expected := 25
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestWrongSyntax(t *testing.T) {
	input := `45 + ((1250 x 100) / 100)`

	out, err := Evaluate(context.TODO(), input)
	if err == nil {
		t.Errorf("expected error, got %v", out)
	}
}

func TestScientific(t *testing.T) {
	input := `a + (b * 1e+06)`

	ctx := context.TODO()
	ctx = SetVariables(ctx, map[string]any{
		"a": 1,
		"b": 1,
	})

	out, err := Evaluate(ctx, input)
	if err != nil {
		t.Error(err)
	}

	if !equalNumber(out, 1.000001e+06) {
		t.Errorf("expected %v, got %v", 1.000001e+06, out)
	}
}

func TestNeg(t *testing.T) {
	input := `2 + -1`

	ctx := context.TODO()

	out, err := Evaluate(ctx, input)
	if err != nil {
		t.Error(err)
	}

	if !equalNumber(out, 1) {
		t.Errorf("expected %v, got %v", 1, out)
	}
}

func TestAiGPT(t *testing.T) {
	input := `ai("gpt", "is the following 42?", ( 21 + 21 ) )`

	_, err := Evaluate(context.TODO(), input)
	if err != nil {
		errMessage := `You didn't provide an API key. You need to provide your API key in an Authorization header using Bearer auth (i.e. Authorization: Bearer YOUR_KEY), or as the password field (with blank username) if you're accessing the API from your browser and are prompted for a username and password. You can obtain an API key from https://platform.openai.com/account/api-keys.`
		if errMessage != err.Error() {
			t.Error(err)
			return
		}
	}
}

func equalNumber(a, b any) bool {
	toFloat64 := func(v any) (float64, bool) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return float64(rv.Int()), true
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return float64(rv.Uint()), true
		case reflect.Float32, reflect.Float64:
			return rv.Float(), true
		default:
			return 0, false
		}
	}

	aFloat, aOk := toFloat64(a)
	bFloat, bOk := toFloat64(b)
	if !aOk || !bOk {
		return false
	}

	// For float comparisons consider using a tolerance to determine equality
	return math.Abs(aFloat-bFloat) < 1e-9
}

func equalBool(a, b any) bool {
	toBool := func(v any) (bool, bool) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Bool:
			return rv.Bool(), true
		default:
			return false, false
		}
	}

	aBool, aOk := toBool(a)
	bBool, bOk := toBool(b)

	return aOk && bOk && aBool == bBool
}
