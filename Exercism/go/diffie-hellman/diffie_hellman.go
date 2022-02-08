package diffiehellman

import (
    "math/big"
    "crypto/rand"
)

func PrivateKey(p *big.Int) *big.Int {
	n,_ := rand.Int(rand.Reader, p)
	x := big.NewInt(0).Mod(n, big.NewInt(0).Add(p, big.NewInt(-1)))
	if x.Cmp(big.NewInt(2)) == -1 {
		return big.NewInt(0).Add(x, big.NewInt(2))
	}
    return x
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	new_p := PrivateKey(p)
	return new_p, PublicKey(new_p, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}