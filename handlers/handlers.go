package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetYear(c *gin.Context) {

	year, err := strconv.ParseUint(c.Params.ByName("year"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid input format\n")
		return
	}

	day, dur := duration(year)

	c.Header("X-Ping", "ping")

	c.Writer.WriteString(day + ": " + strconv.FormatUint(dur, 10) + "\n")

}

func duration(year uint64) (string, uint64) {

	date := time.Date(int(year), 1, 1, 0, 0, 0, 0, time.Local)

	if time.Now().After(date) {

		dur := time.Since(date)
		return "Days gone", uint64(dur.Round(time.Hour).Hours()) / 24

	} else {

		dur := time.Until(date)
		return "Days left", uint64(dur.Round(time.Hour).Hours()) / 24
	}

}
