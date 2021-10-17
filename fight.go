package main

type FightResult struct {
	Winner *Fighter
	Loser  *Fighter
}

func Fight(f1, f2 *Fighter) FightResult {

	for {
		f1A := f1.Attack()
		f2E := f2.Evade()

		if f1A > f2E {
			f2.DealDamage(f1A)
		}

		f2A := f1.Attack()
		f1E := f2.Evade()

		if f2A > f1E {
			f1.DealDamage(f2A)
		}

		if f1.IsDead() {
			return FightResult{
				Winner: f2,
				Loser:  f1,
			}
		}

		if f2.IsDead() {
			return FightResult{
				Winner: f1,
				Loser:  f2,
			}
		}
	}
}
