package strutil

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/dlclark/regexp2"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"math/big"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// TODO -12 不能转到 12
func StrToUInt(num string) uint {
	return uint(StrToInt(num))
}

func StrToUInt8(num string) uint8 {
	return uint8(StrToInt(num))
}

func StrToUInt16(num string) uint16 {
	return uint16(StrToInt(num))
}

func StrToUInt32(num string) uint32 {
	return uint32(StrToInt(num))
}

func StrToUInt64(num string) uint64 {
	return uint64(StrToInt(num))
}

// 字符串转int
func StrToInt(num string) int {
	result, err := strconv.Atoi(num)

	if err != nil {
		float, err := strconv.ParseFloat(num, 64)
		if err == nil {
			return int(float)
		}

		return 0
	}

	return result
}

// 字符串转int32
func StrToInt8(num string) int8 {
	result, err := strconv.ParseInt(num, 10, 8)
	if err != nil {
		float, err := strconv.ParseFloat(num, 8)
		if err == nil {
			return int8(float)
		}

		return 0
	}

	return int8(result)
}

func StrToInt16(num string) int16 {
	result, err := strconv.ParseInt(num, 10, 16)
	if err != nil {
		float, err := strconv.ParseFloat(num, 16)
		if err == nil {
			return int16(float)
		}

		return 0
	}

	return int16(result)
}

func StrToInt32(num string) int32 {
	result, err := strconv.ParseInt(num, 10, 32)
	if err != nil {
		float, err := strconv.ParseFloat(num, 32)
		if err == nil {
			return int32(float)
		}

		return 0
	}

	return int32(result)
}

// 字符串转int64
func StrToInt64(num string) int64 {
	result, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		float, err := strconv.ParseFloat(num, 64)
		if err == nil {
			return int64(float)
		}

		return 0
	}

	return result
}

// int到字符串
func IntToStr(num int) string {
	return strconv.Itoa(num)
}

// int64到字符串
func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

// 任意类型转字符串
func InterfaceToByte(v interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

// 任意类型转字符串
func InterfaceToStr(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

//任意类型转数字

func InterfaceToUInt(v interface{}) uint {
	return StrToUInt(fmt.Sprintf("%v", v))
}

func InterfaceToUInt8(v interface{}) uint8 {
	return StrToUInt8(fmt.Sprintf("%v", v))
}

func InterfaceToUInt16(v interface{}) uint16 {
	return StrToUInt16(fmt.Sprintf("%v", v))
}

func InterfaceToUInt32(v interface{}) uint32 {
	return StrToUInt32(fmt.Sprintf("%v", v))
}

func InterfaceToUInt64(v interface{}) uint64 {
	return StrToUInt64(fmt.Sprintf("%v", v))
}

func InterfaceToInt(v interface{}) int {
	return StrToInt(fmt.Sprintf("%v", v))
}

func InterfaceToInt8(v interface{}) int8 {
	return StrToInt8(fmt.Sprintf("%v", v))
}

func InterfaceToInt16(v interface{}) int16 {
	return StrToInt16(fmt.Sprintf("%v", v))
}

func InterfaceToInt32(v interface{}) int32 {
	return StrToInt32(fmt.Sprintf("%v", v))
}

func InterfaceToInt64(v interface{}) int64 {
	return cast.ToInt64(v)
}

// 转float64
func InterfaceToFloat32(v interface{}) float32 {
	f, e := strconv.ParseFloat(fmt.Sprintf("%v", v), 32)

	if e != nil {
		return 0
	}

	return float32(f)
}

func InterfaceToFloat64(v interface{}) float64 {
	f, e := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)

	if e != nil {
		return 0
	}

	return f
}

func InterfaceToDecimal(v interface{}) decimal.Decimal {
	d, err := decimal.NewFromString(fmt.Sprintf("%v", v))
	if err != nil {
		return decimal.Zero
	}
	return d
}

// 转big.int
func InterfaceToBigInt(v interface{}) *big.Int {
	switch v.(type) {
	case float32, float64:
		return big.NewInt(InterfaceToInt64(v))
	case int, int32, int64:
		return big.NewInt(InterfaceToInt64(v))
	default:
		num := new(big.Int)
		num.SetString(fmt.Sprintf("%v", v), 10)
		return num
	}

	return nil
}

// 16进制字符串转int64
func HexToInt64(hex string) int64 {
	if strings.HasPrefix(hex, "0x") {
		hex = strings.TrimLeft(hex, "0x")
	}

	result, err := strconv.ParseInt(hex, 16, 64)

	if err != nil {
		return 0
	}

	return result
}

// 16进制字符串转BigInt
func HexToBigInt(hex string) *big.Int {
	if strings.HasPrefix(hex, "0x") {
		hex = strings.TrimLeft(hex, "0x")
	}

	num := new(big.Int)
	num.SetString(hex, 16)

	return num
}

// 产生指定长度的数字随机数
func RandomNumStr(len int) (randomNum string) {
	var buffer bytes.Buffer
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		buffer.WriteString(fmt.Sprintf("%d", rand1.Int()%10))
	}

	return buffer.String()
}

