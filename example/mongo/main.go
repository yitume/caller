package main

import (
	"github.com/globalsign/mgo/bson"
	"github.com/yitume/caller"
	"github.com/yitume/caller/pkg/mongo"
	"time"
)

var cfg = `
[callerMongo.default]
	debug = true
	url = "mongodb://127.0.0.1:27017/admin"
`
var (
	mg *mongo.Client
)

// UserVisit 访客记录
type UserVisit struct {
	ID             bson.ObjectId `bson:"_id"`
	Platform       string        `bson:"platform"`
	URL            string        `bson:"url"`
	Referrer       string        `bson:"referrer"`
	ClientID       string        `bson:"clientID"`
	UserID         uint          `bson:"userID"`
	Date           time.Time     `bson:"date"`
	IP             string        `bson:"ip"`
	DeviceWidth    int           `bson:"deviceWidth"`
	DeviceHeight   int           `bson:"deviceHeight"`
	BrowserName    string        `bson:"browserName"`
	BrowserVersion string        `bson:"browserVersion"`
	DeviceModel    string        `bson:"deviceModel"`
	Country        string        `bson:"country"`
	Language       string        `bson:"language"`
	OSName         string        `bson:"osName"`
	OSVersion      string        `bson:"osVersion"`
}

func main() {
	if err := caller.Init(
		[]byte(cfg),
		mongo.New,
	); err != nil {
		panic(err)
	}

	initModel()

	mg.C("userVisit").Insert(&UserVisit{
		ID:             bson.NewObjectId(),
		Platform:       "TEST",
		URL:            "",
		Referrer:       "",
		ClientID:       "",
		UserID:         0,
		Date:           time.Time{},
		IP:             "",
		DeviceWidth:    0,
		DeviceHeight:   0,
		BrowserName:    "",
		BrowserVersion: "",
		DeviceModel:    "",
		Country:        "",
		Language:       "",
		OSName:         "",
		OSVersion:      "",
	})
}

func initModel() {
	mg = mongo.Caller("default")
}
