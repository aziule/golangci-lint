package golinters

import (
	"github.com/aziule/filebuildtag"
	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"golang.org/x/tools/go/analysis"
)

func NewFilebuildtag(settings *config.FilebuildtagSettings) *goanalysis.Linter {
	analyzer := filebuildtag.Analyzer
	var cfg map[string]map[string]interface{}
	if settings != nil {
		cfg = map[string]map[string]interface{}{
			analyzer.Name: {
				filebuildtag.FlagFiletagsName: settings.Filetags,
			},
		}
	}
	return goanalysis.NewLinter(
		analyzer.Name,
		analyzer.Doc,
		[]*analysis.Analyzer{analyzer},
		cfg,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
