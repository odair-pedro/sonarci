package http

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Connection struct {
	HostServer string
	Token      string
	Timeout    time.Duration
}

func NewConnection(hostServer string, token string, timeout time.Duration) *Connection {
	return &Connection{HostServer: hostServer, Token: token, Timeout: timeout}
}

func (connection *Connection) GetHostServer() string {
	return connection.HostServer
}

func (connection *Connection) Delete(endpoint string) <-chan error {
	chErr := make(chan error)

	go func() {
		defer close(chErr)

		client := &http.Client{Timeout: connection.Timeout}

		url := parseUrl(connection.GetHostServer(), endpoint)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			chErr <- err
			return
		}

		req.Header.Add("Authorization", "Basic "+encodeToken(connection.Token))
		resp, err := client.Do(req)
		if err != nil {
			chErr <- err
			return
		}

		defer closeResource(resp.Body)
		if !isStatusSuccess(resp.StatusCode) {
			chErr <- errors.New("Failed request. Status Code: " + resp.Status)
			return
		}
	}()

	return chErr
}

func (connection *Connection) Get(endpoint string) (<-chan []byte, <-chan error) {
	chOut := make(chan []byte, 1)
	chErr := make(chan error, 1)

	go func() {
		defer close(chOut)
		defer close(chErr)

		client := &http.Client{Timeout: connection.Timeout}

		url := parseUrl(connection.GetHostServer(), endpoint)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			chErr <- err
			return
		}

		req.Header.Add("Authorization", "Basic "+encodeToken(connection.Token))
		resp, err := client.Do(req)
		if err != nil {
			chErr <- err
			return
		}

		defer closeResource(resp.Body)
		if !isStatusSuccess(resp.StatusCode) {
			chErr <- errors.New("Failed request. Status Code: " + resp.Status)
		}

		buff, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			chErr <- err
			return
		}

		chOut <- buff
	}()

	return chOut, chErr
}

func (connection *Connection) Post(endpoint string, content []byte, contentType string) (<-chan []byte, <-chan error) {
	chOut := make(chan []byte, 1)
	chErr := make(chan error, 1)

	go func() {
		defer close(chOut)
		defer close(chErr)

		client := &http.Client{Timeout: connection.Timeout}

		url := parseUrl(connection.GetHostServer(), endpoint)
		req, err := http.NewRequest("POST", url, bytes.NewReader(content))
		if err != nil {
			chErr <- err
			return
		}

		req.Header.Add("Authorization", "Basic "+encodeToken(connection.Token))
		req.Header.Add("Content-Type", contentType)
		resp, err := client.Do(req)
		if err != nil {
			chErr <- err
			return
		}

		defer closeResource(resp.Body)
		if !isStatusSuccess(resp.StatusCode) {
			chErr <- errors.New("Failed request. Status Code: " + resp.Status)
		}

		buff, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			chErr <- err
			return
		}

		chOut <- buff
	}()

	return chOut, chErr
}

func parseUrl(host string, endpoint string) string {
	return fmt.Sprintf("%s/%s", strings.TrimRight(host, "/"), strings.TrimLeft(endpoint, "/"))
}

func closeResource(resource io.Closer) {
	err := resource.Close()
	if err != nil {
		log.Panic("Failure to close HTTP resource: ", err.Error())
	}
}

func encodeToken(token string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", token)))
}

func isStatusSuccess(statusCode int) bool {
	const (
		ok                 = 200
		created            = 201
		accepted           = 202
		partialInformation = 203
		noResponse         = 204
	)

	return statusCode == ok || statusCode == created || statusCode == accepted ||
		statusCode == partialInformation || statusCode == noResponse

}
