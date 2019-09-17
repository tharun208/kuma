package manager_test

import (
	"context"
	"github.com/Kong/kuma/api/mesh/v1alpha1"
	"github.com/Kong/kuma/pkg/core/resources/apis/mesh"
	"github.com/Kong/kuma/pkg/core/resources/manager"
	"github.com/Kong/kuma/pkg/core/resources/model"
	"github.com/Kong/kuma/pkg/core/resources/store"
	"github.com/Kong/kuma/pkg/plugins/resources/memory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

type countingResourcesManager struct {
	s           store.ResourceStore
	getQueries  int
	listQueries int
}

func (c *countingResourcesManager) Create(ctx context.Context, res model.Resource, fn ...store.CreateOptionsFunc) error {
	return c.s.Create(ctx, res, fn...)
}

func (c *countingResourcesManager) Update(ctx context.Context, res model.Resource, fn ...store.UpdateOptionsFunc) error {
	return c.s.Update(ctx, res, fn...)
}

func (c *countingResourcesManager) Delete(ctx context.Context, res model.Resource, fn ...store.DeleteOptionsFunc) error {
	return c.s.Delete(ctx, res, fn...)
}

func (c *countingResourcesManager) Get(ctx context.Context, res model.Resource, fn ...store.GetOptionsFunc) error {
	c.getQueries++
	return c.s.Get(ctx, res, fn...)
}

func (c *countingResourcesManager) List(ctx context.Context, list model.ResourceList, fn ...store.ListOptionsFunc) error {
	c.listQueries++
	return c.s.List(ctx, list, fn...)
}

var _ manager.ResourceManager = &countingResourcesManager{}

var _ = Describe("Cached Resource Manager", func() {

	var cachedManager manager.ResourceManager
	var countingManager *countingResourcesManager
	var res *mesh.DataplaneResource
	expiration := 100 * time.Millisecond

	BeforeEach(func() {
		// given
		countingManager = &countingResourcesManager{
			s: memory.NewStore(),
		}
		cachedManager = manager.NewCachedManager(countingManager, expiration)

		// and created resources
		res = &mesh.DataplaneResource{
			Spec: v1alpha1.Dataplane{
				Networking: &v1alpha1.Dataplane_Networking{
					Inbound: []*v1alpha1.Dataplane_Networking_Inbound{
						{
							Interface: "127.0.0.1:80:8080",
						},
					},
				},
			},
		}
		err := cachedManager.Create(context.Background(), res, store.CreateByKey("default", "dp-1", "default"))
		Expect(err).ToNot(HaveOccurred())
	})

	It("should cache Get() queries", func() {
		// when fetched resources multiple times
		fetch := func() mesh.DataplaneResource {
			fetched := mesh.DataplaneResource{}
			err := cachedManager.Get(context.Background(), &fetched, store.GetByKey("default", "dp-1", "default"))
			Expect(err).ToNot(HaveOccurred())
			return fetched
		}

		for i := 0; i < 10; i++ {
			fetch()
		}

		// then real manager should be called only once
		Expect(fetch().Spec).To(Equal(res.Spec))
		Expect(countingManager.getQueries).To(Equal(1))

		// when
		time.Sleep(expiration)

		// then
		Expect(fetch().Spec).To(Equal(res.Spec))
		Expect(countingManager.getQueries).To(Equal(2))
	})

	It("should cache List() queries", func() {
		// when fetched resources multiple times
		fetch := func() mesh.DataplaneResourceList {
			fetched := mesh.DataplaneResourceList{}
			err := cachedManager.List(context.Background(), &fetched, store.ListByMesh("default"))
			Expect(err).ToNot(HaveOccurred())
			return fetched
		}

		for i := 0; i < 10; i++ {
			fetch()
		}

		// then real manager should be called only once
		Expect(fetch().Items).To(HaveLen(1))
		Expect(fetch().Items[0].GetSpec()).To(Equal(&res.Spec))
		Expect(countingManager.listQueries).To(Equal(1))

		// when
		time.Sleep(expiration)

		// then
		Expect(fetch().Items).To(HaveLen(1))
		Expect(fetch().Items[0].GetSpec()).To(Equal(&res.Spec))
		Expect(countingManager.listQueries).To(Equal(2))
	})
})
