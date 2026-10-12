package main

import (
    "html/template"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    "news/docs"
)


var tpl = template.Must(template.ParseFiles("index.html"))


func main() {
    docs.SwaggerInfo.Title = "News Demo API"
    docs.SwaggerInfo.Description = "API для новостного приложения"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:3000"
    docs.SwaggerInfo.BasePath = "/"

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        if err := tpl.Execute(c.Writer, nil); err != nil {
            log.Printf("Template error: %v", err)
        }
    })

    r.GET("/swagger/*any",
        ginSwagger.WrapHandler(swaggerFiles.Handler))

    log.Printf("Server started on :%s", port)
    log.Fatal(r.Run(":" + port))
}