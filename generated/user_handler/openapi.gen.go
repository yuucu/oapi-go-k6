// Package user_handler provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package user_handler

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// User defines model for User.
type User struct {
	IconPath string `json:"icon_path"`

	// Unique identifier for the given user.
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// PostUserJSONBody defines parameters for PostUser.
type PostUserJSONBody struct {
	IconBody *string `json:"icon_body,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// PostUserJSONRequestBody defines body for PostUser for application/json ContentType.
type PostUserJSONRequestBody PostUserJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /user)
	PostUser(ctx echo.Context) error

	// (DELETE /user/{id})
	DeleteUser(ctx echo.Context, id int) error
	// Your GET endpoint
	// (GET /user/{id})
	GetUser(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostUser converts echo context to params.
func (w *ServerInterfaceWrapper) PostUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostUser(ctx)
	return err
}

// DeleteUser converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUser(ctx, id)
	return err
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUser(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/user", wrapper.PostUser)
	router.DELETE(baseURL+"/user/:id", wrapper.DeleteUser)
	router.GET(baseURL+"/user/:id", wrapper.GetUser)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RU0WrbShD9FTP3Pq4t2Q7cm326N6SEUGgKbQolhLKRRtIEaXezOzJxjV6af+kn9Kl/",
	"kx8ps5bjtk6hb32SmJ05c+acnd1A4TrvLFqOoDcQiwY7k34vIwb54r3pfIsR9NUGSsN4UZ1Q4AY0zI+P",
	"/5nO8+lyDgqwM9SCBtNSgbPYETf/1RKbFa7bnb/DQBVhCZpDjwoqCpFfmQ5Bw/9SCAqoBD0/WihozdPZ",
	"G4EDBZFqe+lPDUtwkc+Pp/m/08URDNcKfHAeAxMm+lQ4+8Eb4bkBXnspiBzI1jBse2ygxFgE8kzOgoZL",
	"S3c9TqhEy0IyTCoXJtzgpKYV2kkfMcxA7cDIMtYYBM0mkgdtBgUB73oKMu+V9BxT1XfsrhUwcSt1SfEn",
	"fHdziwXDIDBkK5cajJnOozWeQMEKQ9zSn89y4bI70rCcSUiBtEmSZP1oqXeR5St6GZn+vAQNr13kkYLQ",
	"xsgnrlxLXuEso00lxvuWilSU3UZpvLs0CfjQgpsRo3KhMwwabsiasN7PuTfl1zIeSJKkjd7ZuG21yPND",
	"Qy9ejqlp8GxD5bBNapHxcP7TFH9S4BD9t2X4O2AFGv7K9suVjZuVJfzE6jmyCmp8xpoz5D/KS0Hsu05s",
	"0/De9WFy9uLtBG3pHVkW67wJpkPGMD4TP0A8Pnx+fPj6+OnL+alcfQmlq79zfLsZ+1XZPg172j8v3HCd",
	"GGFY7fr1QV6ehtnrLGtdYdrGRdbLPM/T03A/jex8S3WTVJLth/Vqkeetb8xyXnyEYfgWAAD//5EWuxIL",
	"BQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
