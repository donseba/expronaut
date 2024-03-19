package expronaut

import (
	"context"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/donseba/expronaut/llm"
	"io"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

type bifFunc func(context.Context, ...any) (any, error)
type bif map[string]bifFunc

var BuiltinFunctions = bif{}

func RegisterFunction(name string, function bifFunc) {
	BuiltinFunctions[name] = function
}

func init() {
	BuiltinFunctions = bif{}

	b := BuiltinFunctions

	// mathematical functions
	b["add"] = b.Add       // add two numbers
	b["sub"] = b.Sub       // subtract two numbers
	b["mul"] = b.Mul       // multiply two numbers
	b["div"] = b.Div       // divide two numbers
	b["divint"] = b.DivInt // divide two numbers and return the integer remainder 5//4=1
	b["mod"] = b.Mod       // modulo of two numbers
	b["exp"] = b.Exp       // raise a number to the power of another number
	b["sqrt"] = b.Sqrt     // square root of a number
	b["pow"] = b.Pow       // raise a number to the power of another number
	b["log"] = b.Log       // logarithm of a number
	b["log10"] = b.Log10   // base 10 logarithm of a number
	b["log2"] = b.Log2     // base 2 logarithm of a number
	b["sin"] = b.Sin       // sine of an angle in radians
	b["cos"] = b.Cos       // cosine of an angle in radians
	b["tan"] = b.Tan       // tangent of an angle in radians
	b["asin"] = b.Asin     // arc sine of a value
	b["acos"] = b.Acos     // arc cosine of a value
	b["atan"] = b.Atan     // arc tangent of a value
	b["sinh"] = b.Sinh     // hyperbolic sine of a number
	b["cosh"] = b.Cosh     // hyperbolic cosine of a number
	b["tanh"] = b.Tanh     // hyperbolic tangent of a number
	b["ceil"] = b.Ceil     // round a number up to the nearest integer
	b["floor"] = b.Floor   // round a number down to the nearest integer
	b["round"] = b.Round   // round a number to the nearest integer
	b["abs"] = b.Abs       // absolute value of a number
	b["double"] = b.Double // double a number
	b["root"] = b.Root     // nth root of a number

	b["hypot"] = b.Hypot     // hypotenuse of a right-angled triangle
	b["deg2rad"] = b.Deg2Rad // convert degrees to radians
	b["rad2deg"] = b.Rad2Deg // convert radians to degrees

	// statistical functions
	b["mean"] = b.Mean     // mean of two or more numbers
	b["median"] = b.Median // median of two or more numbers
	b["stddev"] = b.StdDev // standard deviation of two or more numbers
	b["max"] = b.Max       // maximum of two or more numbers
	b["min"] = b.Min       // minimum of two or more numbers

	// array functions
	b["filter"] = b.Filter   // filter an array based on a condition
	b["map"] = b.Map         // apply a function to each element of an array
	b["reduce"] = b.Reduce   // reduce an array to a single value
	b["sum"] = b.Sum         // sum of two or more numbers
	b["shuffle"] = b.Shuffle // shuffle an array
	b["concat"] = b.Concat   // concatenate two or more arrays
	b["reverse"] = b.Reverse // reverse an array
	b["sort"] = b.Sort       // sort an array
	b["unique"] = b.Unique   // remove duplicate elements from an array
	b["slice"] = b.Slice     // slice an array

	// random functions
	b["rand"] = b.Rand // generate a random number

	// date and time functions
	b["date"] = b.Date         // parse a string into a date
	b["time"] = b.Time         // parse a string into a time
	b["datetime"] = b.DateTime // parse a string into a date and time
	b["diffdate"] = b.DiffDate // difference between two dates
	b["difftime"] = b.DiffTime // difference between two times

	// utility functions
	b["len"] = b.Len // length of a string or array
	b["env"] = b.Env // get an environment variable

	// hashing functions
	b["sha256"] = b.Sha256 // SHA-256 hash
	b["sha512"] = b.Sha512 // SHA-512 hash

	// statistical functions
	b["mode"] = b.Mode         // mode of two or more numbers
	b["variance"] = b.Variance // variance of two or more numbers

	// monetary functions
	b["pv"] = b.Pv // present value of an investment at a specified rate of return
	b["fv"] = b.Fv // future value of an investment at a specified rate of return

	// AI functions
	b["ai"] = b.Ai           // call an AI provider to generate a response
	b["predict"] = b.Predict // predict the value of a time series
}

// Abs Calculates the absolute value of a number.
func (bif bif) Abs(ctx context.Context, args ...any) (any, error) {
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

// Acos Computes the arc cosine of a value; returns the angle in radians.
func (bif bif) Acos(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("acos function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Acos(float64(arg)), nil
	case float64:
		return math.Acos(arg), nil
	default:
		return nil, fmt.Errorf("acos function expects a number argument")
	}
}

// Add Adds two numbers together.
func (bif bif) Add(ctx context.Context, args ...any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("add function expects exactly two arguments got %d", len(args))
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

// Ai Calls an AI provider to generate a response.
func (bif bif) Ai(ctx context.Context, args ...any) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("ai function expects two or more arguments")
	}

	llmProvider, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("ai function expects a string argument")
	}

	switch llmProvider {
	case "gpt":
		return llm.ChatGPT(args[1:])
	}

	return nil, nil
}

