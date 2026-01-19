// Example: Basic usage of safeconv
//
// Run with: go run main.go
package main

import (
	"fmt"
	"log"
	"math"

	"github.com/matthew-reed-holden/safeconv"
)

func main() {
	fmt.Println("=== safeconv Basic Examples ===")

	// Example 1: Safe integer conversion
	fmt.Println("1. Integer Conversion (success)")
	val, err := safeconv.Int64ToUint32(1000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Int64(1000) -> Uint32(%d)\n\n", val)

	// Example 2: Detecting overflow
	fmt.Println("2. Detecting Overflow")
	bigVal := int64(math.MaxInt64)
	_, err = safeconv.Int64ToInt32(bigVal)
	if err != nil {
		fmt.Printf("   Int64(%d) -> Int32: %v\n\n", bigVal, err)
	}

	// Example 3: Detecting underflow (negative to unsigned)
	fmt.Println("3. Detecting Underflow")
	negVal := int64(-5)
	_, err = safeconv.Int64ToUint32(negVal)
	if err != nil {
		fmt.Printf("   Int64(%d) -> Uint32: %v\n\n", negVal, err)
	}

	// Example 4: Float to integer with truncation
	fmt.Println("4. Float to Integer (truncate)")
	truncated, err := safeconv.Float64TruncToInt64(3.7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Float64(3.7) truncated -> Int64(%d)\n\n", truncated)

	// Example 5: Float to integer with rounding
	fmt.Println("5. Float to Integer (round)")
	rounded, err := safeconv.Float64RoundToInt64(3.7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("   Float64(3.7) rounded -> Int64(%d)\n\n", rounded)

	// Example 6: Handling NaN
	fmt.Println("6. Handling NaN")
	_, err = safeconv.Float64TruncToInt64(math.NaN())
	if err != nil {
		fmt.Printf("   Float64(NaN) -> Int64: %v\n\n", err)
	}

	// Example 7: Must variant for initialization
	fmt.Println("7. Must Variant (for constants)")
	maxConns := safeconv.MustInt64ToUint32(100)
	fmt.Printf("   MustInt64ToUint32(100) = %d\n\n", maxConns)

	// Example 8: Widening conversion (always succeeds, no error)
	fmt.Println("8. Widening Conversion (no error possible)")
	wide := safeconv.Int32ToInt64(42)
	fmt.Printf("   Int32(42) -> Int64(%d) (no error returned)\n", wide)
}
