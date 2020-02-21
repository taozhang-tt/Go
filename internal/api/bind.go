package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 从 JSON 绑定
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func BindDemo(router *gin.Engine) {
	// 绑定 JSON 的示例 ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 一个 HTML 表单绑定的示例 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 这个将通过 content-type 头去推断绑定器使用哪个依赖。
		if err := c.ShouldBind(&form); err == nil {
			if form.User == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
}

/**
只绑定查询参数
 */
func ShouldBindQuery(router *gin.Engine) {
	router.Any("/should-bind-query", shouldBindQuery)
}

func shouldBindQuery(c *gin.Context) {
	var person Person
	var err error
	if err = c.ShouldBindQuery(&person); err== nil {	//ShouldBindQuery 只绑定查询参数（及url参数）不绑定表单参数
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
		c.JSON(200, person)
	} else {
		c.JSON(400, err.Error())
	}
}

func ShouldBind(router *gin.Engine)  {
	router.Any("/should-bind", shouldBind)
}

func shouldBind(c *gin.Context) {
	var person Person
	// 如果是 `GET`, 只使用 `Form` 绑定引擎 (`query`) 。
	// 如果 `POST`, 首先检查 `content-type` 为 `JSON` 或 `XML`, 然后使用 `Form` (`form-data`) 。
	// 在这里查看更多信息 https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.JSON(200, person)
}

/** 绑定 HTML 复选框
<form action="/bind-checkbox" method="POST">
    <p>Check some colors</p>
    <label for="red">Red</label>
    <input type="checkbox" name="colors[]" value="red" id="red" />
    <label for="green">Green</label>
    <input type="checkbox" name="colors[]" value="green" id="green" />
    <label for="blue">Blue</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue" />
    <input type="submit" />
</form>
 */
func BindCheckBox(router *gin.Engine) {
	router.POST("/bind-checkbox", bindCheckBox)
}

type myForm struct {
	Colors []string `form:"colors[]"`
}

func bindCheckBox(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}