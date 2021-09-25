package logger

import (
	"context"
	"gorepair-rest-api/config"
	"gorepair-rest-api/infrastructures/db"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

type httpLog struct {
	Created time.Time
	Message string
}

// type CustomResponseWriter struct {
// 	w  http.ResponseWriter
// 	Code int
// }

var once sync.Once

var Log *logrus.Logger

func init() {
	once.Do(func() {
		Log = newLogger()
	})
}

func newLogger() *logrus.Logger {
	Log := logrus.New()
	log.Println("Setup Logger")
	if config.Get().LogPath != "" {
		err := os.Mkdir(config.Get().LogPath, 0755)
		if err != nil {
			log.Println("Failed to create log path")
		}
		log.Println("Success to create log path")
	}

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   config.Get().LogPath,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     90, //days
		Level:      logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02T15:04:05.999999999Z07:00",
			ForceColors:     true,
		},
	})

	if err != nil {
		Log.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	Log.SetReportCaller(true)

	Log.AddHook(rotateFileHook)

	return Log
}

// func NewCustomResponseWriter(ww http.ResponseWriter) *CustomResponseWriter {
// 	return &CustomResponseWriter{
// 		w: ww,
// 		Code: 0,
// 	}
// }

// func (w *CustomResponseWriter) Header() http.Header {
// 	return w.w.Header()
// }

// func (w *CustomResponseWriter) Write(b []byte) (int, error) {
// 	return w.w.Write(b)
// }

// func (w *CustomResponseWriter) WriteHeader(statusCode int) {
// 	w.Code = statusCode
// 	w.w.WriteHeader(statusCode)
// }

// func logRequest(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w2 := NewCustomResponseWriter(w)
// 		handler.ServeHTTP(w2, r)
// 		log.Printf("%d %s %s\n", w2.Code, r.Method, r.URL)
// 	})
// }

// func log() {
	
// }

// app.Use("/api", })

func LogToMongo(c *fiber.Ctx)  {
	collection := db.NewMongoClient().DB().Database(config.Get().MongoDb_Name).Collection(config.Get().MongoDb_Collection)
	var claims string
    c.Locals("claims", claims)
    c.Next()
	_, e := collection.InsertOne(context.TODO(), claims)
	if e != nil {
		log.Fatalln(e)
	}
}