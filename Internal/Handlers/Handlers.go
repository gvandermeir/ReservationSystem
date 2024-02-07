package handlers

import datahandler "github.com/gvandermeir/ReservationSystem/Internal/DataHandler"

type Handlers struct {
	dh *datahandler.DataHandler
}

// NewHandlers returns a new handlers
func NewHandlers() (*Handlers, error) {
	dh, err := datahandler.NewDataHandler()
	return &Handlers{dh:dh}, err
}

