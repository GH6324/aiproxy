package router

import (
	"github.com/gin-gonic/gin"
	mcp "github.com/labring/aiproxy/core/controller/mcp"
	"github.com/labring/aiproxy/core/middleware"
)

func SetMCPRouter(router *gin.Engine) {
	mcpRoute := router.Group("/mcp", middleware.MCPAuth)

	mcpRoute.GET("/public/:id/sse", mcp.PublicMCPSSEServer)
	mcpRoute.GET("/public/:id", mcp.PublicMCPStreamable)
	mcpRoute.POST("/public/:id", mcp.PublicMCPStreamable)
	mcpRoute.DELETE("/public/:id", mcp.PublicMCPStreamable)

	mcpRoute.GET("/group/:id/sse", mcp.GroupMCPSSEServer)
	mcpRoute.GET("/group/:id", mcp.GroupMCPStreamable)
	mcpRoute.POST("/group/:id", mcp.GroupMCPStreamable)
	mcpRoute.DELETE("/group/:id", mcp.GroupMCPStreamable)

	router.GET("/sse", middleware.MCPAuth, mcp.HostMCPSSEServer)
	router.POST("/message", mcp.MCPMessage)
	router.GET("/mcp", middleware.MCPAuth, mcp.HostMCPStreamable)
	router.POST("/mcp", middleware.MCPAuth, mcp.HostMCPStreamable)
	router.DELETE("/mcp", middleware.MCPAuth, mcp.HostMCPStreamable)

	publicMcpTestRoute := router.Group("/test-publicmcp")
	{
		publicMcpTestRoute.GET("/:group/:id/sse", mcp.PublicMCPSSEServer)
	}
}
