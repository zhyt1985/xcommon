package _map

import (
	"errors"
	"math"
)

const (
	AMapDistrictURL = "https://restapi.amap.com/v3/config/district?" // 行政区域查询
	AmapKey         = "47ab0ee1accab6e91611a439e3445645"             // 高德key
)
const (
	//BDTOGD ...
	BDTOGD = "BDTOGD"
	//GDTOBD ...
	GDTOBD = "GDTOBD"
	//WGS84TOBD ...
	WGS84TOBD = "WGS84TOBD"
	//BDTOWGS84 ...
	BDTOWGS84 = "BDTOWGS84"
	//GDTOWGS84 ...
	GDTOWGS84 = "GDTOWGS84"
	//WGS84TOGD ...
	WGS84TOGD = "WGS84TOGD"

	// GDTOMERCATOR 高德转莫卡托
	GDTOMERCATOR = "GDTOMERCATOR"
	// MERCATORTOGD 莫卡托转高德
	MERCATORTOGD = "MERCATORTOGD"

	//经纬度转莫卡托
	LLTOMERCATOR = "LLTOMERCATOR"
	//莫卡托转经纬度
	MERCATORTOLL = "MERCATORTOLL"
)

//TransformLatLng ...
func TransformLatLng(coorType string, lng, lat float64) (float64, float64) {
	switch coorType {
	case BDTOGD:
		lng, lat = BD09toGCJ02(lng, lat)
	case GDTOWGS84:
		lng, lat = GCJ02toWGS84(lng, lat)
	default:
	}
	return lng, lat
}

// WGS84坐标系：即地球坐标系，国际上通用的坐标系。
// GCJ02坐标系：即火星坐标系，WGS84坐标系经加密后的坐标系。Google Maps，高德在用。
// BD09坐标系：即百度坐标系，GCJ02坐标系经加密后的坐标系。

const (
	X_PI   = math.Pi * 3000.0 / 180.0
	OFFSET = 0.00669342162296594323
	AXIS   = 6378245.0
)

//GCJ02toWGS84 火星坐标系->WGS84坐标系
func GCJ02toWGS84(lon, lat float64) (float64, float64) {
	if isOutOFChina(lon, lat) {
		return lon, lat
	}

	mgLon, mgLat := delta(lon, lat)

	return lon*2 - mgLon, lat*2 - mgLat
}

//
func isOutOFChina(lon, lat float64) bool {
	return !(lon > 72.004 && lon < 135.05 && lat > 3.86 && lat < 53.55)
}

//
func delta(lon, lat float64) (float64, float64) {
	dlat, dlon := transform(lon-105.0, lat-35.0)
	radlat := lat / 180.0 * math.Pi
	magic := math.Sin(radlat)
	magic = 1 - OFFSET*magic*magic
	sqrtmagic := math.Sqrt(magic)

	dlat = (dlat * 180.0) / ((AXIS * (1 - OFFSET)) / (magic * sqrtmagic) * math.Pi)
	dlon = (dlon * 180.0) / (AXIS / sqrtmagic * math.Cos(radlat) * math.Pi)

	mgLat := lat + dlat
	mgLon := lon + dlon

	return mgLon, mgLat
}

//
func transform(lon, lat float64) (x, y float64) {
	var lonlat = lon * lat
	var absX = math.Sqrt(math.Abs(lon))
	var lonPi, latPi = lon * math.Pi, lat * math.Pi
	var d = 20.0*math.Sin(6.0*lonPi) + 20.0*math.Sin(2.0*lonPi)
	x, y = d, d
	x += 20.0*math.Sin(latPi) + 40.0*math.Sin(latPi/3.0)
	y += 20.0*math.Sin(lonPi) + 40.0*math.Sin(lonPi/3.0)
	x += 160.0*math.Sin(latPi/12.0) + 320*math.Sin(latPi/30.0)
	y += 150.0*math.Sin(lonPi/12.0) + 300.0*math.Sin(lonPi/30.0)
	x *= 2.0 / 3.0
	y *= 2.0 / 3.0
	x += -100.0 + 2.0*lon + 3.0*lat + 0.2*lat*lat + 0.1*lonlat + 0.2*absX
	y += 300.0 + lon + 2.0*lat + 0.1*lon*lon + 0.1*lonlat + 0.1*absX
	return
}

//BD09toGCJ02 百度坐标系->火星坐标系
func BD09toGCJ02(lon, lat float64) (float64, float64) {
	x := lon - 0.0065
	y := lat - 0.006

	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*X_PI)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*X_PI)

	gLon := z * math.Cos(theta)
	gLat := z * math.Sin(theta)

	return gLon, gLat
}

