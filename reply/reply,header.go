package reply

import "net/http"

// SetCookies adds one or more HTTP cookies to the response headers.
// Each cookie is converted to its string representation and added to the Set-Cookie header.
// Empty cookie strings are skipped.
//
// Example:
//
// 	rp.SetCookies(
// 	  http.Cookie{Name: "access_token", Value: accessToken},
// 	  http.Cookie{Name: "refresh_token", Value: refreshToken},
// 	)
func (r *Reply) SetCookies(cookies ...http.Cookie) *Reply {
	for _, c := range cookies {
		if v := c.String(); v != "" {
			r.AddHeader("Set-Cookie", v)
		}
	}
	return r
}

// SetHeader sets a single response header, replacing any existing values.
// Use this when you want to ensure only one value exists for a header.
//
// Example:
//
//	rp.SetHeader("Content-Type", "application/json")
func (r *Reply) SetHeader(key, value string) *Reply {
	r.a.Header().Set(key, value)
	return r
}

// SetHeaders sets multiple response headers at once, replacing any existing values.
// Each key-value pair in the map will replace existing headers with the same key.
//
// Example:
//
//	rp.SetHeaders(map[string]string{
//	    "X-API-Version": "v2.0",
//	    "Cache-Control": "no-cache",
//	})
func (r *Reply) SetHeaders(headers map[string]string) *Reply {
	for k, v := range headers {
		r.SetHeader(k, v)
	}
	return r
}

// AddHeader appends a value to a response header.
// Unlike SetHeader, this allows multiple values for the same header key.
//
// Example:
//
//	rp.AddHeader("Set-Cookie", "session=abc123")
func (r *Reply) AddHeader(key, value string) *Reply {
	r.a.Header().Add(key, value)
	return r
}

// AddHeaders appends multiple values to response headers.
// Each key-value pair is added without replacing existing values.
//
// Example:
//
//	rp.AddHeaders(map[string]string{
//	    "X-Custom-Header": "value1",
//	    "X-Debug-Info": "enabled",
//	})
func (r *Reply) AddHeaders(headers map[string]string) *Reply {
	for k, v := range headers {
		r.AddHeader(k, v)
	}
	return r
}

// DeleteHeader removes a response header by key.
//
// Example:
//
//	rp.DeleteHeader("X-Powered-By")
func (r *Reply) DeleteHeader(key string) *Reply {
	r.a.Header().Del(key)
	return r
}

// DeleteHeaders removes multiple response headers by their keys.
//
// Example:
//
//	rp.DeleteHeaders([]string{"X-Powered-By", "X-API-Version"})
func (r *Reply) DeleteHeaders(keys []string) *Reply {
	for _, k := range keys {
		r.DeleteHeader(k)
	}
	return r
}

// GetHeader retrieves the value of a response header by key.
// Returns an empty string if the header doesn't exist.
//
// Example:
//
//	contentType := rp.GetHeader("Content-Type")
func (r *Reply) GetHeader(key string) string {
	return r.a.Header().Get(key)
}

// GetHeaders retrieves multiple response header values by their keys.
// Returns a slice of values in the same order as the input keys.
// Empty strings are returned for headers that don't exist.
//
// Example:
//
//	values := rp.GetHeaders([]string{"Content-Type", "X-API-Version"})
func (r *Reply) GetHeaders(keys []string) []string {
	values := []string{}
	for _, k := range keys {
		v := r.a.Header().Get(k)
		values = append(values, v)
	}
	return values
}
