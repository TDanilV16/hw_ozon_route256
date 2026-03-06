package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"week3/workshop/internal/pkg/db"
	"week3/workshop/internal/pkg/repository"
)

type ArticleRepos struct {
	db *db.Database
}

func NewArticles(db *db.Database) *ArticleRepos {
	return &ArticleRepos{db: db}
}

func (r *ArticleRepos) Add(ctx context.Context, article *repository.Article) (int64, error) {
	var id int64

	err := r.db.ExecQueryRow(ctx, `INSERT INTO articles(name,rating) VALUES($1, $2) RETURNING id`, article.Name, article.Rating).Scan(&id)
	return id, err
}

func (r *ArticleRepos) GetByID(ctx context.Context, id int64) (*repository.Article, error) {
	var a repository.Article

	err := r.db.Get(ctx, &a, `SELECT id, name, rating, created_at FROM articles WHERE id = $1`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrObjectNotFound
	}

	return &a, err
}
