package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func main() {
	var err error
	conn, err = pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	defer conn.Close(context.Background())
	handleErr(err)

	http.HandleFunc("/", handleReq)
	err = http.ListenAndServe(":8080", nil)
	handleErr(err)
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	address := r.Header.Get("X-Real-Ip")
	if address == "" {
		address = r.Header.Get("X-Forwarded-For")
	}
	if address == "" {
		address = r.RemoteAddr
	}

	ip, _, err := net.SplitHostPort(address)
	if err != nil {
		ip = address
	}

	id := uuid.New().String()
	_, err = conn.Exec(context.Background(), "insert into addresses values ($1, $2)", id, ip)
	handleErr(err)

	io.WriteString(w, id)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
