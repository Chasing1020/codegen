// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T18:13:46+08:00
// File: actuator.go

package router

import (
	"crud/dal"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// HealthHandler godoc
// @Summary      Check http response health
// @Description  Check http response health:
// @Description  curl --location --request 'localhost:8080/actuator/health'
// @Success      200  object  string  success
// @Failure      400  object  string  failed
// @Router       /actuator/health [get]
func HealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}

// RedisHealthHandler godoc
// @Summary      Check redis health
// @Description  Check redis health:
// @Description  curl --location --request 'localhost:8080/actuator/health/redis'
// @Success      200  object  string  success
// @Failure      400  object  string  failed
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
// @Success      200  object  string  success
// @Failure      400  object  string  failed
// @Router       /actuator/health/mysql [get]
func MySQLHealthHandler(c *gin.Context) {
	var result int
	err := dal.DB.Raw("SELECT 1").Scan(&result).Error
	if err != nil || result != 1 {
		c.JSON(500, gin.H{"status": "DOWN"})
	}
	c.JSON(200, gin.H{"status": "UP"})
}

// SessionHealthHandler godoc
// @Summary      Check session health
// @Description  Check session health:
// @Description  curl --location --request 'localhost:8080/actuator/health/session'
// @Success      200  object  string  success
// @Failure      400  object  string  failed
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
