package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/ushios/gengo"
)

type handlerTestSuite struct {
	suite.Suite
}

func Test_handlerTestSuite(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}

func (suite *handlerTestSuite) Test_convert() {
	t := gengo.Meiji.StartAt()
	gengo, no, err := convert(t)
	suite.Assert().NoError(err)
	suite.Assert().Equal("明治", gengo)
	suite.Assert().Equal(no, 1)
}

func (suite *handlerTestSuite) Test_convert_timeZero() {
	t := time.Time{}
	_, _, err := convert(t)
	suite.Assert().Error(err)
}

func (suite *handlerTestSuite) Test_parseTime() {
	expected, err := time.Parse("2006-01-02", "2018-02-12")
	suite.Assert().NoError(err)
	actual, err := parseTime("2018-02-12")
	suite.Assert().NoError(err)
	suite.Assert().Equal(expected, actual)
}

func (suite *handlerTestSuite) Test_parseTime_complete() {
	expected, err := time.Parse("2006-01-02", "2018-01-01")
	suite.Assert().NoError(err)
	actual, err := parseTime("2018-XX-XX")
	suite.Assert().NoError(err)
	suite.Assert().Equal(expected, actual)
}
