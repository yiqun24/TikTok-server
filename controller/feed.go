package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yiqun24/TikTok-server/models"
)

const (
	LIMIT = 30
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	// get video list
	// lateset_time: int64, default:current timestamp
	latesetTime := c.Query("lateset_time")
	// convert latesetTime to timestamp(int64)
	latesetTimeInt, err := strconv.ParseInt(latesetTime, 10, 64)
	if err != nil {
		fmt.Printf("Parse lateset_time failed: %v\n", err)
		latesetTimeInt = time.Now().Unix()
	}
	if latesetTimeInt == 0 {
		latesetTimeInt = time.Now().Unix()
	}
	// access database to get a video list
	var videoList []models.Video
	result := models.DB.Limit(LIMIT).Find(&videoList)
	if result.Error != nil {
		fmt.Printf("Query video list failed: %v\n", result.Error)
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  result.Error.Error(),
			},
		})
	}
	var feedVideos []Video
	for _, video := range videoList {
		var user models.User
		models.DB.First(&user, video.AuthorId)
		feedVideos = append(feedVideos, Video{
			Id: video.ID,
			Author: User{
				Id:            user.ID,
				Name:          user.Username,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
		})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
