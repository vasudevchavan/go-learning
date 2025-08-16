✅ Complete Guide to Interfaces in Go
📘 What is an Interface?

An interface in Go defines a set of method signatures. Any type that implements those methods implicitly satisfies the interface — no need for an explicit declaration like implements in other languages.

🔹 Basic Syntax:
type InterfaceName interface {
	Method1()
	Method2(arg string) int
}

🧠 Key Characteristics
Feature	Description
Implicit implementation	Types implement interfaces by having the methods — no declaration needed.
Duck typing	"If it walks like a duck..." — interface is satisfied by behavior.
Zero-value of interface	nil (both the type and value inside must be nil for interface to be nil)
Use anywhere a value is expected	Can be passed into functions, returned, stored in slices, etc.
📌 Why Use Interfaces?
Use Case	Benefit
Abstraction	Hide the underlying implementation.
Polymorphism	Code that works with different types.
Dependency injection	Easily swap real implementations with mocks in tests.
Flexibility	Accept any type that implements the required behavior.
Decoupling	Keep high-level code independent of low-level implementations.
🚀 Real-World Examples
1. Standard Library Interfaces
Interface	Used In	Description
io.Reader	File input, HTTP bodies, etc.	Reads data
io.Writer	File output, HTTP responses	Writes data
fmt.Stringer	fmt.Println() and logging	Formats output as string
error	Return values	Custom error handling
http.Handler	HTTP servers	Handles HTTP requests
🔍 Example 1: Basic Interface
type Speaker interface {
	Speak() string
}

type Dog struct{}
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}
func (c Cat) Speak() string {
	return "Meow!"
}

func Talk(s Speaker) {
	fmt.Println(s.Speak())
}

🔍 Example 2: Interface with Multiple Methods
type Shape interface {
	Area() float64
	Perimeter() float64
}


Any struct that implements both methods will satisfy Shape.

🔍 Example 3: Interface for Dependency Injection
type Storage interface {
	Save(data string) error
}


This allows different storage backends (file, DB, memory) to be used interchangeably.

🎯 Types of Interfaces
1. Single-method interfaces

Very common in Go

Examples: io.Reader, io.Writer, fmt.Stringer, http.Handler

2. Multi-method interfaces

For related behaviors grouped under one contract

Example: Shape with Area() and Perimeter()

3. Empty Interface (interface{})

Can hold any type — before Go 1.18 (generics), this was the way to handle unknown types.

Common in things like JSON unmarshalling, dynamic values

var x interface{}
x = 5         // int
x = "hello"   // string
x = true      // bool


⚠️ Avoid overusing interface{} — it removes type safety.

🧪 Interface Internals (Advanced)

An interface is internally represented as a pair:

A type (concrete type implementing it)

A value (the actual data)

This is why an interface can be non-nil but still contain nil:

var r io.Reader = (*os.File)(nil)
fmt.Println(r == nil) // false


Use this with care in nil checks.

⚠️ Common Interface Mistakes
Mistake	Why it's bad
Defining interfaces too early	Premature abstraction; wait until you need it.
Large interfaces	Hard to satisfy and test; Go prefers small.
Using interface{} everywhere	Sacrifices type safety
Comparing interfaces incorrectly	Interface may not be nil even if value is
✅ Best Practices
Practice	Benefit
Keep interfaces small	Easier to mock, implement, test
Define interfaces at the consumer	Better decoupling and clearer requirements
Use existing standard interfaces	Avoid reinventing (io.Writer, error, etc.)
Use empty interface only when needed	Maintain type safety
Check nil properly	Use == nil only if both type and value are nil
✅ Interface Comparison Table
Concept	Value Type	Interface
Implementation	Must match struct	Any matching methods
Comparison	By value	By dynamic type + value
Memory	Inline	Pointer-based
Use in slices	Direct	Mixed-type slices
Method calls	Static	Dynamic (via vtable)
🧪 Unit Testing with Interfaces

Interfaces are ideal for writing mockable code.

Example:
type Notifier interface {
	Send(msg string) error
}

// Real email sender
type EmailSender struct{}
func (e EmailSender) Send(msg string) error {
	fmt.Println("Sending email:", msg)
	return nil
}

// Mock for testing
type MockSender struct {
	Messages []string
}
func (m *MockSender) Send(msg string) error {
	m.Messages = append(m.Messages, msg)
	return nil
}

📚 Learn More

Go Blog: Interfaces

Effective Go: Interfaces

Standard library usage

🧩 Summary
Feature	Details
Implements	Implicitly by methods
Use Cases	Abstraction, decoupling, polymorphism, mocking
Interface{}	Special case for any type (use cautiously)
Common Patterns	io.Reader, fmt.Stringer, http.Handler, error
Best Practice	Keep small, define at point of use, test with mocks