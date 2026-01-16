// utils/helpers.go
package internal

import (
	"fmt"
	"strings"
	"time"
)

// StringHelper provides utility functions for string manipulation
type StringHelper struct {
	defaultSeparator string
}

// NewStringHelper creates a new StringHelper instance
func NewStringHelper(separator string) *StringHelper {
	return &StringHelper{
		defaultSeparator: separator,
	}
}

// Reverse reverses a string
func (sh *StringHelper) Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountWords counts words in a string
func (sh *StringHelper) CountWords(s string) int {
	return len(strings.Fields(s))
}

// Capitalize capitalizes first letter of each word
func (sh *StringHelper) Capitalize(s string) string {
	return strings.Title(strings.ToLower(s))
}

// TimeHelper provides utility functions for time operations
type TimeHelper struct{}

// FormatDuration formats a duration into human readable string
func (th *TimeHelper) FormatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
}

// IsWeekend checks if given time is weekend
func (th *TimeHelper) IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// MathHelper provides mathematical utility functions
type MathHelper struct{}

// Fibonacci generates fibonacci sequence up to n terms
func (mh *MathHelper) Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	fib := make([]int, n)
	fib[0], fib[1] = 0, 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

// IsPrime checks if number is prime
func (mh *MathHelper) IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

// Sum returns sum of integers
func (mh *MathHelper) Sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