// Asin Computes the arc sine of a value; returns the angle in radians.
func (bif bif) Asin(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("asin function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Asin(float64(arg)), nil
	case float64:
		return math.Asin(arg), nil
	default:
		return nil, fmt.Errorf("asin function expects a number argument")
	}
}

// Atan Computes the arc tangent of a value; returns the angle in radians.
func (bif bif) Atan(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("atan function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Atan(float64(arg)), nil
	case float64:
		return math.Atan(arg), nil
	default:
		return nil, fmt.Errorf("atan function expects a number argument")
	}
}

// Ceil Rounds a number up to the nearest integer.
func (bif bif) Ceil(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("ceil function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return arg, nil
	case float64:
		return int(math.Ceil(arg)), nil
	default:
		return nil, fmt.Errorf("ceil function expects a number argument")
	}
}

// Cos Calculates the cosine of an angle in radians.
func (bif bif) Cos(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("cos function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Cos(float64(arg)), nil
	case float64:
		return math.Cos(arg), nil
	default:
		return nil, fmt.Errorf("cos function expects a number argument")
	}
}

// Cosh Calculates the hyperbolic cosine of a number.
func (bif bif) Cosh(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("cosh function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Cosh(float64(arg)), nil
	case float64:
		return math.Cosh(arg), nil
	default:
		return nil, fmt.Errorf("cosh function expects a number argument")
	}
}

// Concat Combines two or more arrays into one.
func (bif bif) Concat(ctx context.Context, args ...any) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("concat function expects at least two arguments")
	}

	var result string
	for _, arg := range args {
		switch a := arg.(type) {
		case string:
			result += a
		default:
			return nil, fmt.Errorf("concat function expects string arguments")
		}
	}

	return result, nil
}

