package postgres

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hibiki-horimi/go-todo-api/internal/config"
)

type ctxKey string

const dbCtxKey ctxKey = "db"

type Postgres struct {
	Todo Todo
}

func New() *Postgres {
	return &Postgres{
		Todo: &todo{},
	}
}

func Connect(conf *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		conf.DB.Host,
		conf.DB.User,
		conf.DB.Password,
		conf.DB.Name,
		conf.DB.Port,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().In(time.UTC)
		},
	})
}

func SetDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbCtxKey, db)
}

func DBFromContext(ctx context.Context) *gorm.DB {
	return ctx.Value(dbCtxKey).(*gorm.DB).WithContext(ctx).Debug()
}
