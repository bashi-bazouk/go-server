package main

type Composable interface {
	Compose(Composable) Composable
}


type Monad []Composable

func (l Monad) Compose (r Monad) Monad {
	return append(l, r...)
}





type Applicative interface {
	Compose(Composable) Applicative
	Apply(Applicative) Composable
}


type FreeApplicative []Monad

func (l FreeApplicative) Compose (r Composable) FreeApplicative {
	return nil
}

func (l FreeApplicative) Apply (r FreeApplicative) FreeApplicative {
	res := l
	for i, m := range r {
		if i >= len(m) {
			append(res, m)
		}
		res[i] = l[i].Compose(r[i])
	}
	return res
}

