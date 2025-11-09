package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{svc: svc}
}

func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/catfact", s.handleGetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}

func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		WriteJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, fact)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// func ParseJSON(r *http.Request, payload any) error {
// 	if r.Body == nil {
// 		return fmt.Errorf("request body is empty")
// 	}
// 	return json.NewDecoder(r.Body).Decode(payload)
// 	// Note: We don't close r.Body here because the http package
// 	// automatically handles that after the handler returns.
// }

// WriteJSON is used when your server gets data (like from a database)
// and needs to send it back to the client, usually in response to a
// GET request.

// func WriteError(w http.ResponseWriter, status int, err error) {
// 	WriteJSON(w, status, map[string]string{"error": err.Error()})
// }
