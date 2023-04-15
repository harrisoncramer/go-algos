package error_wrapping

import (
	"errors"
	"fmt"
)

func main() {
	/* Create an original error */
	originalErr := errors.New("Original error")

	/* Wrap the original error with additional context */
	wrappedErr := fmt.Errorf("Wrapped error: %w", originalErr)

	/* Handle the wrapped error. The errors.Is function will call errors.Unwrap()
	repeatedly on the chain until it finds that the original error is present through a strict
	equality check. If it reaches the final error without a match, it'll return false. */
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("Got original error:", originalErr)
	} else {
		fmt.Println("Got wrapped error:", wrappedErr)
	}
}
