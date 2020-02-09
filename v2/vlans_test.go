package netplan

import (
	"testing"

	go_netplan_types "github.com/moznion/go-netplan-types"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSerializeEmptyVLAN(t *testing.T) {
	given := VLAN{}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`{}
`), marshal)

	var unmarshal VLAN
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}

func TestSerializeVLAN(t *testing.T) {
	given := VLAN{
		Device: Device{
			DHCP4: go_netplan_types.NillableBoolOf(true),
			DHCP6: go_netplan_types.NillableBoolOf(false),
		},
		ID:   go_netplan_types.NillableUint16Of(1),
		Link: go_netplan_types.NillableStringOf("link-1"),
	}

	marshal, err := yaml.Marshal(&given)
	assert.NoError(t, err)
	assert.EqualValues(t, []byte(`dhcp4: true
dhcp6: false
id: 1
link: link-1
`), marshal)

	var unmarshal VLAN
	err = yaml.Unmarshal(marshal, &unmarshal)
	assert.NoError(t, err)
	assert.EqualValues(t, given, unmarshal)
}
