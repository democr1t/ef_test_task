package server

// @title Effective Mobile Test Task API
// @version 1.0
// @description API for managing people information

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /people
import (
	"context"
	_ "effective_mobile_test_task/docs"
	"effective_mobile_test_task/internal/app/metrics"
	"effective_mobile_test_task/internal/app/services"
	"effective_mobile_test_task/internal/domain/cache"
	"encoding/json"
	"errors"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
	"strconv"
	"time"
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

	mux.HandleFunc(
		"GET /swagger/",

		httpSwagger.Handler(
			httpSwagger.URL("/swagger/doc.json"), // URL к swagger.json
		),
	)

	// Добавляем endpoint для метрик
	mux.HandleFunc("GET /metrics", s.metrics.MetricsHandler().ServeHTTP)

	slog.Info("Starting server", "addr", addr)
	return http.ListenAndServe(":"+addr, mux)
}

// GetPersonByID godoc
// @Summary Get person by ID
// @Description Get detailed information about a person by their ID
// @Tags people
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} entities.Person "Successfully retrieved person"
// @Failure 400 {string} string "Invalid ID format"
// @Failure 404 {string} string "Person not found"
// @Failure 500 {string} string "Internal server error"
// @Router /people/{id} [get]
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
