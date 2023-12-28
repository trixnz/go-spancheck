package spanlint_test

import (
	"regexp"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jjti/go-spanlint"
)

func Test(t *testing.T) {
	t.Parallel()

	for dir, config := range map[string]*spanlint.Config{
		"base": spanlint.NewDefaultConfig(),
		"enableall": {
			EnableSetStatusCheck:   true,
			EnableRecordErrorCheck: true,
		},
		"disableerrorchecks": {
			EnableSetStatusCheck:             true,
			IgnoreSetStatusCheckSignatures:   regexp.MustCompile("telemetry.Record"),
			EnableRecordErrorCheck:           true,
			IgnoreRecordErrorCheckSignatures: regexp.MustCompile("telemetry.Record"),
		},
	} {
		dir := dir
		t.Run(dir, func(t *testing.T) {
			analysistest.Run(t, "testdata/"+dir, spanlint.NewAnalyzer(config))
		})
	}
}
