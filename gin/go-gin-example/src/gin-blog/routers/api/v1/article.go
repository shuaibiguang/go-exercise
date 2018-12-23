package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取单个文章
func GetArticle (c *gin.Context) {
	id, _ := com.StrTo(c.Param("id")).Int()

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("id 必须大于1")
	code := e.INVALID_PARAMS

	var data interface{}
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _,err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

// 获取多个文章
func GetArticles (c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = com.StrTo(state).Int()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId, _ = com.StrTo(arg).Int()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

// 添加一篇文章
func AddArticle (c *gin.Context) {
	tagId, _ := com.StrTo(c.Query("tag_id")).Int()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state, _ := com.StrTo(c.DefaultQuery("state", "0")).Int()

	valid := validation.Validation{}

	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		// 确认传入进来的标签是否存在
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})

			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改一篇文章
func EditArticle (c *gin.Context) {
	valid := validation.Validation{}

	id, _ := com.StrTo(c.Param("id")).Int()
	tagId, _ := com.StrTo(c.Query("tag_id")).Int()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {

			data := make(map[string]interface{})

			if state != -1 {
				data["state"] = state
			}
			if title != "" {
				data["title"] = title
			}
			if desc != "" {
				data["desc"] = desc
			}
			if content != "" {
				data["content"] = content
			}
			data["modified_by"] = modifiedBy

			edit := func (code *int) {
				models.EditArticle(id, data)
				*code = e.SUCCESS
			}

			if tagId > 0 && models.ExistTagByID(tagId) {
				data["tag_id"] = tagId
				edit(&code)
			} else if tagId > 0 && ! models.ExistTagByID(tagId) {
				code = e.ERROR_NOT_EXIST_TAG
			} else {
				edit(&code)
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})


}

//删除一篇文章
func DeleteArticle (c *gin.Context) {
	id, _ := com.StrTo(c.Param("id")).Int()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}