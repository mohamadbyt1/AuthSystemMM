package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
)
type EllipticCurve struct{
	PubKeyCurve elliptic.Curve
	PrivateKey *ecdsa.PrivateKey
	PublickKey *ecdsa.PublicKey
}
func NewElliptickCurve(curve elliptic.Curve) *EllipticCurve {
	return &EllipticCurve{
		PubKeyCurve: curve,
		PrivateKey: new(ecdsa.PrivateKey),

	}
}
