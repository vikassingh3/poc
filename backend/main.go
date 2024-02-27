package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func main() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no  password set
		DB:       0,  // use default DB
	})

	// Check Redis connectivity
	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Printf("Connected to Redis: %s", pong)

	// Initialize Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Setup routes
	router.GET("/data", getData)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getData(c *gin.Context) {
	// Check if data is cached in Redis
	cachedData, err := redisClient.Get(c.Request.Context(), "cached_data").Result()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"data": cachedData})
		return
	}

	// If data not found in cache, generate synthetic data
	syntheticData := generateSyntheticData()

	// Cache the data in Redis for 5 seconds (for demonstration purposes)
	err = redisClient.Set(c.Request.Context(), "cached_data", syntheticData, 5*time.Second).Err()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to cache data")
		return
	}

	// Return data
	c.JSON(http.StatusOK, gin.H{"data": syntheticData})
}

func generateSyntheticData() string {
	// Generating random data for demonstration purposes
	var data []int
	for i := 0; i < 10; i++ {
		data = append(data, rand.Intn(100))
	}
	return strconv.Itoa(data[0]) + "," + strconv.Itoa(data[1]) + "," + strconv.Itoa(data[2]) + "," +
		strconv.Itoa(data[3]) + "," + strconv.Itoa(data[4]) + "," + strconv.Itoa(data[5]) + "," +
		strconv.Itoa(data[6]) + "," + strconv.Itoa(data[7]) + "," + strconv.Itoa(data[8]) + "," +
		strconv.Itoa(data[9])
}
