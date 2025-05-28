package main

import (
    "log"
    "fmt"
    "github.com/samiyonas/social/internal/env"
    "github.com/samiyonas/social/internal/db"
    "github.com/samiyonas/social/internal/store"
)

func main() {
    // If I have another data that belongs to the app, it will be defined here.
    cnf := config{
        addr: env.GetString("ADDR", ":8080"),
        db: dbConfig{
            addr: env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
            maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
            maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
            maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
        },
    }

    db, err := db.New(
        cnf.db.addr,
        cnf.db.maxOpenConns,
        cnf.db.maxIdleConns,
        cnf.db.maxIdleTime)

    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    log.Println("database connection established")

    store := store.NewStorage(db)

    app := &application{
        config: cnf,
        store: store,
    }

    log.Printf("running on%s", cnf.addr)
    if err := app.run(app.multiplexer()); err != nil {
        log.Fatal(fmt.Errorf("connection error %v", err))
    }
}
