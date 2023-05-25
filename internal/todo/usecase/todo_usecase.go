package usecase

import (
	"context"

	"github.com/ckwhen/go-todo/internal/domain"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type todoUsecase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUsecase(todoRepo domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		todoRepo: todoRepo,
	}
}

func (tu *todoUsecase) GetAll(ctx context.Context) ([]domain.Todo, error) {
	todos, err := tu.todoRepo.GetAll(ctx)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return todos, nil
}

// func (tu *todoUsecase) Find(ctx context.Context) (*domain.Todo, error) {
// 	todos, err := tu.todoRepo.Find(ctx)

// 	if err != nil {
// 		logrus.Error(err)
// 		return nil, err
// 	}

// 	return todos, nil
// }

func (tu *todoUsecase) Store(ctx context.Context, d *domain.Todo) error {
	d.ID = uuid.Must(uuid.NewV4()).String()
	d.Status = domain.TODO_ACTIVE

	if err := tu.todoRepo.Store(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

// func (tu *todoUsecase) UpdateStatus(ctx context.Context, d *domain.Todo) error {
// 	if d.Status == "" {
// 		err := errors.New("Status is blank")
// 		logrus.Error(err)
// 		return err
// 	}

// 	if err := tu.todoRepo.UpdateStatus(ctx, d); err != nil {
// 		logrus.Error(err)
// 		return err
// 	}

// 	return nil
// }
