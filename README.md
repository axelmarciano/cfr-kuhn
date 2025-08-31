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

Today, CFR is the foundation of all **GTO (Game Theory Optimal) solvers**.  
Modern poker solvers (like PioSolver, GTO Wizard, Simple Postflop) usually rely on an improved version called **CFR+**  

👉 For learning and small games like **Kuhn Poker**, plain CFR is totally enough.  
CFR+ is only needed for **scaling up** to huge game trees like No-Limit Hold’em.

---

## References I Used 📚
- [Counterfactual Regret Minimization (Zinkevich et al., 2007)](https://poker.cs.ualberta.ca/publications/NIPS07-cfr.pdf) — the original paper
- [Regret Minimization in Games with Incomplete Information (Neller & Lanctot, 2013)](https://arxiv.org/abs/1305.0023) — easier overview
- [Solving Large Imperfect Information Games Using CFR+](https://arxiv.org/pdf/1407.5042).

---

## Why I Did It 🚀
Mostly for learning and curiosity.  
I like seeing abstract math (Nash equilibria, regrets, convergence) come alive in code.  
Kuhn Poker is small enough that you can actually **watch a solver learn to bluff** in real time. That’s pretty cool 😎
