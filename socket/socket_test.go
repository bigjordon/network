package socket

import (
	"net"
	"testing"

	"github.com/zssky/log"
)

/*https://stackoverflow.com/questions/23205419/how-do-you-print-in-a-go-test-using-the-testing-package
In Go 1.14, go test -v will stream t.Log output as it happens, rather than hoarding it til the end of the test run.

Under Go 1.14 the fmt.Println and t.Log lines are interleaved, rather than waiting for the test to complete,
demonstrating that test output is streamed when go test -v is used.
 */

func TestNetLis(t *testing.T) {
	log.Debug("begin resolve ...")
	addr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err != nil {
		t.Errorf("resolves occure error")
		return
	}

	log.Debug("begin listen ...")
	conn, err := net.ListenIP("ip4:icmp", addr)
	if err != nil {
		log.Errorf("Listen error: %v", err)
		return
	}
	buf := make([]byte, 1024)
	log.Debug("begin read ...")
	for {
		n, a, err := conn.ReadFrom(buf)
		if err != nil {
			log.Errorf("read error %v", err)
			return
		}
		log.Debugf("read %v byte, addr: %v, data: %v", n, a, buf[:n])
	}
}