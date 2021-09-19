package main

import (
    "context"
    "fmt"
    "golang.org/x/sync/errgroup"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func getuser(w http.ResponseWriter, r *http.Request)  {
    //do something
    print("just do something")
}

func main()  {
    eg, ctx := errgroup.WithContext(context.Background())

    servers := []string{
        "127.0.0.1:8080",
        "127.0.0.1:8081",
        "127.0.0.1:8082",
    }
    
    for _, server := range servers {
        eg.Go(func() error {
            http.HandleFunc("/get/user", getuser)
            err := http.ListenAndServe(server, nil)
            if err != nil {
                log.Fatalf("accident happened err:%s", err)
            }
            return nil
        })
    }

    eg.Go(func() error {
        signals := []os.Signal{syscall.SIGINT, syscall.SIGQUIT, syscall.SIGBUS}
        sig := make(chan os.Signal, len(signals))
        signal.Notify(sig, signals...)
        for  {
            select {
            case <- ctx.Done():
                return ctx.Err()

            case <- sig:
                return nil
            }
        }
    })

    err := eg.Wait()

    fmt.Printf("err:%s", err)
}
