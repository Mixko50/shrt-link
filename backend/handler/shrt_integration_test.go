package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"io"
	"net/http/httptest"
	"shrt-server/handler"
	"shrt-server/service"
	"shrt-server/test/mocks"
	"shrt-server/types"
	"shrt-server/types/entity"
	"shrt-server/types/request"
	"shrt-server/types/response"
	"shrt-server/utility/configuration"
	"shrt-server/utility/text"
	"testing"
)

func TestCreateShrtLinkIntegration(t *testing.T) {
	// Arrange
	control := gomock.NewController(t)
	defer control.Finish()

	app := fiber.New()
	payloadValidator := validator.New()
	shrtRepo := mocks.NewMockShrtRepository(control)
	utlity := mocks.NewMockUtility(control)
	shrtService := service.NewShrtService(shrtRepo, utlity)
	shrtHandlerTest := handler.NewShrtHandler(shrtService, payloadValidator, nil)

	type cases struct {
		name          string
		mockSlug      *string
		request       request.CreateShortenLinkRequest
		dbCreatedMock *entity.Shrt
		expected      types.Response[response.CreateShortenLinkResponse]
	}

	testCases := []cases{
		{
			name:     "create link with slug",
			mockSlug: nil,
			request: request.CreateShortenLinkRequest{
				OriginalUrl: "https://www.google.com",
				Slug:        text.Ptr("google_test"),
			},
			dbCreatedMock: &entity.Shrt{
				LongURL: "https://www.google.com",
				Slug:    "google_test",
			},
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: true,
				Payload: &response.CreateShortenLinkResponse{
					OriginalUrl: "https://www.google.com",
					Slug:        "google_test",
				},
			},
		},
		{
			name:     "create link without slug",
			mockSlug: text.Ptr("google_test"),
			request: request.CreateShortenLinkRequest{
				OriginalUrl: "https://www.google.com",
			},
			dbCreatedMock: &entity.Shrt{
				LongURL: "https://www.google.com",
				Slug:    "google_test",
			},
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: true,
				Payload: &response.CreateShortenLinkResponse{
					OriginalUrl: "https://www.google.com",
					Slug:        "google_test",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			app.Post("/create", shrtHandlerTest.CreateShrtLink)

			if testCase.request.Slug != nil {
				shrtRepo.EXPECT().FindBySlug(gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
				shrtRepo.EXPECT().Create(&entity.Shrt{
					LongURL: testCase.request.OriginalUrl,
					Slug:    *testCase.request.Slug,
				}).Return(nil)
			} else {
				utlity.EXPECT().GenerateSlug().Return(*testCase.mockSlug)
				shrtRepo.EXPECT().Create(&entity.Shrt{
					LongURL: testCase.request.OriginalUrl,
					Slug:    *testCase.mockSlug,
				}).Return(nil)
			}

			var buf bytes.Buffer
			_ = json.NewEncoder(&buf).Encode(testCase.request)

			req := httptest.NewRequest("POST", "/create", &buf)
			req.Header.Set("Content-Type", "application/json")

			// Act
			res, _ := app.Test(req)
			defer res.Body.Close()

			// Assert
			body, _ := io.ReadAll(res.Body)
			var actual types.Response[response.CreateShortenLinkResponse]
			_ = json.Unmarshal(body, &actual)
			assert.Equal(t, testCase.expected, actual)
		})
	}

	t.Run("create link with duplicated slug", func(t *testing.T) {
		// Arrange
		app.Post("/create", shrtHandlerTest.CreateShrtLink)

		shrtRepo.EXPECT().FindBySlug(gomock.Any()).Return(&entity.Shrt{
			ID:      1,
			LongURL: "https://www.google.com",
			Slug:    "google_test",
		}, nil)

		request := request.CreateShortenLinkRequest{
			OriginalUrl: "https://www.google.com",
			Slug:        text.Ptr("google_test"),
		}

		expected := types.Response[response.CreateShortenLinkResponse]{
			Success: false,
			Message: text.Ptr("slug already exists"),
		}

		var buf bytes.Buffer
		_ = json.NewEncoder(&buf).Encode(request)

		req := httptest.NewRequest("POST", "/create", &buf)
		req.Header.Set("Content-Type", "application/json")

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		body, _ := io.ReadAll(res.Body)
		var actual types.Response[response.CreateShortenLinkResponse]
		_ = json.Unmarshal(body, &actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("create link with invalid url", func(t *testing.T) {
		// Arrange
		app.Post("/create", shrtHandlerTest.CreateShrtLink)

		request := request.CreateShortenLinkRequest{
			OriginalUrl: "https//www.google",
		}

		expected := types.Response[response.CreateShortenLinkResponse]{
			Success: false,
			Message: text.Ptr("OriginalUrl is invalid"),
		}

		var buf bytes.Buffer
		_ = json.NewEncoder(&buf).Encode(request)

		req := httptest.NewRequest("POST", "/create", &buf)
		req.Header.Set("Content-Type", "application/json")

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		body, _ := io.ReadAll(res.Body)
		var actual types.Response[response.CreateShortenLinkResponse]
		_ = json.Unmarshal(body, &actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("create link with invalid slug", func(t *testing.T) {
		// Arrange
		app.Post("/create", shrtHandlerTest.CreateShrtLink)

		request := request.CreateShortenLinkRequest{
			OriginalUrl: "https://www.google.com",
			Slug:        text.Ptr("google_test@"),
		}

		expected := types.Response[response.CreateShortenLinkResponse]{
			Success: false,
			Message: text.Ptr("slug must be alphanumeric"),
		}

		var buf bytes.Buffer
		_ = json.NewEncoder(&buf).Encode(request)

		req := httptest.NewRequest("POST", "/create", &buf)
		req.Header.Set("Content-Type", "application/json")

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		body, _ := io.ReadAll(res.Body)
		var actual types.Response[response.CreateShortenLinkResponse]
		_ = json.Unmarshal(body, &actual)
		assert.Equal(t, expected, actual)
	})
}

func TestGetOriginalURLIntegration(t *testing.T) {
	// Arrange
	control := gomock.NewController(t)
	defer control.Finish()

	app := fiber.New()
	payloadValidator := validator.New()
	shrtRepo := mocks.NewMockShrtRepository(control)
	utlity := mocks.NewMockUtility(control)
	shrtService := service.NewShrtService(shrtRepo, utlity)
	shrtHandlerTest := handler.NewShrtHandler(shrtService, payloadValidator, nil)

	type cases struct {
		name     string
		mockSlug string
		expected types.Response[response.CreateShortenLinkResponse]
	}

	testCases := []cases{
		{
			name:     "get original url",
			mockSlug: "google_test",
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: true,
				Payload: &response.CreateShortenLinkResponse{
					OriginalUrl: "https://www.google.com",
					Slug:        "google_test",
				},
			},
		},
		{
			name:     "get original url with not found slug",
			mockSlug: "google_test",
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: false,
				Message: text.Ptr("slug not found"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			app.Get("/api/v1/retrieve", shrtHandlerTest.GetOriginalURL)

			if testCase.expected.Success {
				shrtRepo.EXPECT().FindBySlug(testCase.mockSlug).Return(&entity.Shrt{
					ID:      1,
					LongURL: "https://www.google.com",
					Slug:    testCase.mockSlug,
				}, nil)
			} else {
				shrtRepo.EXPECT().FindBySlug(testCase.mockSlug).Return(nil, gorm.ErrRecordNotFound)
			}

			req := httptest.NewRequest("GET", "/api/v1/retrieve?slug="+testCase.mockSlug, nil)

			// Act
			res, _ := app.Test(req)
			defer res.Body.Close()

			// Assert
			body, _ := io.ReadAll(res.Body)
			var actual types.Response[response.CreateShortenLinkResponse]
			_ = json.Unmarshal(body, &actual)
			assert.Equal(t, testCase.expected, actual)
		})
	}

	t.Run("get original url with database error", func(t *testing.T) {
		// Arrange
		app.Get("/api/v1/retrieve", shrtHandlerTest.GetOriginalURL)

		mockSlug := "google_test"
		shrtRepo.EXPECT().FindBySlug(mockSlug).Return(nil, gorm.ErrInvalidDB)

		expected := types.Response[response.CreateShortenLinkResponse]{
			Success: false,
			Message: text.Ptr("something went wrong while processing your request"),
		}

		req := httptest.NewRequest("GET", "/api/v1/retrieve?slug="+mockSlug, nil)

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		body, _ := io.ReadAll(res.Body)
		var actual types.Response[response.CreateShortenLinkResponse]
		_ = json.Unmarshal(body, &actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("get original url with empty slug", func(t *testing.T) {
		// Arrange
		app.Get("/api/v1/retrieve", shrtHandlerTest.GetOriginalURL)

		expected := types.Response[response.CreateShortenLinkResponse]{
			Success: false,
			Message: text.Ptr("slug is required"),
		}

		req := httptest.NewRequest("GET", "/api/v1/retrieve", nil)

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		body, _ := io.ReadAll(res.Body)
		var actual types.Response[response.CreateShortenLinkResponse]
		_ = json.Unmarshal(body, &actual)
		assert.Equal(t, expected, actual)
	})
}

func TestGetOriginalURLToRedirectIntegration(t *testing.T) {
	// Arrange
	control := gomock.NewController(t)
	defer control.Finish()

	app := fiber.New()
	payloadValidator := validator.New()
	shrtRepo := mocks.NewMockShrtRepository(control)
	utlity := mocks.NewMockUtility(control)
	shrtService := service.NewShrtService(shrtRepo, utlity)
	config := configuration.Config{
		BaseUrl: "https://m.mixkomii.com",
	}
	shrtHandlerTest := handler.NewShrtHandler(shrtService, payloadValidator, &config)

	type cases struct {
		name         string
		mockSlug     string
		expectedLink string
	}

	testCases := []cases{
		{
			name:         "get original url",
			mockSlug:     "google_test",
			expectedLink: "https://www.google.com",
		},
		{
			name:         "get original url with not found slug",
			mockSlug:     "google_test_2",
			expectedLink: "https://m.mixkomii.com",
		},
		{
			name:         "get original url with database error",
			mockSlug:     "database_error_slug",
			expectedLink: "https://m.mixkomii.com",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			app.Get("/:slug", shrtHandlerTest.GetOriginalURLToRedirect)

			req := httptest.NewRequest("GET", "/"+testCase.mockSlug, nil)

			if testCase.mockSlug == "google_test" {
				shrtRepo.EXPECT().FindBySlug(testCase.mockSlug).Return(&entity.Shrt{
					ID:      1,
					LongURL: "https://www.google.com",
					Slug:    testCase.mockSlug,
				}, nil)
				shrtRepo.EXPECT().UpdateVisit(uint(1)).Return(nil)
			} else if testCase.mockSlug == "google_test_2" {
				shrtRepo.EXPECT().FindBySlug(testCase.mockSlug).Return(nil, gorm.ErrRecordNotFound)
			} else if testCase.mockSlug == "database_error_slug" {
				shrtRepo.EXPECT().FindBySlug(testCase.mockSlug).Return(nil, gorm.ErrInvalidDB)
			}

			// Act
			res, _ := app.Test(req)

			// Assert
			assert.Equal(t, testCase.expectedLink, res.Header.Get("Location"))
		})
	}
}
