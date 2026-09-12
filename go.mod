module github.com/pekim/impeller

go 1.26.5

// Address deprecation of clang_getDiagnosticCategoryName.
// Avoids a noisy message when generating api.
replace github.com/go-clang/clang-v15 => github.com/pekim/clang-v15 v0.0.0-20240830114552-c0d27ccce9ec

require (
	github.com/dave/jennifer v1.7.1
	github.com/go-clang/clang-v15 v0.0.0-20230222085438-ee3102fa0c71
	github.com/go-webgpu/goffi v0.6.4
	github.com/pekim/go-glfw v0.0.0-20260911115220-7ec34166a7ae
	github.com/stretchr/testify v1.11.1
	golang.org/x/text v0.41.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
