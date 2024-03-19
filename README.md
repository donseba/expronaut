# Expronaut: Magic in Expression Evaluation

Welcome to Expronaut, a fun and dynamic Go package designed for parsing and evaluating expressions with a touch of magic. 
Whether you're crafting conditions for template rendering or just dabbling in the alchemy of logic and math, Expronaut is your companion on this adventurous journey through the realms of syntax and semantics.

## Features

- **Dynamic Expression Parsing:** Dive into expressions with variables, nested properties, and an assortment of operators. From simple arithmetic to complex boolean logic, Expronaut ~~understands it all~~, tries to understand.
- **Flexible Variable Context:** Whether you're working with flat landscapes or exploring the depths of nested objects, Expronaut navigates through your data with ease, bringing context to your expressions.
- **Customizable Evaluation:** Tailor the evaluation context to suit your adventure. Pass in variables and watch as Expronaut conjures up the results you seek.

## Getting Started

Embark on your journey with Expronaut by incorporating it into your Go projects. Here's how to get started:

### Installation

Ensure you have Go installed on your system, then fetch the Expronaut package:

```sh
go get github.com/donseba/expronaut
```

## Usage Examples

### complex example

```go
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
```
This tree visually represents how the operations in the expression are structured and the order in which they would be evaluated, starting from the bottom operations moving up.
```markdown
                              RIGHT_SHIFT
                              /          \
                             /            \
                          LEFT_SHIFT       MODULO
                           /     \          /   \
                          /       \        /     \
                       MINUS       1      2       3
                      /    \
                     /      \
                  PLUS     DIVIDE_INTEGER
                  /   \        /        \
                -2    MULTIPLY          EXPONENT
                        /  \            /      \
                       3    4          2        2
```


### Converting expressions to Go Template Strings
please not this only works for the most basic expressions, for more complex expressions there is still lost to do to make it work.

```go
func TestNewParserVariableBool(t *testing.T) {
   input := `false == foo || ( bar >= baz )`

   lexer := NewLexer(input)

   p := NewParser(lexer)
   tree := p.Parse()

   t.Logf(`%s`, tree.GoTemplate()) // or ( eq false .foo ) ( ge .bar .baz )
}

```

Expronaut transforms your intricate expressions into results or Go template strings, ready for dynamic rendering.

## Supported Operators

### Arithmetic Operators

- **+ (Addition):** Adds two numbers.
- **- (Subtraction):** Subtracts the second number from the first.
- **\* (Multiplication):** Multiplies two numbers.
- **/ (Division):** Divides the first number by the second. Performs floating-point division.
- **// (Integer Division):** Divides the first number by the second, discarding any remainder to return an integer result.
- **% (Modulo):** Returns the remainder of dividing the first number by the second.
- **^ (Exponentiation):** Raises the first number to the power of the second. ( ** is a valid alternative.)

### Bitwise Operators

- **<< (Left Shift):** Shifts the first operand left by the number of bits specified by the second operand.
- **>> (Right Shift):** Shifts the first operand right by the number of bits specified by the second operand.

### Logical Operators

- **&& (Logical AND):** Returns true if both operands are true.
- **|| (Logical OR):** Returns true if at least one of the operands is true.

### Comparison Operators

- **== (Equal):** Returns true if the operands are equal.
- **!= (Not Equal):** Returns true if the operands are not equal.
- **< (Less Than):** Returns true if the first operand is less than the second.
- **<= (Less Than or Equal To):** Returns true if the first operand is less than or equal to the second.
- **> (Greater Than):** Returns true if the first operand is greater than the second.
- **>= (Greater Than or Equal To):** Returns true if the first operand is greater than or equal to the second.

