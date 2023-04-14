package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExperimentSearchApiFilterParsing(t *testing.T) {
	invalidTestCases := []string{
		"(",
		")",
		"())junk",
		")morejunk()",
		"(((junk(((()",
		")((((otherjunk)((()))))))",
	}
	for _, c := range invalidTestCases {
		_, err := parseFilter(c)
		require.Error(t, err)
	}
	validTestCases := [][2]string{
		{"", ""},
		{"()", "()"},
		{"()()", "()()"},
		{"(()())", "(()())"},
		{`string:"value"`, `string = 'value'`},
		{"-number:123456789", "number != 123456789"},
		{"anumber<=12.34", "anumber<=12.34"},
		{"thisnumber>=9.22", "thisnumber>=9.22"},
		{"validation.validation_accuracy>1", "(besttrials.best_validation->'metrics'->'avg_metrics'->>'validation_accuracy')::float8>1"},
		{"validation.validation_loss<-1", "(besttrials.best_validation->'metrics'->'avg_metrics'->>'validation_loss')::float8<-1"},
		{"validation.validation_test>-10.98", "(besttrials.best_validation->'metrics'->'avg_metrics'->>'validation_test')::float8>-10.98"},
		{`string:null`, `string IS NULL`},
		{`-value:null`, `value IS NOT NULL`},
		{`str~"like"`, `str LIKE '%like%'`},
		{`(str~"like" AND -otherstr~"notlike")`, `(str LIKE '%like%' AND otherstr NOT LIKE '%notlike%')`},
		{`(general_column.description~"experiment description" AND (-general_column.id:456 OR -general_column.resourcePool~"test\"s comma value\"s"))`, `(e.config->>'description' LIKE '%experiment description%' AND (e.id != 456 OR e.config->'resources'->>'resource_pool' NOT LIKE '%test\"s comma value\"s%'))`},
		{`(general_column.forkedFrom:5 OR (-validation.error:1 AND hp.hyperparameter<=10))`, `(e.parent_id = 5 OR ((besttrials.best_validation->'metrics'->'avg_metrics'->>'error')::float8 != 1 AND (e.config->'hyperparameters'->'hyperparameter'->>'val')::float8<=10))`},
	}
	for _, c := range validTestCases {
		result, err := parseFilter(c[0])
		require.NoError(t, err)
		require.Equal(t, c[1], result)
	}
}
