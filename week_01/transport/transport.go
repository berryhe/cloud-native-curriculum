package transport

import (
	"fmt"
	"net"
	"net/http"
	"net/textproto"
	"strings"
)

type Response struct {
	http.ResponseWriter
	req    *http.Request
	Status int
}

var Version string

func NewResp(w http.ResponseWriter, r *http.Request) *Response {
	return &Response{
		ResponseWriter: w,
		req:            r,
	}
}
func (resp *Response) handlePrintLog() {
	path := resp.req.URL.Path
	raw := resp.req.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}

	fmt.Printf("PATH: %s | Method: %s | Status: %d | clientIP: %s\n", path, resp.req.Method, resp.Status, resp.ClientIP())
}

func (resp *Response) setHandle(handle http.Header) {

	for k, v := range handle {

		key := textproto.CanonicalMIMEHeaderKey(k)
		if _, ok := resp.Header()[key]; ok {
			resp.Header()[key] = append(resp.Header()[key], v...)
		} else {
			resp.Header()[key] = v
		}

	}
}

func (resp *Response) WriteData(data []byte) {
	defer resp.handlePrintLog()

	if resp.Status == 0 {
		resp.Status = http.StatusOK
	}

	resp.WriteHeader(resp.Status)
	resp.Write(data)
}

func (resp *Response) ClientIP() string {
	if addr := resp.req.Header.Get("X-Forwarded-For"); addr != "" {
		return addr
	}

	if addr := resp.req.Header.Get("X-Real-IP"); addr != "" {
		return addr
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(resp.req.RemoteAddr))
	if err != nil {
		return ""
	}

	remoteIP := net.ParseIP(ip)
	if remoteIP == nil {
		return ""
	}

	return remoteIP.String()
}
func HandleRootPath(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("VERSION", Version)

	resp := NewResp(w, r)
	resp.setHandle(r.Header)
	resp.WriteData([]byte("ok"))
}

func HandleHealthz(w http.ResponseWriter, r *http.Request) {
	resp := NewResp(w, r)
	resp.Status = http.StatusOK
	resp.WriteData([]byte("ok"))
}
