package main
import(
    "bytes"
    "strings"
    "errors"
    "fmt"
    "net/http"
    "os"
    "time"
    "io"
    "io/ioutil"
    "encoding/json"
)

type Server struct {
    port int
    body http.Server
}

type MyClient struct {
    port int
    body http.Client
}

func (s *Server) SetAndStartHttpServer () {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("server: %s /\n", r.Method)
        fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
        fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
        fmt.Println("server: headers:")
        for headerName, headerValue := range r.Header {
            fmt.Printf("\t%s = %s\n", headerName,
                       strings.Join(headerValue, ", "))
        }
        requestByteSlice, err := ioutil.ReadAll(r.Body)
        if err != nil {
            fmt.Printf("server: couldn't read request body: %s\n", err)
        }
        fmt.Printf("server: request body: %s\n", requestByteSlice)
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

func (c *MyClient) GetRequest(requestURL string) {
    response, err := http.Get(requestURL)
    // If there"s a problem, exit the program with return value = 1.
    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        os.Exit(1)
    }
    time.Sleep(50 * time.Millisecond) // Might not be needed, check.
    fmt.Println("client: got response! YEY!")
    fmt.Printf("client: status code %d\n", response.StatusCode)
    // Reading the response's body.
    responseByteSlice, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("client: couldn't read the response: %s\n", err)
        os.Exit(1)
    }
    // Print the json response with proper indentation.
    var byteBuffer bytes.Buffer
    json.Indent(&byteBuffer, responseByteSlice, "", "\t")
    if err != nil {
        fmt.Printf("couldn't pretty print json: %s\n", err)
    }
    fmt.Printf("\nclient: response body:\n")
    byteBuffer.WriteTo(os.Stdout)
}

func ValidateHttpMethod(method string) bool {
    switch method {
    case http.MethodGet:
        return true
    case http.MethodPost:
        return true
    default:
        return false
    }
}

func (c *MyClient) MakeRequest(method string, requestURL string,
                              requestBody io.Reader) {
    // Checking the method we use.
    if isMethodValid := ValidateHttpMethod(method); !isMethodValid {
        fmt.Println("This type of request is not supported.")
        os.Exit(1)
    }
    // Create an http request 'object' with the URL inside.
    request, err := http.NewRequest(method, requestURL, requestBody)
    // If there"s a problem, exit the program with return value = 1.
    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        os.Exit(1)
    }
    /* Setting the content type of the request's body to be json media
    type in our requests header. */
    request.Header.Set("Content-Type", "application/json")

    // Activate the request.
    response, err := c.body.Do(request)
    if err != nil {
        fmt.Printf("error making http request: %s\n", err)
        os.Exit(1)
    }
    time.Sleep(50 * time.Millisecond) // Might not be needed, check.
    fmt.Println("client: got response! YEY!")
    fmt.Printf("client: status code %d\n", response.StatusCode)
    // Reading the response's body.
    responseByteSlice, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("client: couldn't read the response: %s\n", err)
        os.Exit(1)
    }
    // Print the json response with proper indentation.
    var byteBuffer bytes.Buffer
    json.Indent(&byteBuffer, responseByteSlice, "", "\t")
    if err != nil {
        fmt.Printf("couldn't pretty print json: %s\n", err)
    }
    fmt.Printf("\nclient: response body:\n")
    byteBuffer.WriteTo(os.Stdout)
}

func main() {
    myClient := MyClient{
        port: 3333,
        body: http.Client{
            Timeout: 30 * time.Second,
        },
    }
    // // Making a GET request - Test 1: cat facts.
    requestURL := fmt.Sprintf(
        "https://cat-fact.herokuapp.com/facts?animal_type=cat")
    myClient.MakeRequest(http.MethodGet, requestURL, nil)
    fmt.Println("\n")

    // // Making a GET request - Test 2: my user without authentication.
    requestURL := fmt.Sprintf(
        "https://cat-fact.herokuapp.com/users/me")
    myClient.MakeRequest(http.MethodGet, requestURL, nil)
    fmt.Println("\n")

    // Making a POST request - Test 5.
    requestURL = fmt.Sprintf("https://cat-fact.herokuapp.com")
    jsonByteSlice := []byte(`{"client_message": "hemlo, server fren."}`)
    // Creating an io reader 'object' with our message in it.
    bodyReader := bytes.NewReader(jsonByteSlice)
    myClient.MakeRequest(http.MethodPost, requestURL, bodyReader)
}