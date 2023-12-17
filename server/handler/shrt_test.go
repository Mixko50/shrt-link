package handler

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"net/http/httptest"
	"shrt-server/service"
	"shrt-server/test/mocks"
	"shrt-server/types"
	"shrt-server/types/request"
	"shrt-server/types/response"
	"shrt-server/utilities/text"
	"testing"
)

func Test_shrtHandler_CreateShrtLink(t *testing.T) {
	// Arrange
	control := gomock.NewController(t)
	defer control.Finish()

	app := fiber.New()
	payloadValidator := validator.New()
	shrtService := mocks.NewMockShrtService(control)
	shrtService.EXPECT().CreateShrtLink(gomock.Any()).Return(&response.CreateShortenLinkResponse{
		OriginalUrl: "https://www.google.com",
		Slug:        "dfdfdff",
	}, nil).AnyTimes()

	shrtHandlerTest := NewShrtHandler(shrtService, payloadValidator)

	testCases := []struct {
		name     string
		request  *request.CreateShortenLinkRequest
		expected types.Response[response.CreateShortenLinkResponse]
	}{
		{
			name: "create link with valid slug",
			request: &request.CreateShortenLinkRequest{
				OriginalUrl: "https://www.google.com",
				Slug:        text.Ptr("dfdfdff"),
			},
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: true,
				Payload: &response.CreateShortenLinkResponse{
					OriginalUrl: "https://www.google.com",
					Slug:        "dfdfdff",
				},
			},
		},
		{
			name: "create link with empty slug",
			request: &request.CreateShortenLinkRequest{
				OriginalUrl: "https://www.google.com",
			},
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: true,
				Payload: &response.CreateShortenLinkResponse{
					OriginalUrl: "https://www.google.com",
					Slug:        "dfdfdff",
				},
			},
		},
		{
			name: "create link with invalid slug",
			request: &request.CreateShortenLinkRequest{
				OriginalUrl: "https://www.google.com",
				Slug:        text.Ptr("dfdfdff@"),
			},
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: false,
				Message: text.Ptr("Slug should be alphanum"),
			},
		},
		{
			name: "create link with invalid url",
			request: &request.CreateShortenLinkRequest{
				OriginalUrl: "https//www.google.com",
				Slug:        text.Ptr("dfdfdff"),
			},
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: false,
				Message: text.Ptr("OriginalUrl should be url"),
			},
		},
		{
			name: "create link with empty url",
			request: &request.CreateShortenLinkRequest{
				Slug: text.Ptr("dfdfdff"),
			},
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: false,
				Message: text.Ptr("OriginalUrl is required"),
			},
		},
		// Add more test cases as needed
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			app.Post("/create", shrtHandlerTest.CreateShrtLink)

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

		t.Run("create link with empty body", func(t *testing.T) {
			// Arrange
			app.Post("/create", shrtHandlerTest.CreateShrtLink)

			req := httptest.NewRequest("POST", "/create", nil)
			req.Header.Set("Content-Type", "application/json")

			expected := types.Response[response.CreateShortenLinkResponse]{
				Success: false,
				Message: text.Ptr("no data provided"),
			}

			// Act
			res, _ := app.Test(req)
			defer res.Body.Close()

			// Assert
			body, _ := io.ReadAll(res.Body)
			var actual types.Response[response.CreateShortenLinkResponse]
			_ = json.Unmarshal(body, &actual)
			assert.Equal(t, expected, actual)
		})

		t.Run("create link with invalid body", func(t *testing.T) {
			// Arrange
			app.Post("/create", shrtHandlerTest.CreateShrtLink)

			type invalidBody struct {
				OriginalUrl int    `json:"original_url"`
				Slug        string `json:"slug"`
			}

			var buf bytes.Buffer
			_ = json.NewEncoder(&buf).Encode(invalidBody{
				OriginalUrl: 1,
				Slug:        "dfdfdff",
			})

			expected := types.Response[response.CreateShortenLinkResponse]{
				Success: false,
				Message: text.Ptr("original_url should be string"),
			}

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
}

func Test_shrtHandler_GetOriginalURL(t *testing.T) {
	type fields struct {
		shrtService service.ShrtService
		validator   *validator.Validate
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := shrtHandler{
				shrtService: tt.fields.shrtService,
				validator:   tt.fields.validator,
			}
			if err := h.GetOriginalURL(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetOriginalURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_shrtHandler_GetOriginalURLToRedirect(t *testing.T) {
	type fields struct {
		shrtService service.ShrtService
		validator   *validator.Validate
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := shrtHandler{
				shrtService: tt.fields.shrtService,
				validator:   tt.fields.validator,
			}
			if err := h.GetOriginalURLToRedirect(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetOriginalURLToRedirect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
