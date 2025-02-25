package net

import (
	"github.com/stretchr/testify/assert"
	"net/netip"
	"testing"
)

func TestBindingLookupWithLocalhost(t *testing.T) {
	expectedBindIP := "::1"

	bindIP := BindingLookup("localhost")
	assert.Equal(t, bindIP, expectedBindIP)
}

func TestBindingLookupWithIpV4(t *testing.T) {
	expectedBindIP := "127.0.0.1"

	bindIP := BindingLookup("127.0.0.1")
	assert.Equal(t, bindIP, expectedBindIP)
}

func TestBindingLookupWithIpV6(t *testing.T) {
	expectedBindIP := "2001:db8::1:0:0:1"

	bindIP := BindingLookup("2001:db8::1:0:0:1")
	assert.Equal(t, bindIP, expectedBindIP)
}

func TestRegisterAddrWithLocalhost(t *testing.T) {
	expectedAddr, err := netip.ParseAddr("::1")
	if err != nil {
		t.Errorf("invalid expected IP: %+v", expectedAddr)
	}

	addr := RegisterAddr("localhost")
	if addr != expectedAddr {
		t.Errorf("invalid IP: %+v", addr)
	}
	assert.Equal(t, addr, expectedAddr)
}

func TestRegisterAddrWithIpV4(t *testing.T) {
	expectedAddr, err := netip.ParseAddr("127.0.0.1")
	if err != nil {
		t.Errorf("invalid expected IP: %+v", expectedAddr)
	}

	addr := RegisterAddr("127.0.0.1")
	if addr != expectedAddr {
		t.Errorf("invalid IP: %+v", addr)
	}
}

func TestRegisterAddrWithIpV6(t *testing.T) {
	expectedAddr, err := netip.ParseAddr("2001:db8::1:0:0:1")
	if err != nil {
		t.Errorf("invalid expected IP: %+v", expectedAddr)
	}

	addr := RegisterAddr("2001:db8::1:0:0:1")
	if addr != expectedAddr {
		t.Errorf("invalid IP: %+v", addr)
	}
}
