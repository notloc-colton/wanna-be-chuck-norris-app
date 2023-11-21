package randomName

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
	"wanna-be-chuck-norris-app/pkg/httpclient"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TODO: Add more tests for all use cases
// TODO: Add tests for cache functionality
var _ = Describe("Service", func() {
	var (
		mockHTTPClient *httpclient.MockHTTPClient
		svc            Service
	)

	BeforeEach(func() {
		mockHTTPClient = &httpclient.MockHTTPClient{}
		svc = NewService("http://test.com", mockHTTPClient, 0)
	})
	Describe("GetName", func() {
		Context("when the HTTP call is successful", func() {
			BeforeEach(func() {
				json := "{\"first_name\":\"Hasina\",\"last_name\":\"Tanweer\"}"
				mockHTTPClient.Response = &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader([]byte(json))),
				}
			})

			It("should return a name and no error", func() {
				name, err := svc.GetName()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(name).ShouldNot(BeNil())
				Expect(name.FirstName).Should(Equal("Hasina"))
				Expect(name.LastName).Should(Equal("Tanweer"))
			})
		})

		Context("when the HTTP call fails", func() {
			BeforeEach(func() {
				mockHTTPClient.Err = fmt.Errorf("network error")
			})

			It("should return an error", func() {
				_, err := svc.GetName()
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Random Name Suite")
}
