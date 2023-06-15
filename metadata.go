package sample

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	PosMsg string `md:"posMsg,required"` // AoA-postion mqtt msg
}

func (i *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["PosMsg"])
	i.PosMsg = strVal
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"PosMsg": i.PosMsg,
	}
}

type Output struct {
	X int `md:"X"`
	Y int `md:"Y"`
	MAC string `md:"MAC"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error
	o.X, err = coerce.ToInt(values["x"])
	if err != nil {
		return err
	}
	o.Y, err = coerce.ToInt(values["y"])
	if err != nil {
		return err
	}

	o.MAC, err = coerce.ToString(values["MAC"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"x":    o.X,
		"y": o.Y,
		"MAC":   o.MAC,
	}
}

type PositionMessage struct {
	Mac string `json:"MAC"`
	Pos struct {
		MapID string `json:"mapID"`
		X     int    `json:"x"`
		Y     int    `json:"y"`
	} `json:"pos"`
	Status struct {
		Rpt  int `json:"rpt"`
		Sfsw int `json:"sfsw"`
		Chg  int `json:"chg"`
		Batt int `json:"batt"`
		Sqn  int `json:"sqn"`
	} `json:"status"`
}
