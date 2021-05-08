package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type BlockChain struct {
	Bloques []Bloque
}

type Bloque struct {
	Indice int
	Fecha string
	Data string
	Nonce int
	ProviousHash string
	Hash string
}

func NewBloque(indice, nonce int, fecha, data, previous, hash string) *Bloque {
	return &Bloque{
		Indice:       indice,
		Fecha:        fecha,
		Data:         data,
		Nonce:        nonce,
		ProviousHash: previous,
		Hash:         hash,
	}
}

func(b *Bloque) crearHash(){
	var texto string
	var hash string

	//fmt.Println(primeros)
	texto = strconv.Itoa(b.Indice) + b.Fecha + b.ProviousHash + b.Data + strconv.Itoa(b.Nonce)
	hash = encriptar256(texto)
	b.Nonce++
	fmt.Printf(hash)

	b.Hash = hash
}


func(b *Bloque) guardarBloque() {
	f, err := os.Create("Bloques\\"+strconv.Itoa(b.Indice)+".txt")

	text := strconv.Itoa(b.Indice) + "\n" + b.Fecha + "\n" + b.Data + "\n" +strconv.Itoa(b.Nonce) + "\n"+  b.ProviousHash + "\n" + b.Hash
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(text)

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("$$$ Bloque guardado.")
}