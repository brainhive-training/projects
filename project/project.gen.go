// Package project provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package project

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get projects
	// (GET /api/projects)
	ListProjects(c *gin.Context)
	// Create a new project
	// (POST /api/projects)
	CreateProject(c *gin.Context)
	// Get project by pCode
	// (GET /api/projects/{code})
	GetProject(c *gin.Context, code string)
	// Update project
	// (PATCH /api/projects/{code})
	UpdateProject(c *gin.Context, code string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// ListProjects operation middleware
func (siw *ServerInterfaceWrapper) ListProjects(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListProjects(c)
}

// CreateProject operation middleware
func (siw *ServerInterfaceWrapper) CreateProject(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateProject(c)
}

// GetProject operation middleware
func (siw *ServerInterfaceWrapper) GetProject(c *gin.Context) {

	var err error

	// ------------- Path parameter "code" -------------
	var code string

	err = runtime.BindStyledParameter("simple", false, "code", c.Param("code"), &code)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter code: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProject(c, code)
}

// UpdateProject operation middleware
func (siw *ServerInterfaceWrapper) UpdateProject(c *gin.Context) {

	var err error

	// ------------- Path parameter "code" -------------
	var code string

	err = runtime.BindStyledParameter("simple", false, "code", c.Param("code"), &code)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter code: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateProject(c, code)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/api/projects", wrapper.ListProjects)
	router.POST(options.BaseURL+"/api/projects", wrapper.CreateProject)
	router.GET(options.BaseURL+"/api/projects/:code", wrapper.GetProject)
	router.PATCH(options.BaseURL+"/api/projects/:code", wrapper.UpdateProject)
}
