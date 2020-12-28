package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	libredis "github.com/go-redis/redis"
	limiter "github.com/ulule/limiter"
	mgin "github.com/ulule/limiter/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/drivers/store/redis"
	"os"

	"log"
)

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
func setupRouter() *gin.Engine {
	router := gin.Default()
	redisUrl := getEnv("REDIS_URL", "localhost")
	store, error := redis.NewStore(10, "tcp", redisUrl+":6379", "", []byte("secret"))
	if error != nil {
		log.Fatalf("failed to decode: %s", error)
	}
	router.Use(sessions.Sessions("uid", store))
	router.Use(Session())
	// Define a limit rate to 4 requests per hour.
	rate, err := limiter.NewRateFromFormatted("40000-H")
	if err != nil {
		log.Fatal(err)
	}

	// Create a redis client.
	option, err := libredis.ParseURL("redis://" + redisUrl + ":6379/0")
	if err != nil {
		log.Fatal(err)
	}
	client := libredis.NewClient(option)

	// Create a store with the redis client.
	secondStore, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix: "limiter_gin_example",
	})

	// Create a new middleware with the limiter instance.
	middleware := mgin.NewMiddleware(limiter.New(secondStore, rate))

	// Launch a simple server.
	router.ForwardedByClientIP = true
	router.Use(CORS())
	router.Use(middleware)
	//router.Use(static.Serve("/", static.LocalFile("./client/build/", true)))
	router.Use(static.Serve("/create", static.LocalFile("./client/build/", true)))
	router.Static("/converted", "./Storage/live", )

	api := router.Group("/api")
	{
		api.POST("/upload", firstUpload)
		api.POST("/upload/:id", upload)
		api.GET("/process/:id", process)
		api.GET("/download", file)

	}

	return router
}

func StartServer() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
