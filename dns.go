package dns

import (
	"net"
)

func Aaaa(domain string) []string {
	return a(domain)
}

func A(domain string) []string {
	return a(domain)
}

func a(domain string) []string {
	ips := []string{}
	rerecords, _ := net.LookupIP(domain)
	for _, ip := range rerecords {
		ips = append(ips, ip.String())
	}
	return ips
}

func Cname(domain string) string {
	cname, _ := net.LookupCNAME(domain)
	return cname
}

func Ptr(ip string) []string {
	ptr, _ := net.LookupAddr(ip)
	ptrSlice := []string{}
	for _, ptrvalue := range ptr {
		ptrSlice = append(ptrSlice, ptrvalue)
	}
	return ptrSlice
}

func Ns(domain string) []string {
	nsSlice := []string{}
	nameserver, _ := net.LookupNS(domain)
	for _, ns := range nameserver {
		nsSlice = append(nsSlice, ns.Host)
	}
	return nsSlice
}

type MxRecord struct {
	Host string
	Pref uint16
}

func Mx(domain string) []MxRecord {
	mxSlice := []MxRecord{}
	rerecords, _ := net.LookupMX(domain)
	for _, mx := range rerecords {
		mxSlice = append(mxSlice, MxRecord{Host: mx.Host, Pref: mx.Pref})
	}
	return mxSlice
}

// Srv is a DNS SRV record.
func Srv(service, proto, domain string) []*net.SRV {
	_, srvs, err := net.LookupSRV(service, proto, domain)
	if err != nil {
		return nil
	}
	return srvs
}

func Txt(domain string) []string {
	txtSlice := []string{}
	rerecords, _ := net.LookupTXT(domain)
	for _, txt := range rerecords {
		txtSlice = append(txtSlice, txt)
	}
	return txtSlice
}
