// +build integration

package test

import (
	"context"
	"crypto/rand"
	"fmt"

	"go-template/config"
	"go-template/repository/mongodb"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

type SimpleTestStruct struct {
	ID     string   `bson:"_id,omitempty"`
	Title  string   `bson:"title"`
	List   []string `bson:"list"`
	Text   *string  `bson:"text"`
	Number *int     `bson:"number"`
}

type PackageTestSuite struct {
	suite.Suite
	ctx    context.Context
	repo   *mongodb.Repository
	config *mongodb.Config
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	conf := config.New()
	suite.config = &mongodb.Config{
		URI:      conf.MongoDBEndpoint,
		DBName:   fmt.Sprintf("%s-test", conf.MongoDBName),
		CollName: "test",
		Timeout:  conf.MongoDBTimeout,
	}

	var err error
	suite.repo, err = mongodb.New(suite.ctx, suite.config)
	suite.NoError(err)
}

func (suite *PackageTestSuite) SetupTest() {
	var err error
	suite.repo, err = mongodb.New(suite.ctx, suite.config)
	suite.NoError(err)
}

func (suite *PackageTestSuite) TearDownTest() {
	_, _ = suite.repo.Coll.DeleteMany(suite.ctx, bson.M{})
}

func (suite *PackageTestSuite) TearDownSuite() {
	_ = suite.repo.DB.Drop(suite.ctx)
}

func (suite *PackageTestSuite) makeTestStruct(title string, list ...string) (test *SimpleTestStruct) {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	uuid := fmt.Sprintf("%x", b)

	return &SimpleTestStruct{
		ID:    uuid,
		Title: title,
		List:  list,
	}
}
