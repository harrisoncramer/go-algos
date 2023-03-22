package hotdog

import "net/http"

/*
Since this integer-backed type implements the handler interface it's allowed,
even though it's a integer. We technically don't even have to use the 
http.ResponseWriter or http.Request arguments, so long as the interface
is satisfied.
*/
type Hotdog int

func (h Hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(200)
	w.Write([]byte("🌭"))
}
