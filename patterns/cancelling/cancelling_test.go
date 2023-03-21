package patterns

import (
	"context"
	"reflect"
	"testing"
)

func TestCancelling(t *testing.T) {
	t.Run("Should make the request successfully", func(t *testing.T) {
		ctx := context.Background()
		got, err := fetchUserData(ctx, 10, 205)
		if err != nil {
			t.Errorf("Got error during request: %v", err)
		}

		want := &User{name: "sam", email: "sam@google.com"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v but wanted %v", got, want)
		}
	})

	t.Run("Should time out the request when the server is too slow", func(t *testing.T) {
		ctx := context.Background()
		got, err := fetchUserData(ctx, 10, 195)
		if err == nil {
			t.Error("Should have timed out but did not")
		}

		if got != nil {
			t.Errorf("Got %v but wanted %v", got, nil)
		}
	})
}
