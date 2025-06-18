package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"effective_mobile_test_task/internal/models"
	"effective_mobile_test_task/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonHandler struct {
	db *gorm.DB
}

func NewPersonHandler(db *gorm.DB) *PersonHandler {
	return &PersonHandler{db: db}
}

// GetPersons godoc
// @Summary Get list of persons
// @Description Get list of persons with pagination and filtering
// @Tags persons
// @Accept json
// @Produce json
// @Param name query string false "Filter by name"
// @Param surname query string false "Filter by surname"
// @Param age query int false "Filter by age"
// @Param gender query string false "Filter by gender"
// @Param nationality query string false "Filter by nationality"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {array} models.Person
// @Failure 500 {object} map[string]string
// @Router /persons [get]
func (h *PersonHandler) GetPersons(c *gin.Context) {
	var persons []models.Person

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	query := h.db.Model(&models.Person{})
	if name := c.Query("name"); name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if surname := c.Query("surname"); surname != "" {
		query = query.Where("surname LIKE ?", "%"+surname+"%")
	}
	if age := c.Query("age"); age != "" {
		query = query.Where("age = ?", age)
	}
	if gender := c.Query("gender"); gender != "" {
		query = query.Where("gender = ?", gender)
	}
	if nationality := c.Query("nationality"); nationality != "" {
		query = query.Where("nationality = ?", nationality)
	}

	if err := query.Offset(offset).Limit(limit).Find(&persons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons)
}

// CreatePerson godoc
// @Summary Create a new person
// @Description Create a new person with data from external APIs
// @Tags persons
// @Accept json
// @Produce json
// @Param person body models.PersonCreateRequest true "Person data"
// @Success 201 {object} models.Person
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /persons [post]
func (h *PersonHandler) CreatePerson(c *gin.Context) {
	var request models.PersonCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	age, err := services.GetAge(request.Name)
	if err != nil {
		slog.Error("failed to get age:", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get age"})
		return
	}

	gender, err := services.GetGender(request.Name)
	if err != nil {
		slog.Error("failed to get gender", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get gender"})
		return
	}

	nationality, err := services.GetNationality(request.Name)
	if err != nil {
		slog.Error("failed to get nationality", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get nationality"})
		return
	}

	person := models.Person{
		Name:        request.Name,
		Surname:     request.Surname,
		Patronymic:  request.Patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
		CreatedAt:   time.Now(),
	}

	if err := h.db.Create(&person).Error; err != nil {
		slog.Error("failed to create person in database:", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, person)
}

// UpdatePerson godoc
// @Summary Update a person
// @Description Update person data by ID
// @Tags persons
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Param person body models.PersonUpdateRequest true "Person data"
// @Success 200 {object} models.Person
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /persons/{id} [put]
func (h *PersonHandler) UpdatePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("failed to convert id to int on update person", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var request models.PersonUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		slog.Info("failed to bind incoming request to struct", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var person models.Person
	if err := h.db.First(&person, id).Error; err != nil {
		slog.Info("failed to find person in database on updating", "err", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
		return
	}

	if request.Name != "" {
		person.Name = request.Name
	}

	if request.Surname != "" {
		person.Surname = request.Surname
	}

	if request.Patronymic != "" {
		person.Patronymic = request.Patronymic
	}

	if request.Gender != "" {
		person.Gender = request.Gender
	}

	if request.Age != 0 {
		person.Age = request.Age
	}

	if request.Nationality != "" {
		person.Nationality = request.Nationality
	}

	person.UpdatedAt = time.Now()

	if err := h.db.Save(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	slog.Debug("updated person in database", "person", person, "request", request)
	c.JSON(http.StatusOK, person)
}

// DeletePerson godoc
// @Summary Delete a person
// @Description Delete person by ID
// @Tags persons
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /persons/{id} [delete]
func (h *PersonHandler) DeletePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("failed to convert id to int on delete person", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.db.Delete(&models.Person{}, id).Error; err != nil {
		slog.Error("failed to delete person in database:", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	slog.Debug("deleted person in database", "person", id)
	c.JSON(http.StatusOK, gin.H{"message": "person deleted successfully"})
}
