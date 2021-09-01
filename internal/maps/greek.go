package maps

func init() {
	if maps == nil {
		maps = make(map[string]*Map)
	}

	maps[Greek.Name] = &Greek
}

var Greek = Map{
	Name: "greek",

	Chars: map[string]*Char{
		"alpha": {
			U: 'Α',
			L: 'α',
		},
		"beta": {
			U: 'Β',
			L: 'β',
		},
		"gamma": {
			U: 'Γ',
			L: 'γ',
		},
		"delta": {
			U: 'Δ',
			L: 'δ',
		},
		"epsilon": {
			U: 'Ε',
			L: 'ε',
		},
		"zeta": {
			U: 'Ζ',
			L: 'ζ',
		},
		"eta": {
			U: 'Η',
			L: 'η',
		},
		"theta": {
			U: 'Θ',
			L: 'θ',
		},
		"iota": {
			U: 'Ι',
			L: 'ι',
		},
		"kappa": {
			U: 'Κ',
			L: 'κ',
		},
		"lambda": {
			U: 'Λ',
			L: 'λ',
		},
		"mu": {
			U: 'Μ',
			L: 'μ',
		},
		"nu": {
			U: 'Ν',
			L: 'ν',
		},
		"xi": {
			U: 'Ξ',
			L: 'ξ',
		},
		"omicron": {
			U: 'Ο',
			L: 'ο',
		},
		"pi": {
			U: 'Π',
			L: 'π',
		},
		"rho": {
			U: 'Ρ',
			L: 'ρ',
		},
		"sigma": {
			U: 'Σ',
			L: 'σ',
		},
		"tau": {
			U: 'Τ',
			L: 'τ',
		},
		"upsilon": {
			U: 'Υ',
			L: 'υ',
		},
		"phi": {
			U: 'Φ',
			L: 'φ',
		},
		"chi": {
			U: 'Χ',
			L: 'χ',
		},
		"psi": {
			U: 'Ψ',
			L: 'ψ',
		},
		"omega": {
			U: 'Ω',
			L: 'ω',
		},
	},
}
