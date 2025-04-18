package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

const (
	upper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower    = "abcdefghijklmnopqrstuvwxyz"
	digits   = "0123456789"
	specials = "!@#$%^&*()-_=+[]{}<>?/|"
)

func main() {
	length := flag.Int("length", 12, "Longitud de la contraseña")
	includeSymbols := flag.Bool("symbols", false, "Incluir símbolos especiales")
	number := flag.Int("number", 1, "Cantidad de contraseñas a generar")
	flag.Parse()

	charset := upper + lower + digits
	if *includeSymbols {
		charset += specials
	}

	for i := 0; i < *number; i++ {
		pass := GeneratePassword(*length, charset)
		fmt.Println(pass)
	}
}

func GeneratePassword(length int, charset string) string {
	password := ""
	for i := 0; i < length; i++ {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password += string(charset[randIndex.Int64()])
	}
	return password
}