var (
	EARTHRADIUS = 6370996.81
	MCBAND      = []float64{12890594.86, 8362377.87, 5591021, 3481989.83, 1678043.12, 0}
	LLBAND      = []float64{75, 60, 45, 30, 15, 0}
	MC2LL       = [][]float64{
		{1.410526172116255e-008, 8.983055096488720e-006, -1.99398338163310, 2.009824383106796e+002, -1.872403703815547e+002, 91.60875166698430, -23.38765649603339, 2.57121317296198, -0.03801003308653, 1.733798120000000e+007},
		{-7.435856389565537e-009, 8.983055097726239e-006, -0.78625201886289, 96.32687599759846, -1.85204757529826, -59.36935905485877, 47.40033549296737, -16.50741931063887, 2.28786674699375, 1.026014486000000e+007},
		{-3.030883460898826e-008, 8.983055099835780e-006, 0.30071316287616, 59.74293618442277, 7.35798407487100, -25.38371002664745, 13.45380521110908, -3.29883767235584, 0.32710905363475, 6.856817370000000e+006},
		{-1.981981304930552e-008, 8.983055099779535e-006, 0.03278182852591, 40.31678527705744, 0.65659298677277, -4.44255534477492, 0.85341911805263, 0.12923347998204, -0.04625736007561, 4.482777060000000e+006},
		{3.091913710684370e-009, 8.983055096812155e-006, 0.00006995724062, 23.10934304144901, -0.00023663490511, -0.63218178102420, -0.00663494467273, 0.03430082397953, -0.00466043876332, 2.555164400000000e+006},
		{2.890871144776878e-009, 8.983055095805407e-006, -0.00000003068298, 7.47137025468032, -0.00000353937994, -0.02145144861037, -0.00001234426596, 0.00010322952773, -0.00000323890364, 8.260885000000000e+005},
	}
	LL2MC = [][]float64{
		{-0.00157021024440, 1.113207020616939e+005, 1.704480524535203e+015, -1.033898737604234e+016,
			2.611266785660388e+016, -3.514966917665370e+016, 2.659570071840392e+016, -1.072501245418824e+016,
			1.800819912950474e+015, 82.50000000000000},
		{8.277824516172526e-004, 1.113207020463578e+005, 6.477955746671608e+008, -4.082003173641316e+009,
			1.077490566351142e+010, -1.517187553151559e+010, 1.205306533862167e+010, -5.124939663577472e+009,
			9.133119359512032e+008, 67.50000000000000},
		{0.00337398766765, 1.113207020202162e+005, 4.481351045890365e+006, -2.339375119931662e+007,
			7.968221547186455e+007, -1.159649932797253e+008, 9.723671115602145e+007, -4.366194633752821e+007,
			8.477230501135234e+006, 52.50000000000000},
		{0.00220636496208, 1.113207020209128e+005, 5.175186112841131e+004, 3.796837749470245e+006, 9.920137397791013e+005,
			-1.221952217112870e+006, 1.340652697009075e+006, -6.209436990984312e+005, 1.444169293806241e+005, 37.50000000000000},
		{-3.441963504368392e-004, 1.113207020576856e+005, 2.782353980772752e+002, 2.485758690035394e+006, 6.070750963243378e+003,
			5.482118345352118e+004, 9.540606633304236e+003, -2.710553267466450e+003, 1.405483844121726e+003, 22.50000000000000},
		{-3.218135878613132e-004, 1.113207020701615e+005, 0.00369383431289, 8.237256402795718e+005, 0.46104986909093,
			2.351343141331292e+003, 1.58060784298199, 8.77738589078284, 0.37238884252424, 7.45000000000000},
	}
)

func ConvertLL2MC(lng float64, lat float64) (float64, float64, error) {
	var factors []float64
	for index, band := range LLBAND {
		if lat >= band {
			factors = LL2MC[index]
			break
		}
	}
	if len(factors) > 0 {
		var i int
		for i = len(LLBAND) - 1; i >= 0; i-- {
			if lat <= -LLBAND[i] {
				factors = LL2MC[i]
			}
			i = i - 1
		}
	}
	if len(factors) == 0 {
		return 0, 0, errors.New("out of bounds")
	}
	return Convertor(lng, lat, factors)
}

func Convertor(lng float64, lat float64, factors []float64) (float64, float64, error) {
	var (
		x    = factors[0] + factors[1]*math.Abs(lng)
		temp = math.Abs(lat) / factors[9]
	)
	y := factors[2] +
		factors[3]*temp +
		factors[4]*temp*temp +
		factors[5]*temp*temp*temp +
		factors[6]*temp*temp*temp*temp +
		factors[7]*temp*temp*temp*temp*temp +
		factors[8]*temp*temp*temp*temp*temp*temp
	if lng < 0 {
		x *= -1
	} else {
		x *= 1
	}
	if lat < 0 {
		y *= -1
	} else {
		y *= 1
	}
	return x, y, nil
}
