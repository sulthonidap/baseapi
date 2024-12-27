package welcome

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	var messages []string
	messages = append(messages,
		"If everything seems under control, you're not going fast enough. - Mario Andretti",
		"If you are going to be in business, you must learn about money: how it works, how it flows, and how to put it to work for you. - Idowu Koyenikan",
		"If you aren’t embarrassed by the first version of your product, you shipped too late. - Reid Hoffman",
		"You need three things to create a successful startup: to start with good people, to make something customers actually want, and to spend as little money as possible. - Paul Graham",
		"Timing, perseverance, and ten years of trying will eventually make you look like an overnight success. - Biz Stone")

	rand.Seed(time.Now().Unix())
	message := messages[rand.Intn(len(messages))]

	c.JSON(http.StatusOK, gin.H{
		"version": "2.3.21",
		"message": message})
}
