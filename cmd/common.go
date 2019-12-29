package cmd

import (
	"net"
	"net/url"
	"strconv"
)

type Message struct {
	Status string
	Result string
}

func processURI(uri string) (protocol string, host string, port int, path string, err error) {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err.Error())
	}

	protocol = u.Scheme
	host, strPort, err := net.SplitHostPort(u.Host)
	port, err = strconv.Atoi(strPort)
	path = u.Path

	return
}
