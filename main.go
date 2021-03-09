package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handleSum)

	port := os.Getenv("PORT")

	router.Run(":"+port)
}

func handleSum(ctx *gin.Context) {
	res := SumTwoNumbers(4, 6)

	ctx.JSON(http.StatusOK, res)
}

// SumTwoNumbers returns a sum of two integer numbers
// just more stuff for the sake of demonstrating travis
func SumTwoNumbers(a, b int) int {
	return a + b
}