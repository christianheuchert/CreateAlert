package SendMessageToAssets

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())
	input := &Input{IP: "52.45.17.177:802", CustomerId: "1", Username: "afadmin", Password: "admin", StaffIdList: `{"Id": 9064}`, Message: "Test Single-Message 5"}
	// StaffIdList "9064" OR "9064,37685" OR "{ID: 9064}"
	err := tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{} 
	err = tc.GetOutputObject(output)
	assert.Nil(t, err)

	assert.Equal(t, "true", output.Status)
}
