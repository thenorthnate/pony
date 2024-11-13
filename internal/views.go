package pony

import (
	"io"
	"net/http"
)

const (
	ContentTypeHeader = "Content-Type"
)

func (a *app) createViews() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/subscribe/{topic}", a.subscribe)
	mux.HandleFunc("/v1/publish/{topic}", a.publish)
	return mux
}

func (a *app) subscribe(w http.ResponseWriter, r *http.Request) {
	a.logs.Info("got a subscribe request")
}

func (a *app) publish(w http.ResponseWriter, r *http.Request) {
	topic := r.PathValue("topic")
	if topic == "" {
		http.Error(w, "you must specify which topic to publish to", http.StatusBadRequest)
		return
	}
	contentType := r.Header.Get(ContentTypeHeader)
	if contentType == "" {
		http.Error(w, "Content-Type header must be specified", http.StatusBadRequest)
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	values := r.URL.Query()
	if err := a.db.InsertMessage(r.Context(), topic, contentType, data, values); err != nil {
		http.Error(w, "failed to handle incoming message", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
