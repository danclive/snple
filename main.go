package main

import (
	"flag"
	"fmt"

	"github.com/danclive/mqtt-console/api"
	"github.com/danclive/mqtt-console/config"
	"github.com/danclive/mqtt-console/db"
	"github.com/danclive/mqtt-console/log"
	"github.com/sirupsen/logrus"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func run() {

	clientOption := mqtt.NewClientOptions()

	clientOption.AddBroker("tcp://127.0.0.1:1883")
	clientOption.SetClientID("test")
	clientOption.SetUsername("admin")
	clientOption.SetPassword("admin@admin")

	clientOption.SetWill("ddd", "eee", 1, false)

	clientOption.SetOnConnectHandler(func(client mqtt.Client) {
		fmt.Println("connect")
		fmt.Println(client.IsConnected())

		client.Subscribe("testtopic/#", 0, func(client mqtt.Client, msg mqtt.Message) {
			//fmt.Println(client)
			fmt.Println(msg)
		})

	})

	client := mqtt.NewClient(clientOption)

	fmt.Println(client.Connect().Wait())

	select {}
}

func main() {
	log.SetLogFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	log.Info("读取配置文件...")
	config_file := flag.String("c", "config.toml", "config file")
	flag.Parse()
	config.InitConfig(*config_file)
	log.Info("OK!")

	if config.Config.Debug {
		log.SetLogLevel(logrus.DebugLevel)
	}

	// log.Debug("aaaaaaaaaa")
	// log.DebugWithFields("bbb", log.Fields{"aa": "bb"})
	// log.Error("aaaaaaaaaa")
	// log.Info("ccccccccccccc")

	db.InitMongo()
	api.RunApi()
	// log.Fatal("ccdddddddddddd")
	// log.Panic("eeeeeeeeeee")

	//run()

	//emqx.InitEmqxApi()

	/*
		list, err := emqx.ApiList()
		if err != nil {
			panic(err)
		}

		for _, value := range list {
			fmt.Println(value)
		}
	*/
	/*
		list, err := emqx.BrokerList()
		if err != nil {
			panic(err)
		}

		for _, value := range list {
			fmt.Println(value)
		}
	*/
	/*
		item, err := emqx.BrokerItem("emqx@127.0.0.1")
		if err != nil {
			panic(err)
		}

		fmt.Println(item)
	*/
	/*
		list, meta, err := emqx.Banneds(1, 100)
		if err != nil {
			panic(err)
		}

		fmt.Println(list)
		fmt.Println(meta)

		banned := emqx.Banned{
			Who:    "test",
			As:     "client_id",
			Reason: "haha",
			Desc:   "normal banned",
			Until:  1636146187,
		}

		banned2, err := emqx.CreateBanneds(&banned)
		if err != nil {
			panic(err)
		}

		fmt.Println(banned2)
	*/
	/*
		err := emqx.DeleteBanneds("test", "client_id")
		if err != nil {
			panic(err)
		}
	*/
}
