// Package Zerotier provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.3 DO NOT EDIT.
package zerotier

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get Controller Status.
	// (GET /controller)
	GetControllerStatus(ctx echo.Context) error
	// List Networks.
	// (GET /controller/network)
	GetControllerNetworks(ctx echo.Context) error
	// Generate Random Network ID.
	// (POST /controller/network/{controllerID}______)
	GenerateControllerNetwork(ctx echo.Context, controllerID string) error
	// Delete a network.
	// (DELETE /controller/network/{networkID})
	DeleteControllerNetwork(ctx echo.Context, networkID string) error
	// Get Network by ID.
	// (GET /controller/network/{networkID})
	GetControllerNetwork(ctx echo.Context, networkID string) error
	// Create or Update a Network.
	// (POST /controller/network/{networkID})
	SetControllerNetwork(ctx echo.Context, networkID string) error
	// List Network Members.
	// (GET /controller/network/{networkID}/member)
	GetControllerNetworkMembers(ctx echo.Context, networkID string) error
	// Remove a network member.
	// (DELETE /controller/network/{networkID}/member/{nodeID})
	DeleteControllerNetworkMember(ctx echo.Context, networkID string, nodeID string) error
	// Get Network Member Details by ID.
	// (GET /controller/network/{networkID}/member/{nodeID})
	GetControllerNetworkMember(ctx echo.Context, networkID string, nodeID string) error
	// Create or Update a Network Membership.
	// (POST /controller/network/{networkID}/member/{nodeID})
	SetControllerNetworkMember(ctx echo.Context, networkID string, nodeID string) error
	// Get all network memberships.
	// (GET /network)
	GetNetworks(ctx echo.Context) error
	// Leave a network.
	// (DELETE /network/{networkID})
	DeleteNetwork(ctx echo.Context, networkID string) error
	// Get a joined Network membership configuration by Network ID.
	// (GET /network/{networkID})
	GetNetwork(ctx echo.Context, networkID string) error
	// Join a network or update it's configuration by Network ID.
	// (POST /network/{networkID})
	UpdateNetwork(ctx echo.Context, networkID string) error
	// Get all peers.
	// (GET /peer)
	GetPeers(ctx echo.Context) error
	// Get information about a specific peer by Node ID.
	// (GET /peer/{address})
	GetPeer(ctx echo.Context, address string) error
	// Node status and addressing info.
	// (GET /status)
	GetStatus(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetControllerStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetControllerStatus(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetControllerStatus(ctx)
	return err
}

// GetControllerNetworks converts echo context to params.
func (w *ServerInterfaceWrapper) GetControllerNetworks(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetControllerNetworks(ctx)
	return err
}

// GenerateControllerNetwork converts echo context to params.
func (w *ServerInterfaceWrapper) GenerateControllerNetwork(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "controllerID" -------------
	var controllerID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "controllerID", runtime.ParamLocationPath, ctx.Param("controllerID"), &controllerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter controllerID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GenerateControllerNetwork(ctx, controllerID)
	return err
}

// DeleteControllerNetwork converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteControllerNetwork(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "networkID" -------------
	var networkID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "networkID", runtime.ParamLocationPath, ctx.Param("networkID"), &networkID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter networkID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteControllerNetwork(ctx, networkID)
	return err
}

// GetControllerNetwork converts echo context to params.
func (w *ServerInterfaceWrapper) GetControllerNetwork(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "networkID" -------------
	var networkID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "networkID", runtime.ParamLocationPath, ctx.Param("networkID"), &networkID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter networkID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetControllerNetwork(ctx, networkID)
	return err
}

// SetControllerNetwork converts echo context to params.
func (w *ServerInterfaceWrapper) SetControllerNetwork(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "networkID" -------------
	var networkID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "networkID", runtime.ParamLocationPath, ctx.Param("networkID"), &networkID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter networkID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SetControllerNetwork(ctx, networkID)
	return err
}

// GetControllerNetworkMembers converts echo context to params.
func (w *ServerInterfaceWrapper) GetControllerNetworkMembers(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "networkID" -------------
	var networkID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "networkID", runtime.ParamLocationPath, ctx.Param("networkID"), &networkID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter networkID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetControllerNetworkMembers(ctx, networkID)
	return err
}

