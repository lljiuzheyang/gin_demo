package routers

import (
	"github.com/gin-gonic/gin"
	"gin_demo/routers/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api_v1 := r.Group("/api/v1")
	{
		//获取标签列表
		api_v1.GET("/tags", v1.GetTags)
		//新建标签
		api_v1.POST("/tags", v1.AddTag)
		//更新指定标签
		api_v1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		api_v1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
