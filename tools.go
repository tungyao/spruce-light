package spruce

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func SplitString(str []byte, p []byte) [][]byte {
	group := make([][]byte, 0)
	ps := 0
	for i := 0; i < len(str); i++ {
		if str[i] == p[0] && i < len(str)-len(p) {
			if len(p) == 1 {
				group = append(group, str[ps:i])
				ps = i + len(p)
				// return [][]byte{str[:i], str[i+1:]}
			} else {
				for j := 1; j < len(p); j++ {
					if str[i+j] != p[j] || j != len(p)-1 {
						continue
					} else {
						group = append(group, str[ps:i])
						ps = i + len(p)
					}
					// return [][]byte{str[:i], str[i+len(p):]}
				}
			}
		} else {
			continue
		}
	}
	group = append(group, str[ps:])
	return group
}
func FindString(v []byte, p []byte) interface{} {
	// switch v.(type) {
	// case []byte:
	bt := v
	for i := 0; i < len(bt); i++ {
		ist := make([]int, len(p))
		for k, v := range p {
			if i < len(bt)-len(p) && bt[i+k] == v {
				ist[k] = 1
			}
		}
		st := true
		for _, v := range ist {
			if v != 1 {
				st = false
			}
		}
		if st {
			return bt[i+len(p):]
		}
	}
	return nil
	// case string:
	// 	sr := v.(string)
	// }
	// return nil
}
func Equal(one []byte, two []byte) bool {
	if len(one) != len(two) {
		return false
	}
	for k, v := range one {
		if v != two[k] {
			return false
		}
	}
	return true
}
func CreateUUID(length int, xtr []byte, self []byte) []byte {
	str := fmt.Sprintf("%x", xtr)
	strLow := ComplementHex(str[:(len(str)-1)/3], (length-20)*2/3)
	strMid := ComplementHex(str[(len(str)-1)/3:(len(str)-1)*2/3], (length-20)/3)
	<-time.After(1 * time.Nanosecond)
	ti := time.Now().UnixNano()
	return []byte(ComplementHex(fmt.Sprintf("%s%x%s%s", strLow, ti, strMid, self), length))
}
func ComplementHex(s string, x int) string {
	if len(s) == x {
		return s
	} else if len(s) < x {
		s += string(CreateNewId(x - len(s)))
	} else if len(s) > x {
		return s[:x]
	}
	return s
}
func CreateNewId(length int) []byte {
	d := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM012345689"
	da := make([]byte, length)
	for i := 0; i < length; i++ {
		<-time.After(time.Nanosecond * 1)
		da[i] = d[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(15)]
	}
	return da
}
func MD5(b []byte) []byte {
	m := md5.New()
	m.Write(b)
	return []byte(fmt.Sprintf("%x", m.Sum(nil)))
}
