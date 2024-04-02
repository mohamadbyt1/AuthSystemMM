package auth

import (
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"reflect"

)
func (ec *EllipticCurve) GenrateKey() (PrivKey *ecdsa.PrivateKey,PubKey *ecdsa.PublicKey,err error){

	privKey,err := ecdsa.GenerateKey(ec.PubKeyCurve,rand.Reader)
	if err == nil{
		ec.PrivateKey = privKey
		ec.PublickKey = &privKey.PublicKey
	}
	return
}

func (ec *EllipticCurve) EncodePrivate(privKey *ecdsa.PrivateKey)(key string , err error){
	encode, err := x509.MarshalECPrivateKey(privKey)
	if err != nil{
		return
	}
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes:encode})
	key = string(pemEncoded)
	return
}

func (ec *EllipticCurve) EncodePublick(pubKey *ecdsa.PublicKey)(key string , err error){
	encode,err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil{
		return
	}
	pemEncode:= pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes:encode})
	key = string(pemEncode)
	return
}

func (ec *EllipticCurve) DecodePrivate(pemEncodePrivate string)(PrivateKey *ecdsa.PrivateKey,err error){
	blockPriv, _ := pem.Decode([]byte(pemEncodePrivate))
	x509EncodedPriv := blockPriv.Bytes
	PrivateKey,err = x509.ParseECPrivateKey(x509EncodedPriv)
	return
}
func (ec *EllipticCurve)DecodePublick (pemEncoded string)(PublickKey *ecdsa.PublicKey,err error){
	blockPub,_ := pem.Decode([]byte(pemEncoded))
	x509EncodePub := blockPub.Bytes
	genericPublickKey,err := x509.ParsePKIXPublicKey(x509EncodePub)
	PublickKey = genericPublickKey.(*ecdsa.PublicKey)
	return
}
func (ec *EllipticCurve) VerifySignature (pubKey *ecdsa.PublicKey,privKey *ecdsa.PrivateKey)(signature []byte , ok bool, err error) {
	h := md5.New()
	_, err = io.WriteString(h,"this is a message to be signed")
	if err != nil {
		return
	}
	signHash := h.Sum(nil)
	r, s, serr := ecdsa.Sign(rand.Reader,privKey,signHash)
	if serr != nil {
		return []byte(""),false , serr
	}
	signature = r.Bytes()
	signature = append(signature, s.Bytes()...)
	ok = ecdsa.Verify(pubKey, signHash,r,s)
	return
}
func (ec *EllipticCurve) Test(privKey *ecdsa.PrivateKey, pubKey *ecdsa.PublicKey) (err error) {

    encPriv, err := ec.EncodePrivate(privKey)
    if err != nil {
        return
    }
    encPub, err := ec.EncodePublick(pubKey)
    if err != nil {
        return
    }
    priv2, err := ec.DecodePrivate(encPriv)
    if err != nil {
        return
    }
    pub2, err := ec.DecodePublick(encPub)
    if err != nil {
        return
    }

    if !reflect.DeepEqual(privKey, priv2) {
        err = errors.New("private keys do not match")
        return
    }
    if !reflect.DeepEqual(pubKey, pub2) {
        err = errors.New("public keys do not match")
        return
    }

    return
}