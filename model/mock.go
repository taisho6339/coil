package model

import (
	"context"
	"errors"
	"net"

	"github.com/cybozu-go/coil"
	"github.com/cybozu-go/netutil"
)

type mock struct {
	globalPool  *coil.AddressPool
	defaultPool *coil.AddressPool

	offset uint32
}

// NewMock returns a mock model for testing.
func NewMock() Model {
	_, gsubnet, _ := net.ParseCIDR("99.88.77.0/28")
	_, lsubnet, _ := net.ParseCIDR("10.10.0.0/16")
	return &mock{
		globalPool: &coil.AddressPool{
			Subnets:   []*net.IPNet{gsubnet},
			BlockSize: 0,
		},
		defaultPool: &coil.AddressPool{
			Subnets:   []*net.IPNet{lsubnet},
			BlockSize: 5,
		},
	}
}

func (m *mock) GetAllocatedIPs(ctx context.Context, block *net.IPNet) (map[string]net.IP, error) {
	return nil, nil
}

func (m *mock) AllocateIP(ctx context.Context, block *net.IPNet, key string) (net.IP, error) {
	if m.offset > 2 {
		return nil, errors.New("no more IP")
	}

	newIP := netutil.IntToIP4(netutil.IP4ToInt(block.IP) + m.offset)
	m.offset++
	return newIP, nil
}

func (m *mock) FreeIP(ctx context.Context, block *net.IPNet, ip net.IP) error {
	return nil
}

func (m *mock) GetMyBlocks(ctx context.Context, node string) (map[string][]*net.IPNet, error) {
	return nil, nil
}

func (m *mock) AcquireBlock(ctx context.Context, node, poolName string) (*net.IPNet, error) {
	switch poolName {
	case "global":
		_, block, _ := net.ParseCIDR("99.88.77.3/32")
		return block, nil
	case "default":
		_, block, _ := net.ParseCIDR("10.10.0.32/27")
		return block, nil
	}
	return nil, ErrNotFound
}

func (m *mock) ReleaseBlock(ctx context.Context, node, poolName string, block *net.IPNet) error {
	return nil
}

func (m *mock) AddPool(ctx context.Context, name string, subnet *net.IPNet, blockSize int) error {
	return nil
}

func (m *mock) AddSubnet(ctx context.Context, name string, n *net.IPNet) error {
	return nil
}

func (m *mock) GetPool(ctx context.Context, name string) (*coil.AddressPool, error) {
	switch name {
	case "global":
		return m.globalPool, nil
	case "default":
		return m.defaultPool, nil
	}
	return nil, ErrNotFound
}

func (m *mock) RemovePool(ctx context.Context, name string) error {
	return nil
}
