package handlers

import (
	"github.com/abhishekunotech/Hydra-Pearls/utils"
	"net/http"
)


// This function is a handler that creates a GET API to get all the Items Linked to a Ticket.
//
// **Business Logic**: Function takes as an input GET Parameter, __ticketID__ that will identify a ticket and return all the Items attached to it.
//
// Returns data as shown in examples.
func (h *Handler) GetLinkedItems(w http.ResponseWriter, r *http.Request) {
	actionStrg := utils.RequestAbstractGet1(r)
	configStrg := "components.otrs"
	uriStrg := "components.otrs.apis.GetLinkedItems"
	utils.ResponseAbstract(utils.ExecuteCallGet(configStrg,uriStrg,actionStrg),w)

}
