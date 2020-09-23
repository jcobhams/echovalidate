package echovalidate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequired(t *testing.T) {
	key := "username"
	//Success Case
	err := Required(key, "jcobhams")
	assert.Nil(t, err)

	//Error Case
	err = Required(key, "")
	assert.NotNil(t, err)
}

func TestValidEmail(t *testing.T) {
	key := "email"
	//Success Case
	err := ValidEmail(key, "someone@domain.com")
	assert.Nil(t, err)

	//Error Cases
	errorTestCases := []struct {
		Input string
	}{
		{Input: ""},
		{Input: "someone"},
		{Input: "someone@"},
		{Input: "someone@domain"},
		{Input: "someone@domain."},
	}

	for _, testCase := range errorTestCases {
		err = ValidEmail(key, testCase.Input)
		assert.NotNil(t, err)
	}

}

func TestIn(t *testing.T) {
	key := "role"

	roles := []string{
		"ADMIN",
		"MODERATOR",
		"SUPERVISOR",
		"AUDIT",
		"COMPLIANCE",
	}

	//Success Case
	err := In(key, "ADMIN", roles)
	assert.Nil(t, err)

	errorTestCases := []struct {
		Input string
	}{
		{Input: "A DMIN"},
		{Input: "LOGISTICS"},
		{Input: ""},
		{Input: " "},
	}

	for _, testCase := range errorTestCases {
		err = In(key, testCase.Input, roles)
		assert.NotNil(t, err)
	}
}

func TestMinLen(t *testing.T) {
	key := "key"
	minLen := 10

	//Success Case
	err := MinLen(key, 10, minLen)
	assert.Nil(t, err)

	//Error Cases
	errorTestCases := []struct {
		Input int
	}{
		{Input: 9},
		{Input: 0},
		{Input: 1},
	}

	for _, testCase := range errorTestCases {
		err = MinLen(key, testCase.Input, minLen)
		assert.NotNil(t, err)
	}
}

func TestMaxLen(t *testing.T) {
	key := "key"
	maxLen := 10

	//Success Case
	err := MaxLen(key, 9, maxLen)
	assert.Nil(t, err)

	//Error Cases
	errorTestCases := []struct {
		Input int
	}{
		{Input: 11},
		{Input: 12},
		{Input: 100},
	}

	for _, testCase := range errorTestCases {
		err = MaxLen(key, testCase.Input, maxLen)
		assert.NotNil(t, err)
	}
}

func TestValidMongoObjectID(t *testing.T) {
	key := "_id"

	//Success Case
	err := ValidMongoObjectID(key, "5efe124b93f2c8737bc82042")
	assert.Nil(t, err)

	//Error Cases
	errorTestCases := []struct {
		Input string
	}{
		{Input: ""},
		{Input: " "},
		{Input: "a2B3C4D5E6F7G8H9I0J1K2L3M"},
	}

	for _, testCase := range errorTestCases {
		err = ValidMongoObjectID(key, testCase.Input)
		assert.NotNil(t, err)
	}
}

func TestValidDateTime(t *testing.T) {
	key := "date"

	//Success Case
	err := ValidDateTime(key, "2020-01-20", "2006-01-02")
	assert.Nil(t, err)

	// Error Cases
	errorTestCases := []struct {
		Value  string
		Layout string
	}{
		{Value: "2020-25-01", Layout: "2006-01-02"},
		{Value: "", Layout: "2006-01-02"},
		{Value: " ", Layout: "2006-01-02"},
	}

	for _, testCase := range errorTestCases {
		err = ValidDateTime(key, testCase.Value, testCase.Layout)
		assert.NotNil(t, err)
	}
}

func TestValidation_Validate(t *testing.T) {
	v := New()

	//Success Case
	rules := Validator{
		Rules{
			{Required, "username", "jcobhams"},
			{MinLen, "min", 10, 9},
		},
	}

	err := v.Validate(rules)
	assert.Nil(t, err)

	//Error Case
	rules = Validator{
		Rules{
			{Required, "username", ""},
			{MinLen, "min", 10, 9},
		},
	}

	err = v.Validate(rules)
	assert.NotNil(t, err)
}