// Date Parses a string into a date.
func (bif bif) Date(ctx context.Context, args ...any) (any, error) {
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

// DateTime Parses a string into a date and time.
func (bif bif) DateTime(ctx context.Context, args ...any) (any, error) {
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

// Deg2Rad Converts degrees to radians.
func (bif bif) Deg2Rad(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("deg2rad function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return float64(arg) * math.Pi / 180, nil
	case float64:
		return arg * math.Pi / 180, nil
	default:
		return nil, fmt.Errorf("deg2rad function expects a number argument")
	}
}

// Div Divides two numbers.
func (bif bif) Div(ctx context.Context, args ...any) (any, error) {
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

// DiffDate Calculates the difference between two dates.
func (bif bif) DiffDate(ctx context.Context, args ...any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("diffdate function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case time.Time:
		switch b := args[1].(type) {
		case time.Time:
			if a.After(b) {
				return a.Sub(b), nil
			}

			return b.Sub(a), nil
		}
	}

	return nil, nil
}

// DivInt Divides two numbers and returns an integer.
func (bif bif) DivInt(ctx context.Context, args ...any) (any, error) {
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

// DiffTime Calculates the difference between two times.
func (bif bif) DiffTime(ctx context.Context, args ...any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("difftime function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case time.Time:
		switch b := args[1].(type) {
		case time.Time:
			if a.After(b) {
				return a.Sub(b), nil
			}
			return b.Sub(a), nil
		}
	}

	return nil, nil
}

// Double Doubles a number.
func (bif bif) Double(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("double function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		return a * 2, nil
	case float64:
		return a * 2, nil
	}

	return nil, nil
}

// Env Retrieves an environment variable.
func (bif bif) Env(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("env function expects a single argument")
	}

	key, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("env function expects a string argument")
	}

	return os.Getenv(key), nil
}

// Exp Raises a number to the power of another number.
func (bif bif) Exp(ctx context.Context, args ...any) (any, error) {
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

// Filter Filters an array based on a condition.
func (bif bif) Filter(ctx context.Context, args ...any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("filter function expects exactly two arguments: an array and an expression")
	}

	array, ok := args[0].([]any)
	if !ok {
		return nil, fmt.Errorf("first argument to filter must be an array")
	}

	expr, ok := args[1].(string)
	if !ok {
		return nil, fmt.Errorf("second argument to filter must be a string expression")
	}

	var filteredArray []any
	for _, element := range array {
		// Prepare the context for the expression evaluation
		ctx = SetVariables(ctx, map[string]any{"x": element})

		// Evaluate the expression in the context of the current element
		result, err := Evaluate(ctx, expr)
		if err != nil {
			return nil, fmt.Errorf("error evaluating expression '%s': %v", expr, err)
		}

		// Check if the result is true and include the element in the filtered array
		if include, ok := result.(bool); ok && include {
			filteredArray = append(filteredArray, element)
		}
	}

	return filteredArray, nil
}

// Floor Rounds a number down to the nearest integer.
func (bif bif) Floor(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("floor function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return arg, nil
	case float64:
		return int(math.Floor(arg)), nil
	default:
		return nil, fmt.Errorf("floor function expects a number argument")
	}
}

// Fv Calculates the future value of an investment at a specified rate of return.
func (bif bif) Fv(ctx context.Context, args ...any) (any, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("fv function expects exactly three arguments: present value, interest rate, and number of periods")
	}

	var (
		pv   float64
		rate float64
		n    float64
	)

	switch a := args[0].(type) {
	case int:
		pv = float64(a)
	case float64:
		pv = a
	default:
		return nil, fmt.Errorf("first argument (present value) must be a number")
	}

	switch a := args[1].(type) {
	case int:
		rate = float64(a)
	case float64:
		rate = a
	default:
		return nil, fmt.Errorf("second argument (interest rate) must be a number")
	}

	switch a := args[2].(type) {
	case int:
		n = float64(a)
	case float64:
		n = a
	default:
		return nil, fmt.Errorf("third argument (number of periods) must be a number")
	}

	// Calculate the future value
	fvr := math.Pow(1+rate, n)
	fv := pv * fvr

	return fv, nil
}

// Hypot Calculates the hypotenuse of a right-angled triangle.
func (bif bif) Hypot(ctx context.Context, args ...any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("hypot function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		switch b := args[1].(type) {
		case int:
			return math.Hypot(float64(a), float64(b)), nil
		case float64:
			return math.Hypot(float64(a), b), nil
		}
	case float64:
		switch b := args[1].(type) {
		case int:
			return math.Hypot(a, float64(b)), nil
		case float64:
			return math.Hypot(a, b), nil
		}
	}

	return nil, nil
}

// Len Returns the length of a string or array.
func (bif bif) Len(ctx context.Context, args ...any) (any, error) {
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

// Log Calculates the logarithm of a number.
func (bif bif) Log(ctx context.Context, args ...any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("log function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		switch b := args[1].(type) {
		case int:
			return math.Log(float64(a)) / math.Log(float64(b)), nil
		case float64:
			return math.Log(float64(a)) / math.Log(b), nil
		}
	case float64:
		switch b := args[1].(type) {
		case int:
			return math.Log(a) / math.Log(float64(b)), nil
		case float64:
			return math.Log(a) / math.Log(b), nil
		}
	}

	return nil, nil
}

// Log10 Calculates the base 10 logarithm of a number.
func (bif bif) Log10(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("log10 function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Log10(float64(arg)), nil
	case float64:
		return math.Log10(arg), nil
	default:
		return nil, fmt.Errorf("log10 function expects a number argument")
	}
}

// Log2 Calculates the base 2 logarithm of a number.
func (bif bif) Log2(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("log2 function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Log2(float64(arg)), nil
	case float64:
		return math.Log2(arg), nil
	default:
		return nil, fmt.Errorf("log2 function expects a number argument")
	}
}

