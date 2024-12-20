// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// APIVersion Information about a version of the API.
type APIVersion struct {
	Version *string `json:"version,omitempty"`
}

// APIVersions Information about a list of versions of the API.
type APIVersions struct {
	ApiVersions *[]APIVersion `json:"apiVersions,omitempty"`
	UriPrefix   *string       `json:"uriPrefix,omitempty"`
}

// ProblemDetails defines model for ProblemDetails.
type ProblemDetails struct {
	// AdditionalAttributes Any number of additional attributes, as defined in a specification or by an implementation.
	AdditionalAttributes *map[string]string `json:"additionalAttributes,omitempty"`

	// Detail A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail"`

	// Instance A URI reference that identifies the specific occurrence of the problem.
	// It may yield further information if dereferenced.
	Instance *string `json:"instance,omitempty"`

	// Status The HTTP status code for this occurrence of the problem.
	Status int `json:"status"`

	// Title A short, human-readable summary of the problem type. It should not change from occurrence to occurrence of the problem,
	// except for purposes of localization. If type is given and other than "about:blank", this attribute shall also be provided.
	Title *string `json:"title,omitempty"`

	// Type A URI reference according to IETF RFC 3986 [3] that identifies the problem type. It is encouraged that the URI provides
	// human-readable documentation for the problem (e.g. usingHTML) when dereferenced. When this member is not present,
	// its value is assumed to be "about:blank".
	Type *string `json:"type,omitempty"`
}

// ExcludeFields defines model for excludeFields.
type ExcludeFields = string

// Fields defines model for fields.
type Fields = string

// Filter defines model for filter.
type Filter = string

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
