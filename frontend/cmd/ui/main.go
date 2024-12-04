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

func retrieveAccounts(this js.Value, args []js.Value) interface{} {
	pagination := map[string]interface{}{
		"page":    args[0].Int(),
		"perPage": args[1].Int(), // Use "perPage" instead of "size"
	}

	data, _ := json.Marshal(pagination)
	go func() {
		resp, err := httpRequest("POST", "http://localhost:8080/accounts/retrieve", string(data))
		if err != nil {
			js.Global().Get("alert").Invoke("Failed to retrieve accounts: " + err.Error())
			return
		}

		// Parse the response JSON
		var result map[string]interface{}
		if err := json.Unmarshal([]byte(resp), &result); err != nil {
			js.Global().Get("alert").Invoke("Failed to parse accounts response: " + err.Error())
			return
		}

		// Get the "revisions" array from the response
		revisions, ok := result["revisions"].([]interface{})
		if !ok {
			js.Global().Get("alert").Invoke("Invalid accounts response format")
			return
		}

		// Access the table's tbody element in the DOM
		document := js.Global().Get("document")
		tableBody := document.Call("querySelector", "#accounts-table tbody")
		tableBody.Set("innerHTML", "") // Clear existing rows

		// Populate the table with the retrieved accounts
		for _, rev := range revisions {
			revision := rev.(map[string]interface{})
			documentData := revision["document"].(map[string]interface{})

			row := document.Call("createElement", "tr")
			row.Set("innerHTML", `
				<td>`+getString(documentData["_id"])+`</td>
				<td>`+getString(documentData["account_name"])+`</td>
				<td>`+getString(documentData["account_number"])+`</td>
				<td>`+getString(documentData["address"])+`</td>
				<td>`+getFloat(documentData["amount"])+`</td>
				<td>`+getString(documentData["iban"])+`</td>
				<td>`+getString(documentData["type"])+`</td>
			`)
			tableBody.Call("appendChild", row)
		}
	}()
	return nil
}

// Helper function to get string value from interface{}
func getString(value interface{}) string {
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v", value)
}

// Helper function to get float value from interface{}
func getFloat(value interface{}) string {
	if value == nil {
		return "0"
	}
	return fmt.Sprintf("%.2f", value)
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

func httpRequest(method, url, body string) (string, error) {
	fmt.Printf("Making %s request to %s with body: %s\n", method, url, body)

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

	promise := fetch.Invoke(url, options)
	resultChan := make(chan js.Value)
	errChan := make(chan error)

	promise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response := args[0]
		response.Call("text").Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			fmt.Printf("Response received: %s\n", args[0].String())
			resultChan <- args[0]
			return nil
		})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			fmt.Printf("Error extracting response text: %s\n", args[0].String())
			errChan <- fmt.Errorf(args[0].String())
			return nil
		}))
		return nil
	})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Printf("Fetch request failed: %s\n", args[0].String())
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
