# DNS wrapper 

This is a simple wrapper for the net.LookXXX functions. 

## Why?
Sometime, we just  need **clear and simple** way to get the DNS records.

## Usage

```
go get github.com/minroute/dns

dns.A("google.com")
dns.AAAA("google.com")
dns.CNAME("google.com")
dns.Mx("google.com")
dns.Ns("google.com")
dns.Ptr("google.com")
dns.Txt("google.com")
dns.Srv('xmpp-server","tcp","google.com")
```


## License
Mit License
