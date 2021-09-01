package main

import (
	"net/http"
	"strings"
)

const (
	headerClientIP           = "X-Client-IP"
	headerContentLength      = "Content-Length"
	headerContentType        = "Content-Type"
	headerContentTypeOptions = "X-Content-Type-Options"
	headerForwardedFor       = "X-Forwarded-For"
	headerFrameOptions       = "X-Frame-Options"
	headerOrigin             = "Origin"
	headerOriginHTTP         = "HTTP_ORIGIN"
	headerOriginXHTTP        = "HTTP_X_ORIGIN"
	headerPathInfo           = "PATH_INFO"
	headerRealIP             = "X-Real-IP"
	headerRemoteReferer      = "Referer"
	headerStrictTransport    = "Strict-Transport-Security"
	headerXSSProtection      = "X-XSS-Protection"
)

// IPSpoofing type used to detect IP spoofing attacks.
type IPSpoofing struct{}

// NewIPSpoofing creates a new instance of IPSpoofing
func NewIPSpoofing() IPSpoofing {
	return IPSpoofing{}
}

// Handler checks if request of safe against IPSpoofing attacks
func (m IPSpoofing) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			clientIP     = r.Header.Get(headerClientIP)
			forwardedFor = r.Header.Get(headerForwardedFor)
			realIP       = r.Header.Get(headerRealIP)
			isSafe       = false
		)
		if clientIP == "" && realIP == "" {
			next.ServeHTTP(w, r)
			return
		}
		for _, ip := range strings.Split(forwardedFor, ", ") {
			if ip == clientIP || ip == realIP {
				isSafe = true
				break
			}
		}
		if !isSafe {
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set(headerContentType, "text/plain")
				w.Header().Set(headerContentLength, "0")
				w.WriteHeader(http.StatusForbidden)
			}).ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	}
}
