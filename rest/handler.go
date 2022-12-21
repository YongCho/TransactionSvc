package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"pismo.io/db"
	"pismo.io/util"
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

// GetAccount is the HTTP handler function for fetching an account.
func (a *Handler) GetAccount(c *gin.Context) {
	type Req struct {
		ID int32 `uri:"id" binding:"required"`
	}

	type Resp struct {
		ID             int32
		DocumentNumber string `json:"document_number"`
	}

	// Read the account ID in URL.
	var data Req
	err := c.BindUri(&data)
	if err != nil {
		resp := GenericResponse{Error: fmt.Sprintf("Could not parse account ID: %s", err)}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Get the account.
	account, err := a.dbAdapter.GetAccount(data.ID)
	if err != nil {
		resp := GenericResponse{Error: fmt.Sprintf("Could not fetch the account: %s", err)}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// Construct and send back the response.
	resp := Resp{ID: account.ID, DocumentNumber: account.DocumentNumber}
	c.JSON(http.StatusOK, resp)
}

// CreateTransaction is the HTTP handler function creating a new transaction.
func (a *Handler) CreateTransaction(c *gin.Context) {
	type Req struct {
		AccountID       int32 `json:"account_id"`
		OperationTypeID int32 `json:"operation_type_id"`
		Amount          float64
	}

	type Resp struct {
		ID              int32
		AccountID       int32 `json:"account_id"`
		OperationTypeID int32 `json:"operation_type_id"`
		Amount          float64
	}

	// Read the JSON payload.
	var data Req
	err := c.BindJSON(&data)
	if err != nil {
		resp := GenericResponse{Error: fmt.Sprintf("Could not process payload: %s", err)}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Create the transaction. Store the money in cents to avoid float precision issue.
	cents := util.DollarToCents(data.Amount)
	transaction, err := a.dbAdapter.CreateTransaction(data.AccountID, data.OperationTypeID, cents)
	if err != nil {
		resp := GenericResponse{Error: fmt.Sprintf("Could not persist the transaction: %s", err)}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	log.Printf("New Transaction ID: %d, Account ID: %d, Operation Type: %d, Amount: %d (cents)",
		transaction.ID, transaction.AccountID, transaction.OperationTypeID, transaction.Amount)

	// Construct and send back the response.
	dollar := util.CentsToDollar(transaction.Amount)
	resp := Resp{
		ID:              transaction.ID,
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          dollar,
	}
	c.JSON(http.StatusOK, resp)
}
