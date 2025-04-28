package util 


import (
	"testing"

	"reflect"
)
const env_key = "TEST_ENV_KEY"

func TestEnv(t *testing.T) {
	env_value := "123123123123"
	Print("環境変数取得テスト->",env_value)

	SetEnv(env_key,env_value)

	result := GetEnvOrDefault(env_key,"defautl value")

	Print("環境変数取得",result,GetEnvOrDefault("MOCK_ENV","defalut_value"))
}

func TestEnvGet(t *testing.T) {
	Print("環境変数の数値化テスト")

	result := GetEnvOrDefaultToInt("notkey","123")
	Print(result)
	Print(reflect.TypeOf(result))
}