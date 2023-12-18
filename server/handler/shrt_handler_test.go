package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"net/http/httptest"
	"shrt-server/handler"
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
	shrtService.EXPECT().CreateShrtLink(
		&request.CreateShortenLinkRequest{
			OriginalUrl: "https://www.google.com",
			Slug:        text.Ptr("dfdfdff"),
		}).Return(
		&response.CreateShortenLinkResponse{
			OriginalUrl: "https://www.google.com",
			Slug:        "dfdfdff",
		}, nil).AnyTimes()

	shrtService.EXPECT().CreateShrtLink(
		&request.CreateShortenLinkRequest{
			OriginalUrl: "https://www.google.com",
		}).Return(
		&response.CreateShortenLinkResponse{
			OriginalUrl: "https://www.google.com",
			Slug:        "dfdfdff",
		}, nil).AnyTimes()

	shrtService.EXPECT().CreateShrtLink(
		&request.CreateShortenLinkRequest{
			OriginalUrl: "https://www.google.com",
			Slug:        text.Ptr("dup_slug"),
		}).Return(nil, types.ErrSlugAlreadyExists).AnyTimes()

	shrtService.EXPECT().CreateShrtLink(
		&request.CreateShortenLinkRequest{
			OriginalUrl: "https://www.google.com",
			Slug:        text.Ptr("dfdfdff@"),
		}).Return(nil, types.ErrSlugNotAlphanumeric).AnyTimes()

	shrtHandlerTest := handler.NewShrtHandler(shrtService, payloadValidator)

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
		{
			name: "create link with duplicate slug",
			request: &request.CreateShortenLinkRequest{
				OriginalUrl: "https://www.google.com",
				Slug:        text.Ptr("dup_slug"),
			},
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: false,
				Message: text.Ptr(types.ErrSlugAlreadyExists.Error()),
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
				Message: text.Ptr(types.ErrSlugNotAlphanumeric.Error()),
			},
		},
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
	}

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

func Test_shrtHandler_GetOriginalURL(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	app := fiber.New()
	payloadValidator := validator.New()
	shrtService := mocks.NewMockShrtService(control)
	shrtService.EXPECT().GetOriginalURL("dfdfdff").Return(&response.CreateShortenLinkResponse{
		OriginalUrl: "https://www.google.com",
		Slug:        "dfdfdff",
	}, nil).AnyTimes()
	shrtService.EXPECT().GetOriginalURL("mixko").Return(nil, types.ErrSlugNotFound).AnyTimes()
	shrtService.EXPECT().GetOriginalURL("database_error_slug").Return(nil, types.ErrSomethingWentWrong).AnyTimes()

	shrtHandlerTest := handler.NewShrtHandler(shrtService, payloadValidator)

	testCases := []struct {
		name     string
		query    string
		expected types.Response[response.CreateShortenLinkResponse]
	}{
		{
			name:  "get original url with valid slug",
			query: "?slug=dfdfdff",
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: true,
				Payload: &response.CreateShortenLinkResponse{
					OriginalUrl: "https://www.google.com",
					Slug:        "dfdfdff",
				},
			},
		},
		{
			name:  "get original url with empty slug",
			query: "",
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: false,
				Message: text.Ptr(types.ErrSlugIsRequired.Error()),
			},
		},
		{
			name:  "get original url with slug not found",
			query: "?slug=mixko",
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: false,
				Message: text.Ptr(types.ErrSlugNotFound.Error()),
			},
		},
		{
			name:  "get original url with database error",
			query: "?slug=database_error_slug",
			expected: types.Response[response.CreateShortenLinkResponse]{
				Success: false,
				Message: text.Ptr(types.ErrSomethingWentWrong.Error()),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			app.Get("/retrieve", shrtHandlerTest.GetOriginalURL)

			req := httptest.NewRequest("GET", "/retrieve"+testCase.query, nil)

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

}

//
//func Test_shrtHandler_GetOriginalURLToRedirect(t *testing.T) {
//	type fields struct {
//		shrtService service.ShrtService
//		validator   *validator.Validate
//	}
//	type args struct {
//		c *fiber.Ctx
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := shrtHandler{
//				shrtService: tt.fields.shrtService,
//				validator:   tt.fields.validator,
//			}
//			if err := h.GetOriginalURLToRedirect(tt.args.c); (err != nil) != tt.wantErr {
//				t.Errorf("GetOriginalURLToRedirect() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
