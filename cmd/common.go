package cmd

import (
	coap "github.com/moroen/gocoap"
	"net"
	"errors"
	"net/url"
	"strconv"
	"fmt"
	"encoding/json"
	
)

var MalformedUriError = errors.New("Malformed uri")

type Message struct {
	Status string
	Result string
}

func request(uri, payload string) (status, response string) {
	var req coap.RequestParams

	scheme, host, port, path, err := processURI(uri)
	if err != nil {
		if err == MalformedUriError {
			return "MalformedUri", ""
		} else {
			panic(err.Error())
		}
	}

	switch scheme {
	case "coap":
		break
	case "coaps":
		if ident == "" {
			printResponse("MissingIdent", "")
			return
		}

		if key == "" {
			printResponse("MissingKey", "")
			return
		}
		req = coap.RequestParams{Host: host, Port: port, Uri: path, Id: ident, Key: key}
	}

	resp, err := coap.GetRequest(req)
	status = errorToStatus(err)
	return status, string(resp)
}

func printResponse(status, result string) {
	res := Message{Status: status, Result: result }

	jsonObj, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(jsonObj))
}

func errorToStatus(err error) string {
	if err == nil {
		return "ok"
	}

	switch err {
	case coap.UriNotFound:
		return "UriNotFound"
	case coap.ErrorHandshake:
		return "HandshakeError"
	case coap.Unauthorized:
		return "Unauthorized"
	}
	return "UnknownStatus"
}

func processURI(uri string) (protocol string, host string, port int, path string, err error) {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err.Error())
	}

	protocol = u.Scheme
	if protocol == "" {
		err = MalformedUriError
		return
	}

	host, strPort, err := net.SplitHostPort(u.Host)
	if err != nil {
		err = MalformedUriError
		return
	}

	port, err = strconv.Atoi(strPort)
	if err != nil {
		err = MalformedUriError
		return

	}
	path = u.Path

	return
}