// Map Applies a function to each element of an array.
func (bif bif) Map(ctx context.Context, args ...any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("map function expects exactly two arguments: an array and an expression")
	}

	array, ok := args[0].([]any)
	if !ok {
		return nil, fmt.Errorf("first argument to map must be an array")
	}

	expr, ok := args[1].(string)
	if !ok {
		return nil, fmt.Errorf("second argument to map must be a string expression")
	}

	var mappedArray []any
	for i, element := range array {
		// Prepare the context for the expression evaluation
		ctx = SetVariables(ctx, map[string]any{"_i": i, "_x": element})

		//// would it be logical to also pass the next element in the array?
		//ctx = SetVariables(ctx, map[string]any{"_y": element})

		// Evaluate the expression in the context of the current element
		transformedElement, err := Evaluate(ctx, expr)
		if err != nil {
			return nil, fmt.Errorf("error evaluating expression '%s': %v", expr, err)
		}

		// Add the result of the transformation to the mapped array
		mappedArray = append(mappedArray, transformedElement)
	}

	return mappedArray, nil
}

// Max Returns the maximum of two or more numbers.
func (bif bif) Max(ctx context.Context, args ...any) (any, error) {
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

// Mean Calculates the mean of two or more numbers.
func (bif bif) Mean(ctx context.Context, args ...any) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("mean function expects at least one argument")
	}

	var sum float64
	for _, arg := range args {
		switch a := arg.(type) {
		case int:
			sum += float64(a)
		case float64:
			sum += a
		default:
			return nil, fmt.Errorf("mean function expects number arguments")
		}
	}

	return sum / float64(len(args)), nil
}

// Median Calculates the median of two or more numbers.
func (bif bif) Median(ctx context.Context, args ...any) (any, error) {
	var median = func(a []float64) (float64, error) {
		if len(a) == 0 {
			return 0, nil
		}

		if len(a) == 1 {
			return a[0], nil
		}

		if len(a)%2 == 0 {
			return (a[len(a)/2-1] + a[len(a)/2]) / 2, nil
		}

		return a[len(a)/2], nil
	}

	if len(args) < 1 {
		return nil, fmt.Errorf("median function expects at least one argument")
	}

	var a []float64
	for _, arg := range args {
		switch arg := arg.(type) {
		case int:
			a = append(a, float64(arg))
		case float64:
			a = append(a, arg)
		default:
			return nil, fmt.Errorf("median function expects number arguments")
		}
	}

	return median(a)
}

// Min Returns the minimum of two or more numbers.
func (bif bif) Min(ctx context.Context, args ...any) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("min function expects at least one argument")
	}

	// Initialize m with the first argument's value.
	var m float64
	var initialized bool
	for _, arg := range args {
		switch a := arg.(type) {
		case int:
			if !initialized {
				m = float64(a)
				initialized = true
				continue
			}
			if float64(a) < m {
				m = float64(a)
			}
		case float64:
			if !initialized {
				m = a
				initialized = true
				continue
			}
			if a < m {
				m = a
			}
		default:
			return nil, fmt.Errorf("min function expects number arguments, got %T", arg)
		}
	}

	if !initialized {
		return nil, fmt.Errorf("min function requires at least one number argument")
	}

	// Ensure the return type matches the input.
	// If all inputs were int, return an int.
	allInt := true
	for _, arg := range args {
		if _, ok := arg.(float64); ok {
			allInt = false
			break
		}
	}
	if allInt {
		return int(m), nil
	}

	return m, nil
}

// Mod Returns the remainder of a division.
func (bif bif) Mod(ctx context.Context, args ...any) (any, error) {
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

// Mode Returns the mode of two or more numbers.
func (bif bif) Mode(ctx context.Context, args ...any) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("mode function expects at least one argument")
	}

	var mode = func(a []float64) (float64, error) {
		if len(a) == 0 {
			return 0, nil
		}

		m := make(map[float64]int)
		for _, v := range a {
			m[v]++
		}

		var maxim float64
		var maxCount int
		for k, v := range m {
			if v > maxCount {
				maxim = k
				maxCount = v
			}
		}

		return maxim, nil
	}

	var a []float64
	for _, arg := range args {
		switch arg := arg.(type) {
		case int:
			a = append(a, float64(arg))
		case float64:
			a = append(a, arg)
		default:
			return nil, fmt.Errorf("mode function expects number arguments")
		}
	}

	return mode(a)
}

