module github.com/samber/lo/pkg/la

// To use range-over-function clause we need to bump the go version to at least 1.23
// but this is required only for this subpackage so instead of bumping minimal version
// of entire project we create sub-module.
go 1.23

require (
	github.com/samber/lo v1.47.0
	github.com/stretchr/testify v1.9.0
	github.com/thoas/go-funk v0.9.3
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
