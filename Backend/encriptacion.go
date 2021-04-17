package main

import (
	"fmt"
	"github.com/fernet/fernet-go"
)

func encriptar(texto string) (string,string){
	llave := fernet.Key{}
	llave.Generate()
	fmt.Println("Llave de sha256", llave.Encode())
	b := []byte(texto)
	token,_ := fernet.EncryptAndSign(b, &llave)
	tokenTexto := string(token)
	return tokenTexto,llave.Encode()
}

func encriptarConLlave(texto string, key string) string {
	llave,_ := fernet.DecodeKey(key)
	b := []byte(texto)
	token,_ := fernet.EncryptAndSign(b, llave)
	tokenTexto := string(token)
	return tokenTexto
}

func desencriptar(token string, key string)  {
	llave := fernet.MustDecodeKeys(key)
	b := []byte(token)
	textoB := fernet.VerifyAndDecrypt(b, 1000000000, llave)
	texto := string(textoB)
	fmt.Println(texto)
}
