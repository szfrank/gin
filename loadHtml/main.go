package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postsIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"title": "Posts",
	})
}

func usersIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "Users",
	})
}

// func formatAsDate(t time.Time) string {
// 	year, month, day := t.Date()
// 	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
// }

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("loadHtml/templates/**/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/posts/index", postsIndex)
	router.GET("/users/index", usersIndex)

	// router := gin.Default()
	// router.Delims("{[{", "}]}")
	// router.SetFuncMap(template.FuncMap{
	// 	"formatAsDate": formatAsDate,
	// })
	// router.LoadHTMLFiles("./templates/raw.html")

	// router.GET("/raw", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "raw.html", map[string]interface{}{
	// 		"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
	// 	})
	// })
	router.Run(":8080")
}