// Mul Multiplies two numbers together.
func (bif bif) Mul(ctx context.Context, args ...any) (any, error) {
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

// Pow Raises a number to the power of another number.
func (bif bif) Pow(ctx context.Context, args ...any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("pow function expects exactly two arguments")
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

// Predict Predicts the next word in a sentence.
func (bif bif) Predict(ctx context.Context, args ...any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("predict function expects exactly two arguments")
	}

	llmProvider, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("predict function expects a string argument")
	}

	switch llmProvider {
	case "gpt":
		return nil, errors.New("predict function not implemented yet")
	default:
		return nil, fmt.Errorf("predict function expects a string argument")
	}

}

// Pv Calculates the present value of an investment at a specified rate of return.
func (bif bif) Pv(ctx context.Context, args ...any) (any, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("pv function expects exactly three arguments")
	}

	// fv float64, r float64, n float64
	var (
		fv float64
		n  float64
		r  float64
	)

	switch a := args[0].(type) {
	case int:
		fv = float64(a)
	case float64:
		fv = a
	default:
		return nil, fmt.Errorf("pv function expects a number argument")
	}

	switch a := args[1].(type) {
	case int:
		r = float64(a)
	case float64:
		r = a
	default:
		return nil, fmt.Errorf("pv function expects a number argument")
	}

	switch a := args[2].(type) {
	case int:
		n = float64(a)
	case float64:
		n = a
	default:
		return nil, fmt.Errorf("pv function expects a number argument")
	}

	// Calculate (1 + r)^n using the Pow method
	denominator, err := bif.Pow(ctx, 1+r, n)
	if err != nil {
		return 0, err
	}

	// Divide FV by (1 + r)^n
	pv, err := bif.Div(ctx, fv, denominator)
	if err != nil {
		return 0, err
	}

	return pv, nil
}

// Rad2Deg Converts radians to degrees.
func (bif bif) Rad2Deg(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("rad2deg function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return float64(arg) * 180 / math.Pi, nil
	case float64:
		return arg * 180 / math.Pi, nil
	default:
		return nil, fmt.Errorf("rad2deg function expects a number argument")
	}
}

// Rand Generates a random number.
func (bif bif) Rand(ctx context.Context, args ...any) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("rand function expects no arguments")
	}

	rtype, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("rand function expects a string argument")
	}

	n := 0
	if len(args) > 1 {
		n, ok = args[1].(int)
		if !ok {
			return nil, fmt.Errorf("rand function expects an int argument")
		}
	}

	if n > 0 {
		switch rtype {
		case "int":
			return rand.Intn(n), nil
		case "float64":
			return rand.Float64() * float64(n), nil
		}
	} else {
		switch rtype {
		case "int":
			return rand.Int(), nil
		case "float64":
			return rand.Float64(), nil
		}
	}

	return nil, nil
}

// Reduce Reduces an array to a single value.
func (bif bif) Reduce(ctx context.Context, args ...any) (any, error) {
	if len(args) < 2 || len(args) > 3 {
		return nil, fmt.Errorf("reduce function expects two or three arguments: an array, a function, and an optional initial value")
	}

	array, ok := args[0].([]any)
	if !ok || len(array) == 0 {
		return nil, fmt.Errorf("first argument to reduce must be a non-empty array")
	}

	var fun bifFunc
	switch arg := args[1].(type) {
	case string:
		fun, ok = bif[arg]
		if !ok {
			return nil, fmt.Errorf("function %s not supported", arg)
		}
	case bifFunc:
		fun = arg
	case func(context.Context, ...any) (any, error):
		fun = arg
	default:
		return nil, fmt.Errorf("second argument should be a function identifier or a function")
	}

	var accumulator any
	startIdx := 0
	if len(args) == 3 {
		// Use the third argument as the initial value if provided
		accumulator = args[2]
	} else {
		// Otherwise, use the first element of the array as the initial value
		accumulator = array[0]
		startIdx = 1 // Start reducing from the second element
	}

	for _, element := range array[startIdx:] {
		var err error
		in := accumulator
		accumulator, err = fun(ctx, in, element)
		if err != nil {
			return nil, err
		}
	}

	return accumulator, nil
}

