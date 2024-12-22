package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhyanio/kubemap/terraform"
	"github.com/gin-gonic/gin"
)

func RunRoute() {
	// Initialize Gin
	r := gin.Default()

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

	// Serve Static HTML
	r.StaticFile("/", "./ui/index.html")

	// Start HTTP Server
	fmt.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
