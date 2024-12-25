package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhyanio/kubemap/terraform"
	"github.com/gin-gonic/gin"
)

// CORS middleware to handle Cross-Origin Resource Sharing
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		// Handle preflight request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func RunRoute() {
	// Initialize Gin
	r := gin.Default()

	// Use CORS middleware
	r.Use(CORSMiddleware())

	// Parse Terraform state file
	nodes, err := terraform.ParseTfState("./example/terraform.tfstate", "1122113", "us-east-1")
	if err != nil {
		log.Fatalf("Failed to parse Terraform state: %v", err)
	}

	fmt.Println("Parsed nodes:", nodes)

	// Serve Graph Data
	r.GET("/graph", func(c *gin.Context) {
		graph := terraform.Graph{Nodes: nodes}
		c.JSON(http.StatusOK, graph)
	})

	// Start HTTP Server
	fmt.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
