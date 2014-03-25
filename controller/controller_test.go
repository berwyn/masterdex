package controller

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

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
	})

	Describe("Parsing payloads", func() {
		It("should deserialise JSON data", func() {
			setPayload(response, jsonRequest)
			Expect(response.Payload).To(Equal(payload))
		})
	})
})
