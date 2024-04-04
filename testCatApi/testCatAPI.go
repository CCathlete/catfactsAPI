package main
import(
    "errors"
    "fmt"
    'net/http'
    "os"
    'time'
)

const serverPort = 3333

func main() {
    go func() {
        mux := http.NewServerMux()
    }
}