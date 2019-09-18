package main

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/weixuan75/go_model/util"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Obj struct {
	Id         string
	CreateTime *time.Time
}

type User struct {
	Obj
	UserName *string     `json:"user_name" objCheck:"type:string| 字符串;length:0,10| 长度;notNull:1 | 不能为空"`
	Name     string      `json:"name" objCheck:"type:string| 字符串;length:0,10| 长度;notNull:1 | 不能为空"`
	ToAje    *int        `json:"to_aje" objCheck:"type:int| 整数;length:0,10| 长度;notNull:1 | 不能为空"`
	ToCheng  *bool       `json:"to_cheng" checkType:"bool" checkLength:"msg:" checkNull:"msg:"`
	ToPrice  *float64    `json:"to_price" objCheck:"type:float| 浮点数;length:0,10| 长度;notNull:1 | 不能为空"`
	List     []string    `json:"list" objCheck:"type:array| 数组;length:0,10| 长度"`
	ToObje   interface{} `json:"to_obje" checkType:"object" checkLength:"msg:" checkNull:"msg:"`
	Regex    string      `json:"regex" objCheck:"type:regex|正则表达式;regex:^(\\d{15,15}|\\d{16,16}|\\d{17,17}|\\d{18,18}|\\d{19,19}|(\\d{17,17}[x|X]))$"`
}

//checkType 类型：string,int,bool,float,array,object,email,phone,idCard
//checkLength 类型长度：(int float)(0,10).最小，最大 array，string（最短，最长）
func main() {

	fmt.Println(len("type:"))
	fmt.Println(len("length:"))
	fmt.Println(len("null:"))
	var u = new(User)
	u.Name = "杜英杰"
	u.UserName = &u.Name
	f := float64(0.3)
	u.ToPrice = &f
	aje := 10
	u.ToAje = &aje
	t := reflect.TypeOf(*u) // 获取对象的属性类型，获取类型定义里面的所有元素
	fmt.Println(t.Name())
	v := reflect.ValueOf(*u) // 得到实际的值，获取对象存储在类型中的值，还可以去改变值
	if v.IsValid() {
		fmt.Println("不等于空")
	} else {
		fmt.Println("等于空")
	}
	// t.NumField() 计算对象有多少属性
	// t.NumMethod() 返回类型的方法集中导出的方法数
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		val := v.Field(i)
		//switch field.Type.Kind() {
		//case reflect.String:
		//case reflect.Bool:
		//case reflect.Int:
		//case reflect.Int64:
		//case reflect.Float64:
		//	//fmt.Println(val)
		//	continue
		//case reflect.Interface:
		//case reflect.Ptr:
		//	fmt.Println(val.Elem())
		//	continue
		//default:
		//	fmt.Println(val)
		//	continue
		//}
		o := ObjCheck{}
		o.StructTag(field)
		if o.Type == "regex" {

			//判断类型
			switch field.Type.Kind() {
			case reflect.String:
				length := len(val.String())
				if o.Length && (o.Max < int64(length) && o.Min > int64(length)) {
					if len(o.Msg["notNull"]) > 0 {
						errors.New(o.Msg["notNull"])
						panic(o.Msg["notNull"])
					} else {
						errors.New(fieldName + "不能为空")
						panic(fieldName + "不能为空")
					}
				}
				continue
			case reflect.Bool:
				continue
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if o.Length && (o.Max < val.Int() && o.Min > val.Int()) {
					if len(o.Msg["notNull"]) > 0 {
						errors.New(o.Msg["notNull"])
						panic(o.Msg["notNull"])
					} else {
						msg := fieldName + "范围(" + strconv.FormatInt(o.Min, 10) + "," + strconv.FormatInt(o.Max, 10) + ")"
						errors.New(msg)
						panic(msg)
					}
				}
				continue
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				if o.Length && (o.Max < int64(val.Uint()) && o.Min > int64(val.Uint())) {
					if len(o.Msg["notNull"]) > 0 {
						errors.New(o.Msg["notNull"])
						panic(o.Msg["notNull"])
					} else {
						msg := fieldName + "范围(" + strconv.FormatInt(o.Min, 10) + "," + strconv.FormatInt(o.Max, 10) + ")"
						errors.New(msg)
						panic(msg)
					}
				}
				continue
			case reflect.Float32, reflect.Float64:
				if o.Length && (o.Max < int64(val.Float()) && o.Min > int64(val.Uint())) {
					if len(o.Msg["notNull"]) > 0 {
						errors.New(o.Msg["notNull"])
						panic(o.Msg["notNull"])
					} else {
						msg := fieldName + "范围(" + strconv.FormatInt(o.Min, 10) + "," + strconv.FormatInt(o.Max, 10) + ")"
						errors.New(msg)
						panic(msg)
					}
				}
				continue
			case reflect.Interface, reflect.Ptr:
				continue
			}
		} else if o.NotNull && IsBlankValue(val) {
			if len(o.Msg["notNull"]) > 0 {
				errors.New(o.Msg["notNull"])
				panic(o.Msg["notNull"])
			} else {
				errors.New(fieldName + "不能为空")
				panic(fieldName + "不能为空")
			}
		} else {
			//判断类型
			switch field.Type.Kind() {
			case reflect.String:
				length := len(val.String())
				if o.Length && (o.Max < int64(length) && o.Min > int64(length)) {
					if len(o.Msg["notNull"]) > 0 {
						errors.New(o.Msg["notNull"])
						panic(o.Msg["notNull"])
					} else {
						errors.New(fieldName + "不能为空")
						panic(fieldName + "不能为空")
					}
				}
				continue
			case reflect.Bool:
				continue
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if o.Length && (o.Max < val.Int() && o.Min > val.Int()) {
					if len(o.Msg["notNull"]) > 0 {
						errors.New(o.Msg["notNull"])
						panic(o.Msg["notNull"])
					} else {
						msg := fieldName + "范围(" + strconv.FormatInt(o.Min, 10) + "," + strconv.FormatInt(o.Max, 10) + ")"
						errors.New(msg)
						panic(msg)
					}
				}
				continue
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				if o.Length && (o.Max < int64(val.Uint()) && o.Min > int64(val.Uint())) {
					if len(o.Msg["notNull"]) > 0 {
						errors.New(o.Msg["notNull"])
						panic(o.Msg["notNull"])
					} else {
						msg := fieldName + "范围(" + strconv.FormatInt(o.Min, 10) + "," + strconv.FormatInt(o.Max, 10) + ")"
						errors.New(msg)
						panic(msg)
					}
				}
				continue
			case reflect.Float32, reflect.Float64:
				if o.Length && (o.Max < int64(val.Float()) && o.Min > int64(val.Uint())) {
					if len(o.Msg["notNull"]) > 0 {
						errors.New(o.Msg["notNull"])
						panic(o.Msg["notNull"])
					} else {
						msg := fieldName + "范围(" + strconv.FormatInt(o.Min, 10) + "," + strconv.FormatInt(o.Max, 10) + ")"
						errors.New(msg)
						panic(msg)
					}
				}
				continue
			case reflect.Interface, reflect.Ptr:
				continue
			}
		}
		fmt.Println(util.JsonToStr(o))
	}
}

