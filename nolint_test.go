package nolint_test

import (
	"testing"

	"github.com/gqlgo/gqlanalysis"
	"github.com/gqlgo/gqlanalysis/analysistest"
	"github.com/gqlgo/nolint"
	"github.com/vektah/gqlparser/v2/ast"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData(t)
	a := &gqlanalysis.Analyzer{
		Name: "nolint-test",
		Doc:  "nolint-test",
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
	analysistest.Run(t, testdata, a, "a")
}