// Reverse Reverses a string or array.
func (bif bif) Reverse(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("reverse function expects a single argument")
	}

	switch arg := args[0].(type) {
	case string:
		runes := []rune(arg)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes), nil
	case []any:
		for i, j := 0, len(arg)-1; i < j; i, j = i+1, j-1 {
			arg[i], arg[j] = arg[j], arg[i]
		}
		return arg, nil
	default:
		return nil, fmt.Errorf("reverse function expects a string or array argument")
	}
}

// Root Calculates the nth root of a number.
func (bif bif) Root(ctx context.Context, args ...any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("root function expects exactly two arguments")
	}

	switch a := args[0].(type) {
	case int:
		switch b := args[1].(type) {
		case int:
			return math.Pow(float64(a), 1/float64(b)), nil
		case float64:
			return math.Pow(float64(a), 1/b), nil
		}
	case float64:
		switch b := args[1].(type) {
		case int:
			return math.Pow(a, 1/float64(b)), nil
		case float64:
			return math.Pow(a, 1/b), nil
		}
	}

	return nil, nil
}

// Round Rounds a number to the nearest integer.
func (bif bif) Round(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("round function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return arg, nil
	case float64:
		return int(math.Round(arg)), nil
	default:
		return nil, fmt.Errorf("round function expects a number argument")
	}
}

// Sha256 Calculates the SHA-256 hash of the input
func (bif bif) Sha256(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("sha256 function expects a single argument")
	}

	switch arg := args[0].(type) {
	case string:
		h := sha256.New()
		_, _ = io.WriteString(h, arg)
		return string(h.Sum(nil)), nil
	case int:
		h := sha256.New()
		_, _ = io.WriteString(h, fmt.Sprintf("%d", arg))
		return string(h.Sum(nil)), nil
	case float64:
		h := sha256.New()
		_, _ = io.WriteString(h, fmt.Sprintf("%f", arg))
		return string(h.Sum(nil)), nil
	default:
		return nil, fmt.Errorf("sha256 function expects a string argument")
	}
}

// Sha512 Calculates the SHA-512 hash of the input.
func (bif bif) Sha512(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("sha512 function expects a single argument")
	}

	switch arg := args[0].(type) {
	case string:
		h := sha512.New()
		_, _ = io.WriteString(h, arg)
		return string(h.Sum(nil)), nil
	case int:
		h := sha512.New()
		_, _ = io.WriteString(h, fmt.Sprintf("%d", arg))
		return string(h.Sum(nil)), nil
	case float64:
		h := sha512.New()
		_, _ = io.WriteString(h, fmt.Sprintf("%f", arg))
		return string(h.Sum(nil)), nil
	default:
		return nil, fmt.Errorf("sha512 function expects a string argument")
	}
}

// Shuffle Shuffles an array.
func (bif bif) Shuffle(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("shuffle function expects a single argument")
	}

	switch arg := args[0].(type) {
	case []any:
		shuffled := make([]any, len(arg))
		perm := rand.Perm(len(arg))
		for i, v := range perm {
			shuffled[v] = arg[i]
		}
		return shuffled, nil
	default:
		return nil, fmt.Errorf("shuffle function expects an array argument")
	}
}

// Sin Calculates the sine of an angle in radians.
func (bif bif) Sin(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("sin function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Sin(float64(arg)), nil
	case float64:
		return math.Sin(arg), nil
	default:
		return nil, fmt.Errorf("sin function expects a number argument")
	}
}

// Sinh Calculates the hyperbolic sine of a number.
func (bif bif) Sinh(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("sinh function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Sinh(float64(arg)), nil
	case float64:
		return math.Sinh(arg), nil
	default:
		return nil, fmt.Errorf("sinh function expects a number argument")
	}
}

// Slice Returns a portion of an array.
func (bif bif) Slice(ctx context.Context, args ...any) (any, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("slice function expects exactly three arguments")
	}

	switch arg := args[0].(type) {
	case []any:
		start, ok := args[1].(int)
		if !ok {
			return nil, fmt.Errorf("slice function expects an int as the second argument")
		}

		end, ok := args[2].(int)
		if !ok {
			return nil, fmt.Errorf("slice function expects an int as the third argument")
		}

		return arg[start:end], nil
	default:
		return nil, fmt.Errorf("slice function expects an array argument")
	}
}

