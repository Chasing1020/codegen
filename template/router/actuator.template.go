/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/5/5-15:57
File: actuator.go
*/

package router

var ActuatorHeadTemplate = `// Copyright © 2022 {{.Author}} <{{.Email}}>
// Time: {{.Time.Format "2006-01-02T15:04:05Z07:00" }}
// File: actuator.go

package router

import (
	"{{.Package}}/dal"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// HealthHandler godoc
// @Summary      Check http response health
// @Description  Check http response health:
// @Description  curl --location --request 'localhost:8080/actuator/health'
// @Success      200       {object}  model.Resp  "success"
// @Failure      400       {object}  model.Resp  "failed"
// @Router       /actuator/health [get]
func HealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}

// RedisHealthHandler godoc
// @Summary      Check redis health
// @Description  Check redis health:
// @Description  curl --location --request 'localhost:8080/actuator/health/redis'
// @Success      200       {object}  model.Resp  "success"
// @Failure      400       {object}  model.Resp  "failed"
// @Router       /actuator/health/redis [get]
func RedisHealthHandler(c *gin.Context) {
	if dal.RDB.Ping(c).Val() == "" {
		c.JSON(500, gin.H{"status": "DOWN"})
	}
	c.JSON(200, gin.H{"status": "UP"})
}

//MySQLHealthHandler godoc
// @Summary      Check mysql health
// @Description  Check mysql health:
// @Description  curl --location --request 'localhost:8080/actuator/health/mysql'
// @Success      200       {object}  model.Resp  "success"
// @Failure      400       {object}  model.Resp  "failed"
// @Router       /actuator/health/mysql [get]
func MySQLHealthHandler(c *gin.Context) {
	var result int
	err := dal.DB.Session(
		&gorm.Session{Logger: logger.Default.LogMode(logger.Error)},
	).Raw("SELECT 1").Scan(&result).Error
	if err != nil || result != 1 {
		c.JSON(500, gin.H{"status": "DOWN"})
	}
	c.JSON(200, gin.H{"status": "UP"})
}

// SessionHealthHandler godoc
// @Summary      Check session health
// @Description  Check session health:
// @Description  curl --location --request 'localhost:8080/actuator/health/session'
// @Success      200       {object}  model.Resp  "success"
// @Failure      400       {object}  model.Resp  "failed"
// @Router       /actuator/health/session [get]
func SessionHealthHandler(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	err := session.Save()
	if err != nil {
		c.JSON(500, gin.H{"func 'session.Save()' error": err.Error()})
	}
	c.JSON(200, gin.H{"count": count})
}
`
