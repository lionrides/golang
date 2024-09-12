package user

import (
	"context"
	"fmt"
	"golang/internal/domain"
	"log"
)

func (repo *UserRepo) GetAllUser(ctx context.Context) ([]domain.User, error) {
	var user domain.User

	query := "SELECT *  FROM user"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return []domain.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.RoleID, &user.Email, &user.Name)
		if err != nil {
			return []domain.User{}, err
		}
	}
	return []domain.User{}, nil
}

func (repo *UserRepo) CreateUser(user domain.User) (msg string, err error) {

	sqlStatement := "INSERT INTO user (role_id, email, 'password', 'name', last_access, created_at, updated_at, deleted_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`"
	id := 0
	err = repo.DB.QueryRow(sqlStatement, user).Scan(&id)
	if err != nil {
		return "Gagal menambahkan data", err
	}
	fmt.Printf("New record ID is: %d\n", id)
	return "Success", nil
}

func (repo *UserRepo) UpdateUser(id int, user domain.User) (msg string, err error) {

	sqlStatement := `
    UPDATE user
    SET  role_id = $2, email = $3, 'password' = $4, 'name' = $5, last_access =$6, created_at = $7, updated_at $8, deleted_at = $9
    WHERE id = $1;`
	res, err := repo.DB.Exec(sqlStatement, id, user)
	if err != nil {
		return "Gagal mengupdate data user", err
	}

	// Get the number of rows affected
	_, err = res.RowsAffected()
	if err != nil {
		return "Success", nil
	}
}

func (repo *UserRepo) DeleteUser(id int) (msg string, err error) {

	sqlStatement := `
    DELETE FROM user
    WHERE id = $1;`

	res, err := repo.DB.Exec(sqlStatement, 1)
	if err != nil {
		return "Gagal menghapus data user", err
	}

	// Get the number of rows affected
	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d row(s) deleted\n", count)
	return "Success", nil
}
