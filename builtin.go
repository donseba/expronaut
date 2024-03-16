package expronaut

import (
	"fmt"
	"math"
	"time"
)

type bifFunc func([]any) (any, error)
type bif map[string]bifFunc

var BuiltinFunctions = bif{}

func RegisterFunction(name string, function bifFunc) {
	BuiltinFunctions[name] = function
}

func init() {
	BuiltinFunctions = bif{}

	b := BuiltinFunctions

	b["abs"] = b.Abs
	b["add"] = b.Add
	b["date"] = b.Date
	b["datetime"] = b.DateTime
	b["div"] = b.Div
	b["divint"] = b.DivInt
	b["double"] = b.Double
	b["exp"] = b.Exp
	b["filter"] = b.Filter
	b["len"] = b.Len
	b["map"] = b.Map
	b["max"] = b.Max
	b["min"] = b.Min
	b["mod"] = b.Mod
	b["mul"] = b.Mul
	b["reduce"] = b.Reduce
	b["sub"] = b.Sub
	b["sqrt"] = b.Sqrt
	b["sum"] = b.Sum
	b["time"] = b.Time
}

func (b bif) Date(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("date function expects a single argument")
	}

	sd, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("date function expects a string argument")
	}

	t, err := time.Parse(time.DateOnly, sd)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (b bif) DateTime(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("datetime function expects a single argument")
	}

	sd, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("datetime function expects a string argument")
	}

	t, err := time.Parse(time.DateTime, sd)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (b bif) Time(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("time function expects a single argument")
	}

	st, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("time function expects a string argument")
	}

	t, err := time.Parse(time.TimeOnly, st)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (b bif) Sub(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("sub function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		switch b := args[1].(type) {
		case int:
			return a - b, nil
		case float64:
			return float64(a) - b, nil
		}
	case float64:
		switch b := args[1].(type) {
		case int:
			return a - float64(b), nil
		case float64:
			return a - b, nil
		}
	}
	return nil, nil
}

func (b bif) Sqrt(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("time function expects a single argument")
	}

	switch arg := args[0].(type) {
	case float64:
		return math.Sqrt(arg), nil
	case int:
		return math.Sqrt(float64(arg)), nil
	default:
		return nil, fmt.Errorf("sqrt function expects a number argument")
	}
}

func (b bif) Max(args []any) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("max function expects at least one argument")
	}

	var m float64
	for i, arg := range args {
		switch a := arg.(type) {
		case int:
			if i == 0 {
				m = float64(a)
				continue
			}

			if float64(a) > m {
				m = float64(a)
			}
		case float64:
			if a > m {
				m = a
			}
		default:
			return nil, fmt.Errorf("max function expects number arguments")
		}
	}

	return m, nil
}

func (b bif) Min(args []any) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("min function expects at least one argument")
	}

	var m float64
	for i, arg := range args {
		switch a := arg.(type) {
		case int:
			if i == 0 {
				m = float64(a)
				continue
			}

			if float64(a) < m {
				m = float64(a)
			}
		case float64:
			if a < m {
				m = a
			}
		default:
			return nil, fmt.Errorf("min function expects number arguments")
		}
	}

	return m, nil
}

func (b bif) Mod(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("mod function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		switch b := args[1].(type) {
		case int:
			return a % b, nil
		case float64:
			return a % int(b), nil
		}
	case float64:
		switch b := args[1].(type) {
		case int:
			return int(a) % b, nil
		case float64:
			return int(a) % int(b), nil
		}
	}

	return nil, nil
}

func (b bif) Mul(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("mul function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		switch b := args[1].(type) {
		case int:
			return a * b, nil
		case float64:
			return float64(a) * b, nil
		}
	case float64:
		switch b := args[1].(type) {
		case int:
			return a * float64(b), nil
		case float64:
			return a * b, nil
		}
	}

	return nil, nil
}

func (b bif) Abs(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("abs function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return int(math.Abs(float64(arg))), nil
	case float64:
		return math.Abs(arg), nil
	default:
		return nil, fmt.Errorf("abs function expects a number argument")
	}
}

func (b bif) Len(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("len function expects a single argument")
	}

	switch arg := args[0].(type) {
	case string:
		return len(arg), nil
	case []any:
		return len(arg), nil
	default:
		return nil, fmt.Errorf("len function expects a string or array argument")
	}
}

