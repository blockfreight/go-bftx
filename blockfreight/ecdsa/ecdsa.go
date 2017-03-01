package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"math/big"
	"os"
	//"reflect"
	"strconv"

	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bf_tx"
	"github.com/davecgh/go-spew/spew"
)

func Sign_BFTX(bftx bf_tx.BF_TX) bf_tx.BF_TX{

    content := bf_tx.BFTXContent(bftx)
    
	pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256
	
	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	//var pubkey ecdsa.PublicKey
	pubkey := privatekey.PublicKey
	
	//fmt.Println("Private Key :")
	//fmt.Printf("%x \n\n", privatekey)
	
	//fmt.Println("Public Key :")
	//fmt.Printf("%x \n\n", pubkey)
	
	// Sign ecdsa style
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)
	
	io.WriteString(h, content)
	signhash := h.Sum(nil)
	//fmt.Println("signhash:",signhash, reflect.TypeOf(signhash))
	
	r, s, serr := ecdsa.Sign(rand.Reader, privatekey, signhash)
	if serr != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	signature := r.Bytes()
	//fmt.Println("signature 1:",signature)
	signature = append(signature, s.Bytes()...)
	//fmt.Printf("%x\n\n", signature)
	//fmt.Println(signature)
	sign := ""
	for i, _ := range signature {
		sign += strconv.Itoa(int(signature[i]))
	}
	//fmt.Println(sign)
	
	// Verify
	verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
	//fmt.Println(verifystatus) // should be true

	//Set Private Key and Sign to BFTX
	bftx.PrivateKey = *privatekey
	bftx.Signhash = signhash
	bftx.Signature = sign
	bftx.Signed = verifystatus
	printJson := false
	if printJson { spew.Dump(bftx) }
	return bftx
}