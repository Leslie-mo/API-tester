package handlers

import (
	"log"
	"net/http"
)

func InvokeHandle(w http.ResponseWriter, r *http.Request) {

	handFlg, errMsg, httpstatus := RequestHandler(w, r)

	if handFlg {
		log.Println("The message was successfully sent and received.")

	} else {
		log.Println("Failed to receive message.")
		FailedRequest(w, errMsg, httpstatus)
	}

}
