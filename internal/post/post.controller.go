package post

import (
	"github.com/gin-gonic/gin"
	"github.com/shardent/messec-be/infra/database"
)

func GetAllPostByUserId(c *gin.Context) {
    userId := c.Param("user_id")
    var posts *[]Post

    err := GetPostsByUserId(&posts[], userId)

}
