package functional_options

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

/* When designing an API, it's helpful to have patterns for adding optional configuration. This is achieved via the functional options pattern, which lets us pass a series of functions to a creator function, which mutates the underlying default option object. This lets consumers optionally configure the server and use the defaults easily */

type options struct {
	port *int
	name *string
}

type Option func(o *options) error

/* This is a setter function. It creates a closure with the first argument, and then mutates the options struct */
func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

func WithName(name string) Option {
	return func(options *options) error {
		if name == "" {
			return errors.New("Name cannot be blank")
		}

		options.name = &name
		fmt.Printf("Name is %s", *options.name)
		return nil
	}
}

/* The creator function calls the returned function of each setter, applying the closure arguments to the struct */
func NewServer(setters ...Option) (error, *http.Server, int) {
	var options options
	for _, setter := range setters {
		err := setter(&options)
		if err != nil {
			return err, nil, 0
		}
	}

	/* The opts object has been mutated by each function */
	defaultPort := 13902
	if options.port == nil {
		options.port = &defaultPort
	} else if *options.port == 0 {
		options.port = randomPort()
	}

	defaultName := "world"
	if options.name == nil {
		options.name = &defaultName
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Hello, %s", *options.name)))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *options.port),
		Handler: mux,
	}

	return nil, server, *options.port
}

/* Helper */
func randomPort() *int {
	min := 1025
	result := rand.Intn(65535-min) + min
	return &result
}
