// Example: Error handling patterns with safeconv
//
// Run with: go run main.go
package main

import (
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/heisenbergs-uncertainty/safeconv"
)

func main() {
	fmt.Println("=== safeconv Error Handling Examples ===\n")

	// Example 1: Using errors.Is for quick checks
	fmt.Println("1. Using errors.Is()")
	demonstrateErrorsIs()

	// Example 2: Using errors.As for detailed info
	fmt.Println("\n2. Using errors.As()")
	demonstrateErrorsAs()

	// Example 3: Switch on error types
	fmt.Println("\n3. Switch on Error Types")
	demonstrateErrorSwitch()

	// Example 4: Real-world usage pattern
	fmt.Println("\n4. Real-World Pattern")
	demonstrateRealWorld()
}

func demonstrateErrorsIs() {
	testCases := []int64{-5, 1000, math.MaxInt64}

	for _, x := range testCases {
		_, err := safeconv.Int64ToUint32(x)
		switch {
		case err == nil:
			fmt.Printf("   Int64(%d) -> success\n", x)
		case errors.Is(err, safeconv.ErrUnderflow):
			fmt.Printf("   Int64(%d) -> ErrUnderflow (negative value)\n", x)
		case errors.Is(err, safeconv.ErrOverflow):
			fmt.Printf("   Int64(%d) -> ErrOverflow (too large)\n", x)
		}
	}
}

func demonstrateErrorsAs() {
	_, err := safeconv.Int64ToUint32(-42)
	if err != nil {
		var convErr *safeconv.ConversionError
		if errors.As(err, &convErr) {
			fmt.Printf("   From:  %s\n", convErr.From)
			fmt.Printf("   To:    %s\n", convErr.To)
			fmt.Printf("   Value: %v\n", convErr.Value)
			fmt.Printf("   Err:   %v\n", convErr.Err)
		}
	}
}

func demonstrateErrorSwitch() {
	testValues := []float64{3.14, math.NaN(), math.Inf(1), 1e20}

	for _, x := range testValues {
		_, err := safeconv.Float64TruncToInt32(x)
		switch {
		case err == nil:
			fmt.Printf("   Float64(%v) -> success\n", x)
		case errors.Is(err, safeconv.ErrNaN):
			fmt.Printf("   Float64(%v) -> ErrNaN\n", x)
		case errors.Is(err, safeconv.ErrInfinity):
			fmt.Printf("   Float64(%v) -> ErrInfinity\n", x)
		case errors.Is(err, safeconv.ErrOverflow):
			fmt.Printf("   Float64(%v) -> ErrOverflow\n", x)
		}
	}
}

func demonstrateRealWorld() {
	// Simulate processing user input
	userID := int64(12345)
	amount := 99.99

	// Convert user ID to uint32 for database
	dbUserID, err := safeconv.Int64ToUint32(userID)
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		return
	}

	// Convert amount to cents (integer)
	cents, err := safeconv.Float64RoundToInt64(amount * 100)
	if err != nil {
		log.Printf("Invalid amount: %v", err)
		return
	}

	fmt.Printf("   User ID: %d (uint32)\n", dbUserID)
	fmt.Printf("   Amount:  %d cents (int64)\n", cents)
}
