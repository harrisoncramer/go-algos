package hotdog

import (
	"net/http"
  "net/http/httptest"
	"testing"
)

func TestHotDog(t *testing.T) {

  request, err := http.NewRequest(http.MethodGet, "/", nil)
  if err != nil {
    t.Errorf("Got error during request creation: %v", err)
  }

  recorder := httptest.NewRecorder()
	dawg := Hotdog(3)
  dawg.ServeHTTP(recorder, request)

  status := recorder.Code;
  if status != 200 {
    t.Errorf("Expected 200 status, got %d", status)
  }

  response := recorder.Body.String()
  if response != "🌭" {
    t.Errorf("Expected 🌭 but got %s", response)
  }


}
