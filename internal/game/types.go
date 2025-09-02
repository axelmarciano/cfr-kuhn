package game

type Card string

const (
	Jack  Card = "J"
	Queen Card = "Q"
	King  Card = "K"
)

type Action string

const (
	Check Action = "c"
	Bet   Action = "b"
	Fold  Action = "f"
	Call  Action = "k"
)

type InfoSet struct {
	Key         string
	Actions     []Action
	RegretSum   map[Action]float64
	StrategySum map[Action]float64
}

type Player int

const (
	Player1 Player = 1
	Player2 Player = 2
)

type Deal struct{ P1, P2 Card }

type State struct {
	History string
	Player  Player
	Deal    Deal
}
