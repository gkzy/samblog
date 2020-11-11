package middleware

import (
	"github.com/gkzy/gow"
	"time"
)

var (
	ver int64
)

func init() {
	ver = time.Now().Unix()

}

// Version
func Version() gow.HandlerFunc {
	return func(c *gow.Context) {
		c.Data["ver"] = ver
		c.Data["year"] = time.Now().Year()
		c.Next()
	}
}
