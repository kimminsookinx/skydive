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

// LogicalRouterPort defines the type used by both libovsdb and skydive for table Logical_Router_Port
// easyjson:json
// gendecoder
type LogicalRouterPort struct {
	UUID           string            `ovsdb:"_uuid" json:",omitempty" `
	Enabled        []bool            `ovsdb:"enabled" json:",omitempty" `
	ExternalIDs    map[string]string `ovsdb:"external_ids" json:",omitempty" `
	GatewayChassis []string          `ovsdb:"gateway_chassis" json:",omitempty" `
	HaChassisGroup []string          `ovsdb:"ha_chassis_group" json:",omitempty" `
	Ipv6Prefix     []string          `ovsdb:"ipv6_prefix" json:",omitempty" `
	Ipv6RaConfigs  map[string]string `ovsdb:"ipv6_ra_configs" json:",omitempty" `
	MAC            string            `ovsdb:"mac" json:",omitempty" `
	Name           string            `ovsdb:"name" json:",omitempty" `
	Networks       []string          `ovsdb:"networks" json:",omitempty" `
	Options        map[string]string `ovsdb:"options" json:",omitempty" `
	Peer           []string          `ovsdb:"peer" json:",omitempty" `

	ExternalIDsMeta   graph.Metadata `json:",omitempty" field:"Metadata"`
	Ipv6RaConfigsMeta graph.Metadata `json:",omitempty" field:"Metadata"`
	OptionsMeta       graph.Metadata `json:",omitempty" field:"Metadata"`
}

func (t *LogicalRouterPort) Metadata() graph.Metadata {
	// Generate Metadata from maps
	t.ExternalIDsMeta = graph.NormalizeValue(t.ExternalIDs).(map[string]interface{})
	t.Ipv6RaConfigsMeta = graph.NormalizeValue(t.Ipv6RaConfigs).(map[string]interface{})
	t.OptionsMeta = graph.NormalizeValue(t.Options).(map[string]interface{})

	return graph.Metadata{
		"Type":    "LogicalRouterPort",
		"Manager": "ovn",
		"UUID":    t.GetUUID(),
		"Name":    t.GetName(),
		"OVN":     t,
	}
}

func (t *LogicalRouterPort) GetUUID() string {
	return t.UUID
}

func (t *LogicalRouterPort) GetName() string {
	if name := t.Name; name != "" {
		return name
	}
	return t.GetUUID()
}

// LogicalRouterPortDecoder implements a json message raw decoder
func LogicalRouterPortDecoder(raw json.RawMessage) (getter.Getter, error) {
	var t LogicalRouterPort
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, fmt.Errorf("unable to unmarshal LogicalRouterPort metadata %s: %s", string(raw), err)
	}
	return &t, nil
}
