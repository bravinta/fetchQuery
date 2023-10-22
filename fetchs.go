package fetchQuery

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Response struct {
	UnmarshalData interface{}
	Data          []byte
	Status        int
	Headers       http.Header
}

func (resp *Response) JSON(v *interface{}) (err error) {
	err = json.Unmarshal(resp.Data, &v)
	if err == nil {
		resp.UnmarshalData = v
	}

	return
}

func postX(url string, data map[string]interface{}, timeout time.Duration, headers http.Header, query ...url.Values) (resp *Response, err error) {
	config := &config{
		URL:     url,
		Method:  http.MethodPost,
		Body:    data,
		Timeout: timeout,
		Headers: headers,
	}

	if len(query) != 0 {
		config.Query = query[0]
	}
	return request(config)
}

func putX(url string, data map[string]interface{}, timeout time.Duration, headers http.Header, query ...url.Values) (resp *Response, err error) {
	config := &config{
		URL:     url,
		Method:  http.MethodPut,
		Body:    data,
		Headers: headers,
		Timeout: timeout,
	}

	if len(query) != 0 {
		config.Query = query[0]
	}
	return request(config)
}

func delX(url string, data map[string]interface{}, timeout time.Duration, headers http.Header, query ...url.Values) (resp *Response, err error) {
	config := &config{
		URL:     url,
		Method:  http.MethodDelete,
		Body:    data,
		Headers: headers,
		Timeout: timeout,
	}

	if len(query) != 0 {
		config.Query = query[0]
	}
	return request(config)
}

func getX(url string, timeout time.Duration, headers http.Header, query ...url.Values) (resp *Response, err error) {
	config := &config{
		URL:     url,
		Method:  http.MethodGet,
		Timeout: timeout,
		Headers: headers,
	}

	if len(query) != 0 {
		config.Query = query[0]
	}
	return request(config)
}

type config struct {
	URL     string
	Headers http.Header
	Method  string
	Body    map[string]interface{}
	Timeout time.Duration
	Query   url.Values
}

func request(c *config) (*Response, error) {
	if c.Method == "" {
		c.Method = http.MethodGet
	}
	resp, err := http.NewRequest(c.Method, c.URL, getRequestBody(c))
	if err != nil {
		// Manejar el error en caso de que ocurra
		return &Response{}, err
	}

	resp.Header = c.Headers

	// Realizar la solicitud HTTP
	client := &http.Client{
		Timeout: c.Timeout,
	}
	response, err := client.Do(resp)
	if err != nil {
		// Manejar el error de la solicitud
		return &Response{}, err
	}
	defer response.Body.Close()

	// Leer el cuerpo de la respuesta y convertirlo en []byte
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	dataResponse := &Response{
		Status:  response.StatusCode,
		Data:    bodyBytes,
		Headers: response.Header,
	}
	dataResponse.UnmarshalData = string(dataResponse.Data)
	return dataResponse, nil
}

func getRequestBody(c *config) (reader io.Reader) {
	if c.Body == nil {
		return
	}

	// Serializa el mapa a JSON
	jsonData, err := json.Marshal(c.Body)
	if err != nil {
		fmt.Println("Error al serializar el mapa a JSON:", err)
		return
	}

	// Crea un io.Reader a partir del JSON serializado
	reader = strings.NewReader(string(jsonData))

	return
}
