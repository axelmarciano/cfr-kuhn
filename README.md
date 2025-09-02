# Kuhn Poker + CFR 🎲

I built this little project to learn how **Counterfactual Regret Minimization (CFR)** works in practice.  
Instead of going straight into huge poker solvers, I wanted to start small with **Kuhn Poker**, which is a super tiny poker game that still has bluffing, betting, and strategy.

---

## What’s Kuhn Poker? 🃏
- 2 players
- Deck of 3 cards: Jack, Queen, King
- Each player antes 1 chip
- Each gets 1 private card (the 3rd card is hidden)
- Simple betting round:
  - First player can check or bet 1 chip
  - If first player checks, the second can check (go to showdown) or bet
  - If a bet happens, the other can call or fold
- Showdown: higher card wins the pot

---

## What’s CFR? 🤖
CFR = **Counterfactual Regret Minimization**.  
It’s an algorithm that:
- Keeps track of how much you “regret” not taking certain actions
- Increases the chances of actions that had positive regret
- Decreases the chances of actions with negative regret
- Over many iterations, the average strategy converges towards a **Nash equilibrium**

---
## Results after 5M iterations 📊
*(The overall regret did not reach the 1e-4 threshold.)*

### Player 1 (P1)

| InfoSet   | Action probabilities       | Note |
|-----------|----------------------------|------|
| Q         | check **100%**, bet 0%     | Q never bets at equilibrium |
| J         | bet ~22%, check ~78%       | Bluffs ~22% with J (≤ 1/3, consistent) |
| K         | bet ~66%, check ~34%       | Value-bets ~66% (≈ 3 × bluff freq) |
| Q \| cb   | call ~55%, fold ~45%       | Calls with Q ~α+1/3 (here 0.22+0.33 ≈ 0.55) |
| J \| cb   | fold **100%**              | J always folds vs bet (can’t win) |
| K \| cb   | call **100%**              | K always calls (best hand) |

### Player 2 (P2)

| InfoSet   | Action probabilities       | Note |
|-----------|----------------------------|------|
| J \| c    | bet ~33%, check ~67%       | Semi-bluff with J 1/3 of the time |
| J \| b    | fold **100%**              | J folds if raised |
| Q \| c    | check **100%**             | Q checks, never bets itself |
| Q \| b    | call ~33%, fold ~67%       | Defends vs bluff by calling 1/3 |
| K \| c    | bet **100%**               | K always value-bets |
| K \| b    | call **100%**              | K always calls vs bet |

---

### 🧐 Interpretation
- **Bluff/value balance:** P1 bluffs ~22% with J and value-bets ~66% with K → ratio matches theory.
- **Defense:** P2 calls with Q only 1/3 of the time, which prevents P1 from bluffing too much.
- **Consistency:** Every deviation (e.g. J folding always, K betting/calling always) matches the known Nash equilibrium of Kuhn Poker.
---

## References I Used 📚
- [Counterfactual Regret Minimization (Zinkevich et al., 2007)](https://poker.cs.ualberta.ca/publications/NIPS07-cfr.pdf) — the original paper
- [Regret Minimization in Games with Incomplete Information (Neller & Lanctot, 2013)](https://arxiv.org/abs/1305.0023) — easier overview



