package data

import (
	"context"
	"crypto/rsa"
	"database/sql"
	"encoding/base64"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
	"github.com/wzyjerry/auth/internal/conf"
	"github.com/wzyjerry/auth/internal/ent"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewPrivateKey)

// Data .
type Data struct {
	db    *PostgresClient
	redis *redis.Client
}

func NewPrivateKey(c *conf.Security, logger log.Logger) *rsa.PrivateKey {
	log := log.NewHelper(logger)
	result, err := base64.StdEncoding.DecodeString(os.Getenv(c.PrivateKey))
	if err != nil {
		log.Fatalf("failed parse base64 private key: %v", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(result)
	if err != nil {
		log.Fatalf("failed parse private key: %v", err)
	}
	return privateKey
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	db, err := sql.Open("pgx", os.Getenv(c.Postgres.Env))
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	ctx := context.Background()
	opt, err := redis.ParseURL(os.Getenv(c.Redis.Env))
	if err != nil {
		log.Fatalf("failed parse redis url: %v", err)
	}
	client := redis.NewClient(opt)
	client.Info(ctx)
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("failed connecting to redis: %v", err)
	}

	cleanup := func() {
		db.Close()
		client.Close()
		log.Info("closing the data resources")
	}
	data := &Data{
		db:    (*PostgresClient)(ent.NewClient(ent.Driver(drv))),
		redis: client,
	}
	// 数据库迁移
	if err := data.db.Schema.Create(ctx, schema.WithAtlas(true), schema.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return data, cleanup, nil
}

type PostgresClient ent.Client

func (c *PostgresClient) GenerateId() string {
	return primitive.NewObjectID().Hex()
}

func (c *PostgresClient) WithTx(ctx context.Context, f func(tx *ent.Tx) error) error {
	tx, err := (*ent.Client)(c).Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := f(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
