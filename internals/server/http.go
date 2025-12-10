package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/codeshelldev/gotl/pkg/logger"
	"github.com/codeshelldev/wol-ve/internals/config"
	"github.com/codeshelldev/wol-ve/internals/ve"
	"github.com/codeshelldev/wol-ve/utils/pingutils"
	"github.com/gorilla/websocket"
)

type RequestBody struct {
    ID        	string 	`json:"id,omitempty"`
    IP        	string 	`json:"ip,omitempty"`
	StartupTime	*int	`json:"startupTime,omitempty"`
}

func httpHandler(w http.ResponseWriter, req *http.Request) {
    var body RequestBody

    err := json.NewDecoder(req.Body).Decode(&body)
    if err != nil {
        logger.Error("Could not get Request Body: ", err)
        http.Error(w, "Bad Request: invalid body", http.StatusBadRequest)
        return
    }

    if body.ID == "" {
        http.Error(w, "Bad Request: missing required fields", http.StatusBadRequest)
        return
    }

	clientID := createID()

    resp := map[string]any{
        "client_id": clientID,
    }

	respBytes, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	logger.Debug("Sending client_id to client")

    w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
    w.Write(respBytes)

	f, ok := w.(http.Flusher)
	if ok { 
		f.Flush() 
	}

	logger.Debug("Waiting for client to establish websocket connection")

	client, err := waitForClient(clientID, time.Duration(20 * time.Second))

	if err != nil {
		logger.Error("Could not get client: ", err.Error())
		return
	}

	if body.IP != "" {
		logger.Debug("Pinging virtual host")

		reachable, err := tryPing(client, body.IP, 
			func() (bool, error) {
				sendToClient(client, map[string]any{
					"success": true,
					"message": "Virtual host is reachable.",
				})
				return true, nil
			},
			func() (bool, error) {
				sendToClient(client, map[string]any{
					"success": false,
					"message": "Virtual host is unreachable.",
				})
				return false, nil
			},
		)

		logger.Debug("Virtual host is unreachable")

		if reachable || err != nil {
			closeClient(client)
			return
		}
	}

	sendToClient(client, map[string]any{
		"success": false,
		"message": "Starting virtual host...",
	})

	err = ve.StartVirtualHost(body.ID)

	if err != nil {
		logger.Error("Error starting virtual host: ", err.Error())

		sendToClient(client, map[string]any{
			"success": false,
			"error": true,
			"message": "Could not start virtual host",
		})

		closeClient(client)
		return
	}

	if body.IP != "" {
		logger.Debug("Pinging virtual host again")

		if body.StartupTime != nil {
			time.Sleep(time.Duration(*body.StartupTime) * time.Second)

			reachable, _ := tryPing(client, body.IP, 
				func() (bool, error) {
					sendToClient(client, map[string]any{
						"success": true,
						"message": "Virtual host is now reachable.",
					})
					return true, nil
				},
				func() (bool, error) {
					sendToClient(client, map[string]any{
						"success": false,
						"error": true,
						"message": "Virtual host is still unreachable.",
					})
					return false, nil
				},
			)

			if reachable {
				logger.Debug("Virtual host is now reachable.")
			} else {
				logger.Debug("Virtual host is still unreachable.")
			}

			closeClient(client)
			return
		}

		success, err := tryPingInterval(client, config.ENV.PING_INTERVAL, config.ENV.PING_RETRIES, body.IP)

		if success {
			logger.Debug("Virtual host is now reachable.")
			sendToClient(client, map[string]any{
				"success": true,
				"message": "Virtual Host is now reachable.",
			})
		} else if !success && err == nil {
			logger.Debug("Virtual Host is still unreachable.")
			sendToClient(client, map[string]any{
				"success": false,
				"error": true,
				"message": "Virtual Host is still unreachable.",
			})
		}
	}

	closeClient(client)
}

func tryPingInterval(client *websocket.Conn, interval, retries int, addr string) (bool, error) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	count := 0

	success := false

	for range ticker.C {
		count++

		_, err := tryPing(client, addr, 
			func() (bool, error) { 
				success = true
				return true, nil 
			}, func() (bool, error) { return false, nil },
		)

		if err != nil {
			closeClient(client)
			return false, err
		}

		if count >= retries {
			break
		}
	}

	return success, nil
}

func tryPing(client *websocket.Conn, addr string, reachable, unreachable func() (bool, error)) (bool, error) {
	hostReachable, err := pingutils.Ping(addr)

	if err != nil {
		logger.Error("Error pinging ", addr, ": ", err.Error())

		sendToClient(client, map[string]any{
			"success": false,
			"error": true,
			"message": "Could not ping host",
		})
		return false, err
	}

	if hostReachable {
		return reachable()
	} else {
		return unreachable()
	}
}