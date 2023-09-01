package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"smarthome/auth"
	"smarthome/interfaces"
	"smarthome/repository"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-ping/ping"
)

func Authenticator() gin.HandlerFunc {
	var session *auth.Session = auth.New("http://fritz.box", "smarthome", "smartTestHome1")

	return func(c *gin.Context) {
		var sid, err = session.GetSID()

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.Set("sid", sid)
		c.Next()
	}
}

func main() {
	var r *gin.Engine = gin.New()
/*
	var inet, netErr = net.InterfaceByName("eth0")
	if netErr != nil {
		fmt.Println(netErr)
		return
	}

	var client, wolErr = wol.NewRawClient(inet)
	if wolErr != nil {
		fmt.Println(wolErr)
		return
	}

	var target, macErr = net.ParseMAC("A8:A1:59:19:0E:A4")
	if macErr != nil {
		fmt.Println(macErr)
		return
	}*/

	var deviceRepository *repository.Repository = repository.NewRepository()

	r.GET("/startup", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, client.Wake(target))
	})

	r.GET("/shutdown", func(c *gin.Context) {
		var shutdownCmd = exec.Command("ssh", "shutdown@iris", "sudo systemctl poweroff")
		c.JSON(http.StatusOK, shutdownCmd.Start())
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

	r.GET("/devices", Authenticator(), func(c *gin.Context) {
		var sid string = c.MustGet("sid").(string)
		var devices, err = deviceRepository.GetDevices(sid)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, devices)
	})

	r.PATCH("/devices", Authenticator(), func(c *gin.Context) {
		var sid string = c.MustGet("sid").(string)
		var id string = c.Query("id")
		var deviceType string = c.Query("type")

		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		var deviceState *interfaces.DeviceStateDto
		err = json.Unmarshal([]byte(data), &deviceState)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		err = deviceRepository.SetDevice(sid, id, deviceType, deviceState)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.Status(http.StatusOK)
	})

	r.Use(static.Serve("/", static.LocalFile("./public", false)))

	r.Run(":8000")
}
