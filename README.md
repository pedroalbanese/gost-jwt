# gost-jwt
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/gost-jwt/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/gost-jwt?status.png)](http://godoc.org/github.com/pedroalbanese/gost-jwt)
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/gost-jwt/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/gost-jwt/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/gost-jwt)](https://goreportcard.com/report/github.com/pedroalbanese/gost-jwt)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/gost-jwt)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/gost-jwt)](https://github.com/pedroalbanese/gost-jwt/releases)

JSON Web Tokens are an open, industry standard method (RFC 7519) for representing claims securely between two parties. This version only works with HMAC-Streebog (RFC 6986) signatures. ECDH_ES variant, consists in the direct use of a shared symmetric key as the Content Encryption Key (CEK) for the block encryption step (recommended). 

## Usage of gost-jwt:
<pre>  One of the following flags is required: sign, verify
  -alg string
        signing algorithm identifier (default "HG256")
  -claim value
        add additional claims. may be used more than once (default {})
  -compact
        output compact JSON
  -debug
        print out all kinds of debug data
  -header value
        add additional header params. may be used more than once (default {})
  -key string
        path to key file or '-' to read from stdin
  -show
        Show header
  -sign string
        path to claims object to sign, '-' to read from stdin, or '+' to use only -claim args
  -verify string
        path to JWT token to verify or '-' to read from stdin
</pre>
        
## License

This project is licensed under the ISC License.

#### Copyright (c) 2020-2022 Pedro F. Albanese - ALBANESE Research Lab
