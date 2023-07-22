package web

import (
	"backend/modules/api/endpoints/messages"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
)

const (
	ContentTypeOctetStream = "application/octet-stream"
)

func GenerateInternalServerError() error {
	return echo.NewHTTPError(http.StatusInternalServerError)
}

func GenerateUnauthorizedError() error {
	return echo.NewHTTPError(http.StatusUnauthorized)
}

func GenerateError(c echo.Context, httpStatus int, errorCode messages.ErrorCode) error {
	errorMessage := &messages.Wrapper{
		Message: &messages.Wrapper_ErrorResponse{
			ErrorResponse: &messages.Error{
				ErrorCode: errorCode,
			},
		},
	}
	content, err := proto.Marshal(errorMessage)
	if err != nil {
		c.Logger().Error("Error message marshaling failed: ", err)
		return GenerateInternalServerError()
	}
	return c.Blob(httpStatus, ContentTypeOctetStream, content)
}

func GenerateResponse(c echo.Context, response *messages.Wrapper) error {
	content, err := proto.Marshal(response)
	if err != nil {
		c.Logger().Error("Response message marshaling failed: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.Blob(http.StatusOK, ContentTypeOctetStream, content)
}
