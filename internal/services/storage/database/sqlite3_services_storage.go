package database

import (
	"context"

	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
	"github.com/rcmendes/crud-example-go/internal/services/core/errors"
)

type SQLite3ServicesStorage struct{}

func (storage *SQLite3ServicesStorage) FindAll(ctx context.Context) (entities.ServiceList, error) {
	query := "SELECT id, created_at, updated_at, name, description FROM services"
	rows, err := DB.QueryxContext(ctx, query)

	if err != nil {
		return nil, errors.DBQueryError(err)
	}

	var list entities.ServiceList

	for rows.Next() {
		var model ServiceModel
		err = rows.StructScan(&model)
		if err != nil {
			return nil, errors.DBQueryError(err)
		}

		service, err := model.ToEntity()
		if err != nil {
			return nil, err
		}

		list = append(list, service)
	}

	return list, nil
}

func (storage *SQLite3ServicesStorage) ExistsByName(ctx context.Context, name string) (bool, error) {
	query := "SELECT id FROM services WHERE name LIKE ?"
	row := DB.QueryRowxContext(ctx, query, name)

	hasReturn := row != nil

	var err error

	if row != nil {
		err = errors.DBQueryError(row.Err())
	}

	return hasReturn, err
}

func (storage *SQLite3ServicesStorage) Insert(ctx context.Context, service entities.Service) error {
	query := "INSERT INTO services(id, created_at, updated_at, name, description) VALUES(?,?,?,?,?)"
	_, err := DB.ExecContext(ctx, query, service.ID().String(), service.CreatedAt(),
		service.UpdatedAt(), service.Name(), service.Description())

	if err != nil {
		return errors.DBQueryError(err)
	}

	return nil
}
