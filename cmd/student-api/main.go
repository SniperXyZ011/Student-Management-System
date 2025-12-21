package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SniperXyZ011/Student-Management-System/internal/config"
	"github.com/SniperXyZ011/Student-Management-System/internal/http/handlers/student"
	"github.com/SniperXyZ011/Student-Management-System/internal/storage/sql"
	// "github.com/SniperXyZ011/Student-Management-System/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setup
	// storage, err := sqlite.New(cfg)
	storage, err := sql.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initilized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.DeleteById(storage))
	router.HandleFunc("PUT /api/students/{id}", student.EditById(storage))

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server has started :", slog.String("Address", cfg.Addr)) //we are using slog, which is nothing but structured log, kafi accha hai for production logs ke liye, ex : 2025/12/17 23:20:56 INFO Server has started, %s Address=localhost:8080

	// err := server.ListenAndServe() //basic way of running server
	// if err != nil {
	// 	log.Fatal("Failed to start server")
	// }

	//in production gracefully server terminate hona chiye uske liye we use go routines
	//we will use channels to keep it blocking so that ye instantly terminate na ho jaye
	done := make(chan os.Signal, 1) // ye os.singal type ka hai mtlb jo signal aate hai, like kucch signal dikhta hoga na server start ya end krte time waisa, also we have made it buffered with size 1

	//yha par ab ham signal channel mai daalegai
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) //so jiase hi ye signals encounter honge ye iss channel ke upar bejh dega
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done //yha pr ab isko nikal degai, now ab ye work aise krega ki jise hi iss server ko terminate krna ka signal jaygea iss channel mai then ye unblock hoga

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //agr server infinte loop mai fas jaye terminate hote time and 5sec se jada time lag jayge to fir ye error ctx bejh dega and uss situation ko acche se handle kr lega
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown the server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}