// Sort Sorts an array.
func (bif bif) Sort(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("sort function expects a single argument")
	}

	switch arg := args[0].(type) {
	case []any:
		sorted := make([]any, len(arg))
		copy(sorted, arg)

		switch sorted[0].(type) {
		case int:
			ints := make([]int, len(sorted))
			for i, v := range sorted {
				ints[i] = v.(int)
			}
			sort.Ints(ints)
			return ints, nil
		case float64:
			floats := make([]float64, len(sorted))
			for i, v := range sorted {
				floats[i] = v.(float64)
			}
			sort.Float64s(floats)
			return floats, nil
		case string:
			strings := make([]string, len(sorted))
			for i, v := range sorted {
				strings[i] = v.(string)
			}

			sort.Strings(strings)
			return strings, nil
		}
	default:
		return nil, fmt.Errorf("sort function expects an array argument")
	}

	return nil, nil
}

// StdDev Calculates the standard deviation of two or more numbers.
func (bif bif) StdDev(ctx context.Context, args ...any) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("stddev function expects at least one argument")
	}

	var sum float64
	for _, arg := range args {
		switch a := arg.(type) {
		case int:
			sum += float64(a)
		case float64:
			sum += a
		default:
			return nil, fmt.Errorf("stddev function expects number arguments")
		}
	}

	mean := sum / float64(len(args))

	var variance float64
	for _, arg := range args {
		switch a := arg.(type) {
		case int:
			variance += math.Pow(float64(a)-mean, 2)
		case float64:
			variance += math.Pow(a-mean, 2)
		}
	}

	return math.Sqrt(variance / float64(len(args))), nil
}

// Sub Subtracts two numbers.
func (bif bif) Sub(ctx context.Context, args ...any) (any, error) {
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

// Sqrt Calculates the square root of a number.
func (bif bif) Sqrt(ctx context.Context, args ...any) (any, error) {
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

// Sum Returns the sum of two or more numbers.
func (bif bif) Sum(ctx context.Context, args ...any) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("sum function expects a single argument")
	}

	var sum float64

	for _, arg := range args {
		switch a := arg.(type) {
		case int:
			sum += float64(a)
		case float64:
			sum += a
		case []any:
			ssum, _ := bif.Sum(ctx, a...)
			sum += ssum.(float64)
		default:
			return nil, fmt.Errorf("sum function expects number arguments")
		}
	}

	return sum, nil
}

// Unique Returns unique elements from an array.
func (bif bif) Unique(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("unique function expects a single argument")
	}

	switch arg := args[0].(type) {
	case []any:
		unique := make(map[any]bool)
		var result []any
		for _, a := range arg {
			if !unique[a] {
				unique[a] = true
				result = append(result, a)
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unique function expects an array argument")
	}
}

// Tan Calculates the tangent of an angle in radians.
func (bif bif) Tan(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("tan function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Tan(float64(arg)), nil
	case float64:
		return math.Tan(arg), nil
	default:
		return nil, fmt.Errorf("tan function expects a number argument")
	}
}

// Tanh Calculates the hyperbolic tangent of a number.
func (bif bif) Tanh(ctx context.Context, args ...any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("tanh function expects a single argument")
	}

	switch arg := args[0].(type) {
	case int:
		return math.Tanh(float64(arg)), nil
	case float64:
		return math.Tanh(arg), nil
	default:
		return nil, fmt.Errorf("tanh function expects a number argument")
	}
}

// Time Parses a string into a time.
func (bif bif) Time(ctx context.Context, args ...any) (any, error) {
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

// Variance Calculates the variance of two or more numbers.
func (bif bif) Variance(ctx context.Context, args ...any) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("variance function expects at least two arguments")
	}

	var sum float64
	for _, arg := range args {
		switch a := arg.(type) {
		case int:
			sum += float64(a)
		case float64:
			sum += a
		default:
			return nil, fmt.Errorf("variance function expects number arguments")
		}
	}

	mean := sum / float64(len(args))

	var variance float64
	for _, arg := range args {
		switch a := arg.(type) {
		case int:
			variance += math.Pow(float64(a)-mean, 2)
		case float64:
			variance += math.Pow(a-mean, 2)
		}
	}

	// Use len(args) - 1 for sample variance
	return variance / float64(len(args)-1), nil
}
