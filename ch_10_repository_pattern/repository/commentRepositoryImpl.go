package repository

import (
	"belajar-golang-database/ch_10_repository_pattern/entities"
	"context"
	"database/sql"
	"errors"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

// membuat constructornya
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Save(ctx context.Context, comment entities.Comment) (entities.Comment, error) {

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := repo.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int(id)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int) (entities.Comment, error) {
	query := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, query, id)

	comment := entities.Comment{} // membnuat comment kosong
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comment, err
		}
		return comment, nil // return comment
	} else {
		return comment, errors.New("comment not found")
	}
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entities.Comment, error) {
	query := "SELECT id, email, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entities.Comment
	for rows.Next() {
		comment := entities.Comment{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment) // push comment
	}
	return comments, nil
}
