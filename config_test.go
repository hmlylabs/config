package config

import (
	"os"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) TestCreateNewConfigSucceeds() {
	os.Setenv("SERVER_NAME", "config")
	os.Setenv("PORT", "80")
	os.Setenv("DATABASE_URL", "postgres://database-url/5432")
	os.Setenv("SHADOW_URL", "postgres://database-url/shadow-url/5432")
	config, err := LoadConfig()
	utils.AssertEqual(s.T(), nil, err)
	utils.AssertEqual(s.T(), "config", config.ServerName)
	utils.AssertEqual(s.T(), "80", config.Port)
	utils.AssertEqual(s.T(), "postgres://database-url/5432", config.DatabaseUrl)

}

func (s *TestSuite) TestCreateNewConfigFails() {
	_, err := LoadConfig()
	s.NotNil(err)
}
