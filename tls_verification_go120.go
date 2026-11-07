//go:build go1.20

package ipgeolocation

import (
	"crypto/tls"
	"crypto/x509"
)

func safeTLSVerificationError(err error, key string) (error, bool) {
	e, ok := err.(*tls.CertificateVerificationError)
	if !ok {
		return nil, false
	}
	certs := make([]*x509.Certificate, len(e.UnverifiedCertificates))
	for i, cert := range e.UnverifiedCertificates {
		certs[i] = safeCertificate(cert, key)
	}
	return &tls.CertificateVerificationError{UnverifiedCertificates: certs, Err: safeCause(e.Err, key)}, true
}
