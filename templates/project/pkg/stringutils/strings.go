/**
 * Created by Wangwei on 2019-05-30 20:04.
 */

package stringutils

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GenPayCode(prefix string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := prefix
	for i := len(prefix); i < 18; i++ {
		code = code + strconv.Itoa(r.Intn(10))
	}

	return code
}

func GenNo() string {
	_uuid := uuid.NewV4().String()
	_uuid = strings.Replace(_uuid, "-", "", -1)
	no := Substr(_uuid, 0, 32)
	return no
}

/**
生成订单号
*/
func GetOrderNo(l int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prefix := "EP" + time.Now().Format("20060102150405") //16
	code := prefix
	for i := len(prefix); i < l; i++ {
		code = code + strconv.Itoa(r.Intn(10))
	}
	return code
}

func GetRandomString(l int, prefix string) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return prefix + string(result)
}

func GetAuthtoken(l int, prefix string) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return prefix + string(result)
}

func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}

	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

//RandomStr 获取一个随机字符串
func RandomStr() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// LocalIP 获取机器的IP
func LocalIP() string {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return ""
}
