package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Kong/kuma/pkg/core/resources/model"
	"github.com/Kong/kuma/pkg/core/resources/store"
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

type cachedManager struct {
	delegate ResourceManager
	cache    *cache.Cache
}

const getCacheKeyPrefix = "GET:"
const ListCacheKeyPrefix = "LIST:"

func NewCachedManager(delegate ResourceManager, expiration time.Duration) ResourceManager {
	return &cachedManager{
		delegate: delegate,
		cache:    cache.New(expiration, time.Duration(int64(float64(expiration)*0.9))),
	}
}

var _ ResourceManager = &cachedManager{}

func (c *cachedManager) Create(ctx context.Context, res model.Resource, fn ...store.CreateOptionsFunc) error {
	return c.delegate.Create(ctx, res, fn...)
}

func (c *cachedManager) Update(ctx context.Context, res model.Resource, fn ...store.UpdateOptionsFunc) error {
	return c.delegate.Update(ctx, res, fn...)
}

func (c *cachedManager) Delete(ctx context.Context, res model.Resource, fn ...store.DeleteOptionsFunc) error {
	return c.delegate.Delete(ctx, res, fn...)
}

func (c *cachedManager) clearCache(resKey model.ResourceKey) error {
	for key, _ := range c.cache.Items() {
		if strings.HasPrefix(getCacheKeyPrefix, key) {
			optsJson := strings.TrimPrefix(getCacheKeyPrefix, key)
			opts := store.GetOptions{}
			if err := json.Unmarshal([]byte(optsJson), &opts); err != nil {
				return err
			}
			if opts.Name == resKey.Name && opts.Namespace == resKey.Namespace && opts.Mesh == resKey.Mesh {
				c.cache.Delete(key)
			}
		} else if strings.HasPrefix(ListCacheKeyPrefix, key) {
			optsJson := strings.TrimPrefix(ListCacheKeyPrefix, key)
			opts := store.ListOptions{}
			if err := json.Unmarshal([]byte(optsJson), &opts); err != nil {
				return err
			}
			if opts.Mesh == resKey.Mesh || opts.Namespace == resKey.Namespace {
				c.cache.Delete(key)
			}
		}
	}
	return nil
}

func (c *cachedManager) Get(ctx context.Context, res model.Resource, fn ...store.GetOptionsFunc) error {
	opts := store.NewGetOptions(fn...)
	bytes, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("%s%s", getCacheKeyPrefix, string(bytes))
	obj, found := c.cache.Get(cacheKey)
	if !found {
		if err := c.delegate.Get(ctx, res, fn...); err != nil {
			return err
		}
		c.cache.SetDefault(cacheKey, res)
	} else {
		cached := obj.(model.Resource)
		if err := res.SetSpec(cached.GetSpec()); err != nil {
			return err
		}
		res.SetMeta(cached.GetMeta())
	}
	return nil
}

func (c cachedManager) List(ctx context.Context, list model.ResourceList, fn ...store.ListOptionsFunc) error {
	opts := store.NewListOptions(fn...)
	bytes, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("%s%s", ListCacheKeyPrefix, string(bytes))
	obj, found := c.cache.Get(cacheKey)
	if !found {
		if err := c.delegate.List(ctx, list, fn...); err != nil {
			return err
		}
		c.cache.SetDefault(cacheKey, list.GetItems())
	} else {
		resources := obj.([]model.Resource)
		for _, res := range resources {
			if err := list.AddItem(res); err != nil {
				return err
			}
		}
	}
	return nil
}
