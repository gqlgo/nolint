package nolint

import (
	"reflect"
	"strings"

	"github.com/gqlgo/gqlanalysis"
)

const doc = "nolint set reporter to Pass which ignore a diagnostic with nolint comment"

var (
	flagOff       bool
	flagDirective string
)

func init() {
	Analyzer.Flags.BoolVar(&flagOff, "off", false, "true means nolint comment does not work")
	Analyzer.Flags.StringVar(&flagDirective, "directive", "nolint:", "prefix of nolint directive")
}

// Analyzer provides reporters which report diagnostic which is not related with any nolint comments.
//
//	var Analyzer = &gqlanalysis.Analyzer{
//		Name: "mylinter",
//		Doc:  "document",
//		Requires: []*gqlanalysis.Analyzer{
//			nolint.Analyzer,
//		},
//		Run: func(pass *gqlanalysis.Pass) (interface{}, error) {
//			pass.Report = pass.ResultOf[nolint.Analyzer].(*nolint.Reporters).New(pass)
//
//			return nil, nil
//		},
//	}
var Analyzer = &gqlanalysis.Analyzer{
	Name:       "nolint",
	Doc:        doc,
	Run:        run,
	ResultType: reflect.TypeOf((*Reporters)(nil)),
}

// Reporters creates a reporter which reports diagnostic which is not related with any nolint comments.
type Reporters struct {
	// TODO(tenntenn): comment map
}

// New creates a reporter which reports diagnostic which is not related with any nolint comments.
func (rs *Reporters) New(pass *gqlanalysis.Pass) func(*gqlanalysis.Diagnostic) {
	org := pass.Report
	return func(d *gqlanalysis.Diagnostic) {
		// -nolint.off
		if flagOff {
			org(d)
			return
		}

		for _, cmt := range pass.Comments {
			withoutMark := strings.TrimLeft(cmt.Value, "# ")
			// same line
			// TODO(tentenn): one line above
			if d.Pos.Line == cmt.Pos.Line &&
				strings.HasPrefix(withoutMark, flagDirective) {
				return
			}

			if d.Pos.Line <= cmt.Pos.Line {
				break
			}
		}

		org(d)
	}
}

func run(pass *gqlanalysis.Pass) (interface{}, error) {
	return new(Reporters), nil
}
