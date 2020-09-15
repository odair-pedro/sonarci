package http

import (
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
	Server  string
	Token   string
	Timeout time.Duration
}

func NewConnection(server string, token string, timeout time.Duration) *Connection {
	return &Connection{Server: server, Token: token, Timeout: timeout}
}

//
//func (connection *Connection) DoGet(route string) ([]byte, error) {
//	client := &http.Client{Timeout: connection.Timeout}
//
//	url := fmt.Sprintf("%s/%s", strings.TrimRight(connection.Server, "/"), strings.TrimLeft(route, "/"))
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		log.Fatalln(err)
//		return
//	}
//
//	req.Header.Add("Authorization", "Basic "+encodeToken(connection.Token))
//
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Panicln(err)
//		return
//		//return err
//	}
//
//	defer closeResource(resp.Body)
//
//	if !isStatusSuccess(resp.StatusCode) {
//		//log.Println("Failed request. Status Code: ", resp.Status)
//		//ch = make(chan error)
//		return errors.New("Failed request. Status Code: " + resp.Status)
//	}
//
//	buff, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatalln(err)
//		return
//	}
//
//	log.Println("Length: " + string(buff))
//	log.Println("begin return channel DoGet")
//	ch <- buff
//	log.Println("end return channel DoGet")
//}

func (connection *Connection) DoGet(route string) (<-chan []byte, <-chan error) {
	chOut := make(chan []byte, 1)
	chErr := make(chan error, 1)

	go func() {
		defer close(chOut)
		defer close(chErr)

		client := &http.Client{Timeout: connection.Timeout}

		url := fmt.Sprintf("%s/%s", strings.TrimRight(connection.Server, "/"), strings.TrimLeft(route, "/"))
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
			return
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

func closeResource(resource io.Closer) {
	err := resource.Close()
	if err != nil {
		log.Println("Failure to close HTTP resource: ", err.Error())
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
