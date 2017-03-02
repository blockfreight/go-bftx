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
	"strconv"

	"github.com/blockfreight/blockfreight-alpha/blockfreight/bft/bf_tx"
	"github.com/davecgh/go-spew/spew"
)

func Sign_BF_TX(bft_tx bf_tx.BF_TX) bf_tx.BF_TX{

    content := bf_tx.BF_TXContent(bft_tx)
    
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
	
	r, s, serr := ecdsa.Sign(rand.Reader, privatekey, signhash)
	if serr != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	signature := r.Bytes()
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

	//Set Private Key and Sign to BF_TX
	bft_tx.PrivateKey = *privatekey
	bft_tx.Signhash = signhash
	bft_tx.Signature = sign
	bft_tx.Signed = verifystatus
	printJson := false
	if printJson { spew.Dump(bft_tx) }
	return bft_tx
}