package interfaces

import (
	"context"
	"effective_mobile_test_task/internal/domain/entities"
)

type PersonService interface {
	// Create добавляет нового человека
	Create(ctx context.Context, person *entities.Person) (*entities.Person, error)

	// GetByID возвращает человека по ID
	GetByID(ctx context.Context, id int) (*entities.Person, error)

	// Update обновляет данные существующего человека
	Update(ctx context.Context, id int, person *entities.Person) (*entities.Person, error)

	// Delete удаляет человека по ID
	Delete(ctx context.Context, id int) error

	//// List возвращает список людей с фильтрацией и пагинацией
	//List(ctx context.Context, filters FilterParams, pagination Pagination) ([]Person, int, error)

	// Enrich дополняет данные человека (возраст, пол, национальность)
	Enrich(ctx context.Context, person *entities.Person) (*entities.Person, error)
}
