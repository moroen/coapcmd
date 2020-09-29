package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"

	coap "github.com/moroen/gocoap/v3"
)

var MalformedUriError = errors.New("Malformed uri")

type Message struct {
	Status string
	Result string
}

const GET = 1
const PUT = 2
const POST = 3

func request(method int, uri, payload string) (status, response string) {
	var req coap.RequestParams
	var resp []byte
	var err error

	scheme, host, port, path, err := processURI(uri)
	if err != nil {
		if err == MalformedUriError {
			return "MalformedUri", ""
		} else {
			panic(err.Error())
		}
	}

	req = coap.RequestParams{Host: host, Port: port, Uri: path}

	if scheme == "coaps" {
		if ident == "" {
			printResponse("MissingIdent", "")
			return
		}

		if key == "" {
			printResponse("MissingKey", "")
			return
		}

		req.Id = ident
		req.Key = key
	}

	if payload != "" {
		req.Payload = payload
	}

	switch method {
	case GET:
		resp, err = coap.GetRequest(req)
	case PUT:
		resp, err = coap.PutRequest((req))
	case POST:
		resp, err = coap.PostRequest((req))
	}

	status = errorToStatus(err)
	return status, string(resp)
}

func printResponse(status, result string) {
	res := Message{Status: status, Result: result}

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
	case coap.MethodNotAllowed:
		return "MethodNotAllowed"
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
