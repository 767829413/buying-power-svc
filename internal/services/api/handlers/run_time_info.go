package handlers

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"runtime"
	"sync"
)

//RunTimeInfo - provides basic information on the service
type RunTimeInfo struct {
	TotalConnections      uint64            `json:"total_connections"`
	ConnectionsPerAccount map[string]uint64 `json:"connections_per_account"`
	Goroutines            int               `json:"goroutines"`
	lock                  sync.Mutex        `json:"-"`
}

//NewRunTimeInfo - creates new instance of runtime info
func NewRunTimeInfo() *RunTimeInfo {
	return &RunTimeInfo{
		TotalConnections:      0,
		ConnectionsPerAccount: map[string]uint64{},
		Goroutines:            0,
		lock:                  sync.Mutex{},
	}
}

//AddConnection - adds new ws connection to info
func (i *RunTimeInfo) AddConnection(accountID string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.TotalConnections++
	i.ConnectionsPerAccount[accountID] = i.ConnectionsPerAccount[accountID] + 1
}

//RemoveConnection - removes connection from info
func (i *RunTimeInfo) RemoveConnection(accountID string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.TotalConnections--
	i.ConnectionsPerAccount[accountID] = i.ConnectionsPerAccount[accountID] - 1
	if i.ConnectionsPerAccount[accountID] == 0 {
		delete(i.ConnectionsPerAccount, accountID)
	}
}

//Render - renders to json info
func (i *RunTimeInfo) Render() []byte {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.Goroutines = runtime.NumGoroutine()
	asJSON, err := json.Marshal(i)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal runtime info"))
	}

	return asJSON
}

//Info - renders runtime info
func Info(w http.ResponseWriter, r *http.Request) {
	// TODO: allow only local requests
	response := GetRunTimeInfo(r).Render()
	_, _ = w.Write(response)
}
