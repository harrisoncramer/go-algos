package patterns

import (
	"context"
	"testing"
)

func TestCancelling(t *testing.T) {
	t.Run("Should make the request successfully", func(t *testing.T) {
		ctx := context.Background()
		got, err := fetchUserData(ctx, 10, 510)
		if err != nil {
			t.Errorf("Got error during request: %v", err)
		}

		want := 100
		if got != want {
			t.Errorf("Got %d but wanted %d", got, want)
		}
	})

	t.Run("Should time out the request when the server is too slow", func(t *testing.T) {
		ctx := context.Background()
		got, err := fetchUserData(ctx, 10, 400)
		if err == nil {
			t.Error("Should have timed out but did not")
		}

		want := 0
		if got != want {
			t.Errorf("Got %d but wanted %d", got, want)
		}
	})
}
