package main

import "fmt"

func main() {
	// ===========================
	// 1️⃣ BASIC VARIABLE DECLARATIONS
	// ===========================

	var name string = "Jaideep Singh Rathore"
	// Explicit declaration: name, type, and value provided.

	var name2 = "Aditi Jeena"
	// Type inference: Go automatically detects the type as 'string'.

	name3 := "Chunu Munnu Chinni"
	// Short declaration: := can only be used inside a function (not at package level).

	var name4 string
	name4 = "Test"
	// Declare first, assign later (must specify type at declaration).

	fmt.Println("The input string 1 is:", name)
	fmt.Println("The input string 2 is:", name2)
	fmt.Println("The input string 3 is:", name3)
	fmt.Println("The input string 4 is:", name4)

	// ===========================
	// 2️⃣ BASIC TYPES AND ZERO VALUES
	// ===========================

	var age int = 40
	var isAdult bool = true
	var perSecIncome float32 = 1.334

	fmt.Println("The input int is:", age)
	fmt.Println("The input bool is:", isAdult)
	fmt.Println("The input float is:", perSecIncome)

	// Zero value concept — variables declared without initialization
	// get a default "zero value" depending on their type.
	var defaultInt int
	var defaultBool bool
	var defaultString string
	var defaultFloat float64

	fmt.Println("Default int value:", defaultInt)       // 0
	fmt.Println("Default bool value:", defaultBool)     // false
	fmt.Println("Default string value:", defaultString) // ""
	fmt.Println("Default float value:", defaultFloat)   // 0

	// ===========================
	// 3️⃣ CONSTANTS
	// ===========================
	const Pi float64 = 3.14159
	const Greeting = "Hello, Go Learner!"

	fmt.Println("Constant Pi value:", Pi)
	fmt.Println("Constant Greeting value:", Greeting)

	// Constants cannot be changed once defined:
	// Pi = 3.14 // ❌ This would cause a compilation error

	// ===========================
	// 4️⃣ MULTIPLE VARIABLE DECLARATION
	// ===========================
	var x, y, z int = 10, 20, 30
	a, b, c := "Apple", "Banana", "Cherry"

	fmt.Println("Multiple int variables:", x, y, z)
	fmt.Println("Multiple string variables:", a, b, c)

	// ===========================
	// 5️⃣ TYPE CONVERSION
	// ===========================
	var smallNumber int = 10
	var bigNumber float64 = float64(smallNumber) * 1.5
	fmt.Println("Converted int to float64:", bigNumber)

	// Converting numeric to string using fmt.Sprintf
	var incomeString string = fmt.Sprintf("%.2f", perSecIncome)
	fmt.Println("Formatted income as string:", incomeString)

	// ===========================
	// 6️⃣ VARIABLE SHADOWING
	// ===========================
	// If you use := with an already declared variable name inside a new scope,
	// it creates a *new* variable (shadows the old one).
	number := 5
	fmt.Println("Outer number:", number)
	{
		number := 10 // Shadows the outer 'number'
		fmt.Println("Inner (shadowed) number:", number)
	}
	fmt.Println("Outer number after shadowing:", number)

	// ===========================
	// 7️⃣ SHORT SUMMARY
	// ===========================
	fmt.Println("\nSummary:")
	fmt.Println("- Use 'var' for flexible declarations and 'const' for constants.")
	fmt.Println("- Use ':=' for short declarations inside functions.")
	fmt.Println("- Uninitialized vars get zero values.")
	fmt.Println("- Type conversion is explicit in Go.")
	fmt.Println("- Be cautious of variable shadowing inside nested blocks.")
}
