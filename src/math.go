package mymath

import (
	"fmt"
	"math/big"
	"math/cmplx"
	"math/rand"
	"time"
)

func MathTesting() {
	bigInt()
	primeCheck()
	bigFloat()
	bigRat()
	complexNums()
	random()
}

func bigInt() {
	//n := 2340300000000000000000000000000000
	n := new(big.Int)
	n.SetString("2340300000000000000000000000000000", 10)
	fmt.Printf("n = %v - %T\n", n, n)

	m := big.NewInt(129)
	fmt.Printf("m = %v - %T\n", n, n)

	n.Add(n, m)
	fmt.Printf("n = %v - %T\n", n, n)

	o := new(big.Int).Mul(n, m)
	fmt.Printf("o = %v - %T\n", o, o)

	fmt.Printf("n.Cmp(o): %d\n", n.Cmp(o))
	fmt.Printf("n.Cmp(m): %d\n", n.Cmp(m))
	fmt.Printf("n.Cmp(n): %d\n", n.Cmp(n))

}

func primeCheck() {
	primes := []*big.Int{
		big.NewInt(3),
		big.NewInt(212),
		big.NewInt(435),
		big.NewInt(350),
		big.NewInt(321),
	}

	for _, p := range primes {
		fmt.Printf("%v a prime? %t\n", p, p.ProbablyPrime(1))
	}
}

func bigFloat() {
	var pi big.Float
	pi.Parse("3.14159265358979323419284028340912830482378491278349123048120348120347", 10)
	fmt.Printf("pi = %.10g\n", &pi)
}

func bigRat() {
	n := big.NewRat(1,2)
	m := big.NewRat(100,2000)
	fmt.Printf("n = %v/%v\n", n.Num(), n.Denom())
	fmt.Printf("n = %v/%v\n", m.Num(), m.Denom())
	fmt.Printf("n.Cmp(m) = %v\n", n.Cmp(m))

	o := new(big.Rat)
	_, _ = fmt.Sscan(".8", o)
	fmt.Printf("o = %v\n", o.Num(), o.Denom())
	fmt.Printf("n.Cmp(o) = %v\n", n.Cmp(o))
}

func complexNums() {
	n := complex(3,5)
	m := 3 +6i

	sum := n+m
	abs := cmplx.Abs(sum)
	fmt.Printf("sum: %v, %T, Abs: %v, %T\n", sum, sum, abs, abs)

	prod := n+m
	icos := cmplx.Acos(prod)
	fmt.Printf("prod: %v, %T, Acos: %v, %T\n", prod, prod, icos, icos)

	pow := cmplx.Pow(n,m)
	fmt.Printf("pow: %v, %T\n", pow, pow)
}

func random() {
	sentences := []string{
		"sdf sdaf as Sasdsdf s f.",
		"Onerfne uef ksdf kuy nwefohsdf.",
		"ASDjfsdpen ifok jasdf nasdf  lasdf.",
		"Dsd lkjsdf sdk.",
		"FSD osdf sdfpj afsd sdflf klj.",
	}

	var n int

	n = rand.Intn(len(sentences))
	fmt.Println("Random Sentences:", sentences[n], n)

	rand.Seed(12332)
	n = rand.Intn(len(sentences))
	fmt.Println("Random Sentences:", sentences[n], n)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n = r.Intn(len(sentences))
	fmt.Println("Random Sentences:", sentences[n], n)
}

