package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/abhishekunotech/Hydra-Pearls/utils"
	"github.com/abhishekunotech/Hydra-Pearls/lerna"
	"github.com/antigloss/go/logger"
	"net/http"
)


type CustomerArticleReply_Request struct{
	Password	        string	`json:"Password,omitempty"`
	CustomerUserLogin	string	`json:"CustomerUserLogin,omitempty"`
	ArticleID	        int	    `json:"ArticleID,omitempty"`
	ArticleState         string `json:"ArticleState,omitempty"`
    Body                string `json:"Body,omitempty"`
    Subject             string `json:"Subject,omitempty"`
    TicketNumber        int    `json:"TicketNumber,omitempty"`
    StateID             int    `json:"StateID,omitempty"`
    PriorityID             int    `json:"PriorityID,omitempty"`
	Attachment	    *[]CAR_Attachment	`json:"Attachment,omitempty"`
	DynamicField	*[]CAR_DynamicField	`json:"DynamicField,omitempty"`
}

type CAR_Attachment struct{
	Content	string	`json:"Content,omitempty"`
	ContentType	string	`json:"ContentType,omitempty"`
	Filename	string	`json:"Filename,omitempty"`

}

type CAR_DynamicField	struct{
	Name	string	`json:"Name,omitempty"`
	Value	string	`json:"Value,omitempty"`
}

// This function is a handler that creates a POST request to update Ticket Common Agent Functions
//
// **Business Logic**: Function takes as an input a JSON Body and generates the response
//
// Returns data as shown in examples
func (h *Handler) PostCustomerArticleReply(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t  CustomerArticleReply_Request
	err := decoder.Decode(&t)
	if err != nil {
		logger.Error("Error Occured in Decoding Post Request")
		logger.Error(err.Error())
	}
	defer r.Body.Close()

	utils.ResponseAbstract(postCustomerArticleReply(t),w)
}

func postCustomerArticleReply(T CustomerArticleReply_Request) []uint8{

	//Get Lerna Running
	ConfObj := lerna.ReturnConfigObject()

	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.postcustomerarticlereply").GetString("uri")

	url := felicitybaseurl + felicityapiuri 
	j, err := json.Marshal(T)
//	logger.Info(T)
	if err != nil {
		logger.Error("Error in Marshaling JsON")
		logger.Error(err.Error())
	}

	b := bytes.NewBuffer(j)

	return utils.MakeHTTPPostCall(url,b)

}