### Builtin functions 
- **add (Addition):** Adds two numbers (Considered as a function call, `add(1, 2)`). The first argument is the first number. The second argument is the second number.
- **sub (Subtraction):** Subtracts two numbers (Considered as a function call, `sub(5, 3)`). The first argument is the first number. The second argument is the second number.
- **mul (Multiplication):** Multiplies two numbers (Considered as a function call, `mul(2, 3)`). The first argument is the first number. The second argument is the second number.
- **div (Division):** Divides two numbers (Considered as a function call, `div(6, 2)`). The first argument is the first number. The second argument is the second number.
- **divint (Integer Division):** Divides two numbers and returns the integer result (Considered as a function call, `divint(5, 2)`). The first argument is the first number. The second argument is the second number.
- **mod (Modulo):** Calculates the remainder of dividing two numbers (Considered as a function call, `mod(5, 2)`). The first argument is the first number. The second argument is the second number.
- **exp (Exponentiation):** Raises a number to the power of another number (Considered as a function call, `exp(2, 3)`). The first argument is the base. The second argument is the exponent.
- **sqrt (Square Root):** Calculates the square root of a number (Considered as a function call, sqrt(x)).
- **pow (Power):** Raises a number to the power of another number (Considered as a function call, `pow(2, 3)`). The first argument is the base. The second argument is the exponent.
- **log (Logarithm):** Calculates the natural logarithm of a number (Considered as a function call, `log(10)`). The argument is the number.
- **log10 (Logarithm Base 10):** Calculates the base 10 logarithm of a number (Considered as a function call, `log10(100)`). The argument is the number.
- **log2 (Logarithm Base 2):** Calculates the base 2 logarithm of a number (Considered as a function call, `log2(8)`). The argument is the number.
- **sin (Sine):** Calculates the sine of an angle in radians (Considered as a function call, `sin(3.14159)`). The argument is the number of radians.
- **cos (Cosine):** Calculates the cosine of an angle in radians (Considered as a function call, `cos(3.14159)`). The argument is the number of radians.
- **tan (Tangent):** Calculates the tangent of an angle in radians (Considered as a function call, `tan(3.14159)`). The argument is the number of radians.
- **asin (Arc Sine):** Calculates the arc sine of a value (Considered as a function call, `asin(1)`). The argument is the value.
- **acos (Arc Cosine):** Calculates the arc cosine of a value (Considered as a function call, `acos(1)`). The argument is the value.
- **atan (Arc Tangent):** Calculates the arc tangent of a value (Considered as a function call, `atan(1)`). The argument is the value.
- **sinh (Hyperbolic Sine):** Calculates the hyperbolic sine of a number (Considered as a function call, `sinh(1)`). The argument is the number.
- **cosh (Hyperbolic Cosine):** Calculates the hyperbolic cosine of a number (Considered as a function call, `cosh(1)`). The argument is the number.
- **tanh (Hyperbolic Tangent):** Calculates the hyperbolic tangent of a number (Considered as a function call, `tanh(1)`). The argument is the number.
- **ceil (Ceiling):** Rounds a number up to the nearest integer (Considered as a function call, `ceil(3.14)`). The argument is the number.
- **floor (Floor):** Rounds a number down to the nearest integer (Considered as a function call, `floor(3.14)`). The argument is the number.
- **round (Round):** Rounds a number to the nearest integer (Considered as a function call, `round(3.14)`). The argument is the number.
- **abs (Absolute):** Calculates the absolute value of a number (Considered as a function call, `abs(-5)`). The argument is the number.
- **double (Double):** Doubles a number (Considered as a function call, `double(5)`). The argument is the number.
- **root (Root):** Calculates the nth root of a number (Considered as a function call, `root(27, 3)`). The first argument is the number. The second argument is the root.
- 
- **rand (Random):** Generates a random number (Considered as a function call, `rand(1, 10)`). The first argument is the minimum value. The second argument is the maximum value.
- **len (Length):** Returns the length of a string or array (Considered as a function call, `len("hello")`).
- **env (Environment):** Gets an environment variable (Considered as a function call, `env("HOME")`). The argument is the environment variable to get.


### statistical functions
- **mode (Mode):** Calculates the mode of a list of numbers (Considered as a function call, `mode(int[1,2,3,4,5,5,4,3,2,1])`). The argument is the list of numbers.
- **variance (Variance):** Calculates the variance of a list of numbers (Considered as a function call, `variance(int[1,2,3,4,5])`). The argument is the list of numbers.

