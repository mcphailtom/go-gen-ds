package tree

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TreeTestSuite struct {
	suite.Suite
}

func (tts *TreeTestSuite) SetupSuite() {
}

func TestTreeTestSuite(t *testing.T) {
	suite.Run(t, new(TreeTestSuite))
}
