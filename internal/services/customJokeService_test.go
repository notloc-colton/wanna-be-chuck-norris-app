package services

import (
	"context"
	"errors"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"wanna-be-chuck-norris-app/internal/vendors/jokes"
	"wanna-be-chuck-norris-app/internal/vendors/randomName"
)

type MockRandomNameService struct {
	NameResponse randomName.NameResponse
	Err          error
}

func (m *MockRandomNameService) GetName() (*randomName.NameResponse, error) {
	return &m.NameResponse, m.Err
}

type MockJokeService struct {
	JokeResponse jokes.JokeResponse
	Err          error
}

func (m *MockJokeService) GetJoke(firstName, lastName string) (*jokes.JokeResponse, error) {
	return &m.JokeResponse, m.Err
}

var _ = Describe("CustomJokeService", func() {
	var (
		mockRandomNameService *MockRandomNameService
		mockJokeService       *MockJokeService
		customJokeService     *customJokeService
	)

	BeforeEach(func() {
		mockRandomNameService = &MockRandomNameService{}
		mockJokeService = &MockJokeService{}
		customJokeService = NewCustomJokeService(mockRandomNameService, mockJokeService)
	})

	Describe("GetJoke", func() {
		Context("when both name and joke retrieval are successful", func() {
			It("should return a custom joke", func() {
				mockRandomNameService.NameResponse = randomName.NameResponse{FirstName: "John", LastName: "Doe"}
				mockJokeService.JokeResponse.Value.Joke = "Funny Joke"
				joke, err := customJokeService.GetJoke(context.Background())
				Expect(err).NotTo(HaveOccurred())
				Expect(joke).To(Equal("Funny Joke"))
			})
		})

		Context("when name retrieval fails", func() {
			It("should return an error", func() {
				mockRandomNameService.Err = errors.New("name service error")
				_, err := customJokeService.GetJoke(context.Background())
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when joke retrieval fails", func() {
			It("should return an error", func() {
				mockRandomNameService.NameResponse = randomName.NameResponse{FirstName: "John", LastName: "Doe"}
				mockJokeService.Err = errors.New("joke service error")
				_, err := customJokeService.GetJoke(context.Background())
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Custom Joke Service Suite")
}
