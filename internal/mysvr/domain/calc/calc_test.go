package calc_test

import (
	"backend/internal/mysvr/domain/calc"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CalcTestSuite struct {
	suite.Suite
}

func (s *CalcTestSuite) TestAdd() {
	var tests = []struct {
		a      int64
		b      int64
		result int64
	}{
		{1, 1, 2},
		{1, 2, 3},
	}

	for _, test := range tests {
		assert.Equal(s.T(), calc.Add(test.a, test.b), test.result)
	}
}

func TestCalc(t *testing.T) {
	suite.Run(t, new(CalcTestSuite))
}
