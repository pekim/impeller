module github.com/pekim/impeller

go 1.26.5

replace github.com/pekim/gl-purego => ../gl-purego

replace github.com/go-clang/clang-v15 => github.com/pekim/clang-v15 v0.0.0-20240830114552-c0d27ccce9ec

require (
	github.com/dave/jennifer v1.7.1
	github.com/go-clang/clang-v15 v0.0.0-20230222085438-ee3102fa0c71
	github.com/go-webgpu/goffi v0.6.3
	github.com/pekim/gl-purego v0.0.0
	github.com/stretchr/testify v1.11.1
	golang.org/x/text v0.41.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
