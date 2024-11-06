package serve

import "net/http"

func (a *app) subscribe(w http.ResponseWriter, r *http.Request) {
	a.logs.Info("got a subscribe request")
}
