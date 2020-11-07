package controller

import (
	"errors"
	"fmt"
	"strconv"

	coordTransform "github.com/qichengzx/coordtransform"
	"github.com/urfave/cli/v2"
)

//TargetCoordinate coordinate 目标坐标系
var TargetCoordinate string

//Trans 火星转真实
func Trans(c *cli.Context) error {
	if c.Args().Len() == 2 {
		lng, err := strconv.ParseFloat(c.Args().First(), 64)
		lat, err := strconv.ParseFloat(c.Args().Get(1), 64)
		if err != nil {
			return err
		}
		resLng, resLat := 0.0, 0.0
		tag := "gcj02"
		switch TargetCoordinate {
		case "w2g":
			resLng, resLat = coordTransform.WGS84toGCJ02(lng, lat)
			break
		case "g2w":
			tag = "wgs84"
			resLng, resLat = coordTransform.GCJ02toWGS84(lng, lat)
			break
		case "w2b":
			tag = "bd09"
			resLng, resLat = coordTransform.WGS84toBD09(lng, lat)
			break
		case "b2w":
			tag = "wgs84"
			resLng, resLat = coordTransform.BD09toWGS84(lng, lat)
			break
		case "g2b":
			tag = "bd09"
			resLng, resLat = coordTransform.GCJ02toBD09(lng, lat)
			break
		case "b2g":
			resLng, resLat = coordTransform.BD09toGCJ02(lng, lat)
			break
		default:
			return errors.New("please input correct coordinate target")
		}

		fmt.Printf("%s Lng: %0.9f", tag, resLng)
		fmt.Println()
		fmt.Printf("%s Lat: %0.9f", tag, resLat)
	}
	return nil
}
