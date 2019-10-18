package main


import (
        "fmt"
        "log"
        "os/exec"
        "net/http"
        "io"
        //"io/ioutil"
)


type handler struct {
}


func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
        fmt.Println(req.URL.Path)

        if req.URL.Path == "/" {
                w.Write([]byte("<h1>I be healthful</h1>\n"))
        } else if req.URL.Path == "/alignmentHeader" {

                params := req.URL.Query()
                url := params["url"][0]

                handle("alignmentHeader.sh", []string{url}, w)
        } else if req.URL.Path == "/baiReadDepth" {

                params := req.URL.Query()
                url := params["url"][0]

                handle("baiReadDepth.sh", []string{url}, w)
        } else {

                w.Write([]byte("Hi there\n"))
        }
}

func handle(scriptName string, args []string, res http.ResponseWriter) {
        scriptPath := fmt.Sprintf("./scripts/%s", scriptName)
        cmd := exec.Command(scriptPath, args...)
        stdout, err := cmd.StdoutPipe()
        if err != nil {
                log.Println(err)
        }

        if err := cmd.Start(); err != nil {
                log.Println(err)
        }

        io.Copy(res, stdout)

        if err := cmd.Wait(); err != nil {
                log.Println(err)
        }
}


func main() {
        fmt.Println("Hi there")

        h := &handler{};

        server := &http.Server{
                Addr:           ":9001",
                Handler:        h,
        }

        log.Fatal(server.ListenAndServe())
}
