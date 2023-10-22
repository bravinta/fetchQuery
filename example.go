package fetchQuery

import (
	"fmt"
	"net/http"
	"time"
)

func exampleQueryPOST() {
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	test := NewInstance("http://127.0.0.1:8080", headers, 5*time.Second)
	resp, err5 := test.Post("hosts", map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	})
	if err5 != nil {
		panic(err5)
	}

	fmt.Println(resp)
}

func ExampleQueryGET() {
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	test := NewInstance("http://127.0.0.1:8080", headers, 5*time.Second)

	resp, err5 := test.Get("hosts")
	if err5 != nil {
		panic(err5)
	}

	fmt.Println(string(resp.Data))
	fmt.Println(resp.UnmarshalData)
}
