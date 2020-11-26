package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var (
	_CONFIGS map[string]interface{}
	_RUNONCE sync.Once
)

func init() {
	_CONFIGS = make(map[string]interface{})
}

// 初始化所有注册的配置struct。
func _loadStructConfig() {
	for key, ptrStruct := range _CONFIGS {
		err := viper.UnmarshalKey(key, ptrStruct)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed load config %s: %v\n", key, err)
			os.Exit(1)
		}
		// 处理env
		ptrVal := reflect.ValueOf(ptrStruct)
		if ptrVal.Kind() != reflect.Ptr {
			continue
		}

		val := reflect.Indirect(ptrVal)
		if val.Kind() != reflect.Struct {
			continue
		}

		for i := 0; i < val.NumField(); i++ {
			name := key + "." + val.Type().Field(i).Name
			name = strings.ToUpper(strings.ReplaceAll(name, ".", "_"))
			envVal, ok := os.LookupEnv(name)
			if !ok {
				continue
			}

			fieldVal := reflect.ValueOf(ptrStruct).Elem().FieldByName(val.Type().Field(i).Name)
			if !fieldVal.CanSet() {
				continue
			}

			switch val.Field(i).Kind() {
			case reflect.String:
				fieldVal.SetString(cast.ToString(envVal))
			case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
				fieldVal.SetInt(cast.ToInt64(envVal))
			case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
				fieldVal.SetUint(cast.ToUint64(envVal))
			case reflect.Float64, reflect.Float32:
				fieldVal.SetFloat(cast.ToFloat64(envVal))
			case reflect.Bool:
				fieldVal.SetBool(cast.ToBool(envVal))
			case reflect.Slice:
				fieldVal.Set(reflect.ValueOf(strings.Fields(envVal)))
			default:
				continue
			}
		}
	}
}

// 注册需要解析为struct的配置项。
// * key: 配置项路径。
// * ptrStruct: 配置struct的引用。
func Register(key string, ptrStruct interface{}) {
	ptrVal := reflect.ValueOf(ptrStruct)
	if ptrVal.Kind() == reflect.Ptr {
		val := reflect.Indirect(ptrVal)
		if val.Kind() == reflect.Struct {
			_CONFIGS[key] = ptrStruct
			return
		}
	}

	fmt.Fprintf(os.Stderr, "config.Register need pointer of struct.\n")
	os.Exit(1)
}
