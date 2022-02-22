package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

//  event struct store information about the events available on the exchange.
type event struct {
	id       string
	name     string
	amount   float64
	outcomes map[string]outcome
}

// outcome struct describe the total amount bet on particular outcome
type outcome struct {
	id     string
	name   string
	amount float64
	payout float64
}

// player struct holds information about the order submitted by the user
type player struct {
	id        string
	amount    float64
	odd       float64
	eventID   string
	outcomeID string
	created   time.Time
}

type order struct {
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	fmt.Println("STARTING EXCHANGE...")

	// GENERATE THE EVENT
	ev := event{
		id:     "event-1",
		name:   "5K",
		amount: 0.0,
	}

	// GENERATE OUTCOME
	winOutcome := outcome{
		id:     "out-1",
		name:   "WIN",
		amount: 0.0,
		payout: 0.0,
	}

	lossOutcome := outcome{
		id:     "out-2",
		name:   "LOSS",
		amount: 0.0,
		payout: 0.0,
	}

	//  add the out come
	ev.outcomes = make(map[string]outcome)
	ev.outcomes[winOutcome.id] = winOutcome
	ev.outcomes[lossOutcome.id] = lossOutcome

	fmt.Println("PLAYER BET SIMULATION...")
	playerList := simulationStart()

	for i := range playerList {
		// out come calculation
		st := ev.outcomes[playerList[i].outcomeID]
		st.amount += playerList[i].amount

		st.payout += playerList[i].amount * playerList[i].odd

		ev.outcomes[playerList[i].outcomeID] = st

		//  event calculation
		ev.amount += playerList[i].amount
	}

	ev.amount = math.Floor(ev.amount*100) / 100
	for i := range ev.outcomes {
		st := ev.outcomes[i]
		st.amount = math.Floor(st.amount*100) / 100

		st.payout = math.Floor(st.payout*100) / 100

		ev.outcomes[i] = st
	}

	fmt.Println(ev)

}

func simulationStart() []player {
	playerList := make([]player, 0)
	//  generate player random id , name , amount and odd

	for i := 0; i < 1000; i++ {
		playerList = append(playerList, player{
			id:        genRandomId(),
			amount:    getRandomStake(0.01, 10.99), // 0.00 - 10.99
			odd:       getGetRandomOdd([]float64{1.5, 2.5}),
			eventID:   "event-1",                                       // 1.5 or 2.5
			outcomeID: getGetRandomOutcome([]string{"out-1", "out-2"}), // out-1 or out-2
			created:   time.Now(),
		})
	}

	return playerList
}

func genRandomId() string {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Panic(err)
	}
	return u.String()
}

func getRandomStake(min, max float64) float64 {
	return math.Floor(((rand.Float64()*(max-min+1))+min)*100) / 100
}

func getGetRandomOdd(odd []float64) float64 {

	return odd[rand.Intn(len(odd))]
}

func getGetRandomOutcome(out []string) string {

	return out[rand.Intn(len(out))]
}
