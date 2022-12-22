package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"pismo.io/util"
)

// TestCreateGetAccount checks whether an account can be created (POST) and fetched (GET).
func TestCreateGetAccount(t *testing.T) {
	// Make a POST request to create a new account.
	docNumber := "test_document_1"
	payload := fmt.Sprintf(`{"document_number": "%s"}`, docNumber)
	url := fmt.Sprintf("http://localhost:%d/accounts", util.Env.GetListenPort())
	newAcctResp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Errorf("Request failed to URL %s: %s", url, err)
		return
	}

	defer newAcctResp.Body.Close()

	// Validate the HTTP response status.
	if newAcctResp.StatusCode != http.StatusOK {
		t.Errorf("Received HTTP status %d from URL %s", newAcctResp.StatusCode, url)
		return
	}

	// Read and parse the response payload.
	type AccountsResp struct {
		ID             int    `json:"id"`
		DocumentNumber string `json:"document_number"`
	}
	var newAccount AccountsResp
	dataBytes, err := io.ReadAll(newAcctResp.Body)
	if err != nil {
		t.Errorf("Could not read response body from URL %s: %s", url, err)
		return
	}

	err = json.Unmarshal(dataBytes, &newAccount)
	if err != nil {
		t.Errorf("Could not unmarshal JSON response from URL %s: %s", url, err)
		return
	}

	// Make sure we've got a valid account ID.
	if newAccount.ID <= 0 {
		t.Errorf("Received invalid account ID: %d", newAccount.ID)
		return
	}

	// Make sure the document number matches what we requested.
	if newAccount.DocumentNumber != docNumber {
		t.Errorf("Received unexpected document number: %s", newAccount.DocumentNumber)
		return
	}

	// Try fetching the newly created account using a GET request.
	url = fmt.Sprintf("http://localhost:%d/accounts/%d", util.Env.GetListenPort(), newAccount.ID)
	getAcctResp, err := http.Get(url)
	if err != nil {
		t.Errorf("Request failed to URL %s: %s", url, err)
		return
	}

	defer getAcctResp.Body.Close()

	// Validate the HTTP response status.
	if getAcctResp.StatusCode != http.StatusOK {
		t.Errorf("Received HTTP status %d from URL %s", getAcctResp.StatusCode, url)
		return
	}

	var existingAccount AccountsResp
	dataBytes, err = io.ReadAll(getAcctResp.Body)
	if err != nil {
		t.Errorf("Could not read response body from URL %s: %s", url, err)
		return
	}

	err = json.Unmarshal(dataBytes, &existingAccount)
	if err != nil {
		t.Errorf("Could not unmarshal JSON response from URL %s: %s", url, err)
		return
	}

	// Check whether the fetched account matches the previously created account.
	if existingAccount.ID != newAccount.ID {
		t.Errorf("Account ID does not match the newly created account. Received %d, expected: %d",
			existingAccount.ID, newAccount.ID)
		return
	}
	if existingAccount.DocumentNumber != newAccount.DocumentNumber {
		t.Errorf("Document number does not match the newly created account. Received %s, expected %s",
			existingAccount.DocumentNumber, newAccount.DocumentNumber)
		return
	}
}
