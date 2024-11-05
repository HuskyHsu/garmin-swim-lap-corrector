package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/tormoder/fit"
)

func decodeFit(this js.Value, p []js.Value) interface{} {
	if len(p) < 1 {
		return "Error: No file buffer provided"
	}

	data := make([]byte, p[0].Length())
	js.CopyBytesToGo(data, p[0])

	fitFile, err := fit.Decode(bytes.NewReader(data))
	if err != nil {
		return fmt.Sprintf("Decode error: %v", err)
	}

	activity, err := fitFile.Activity()
	if err != nil {
		return fmt.Sprintf("Activity error: %v", err)
	}

	jsonData, err := json.Marshal(activity)
	if err != nil {
		return fmt.Sprintf("JSON encoding error: %v", err)
	}

	return string(jsonData)
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("decodeFit", js.FuncOf(decodeFit))

	<-c
}
