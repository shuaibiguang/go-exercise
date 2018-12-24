package gredis

import (
	"gin-blog/pkg/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle: setting.RedisSetting.MaxIdle, // 最大空闲链接数
		MaxActive: setting.RedisSetting.MaxActive,  // 在给定时间内，允许分配的最大连接数 (当为0时， 没有限制)
		IdleTimeout: setting.RedisSetting.IdleTimeout,  // 在给定时间内将会保持空闲状态， 若到达时间限制则关闭链接（当为0时没有限制）
		Dial: func () (redis.Conn, error) { // 提供创建和配置应用程序链接的一个函数
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow:func (c redis.Conn, t time.Time) error {   // 可选的应用程序检查健康功能
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

func Set(key string, data interface{}, time int) (bool, error) {
	// 在连接池中获取一个活跃连接
	conn := RedisConn.Get()
	defer conn.Close()

	// conn.Do() 向redis服务器 发送命令并返回收到的答复
	// redis.Bool 将命令返回转为布尔值， 其他的同理
	value, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	reply, err := redis.Bool(conn.Do("SET", key, value))
	conn.Do("EXPIRE", key, time)

	return reply, err
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func Delete (key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*" + key + "*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}