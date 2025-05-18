package main

import (
	"fmt"
	"jetbra-free/internal/certificate"
	"jetbra-free/internal/core"
	"jetbra-free/internal/util"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func init() {
	binDir, _ := os.Executable()
	binDir = filepath.Dir(binDir)

	err := util.ExtractAssets(binDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to extract assets: %v", err))
	}
	core.PluginsInit()
	certificate.Run()
	core.LicenseInit()
}

func main() {
	fmt.Println("     _      _   _                 _____              ")
	fmt.Println("    | | ___| |_| |__  _ __ __ _  |  ___| __ ___  ___ ")
	fmt.Println(" _  | |/ _ \\ __| '_ \\| '__/ _` | | |_ | '__/ _ \\/ _ \\")
	fmt.Println("| |_| |  __/ |_| |_) | | | (_| | |  _|| | |  __/  __/")
	fmt.Println(" \\___/ \\___|\\__|_.__/|_|  \\__,_| |_|  |_|  \\___|\\___|")
	address := "127.0.0.1:8080"
	fmt.Printf("\nServer running at: http://%s\n", address)
	gin.SetMode(gin.ReleaseMode)
	// init route
	r := gin.Default()
	r.Use(cors())

	binDir, _ := os.Executable()
	binDir = filepath.Dir(binDir)
	r.Static("/static", filepath.Join(binDir, "static"))

	r.LoadHTMLGlob(filepath.Join(binDir, "templates/*"))

	r.GET("/", core.Index)
	r.POST("/generateLicense", core.GenerateLicense)
	r.Run(address)
}