// 产生指定长度的16进制字符串
func RandomHexStr(len int) (str string) {
	var hexArr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	var buffer bytes.Buffer

	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		buffer.WriteString(hexArr[rand1.Int()%16])
	}

	return buffer.String()
}

// 将字符串填充到指定位数
func FillToLen(str string, length int) string {
	return FillToLenByChar(str, length, "0")
}

func FillToLenByChar(str string, length int, char string) string {
	fillStr := str

	for i := 0; i < length-len(str); i++ {
		fillStr = fmt.Sprintf("%s%s", char, fillStr)
	}

	return fillStr
}

// 隐藏字符串以*代替
// 包含fromIndex，但是不包含toIndex 和 slice的规则一致
func HidStr(str string, prefixLen int, suffixLen int) string {
	var hidStrCount = 0
	var buffer bytes.Buffer
	for i := 0; i < len(str); i++ {
		if i >= prefixLen && i < len(str)-suffixLen {
			if hidStrCount < 4 {
				hidStrCount++
				buffer.WriteString("*")
			}
		} else {
			buffer.WriteString(str[i : i+1])
		}
	}

	return buffer.String()
}

// 下划线转大驼峰
func UnderscoreToCamel(name string) string {

	if name == "" {
		return ""
	}

	temp := strings.Split(name, "_")
	var s string
	for _, v := range temp {
		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z') { //首字母大写
				vv[0] -= 32
			}
			s += string(vv)
		}
	}

	return s
}

/*
*
struct 转 map
*/
func StructToMap(obj interface{}) (data map[string]interface{}) {
	bytes, _ := jsoniter.Marshal(obj)
	jsoniter.Unmarshal(bytes, &data)
	return data
}

var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

// 去除字符串中的中文字符
func StrFilterChinese(src string) string {

	str := ""
	for _, c := range src {
		if !hzRegexp.MatchString(string(c)) {
			str += string(c)
		}
	}

	return str
}

// 提取字符串中的数字 支持提取 浮点数
func StrFilterNum(src string) string {

	str := ""
	j := 0
	a := ""
	for i, c := range src {
		_, err := strconv.Atoi(string(c))
		// 是数字
		if err == nil {
			str += string(c)
			j = i
		} else {
			//非数字，则判断是否是. 小数点，如果j 即上一行是数字，则将.加入
			if string(c) == "." && a != "." {
				if i-j == 1 && a != "." {
					str += string(c)
					j = i
					a = string(c)
				}
			}
		}
	}

	return str
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// Udl2LCC 下划线转小驼峰 underline to lower camel case
func Udl2LCC(s string) (camelCase string) {
	camelCase = FirstLower(Udl2UCC(s))
	return camelCase
}

// Udl2UCC 下划线转大驼峰 underline to upper camel case
func Udl2UCC(s string) (camelCase string) {
	isToUpper := false
	for k, v := range s {
		if k == 0 {
			camelCase = strings.ToUpper(string(s[0]))
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return FirstUpper(camelCase)
}

// Cc2Udl 大，小驼峰转下划线
func Cc2Udl(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// StrIsNumber 判断字符串是否纯数字
func StrIsNumber(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsNumber(x) {
			return false
		}
	}
	return true
}

// 将map转为struct
func MapToStruct(mapValue interface{}, structValue interface{}) (err error) {
	err = mapstructure.Decode(mapValue, structValue)
	return err
}

// string list to list
func Eval(str string) (evalVal string) {
	evalVal = strings.Replace(str, "\\", "", -1)
	evalVal = strings.Replace(evalVal, "\"[", "[", -1)
	evalVal = strings.Replace(evalVal, "]\"", "]", -1)
	return evalVal
}

func TrimSpace(objJsonStr string) string {
	reg := regexp2.MustCompile(`"[\s\t\n]*|[\s\t\n]*"`, 0)
	replaceStr, _ := reg.Replace(objJsonStr, `"`, -1, -1)
	replaceStr = strings.Replace(replaceStr, `\\t`, "", -1)
	replaceStr = strings.Replace(replaceStr, `\t`, "", -1)
	//replaceStr = strings.Replace(replaceStr, "\\", "", -1)
	return replaceStr
}

// ReFormatMoney 格式化金额
// 格式化规则: 从右边开始按照每3个长度切割字符串并使用 `,` 拼接
// e.g: input:12000 -> output:12,000 ; input:1200000 -> output:1,200,000
func ReFormatMoney(str string) string {
	chunkSize := 3
	strLength := len(str)
	chunks := make([]string, 0)

	for i := strLength; i > 0; i -= chunkSize {
		startIndex := i - chunkSize
		if startIndex < 0 {
			startIndex = 0
		}
		chunk := str[startIndex:i]
		chunks = append([]string{chunk}, chunks...)
	}

	return strings.Join(chunks, ",")
}
