package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"pismo.io/db"
)

// Handler implements the HTTP handler logic.
type Handler struct {
	dbAdapter *db.DBAdapter
}

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type CreateAccountResponse struct {
	ID             int32
	DocumentNumber string `json:"document_number"`
}

type GenericResponse struct {
	Status string `json:",omitempty"`
	Error  string `json:",omitempty"`
}

// NewHandler creates a new instance of Handler.
func NewHandler(dbAdapter *db.DBAdapter) (*Handler, error) {
	return &Handler{
		dbAdapter: dbAdapter,
	}, nil
}

// Error is a helper function to send an common error response in JSON format.
func (a *Handler) Error(w http.ResponseWriter, msg string, code int) {
	resp := GenericResponse{Error: msg}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

// CreateAccount is the HTTP handler function for creating a new account.
func (a *Handler) CreateAccount(w http.ResponseWriter, req *http.Request) {
	// Only accept POST.
	if req.Method != http.MethodPost {
		a.Error(w, fmt.Sprintf("%s not allowed", req.Method), http.StatusMethodNotAllowed)
		return
	}

	// Read the request payload.
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		a.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var data CreateAccountRequest
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		a.Error(w, fmt.Sprintf("Could not unmarshal JSON payload: %s", err), http.StatusBadRequest)
		return
	}

	// Create the new account.
	account, err := a.dbAdapter.CreateAccount(data.DocumentNumber)
	if err != nil {
		a.Error(w, fmt.Sprintf("Could not persist new account: %s", err), http.StatusInternalServerError)
		return
	}

	log.Printf("New Account ID: %d, Document Number: %s", account.ID, account.DocumentNumber)

	// Construct and send back the response.
	resp := CreateAccountResponse{ID: account.ID, DocumentNumber: account.DocumentNumber}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		a.Error(w, fmt.Sprintf("Could not serialize response: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(respBytes))
}

// GetAccount is the HTTP handler function for fetching an account.
func (a *Handler) GetAccount(w http.ResponseWriter, req *http.Request) {
	// Only accept POST.
	if req.Method != http.MethodPost {
		a.Error(w, fmt.Sprintf("%s not allowed", req.Method), http.StatusMethodNotAllowed)
		return
	}

	// Read the request payload.
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		a.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var data CreateAccountRequest
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		a.Error(w, fmt.Sprintf("Could not unmarshal JSON payload: %s", err), http.StatusBadRequest)
		return
	}

	// Create the new account.
	account, err := a.dbAdapter.CreateAccount(data.DocumentNumber)
	if err != nil {
		a.Error(w, fmt.Sprintf("Could not persist new account: %s", err), http.StatusInternalServerError)
		return
	}

	log.Printf("New Account ID: %d, Document Number: %s", account.ID, account.DocumentNumber)

	// Construct and send back the response.
	resp := CreateAccountResponse{ID: account.ID, DocumentNumber: account.DocumentNumber}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		a.Error(w, fmt.Sprintf("Could not serialize response: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(respBytes))
}
