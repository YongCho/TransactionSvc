package rest

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	txn "pismo.io/transaction"
	"pismo.io/util"
)

// Handler implements the REST API handler logic.
type Handler struct {
	txnMgr *txn.TransactionMgr
}

// GenericResponse is the response payload for any requests
// that do not return specific data.
type GenericResponse struct {
	Status string `json:",omitempty"`
	Error  string `json:",omitempty"`
}

// NewHandler creates a new instance of Handler.
func NewHandler(txnMgr *txn.TransactionMgr) (*Handler, error) {
	return &Handler{
		txnMgr: txnMgr,
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
	account, err := a.txnMgr.CreateAccount(data.DocumentNumber)
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

// GetAccount is the HTTP handler function for fetching an existing account.
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
	account, err := a.txnMgr.GetAccount(data.ID)
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
		CreatedAt       string
	}

	// Read the JSON payload.
	var data Req
	err := c.BindJSON(&data)
	if err != nil {
		resp := GenericResponse{Error: fmt.Sprintf("Could not process payload: %s", err)}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Create the transaction.
	transaction, err := a.txnMgr.CreateTransaction(data.AccountID, txn.OpType(data.OperationTypeID), data.Amount)
	if err != nil {
		resp := GenericResponse{Error: fmt.Sprintf("Could not persist the transaction: %s", err)}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	createdAtStr := transaction.CreatedAt.Format(time.RFC3339)
	log.Printf("New Transaction ID: %d, Account ID: %d, Operation Type: %d, Amount: %d (cents), CreatedAt: %s",
		transaction.ID, transaction.AccountID, transaction.OperationTypeID, transaction.Amount, createdAtStr)

	// Construct and send back the response.
	dollar := util.CentsToDollar(transaction.Amount)
	resp := Resp{
		ID:              transaction.ID,
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          dollar,
		CreatedAt:       createdAtStr,
	}
	c.JSON(http.StatusOK, resp)
}
