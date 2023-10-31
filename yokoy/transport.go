package yokoy

import "net/http"

type Transport struct {
	T     http.RoundTripper
	token string
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Set("X-Yk-Auth-Method", "yokoy") // required by yokoy api

	return t.T.RoundTrip(req)
}

func NewTransport(token string) *Transport {
	return &Transport{
		T:     http.DefaultTransport,
		token: token,
	}
}
