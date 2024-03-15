# Expronaut: Magic in Expression Evaluation

Welcome to Expronaut, a fun and dynamic Go package designed for parsing and evaluating expressions with a touch of magic. 
Whether you're crafting conditions for template rendering or just dabbling in the alchemy of logic and math, Expronaut is your companion on this adventurous journey through the realms of syntax and semantics.

## Features

- **Dynamic Expression Parsing:** Dive into expressions with variables, nested properties, and an assortment of operators. From simple arithmetic to complex boolean logic, Expronaut ~~understands it all~~, tries to understand.
- **Flexible Variable Context:** Whether you're working with flat landscapes or exploring the depths of nested objects, Expronaut navigates through your data with ease, bringing context to your expressions.
- **Customizable Evaluation:** Tailor the evaluation context to suit your adventure. Pass in variables and watch as Expronaut conjures up the results you seek.

## Getting Started

Embark on your journey with Expronaut by incorporating it into your Go projects. Here's how to get started:

## Installation

Ensure you have Go installed on your system, then fetch the Expronaut package:

```sh
go get github.com/donseba/expronaut
```

## Usage Examples
Evaluating Expressions in Templates

```go
package main

import (
    "bytes"
    "text/template"
    "testing"

    "github.com/donseba/expronaut"
)

func TestExp(t *testing.T) {
    content := `{{ if exp "foo == 5" "foo" .foo }} OK {{ else }} NOT OK {{ end }}`

    templ, err := template.New("test").Funcs(map[string]any{
        "exp": expronaut.Exp,
    }).Parse(content)
    if err != nil {
        t.Fatal(err)
    }

    data := map[string]any{
        "foo": 5,
    }

    var wr bytes.Buffer
    err = templ.ExecuteTemplate(&wr, "test", data)
    if err != nil {
        t.Fatal(err)
    }

    if wr.String() != " OK " {
        t.Fatalf("expected OK, got %s", wr.String())
    }
}
```
## Converting expressions to Go Template Strings

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
- 
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
- **sqrt (Square Root):** Calculates the square root of a number (Considered as a function call, sqrt(x)).
- **date (Date):** Returns the current date (Considered as a function call, date()). **"2006-01-02"** is the format.
- **time (Time):** Returns the current time (Considered as a function call, time()). **"15:04"** is the format.
- **datetime (Date Time):** Returns the current date and time (Considered as a function call, datetime()). **"2006-01-02 15:04"** is the format.

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

##  The Complementary FuncMap: A Swashbuckling Toolkit

To enhance your Go template experience, Expronaut comes with a complementary funcMap. This map is a collection of arithmetic functions tailored to work within your templates, allowing you to perform calculations directly:

```go
var funcMap = template.FuncMap{
    "add": add,
    "sub": sub,
    "div": div,
    "mul": mul,
    "mod": mod,
    "sqrt": sqrt,
}
```

These functions are designed to handle both integers and floats, ensuring type compatibility and smooth sailing:

- **add:** Summons two values, casting them into harmony.
- **sub:** Finds the difference between two numbers, navigating through type tides.
- **div:** Divides one value by another, steering clear of the rocks of division by zero.
- **mul:** Multiplies values, amplifying their essence.
- **mod:** Seeks the remainder, exploring the cyclical nature of numbers.

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

## Dive Deeper

Dive into the world of Expronaut, where documentation, test cases, and examples await to guide you through the capabilities of this powerful tool.

Designed out of necessity for a tool capable of interpreting expressions within a larger package aimed at parsing PHP Blade templates and executing them in Go, Expronaut has evolved beyond its practical origins. Whether you're tackling a specific problem or simply indulging in the joy of exploring expression evaluation, Expronaut brings a unique blend of utility and enchantment to your Go projects.

Embark on a journey with Expronaut, a place where logic intertwines with a touch of magic, unlocking endless possibilities for your expressions. Let the adventure begin!