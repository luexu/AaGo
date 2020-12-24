package aa

import (
	"fmt"
	"log"
	"time"

	"github.com/luexu/AaGo/util"
)

type Configuration struct {
	//SID          string // service + ':' + service id, e.g. user01:12
	Service string `yaml:"service"`
	//ServerID     string `yaml:"server_id"`
	Env          string `yaml:"env"`         // dev test preprod product
	TimezoneID   string `yaml:"timezone_id"` // e.g. "Asia/Shanghai"
	TimeLocation *time.Location
	TimeFormat   string `yaml:"time_format"` // e.g. "2006-02-01 15:04:05"
	Mock         bool   `yaml:"mock"`        // using mock

}

func (app *Aa) ParseToConfiguration() {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.Configuration.Env = app.Config.Get("env").String()
	app.Configuration.Service = app.Config.Get("service").String()
	//serverID := app.Config.Get("server_id").String()
	//app.Configuration.ServerID = serverID
	//app.Configuration.SID = svc + ":" + serverID

	if tz := app.Config.Get("timezone_id").String(); tz != "" {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			log.Println("invalid timezone: " + tz + ", error: " + err.Error())
		} else {
			app.Configuration.TimezoneID = tz
			app.Configuration.TimeLocation = loc
		}
	}

	mock, _ := app.Config.Get("mock").Bool()
	app.Configuration.Mock = mock
}

func (c Configuration) Log() {
	msg := fmt.Sprintf("service %s has started! env: %s timezone_id: %s mock: %v git_ver: %s", c.Service, c.Env, c.TimezoneID, c.Mock, util.GitVersion())
	log.Println(msg)
	fmt.Println(msg)
}
