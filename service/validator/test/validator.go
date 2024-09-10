package test

type SimpleTestStruct struct {
	Title string `validate:"required"`
	Body  string `validate:"max=5"`
}

func (suite *PackageTestSuite) TestValidatorValidateValid() {
	given := &SimpleTestStruct{
		Title: "test",
		Body:  "AAA",
	}

	err := suite.validator.Validate(given)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestValidatorValidateInvalid() {
	given := &SimpleTestStruct{
		Title: "",
		Body:  "AAAAAAA",
	}

	err := suite.validator.Validate(given)

	suite.Error(err)
}
