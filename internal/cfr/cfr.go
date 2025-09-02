package cfr

import (
	"bufio"
	"cfr-kuhn/internal/game"
	"fmt"
	"math"
	"os"
	"strings"
)

func computeStrategyFromRegrets(infoset *game.InfoSet) map[game.Action]float64 {
	strategy := make(map[game.Action]float64)
	regretSum := 0.0
	for _, a := range infoset.Actions {
		if infoset.RegretSum[a] > 0 {
			regretSum += infoset.RegretSum[a]
		}
	}
	for _, action := range infoset.Actions {
		if regretSum > 0 {
			strategy[action] = math.Max(infoset.RegretSum[action], 0) / regretSum
		} else {
			strategy[action] = 1.0 / float64(len(infoset.Actions))
		}
	}
	return strategy
}

func nextState(s game.State, a game.Action) game.State {
	return game.State{
		History: fmt.Sprintf("%s%s", s.History, a),
		Player:  game.NextPlayer[s.Player],
		Deal:    s.Deal,
	}
}

func CFR(s game.State, reachP1, reachP2, reachC float64) float64 {
	if game.IsTerminalState(s) {
		return float64(game.ComputePayOff(s))
	}
	infoset := game.ComputeInfoSet(s)
	strategy := computeStrategyFromRegrets(infoset)
	actionUtil := make(map[game.Action]float64)
	nodeUtil := 0.0
	for a, p := range strategy {
		ns := nextState(s, a)
		if s.Player == game.Player1 {
			actionUtil[a] = CFR(ns, reachP1*p, reachP2, reachC)
		} else {
			actionUtil[a] = CFR(ns, reachP1, reachP2*p, reachC)
		}
		nodeUtil += p * actionUtil[a]
	}
	reachProbability := reachC
	if s.Player == game.Player1 {
		reachProbability = reachProbability * reachP2
	} else {
		reachProbability = reachProbability * reachP1
	}
	sign := 1.0
	if s.Player == game.Player2 {
		sign = -1.0
	}
	for _, action := range infoset.Actions {
		immediateRegret := reachProbability * (sign * (actionUtil[action] - nodeUtil))
		infoset.RegretSum[action] += immediateRegret
		if s.Player == game.Player2 {
			infoset.StrategySum[action] += reachP2 * strategy[action]
		} else {
			infoset.StrategySum[action] += reachP1 * strategy[action]
		}
	}
	return nodeUtil
}

func launchIteration() {
	for _, deal := range game.AllDeals {
		root := game.State{
			History: "",
			Player:  game.Player1,
			Deal:    deal,
		}
		CFR(root, 1, 1, 1.0/float64(len(game.AllDeals)))
	}
}

var out = bufio.NewWriter(os.Stderr)

func computeOverallRegret(infosets map[string]*game.InfoSet, round int) float64 {
	if round == 0 {
		return math.Inf(1)
	}
	globalRegret := 0.0
	for _, infoset := range infosets {
		maxRegret := 0.0
		for _, r := range infoset.RegretSum {
			if r > maxRegret {
				maxRegret = r
			}
		}
		globalRegret += maxRegret
	}
	return globalRegret / float64(round)
}

func progressBar(label string, value, max float64, width int) string {
	ratio := value / max
	if ratio > 1 {
		ratio = 1
	}
	filled := int(ratio * float64(width))
	return fmt.Sprintf("%s [%s%s] %6.2f%%",
		label,
		strings.Repeat("=", filled),
		strings.Repeat(" ", width-filled),
		ratio*100,
	)
}

func Launch() {
	const maxRounds = 5_000_000
	const threshold = 1e-4
	const barWidth = 40

	infosets := game.ComputeAllInfoSets()
	round := 0

	for {
		launchIteration()
		round++

		overall := computeOverallRegret(infosets, round)
		progressTowards := 0.0
		if overall > 0 && !math.IsInf(overall, 0) {
			progressTowards = threshold / overall
			if progressTowards > 1 {
				progressTowards = 1
			}
		}

		if round%1000 == 0 || round == 1 {
			fmt.Fprintf(out, "%s  %s  curr=%.6g thr=%.6g\n",
				progressBar("Rounds", float64(round), float64(maxRounds), barWidth),
				progressBar("Regret", progressTowards, 1, barWidth),
				overall, threshold,
			)
			fmt.Fprint(out, "\x1b[1A")
			out.Flush()
		}

		if round >= maxRounds || overall < threshold {
			fmt.Fprintln(out)
			out.Flush()
			break
		}
	}

	for _, infoset := range infosets {
		strategy := game.AverageStrategy(infoset)
		fmt.Printf("Strategy for %s\n", infoset.Key)
		for action, strat := range strategy {
			fmt.Printf("  Action %s: %f\n", action, strat)
		}
	}
}
