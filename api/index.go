package handler

import (
	"net/http"
	"sync"
	"wirehound/internal/engine"
	"wirehound/internal/handlers"
)

var (
	eng  *engine.Engine
	once sync.Once
)

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		eng = engine.New()
	})

	mux := http.NewServeMux()
	handlers.RegisterRoutes(mux, eng)
	mux.ServeHTTP(w, r)
}
