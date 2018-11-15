package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type WebServer struct {
	Logger *log.Logger
}

const usage = `Usage: server <options>

  Options:
    -p <port>     : webserver listening port
    -l <log path> : log filepath

`

func mainInternal() error {
	var port int
	var logPath string
	flag.IntVar(&port, "p", 8080, "listening port")
	flag.StringVar(&logPath, "l", "/tmp/webserver.log", "log filepath")
	flag.Usage = func() {
		fmt.Println(usage)
	}
	flag.Parse()

	// create log directory is doesn't exist
	logDir := filepath.Base(logPath)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err = os.MkdirAll(logDir, os.FileMode(0755)); err != nil {
			return err
		}
	}
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	// initialize
	logger := log.New(io.Writer(f), "", 0)
	webServer := &WebServer{logger}

	webServer.Log("Starting WebWeb Server")

	httpServer := &http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		Handler:     webServer,
		ReadTimeout: 5 * time.Second,
	}
	if err = httpServer.ListenAndServe(); err != nil {
		webServer.Log("ListenAndServe error: %v", err)
		return err
	}
	return nil
}

func (srv *WebServer) Log(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println(msg)
	srv.Logger.Println(time.Now().Format(time.RFC3339), msg)
}

func (srv *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	srv.Log("%s %s", r.Method, buf.String())
	// TODO implement webserver logic
	//
	w.WriteHeader(http.StatusOK)
}

func main() {
	if err := mainInternal(); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	os.Exit(0)
}
