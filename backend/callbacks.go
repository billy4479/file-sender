package main

import (
	"io"
	"mime/multipart"
	"sync"
)

type callbackData struct {
	dlCallback     func(io.Writer) error
	statusCallback func() string
	called         bool
}

var callbacks = make(map[string]*callbackData)

func makeCallback(uuid string, wg *sync.WaitGroup, h *multipart.FileHeader) *callbackData {
	status := "Not connected"
	statusFn := func() string {
		return status
	}

	callbacks[uuid] = &callbackData{
		dlCallback: func(w io.Writer) error {
			defer wg.Done()
			status = "Reading uploaded file"

			f, err := h.Open()
			if err != nil {
				status = "Error: " + err.Error()
				return err
			}

			status = "Copying from upload to download"
			_, err = io.Copy(w, f)
			if err != nil {
				status = "Error: " + err.Error()
				return err
			}

			delete(callbacks, uuid)
			f.Close()

			status = "Done"
			return nil
		},
		statusCallback: statusFn,
		called:         false,
	}

	return callbacks[uuid]
}
