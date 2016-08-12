package graph


type Graph64 map[uint64]map[uint64]*interface{}

func (g Graph64) GetArrow (domain uint64, codomain uint64) (*interface{}, bool) {
	if codomains, ok := g[domain]; ok {
		return codomains[codomain]
	} else {
		return nil, false
	}
}


func (g Graph64) SetArrow (domain uint64, codomain uint64, ident *interface{}) {
	g[domain][codomain] = ident
}


func (g Graph64) DeleteArrow (domain uint64, codomain uint64) {
	if codomains, ok := g[domain]; ok {
		delete(codomains, codomain)
		if len(codomains) == 0 {
			delete(g, domain)
		}
	}
}


func (g Graph64) MapReduce (mr MR, target *Graph64) error {
	for domain, codomains := range g {
		mapped_domain := mr.Map(domain)
		for codomain, ident := range codomains {
			mapped_codomain := mr.Map(codomain)
			acc, isPopulated := (*target).GetArrow(mapped_domain, mapped_codomain)
			if isPopulated {
				acc = mr.Reduce(acc, ident)
			} else {
				acc= ident
			}
			(*target).SetArrow(mapped_domain, mapped_codomain, acc)
		}
	}
}