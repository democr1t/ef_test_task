package postgres

import (
	"context"
	"effective_mobile_test_task/internal/domain/entities"
	"effective_mobile_test_task/internal/domain/repository"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PersonRepository struct {
	pool *pgxpool.Pool
}

func (r *PersonRepository) Create(ctx context.Context, person *entities.Person) (*entities.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PersonRepository) Update(ctx context.Context, id int, person *entities.Person) (*entities.Person, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PersonRepository) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (r *PersonRepository) List(ctx context.Context, filters repository.FilterParams, pagination repository.Pagination) ([]entities.Person, int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PersonRepository) Exists(ctx context.Context, name, surname string, patronymic *string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewPersonRepository(pool *pgxpool.Pool) *PersonRepository {
	return &PersonRepository{pool: pool}
}

func (r *PersonRepository) GetByID(ctx context.Context, id int) (*entities.Person, error) {
	slog.Debug("start repository get by id")

	query := "SELECT * FROM persons WHERE id = $1"

	var person entities.Person

	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&person.ID,
		&person.Name,
		&person.Surname,
		&person.Patronymic,
		&person.Gender,
		&person.Age,
		&person.Nationality,
		&person.CreatedAt,
		&person.UpdatedAt,
	)
	fmt.Println(person)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			slog.Debug("PersonRepository.GetByID: no person found")
			return nil, errors.New("person not found")
		}

		return nil, fmt.Errorf("failed to get person by id: %w", err)
	}

	return &person, nil
}
