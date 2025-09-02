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

Here’s the **average strategy** learned by CFR in my implementation for 5M iterations (The overall regret did not reach the 1e-4 threshold.)

### Player 1 (P1)

| InfoSet   | Action probabilities                  |
|-----------|---------------------------------------|
| Q         | check **100%**, bet 0%                |
| J         | bet ~22%, check ~78%                  |
| K         | bet ~66%, check ~34%                  |
| Q \| cb   | call ~55%, fold ~45%                  |
| J \| cb   | fold **100%**                         |
| K \| cb   | call **100%**                         |

### Player 2 (P2)

| InfoSet   | Action probabilities                  |
|-----------|---------------------------------------|
| J \| c    | bet ~33%, check ~67%                  |
| J \| b    | fold **100%**                         |
| Q \| c    | check **100%**                        |
| Q \| b    | call ~33%, fold ~67%                  |
| K \| c    | bet **100%**                          |
| K \| b    | call **100%**                         |


### 🧐 Interpretation
- **Bluffing:** P1 bluffs ~22% of the time with J, and value-bets ~66% with K.
- **Consistency:** This matches the theoretical equilibrium where bluff frequency with J is α and value-bet with K is 3α. Here α ≈ 0.22.
- **Defense:** P2 defends by calling with Q about 1/3 of the time, and bluffing with J 1/3 after a check.
- **Equilibrium:** These strategies line up with the published Nash equilibrium of Kuhn Poker (see Neller & Lanctot 2013).

👉 Overall, CFR converged correctly towards the known Nash equilibrium, where Player 1’s expected value is about **-1/18 ≈ -0.0556**.

---

## References I Used 📚
- [Counterfactual Regret Minimization (Zinkevich et al., 2007)](https://poker.cs.ualberta.ca/publications/NIPS07-cfr.pdf) — the original paper
- [Regret Minimization in Games with Incomplete Information (Neller & Lanctot, 2013)](https://arxiv.org/abs/1305.0023) — easier overview



