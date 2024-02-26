package infrastructure

import (
	"github.com/caarlos0/env/v10"
	"started-fiber/internal/article"
	"started-fiber/internal/author"
	"started-fiber/internal/config"
	"started-fiber/internal/domain"
	"started-fiber/pkg/xlogger"
)

var (
	cfg config.Config

	authorRepository  domain.AuthorRepository
	articleRepository domain.ArticleRepository

	articleService domain.ArticleService
)

func init() {
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	xlogger.Setup(cfg)
	dbSetup()

	authorRepository = author.NewMysqlAuthorRepository(db)
	articleRepository = article.NewMysqlArticleRepository(db)

	articleService = article.NewArticleService(articleRepository, authorRepository)
}
