package server

import (
	"context"
	"effective_mobile_test_task/internal/app/metrics"
	"effective_mobile_test_task/internal/app/services"
	"effective_mobile_test_task/internal/domain/cache"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"log/slog"
)

var ErrNotFound = errors.New("not found")

type Server struct {
	service *services.PersonService
	metrics *metrics.MetricsService
	cache   cache.Cache
}

func NewServer(service *services.PersonService, metrics *metrics.MetricsService, cache cache.Cache) *Server {
	return &Server{
		service: service,
		metrics: metrics,
		cache:   cache,
	}
}

func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()

	// Регистрируем обработчики с инструментированием метрик
	mux.HandleFunc(
		"GET /people/{id}",
		s.metrics.InstrumentHandler("GET", "/people/{id}", s.handleGetPersonByID),
	)

	// Добавляем endpoint для метрик
	mux.HandleFunc("GET /metrics", s.metrics.MetricsHandler().ServeHTTP)

	slog.Info("Starting server", "addr", addr)
	return http.ListenAndServe(":"+addr, mux)
}

func (s *Server) handleGetPersonByID(w http.ResponseWriter, r *http.Request) {

	if val, err := s.cache.Get(""); err != nil {
		//TODO implement get value from cache and return to users
		fmt.Println(val)
	}

	slog.Debug("start processing get person by ID")
	idStr := r.PathValue("id")
	slog.Debug("get person by id", "id", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	person, err := s.service.GetByID(ctx, id)

	slog.Debug("person", slog.Any("persons", person))

	if err != nil {
		switch {
		case errors.Is(err, ErrNotFound):
			http.Error(w, "Person not found", http.StatusNotFound)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	//TODO implement set in cache before response to user mb async
	err = s.cache.Set("", "")

	if err != nil {

	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(person); err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
