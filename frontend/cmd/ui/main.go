//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("WebAssembly Loaded!")

	// Expose Go functions to JavaScript
	js.Global().Set("createAccount", js.FuncOf(createAccount))
	js.Global().Set("retrieveAccounts", js.FuncOf(retrieveAccounts))
	js.Global().Set("healthCheck", js.FuncOf(healthCheck))

	// Keep the application running
	select {}
}

// CreateAccount sends a PUT request to the backend to create an account
func createAccount(this js.Value, args []js.Value) interface{} {
	account := map[string]interface{}{
		"account_name":   args[0].String(),
		"account_number": args[1].String(),
		"address":        args[2].String(),
		"amount":         args[3].Float(),
		"iban":           args[4].String(),
		"type":           args[5].String(),
	}

	data, _ := json.Marshal(account)
	go func() {
		resp, err := httpRequest("PUT", "http://localhost:8080/accounts", string(data))
		if err != nil {
			js.Global().Get("alert").Invoke("Failed to create account: " + err.Error())
			return
		}
		js.Global().Get("alert").Invoke("Account created successfully: " + resp)
	}()
	return nil
}

// RetrieveAccounts sends a POST request to retrieve accounts
func retrieveAccounts(this js.Value, args []js.Value) interface{} {
	pagination := map[string]interface{}{
		"page": args[0].Int(),
		"size": args[1].Int(),
	}

	data, _ := json.Marshal(pagination)
	go func() {
		resp, err := httpRequest("POST", "http://localhost:8080/accounts/retrieve", string(data))
		if err != nil {
			js.Global().Get("alert").Invoke("Failed to retrieve accounts: " + err.Error())
			return
		}
		js.Global().Get("alert").Invoke("Accounts retrieved: " + resp)
	}()
	return nil
}

// HealthCheck sends a GET request to check backend health
func healthCheck(this js.Value, args []js.Value) interface{} {
	go func() {
		resp, err := httpRequest("GET", "http://localhost:8080/healthz", "")
		if err != nil {
			js.Global().Get("alert").Invoke("Health check failed: " + err.Error())
			return
		}
		js.Global().Get("alert").Invoke("Health check response: " + resp)
	}()
	return nil
}

// Helper function to make HTTP requests
func httpRequest(method, url, body string) (string, error) {
	fetch := js.Global().Get("fetch")
	options := map[string]interface{}{
		"method": method,
		"headers": map[string]interface{}{
			"Content-Type": "application/json",
		},
	}
	if body != "" {
		options["body"] = body
	}

	// Create a JavaScript Promise and handle it
	promise := fetch.Invoke(url, options)
	resultChan := make(chan js.Value)
	errChan := make(chan error)

	promise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response := args[0]
		// Extract the text response from the fetch response
		response.Call("text").Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resultChan <- args[0]
			return nil
		})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			errChan <- fmt.Errorf(args[0].String())
			return nil
		}))
		return nil
	})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		errChan <- fmt.Errorf(args[0].String())
		return nil
	}))

	select {
	case result := <-resultChan:
		return result.String(), nil
	case err := <-errChan:
		return "", err
	}
}
