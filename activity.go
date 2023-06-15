package sample

import (
	"encoding/json"
	"fmt"

	"github.com/project-flogo/core/activity"
	//"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	posMsg := input.PosMsg
	data := PositionMessage{}
	json.Unmarshal([]byte(posMsg), &data)


	fmt.Println(data.Mac)
	ctx.Logger().Info(data.Mac)

	fmt.Println("fmt.Println59")
	ctx.Logger().Info("ctx.Logger().Info60")

	output := &Output{X: data.Pos.X, Y: data.Pos.Y, MAC: data.Mac}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
