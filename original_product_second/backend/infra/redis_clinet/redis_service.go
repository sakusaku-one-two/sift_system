package redis_client

import (
	"backend/application/dto"
	"fmt"
	"github.com/redis/go-redis/v9"
	"reflect"
	"sync"
)

var (
	REDIS_CLIENT *redis.Client //共有クライアント
)

const (
	ACTIVE_USERS_REDIS_KEY = uitl.GetEnvOrDefault("ACTIVE_USERS_REDIS_KEY", "ACTIVE_USERS")
)

//----------------------------[reidsで共有＆配信する構造体を包む構造体（RedisServiceのActiveChannelsで利用）]--------------------------------

type PubSubModels[Model any] struct {
	Value Model
	Type  reflect.Type
}

func NewPubSubModels[Model any](value Model) PubSubModels {
	type_ := reflect.TypeOf(Model)

	model := PubSubModels[Model]{
		Value: value,
		Type:  type_,
	}

	return model
}

func ModleFromModle[Model any](json []byte) Model {
	value, _ := dto.JsonToModle[Modle](json)
	return NewPubSubModels[Model](*value)
}

func (pbm *PubSubModels[Model]) ToJson() []byte {
	json, _ := dto.ModelToJson[Model](pbm.Value)
	return json
}

//----------------------------[]--------------------------------

type RedisServiece struct {
	RDB                *redis.Client
	active_channels_mu sync.Mutex
	ActiveChannels     map[string][]PubSubModels[any]
	active_usres_mu    sync.Mutex
	ActiveUsers        map[string][]PubSubModels[any]
	is_init            bool //初期化のフラグ
}

func NewRedisService() (*RedisServiece, error) {

	if REDIS_CLIENT == nil {
		REDIS_CLIENT = GetCLient()
	}

	redis_service := &RedisService{
		RDB:                REDIS_CLIENT,
		ActiveChannels:     make(map[string][]PubSubModels[any]),
		ActiveUsers:        make(map[string][]PubSubModels[any]),
		active_channels_mu: sync.Mutext{},
		active_usres_mu:    sync.Mutext{},
	}

	if err := redis_service.Init(); err != nil {
		//初期化失敗
		return nil, err
	}

	return redis_service, nil
}

func (rs *RedisServiece) Init() error {
	//ECSで立ち上がった際にアクティブなユーザーIDや認証情報、PUBSUBのチャネルを受け取る
	rs.is_init = True
	return nil
}
