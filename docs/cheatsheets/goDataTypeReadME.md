# 📘 Go Data Types: Reference vs Value Types

This document summarizes the **value and reference types in Go**, along with their **advantages**, use cases, and behavior when passed to functions or compared.

---

## 🔍 Value vs Reference Types

| Category   | Types                                                                 |
|------------|-----------------------------------------------------------------------|
| Value      | `int`, `float`, `bool`, `string`, `struct`, `array`                  |
| Reference  | `map`, `slice`, `channel`, `function`, `interface`, `pointer`        |

---

## ✅ Value Types

These are **copied** when assigned or passed to functions.

| Type        | Use Case                      | Advantages                          | When to Use                         |
|-------------|-------------------------------|-------------------------------------|-------------------------------------|
| `int`, `bool`, `string` | Basic immutable values      | Fast, simple, safe copy             | Default choice                      |
| `struct`    | Grouping related data         | Strong typing, safe from mutation   | Data models, small objects          |
| `array`     | Fixed-size sequences          | Inline memory layout                | Rare, use only when size is fixed   |

---

## ✅ Reference Types

These store **pointers** to data; assignment or passing shares memory.

| Type        | Use Case                      | Advantages                          | When to Use                         |
|-------------|-------------------------------|-------------------------------------|-------------------------------------|
| `slice`     | Dynamic list/array            | Resizable, efficient indexing       | Default for sequences               |
| `map`       | Key-value lookup              | Fast hashing                        | Dictionaries, indexes               |
| `channel`   | Goroutine communication       | Safe data sharing in concurrency    | Pipelines, goroutine coordination   |
| `pointer`   | Direct memory access          | Avoid copying large structs         | Mutating shared or large data       |
| `interface` | Abstraction and polymorphism  | Flexible, clean APIs                | Interfaces, test mocking            |

---

## 🧪 Comparison Behavior

| Type       | Compared by Value | Compared by Reference | Notes                              |
|------------|-------------------|------------------------|------------------------------------|
| Pointer    | ❌                | ✅                     | `a == b` compares address          |
| Struct     | ✅                | ❌                     | Unless using `*struct`             |
| Slice      | ❌                | ✅                     | Only `== nil` allowed              |
| Map        | ❌                | ✅                     | Only `== nil` allowed              |
| Array      | ✅                | ❌                     | Fixed-length only                  |
| Interface  | ❌ (content)      | ✅ (reference)         | Can compare with `==`              |

---

## 📌 Key Takeaways

- Use **value types** when you want safe copying and immutability.
- Use **reference types** when performance or shared state is critical.
- You **can't sort or order a map**, but you can convert it to a slice of key-value pairs.
- Be cautious of shared memory with slices and maps — changes affect all references.
- Interfaces are great for **abstraction**, but be mindful of type assertions.

---
