package ipgeolocation

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"os"
	"strings"
	"syscall"
)

func (c Config) Format(out fmt.State, _ rune) {
	fmt.Fprint(out, redactKey(fmt.Sprint(c.SafeMap()), strings.TrimSpace(c.APIKey)))
}

type clientSettings struct{ config normalizedConfig }

func (settings clientSettings) Format(out fmt.State, verb rune) {
	Config(settings.config).Format(out, verb)
}

func redactKey(text, key string) string {
	if key == "" || text == "" {
		return text
	}
	const mask = "[REDACTED]"
	text = strings.ReplaceAll(text, key, mask)
	if strings.ContainsAny(text, "%+") {
		for _, encoded := range []string{url.QueryEscape(key), url.PathEscape(key)} {
			if encoded != key {
				text = strings.ReplaceAll(text, encoded, mask)
			}
		}
		decoded := unescapeDiagnostic(text, false)
		if strings.Contains(decoded, key) || strings.Contains(key, " ") && strings.Contains(unescapeDiagnostic(text, true), key) {
			return mask
		}
	}
	return text
}

func unescapeDiagnostic(text string, plusAsSpace bool) string {
	var out strings.Builder
	for i := 0; i < len(text); i++ {
		if text[i] == '%' && i+2 < len(text) {
			if decoded, err := url.PathUnescape(text[i : i+3]); err == nil {
				out.WriteString(decoded)
				i += 2
				continue
			}
		}
		if plusAsSpace && text[i] == '+' {
			out.WriteByte(' ')
			continue
		}
		out.WriteByte(text[i])
	}
	return out.String()
}

func safeTransportError(err error, key string) error {
	if key == "" {
		return err
	}
	switch e := err.(type) {
	case *RequestTimeoutError:
		return &RequestTimeoutError{TransportError{Message: redactKey(e.Message, key), Err: safeCause(e.Err, key)}}
	case *TransportError:
		return &TransportError{Message: redactKey(e.Message, key), Err: safeCause(e.Err, key)}
	default:
		return safeCause(err, key)
	}
}

func safeCause(err error, key string) error {
	if err == nil {
		return nil
	}
	if safe, ok := safeTLSVerificationError(err, key); ok {
		return safe
	}
	switch e := err.(type) {
	case *url.Error:
		address := ""
		if parsed, parseErr := url.Parse(e.URL); parseErr == nil {
			parsed.RawQuery, parsed.Fragment, parsed.ForceQuery, parsed.User = "", "", false, nil
			address = redactKey(parsed.String(), key)
		}
		return &url.Error{Op: redactKey(e.Op, key), URL: address, Err: safeCause(e.Err, key)}
	case x509.UnknownAuthorityError:
		return x509.UnknownAuthorityError{Cert: safeCertificate(e.Cert, key)}
	case x509.HostnameError:
		return x509.HostnameError{Certificate: safeCertificate(e.Certificate, key), Host: redactKey(e.Host, key)}
	case x509.CertificateInvalidError:
		return x509.CertificateInvalidError{Cert: safeCertificate(e.Cert, key), Reason: e.Reason, Detail: redactKey(e.Detail, key)}
	case x509.SystemRootsError:
		return x509.SystemRootsError{Err: safeCause(e.Err, key)}
	case *net.OpError:
		copy := *e
		copy.Op, copy.Net = redactKey(e.Op, key), redactKey(e.Net, key)
		copy.Err = safeCause(e.Err, key)
		copy.Source, copy.Addr = safeAddr(e.Source, key), safeAddr(e.Addr, key)
		return &copy
	case *net.DNSError:
		return &net.DNSError{Err: redactKey(e.Err, key), Name: redactKey(e.Name, key), Server: redactKey(e.Server, key), IsTimeout: e.IsTimeout, IsTemporary: e.IsTemporary, IsNotFound: e.IsNotFound}
	case *os.SyscallError:
		return &os.SyscallError{Syscall: redactKey(e.Syscall, key), Err: safeCause(e.Err, key)}
	case syscall.Errno:
		return e
	}
	for _, sentinel := range []error{context.Canceled, context.DeadlineExceeded, io.EOF, io.ErrUnexpectedEOF, net.ErrClosed} {
		if errors.Is(err, sentinel) {
			return sentinel
		}
	}
	var timeout net.Error
	if errors.As(err, &timeout) && timeout.Timeout() {
		return context.DeadlineExceeded
	}
	return errors.New(redactKey(err.Error(), key))
}

func safeCertificate(cert *x509.Certificate, key string) *x509.Certificate {
	if cert != nil && len(cert.Raw) > 0 {
		der := string(cert.Raw)
		if redactKey(der, key) == der {
			if parsed, err := x509.ParseCertificate(cert.Raw); err == nil {
				return parsed
			}
		}
	}
	return &x509.Certificate{}
}

type errorAddress struct{ network, address string }

func (a errorAddress) Network() string { return a.network }
func (a errorAddress) String() string  { return a.address }

func safeAddr(addr net.Addr, key string) net.Addr {
	if addr == nil {
		return nil
	}
	address, network := addr.String(), addr.Network()
	clean := errorAddress{redactKey(network, key), redactKey(address, key)}
	switch addr.(type) {
	case *net.IPAddr, *net.IPNet, *net.TCPAddr, *net.UDPAddr, *net.UnixAddr:
		if clean.address == address && clean.network == network {
			return addr
		}
	}
	return clean
}
