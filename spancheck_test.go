package spancheck_test

import (
	"regexp"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/jjti/go-spancheck"
)

func Test(t *testing.T) {
	t.Parallel()

	for dir, config := range map[string]*spancheck.Config{
		"base": spancheck.NewConfig(),
		"disableerrorchecks": {
			EnableSetStatusCheck:             true,
			IgnoreSetStatusCheckSignatures:   regexp.MustCompile("telemetry.Record"),
			EnableRecordErrorCheck:           true,
			IgnoreRecordErrorCheckSignatures: regexp.MustCompile("telemetry.Record"),
		},
		"enableall": {
			EnableAll: true,
		},
		"enablechecks": {
			EnableSetStatusCheck:   true,
			EnableRecordErrorCheck: true,
		},
	} {
		dir := dir
		t.Run(dir, func(t *testing.T) {
			analysistest.Run(t, "testdata/"+dir, spancheck.NewAnalyzerWithConfig(config))
		})
	}
}