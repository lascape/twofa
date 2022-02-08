# towfa

Generate a generic library of 2FA tokens compatible with Google Authenticator

```shell
go get -u github.com/golandscape/twofa
$twofa "you secret"
result:{Secret:WRQLGCZHKY6HMEL4 Expire:27 Code:409056}
```

Relevant RFCs:
* http://tools.ietf.org/html/rfc4226
* http://tools.ietf.org/html/rfc6238