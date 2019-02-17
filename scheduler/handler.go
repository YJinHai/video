package main

import (
	"fmt"
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"time"
)

func DeleteVideo(c *gin.Context) {
	vid := c.Param("vid")

	jobrunner.In(10*time.Second, ReminderEmails{
		vid:vid,
	})
	c.JSON(200, "ok")
}

// Job Specific Functions
type ReminderEmails struct {
	vid string
}

// ReminderEmails.Run() will get triggered automatically.
func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email
	fmt.Printf("Every 5 sec send reminder emails:%s \n",e.vid)
}


