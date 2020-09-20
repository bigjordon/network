package socket

import (
	"net"
	"testing"
)

func TestNetLis(t *testing.T) {
	t.Log("begin resolve ...")
	addr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err != nil {
		t.Errorf("resolves occure error")
		return
	}

	t.Log("begin listen ...")
	conn, err := net.ListenIP("ip4:icmp", addr)
	if err != nil {
		t.Errorf("Listen error: %v", err)
		return
	}
	buf := make([]byte, 1024)
	t.Log("begin read ...")
	for {
		n, a, err := conn.ReadFrom(buf)
		if err != nil {
			t.Errorf("read error %v", err)
			return
		}
		t.Logf("read %v byte, addr: %v, data: %v", n, a, buf[:n])
	}
}
