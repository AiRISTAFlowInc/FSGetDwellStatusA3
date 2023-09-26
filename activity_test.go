package getDwellStatus

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
	input := &Input{IP: "3.212.201.170:802", CustomerId: "2047", Username: "afadmin", Password: "admin", MAC: "C4:CB:6B:23:24:D0",GracePeriod: "2", ZoneItem: "Patient Room"}
	// ZoneItem: "26960" OR "{\"ZoneID\":26960,\"ZoneName\":\"Entrance\",\"ZoneType\":\"Open\"}" OR "Entrance"
	//TESTING TIMES MUST BE SET in the activty to have true returned
		// StartDateTime = "2023-07-27 16:47:48.807"
		// EndDateTime = "2023-07-27 16:57:48.807"
		// //TESTING TIMES
	err := tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{} 
	err = tc.GetOutputObject(output)
	assert.Nil(t, err)

	assert.Equal(t, true, output.DwellStatus)
}
