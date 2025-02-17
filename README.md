# Go Range Loop Panic on Empty Slice
This example demonstrates a common error in Go when using range loops on empty slices.  Attempting to iterate over an empty slice with a range loop causes a panic because the loop tries to access elements that don't exist.

## Bug
The `bug.go` file contains a function `myFunc` that iterates over a slice using a `for...range` loop. When this function is called with an empty slice, it results in a runtime panic.

## Solution
The `bugSolution.go` file provides a corrected version. It checks for an empty slice before iterating using `if len(a) > 0`. This prevents the panic by skipping the iteration when the slice is empty.

This demonstrates the importance of checking slice lengths or using other mechanisms (like a `for` loop with an index and length check) before looping in Go to prevent such panics.