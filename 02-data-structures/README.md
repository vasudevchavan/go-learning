# 🟢 Go Data Structures & Data Types

Master Go's built-in data types, composite types, and common data structures.

---

## 📚 Topics Covered
- **basics/**: Strings and basic operations  
- **arrays-slices/**: Arrays and slices manipulation  
- **maps/**: Hash maps and key-value operations  
- **structs/**: Custom data types and methods  
- **interfaces/**: Polymorphism and abstraction  

**Key Concepts:**  
- Value vs Reference types  
- Memory management  
- Type safety  

---

## 1️⃣ Primitive Data Types

| Type | Description | Example |
|------|------------|---------|
| `int` | Integer (platform-dependent size) | `var a int = 10` |
| `int8, int16, int32, int64` | Fixed-size signed integers | `var b int32 = 1000` |
| `uint, uint8, uint16, uint32, uint64` | Unsigned integers | `var c uint8 = 255` |
| `float32, float64` | Floating-point numbers | `var f float64 = 3.14` |
| `complex64, complex128` | Complex numbers | `var z complex128 = 2 + 3i` |
| `bool` | Boolean | `var flag bool = true` |
| `string` | UTF-8 string | `var name string = "Go"` |
| `byte` | Alias for `uint8` | `var b byte = 'a'` |
| `rune` | Alias for `int32` (Unicode code point) | `var r rune = '世'` |

**Example:**
```go
var age int = 25
var pi float64 = 3.1415
var name string = "Alice"
var isActive bool = true
```


## 2️⃣ Composite Data Types

### a) Array – Fixed size
```go
	
	var arr [5]int
	arr[0] = 10
	fmt.Println(arr)  // [10 0 0 0 0]
	Access: arr[index]
	Size fixed at compile-time
```
###	b) Slice – Dynamic array

```go
	slice := []int{1,2,3}
	slice = append(slice, 4)
	fmt.Println(slice)  // [1 2 3 4]
	Can grow/shrink dynamically
	Backed by arrays
```
###	c) Map – Key-value store
	m := make(map[string]int)
	m["Alice"] = 25
	fmt.Println(m["Alice"])  // 25
	Lookup: m[key] → O(1) average
	Can act as a Set: map[type]bool
###	d) Struct – Custom data type
	type Person struct {
	    Name string
	    Age  int
	}
	p := Person{Name: "Alice", Age: 25}
	fmt.Println(p.Name, p.Age)
	Can implement linked lists, trees, graphs
###	e) Pointer – Memory reference
	a := 10
	p := &a
	fmt.Println(*p)  // 10
###	f) Interface – Abstraction
	type Shape interface {
	    Area() float64
	}
	Defines a set of methods
	Allows polymorphism



## 3️⃣ Common Data Structures in Go
|Structure	|Implementation|
|-----------|--------------|
|Stack / Queue	|Use slice (append + pop)|
|Linked List	|Use struct + pointer|
|Heap / Priority |Queue	container/heap package|
|Set	|map[type]bool|


### Set Example:
	set := make(map[int]bool)
	set[1] = true
	set[2] = true
	if set[1] {
		fmt.Println("1 exists")
	}

## 4️⃣ Type Conversion
	Go does not do implicit type conversion.
	var a int = 10
	var b float64 = float64(a) + 3.5
	fmt.Println(b)  // 13.5

## 5️⃣ Complexity Notes
|Structure |Access|	Insert|	Delete|	Notes|
|----------|-------|-------|-------|------|
|Array	|O(1)|	|O(n)|	O(n)|	Fixed size|
|Slice	|O(1)|	|O(1)| amortized	O(n)	|Dynamic array|
|Map	|O(1)| |avg|	O(1)	|O(1)	|Hash map|
|Struct	|O(1)| |access to fields|	N/A	N/A	Composite type|
|Pointer	|O(1)|	|N/A	|N/A	|Memory reference|



