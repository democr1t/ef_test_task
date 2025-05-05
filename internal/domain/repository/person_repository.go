package repository

import (
	"context"
	"effective_mobile_test_task/internal/domain/entities"
	"errors"
)

var (
	ErrPersonNotFound = errors.New("person not found")
	ErrDuplicateEntry = errors.New("duplicate person entry")
)

type FilterParams struct {
	Name        *string
	Surname     *string
	Patronymic  *string
	Gender      *string
	MinAge      *int
	MaxAge      *int
	Nationality *string
}

type Pagination struct {
	Limit  int
	Offset int
}

type PersonRepository interface {
	// Create создает новую запись о человеке и возвращает созданную сущность
	Create(ctx context.Context, person *entities.Person) (*entities.Person, error)

	// GetByID возвращает человека по его ID
	GetByID(ctx context.Context, id int) (*entities.Person, error)

	// Update обновляет данные существующего человека
	Update(ctx context.Context, id int, person *entities.Person) (*entities.Person, error)

	// Delete удаляет запись о человеке по ID
	Delete(ctx context.Context, id int) error

	// List возвращает список людей с учетом фильтров и пагинации
	// Возвращает список людей и общее количество записей (для пагинации)
	List(ctx context.Context, filters FilterParams, pagination Pagination) ([]entities.Person, int, error)

	// Exists проверяет, существует ли человек с такими же name, surname и patronymic
	Exists(ctx context.Context, name, surname string, patronymic *string) (bool, error)
}
