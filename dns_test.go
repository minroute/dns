package dns

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Printf("%+v", A("www.baidu.com"))
	fmt.Printf("%+v", AAAA("www.baidu.com"))
}

func TestMX(t *testing.T) {
	fmt.Printf("%+v", Mx("minroute.com"))
}

func TestNs(t *testing.T) {
	fmt.Printf("%+v", Ns("minroute.com"))
}

func TestTxt(t *testing.T) {
	fmt.Printf("%+v", Txt("minroute.com"))
}

func TestCname(t *testing.T) {
	fmt.Printf("%+v", Cname("www.minroute.com"))
}

func TestPtr(t *testing.T) {
	fmt.Printf("%+v", Ptr("159.69.65.104"))
}

func TestSrv(t *testing.T) {
	fmt.Printf("%+v", Srv("xmpp-server", "tcp", "golang.org"))
}
