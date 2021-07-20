package xtime

import "time"

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp         = "Jan _2 15:04:05"
	StampMilli    = "Jan _2 15:04:05.000"
	StampMicro    = "Jan _2 15:04:05.000000"
	StampNano     = "Jan _2 15:04:05.000000000"
	YYmmdd1       = "20060102"       //20060102
	YYmmdd2       = "2006-01-02"     //2006-01-02
	YYmmddHHmmss1 = "20060102150405" //20060102150405
	YYmmddHHmm    = "2006-01-02 15:04"
	YYmmddHHmmss2 = "2006-01-02 15:04:05" //2006-01-02 15:04:05
	YYmmddHH      = "2006-01-02 15"       //2006-01-02 15:
	HHmmss        = "15:04:05"
)

// TimeParseString 接口名称：时间戳转字符传格式
//参数名称：format:
//YYmmdd1       = "20060102"
//YYmmdd2       = "2006-01-02"
//YYmmddHHmmss1 = "20060102150405"
//YYmmddHHmmss2 = "2006-01-02 15:04:05"
func TimeParseString(v int64, format string) string {
	result := time.Unix(v, 0).Format(format)
	return result
}

// StringParseUnix 转成时间戳
//format:
//YYmmdd1       = "20060102"
//YYmmdd2       = "2006-01-02"
//YYmmddHHmmss1 = "20060102150405"
//YYmmddHHmmss2 = "2006-01-02 15:04:05"
func StringParseUnix(v string, format string) (int64, error) {
	var (
		err error
		tm  time.Time
	)
	localTempleate := Templeate()
	if tm, err = time.ParseInLocation(format, v, localTempleate); err != nil {
		return 0, err
	}
	return tm.Unix(), nil
}

// StringParseTime 字符串转成Time
// m：时间 format：参照格式
func StringParseTime(m string, format string) (tm time.Time, err error) {
	tm, err = time.ParseInLocation(format, m, Templeate())
	if err != nil {
		return
	}
	return
}

// TimeFormat时间格式化
func TimeFormat(tm time.Time, format string) (time.Time, error) {
	return time.ParseInLocation(format, tm.Format(format), Templeate())
}

// CurrentTime 当期时间
func CurrentTime() int64 {
	localTempleate := Templeate()
	return time.Now().In(localTempleate).Unix()
}

const (
	// TimeTeplateChina 中国市区
	TimeTeplateChina = "Asia/Chongqing"
	// TimeTeplateAmerica 美国洛杉矶
	TimeTeplateAmerica = "America/Los_Angeles"
	//TimeTeplateServer 服务器市区
	TimeTeplateServer = "Local"
)

var (
	// TimeTeplateDefault 模板
	TimeTeplateDefault = TimeTeplateChina
)

// Templeate 时间模板
func Templeate() *time.Location {
	local, _ := time.LoadLocation(TimeTeplateDefault)
	return local
}
