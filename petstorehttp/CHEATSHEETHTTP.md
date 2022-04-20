# CheatSheet
This CheatSheet provides some common uses and tips for `net/http` package.

### Import
```go
import "net/http"
```

### Start DEFAULT Server
```go
http.ListenAndServe(":8080", nil)
```

### Using HandleFunc
```go
http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request){
    // process request.
    // do more stuffs.
    ...
})
```

### Using Handle
```go
type PetsHandler struct {}

func (p PetsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // process request.
    // do more stuffs.
    ...
}

http.Handle("/pets", PetsHandler{})
```

### NotFound Response
```go
http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request){
    // process request.
    // do more stuffs.
    ...
    // if pets were not found...
    http.NotFound(w, r)
})
```

### Error Response
```go
http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request){
    // process request.
    // do more stuffs.
    ...
    // if pets were not found...
    http.Error(w, "404 pets not found", http.StatusNotFound)
})
```

### HTTP Method
```go
http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request){
    // process request.
    if r.Method == "GET" {
        // process GET request only.
    }
})
```

### Get URL query values
```go
http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request){
    // process request.
    r.URL.Values()
})
```

### Create a custom server
```go
s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

log.Fatal(s.ListenAndServe())
```

### Handle static files content
```go
http.Handle("/", http.FileServer(http.Dir("/assets")))
```