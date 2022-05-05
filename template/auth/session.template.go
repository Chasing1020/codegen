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
	"crud/conf"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
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
`