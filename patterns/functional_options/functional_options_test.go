package functional_options

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestFunctionalOptions(t *testing.T) {
	t.Run("Example implementation of the functions", func(t *testing.T) {
		t.SkipNow()
		rand.Seed(time.Now().Unix()) /* Used for helper func */

		err, server, port := NewServer(WithPort(0), WithName("Harry"))

		if err != nil {
			log.Fatal(err)
		}

		var e chan (error)

		go func() {
			err = server.ListenAndServe()
			if err != nil {
				e <- err
			}
		}()

		fmt.Printf("Server running on port %d", port)
		<-e
	})
}
