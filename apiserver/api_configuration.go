/*
 * Booking App
 *
 * API to access and configure the Booking App
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ConfigurationAPIController binds http requests to an api service and writes the service results to the http response
type ConfigurationAPIController struct {
	service      ConfigurationAPIServicer
	errorHandler ErrorHandler
}

// ConfigurationAPIOption for how the controller is set up.
type ConfigurationAPIOption func(*ConfigurationAPIController)

// WithConfigurationAPIErrorHandler inject ErrorHandler into controller
func WithConfigurationAPIErrorHandler(h ErrorHandler) ConfigurationAPIOption {
	return func(c *ConfigurationAPIController) {
		c.errorHandler = h
	}
}

// NewConfigurationAPIController creates a default api controller
func NewConfigurationAPIController(s ConfigurationAPIServicer, opts ...ConfigurationAPIOption) Router {
	controller := &ConfigurationAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ConfigurationAPIController
func (c *ConfigurationAPIController) Routes() Routes {
	return Routes{
		"DeleteConfigurationById": Route{
			strings.ToUpper("Delete"),
			"/v1/configs/{config-id}",
			c.DeleteConfigurationById,
		},
		"GetConfigurationById": Route{
			strings.ToUpper("Get"),
			"/v1/configs/{config-id}",
			c.GetConfigurationById,
		},
		"GetConfigurations": Route{
			strings.ToUpper("Get"),
			"/v1/configs",
			c.GetConfigurations,
		},
		"PostConfiguration": Route{
			strings.ToUpper("Post"),
			"/v1/configs",
			c.PostConfiguration,
		},
		"PutConfigurationById": Route{
			strings.ToUpper("Put"),
			"/v1/configs/{config-id}",
			c.PutConfigurationById,
		},
	}
}

// DeleteConfigurationById - Deletes a configuration
func (c *ConfigurationAPIController) DeleteConfigurationById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	configIdParam, err := parseNumericParameter[int64](
		params["config-id"],
		WithRequire[int64](parseInt64),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteConfigurationById(r.Context(), configIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetConfigurationById - Get configuration
func (c *ConfigurationAPIController) GetConfigurationById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	configIdParam, err := parseNumericParameter[int64](
		params["config-id"],
		WithRequire[int64](parseInt64),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetConfigurationById(r.Context(), configIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetConfigurations - Get configurations
func (c *ConfigurationAPIController) GetConfigurations(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetConfigurations(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// PostConfiguration - Creates a configuration
func (c *ConfigurationAPIController) PostConfiguration(w http.ResponseWriter, r *http.Request) {
	configurationParam := Configuration{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&configurationParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertConfigurationRequired(configurationParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertConfigurationConstraints(configurationParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostConfiguration(r.Context(), configurationParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// PutConfigurationById - Updates a configuration
func (c *ConfigurationAPIController) PutConfigurationById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	configIdParam, err := parseNumericParameter[int64](
		params["config-id"],
		WithRequire[int64](parseInt64),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	configurationParam := Configuration{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&configurationParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertConfigurationRequired(configurationParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertConfigurationConstraints(configurationParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PutConfigurationById(r.Context(), configIdParam, configurationParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
