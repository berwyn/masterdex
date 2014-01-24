package controller

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Controller", func() {

	var (
		request     *http.Request
		jsonRequest *http.Request
	)

	BeforeEach(func() {
		request, _ = http.NewRequest("GET", "localhost", nil)
		jsonRequest, _ = http.NewRequest("GET", "localhost", nil)
		jsonRequest.Header.Add("Content-Type", "application/json")
		jsonRequest.Header.Add("Accept", "application/json")
	})

	Describe("Checking for JSON reqeusts", func() {
		Context("HTML requests", func() {
			It("should reject HTML Content-Types", func() {
				Expect(hasJSON(request)).To(Equal(false))
			})

			It("should reject HTML Accept headers", func() {
				Expect(useJSON(request)).To(Equal(false))
			})
		})

		Context("JSON requests", func() {
			It("should accept JSON Content-Types", func() {
				Expect(hasJSON(jsonRequest)).To(Equal(true))
			})

			It("should accept JSON Accept headers", func() {
				Expect(useJSON(jsonRequest)).To(Equal(true))
			})
		})
	})

})
