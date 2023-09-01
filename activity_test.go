package getAllByDept

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
	input := &Input{IP: "52.45.17.177:802", CustomerId: "1", Username: "afadmin", Password: "admin", DepartmentItem: "3098"}
	// DepartmentItem options: "Administration " OR "3098"
	// `{
	// 	"CustomerId": 123,
	// 	"DateCreated": "2023-08-09T12:34:56Z",
	// 	"DateUpdated": "2023-08-09T13:45:21Z",
	// 	"Description": "Example department",
	// 	"EnableTenancy": true,
	// 	"Name": "Engineering",
	// 	"TenantId": "987654",
	// 	"ElapsedTimeInMillseconds": 1234.567,
	// 	"ErrorMessage": "Error occurred",
	// 	"SuccessMessage": "Department created successfully",
	// 	"HasError": false,
	// 	"Id": 3098
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
