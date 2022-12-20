package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os/exec"
	"sample-app/dataManager"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-ping/ping"
	"github.com/mdlayher/wol"
)

func main() {
	var r *gin.Engine = gin.Default()
	var dataManager dataManager.DataManager = dataManager.New()
	var inet, _ = net.InterfaceByName("eth0")
	var client, _ = wol.NewRawClient(inet)
	var target, _ = net.ParseMAC("A8:A1:59:19:0E:A4")
	var shutdownCmd = exec.Command("ssh", "shutdown@iris", "sudo systemctl poweroff")

	r.GET("/startup", func(c *gin.Context) {
		c.JSON(http.StatusOK, client.Wake(target))
	})

	r.GET("/shutdown", func(c *gin.Context) {
		c.JSON(http.StatusOK, shutdownCmd.Run())
	})

	r.GET("/status", func(c *gin.Context) {
		var pinger *ping.Pinger = ping.New("iris")
		var counter uint8 = 1
		var available bool = false

		for counter <= 2 {
			var err error = pinger.Resolve()

			if err == nil {
				fmt.Println("look DNS ok")
				break
			}

			fmt.Println("look DNS failed, retry")

			counter++
			time.Sleep(3 * time.Second)
		}

		if counter == 2 {
			c.JSON(http.StatusOK, gin.H{"isRunning": false})
			return
		}

		pinger.Count = 2
		pinger.Timeout = time.Second
		pinger.OnRecv = func(pkt *ping.Packet) {
			pinger.Stop()
			available = true
		}

		if pinger.Run() != nil {
			pinger.Stop()
		}

		c.JSON(http.StatusOK, gin.H{"isRunning": available})
	})

	// TODO Get user value
	r.GET("/rooms/:roomName", func(c *gin.Context) {
		var roomName string = c.Params.ByName("roomName")
		var temperature string = dataManager.GetValueOf(roomName)

		c.JSON(http.StatusOK, gin.H{"roomName": roomName, "temperature": temperature})
	})

	r.Any("/heating-control/*proxyPath", ReverseProxy())

	r.Use(static.Serve("/", static.LocalFile("./public", false)))

	r.Run(":8000")
}

func ReverseProxy() gin.HandlerFunc {

	target := "iris:80"

	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
			req.URL.Path = c.Param("proxyPath")
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