// DeleteControllerNetworkMember converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteControllerNetworkMember(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "networkID" -------------
	var networkID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "networkID", runtime.ParamLocationPath, ctx.Param("networkID"), &networkID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter networkID: %s", err))
	}

	// ------------- Path parameter "nodeID" -------------
	var nodeID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "nodeID", runtime.ParamLocationPath, ctx.Param("nodeID"), &nodeID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter nodeID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteControllerNetworkMember(ctx, networkID, nodeID)
	return err
}

// GetControllerNetworkMember converts echo context to params.
func (w *ServerInterfaceWrapper) GetControllerNetworkMember(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "networkID" -------------
	var networkID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "networkID", runtime.ParamLocationPath, ctx.Param("networkID"), &networkID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter networkID: %s", err))
	}

	// ------------- Path parameter "nodeID" -------------
	var nodeID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "nodeID", runtime.ParamLocationPath, ctx.Param("nodeID"), &nodeID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter nodeID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetControllerNetworkMember(ctx, networkID, nodeID)
	return err
}

// SetControllerNetworkMember converts echo context to params.
func (w *ServerInterfaceWrapper) SetControllerNetworkMember(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "networkID" -------------
	var networkID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "networkID", runtime.ParamLocationPath, ctx.Param("networkID"), &networkID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter networkID: %s", err))
	}

	// ------------- Path parameter "nodeID" -------------
	var nodeID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "nodeID", runtime.ParamLocationPath, ctx.Param("nodeID"), &nodeID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter nodeID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SetControllerNetworkMember(ctx, networkID, nodeID)
	return err
}

// GetNetworks converts echo context to params.
func (w *ServerInterfaceWrapper) GetNetworks(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNetworks(ctx)
	return err
}

// DeleteNetwork converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteNetwork(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "networkID" -------------
	var networkID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "networkID", runtime.ParamLocationPath, ctx.Param("networkID"), &networkID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter networkID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteNetwork(ctx, networkID)
	return err
}

// GetNetwork converts echo context to params.
func (w *ServerInterfaceWrapper) GetNetwork(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "networkID" -------------
	var networkID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "networkID", runtime.ParamLocationPath, ctx.Param("networkID"), &networkID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter networkID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNetwork(ctx, networkID)
	return err
}

// UpdateNetwork converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateNetwork(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "networkID" -------------
	var networkID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "networkID", runtime.ParamLocationPath, ctx.Param("networkID"), &networkID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter networkID: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateNetwork(ctx, networkID)
	return err
}

// GetPeers converts echo context to params.
func (w *ServerInterfaceWrapper) GetPeers(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPeers(ctx)
	return err
}

// GetPeer converts echo context to params.
func (w *ServerInterfaceWrapper) GetPeer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPeer(ctx, address)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
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

	router.GET(baseURL+"/controller", wrapper.GetControllerStatus)
	router.GET(baseURL+"/controller/network", wrapper.GetControllerNetworks)
	router.POST(baseURL+"/controller/network/:controllerID______", wrapper.GenerateControllerNetwork)
	router.DELETE(baseURL+"/controller/network/:networkID", wrapper.DeleteControllerNetwork)
	router.GET(baseURL+"/controller/network/:networkID", wrapper.GetControllerNetwork)
	router.POST(baseURL+"/controller/network/:networkID", wrapper.SetControllerNetwork)
	router.GET(baseURL+"/controller/network/:networkID/member", wrapper.GetControllerNetworkMembers)
	router.DELETE(baseURL+"/controller/network/:networkID/member/:nodeID", wrapper.DeleteControllerNetworkMember)
	router.GET(baseURL+"/controller/network/:networkID/member/:nodeID", wrapper.GetControllerNetworkMember)
	router.POST(baseURL+"/controller/network/:networkID/member/:nodeID", wrapper.SetControllerNetworkMember)
	router.GET(baseURL+"/network", wrapper.GetNetworks)
	router.DELETE(baseURL+"/network/:networkID", wrapper.DeleteNetwork)
	router.GET(baseURL+"/network/:networkID", wrapper.GetNetwork)
	router.POST(baseURL+"/network/:networkID", wrapper.UpdateNetwork)
	router.GET(baseURL+"/peer", wrapper.GetPeers)
	router.GET(baseURL+"/peer/:address", wrapper.GetPeer)
	router.GET(baseURL+"/status", wrapper.GetStatus)
}