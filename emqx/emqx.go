package emqx

import (
	"errors"
	"fmt"

	"github.com/danclive/rest"
)

var baseUrl = "http://127.0.0.1:8080/api/v3"
var appID = "4383ae64e969d8"
var appPassword = "Mjg2OTM3OTQ3NTkzNDU3NjM4ODEwMDQ4Mzk5NjIxNDg4NjE"

var Rest *rest.Rest

func InitEmqxApi() {
	Rest = rest.NewRest().BaseUrl(baseUrl)

	Rest.Before(func(req *rest.Req) {
		req.BasicAuth(appID, appPassword)
	})

	Rest.After(func(res *rest.Res) {

	})
}

type Meta struct {
	Count int `json:"count"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Api struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Path   string `json:"path"`
	Desc   string `json:"descr"`
}

func Apis() ([]Api, error) {
	var result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    []Api  `json:"data"`
	}

	res, err := Rest.Get("/").Send()
	if err != nil {
		return nil, err
	}

	err = res.Json(&result)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("%s", res.StatusCode())
	}

	if result.Code > 0 {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

type Broker struct {
	Datetime   string `json:"datetime"`
	Node       string `json:"node"`
	NodeStatus string `json:"node_status"`
	OtpRelease string `json:"otp_release"`
	SysDesc    string `json:"sysdescr"`
	Uptime     string `json:"uptime"`
}

func Brokers() ([]Broker, error) {
	var result struct {
		Code    int      `json:"code"`
		Message string   `json:"message"`
		Data    []Broker `json:"data"`
	}

	res, err := Rest.Get("/brokers/").Send()
	if err != nil {
		return nil, err
	}

	err = res.Json(&result)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("%s", res.StatusCode())
	}

	if result.Code > 0 {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}

func BrokerItem(node string) (Broker, error) {
	var result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    Broker `json:"data"`
	}

	result.Data.Node = node

	res, err := Rest.Get("/brokers/" + node).Send()
	if err != nil {
		return Broker{}, err
	}

	err = res.Json(&result)
	if err != nil {
		return Broker{}, err
	}

	if res.StatusCode() != 200 {
		return Broker{}, fmt.Errorf("%s", res.StatusCode())
	}

	if result.Code > 0 {
		return Broker{}, errors.New(result.Message)
	}

	return result.Data, nil

}

type Node struct {
	Connections      string `json:"connections"`
	Load1            string `json:"load1"`
	Load5            string `json:"load5"`
	Load15           string `json:"load15"`
	MaxFds           string `json:"max_fds"`
	MemoryTotal      string `json:"memory_total"`
	MemoryUsed       string `json:"memory_used"`
	Name             string `json:"name"`
	Node             string `json:"node"`
	NodeStatus       string `json:"node_status"`
	OtpRelease       string `json:"otp_release"`
	ProcessAvailable int    `json:"process_available"`
	ProcessUsed      int    `json:"process_used"`
	Uptime           string `json:"uptime"`
	Vsersion         string `json:"version"`
}

func Nodes() {

}

func NodeItem() {

}

type Connection struct {
	CleanStart  bool   `json:"clean_start"`
	ClientID    string `json:"client_id"`
	ConnectedAt string `json:"connected_at"`
	IPAddress   string `json:"ipaddress"`
	IsBridge    bool   `json:"is_bridge"`
	IsSuper     bool   `json:"is_super"`
	KeepAlive   int    `json:"keepalive"`
	MountPoint  string `json:"mount_point"`
	Node        string `json:"node"`
	Peercert    string `json:"peercert"`
	Port        int    `json:"port"`
	ProtoName   string `json:"proto_name"`
	ProtoVer    int    `json:"proto_ver"`
	UserName    string `json:"username"`
	WillTopic   string `json:"will_topic"`
	Zone        string `json:"external"`
}

func BrokerConnections() {

}

func NodeConnections() {

}

func ConnectionItem(clientid string) {

}

func NodeConnectionItem(node, clientid string) {

}

func DeleteConnection(clientid string) {

}

type Session struct {
	AwaitingRelLen     int    `json:"awaiting_rel_len"`
	Binding            string `json:"binding"`
	CleanStart         bool   `json:"clean_start"`
	ClientID           string `json:"client_id"`
	CreatedAt          string `json:"created_at"`
	DeliverMsg         int    `json:"deliver_msg"`
	EnqueueMsg         int    `json:"enqueue_msg"`
	ExpiryInterval     int    `json:"expiry_interval"`
	HeapSize           int    `json:"head_size"`
	InflightLen        int    `json:"inflight_len"`
	MailBox            int    `json:"mail_box"`
	MaxAwaitingRel     int    `json:"max_awaiting_rel"`
	MaxInflight        int    `json:"max_inflight"`
	MaxMQueue          int    `json:"max_mqueue"`
	MaxSubscriptions   int    `json:"max_subscriptions"`
	MQqueueDrppped     int    `json:"mqueue_dropped"`
	MQqueueLen         int    `json:"mqueue_len"`
	Node               string `json:"node"`
	Reductions         int    `json:"reductions"`
	SubscriptionsCount int    `json:"subscriptions_count"`
	Username           string `json:"username"`
}

func BrokerSessions() {

}

func SessionItem(clientid string) {

}

func NodeSessions() {

}

func NodeSessionItem() {

}

type Subscription struct {
	ClientID string `json:"client_id"`
	Node     string `json:"node"`
	Qos      int    `json:"qos"`
	Topic    string `json:"topic"`
}

func BrokerSubscriptions() {

}

func SubscriptionItem() {

}

func NodeSubsctiptions() {

}

func NodeSubsctiptionItem() {

}

type Listener struct {
	Acceptors     int           `json:"acceptors"`
	CurrentConns  int           `json:"current_conns"`
	ListenOn      string        `json:"8883"`
	MaxConns      int           `json:"102400"`
	Protocol      string        `json:"protocol"`
	ShutdownCount ShutdownCount `json:"shutdown_count"`
}

type ShutdownCount struct {
	Closed int `json:"closed"`
	Kicked int `json:"kicked"`
}

func BrokerListeners() {

}

func NodeListeners() {

}
