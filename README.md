[![GoDoc](https://godoc.org/github.com/danielstutzman/go-belugacdn?status.svg)](https://godoc.org/github.com/danielstutzman/go-belugacdn)

# Go SDK for BelugaCDN API

## Supported
* Username and password authentication
* Site Configuration API:
  * Create Site [(API docs)](https://docs.belugacdn.com/docs/create-new-site)
  * List Sites [(API docs)](https://docs.belugacdn.com/docs/list-sites)
  * Update Site [(API docs)](https://docs.belugacdn.com/docs/update-site)
  * Delete Site [(API docs)](https://docs.belugacdn.com/v2/docs/delete-cdn-site)
* SSL Certificate API:
  * Create SSL Certificate [(API docs)](https://docs.belugacdn.com/v2/docs/create-ssl-certificate)
  * List SSL Certificates [(API docs)](https://docs.belugacdn.com/v2/docs/list-ssl-certificates)
  * Delete SSL Certificate [(API docs)](https://docs.belugacdn.com/v2/docs/delete-ssl-certificate)

## Not supported yet
* API token authentication
* Update SSL Certificate
* Cache Invalidation

## Installation

```
go get -u github.com/danielstutzman/go-belugacdn
```

## See also

* [beluga_py](https://github.com/jaddison/beluga_py)
* [python_beluga](https://github.com/belugacdn/python-beluga)
