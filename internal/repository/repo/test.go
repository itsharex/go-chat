package repo

import (
	"go-chat/internal/pkg/ichat"
	"go-chat/internal/repository/model"
	"gorm.io/gorm"
)

type Test struct {
	ichat.Repo[model.Article]
}

func NewTest(db *gorm.DB) *Test {
	return &Test{Repo: ichat.Repo[model.Article]{Db: db}}
}

func (t *Test) name() {}