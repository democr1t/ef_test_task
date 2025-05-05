package services

import (
	"context"
	"effective_mobile_test_task/internal/domain/entities"
	"effective_mobile_test_task/internal/domain/repository"
	"errors"
	"log/slog"
)

type PersonService struct {
	personRepository repository.PersonRepository
}

func NewPersonService(repo repository.PersonRepository) *PersonService {
	return &PersonService{
		personRepository: repo,
	}
}

func (s *PersonService) GetByID(ctx context.Context, id int) (*entities.Person, error) {
	// Валидация входных данных
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	// Получаем данные из репозитория
	dbPerson, err := s.personRepository.GetByID(ctx, id)
	slog.Debug("person from db repo", slog.Any("dbPerson", dbPerson))
	if err != nil {
		slog.Error(err.Error())
		if errors.Is(err, repository.ErrPersonNotFound) {
			return nil, errors.New("person not found")
		}
		return nil, err
	}
	person := dbPerson
	// Конвертируем модель репозитория в сервисную модель
	//person := &entities.Person{
	//	ID:          dbPerson.ID,
	//	Name:        dbPerson.Name,
	//	Surname:     dbPerson.Surname,
	//	Patronymic:  dbPerson.Patronymic,
	//	Gender:      dbPerson.Gender,
	//	Age:         dbPerson.Age,
	//	Nationality: dbPerson.Nationality,
	//}
	slog.Debug("person after convert", slog.Any("person", person))
	return person, nil
}
