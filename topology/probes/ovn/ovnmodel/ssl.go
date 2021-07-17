//go:generate go run github.com/skydive-project/skydive/graffiti/gendecoder -package github.com/skydive-project/skydive/topology/probes/ovn/ovnmodel
//go:generate go run github.com/mailru/easyjson/easyjson $GOFILE

// Code generated by ovnmetagen
// DO NOT EDIT.

package ovnmodel

import (
	"encoding/json"
	"fmt"

	"github.com/skydive-project/skydive/graffiti/getter"
	"github.com/skydive-project/skydive/graffiti/graph"
)

// SSL defines the type used by both libovsdb and skydive for table SSL
// easyjson:json
// gendecoder
type SSL struct {
	UUID            string            `ovsdb:"_uuid" json:",omitempty" `
	BootstrapCaCert bool              `ovsdb:"bootstrap_ca_cert" json:",omitempty" `
	CaCert          string            `ovsdb:"ca_cert" json:",omitempty" `
	Certificate     string            `ovsdb:"certificate" json:",omitempty" `
	ExternalIDs     map[string]string `ovsdb:"external_ids" json:",omitempty" `
	PrivateKey      string            `ovsdb:"private_key" json:",omitempty" `
	SSLCiphers      string            `ovsdb:"ssl_ciphers" json:",omitempty" `
	SSLProtocols    string            `ovsdb:"ssl_protocols" json:",omitempty" `

	ExternalIDsMeta graph.Metadata `json:",omitempty" field:"Metadata"`
}

func (t *SSL) Metadata() graph.Metadata {
	// Generate Metadata from maps
	t.ExternalIDsMeta = graph.NormalizeValue(t.ExternalIDs).(map[string]interface{})

	return graph.Metadata{
		"Type":    "SSL",
		"Manager": "ovn",
		"UUID":    t.GetUUID(),
		"Name":    t.GetName(),
		"OVN":     t,
	}
}

func (t *SSL) GetUUID() string {
	return t.UUID
}

func (t *SSL) GetName() string {
	if name := t.UUID; name != "" {
		return name
	}
	return t.GetUUID()
}

// SSLDecoder implements a json message raw decoder
func SSLDecoder(raw json.RawMessage) (getter.Getter, error) {
	var t SSL
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, fmt.Errorf("unable to unmarshal SSL metadata %s: %s", string(raw), err)
	}
	return &t, nil
}
