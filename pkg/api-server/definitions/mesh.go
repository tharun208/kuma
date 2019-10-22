package definitions

import (
	"bytes"
	"encoding/json"
	"github.com/Kong/kuma/api/mesh/v1alpha1"
	"github.com/Kong/kuma/pkg/core"
	"github.com/Kong/kuma/pkg/core/resources/apis/mesh"
	"github.com/Kong/kuma/pkg/core/resources/model"
	"github.com/Kong/kuma/pkg/core/resources/model/rest"
	"github.com/gogo/protobuf/jsonpb"
)

var MeshWsDefinition = ResourceWsDefinition{
	Name: "Mesh",
	Path: "meshes",
	ResourceFactory: func() model.Resource {
		return &mesh.MeshResource{}
	},
	ResourceListFactory: func() model.ResourceList {
		return &mesh.MeshResourceList{}
	},
	Sample: meshRestResource{
		ResourceMeta: rest.ResourceMeta{
			Type: "Mesh",
			Name: "default",
		},
		Mesh: v1alpha1.Mesh{
			Mtls: &v1alpha1.Mesh_Mtls{
				Ca: &v1alpha1.CertificateAuthority{
					Type: &v1alpha1.CertificateAuthority_Builtin_{
						Builtin: &v1alpha1.CertificateAuthority_Builtin{},
					},
				},
				Enabled: true,
			},
			Logging: &v1alpha1.Logging{
				DefaultBackend: "logstash",
				Backends: []*v1alpha1.LoggingBackend{
					{
						Name:   "logstash",
						Format: "",
						Type: &v1alpha1.LoggingBackend_Tcp_{
							Tcp: &v1alpha1.LoggingBackend_Tcp{
								Address: "logstash:1234",
							},
						},
					},
				},
			},
		},
	},
	SampleList: nil,
}

type meshRestResource struct {
	rest.ResourceMeta
	v1alpha1.Mesh
}

func (m meshRestResource) MarshalJSON() ([]byte, error) {
	core.Log.Info("TEST")
	meta, err := json.Marshal(m.ResourceMeta)
	if err != nil {
		return nil, err
	}
	//if m == nil {
	//	return meta, nil
	//}

	var buf bytes.Buffer
	if err := (&jsonpb.Marshaler{}).Marshal(&buf, &m); err != nil {
		return nil, err
	}
	spec := buf.Bytes()

	var obj map[string]json.RawMessage
	if err := json.Unmarshal(meta, &obj); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(spec, &obj); err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

func (m meshRestResource) UnmarshalJSON(data []byte) error {
	core.Log.Info("TEST2")
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	//if m == nil {
	//	return nil
	//}
	if err := (&jsonpb.Unmarshaler{AllowUnknownFields: true}).Unmarshal(bytes.NewReader(data), &m); err != nil {
		return err
	}
	return nil
}

type meshRestListResource struct {
	Items []meshRestResource `json:"items"`
}
