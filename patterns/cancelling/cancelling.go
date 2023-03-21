package patterns

import (
	"context"
	"errors"
	"time"
)

type User struct {
	name  string
	email string
}

type Response struct {
	value User
	error error
}

func fetchUserData(ctx context.Context, userId int, timeout time.Duration) (*User, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*timeout)
	defer cancel()

	respch := make(chan Response)

	go func() {
		user, err := slowApiCall(userId)
		respch <- Response{
			value: user,
			error: err,
		}
	}()

	for {
		select {
		case resp := <-respch:
			return &resp.value, resp.error
		case <-ctx.Done():
			return nil, errors.New("Request timed out")
		}
	}
}

/* Some slow API call that eventually returns the user */
func slowApiCall(userId int) (User, error) {
	time.Sleep(200 * time.Millisecond)
	res := User{
		name:  "sam",
		email: "sam@google.com",
	}

	return res, nil

}
