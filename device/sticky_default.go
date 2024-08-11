//go:build !linux

package device

import (
	"github.com/kterentev/wireguard-go/conn"
	"github.com/kterentev/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