### trigonometric functions
- **hypot (Hypotenuse):** Calculates the hypotenuse of a right-angled triangle (Considered as a function call, `hypot(3, 4)`). The first argument is the length of the first side. The second argument is the length of the second side.
- **deg2rad (Degrees to Radians):** Converts degrees to radians (Considered as a function call, `deg2rad(180)`). The argument is the number of degrees.
- **rad2deg (Radians to Degrees):** Converts radians to degrees (Considered as a function call, `rad2deg(3.14159)`). The argument is the number of radians.

### financial functions
- **pv (Present Value):** Calculates the present value of an investment (Considered as a function call, `pv(0.05, 5, 1000)`). The first argument is the rate of return. The second argument is the number of periods. The third argument is the future value.
- **fv (Future Value):** Calculates the future value of an investment (Considered as a function call, `fv(0.05, 5, 1000)`). The first argument is the rate of return. The second argument is the number of periods. The third argument is the present value.

### time functions
- **date (Date):** Returns the current date (Considered as a function call, `date()`). **"2006-01-02"** is the format.
- **time (Time):** Returns the current time (Considered as a function call, `time()`). **"15:04"** is the format.
- **datetime (Date Time):** Returns the current date and time (Considered as a function call, `datetime()`). **"2006-01-02 15:04"** is the format.
- **diffdate (Diff Date):** Calculates the difference between two dates (Considered as a function call, `diffdate("2022-01-01", "2022-01-02")`). The first argument is the start date. The second argument is the end date.
- **difftime (Diff Time):** Calculates the difference between two times (Considered as a function call, `difftime("15:04", "16:04")`). The first argument is the start time. The second argument is the end time.
- 

### Statistical functions
- **mean (Mean):** Calculates the mean of a list of numbers (Considered as a function call, `mean(int[1,2,3,4,5])`). The argument is the list of numbers.
- **median (Median):** Calculates the median of a list of numbers (Considered as a function call, `median(int[1,2,3,4,5])`). The argument is the list of numbers.
- **stddev (Standard Deviation):** Calculates the standard deviation of a list of numbers (Considered as a function call, `stddev(int[1,2,3,4,5])`). The argument is the list of numbers.
- **max (Max):** Calculates the maximum of a list of numbers (Considered as a function call, `max(int[1,2,3,4,5])`). The argument is the list of numbers.
- **min (Min):** Calculates the minimum of a list of numbers (Considered as a function call, `min(int[1,2,3,4,5])`). The argument is the list of numbers.

### Array functions
- **map (Map):** Applies a function to each element of a list (Considered as a function call, `map(int[1,2,3,4,5], double)`). The second argument is the function to apply to the list. The first argument is the list of numbers.
- **filter (Filter):** Filters a list based on a condition (Considered as a function call, `filter(int[1,2,3,4,5], "gt", 3)`). The second argument is the function to apply to the list. The first argument is the list of numbers.
- **reduce (Reduce):** Reduces a list of numbers to a single value (Considered as a function call, `reduce(int[1,2,3,4,5],"add", 0)`). The second argument is the function to apply to the list. The first argument is the list of numbers.
- **sum (Sum):** Sums a list of numbers (Considered as a function call, `sum(int[1,2,3,4,5])`). The argument is the list of numbers.
- **shuffle (Shuffle):** Shuffles a list of numbers (Considered as a function call, `shuffle(int[1,2,3,4,5])`). The argument is the list of numbers.
- **concat (Concat):** Concatenates two lists of numbers (Considered as a function call, `concat(int[1,2,3], int[4,5])`). The arguments are the lists of numbers.
- **reverse (Reverse):** Reverses a list of numbers (Considered as a function call, `reverse(int[1,2,3,4,5])`). The argument is the list of numbers.
- **sort (Sort):** Sorts a list of numbers (Considered as a function call, `sort(int[5,4,3,2,1])`). The argument is the list of numbers.
- **unique (Unique):** Removes duplicate numbers from a list (Considered as a function call, `unique(int[1,2,3,4,5,5,4,3,2,1])`). The argument is the list of numbers.
- **slice (Slice):** Slices a list of numbers (Considered as a function call, `slice(int[1,2,3,4,5], 1, 3)`). The first argument is the list of numbers. The second argument is the start index. The third argument is the end index.

