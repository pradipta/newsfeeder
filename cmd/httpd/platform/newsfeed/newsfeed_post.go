package newsfeed

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type newsFeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json: "post"`
}

func AddNews(f *Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsFeedPostRequest{}
		c.Bind(&requestBody)
		item :=
			Item{
				Title: requestBody.Title,
				Post:  requestBody.Post,
			}
		f.Add(item)

		c.JSON(http.StatusOK, item)
	}
}
