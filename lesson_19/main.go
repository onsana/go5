package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type ToDo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type ToDoStore struct {
	tasks       []ToDo
	redisClient *redis.Client
}

var ctx = context.Background()

func (store *ToDoStore) getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.tasks)
}

func (store *ToDoStore) getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID := params["id"]

	// Check Redis cache first
	cachedTask, err := store.redisClient.Get(ctx, taskID).Result()
	if err == redis.Nil {
		// Cache miss: fetch from database (memory in this case)
		for _, item := range store.tasks {
			if item.ID == taskID {
				// Store in Redis cache
				cacheData, _ := json.Marshal(item)
				store.redisClient.Set(ctx, taskID, cacheData, 10*time.Minute).Err()
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		// If task not found
		json.NewEncoder(w).Encode(&ToDo{})
	} else if err != nil {
		// Error handling Redis connection
		log.Printf("Failed to connect to Redis: %v", err)
		json.NewEncoder(w).Encode(&ToDo{})
	} else {
		// Cache hit: return cached task
		var task ToDo
		_ = json.Unmarshal([]byte(cachedTask), &task)
		json.NewEncoder(w).Encode(task)
	}
}

func (store *ToDoStore) createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task ToDo
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(100000000))
	store.tasks = append(store.tasks, task)

	// Store new task in Redis
	cacheData, _ := json.Marshal(task)
	store.redisClient.Set(ctx, task.ID, cacheData, 10*time.Minute).Err()

	json.NewEncoder(w).Encode(task)
}

func (store *ToDoStore) updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID := params["id"]

	for index, item := range store.tasks {
		if item.ID == taskID {
			store.tasks = append(store.tasks[:index], store.tasks[index+1:]...)
			var task ToDo
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = taskID
			store.tasks = append(store.tasks, task)

			// Update Redis cache
			cacheData, _ := json.Marshal(task)
			store.redisClient.Set(ctx, task.ID, cacheData, 10*time.Minute).Err()

			json.NewEncoder(w).Encode(task)
			return
		}
	}
}

func (store *ToDoStore) deleteTask(w http.ResponseWriter, r *Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID := params["id"]

	for index, item := range store.tasks {
		if item.ID == taskID {
			store.tasks = append(store.tasks[:index], store.tasks[index+1:]...)
			// Remove from Redis cache
			store.redisClient.Del(ctx, taskID)
			break
		}
	}
	json.NewEncoder(w).Encode(store.tasks)
}

func main() {
	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis address
		DB:   0,                // use default DB
	})

	// Ensure Redis is connected
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	store := &ToDoStore{
		tasks: []ToDo{
			{ID: "1", Title: "Task One"},
			{ID: "2", Title: "Task Two"},
		},
		redisClient: redisClient,
	}

	r := mux.NewRouter()
	r.HandleFunc("/tasks", store.getTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", store.getTask).Methods("GET")
	r.HandleFunc("/tasks", store.createTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", store.updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", store.deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
