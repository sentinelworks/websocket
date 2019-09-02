package restclient

import (
	"fmt"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"crypto/tls"
        "crypto/x509"
	"net/http/httputil"
)

type Params map[string]string

var UnexpectedStatus = errors.New("Server returned unexpected status.")

type RequestResponse struct {
	Url      string        // Raw URL string
	Method   string        // HTTP method to use
	Xauth    *string        // HTTP method to use
	RawInput *string         // Raw text of server response (JSON or otherwise)
	//RawInput *[]byte         // Raw text of server response (JSON or otherwise)
	Userinfo *url.Userinfo // Optional username/password to authenticate this request
	Params   Params        // URL parameters for GET requests (ignored otherwise)
	Header   *http.Header  // HTTP Headers to use (will override defaults)

	ExpectedStatus int

	Data   interface{} // Data to JSON-encode and POST
	Result interface{} // Successful response is unmarshalled into Result
	Error  interface{} // Error response is unmarshalled into Error

	Timestamp    time.Time      // Time when HTTP request was sent
	RawText      string         // Raw text of server response (JSON or otherwise)
	Status       int            // HTTP status for executed request
	HttpResponse *http.Response // Response object from http package
}

type Client struct {
	HttpClient      *http.Client
	UnsafeBasicAuth bool // Allow Basic Auth over unencrypted HTTP
	Log             bool // Log request and response
}

func New() *Client {
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    myclient := &http.Client{Transport: tr}
		//myclient := new(http.Client)


	return &Client{
		HttpClient:      myclient,
		UnsafeBasicAuth: false,
		Log: true,
//		Log: false,
	}
}

func (c *Client) SetTransport(transport http.RoundTripper) *Client {
        if transport != nil {
                c.HttpClient.Transport = transport
        }
        return c
}

func (c *Client) Do(rr *RequestResponse) (status int, err error) {
	rr.Method = strings.ToUpper(rr.Method)
	u, err := url.Parse(rr.Url)
	if err != nil {
		log.Println(err)
		return
	}

	if rr.Method == "GET" && rr.Params != nil {
		vals := u.Query()
		for k, v := range rr.Params {
			vals.Set(k, v)
		}
		u.RawQuery = vals.Encode()
	}

	rr.Timestamp = time.Now()
	m := string(rr.Method)
	var req *http.Request

	if rr.Data == nil {
		if rr.RawInput != nil {
			if c.Log {
				log.Println(*rr.RawInput)
			}
			b:= []byte(*rr.RawInput)
			buf := bytes.NewBuffer(b)
			req, err = http.NewRequest(m, u.String(), buf)
			if err != nil {
				log.Println(err)
				return
			}
		} else {
			if c.Log {
				log.Println("**** no input")
			}
			req, _ = http.NewRequest(m, u.String(), nil)
		}
	} else {
		var b []byte
		b, err = json.Marshal(&rr.Data)
		if err != nil {
			log.Println(err)
			return
		}
		buf := bytes.NewBuffer(b)
		req, err = http.NewRequest(m, u.String(), buf)
		if err != nil {
			log.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/json")
	}

	if rr.Xauth != nil {
		if c.Log {
			log.Println("***** start xauth header ", *rr.Xauth)
		}
		req.Header.Set("X-Auth-Token", *rr.Xauth)
		//req.Header.Add("X-Auth-Token", rr.Xauth)
	}

	if rr.Header != nil {
		for key, values := range *rr.Header {
			if len(values) > 0 {
				req.Header.Set(key, values[0]) // Possible to overwrite Content-Type
			}
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Add("Accept", "application/json")
	}

	if rr.Userinfo != nil {
		if !c.UnsafeBasicAuth && u.Scheme != "https" {
			err = errors.New("Unsafe to use HTTP Basic authentication without HTTPS")
			return
		}
		pwd, _ := rr.Userinfo.Password()
		req.SetBasicAuth(rr.Userinfo.Username(), pwd)
	}

	if c.Log {
		dump, err11 := httputil.DumpRequest(req, true)
		if err11 != nil {
			log.Print("Dump error: ")
		}

		fmt.Printf("Request data: %s\n\n", dump)

		log.Println("--------------------------------------------------------------------------------")
		log.Println("REQUEST")
		log.Println("--------------------------------------------------------------------------------")
		prettyPrint(req)
		log.Print("Payload: ")
		prettyPrint(rr.Data)
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	status = resp.StatusCode
	rr.HttpResponse = resp
	rr.Status = resp.StatusCode
	var data []byte
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	rr.RawText = strings.TrimSpace(string(data))
	json.Unmarshal(data, &rr.Error) // Ignore errors
	if rr.RawText != "" && status < 300 {
		err = json.Unmarshal(data, &rr.Result) // Ignore errors
	}
	if c.Log {
		log.Println("--------------------------------------------------------------------------------")
		log.Println("RESPONSE")
		log.Println("--------------------------------------------------------------------------------")
		log.Println("Status: ", status)
		if rr.RawText != "" {
			raw := json.RawMessage{}
			if json.Unmarshal(data, &raw) == nil {
				prettyPrint(&raw)
			} else {
				prettyPrint(rr.RawText)
			}
		} else {
			log.Println("Empty response body")
		}

	}
	if rr.ExpectedStatus != 0 && status != rr.ExpectedStatus {
		log.Printf("Expected status %s but got %s", rr.ExpectedStatus, status)
		return status, UnexpectedStatus
	}
	return
}

var (
	defaultClient = New()
)

func Do(rr *RequestResponse) (status int, err error) {
	return defaultClient.Do(rr)
}

func (c *Client) getTLSConfig() (*tls.Config, error) {
        transport, err := c.getTransport()
        if err != nil {
                return nil, err
        }
        if transport.TLSClientConfig == nil {
                transport.TLSClientConfig = &tls.Config{}
        }
        return transport.TLSClientConfig, nil
}


func (c *Client) getTransport() (*http.Transport, error) {
        if c.HttpClient.Transport == nil {
                c.SetTransport(new(http.Transport))
        }

        if transport, ok := c.HttpClient.Transport.(*http.Transport); ok {
                return transport, nil
        }
        return nil, errors.New("current transport is not an *http.Transport instance")
}



func (c *Client) SetTLSClientConfig(config *tls.Config) *Client {
        transport, err := c.getTransport()
        if err != nil {
                log.Println("ERROR %v", err)
                return c
        }
        transport.TLSClientConfig = config
        return c
}

func (c *Client) SetCertificates(certs ...tls.Certificate) *Client {
        config, err := c.getTLSConfig()
        if err != nil {
                log.Println("ERROR %v", err)
                return c
        }
        config.Certificates = append(config.Certificates, certs...)
        return c
}

func (c *Client) SetRootCertificate(pemFilePath string) *Client {
        rootPemData, err := ioutil.ReadFile(pemFilePath)
        if err != nil {
                log.Println("ERROR %v", err)
                return c
        }

        config, err := c.getTLSConfig()
        if err != nil {
                log.Println("ERROR %v", err)
                return c
        }
        if config.RootCAs == nil {
                config.RootCAs = x509.NewCertPool()
        }

        config.RootCAs.AppendCertsFromPEM(rootPemData)

        return c
}

