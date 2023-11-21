package jokes

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
var _ = Describe("Service", func() {
	var (
		mockHTTPClient *httpclient.MockHTTPClient
		svc            *service
	)

	BeforeEach(func() {
		mockHTTPClient = &httpclient.MockHTTPClient{}
		svc = NewService("http://test.com", mockHTTPClient)
	})

	Describe("GetJoke", func() {
		Context("when the HTTP call is successful", func() {
			BeforeEach(func() {
				json := "{\"type\":\"success\",\"value\":{\"categories\":[\"nerdy\"],\"id\":1700536212,\"joke\":\"John Doe can't test for equality because he has no equal.\"}}"
				mockHTTPClient.Response = &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader([]byte(json))),
				}
			})

			It("should return a joke and no error", func() {
				res, err := svc.GetJoke("John", "Doe")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res).ShouldNot(BeNil())
				Expect(res.Value.Joke).Should(Equal("John Doe can't test for equality because he has no equal."))
			})
		})

		Context("when the HTTP call fails", func() {
			BeforeEach(func() {
				mockHTTPClient.Err = fmt.Errorf("network error")
			})

			It("should return an error", func() {
				_, err := svc.GetJoke("John", "Doe")
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})

func TestJokeService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Joke Service Suite")
}
