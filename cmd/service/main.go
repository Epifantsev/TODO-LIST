package main

import (
	"final_task/internal/db"
	"final_task/internal/handlers"
	"os"
	"path/filepath"

	"final_task/internal/repository"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "modernc.org/sqlite"
)

func main() {
	port, ok := checkVar("TODO_PORT")
	if !ok {
		port = ":7540"
	}

	dbPath := databaseFileExist()

	dBase := db.New(dbPath)
	defer dBase.Close()

	repo := repository.New(dBase)
	migration(repo)

	handler := handlers.New(repo)

	r := chi.NewRouter()

	r.Handle("/*", http.FileServer(http.Dir("./web")))

	r.Post("/api/task", handler.AddTask)
	r.Post("/api/task/done", handler.DoneTask)

	r.Put("/api/task", handler.PutTask)

	r.Delete("/api/task", handler.DeleteTask)

	r.Get("/api/nextdate", handlers.NextDate)
	r.Get("/api/tasks", handler.GetTasks)
	r.Get("/api/task", handler.GetTask)

	log.Printf("Сервер успешно запущен на порту %s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

func migration(rep *repository.Repository) {
	if err := rep.CreateScheduler(); err != nil {
		log.Fatal(err)
	}

	if err := rep.IndexDate(); err != nil {
		log.Fatal(err)
	}
}

func databaseFileExist() string {
	var err error

	appPath, ok := checkVar("TODO_DBFILE")
	if !ok {
		appPath, err = os.Executable()
		if err != nil {
			log.Fatal(err)
		}
	}

	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")

	_, err = os.Stat(dbFile)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	if os.IsNotExist(err) {
		_, err = os.Create(dbFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dbFile
}

func checkVar(value string) (string, bool) {
	result := os.Getenv(value)

	if result == "" {
		return result, false
	}

	return result, true
}
