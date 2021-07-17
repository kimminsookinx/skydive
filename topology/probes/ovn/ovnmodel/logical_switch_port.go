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

// LogicalSwitchPort defines the type used by both libovsdb and skydive for table Logical_Switch_Port
// easyjson:json
// gendecoder
type LogicalSwitchPort struct {
	UUID             string            `ovsdb:"_uuid" json:",omitempty" `
	Addresses        []string          `ovsdb:"addresses" json:",omitempty" `
	Dhcpv4Options    []string          `ovsdb:"dhcpv4_options" json:",omitempty" `
	Dhcpv6Options    []string          `ovsdb:"dhcpv6_options" json:",omitempty" `
	DynamicAddresses []string          `ovsdb:"dynamic_addresses" json:",omitempty" `
	Enabled          []bool            `ovsdb:"enabled" json:",omitempty" `
	ExternalIDs      map[string]string `ovsdb:"external_ids" json:",omitempty" `
	HaChassisGroup   []string          `ovsdb:"ha_chassis_group" json:",omitempty" `
	Name             string            `ovsdb:"name" json:",omitempty" `
	Options          map[string]string `ovsdb:"options" json:",omitempty" `
	ParentName       []string          `ovsdb:"parent_name" json:",omitempty" `
	PortSecurity     []string          `ovsdb:"port_security" json:",omitempty" `
	Tag              []int             `ovsdb:"tag" json:",omitempty" `
	TagRequest       []int             `ovsdb:"tag_request" json:",omitempty" `
	Type             string            `ovsdb:"type" json:",omitempty" `
	Up               []bool            `ovsdb:"up" json:",omitempty" `

	ExternalIDsMeta graph.Metadata `json:",omitempty" field:"Metadata"`
	OptionsMeta     graph.Metadata `json:",omitempty" field:"Metadata"`
}

func (t *LogicalSwitchPort) Metadata() graph.Metadata {
	// Generate Metadata from maps
	t.ExternalIDsMeta = graph.NormalizeValue(t.ExternalIDs).(map[string]interface{})
	t.OptionsMeta = graph.NormalizeValue(t.Options).(map[string]interface{})

	return graph.Metadata{
		"Type":    "LogicalSwitchPort",
		"Manager": "ovn",
		"UUID":    t.GetUUID(),
		"Name":    t.GetName(),
		"OVN":     t,
	}
}

func (t *LogicalSwitchPort) GetUUID() string {
	return t.UUID
}

func (t *LogicalSwitchPort) GetName() string {
	if name := t.Name; name != "" {
		return name
	}
	return t.GetUUID()
}

// LogicalSwitchPortDecoder implements a json message raw decoder
func LogicalSwitchPortDecoder(raw json.RawMessage) (getter.Getter, error) {
	var t LogicalSwitchPort
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, fmt.Errorf("unable to unmarshal LogicalSwitchPort metadata %s: %s", string(raw), err)
	}
	return &t, nil
}
