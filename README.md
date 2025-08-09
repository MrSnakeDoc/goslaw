# Goslaw

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Goslaw is a Go library that brings JavaScript-like Array methods to Go slices, making collection manipulation more expressive and functional.

## 🚀 Installation

```bash
go get -u github.com/MrSnakeDoc/goslaw
```

## 📋 Prerequisites

- Go 1.24+ (required for generics and use of new built-in functions like `min`, `max`, etc.)

## 📖 Usage

```go
package main

import (
	"fmt"
	
	"github.com/MrSnakeDoc/goslaw"
)

func main() {
	// Example with Map
	numbers := []int{1, 2, 3, 4, 5}
	doubled := goslaw.Map(numbers, func(x int) int {
		return x * 2
	})
	fmt.Println(doubled) // Output: [2 4 6 8 10]
	
	// Example with Filter
	evens := goslaw.Filter(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println(evens) // Output: [2 4]
	
	// Example with Reduce
	sum := goslaw.Reduce(numbers, 0, func(acc, x int) int {
		return acc + x
	})
	fmt.Println(sum) // Output: 15
	
	// Example with ForEach
	goslaw.ForEach(numbers, func(x int) {
		fmt.Printf("%d ", x)
	}) // Output: 1 2 3 4 5
	
	// Combining operations
	result := goslaw.Reduce(
		goslaw.Filter(
			goslaw.Map(numbers, func(x int) int { return x * x }),
			func(x int) bool { return x > 10 },
		),
		0,
		func(acc, x int) int { return acc + x },
	)
	fmt.Println("\nSum of squares greater than 10:", result) // Output: 41 (16 + 25)
}
```

## 🛠️ Available Functions

This library provides the following functions to work with slices:

### 🔄 Transformation Functions

| Function | Description | JS Equivalent |
|----------|-------------|---------------|
| `Map` | Transforms each element in a slice using the provided function | `Array.map()` |
| `FlatMap` | Maps each element using a function, then flattens the result | `Array.flatMap()` |
| `Flat` | Flattens a nested slice structure | `Array.flat()` |
| `Reverse` | Returns a new slice with elements in reversed order | `Array.reverse()` |
| `ToReversed` | Returns a new slice with elements in reversed order | `Array.toReversed()` |
| `Sort` | Sorts the elements in a slice (modifies the original slice) | `Array.sort()` |
| `With` | Returns a new slice with a specified element replaced at a given index | `Array.with()` |
| `Concat` | Concatenates multiple slices into one | `Array.concat()` |
| `Fill` | Fills all elements in a slice with a static value | `Array.fill()` |
| `CopyWithin` | Copies part of a slice to another position in-place | `Array.copyWithin()` |
| `Splice` | Removes and/or adds elements at a given index | `Array.splice()` |


### 🔄 Iteration Functions

| Function | Description | JS Equivalent |
|----------|-------------|---------------|
| `ForEach` | Executes a provided function once for each slice element | `Array.forEach()` |
| `Reduce` | Reduces a slice to a single value by applying a function against an accumulator | `Array.reduce()` |
| `ReduceRight` | Like Reduce but processes the slice from right to left | `Array.reduceRight()` |

### 🔍 Search Functions

| Function | Description | JS Equivalent |
|----------|-------------|---------------|
| `Find` | Returns the first element that satisfies the provided testing function | `Array.find()` |
| `FindIndex` | Returns the index of the first element that satisfies the provided testing function | `Array.findIndex()` |
| `Filter` | Creates a new slice with all elements that pass the test function | `Array.filter()` |
| `Every` | Tests whether all elements in the slice pass the provided function | `Array.every()` |
| `Some` | Tests whether at least one element in the slice passes the provided function | `Array.some()` |
| `Includes` | Determines whether a slice includes a certain value | `Array.includes()` |
| `IndexOf` | Returns the first index at which a given element can be found | `Array.indexOf()` |
| `EveryIndex` | Returns the index of the first element that satisfies the provided testing function | `Array.everyIndex()` |

### ✏️ Modification Functions

| Function | Description | JS Equivalent |
|----------|-------------|---------------|
| `Pop` | Removes the last element from a slice and returns that element | `Array.pop()` |
| `Shift` | Removes the first element from a slice and returns that element | `Array.shift()` |
| `Unshift` | Adds one or more elements to the beginning of a slice | `Array.unshift()` |
| `Slice` | Returns a shallow copy of a portion of a slice | `Array.slice()` |

### 🗂️ Object Functions

| Function | Description | JS Equivalent |
|----------|-------------|---------------|
| `Keys` | Returns the keys of a map | `Object.keys()` |
| `Values` | Returns the values of a map | `Object.values()` |
| `Entries` | Returns key-value pairs from a map as a slice of structs | `Object.entries()` |

### 📝 String Functions

| Function | Description | JS Equivalent |
|----------|-------------|---------------|
| `Join` | Joins all elements of a slice into a string | `Array.join()` |
| `ToString` | Converts a slice to its string representation | `Array.toString()` |

## 🤔 Why Goslaw?

Go is a powerful but minimalist language. Its standard library offers basic functionality for working with slices, but often lacks the higher-order operations found in other languages like JavaScript. Goslaw  bridges this gap by providing a familiar and functional API for slice manipulation.

## 🔧 Performance

Goslaw is designed to be efficient while offering an elegant API. All functions use Go generics, which means there's no runtime cost associated with type assertions or reflections.

## 🤝 Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## 📄 License

This project is licensed under the MIT License.
