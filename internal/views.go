package pony

import "net/http"

func (a *app) createViews() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/subscribe", a.subscribe)
	return mux
}

func (a *app) subscribe(w http.ResponseWriter, r *http.Request) {
	a.logs.Info("got a subscribe request")
}
