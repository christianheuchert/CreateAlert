package getAllByGroup

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
	input := &Input{IP: "52.45.17.177:802", CustomerId: "1", Username: "afadmin", Password: "admin", GroupItem: "5620"}
	// GroupItem = "5620" OR "Assets" OR 
	// `{
    //   "Icon": "assets\\icons\\792794ff-17db-46bc-b1fd-8110cf81f061.png",
    //   "MultiAssign": false,
    //   "CustomerId": 0,
    //   "DateCreated": "2023-06-22T16:40:25.43",
    //   "DateUpdated": "2023-06-22T16:40:25.43",
    //   "Description": "",
    //   "EnableTenancy": false,
    //   "Name": "Acute Care",
    //   "TenantId": "",
    //   "ElapsedTimeInMillseconds": 0,
    //   "ErrorMessage": "",
    //   "SuccessMessage": "",
    //   "HasError": false,
    //   "Id": 5644
    // }`

	err := tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{} 
	err = tc.GetOutputObject(output)
	assert.Nil(t, err)

	assert.NotNil(t, output.Staff)
}
