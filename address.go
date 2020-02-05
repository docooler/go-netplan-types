package netplan

import (
	"fmt"
	"strconv"
	"strings"
)

type Address struct {
	Address   string
	PrefixLen *NilableUint8
}

func (addr *Address) MarshalYAML() (interface{}, error) {
	if addr.PrefixLen == nil || !addr.PrefixLen.isAssigned {
		return addr.Address, nil
	}

	return fmt.Sprintf("%s/%d", addr.Address, addr.PrefixLen.val), nil
}

func (addr *Address) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	if err := unmarshal(&val); err != nil {
		return err
	}

	ss := strings.Split(val, "/")
	if len(ss) <= 0 {
		return nil
	}

	addr.Address = ss[0]
	if len(ss) >= 2 {
		prefixLen, err := strconv.ParseUint(ss[1], 10, 8)
		if err != nil {
			return err
		}
		addr.PrefixLen = NilableUint8Of(uint8(prefixLen))
	}

	return nil
}