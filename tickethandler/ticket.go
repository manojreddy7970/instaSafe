package tickethandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Notes struct {
	Note_2000 int `json:"note_2000"`
	Note_500  int `json:"note_500"`
	Note_100  int `json:"note_100"`
	Note_50   int `json:"note_50"`
	Note_20   int `json:"note_20"`
	Note_10   int `json:"note_10"`
	Note_5    int `json:"note_5"`
}

type Coins struct {
	Coin_10 int `json:"coin_10"`
	Coin_5  int `json:"coin_5"`
	Coin_2  int `json:"coin_2"`
	Coin_1  int `json:"coin_1"`
}

type PFTicket struct {
	Ticket int `json:"tickets"`
	Notes
	Coins
}

type Response struct {
	Ticket int `json:"tickets"`
	Coins
}

var response Response

func EvaluatePFTHandler(w http.ResponseWriter, r *http.Request) {
	platformTicket := PFTicket{}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Sscan(r.Form.Get("tickets"), &platformTicket.Ticket)
	fmt.Sscan(r.Form.Get("note_2000"), &platformTicket.Note_2000)
	fmt.Sscan(r.Form.Get("note_500"), &platformTicket.Note_500)
	fmt.Sscan(r.Form.Get("note_100"), &platformTicket.Note_100)
	fmt.Sscan(r.Form.Get("note_50"), &platformTicket.Note_50)
	fmt.Sscan(r.Form.Get("note_20"), &platformTicket.Note_20)
	fmt.Sscan(r.Form.Get("note_10"), &platformTicket.Note_10)
	fmt.Sscan(r.Form.Get("note_5"), &platformTicket.Note_5)
	fmt.Sscan(r.Form.Get("coin_10"), &platformTicket.Coin_10)
	fmt.Sscan(r.Form.Get("coin_5"), &platformTicket.Coin_5)
	fmt.Sscan(r.Form.Get("coin_2"), &platformTicket.Coin_2)
	fmt.Sscan(r.Form.Get("coin_1"), &platformTicket.Coin_1)

	totalMoney := platformTicket.Note_2000*2000 + platformTicket.Note_500*500 +
		platformTicket.Note_100*100 + platformTicket.Note_50*50 + platformTicket.Note_20*20 +
		platformTicket.Note_10*10 + platformTicket.Note_5*5 + platformTicket.Coin_10*10 +
		platformTicket.Coin_5*5 + platformTicket.Coin_2*2 + platformTicket.Coin_1

	fmt.Println(totalMoney)

	totalChange := totalMoney - platformTicket.Ticket

	if totalChange >= 0 {
		totalChangeInCoins := Coins{}

		totalChangeInCoins.Coin_10 = totalChange / 10
		reminder := totalChange % 10
		totalChangeInCoins.Coin_5 = reminder / 5
		reminder %= 5
		totalChangeInCoins.Coin_2 = reminder / 2
		reminder %= 2
		totalChangeInCoins.Coin_1 = reminder

		response = Response{Ticket: platformTicket.Ticket, Coins: totalChangeInCoins}
	} else {
		response = Response{Ticket: totalMoney, Coins: Coins{}}
	}
	
	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func GetBalancePFTHandler(w http.ResponseWriter, r *http.Request) {
	balanceCoinsResponse, err := json.Marshal(response)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(balanceCoinsResponse)
	response = Response{}
}
