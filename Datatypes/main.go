package main // Every Go program must start with a package declaration.
// "main" means this is an executable program (not a library).

import (
	"fmt" // fmt provides functions for formatted I/O (printing, scanning, etc.)
)

// ========== 1. BASIC TYPES ==========

// Example: Integers (whole numbers, both signed and unsigned)
func demoIntegers() {
	var x int = 42   // int stores whole numbers, size depends on the system (32 or 64 bits)
	var y uint = 100 // uint is an unsigned int (only positive numbers)
	fmt.Println("Integers => x:", x, "y:", y)
}

// Example: Floating-point numbers (decimals)
func demoFloats() {
	var pi float32 = 3.14    // float32 gives ~7 digits of precision
	var e float64 = 2.71828  // float64 gives ~15 digits of precision (recommended)
	fmt.Println("Floats => pi:", pi, "e:", e)
}

// Example: Boolean (true/false values)
func demoBooleans() {
	var isActive bool = true // boolean type can only be true or false
	fmt.Println("Boolean => isActive:", isActive)
}

// Example: Strings (text values, UTF-8 encoded)
func demoStrings() {
	name := "GoLang"                       // shorthand declaration with :=
	fmt.Println("String => name:", name, "length:", len(name)) // len() returns number of bytes
}

// ========== 2. COMPOSITE TYPES ==========

// Example: Arrays (fixed-size list of same type)
func demoArrays() {
	var nums [3]int = [3]int{10, 20, 30} // array of size 3, containing ints
	fmt.Println("Array =>", nums)
}

// Example: Slices (dynamic-size list, more common than arrays)
func demoSlices() {
	slice := []string{"apple", "banana"} // slice can grow or shrink
	slice = append(slice, "cherry")      // append adds new elements
	fmt.Println("Slice =>", slice)
}

// Example: Maps (key-value store, like dictionary in Python)
func demoMaps() {
	m := map[string]int{"Alice": 25, "Bob": 30} // keys are strings, values are ints
	fmt.Println("Map => Alice's age:", m["Alice"])
}

// Example: Structs (custom data type, like classes without methods)
type User struct {
	Name string // field "Name" of type string
	Age  int    // field "Age" of type int
}

func demoStructs() {
	u := User{Name: "Jaideep", Age: 29} // initialize struct with field values
	fmt.Println("Struct =>", u)
}

// ========== 3. REFERENCE TYPES ==========

// Example: Pointers (store memory address of a variable)
func demoPointers() {
	x := 10    // normal integer variable
	p := &x    // &x means "address of x"
	fmt.Println("Pointer => x:", x, "p:", p, "value at p:", *p) // *p gives value at address
}

// Example: Functions as values (you can assign functions to variables)
func add(a, b int) int {
	return a + b // function that returns sum of a and b
}

func demoFunctions() {
	f := add                     // assign function "add" to variable f
	fmt.Println("Function => sum:", f(3, 4)) // call f like a normal function
}

// Example: Interfaces (define behavior, not data)
type Shape interface {
	Area() float64 // any type with Area() method is a Shape
}

// Concrete type implementing Shape interface
type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func demoInterfaces() {
	var s Shape = Square{4} // assign a Square to a Shape variable
	fmt.Println("Interface => square area:", s.Area())
}

// ========== 4. SPECIAL TYPES ==========

// Example: Byte and Rune
func demoByteRune() {
	var b byte = 'A'   // byte is alias for uint8, usually represents ASCII characters
	var r rune = '好' // rune is alias for int32, used for Unicode characters
	fmt.Println("Byte & Rune => byte:", b, "rune:", r)
}

// Example: Constants
func demoConstants() {
	const Pi = 3.14159 // constant value (cannot be changed later)
	fmt.Println("Constant => Pi:", Pi)
}

// Example: Nil (zero value for pointers, maps, slices, interfaces, etc.)
func demoNil() {
	var p *int // declare a pointer to int, but not initialized
	if p == nil {
		fmt.Println("Nil => pointer is nil")
	}
}

// ========== MAIN FUNCTION ==========
// Execution starts here
func main() {
	// Uncomment the functions you want to run one by one:

	 demoIntegers()
	 demoFloats()
	 demoBooleans()
	 demoStrings()

	 demoArrays()
	 demoSlices()
	 demoMaps()
	 demoStructs()

	 demoPointers()
	 demoFunctions()
	 demoInterfaces()

	 demoByteRune()
	 demoConstants()
	 demoNil()
}
