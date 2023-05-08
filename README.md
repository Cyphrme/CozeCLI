# `coze` 
`coze` - a CLI client for Coze, a cryptographic JSON messaging specification.

See also the [main Coze specification repository](https://github.com/Cyphrme/Coze)

## Installation
Installation from source requires Go.  Install executable `coze` from local copy
of the repository:

```sh
go install coze.go
```

Or alternatively (on non-Nix systems):

```
go install github.com/cyphrme/coze_cli@latest && mv $GOBIN/coze_cli coze
```
Note that `$GOPATH` and/or `$GOBIN` needs to be in `$PATH`.
Otherwise Go usually installs to `/usr/local/bin`.

The above installs the executable as `coze`, but `go install
github.com/cyphrme/coze_cli` installs as `coze_cli`.  After the `go build -o`
[proposal is implemented](https://github.com/golang/go/issues/44469), in the
future install should be possible with `go install -o coze
github.com/cyphrme/coze_cli@latest`.


# Development and Testing 

## Go mod
The directive `replace` in `go.mod` is useful for local developement of Go coze repo and the cli coze repo.  Usually have this in `go.mod` for
`go.work` to work.  However, don't commit with `replace` as it will break `go install`.
// Go work/go mod/go has a bug: https://github.com/golang/go/issues/54264
// Fix this once they fix go workspace
replace github.com/cyphrme/coze@latest => ../coze


### sign
```sh
go run coze.go sign '{"pay":{"msg":"Hi!","alg":"ES256","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"}}' '{"alg":"ES256","x":"tmdb5tBJlKaCOTgvsZvtlf4XCL8MyasTdqKDYzdRsc6p898M4IuvQAsEthE624-jNyWzR4BLM29eupkxu80zGw","d":"kqRoh9VcH_SW3kQrSxwt4gP_o66PYK96xMuLk6Hri08","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"}'
```

Expected output is a signed coze, e.g. `{"pay":{"msg":"Hi!","alg":"ES256","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"},"sig":"pmWSJdfxsG-26rLBHxRp6qLARZROxDFCnGlPpDtAFqZZ6sIs3-x8BZ4FTf478DG0kdPM8QVotfEl2lBzdVRS2g"}`.


### signpay
```sh
go run coze.go signpay '{"msg":"Hi!","alg":"ES256","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"}' '{"alg":"ES256","x":"tmdb5tBJlKaCOTgvsZvtlf4XCL8MyasTdqKDYzdRsc6p898M4IuvQAsEthE624-jNyWzR4BLM29eupkxu80zGw","d":"kqRoh9VcH_SW3kQrSxwt4gP_o66PYK96xMuLk6Hri08","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"}'
```

Expected output is a signed coze, e.g. `{"pay":{"msg":"Hi!","alg":"ES256","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"},"sig":"pmWSJdfxsG-26rLBHxRp6qLARZROxDFCnGlPpDtAFqZZ6sIs3-x8BZ4FTf478DG0kdPM8QVotfEl2lBzdVRS2g"}`.


### verify
```sh
go run coze.go verify '{"pay":{"msg":"Hi!","alg":"ES256","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"},"sig":"pmWSJdfxsG-26rLBHxRp6qLARZROxDFCnGlPpDtAFqZZ6sIs3-x8BZ4FTf478DG0kdPM8QVotfEl2lBzdVRS2g"}' '{"alg":"ES256","x":"tmdb5tBJlKaCOTgvsZvtlf4XCL8MyasTdqKDYzdRsc6p898M4IuvQAsEthE624-jNyWzR4BLM29eupkxu80zGw","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"}'
```

Expected output is the string `true` or `false`


### newkey
Parameter specifying key alg is optional (ES256 currently default)

```sh
go run coze.go newkey
```

or

```sh
go run coze.go newkey ES384
```

Expected output is a new private Coze key, e.g.

`{"alg":"ES256","x":"tmdb5tBJlKaCOTgvsZvtlf4XCL8MyasTdqKDYzdRsc6p898M4IuvQAsEthE624-jNyWzR4BLM29eupkxu80zGw","d":"kqRoh9VcH_SW3kQrSxwt4gP_o66PYK96xMuLk6Hri08","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"}`



newkey will print the key.  If not wanting to print the key, pipe the output to file.

```
go run coze.go newkey ES384 > zami.czk
```



### tmb
```sh
go run coze.go tmb '{"alg":"ES256","x":"tmdb5tBJlKaCOTgvsZvtlf4XCL8MyasTdqKDYzdRsc6p898M4IuvQAsEthE624-jNyWzR4BLM29eupkxu80zGw"}'
```

Expected output is the b64 thumbprint: `QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw`


### meta
```sh
go run coze.go meta '{"pay":{"msg":"Hi!","alg":"ES256","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"},"sig":"pmWSJdfxsG-26rLBHxRp6qLARZROxDFCnGlPpDtAFqZZ6sIs3-x8BZ4FTf478DG0kdPM8QVotfEl2lBzdVRS2g"}'
```

Expected output is an object with `can`, `cad`, and `czd` (e.g. `{"can":["msg","alg","tmb"],"cad":"-_uFn_bu_N8AMeLFU1dZ6ImQRS11J-9HAt7gyW6PqFU","czd":"KqQzYSDp956CFEd3XR8SOB0ue8I7TOlpKNsfEqv-cPU"}`)


### revoke
```sh
go run coze.go revoke '{"alg":"ES256","x":"tmdb5tBJlKaCOTgvsZvtlf4XCL8MyasTdqKDYzdRsc6p898M4IuvQAsEthE624-jNyWzR4BLM29eupkxu80zGw","d":"kqRoh9VcH_SW3kQrSxwt4gP_o66PYK96xMuLk6Hri08","tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw"}'
```

Expected output is a revoke coze (e.g.
`{"pay":{"alg":"ES256","iat":1683221243,"tmb":"QV_dSgjtGP7kiZLxqhPsp5P9Gufgv7rwzFGuUCjm_Zw","typ":"cyphr.me/key/revoke","rvk":1683221243},"sig":"rTVGIQcTExRFW8CSa-eKrDfEV4FlnmowI-BKBYvf_AhgCzSH0XFHDuQzLEoojIe3bigUkdDGfgTgxOrHRFsMxg"}`).


# TODO
// Coze CLI todo
- Specify key files with flags and/or pipe key in. 

Flags Ideas:
-alg for "include alg"
-tmb for "include thumbprint"
-iat for "include iat"
-typ <type> for "include typ with this value".

Or perhaps the opposite.  --noalg --notmb --noiat --notyp or --no for not including any no explicitly specified fields.  