func (b bif) Sum(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("sum function expects a single argument")
	}

	switch arg := args[0].(type) {
	case []any:
		var sum float64
		for _, a := range arg {
			switch a := a.(type) {
			case int:
				sum += float64(a)
			case float64:
				sum += a
			default:
				return nil, fmt.Errorf("sum function expects an array of number arguments")
			}
		}
		return sum, nil
	default:
		return nil, fmt.Errorf("sum function expects an array argument")
	}
}

func (b bif) Filter(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("filter function expects two arguments")
	}

	switch arg := args[0].(type) {
	case []any:
		var filtered []any
		for _, a := range arg {
			switch a := a.(type) {
			case int:
				if a == args[1] {
					filtered = append(filtered, a)
				}
			case float64:
				if a == args[1] {
					filtered = append(filtered, a)
				}
			case string:
				if a == args[1] {
					filtered = append(filtered, a)
				}
			}
		}
		return filtered, nil
	default:
		return nil, fmt.Errorf("filter function expects an array argument")
	}
}

func (b bif) Reduce(args []any) (any, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("reduce function expects exactly three arguments: an array, a function, and an initial value")
	}

	array, ok := args[0].([]any)
	if !ok {
		return nil, fmt.Errorf("first argument to reduce must be an array")
	}

	var fun bifFunc
	switch arg := args[1].(type) {
	case string:
		fun, ok = b[arg]
		if !ok {
			return nil, fmt.Errorf("function %s not supported", arg)
		}
	case bifFunc:
		fun = arg
	case func([]any) (any, error):
		fun = arg
	default:
		return nil, fmt.Errorf("second argument should be `func([]any) (any, error)` got %T", arg)
	}

	accumulator := args[2]

	for _, element := range array {
		var err error
		accumulator, err = fun([]any{accumulator, element})
		if err != nil {
			return nil, err
		}
	}

	return accumulator, nil
}

func (b bif) Map(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("map function expects exactly two arguments: an array and a function")
	}

	array, ok := args[0].([]any)
	if !ok {
		return nil, fmt.Errorf("first argument to map must be an array")
	}

	var fun bifFunc
	switch arg := args[1].(type) {
	case string:
		fun, ok = b[arg]
		if !ok {
			return nil, fmt.Errorf("function %s not supported", arg)
		}
	case bifFunc:
		fun = arg
	case func([]any) (any, error):
		fun = arg
	default:
		return nil, fmt.Errorf("second argument should be `func([]any) (any, error)` got %T", arg)
	}

	var result []any
	for _, element := range array {
		r, err := fun([]any{element})
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}

func (b bif) Div(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("div function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		switch b := args[1].(type) {
		case int:
			return a / b, nil
		case float64:
			return float64(a) / b, nil
		}
	case float64:
		switch b := args[1].(type) {
		case int:
			return a / float64(b), nil
		case float64:
			return a / b, nil
		}
	}
	return nil, nil
}

func (b bif) DivInt(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("div function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		switch b := args[1].(type) {
		case int:
			return a / b, nil
		case float64:
			return int(a) / int(b), nil
		}
	case float64:
		switch b := args[1].(type) {
		case int:
			return int(a) / int(b), nil
		case float64:
			return int(a) / int(b), nil
		}
	}
	return nil, nil
}

func (b bif) Double(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("add function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		return a * 2, nil
	case float64:
		return a * 2, nil
	}

	return nil, nil
}

func (b bif) Exp(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("exp function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		switch b := args[1].(type) {
		case int:
			return math.Pow(float64(a), float64(b)), nil
		case float64:
			return math.Pow(float64(a), b), nil
		}
	case float64:
		switch b := args[1].(type) {
		case int:
			return math.Pow(a, float64(b)), nil
		case float64:
			return math.Pow(a, b), nil
		}
	}

	return nil, nil
}

func (b bif) Add(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("add function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		switch b := args[1].(type) {
		case int:
			return a + b, nil
		case float64:
			return float64(a) + b, nil
		}
	case float64:
		switch b := args[1].(type) {
		case int:
			return a + float64(b), nil
		case float64:
			return a + b, nil
		}
	case string:
		switch b := args[1].(type) {
		case string:
			return a + b, nil
		}
	}
	return nil, nil
}