//获取对象属性
func get(obj interface{}) {
	// 获取对象tag
	t := reflect.TypeOf(obj)
	objType := t.Elem()
	for i := 0; i < t.NumField(); i++ {
		field := objType.Field(i)
		structTag := field.Tag
		fmt.Println(structTag.Get("json"))
	}
	//fmt.Println(t.Len())
	//fmt.Println(t.Bits())
	// 获取对象属性
	value := reflect.ValueOf(obj)
	typ := value.Type()
	for i := 0; i < value.NumMethod(); i++ {
		fmt.Println(fmt.Sprintf("method[%d]%s and type is %v", i, typ.Method(i).Name, typ.Method(i).Type))
	}
}

type ObjCheck struct {
	Type     string            // 校验类型
	Length   bool              // 开启大小长度校验
	Min      int64             // 最小长度
	Max      int64             // 最大长度
	MinFloat float64           // 最小长度
	MaxFloat float64           // 最大长度
	NotNull  bool              // 是否为空
	Regex    string            // 正则表达式
	Msg      map[string]string // 信息提示
}

func (o *ObjCheck) StructTag(field reflect.StructField) {
	tag := field.Tag
	tagStr := tag.Get("objCheck")
	tagStr = strings.ReplaceAll(tagStr, " ", "")
	// 判断是否是正则表达式
	msg := make(map[string]string, 1)
	if strings.Index(tagStr, "type:regex") == 0 {
		o.Type = "regex"
		msg["msg"] = ""
		o.Regex = tagStr[strings.Index(tagStr, "regex:")+len("regex:"):]
	} else {
		tagArr := strings.Split(tagStr, ";")
		msg = make(map[string]string, len(tagArr))
		for _, a := range tagArr {
			if strings.Index(a, "type:") == 0 {
				k, v := splitTwo(a[5:], "|")
				msg["type"] = *v
				o.Type = *k
			} else if strings.Index(a, "length:") == 0 {
				k, v := splitTwo(a[5:], "|")
				if k != nil {
					if strings.Index(*k, ",") > 0 {
						min, max := splitTwo(*k, ",")
						min1, _ := strconv.ParseFloat(*min, 0)
						max1, _ := strconv.ParseFloat(*max, 0)
						o.MinFloat = min1
						o.MaxFloat = max1
						min2, _ := strconv.ParseInt(*min, 10, 64)
						max2, _ := strconv.ParseInt(*max, 10, 64)
						o.Min = min2
						o.Max = max2
						o.Length = true
					} else {
						max, _ := strconv.ParseInt(*k, 10, 64)
						o.Max = max
						max1, _ := strconv.ParseFloat(*k, 0)
						o.MaxFloat = max1
						o.Length = true
					}
				}
				msg["length"] = *v

			} else if strings.Index(a, "notNull:") == 0 {
				k, v := splitTwo(a[8:], "|")
				if k != nil && *k == "1" {
					o.NotNull = true
				}
				msg["notNull"] = *v
			}
		}
	}
	o.Msg = msg
}
func splitTwo(str, sep string) (k, v *string) {
	if strings.Index(str, sep) > 0 {
		a := strings.Split(str, sep)
		k = &a[0]
		v = &a[1]
	} else {
		k = &str
	}
	return
}

func IsBlankValue(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		values := strings.Replace(value.String(), " ", "", -1)
		// 去除换行符
		values = strings.Replace(values, "\n", "", -1)
		values = strings.Replace(values, "\t", "", -1)
		return len(values) == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
func IsBlank(model interface{}) bool {
	return IsBlankValue(reflect.ValueOf(model))
}
