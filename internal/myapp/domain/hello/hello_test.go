package hello_test

import (
	"backend/internal/myapp/domain/hello"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HelloTestSuite struct {
	suite.Suite
}

func (s *HelloTestSuite) TestHello() {
	var tests = []struct {
		name     string
		expected string
	}{
		{"", "Hello!\n"},
		{"Mike", "Hello Mike!\n"},
	}

	for _, test := range tests {
		assert.Equal(s.T(), hello.Hello(test.name), test.expected)
	}
}

func TestHello(t *testing.T) {
	suite.Run(t, new(HelloTestSuite))
}
