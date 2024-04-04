package main
import(
    "errors"
    "fmt"
    "net/http"
    "os"
    "time"
    "io/ioutil"
)

type Server struct {
    port int
    body http.Server
}

type MyClient struct {
    port int
}

func (s *Server) SetAndStartHttpServer () {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("server: %s /\n", r.Method)
        fmt.Fprintf(w, "{'message': 'hemlo'}")
    })
    s.body = http.Server{
        Addr: fmt.Sprintf(":%d", s.port),
        Handler: mux,
    }
    if err := s.body.ListenAndServe(); err != nil {
        if !errors.Is(err, http.ErrServerClosed) {
            fmt.Printf("error running http server: %s\n", err)
        }
    }
}

func (c *MyClient) GetRequest() {
    requestURL := fmt.Sprintf("http://localhost:%d", c.port)
    response, err := http.Get(requestURL)
    // If there"s a problem, exit the program with return value = 1.
    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        os.Exit(1)
    }
    fmt.Println("client: got response!")
    fmt.Printf("client: status code %d\n", response.StatusCode)
}

func (c *MyClient) MakeRequest() {
    requestURL := fmt.Sprintf("http://localhost:%d", c.port)
    // Create an http request 'object' with the URL inside.
    request, err := http.NewRequest(http.MethodGet, requestURL, nil)
    // If there"s a problem, exit the program with return value = 1.
    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        os.Exit(1)
    }
    // Activate the request.
    response, err := http.DefaultClient.Do(request)
    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        os.Exit(1)
    }
    fmt.Println("client: got response! YEY!")
    fmt.Printf("client: status code %d\n", response.StatusCode)
    // Reading the response's body.
    responseByteSlice, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("client: couldn't read the response: %s\n", err)
        os.Exit(1)
    }
    fmt.Printf("client: response body: %s", responseByteSlice)
}

func main() {
    myServer := Server{port: 3333,
                      body: http.Server{},
                      }
    // Start in a separate goroutine.
    go myServer.SetAndStartHttpServer()
    // Give the server time to start.
    time.Sleep(100 * time.Millisecond)

    myClient := MyClient{port: 3333}
    myClient.MakeRequest()
}