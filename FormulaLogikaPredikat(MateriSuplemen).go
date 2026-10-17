package main

func main() {
	var x, y, z, a, f, B, E, M, Akebalik, Ekebalik string
	pertama := (Akebalik && Ekebalik(M(x, y) && Akebalik*z(!M(z, y)) || E(x, z))) || ()

}

//Formula Penguraian top-down

// ∀x∃y (M(x, y) ∧ ∀z (M(z, y) → E(x, z))).
// (M(x, y) ∧ ∀z (M(z, y) → E(x, z)))
// (M(z, y) → E(x, z))
// (x, y)
// (z, y)
// (x, z)
// ∀x∃y
// M
// ∀z
// E
// x
// y
// z

// ∀x∃y∀z (M(x, y) ∧ (M(z, y) → E(x, z))).
// (M(x, y) ∧ (M(z, y) → E(x, z)))
// (M(z, y) → E(x, z))
// (x, y)
// (z, y)
// (x, z)
// ∀x∃y∀z
// M
// E
// x
// y
// z


// ∀x (∃y S(x, f(y)) → B(x, f(a))).
// (∃y S(x, f(y)) → B(x, f(a)))
// (x, f(y))
// (x, f(a))
// (y)
// (a)
// ∀x
// ∃y
// S
// x
// f
// y
// B
// a


//Formula parse tree

∀x
└── ∃y
    └── ∧
        ├── M
        │   ├── x
        │   └── y
        └── ∀z
            └── →
                ├── M
                │   ├── z
                │   └── y
                └── E
                    ├── x
                    └── z


∀x
└── ∃y
    └── ∀z
        └── ∧
            ├── M
            │   ├── x
            │   └── y
            └── →
                ├── M
                │   ├── z
                │   └── y
                └── E
                    ├── x
                    └── z


∀x
└── →
    ├── ∃y
    │   └── S
    │       ├── x
    │       └── f
    │           └── y
    └── B
        ├── x
        └── f
            └── a	