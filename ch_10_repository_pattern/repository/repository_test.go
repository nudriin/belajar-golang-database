package repository

import (
	"belajar-golang-database/ch_10_repository_pattern/entities"
	"belajar-golang-database/database"
	"context"
	"fmt"
	"testing"
)

func TestCommentRepoSave(t *testing.T) {
	// Membuat repository
	repo := NewCommentRepository(database.GetConnect())

	ctx := context.Background()
	comment := entities.Comment{
		Email:   "test@test.com",
		Comment: "test comment",
	}
	result, err := repo.Save(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentRepoFindById(t *testing.T) {
	// Membuat repository
	// 35
	repo := NewCommentRepository(database.GetConnect())

	ctx := context.Background()
	result, err := repo.FindById(ctx, 35)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentRepoFindByAll(t *testing.T) {
	// Membuat repository
	repo := NewCommentRepository(database.GetConnect())

	ctx := context.Background()
	result, err := repo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)

	}
}
