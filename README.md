# nolint

[![pkg.go.dev][gopkg-badge]][gopkg]

`nolint` is an analyzer which provides a reporter which ignores diagnostics with nolint comment.

```graphql
query GetA() {
    a {
	name # nolint: ignore it
    }
}
```

## How to use

[nolint.Analyzer](https://pkg.go.dev/github.com/gqlgo/nolint) can be set to `Requires` field of an analyzer. `(*nolint.Reporters).New` creates a reporter which ignore diagnostics with nolint comment. The reporter can be set to `(*gqlanalysis.Pass).Report` field.

```go
var Analyzer = &gqlanalysis.Analyzer{
	Name: "mylint",
	Doc:  "mylint",
	Requires: []*gqlanalysis.Analyzer{
		nolint.Analyzer,
	},
	Run: func(pass *gqlanalysis.Pass) (interface{}, error) {
		pass.Report = pass.ResultOf[nolint.Analyzer].(*nolint.Reporters).New(pass)

		for _, q := range pass.Queries {
			for _, f := range q.Fragments {
				for _, s := range f.SelectionSet {
					field, _ := s.(*ast.Field)
					if field != nil {
						pass.Reportf(field.Position, "NG")
					}
				}
			}
		}

		return nil, nil
	},
}
```

## Author

[![Appify Technologies, Inc.](appify-logo.png)](http://github.com/appify-technologies)

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/nolint
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/nolint?status.svg
