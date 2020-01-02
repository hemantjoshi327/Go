package main

import (
"encoding/json"
"fmt"
"net/http"
"github.com/gorilla/mux"
)

func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		user, pass, ok := r.BasicAuth()
		fmt.Println("username: ", user)
		fmt.Println("password: ", pass)
		fmt.Println("OK:", ok)
		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(w, r)
	}
}

func checkUsernameAndPassword(username, password string) bool {
	return username == "abc" && password == "123"
}

type Confirmation struct {
Message string
MessageId string
SubscribeURL string
Timestamp string
Token string
TopicArn string
Type string
}

type Notification struct {
Type string
MessageId string
TopicArn string
Subject string
Message string
Timestamp string
SignatureVersion string
Signature string
SigningCertURL string
UnsubscribeURL string
}

func DemoAPI(response http.ResponseWriter, request *http.Request) {
	//request.Body.Read()
	decoder := json.NewDecoder(request.Body)
	messageType := request.Header.Get("x-amz-sns-message-type")
	fmt.Println(messageType)

	if messageType == "SubscriptionConfirmation" {
	var c Confirmation
	err := decoder.Decode(&c)
		if err != nil {
			panic(err)
		}
	fmt.Println(c)
	} else if messageType == "Notification" {

		var n Notification
		err := decoder.Decode(&n)
			if err != nil {
				panic(err)
			}
		fmt.Println(n)
	}
	fmt.Fprint(response, "Demo API")
}

func main() {
	router := mux.NewRouter()
	router.Handle("/api/demo", BasicAuthMiddleware(http.HandlerFunc(DemoAPI))).Methods("POST")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}