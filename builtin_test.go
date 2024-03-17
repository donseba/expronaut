package expronaut

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestBif_Ai(t *testing.T) {
	t.Skip("not implemented")
}

func TestBif_Abs(t *testing.T) {
	input := `abs(-5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 5
	if !equalNumber(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestBif_Acos(t *testing.T) {
	input := `acos(0.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 1.0471975511965979
	if !equalNumber(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestBif_Add(t *testing.T) {
	input := `add(5, 5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 10
	if !equalNumber(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestBif_Asin(t *testing.T) {
	input := `asin(0.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 0.5235987755982989
	if !equalNumber(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestBif_Atan(t *testing.T) {
	input := `atan(0.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 0.4636476090008061
	if !equalNumber(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestBif_Ceil(t *testing.T) {
	input := `ceil(5.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 6
	if !equalNumber(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestBif_Concat(t *testing.T) {
	input := `concat("a", "b")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := "ab"
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestBif_Cos(t *testing.T) {
	input := `cos(0.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	result, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 0.8775825618903728
	if !equalNumber(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestBif_Cosh(t *testing.T) {
	input := `cosh(0.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 1.1276259652063807
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Date(t *testing.T) {
	input := `date("2016-01-01")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := "2016-01-01 00:00:00 +0000 UTC"
	if out.(time.Time) != time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_DateTime(t *testing.T) {
	input := `datetime("2016-01-01 12:00:00")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := "2016-01-01 12:00:00 +0000 UTC"
	if out.(time.Time) != time.Date(2016, 1, 1, 12, 0, 0, 0, time.UTC) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Deg2Rad(t *testing.T) {
	input := `deg2rad(180)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 3.141592653589793
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_DiffDate(t *testing.T) {
	input := `diffdate(date("2016-01-01"), date("2016-01-02"))`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 24 * time.Hour
	if out.(time.Duration) != expected {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_DiffTime(t *testing.T) {
	input := `difftime(datetime("2016-01-01 12:00:00"), datetime("2016-01-01 12:00:01"))`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 1 * time.Second
	if out.(time.Duration) != expected {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Div(t *testing.T) {
	input := `div(5, 5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 1
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_DivInt(t *testing.T) {
	input := `divint(5, 4)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 1
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Double(t *testing.T) {
	input := `double(5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 10
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Exp(t *testing.T) {
	input := `exp(5,2)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 25
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Filter(t *testing.T) {
	input := `filter(int[1,2,3,4,5], "x > 3")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(context.TODO())
	if err != nil {
		t.Error(err)
	}

	expected := []any{4, 5}

	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Floor(t *testing.T) {
	input := `floor(5.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 5
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Hypot(t *testing.T) {
	input := `hypot(3, 4)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 5
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Len(t *testing.T) {
	input := `len("hello")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 5
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Log(t *testing.T) {
	input := `log(100,10)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 2
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Log10(t *testing.T) {
	input := `log10(100)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 2
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Log2(t *testing.T) {
	input := `log2(8)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 3
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Map(t *testing.T) {
	input := `map(int[1,2,3,4,5], "_x * 2")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(context.TODO())
	if err != nil {
		t.Error(err)
	}

	expected := []any{2, 4, 6, 8, 10}

	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_MapSqrt(t *testing.T) {
	input := `map(int[1,4,9], "sqrt(_x)")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(context.TODO())
	if err != nil {
		t.Error(err)
	}

	expected := []any{1, 2, 3}

	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_MapCustom(t *testing.T) {
	RegisterFunction("custom", func(ctx context.Context, args ...any) (any, error) {
		var out int

		for i, x := range args {
			out = x.(int) * (i + 1)
		}

		return out, nil
	})

	input := `map(int[1,4,9], "custom(_i,_x)")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(context.TODO())
	if err != nil {
		t.Error(err)
	}

	expected := []any{2, 8, 18}

	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Max(t *testing.T) {
	input := `max(1,2,3,4,5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	expected := 5
	if !equalNumber(out, expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Mean(t *testing.T) {
	input := `mean(1,2,3,4,5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	if !equalNumber(out, 3) {
		t.Errorf("expected 3, got %v", out)
	}
}

func TestBif_Median(t *testing.T) {
	input := `median(1,2,3,4,5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	if !equalNumber(out, 3) {
		t.Errorf("expected 3, got %v", out)
	}
}

func TestBif_Min(t *testing.T) {
	input := `min(1,2,3,4,5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	if !equalNumber(out, 1) {
		t.Errorf("expected 1, got %v", out)
	}
}

func TestBif_Mod(t *testing.T) {
	input := `mod(5, 2)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	if !equalNumber(out, 1) {
		t.Errorf("expected 1, got %v", out)
	}
}

func TestBif_Mode(t *testing.T) {
	input := `mode(1,2,3,4,5,5,5,5,5,5,5,5,5,5,5,5,5,5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	if !equalNumber(out, 5) {
		t.Errorf("expected 5, got %v", out)
	}
}

func TestBif_Mul(t *testing.T) {
	input := `mul(5, 5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}

	if !equalNumber(out, 25) {
		t.Errorf("expected 25, got %v", out)
	}
}

func TestBif_Pow(t *testing.T) {
	input := `pow(5, 2)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 25) {
		t.Errorf("expected 25, got %v", out)
	}
}

func TestBif_Predict(t *testing.T) {
	t.Skip("not implemented")
}

func TestBif_Rad2Deg(t *testing.T) {
	input := `rad2deg(3.141592653589793)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 180) {
		t.Errorf("expected 180, got %v", out)
	}
}

func TestBif_Rand(t *testing.T) {
	input := `rand("int",10)`

	lexer := NewLexer(input)
	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(context.TODO())
	if err != nil {
		t.Error(err)
	}

	if out.(int) > 10 {
		t.Errorf("expected 1-10, got %v", out)
	}
}

func TestBif_Reduce(t *testing.T) {
	input := `reduce(int[1,2,3,4,5], "add")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 15) {
		t.Errorf("expected 15, got %v", out)
	}
}

func TestBif_ReduceAdd10(t *testing.T) {
	input := `reduce(int[1,2,3,4,5], "add", 10)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 25) {
		t.Errorf("expected 25, got %v", out)
	}
}

func TestBif_ReduceMul(t *testing.T) {
	input := `reduce(int[1,2,3,4,5], "mul")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 120) {
		t.Errorf("expected 120, got %v", out)
	}
}

func TestBif_ReduceMin(t *testing.T) {
	input := `reduce(int[1,2,3,4,5], "min")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 1) {
		t.Errorf("expected 1, got %v", out)
	}
}

func TestBif_Reverse(t *testing.T) {
	input := `reverse(int[1,2,3,4,5])`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	expected := []any{5, 4, 3, 2, 1}
	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Round(t *testing.T) {
	input := `round(5.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 6) {
		t.Errorf("expected 6, got %v", out)
	}
}

func TestBif_Root(t *testing.T) {
	input := `root(25,2)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 5) {
		t.Errorf("expected 5, got %v", out)
	}
}

func TestBif_Shuffle(t *testing.T) {
	input := `shuffle(int[1,2,3,4,5])`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if len(out.([]any)) != 5 {
		t.Errorf("expected 5, got %v", out)
	}
}

func TestBif_Sin(t *testing.T) {
	input := `sin(0.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 0.479425538604203) {
		t.Errorf("expected 0.479425538604203, got %v", out)
	}
}

func TestBif_Sinh(t *testing.T) {
	input := `sinh(0.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 0.5210953054937474) {
		t.Errorf("expected 0.5210953054937474, got %v", out)
	}
}

func TestBif_Slice(t *testing.T) {
	input := `slice(int[1,2,3,4,5], 1, 3)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	expected := []any{2, 3}
	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Sort(t *testing.T) {
	input := `sort(int[5,4,3,2,1])`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	expected := []any{1, 2, 3, 4, 5}
	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Sqrt(t *testing.T) {
	input := `sqrt(25)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 5) {
		t.Errorf("expected 5, got %v", out)
	}
}

func TestBif_StdDev(t *testing.T) {
	input := `stddev(1,2,3,4,5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 1.4142135623730951) {
		t.Errorf("expected 1.4142135623730951, got %v", out)
	}
}

func TestBif_Sub(t *testing.T) {
	input := `sub(5, 5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 0) {
		t.Errorf("expected 0, got %v", out)
	}
}

func TestBif_Sum(t *testing.T) {
	input := `sum(1,2,3,4,5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 15) {
		t.Errorf("expected 15, got %v", out)
	}
}

func TestBif_SumArray(t *testing.T) {
	input := `sum(1,2,3,4,5, int[1,2,3,4,5])`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 30) {
		t.Errorf("expected 15, got %v", out)
	}
}

func TestBif_Tan(t *testing.T) {
	input := `tan(0.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 0.5463024898437905) {
		t.Errorf("expected 0.5463024898437905, got %v", out)
	}
}

func TestBif_Tanh(t *testing.T) {
	input := `tanh(0.5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 0.46211715726000974) {
		t.Errorf("expected 0.46211715726000974, got %v", out)
	}
}

func TestBif_Time(t *testing.T) {
	input := `time("12:00:00")`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	expected := "0000-01-01 12:00:00 +0000 UTC"
	if out.(time.Time) != time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Unique(t *testing.T) {
	input := `unique(int[1,2,3,4,5,5,5,5,5,5,5,5,5,5,5,5,5,5])`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	expected := []any{1, 2, 3, 4, 5}
	if fmt.Sprintf("%v", out) != fmt.Sprintf("%v", expected) {
		t.Errorf("expected %v, got %v", expected, out)
	}
}

func TestBif_Variance(t *testing.T) {
	input := `variance(1,2,3,4,5)`
	lexer := NewLexer(input)

	p := NewParser(lexer)
	tree := p.Parse()

	out, err := tree.Evaluate(nil)
	if err != nil {
		t.Error(err)
	}
	if !equalNumber(out, 2.5) {
		t.Errorf("expected 2.5, got %v", out)
	}
}
