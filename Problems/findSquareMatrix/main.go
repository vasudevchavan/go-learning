// Explaination & code generated from ChatGPT.
package main

import "fmt"

func maximalSquare(matrix [][]int) (int, [2]int, [2]int) {
	if len(matrix) == 0 {
		return 0, [2]int{-1, -1}, [2]int{-1, -1}
	}

	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	maxSquare := 0
	endRow, endCol := 0, 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
				}
				if dp[i][j] > maxSquare {
					maxSquare = dp[i][j]
					endRow, endCol = i, j
				}
			}
		}
	}

	if maxSquare == 0 {
		return 0, [2]int{-1, -1}, [2]int{-1, -1}
	}

	topLeft := [2]int{endRow - maxSquare + 1, endCol - maxSquare + 1}
	bottomRight := [2]int{endRow, endCol}
	return maxSquare, topLeft, bottomRight
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func main() {
	matrix := [][]int{
		{1, 0, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 0},
		{1, 1, 1, 1, 1, 1, 1},
	}

	size, topLeft, bottomRight := maximalSquare(matrix)
	fmt.Println("Largest square size:", size)
	fmt.Println("Top-left coordinate:", topLeft)
	fmt.Println("Bottom-right coordinate:", bottomRight)

	fmt.Println("\nSquare:")
	for i := topLeft[0]; i <= bottomRight[0]; i++ {
		for j := topLeft[1]; j <= bottomRight[1]; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}






// you are given a binary matrix of 0’s and 1’s.
// You need to find the largest square sub-matrix made entirely of 1’s, and return:
// The size of that square
// The coordinates (top-left and bottom-right)
// 🧠 Core Idea (Dynamic Programming)
// We use DP because each cell’s value depends on its neighbors:
// Above (dp[i-1][j])
// Left (dp[i][j-1])
// Top-left diagonal (dp[i-1][j-1])
// If we know the size of the largest square ending at each of these neighbors,
// we can easily find the largest square ending at the current cell.
// 🧱 Step 1 – Setup
// rows, cols := len(matrix), len(matrix[0])
// dp := make([][]int, rows)
// for i := range dp {
// 	dp[i] = make([]int, cols)
// }
// We create a DP matrix (dp) of the same size as matrix.
// dp[i][j] = the size of the largest square that ends at position (i, j).
// Also, we keep:
// maxSquare := 0
// endRow, endCol := 0, 0
// These will store:
// maxSquare → size of the largest square found so far
// (endRow, endCol) → bottom-right corner of that square
// 🧩 Step 2 – Fill the DP Table
// We iterate through every cell:
// for i := 0; i < rows; i++ {
// 	for j := 0; j < cols; j++ {
// 		if matrix[i][j] == 1 {
// 			if i == 0 || j == 0 {
// 				dp[i][j] = 1
// 			} else {
// 				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
// 			}

// 			if dp[i][j] > maxSquare {
// 				maxSquare = dp[i][j]
// 				endRow, endCol = i, j
// 			}
// 		}
// 	}
// }
// Let’s unpack that:
// ✅ Case 1: First row or first column
// If the current cell is in the first row or first column:
// There can’t be a square larger than 1×1 ending here.
// dp[i][j] = 1
// ✅ Case 2: Other cells
// If matrix[i][j] == 1, we check its neighbors:
// top        = dp[i-1][j]
// left       = dp[i][j-1]
// top-left   = dp[i-1][j-1]
// Then:
// dp[i][j] = 1 + min(top, left, top-left)
// Why?
// Because a square of side k can only be formed if all three neighbors can form a square of at least side k-1.
// So, we take the smallest of those three squares and add 1.
// 📘 Example Walkthrough
// For this part of the matrix:
// 1 1 1
// 1 1 1
// 1 1 1
// Let’s look at the bottom-right corner (last cell):
// top = 2 (square ending above)
// left = 2 (square ending to the left)
// top-left = 2
// → dp[i][j] = 1 + min(2, 2, 2) = 3
// So, the bottom-right cell represents a 3×3 square of 1’s ending there.
// 🧮 Step 3 – Track the Maximum
// Whenever we find a new largest square:
// if dp[i][j] > maxSquare {
// 	maxSquare = dp[i][j]
// 	endRow = i
// 	endCol = j
// }
// We update the maximum size and remember its ending coordinates.
// 🧭 Step 4 – Compute Coordinates
// After processing all cells:
// Bottom-right of the largest square = (endRow, endCol)
// Top-left = (endRow - size + 1, endCol - size + 1)
// topLeft := [2]int{endRow - maxSquare + 1, endCol - maxSquare + 1}
// bottomRight := [2]int{endRow, endCol}
// 🧩 Step 5 – Example Result
// For your input matrix:
// 1 0 1 1 1 1 1
// 0 1 0 1 1 1 0
// 1 0 1 1 1 0 1
// 1 0 1 1 1 0 0
// 1 1 1 1 1 1 1
// We get:
// Largest square size: 3
// Top-left coordinate: [2 2]
// Bottom-right coordinate: [4 4]
// And the 3×3 square is:
// 1 1 1
// 1 1 1
// 1 1 1
// ⚙️ Step 6 – Time and Space Complexity
// Type	Complexity
// Time	O(rows × cols) — each cell computed once
// Space	O(rows × cols) — can be optimized to O(cols)
// 🧠 Intuitive Summary
// Concept	Description
// DP Idea	Each cell stores the size of the largest square ending there
// Transition	dp[i][j] = 1 + min(top, left, top-left)
// Tracking	Keep updating maxSquare and coordinates
// Result	Largest square size + top-left & bottom-right coordinates
