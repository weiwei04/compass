package clicommands

import (
	"errors"
	"testing"
)

var err = errors.New("error")

func Test_splitSpaceChart(t *testing.T) {
	cases := []struct {
		Input string
		Space string
		Chart string
		Err   error
	}{
		{
			Input: "weiwei04/redis",
			Space: "weiwei04",
			Chart: "redis",
			Err:   nil,
		},
		{
			Input: "weiwei04",
			Space: "weiwei04",
			Chart: "",
			Err:   nil,
		},
		{
			Input: "hub/weiwei04/redis",
			Space: "",
			Chart: "",
			Err:   err,
		},
	}

	for _, c := range cases {
		space, chart, err := splitSpaceChart(c.Input)
		if c.Err == nil {
			if err != nil {
				t.Errorf("splitSpaceChart(%s), expect err:nil, go err:%s",
					c.Input, err)
			}
			if space != c.Space {
				t.Errorf("splitSpaceChart(%s), expect space:%s, got space:%s",
					c.Input, c.Space, space)
			}
			if chart != c.Chart {
				t.Errorf("splitSpaceChart(%s), expect chart:%s, got chart:%s",
					c.Input, c.Chart, chart)
			}
		} else {
			if err == nil {
				t.Errorf("splitSpaceChart(%s), expect err occurred, got err:nil",
					c.Input)
			}
		}
	}
}

func Test_splitSpaceChartVer(t *testing.T) {
	cases := []struct {
		Input   string
		Space   string
		Chart   string
		Version string
		Err     error
	}{
		{
			Input:   "myspace/mychart:0.1.0",
			Space:   "myspace",
			Chart:   "mychart",
			Version: "0.1.0",
			Err:     nil,
		},
		{
			Input: "myspace/mychart",
			Err:   err,
		},
		{
			Input: "myspace",
			Err:   err,
		},
		{
			Input: "mychart:0.1.0",
			Err:   err,
		},
	}

	for _, c := range cases {
		space, chart, ver, err := splitSpaceChartVer(c.Input)
		if c.Err == nil {
			if err != nil {
				t.Errorf("splitSpaceChartVer(%s), expect err:nil, go err:%s",
					c.Input, err)
			}
			if space != c.Space {
				t.Errorf("splitSpaceChartVer(%s), expect space:%s, got space:%s",
					c.Input, c.Space, space)
			}
			if chart != c.Chart {
				t.Errorf("splitSpaceChartVer(%s), expect chart:%s, got chart:%s",
					c.Input, c.Chart, chart)
			}
			if ver != c.Version {
				t.Errorf("splitSpaceChartVer(%s), expect version:%s, got version:%s",
					c.Input, c.Version, ver)
			}
		} else {
			if err == nil {
				t.Errorf("splitSpaceChartVer(%s), expect err occurred, got err:nil",
					c.Input)
			}
		}
	}
}
