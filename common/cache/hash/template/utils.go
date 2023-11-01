package template

import (
	"context"
	"github.com/google/uuid"
	"go-micro-framework/common/conf"
	"go-micro-framework/global/cache/cluster"
)

func Create(cacheConf *conf.CacheConf) error {
	cluster, err := cluster.GetClient(cacheConf, GetWhich())
	if err != nil {
		return err
	}

	err = cluster.HMSet(
		context.Background(),
		Format,
		map[string]interface{}{
			Field1: uuid.New().String(),
			Field2: uuid.New().String(),
			Field3: uuid.New().String(),
		},
	).Err()
	if err != nil {
		return err
	}

	return cluster.Expire(context.Background(), Format, Expire).Err()
}
