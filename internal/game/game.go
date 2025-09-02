package game

import "fmt"

var infoSets map[string]*InfoSet

func computeInfoSetKey(s State) string {
	key := ""
	if s.Player == Player1 {
		key += fmt.Sprintf("%s:%s", "P1", s.Deal.P1)
	} else {
		key += fmt.Sprintf("%s:%s", "P2", s.Deal.P2)
	}
	if s.History == "" {
		return key
	}
	return fmt.Sprintf("%s|%s", key, s.History)
}

func computeAvailableActions(s State) []Action {
	if s.History == "" {
		return []Action{Check, Bet}
	} else if s.History == "c" {
		return []Action{Check, Bet}
	} else if s.History == "b" {
		return []Action{Call, Fold}
	} else if s.History == "cb" {
		return []Action{Call, Fold}
	}
	return []Action{}
}

func IsTerminalState(s State) bool {
	if s.History == "cc" || s.History == "bk" || s.History == "cbf" || s.History == "bf" || s.History == "cbk" {
		return true
	}
	return false
}

func getShowdownWinner(s State) Player {
	if s.Deal.P1 == Jack {
		return Player2
	}
	if s.Deal.P1 == Queen {
		if s.Deal.P2 == Jack {
			return Player1
		}
		return Player2
	}
	return Player1
}

func isShowdown(s State) bool {
	return s.History == "cc" || s.History == "bk" || s.History == "cbk"
}

func whoFold(s State) Player {
	if s.History == "bf" {
		return Player2
	}
	if s.History == "cbf" {
		return Player1
	}
	return 0
}

func ComputePayOff(s State) int {
	if !IsTerminalState(s) {
		return 0
	}
	isShowdownState := isShowdown(s)
	if isShowdownState {
		hasWonShowdown := getShowdownWinner(s) == Player1
		if s.History == "cc" {
			if hasWonShowdown {
				return 1
			}
			return -1
		}
		if hasWonShowdown {
			return 2
		} else {
			return -2
		}
	}
	folderedPlayer := whoFold(s)
	if folderedPlayer == Player1 {
		return -1
	}
	return 1
}

func ComputeInfoSet(s State) *InfoSet {
	key := computeInfoSetKey(s)
	actions := computeAvailableActions(s)
	if infoSet, exists := infoSets[key]; exists {
		return infoSet
	}
	regretSum := make(map[Action]float64)
	strategySum := make(map[Action]float64)
	for _, action := range actions {
		regretSum[action] = 0.0
		strategySum[action] = 0.0
	}
	infoSet := &InfoSet{
		Key:         key,
		Actions:     actions,
		RegretSum:   regretSum,
		StrategySum: strategySum,
	}
	infoSets[key] = infoSet
	return infoSet
}

var AllDeals = []Deal{
	{P1: Jack, P2: Queen},
	{P1: Jack, P2: King},
	{P1: Queen, P2: Jack},
	{P1: Queen, P2: King},
	{P1: King, P2: Jack},
	{P1: King, P2: Queen},
}

var NextPlayer = map[Player]Player{Player1: Player2, Player2: Player1}

func ComputeAllInfoSets() map[string]*InfoSet {
	infoSets = make(map[string]*InfoSet)

	var visit func(s State)
	visit = func(s State) {
		if IsTerminalState(s) {
			return
		}
		ComputeInfoSet(s)
		actions := computeAvailableActions(s)
		for _, action := range actions {
			nextState := s
			nextState.History = fmt.Sprintf("%s%s", nextState.History, action)
			nextState.Player = NextPlayer[nextState.Player]
			visit(nextState)
		}
	}
	for _, deal := range AllDeals {
		initialState := State{
			History: "",
			Player:  Player1,
			Deal:    deal,
		}
		visit(initialState)
	}
	return infoSets
}

func AverageStrategy(is *InfoSet) map[Action]float64 {
	avg := make(map[Action]float64)
	normalizing := 0.0
	for _, a := range is.Actions {
		normalizing += is.StrategySum[a]
	}
	for _, a := range is.Actions {
		if normalizing > 0 {
			avg[a] = is.StrategySum[a] / normalizing
		} else {
			avg[a] = 1.0 / float64(len(is.Actions))
		}
	}
	return avg
}
