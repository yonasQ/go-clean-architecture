package dbinstance

import (
	"context"
	"project-structure-template/internal/constants/model/dto"
)

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, first_name, middle_name, last_name, email, status, created_at, updated_at FROM users WHERE deleted_at IS NULL
`

func (q *DBInstance) GetAllUsers(ctx context.Context) ([]dto.User, error) {
	rows, err := q.Pool.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []dto.User{}
	for rows.Next() {
		i := dto.User{}
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.MiddleName,
			&i.LastName,
			&i.Email,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
