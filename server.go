package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
)

const (
	EventTitle  = "title"
	EventLyrics = "lyrics"
)

type Event struct {
	data interface{}
	name string
}

var (
	serverStarted bool
	serverMutex   sync.Mutex
	eventChan     = make(chan Event)
)

func sendEvent(e Event) {
	eventChan <- e
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func isPortOpen(port int) bool {
	address := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return false // Port is already in use
	}
	defer ln.Close()
	return true // Port is available
}

func startServer() {
	serverMutex.Lock()
	defer serverMutex.Unlock()

	if serverStarted {
		log.Println("[lrClient:server] already started")
		return
	}

	port := getConfig().ServerPort

	if !isPortOpen(port) {
		log.Printf("[lrClient:server] port %d is already in use\n", port)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "SSE not supported", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")

		for newEvent := range eventChan {
			event, err := formatServerSentEvent(newEvent.name, newEvent.data)
			if err != nil {
				fmt.Println(err)
				break
			}

			_, err = fmt.Fprint(w, event)
			if err != nil {
				fmt.Println(err)
				break
			}

			flusher.Flush()
		}
	})

	log.Printf("[lrClient:server] started on port %d\n", port)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	serverStarted = true
}

func formatServerSentEvent(event string, data interface{}) (string, error) {
	buff := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buff)

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("event: %s\n", event))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", buff.String()))

	return sb.String(), nil
}
