package repo

import (
	"github.com/Nevator27/um-help/config"
	"github.com/Nevator27/um-help/repo/mysql"
	"github.com/Nevator27/um-help/repo/redis"
)

type RepoManager struct {
	MySQL *mysql.Repo
	Redis *redis.Repo
}

func New(cfg *config.Config) (*RepoManager, error) {
	mysql, err := mysql.New(cfg)
	if err != nil {
		return nil, err
	}

	redis, err := redis.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RepoManager{
		MySQL: mysql,
		Redis: redis,
	}, nil
}
