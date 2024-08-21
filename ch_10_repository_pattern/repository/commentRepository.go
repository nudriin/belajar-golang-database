package repository

import (
	"belajar-golang-database/ch_10_repository_pattern/entities"
	"context"
)

type CommentRepository interface {
	Save(ctx context.Context, comment entities.Comment) (entities.Comment, error)
	FindById(ctx context.Context, id int) (entities.Comment, error)
	FindAll(ctx context.Context) ([]entities.Comment, error)
}
