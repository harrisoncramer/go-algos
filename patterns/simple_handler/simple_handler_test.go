package simplehandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimpleHandler(t *testing.T) {
	t.Run("Should test a handler", func(t *testing.T) {
		request, err := http.NewRequest("GET", "/health-check", nil)
		if err != nil {
			t.Errorf("Got error during request creation: %v", err)
		}

		/* Our recorder satisfies the http.ResponseWriter interface.
		The recorder is an implementation of http.ResponseWriter that
		records its mutations for later inspection in tests. */

		recorder := httptest.NewRecorder()

		/*
		   The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
		   It's based on the Handler interface, which looks like this:

		   type Handler interface {
		     ServeHTTP(ResponseWriter, *Request)
		   }

		   If anything satisfies this interface, it's also a handler type. Our normal function does not
		   have the ServeHTTP method, but http.HandlerFunc attaches it. We can then invoke the ServeHTTP
		   method and pass in our recorder and the request.
		*/

		handler := http.HandlerFunc(BasicHandler)
		handler.ServeHTTP(recorder, request)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("Got status %v but wanted %v", recorder.Code, http.StatusOK)
		}

		expected := `{ "alive": true }`
		body := recorder.Body.String()
		if body != expected {
			t.Errorf("Body does not match, got %s but wanted %s", body, expected)
		}

	})
}
