module github.com/samber/lo

go 1.25

//
// Dev dependencies are excluded from releases. Please check CI.
//

require (
	github.com/stretchr/testify v1.12.1
	github.com/thoas/go-funk v0.9.3
	go.uber.org/goleak v1.2.1
	golang.org/x/text v0.22.0
)

require go.yaml.in/yaml/v3 v3.0.5 // indirect
