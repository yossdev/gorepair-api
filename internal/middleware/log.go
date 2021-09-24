package middleware

// import (
// 	"crypto/tls"
// 	"gorepair-rest-api/config"
// 	"log"
// 	"net"
// 	"sync"
// 	"time"

// 	mgo "gopkg.in/mgo.v2"

// 	"github.com/sirupsen/logrus"
// 	"github.com/weekface/mgorus"
// )

// var once sync.Once

// var Log *logrus.Logger

// func init() {
// 	once.Do(func() {
// 		Log = mongoLogger()
// 	})
// }

// // Save log to MongoDB (Local for now)
// func mongoLogger() *logrus.Logger {
// 	s, err := mgo.DialWithInfo(&mgo.DialInfo{
// 		Addrs:    []string{config.Get().MongoDb_Address},
// 		Timeout:  5 * time.Second,
// 		Database: config.Get().MongoDb_Name,
// 		Username: config.Get().MongoDb_Username,
// 		Password: config.Get().MongoDb_Password,
// 		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
// 			conn, err := tls.Dial("tcp", addr.String(), &tls.Config{InsecureSkipVerify: true})
// 			return conn, err
// 		},
// 	})
// 	if err != nil {
// 		log.Fatalf("Can't create session: %s\n", err)
// 	}

// 	c := s.DB("db").C(config.Get().MongoDb_Collection)

// 	Log := logrus.New()
// 	log.Println("Setup Logger")
// 	hooker := mgorus.NewHookerFromCollection(c)

// 	Log.Hooks.Add(hooker)

// 	Log.WithFields(logrus.Fields{
// 		"name": "zhangsan",
// 		"age":  28,
// 	}).Error("Error Logrus!")

// 	return Log
// }