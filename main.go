package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Slack_Name		string		`json:"slack_name"`
	Current_Day		string		`json:"current_day"`
	UTC_Time		time.Time	`json:"utc_time"`
	Track			string		`json:"track"`
	Github_File_URL	string		`json:"github_file_url"`
	Github_Repo_URL	string		`json:"github_repo_url"`
	Status_Code		int			`json:"status_code"`

}

func main() {
	router := gin.Default()

	router.GET("/api", func(ctx *gin.Context) {
	
		slack_name := ctx.Query("slack_name")
		track := ctx.Query("track")

		res := &Response{
			Slack_Name:      slack_name,
			Current_Day:     time.Now().Weekday().String(),
			UTC_Time:        time.Now().UTC(),
			Track:           track,
			Github_File_URL: "https://github.com/ashiruhabeeb/hng-internship/blob/main/main.go",
			Github_Repo_URL: "https://github.com/ashiruhabeeb/hng-internship",
			Status_Code:     200,
		}

		ctx.JSON(http.StatusOK, res)
	})

	router.Run(":8090")
}