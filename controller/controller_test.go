package controller

import (
	"bytes"
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

type MockDatastore struct{
	store map[string]interface{}
}

func (datastore MockDatastore) Find(id string) (interface{}, error) {
	if entity, ok := datastore.store[id]; ok {
		return entity, nil
	}
	return new(struct{}), errors.New("Doesn't exist")
}

func (datastore MockDatastore) Insert(entity interface{}) (interface{}, error) {
	if entity != nil {
		return entity, nil
	}
	return new(struct{}), errors.New("Nil insert")
}

func (datastore MockDatastore) Update(entity interface{}) (interface{}, error) {
	
}

var _ = Describe("Controller", func() {

	var (
		request     *http.Request
		jsonRequest *http.Request
		response    *Request
		payload     []byte
	)

	BeforeEach(func() {
		payload = []byte(`{"first": "value1", "second": "value2"}`)
		request, _ = http.NewRequest("GET", "localhost", nil)
		jsonRequest, _ = http.NewRequest("POST", "localhost", bytes.NewReader(payload))
		jsonRequest.Header.Add("Content-Type", mime_type_json)
		jsonRequest.Header.Add("Accept", mime_type_json)
		response = &Request{}
	})

	Describe("Recognising the accept type", func() {
		It("should recognise text/html", func() {
			setResponseType(response, request)
			Expect(response.ResponseType).To(Equal(mime_type_html))
		})

		It("should recognise application/json", func() {
			setResponseType(response, jsonRequest)
			Expect(response.ResponseType).To(Equal(mime_type_json))
		})

		It("should choose whichever content type has higher priority", func() {
			htmlRequest, _ := http.NewRequest("POST", "localhost", nil)
			jsonRequest, _ := http.NewRequest("POST", "localhost", nil)
			htmlRequest.Header.Add("Accept", fmt.Sprintf("%v,%v", mime_type_html, mime_type_json))
			jsonRequest.Header.Add("Accept", fmt.Sprintf("%v,%v", mime_type_json, mime_type_html))

			setResponseType(response, htmlRequest)
			Expect(response.ResponseType).To(Equal(mime_type_html))

			setResponseType(response, jsonRequest)
			Expect(response.ResponseType).To(Equal(mime_type_json))
		})
	})

	Describe("Parsing payloads", func() {
		It("should deserialise JSON data", func() {
			setPayload(response, jsonRequest)
			Expect(response.Payload).To(Equal(payload))
		})
	})
})
