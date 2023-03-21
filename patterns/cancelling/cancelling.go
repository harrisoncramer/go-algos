package patterns

import (
	"context"
	"errors"
	"time"
)

/* We can use the context package to time out a request after a given amount of time */
type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userId int, duration time.Duration) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*duration)
	defer cancel()

	respch := make(chan Response)
	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow(userId)
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, errors.New("Request timed out")
		case resp := <-respch:
			return resp.value, resp.err

		}
	}

}

func fetchThirdPartyStuffWhichCanBeSlow(userId int) (int, error) {
	time.Sleep(500 * time.Millisecond)

	/* We could do some sort of real work */
	return userId * 10, nil
}
