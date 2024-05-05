package api

import (
	"encoding/JSON"
	"net/http"
)

// this file contains the parameters and response structures



type CoinBalanceParams struct{
	Username string

}


type CoinBalanceResponse struct{
	Code int //status code
	Balance int
}

type Error struct {
	Code int

	Message string
}


