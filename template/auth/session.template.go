/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/5/5-13:13
File: session_tpl.go
*/

package auth

var Template = `// Copyright © 2022 {{.Author}} <{{.Email}}>
// Time: {{ .Time.Format "2006-01-02T15:04:05Z07:00" }}
// File: session.go

package auth

import (
	"{{.Package}}/conf"
	"{{.Package}}/dal"
	"{{.Package}}/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"strings"
)

var Session gin.HandlerFunc

func init() {
	store := cookie.NewStore([]byte("secret"))
	store, err := redis.NewStore(10, "tcp",
		conf.Conf.Redis.Addr,
		conf.Conf.Redis.Password,
		[]byte("secret"))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{
		MaxAge: 60 * 60 * 24 * 7, // expire in a week
	})
	Session = sessions.Sessions("user_session", store)
}

func CookieRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("logged_in")
	if user == nil {
		c.AbortWithStatusJSON(401, model.Resp{Code:401, Message: "Authentication failed"})
		return
	}
	c.Next()
}

// Login godoc
// @Summary      Login
// @Description  Parses a form and checks for specific data
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        username  body    string      false  "username"  default(0)
// @Param        password  query   string      false  "password"  default(0)
// @Success      200       object  model.Resp  success
// @Failure      401       object  model.Resp  failed
// @Failure      500       object  model.Resp  failed
// @Router       /login [post]
func Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(403, model.Resp{Code:403, Message: "Parameters can't be empty"})
		return
	}

	var students model.Student
	if dal.DB.Where("username = ? AND password = ?", username, password).Find(&students).RowsAffected == 0 {
		c.JSON(401, model.Resp{Code:403, Message: "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set("logged_in", username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(200, model.Resp{Code:200, Message: "Login successful"})
}

// Logout godoc
// @Summary      Logout
// @Description  Logout and delete the session
// @Produce      json
// @Success      200  object  model.Resp  success
// @Failure      401  object  model.Resp  failed
// @Failure      500  object  model.Resp  failed
// @Router       /logout [get]
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("logged_in")
	if user == nil {
		c.JSON(401, model.Resp{Code:403, Message: "Authentication failed"})
		return
	}
	session.Delete("logged_in")
	if err := session.Save(); err != nil {
		c.JSON(500, model.Resp{Code:500, Message: "Internel server error"})
		return
	}
	c.JSON(200, model.Resp{Code:200, Message: "Logout successful"})
}
`