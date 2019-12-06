package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("file1").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	r := gin.Default()

	r.SetHTMLTemplate(html)

	r.GET("/welcome", func(c *gin.Context) {
        if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(http.StatusOK, "file1", gin.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
