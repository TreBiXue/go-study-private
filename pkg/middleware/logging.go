package middleware

import (
	"go-studying/models"
	"go-studying/repo"
	"time"

	"github.com/gin-gonic/gin"
)

func RecordUaAndTime(logrepo repo.IAccessLogRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		// logger, err := zap.NewProduction()
		// if err != nil {
		// 	log.Fatal(err.Error())
		// }
		oldTime := time.Now()
		ua := c.GetHeader("User-Agent")
		c.Next()
		accesslog := models.AccessLog{}
		accesslog.URL = c.Request.URL.String()
		accesslog.Path = c.Request.URL.Path
		accesslog.Ua = ua
		accesslog.Status = c.Writer.Status()
		accesslog.Elapsed = float64(time.Since(oldTime).Seconds())
		logrepo.SaveAccessLog(c, accesslog)

	}
}

// func RecordUaAndTime1(c *gin.Context) {
// 	// logger, err := zap.NewProduction()
// 	// if err != nil {
// 	// 	log.Fatal(err.Error())
// 	// }
// 	oldTime := time.Now()
// 	ua := c.GetHeader("User-Agent")
// 	c.Next()
// 	accesslog := models.AccessLog{}
// 	accesslog.URL = c.Request.URL.String()
// 	accesslog.Path = c.Request.URL.Path
// 	accesslog.Ua = ua
// 	accesslog.Status = c.Writer.Status()
// 	accesslog.Elapsed = float32(time.Since(oldTime).Seconds())
// 	logrepo.SaveAccessLog(c, accesslog)

// 	// logger.Info("incoming request",
// 	// 	zap.String("URL", c.Request.URL.String()),
// 	// 	zap.String("path", c.Request.URL.Path),
// 	// 	zap.String("Ua", ua),
// 	// 	zap.Int("status", c.Writer.Status()),
// 	// 	zap.Duration("elapsed", time.Since(oldTime)),
// 	// )
// }