## Encrypted Expressions
- **sha256 (SHA-256):** Calculates the SHA-256 hash of a string (Considered as a function call, `sha256("hello")`). The argument is the string to hash.
- **sha512 (SHA-512):** Calculates the SHA-512 hash of a string (Considered as a function call, `sha512("hello")`). The argument is the string to hash.


##  The Complementary BuiltInFunctions (bifs): As in code

```go
func init() {
    BuiltinFunctions = bif{}
    
    b := BuiltinFunctions
	
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
}
```

These functions are designed to handle both integers and floats, ensuring type compatibility and smooth sailing, 
below is a subset of all the methods available in the `bif` struct.

- **abs:** Unveils the absolute essence of a value, stripping away the veil of negativity.
- **date:** Translates a temporal sequence into a calendar date, anchoring fleeting moments.
- **datetime:** Merges date and time, capturing the full spectrum of a moment's presence.
- **div:** Divides two numerical values.
- **divint:** Divides with the precision of integers, discarding any fractional whispers.
- **double:** Echoes a value into twice its magnitude, reflecting its potential.
- **exp:** Elevates numbers to the power of another, scaling the heights of exponential growth.
- **filter:** Sifts through collections with a discerning eye, selecting only those that resonate.
- **len:** Measures the length, revealing the extent of data's expanse.
- **map:** Transforms each element with a spell of modification, rebirthing them anew.
- **max:** Ascends to the peak, finding the pinnacle value in a sea of numbers.
- **min:** Delves into the depths, uncovering the lowest ebb amidst numerical waves.
- **mul:** Fuses values in a dance of multiplication, celebrating their combined strength.
- **reduce:** Weaves through an array with a thread of operation, binding it into a single essence.
- **sub:** Draws apart numbers, navigating the distance between their values.
- **sqrt:** Unravels the square, bringing forth the root from the depths of its square cloister.
- **sum:** Gathers scattered numbers into a collective embrace, uniting them into one.
- **time:** Captures the flow of seconds, minutes, and hours, crystallizing them into a timestamp.

## Example: Summoning Arithmetic in Templates

```go
templ, err := template.New("example").Funcs(funcMap).Parse(`
    Result: {{ add .a .b }} | {{ div .c .d }}
    `)

    if err != nil {
        log.Fatalf("Failed to parse template: %v", err)
    }

data := map[string]any{
    "a": 7,
    "b": 5,
    "c": 10,
    "d": 2,
}

var wr bytes.Buffer
if err := templ.Execute(&wr, data); err != nil {
    log.Fatalf("Failed to execute template: %v", err)
}

fmt.Println(wr.String())
// Output: Result: 12 | 5
```

## Embark on an Expronaut Adventure

Step into the realm of Expronaut, a haven where extensive documentation, comprehensive test cases, and illustrative examples shine a light on the boundless capabilities of this enchanting tool.

Born from the essential need to decode and execute expressions within the intricate universe of PHP Blade templates in Go, Expronaut transcends its initial purpose, morphing into a vessel of exploration and discovery. It's not just a tool; it's a gateway to solving complex challenges and savoring the thrill of expression evaluation, all within the rich landscape of Go programming.

With Expronaut, embark on a voyage where logic seamlessly melds with magic, crafting a world brimming with limitless possibilities. Here, expressions aren't just evaluated—they're brought to life, setting the stage for an epic journey of coding sorcery and innovation.

Let your curiosity be your compass as you navigate through the wonders of Expronaut. Here, in the confluence of practicality and imagination, your Go projects will find their wings. The quest begins now.

## what's next
Integration with LLM's is one of the next steps, this will allow for a more dynamic and flexible way of working with expressions.
It would be nice to do something like : 
```go
	input := `ai("gpt", "is the following 42?", ( 21 + 21 ) )`
```

## License

Expronaut is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
