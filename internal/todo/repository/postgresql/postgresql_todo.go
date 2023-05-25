package postgresql

import (
	"context"
	"database/sql"

	"github.com/ckwhen/go-todo/internal/domain"

	"github.com/sirupsen/logrus"
)

type postgresqlTodoRepository struct {
	db *sql.DB
}

func NewPostgresqlTodoRepository(db *sql.DB) domain.TodoRepository {
	return &postgresqlTodoRepository{db}
}

func (p *postgresqlTodoRepository) GetAll(ctx context.Context) ([]domain.Todo, error) {
	rows, err := p.db.Query("SELECT id, task, status, created_at, updated_at FROM todos")
	todos := make([]domain.Todo, 0)

	if err != nil {
		logrus.Error(err)
		return todos, err
	}

	for rows.Next() {
		todo := new(domain.Todo)

		err = rows.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			logrus.Error(err)
			return todos, err
		}

		todos = append(todos, *todo)
	}

	return todos, nil
}

// func (p *postgresqlTodoRepository) Find(ctx context.Context) (*domain.Todo, error) {
// 	row := p.db.QueryRow("SELECT * FROM todos")
// 	d := &domain.Todo{}

// 	if err := row.Scan(&d.ID, &d.Task, &d.Status); err != nil {
// 		logrus.Error(err)
// 		return nil, err
// 	}
// 	return d, nil
// }

func (p *postgresqlTodoRepository) Store(ctx context.Context, d *domain.Todo) error {
	_, err := p.db.Exec(
		"INSERT INTO todos (id, task, status) VALUES ($1, $2, $3)",
		d.ID, d.Task, d.Status,
	)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// func (p *postgresqlTodoRepository) UpdateStatus(ctx context.Context, d *domain.Todo) error {
// 	_, err := p.db.Exec(
// 		"UPDATE todos SET status=$1 WHERE id=$2",
// 		d.Status, d.ID,
// 	)

// 	if err != nil {
// 		logrus.Error(err)
// 		return err
// 	}

// 	return nil
// }
