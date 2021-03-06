package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/abhishekunotech/Hydra-Pearls/utils"
	"github.com/abhishekunotech/Hydra-Pearls/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)
// This Type defines the Input JSON Body for creating a Pickup Ticket Request
//
// It includes ticket parameters
type Pickup_Request struct{
	UserLogin	string	`json:"UserLogin"`
	Password	string	`json:"Password"`
	Action		string	`json:"Action"`
	TicketIDs	[]string	`json:"TicketIDs"`
	Subaction	string	`json:"Subaction"`
	ArticleType	string	`json:"ArticleType"`
	Unlock		string	`json:"Unlock"`
	OwnerID		string	`json:"OwnerID"`
	Subject		string	`json:"Subject"`
	Body		string	`json:"Body"`
}
// This function is a handler that creates a POST request to assign an unassigned ticket to yourself.
//
// **Business Logic**: Function takes as an input a JSON Body and uses the Ticket details in Request Body to generate response
//
// Returns data as shown in examples
func (h *Handler) PostPickupAgentTicket(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t Pickup_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(pickupAgentTicket(t),w)
}

func pickupAgentTicket(T Pickup_Request) []uint8 {

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.postpickupagentticket").GetString("uri")
	url := felicitybaseurl + felicityapiuri  
	j, err := json.Marshal(T)

	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)

	return utils.MakeHTTPPostCall(url,b)
}
