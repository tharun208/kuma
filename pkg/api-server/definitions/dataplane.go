package definitions

import (
	"github.com/Kong/kuma/api/mesh/v1alpha1"
	"github.com/Kong/kuma/pkg/core/resources/apis/mesh"
	"github.com/Kong/kuma/pkg/core/resources/model"
	"github.com/Kong/kuma/pkg/core/resources/model/rest"
)

var DataplaneWsDefinition = ResourceWsDefinition{
	Name: "Dataplane",
	Path: "dataplanes",
	ResourceFactory: func() model.Resource {
		return &mesh.DataplaneResource{}
	},
	ResourceListFactory: func() model.ResourceList {
		return &mesh.DataplaneResourceList{}
	},
	Sample: dataplaneRestResource{
		ResourceMeta: rest.ResourceMeta{
			Type: "Dataplane",
			Name: "web-01",
			Mesh: "default",
		},
		Dataplane: v1alpha1.Dataplane{
			Networking: &v1alpha1.Dataplane_Networking{
				Inbound: []*v1alpha1.Dataplane_Networking_Inbound{
					{
						Interface: "127.0.0.1:8080:9090",
						Tags: map[string]string{
							"service": "web",
							"version": "v1",
						},
					},
					{
						Interface: "127.0.0.1:8081:9091",
						Tags: map[string]string{
							"service": "web-api",
						},
					},
				},
				Outbound: []*v1alpha1.Dataplane_Networking_Outbound{
					{
						Interface: ":30000",
						Service:   "backend",
					},
					{
						Interface: ":40000",
						Service:   "stats",
					},
				},
				TransparentProxying: &v1alpha1.Dataplane_Networking_TransparentProxying{
					RedirectPort: 1234,
				},
			},
		},
	},
	SampleList: dataplaneRestListResource{},
}

type dataplaneRestResource struct {
	rest.ResourceMeta
	v1alpha1.Dataplane
}

type dataplaneRestListResource struct {
	Items []dataplaneRestResource `json:"items"`
}
