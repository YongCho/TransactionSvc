package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"pismo.io/db"
)

// Handler implements the HTTP handler logic.
type Handler struct {
	dbAdapter *db.DBAdapter
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

// CreateAccount is the HTTP handler function for creating a new account.
func (a *Handler) CreateAccount(c *gin.Context) {
	type Req struct {
		DocumentNumber string `json:"document_number"`
	}

	type Resp struct {
		ID             int32
		DocumentNumber string `json:"document_number"`
	}

	// Read the JSON payload.
	var data Req
	err := c.BindJSON(&data)
	if err != nil {
		resp := GenericResponse{Error: fmt.Sprintf("Could not process payload: %s", err)}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Create the new account.
	account, err := a.dbAdapter.CreateAccount(data.DocumentNumber)
	if err != nil {
		resp := GenericResponse{Error: fmt.Sprintf("Could not persist new account: %s", err)}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	log.Printf("New Account ID: %d, Document Number: %s", account.ID, account.DocumentNumber)

	// Construct and send back the response.
	resp := Resp{ID: account.ID, DocumentNumber: account.DocumentNumber}
	c.JSON(http.StatusOK, resp)
}
