package ParseUnknownXpertMessage

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
	input := &Input{
		XpertMessageJSON: "{\"MAC\":\"C4:CB:6B:22:1B:DA\",\"pos\":{\"mapID\":\"14741\",\"x\":1168,\"y\":565},\"status\":{\"rpt\":0,\"sfsw\":0,\"chg\":0,\"batt\":40,\"sqn\":0}}",
		DeviceMACTarget: "XpertMessage.MAC",
		TimestampTarget: "",
		DeviceLogIdTarget: "",
		StatusReportReasonTarget: "",
		BatteryLevelTarget: "XpertMessage.status.batt",
		TemperatureTarget: "",
		HumidityTarget: "",
		MapIdTarget: "XpertMessage.pos.mapID",
		XTarget: "XpertMessage.pos.x",
		YTarget: "XpertMessage.pos.y",
		ZoneTarget: "",
		GeoLattitudeTarget: "",
		GeoLongitudeTarget: "",
		ItemIdTarget: "",
		DisplayNameTarget: "",
	}
	err := tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{}
	err = tc.GetOutputObject(output)
	assert.Nil(t, err)
	assert.Equal(t, "C4:CB:6B:22:1B:DA", output.DeviceMAC)
}
