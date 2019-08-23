module main

go 1.12

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190215142949-d0b11bdaac8a
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190222072716-a9d3bda3a223
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 => github.com/go-check/check v0.0.0-20161208181325-20d25e280405
	gopkg.in/go-playground/validator.v8 v8.18.2 => github.com/go-playground/validator v8.18.2+incompatible
	gopkg.in/yaml.v2 v2.2.2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8
	scws v1.0.0 => ./scws
)

require github.com/gin-gonic/gin v1.4.0

require github.com/hetao29/goscws v0.0.0-20130817132951-16c49c56e273

require scws v1.0.0
