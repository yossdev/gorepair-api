package middleware

import (
	"context"
	"fmt"
	"gorepair-rest-api/config"
	"gorepair-rest-api/infrastructures/db"
	"gorepair-rest-api/infrastructures/logger"
	"time"

	"github.com/gofiber/fiber/v2"
)

type logReq struct {
	ReqId 	  uint64
	Timestamp time.Time
	RemoteIP  string
	Hostname  string
	Protocol  string
	Method 	  string
	Path 	  string
	Duration  string
}

type LogMethod interface {
	LogReqRes(ctx *fiber.Ctx) error
}

type logMongo struct {
	DB db.MongoDB
}

func NewLogMongo(DB db.MongoDB) LogMethod {
	return &logMongo{
		DB: DB,
	}
}

func (u *logMongo) LogReqRes(ctx *fiber.Ctx) error {
	id := ctx.Context().ID()
	time := ctx.Context().ConnTime()
	ip := ctx.IP()
	hostname := ctx.Context().URI().Host()
	protocol := ctx.Protocol()
	method := ctx.Context().Method()
	path := ctx.Context().Path()
	duration := ctx.Context().Time()
	diff := duration.Sub(time)

	data := logReq{
		ReqId: id,
		Timestamp: time,
		RemoteIP: ip,
		Hostname: string(hostname),
		Protocol: protocol,
		Method: string(method),
		Path: string(path),
		Duration: fmt.Sprintf("%v", diff),
	}

	//save log to mongo db
	go func() {
		session := u.DB.DB().Database(config.Get().MongoDb_Name).Collection(config.Get().MongoDb_Collection)
		_, err := session.InsertOne(context.TODO(), data)
		if err != nil {
			logger.Log.Infoln("Failed to save logResReq to mongo, with err: ", err)
		} else {
			logger.Log.Infoln("Successfully to save logResReq to mongo")
		}
	}()

	return ctx.Next()
}