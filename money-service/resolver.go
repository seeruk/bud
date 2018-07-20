package moneyapi

import "github.com/seeruk/bud/bud"

type Resolver struct{
	budResolver bud.Resolver
}

func NewResolver(budResolver bud.Resolver) *Resolver {
	return &Resolver{
		budResolver: budResolver,
	}
}
