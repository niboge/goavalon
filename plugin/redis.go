// Redis 基础操作封装
package plugin

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	red "github.com/gomodule/redigo/redis"

	"errors"
	"log"
)

type Redis struct {
	zone     string
	host     string
	pass     string
	idx      int
	maxIdle  int
	maxConn  int
	idleTime int
	keySpace string
	pool     *red.Pool
}

// 执行一条Redis命令，只操作一个key
func (this *Redis) Exec(cmd string, key string, args ...interface{}) (interface{}, error) {
	var params []interface{}
	conn, connErr := this.Connect()

	if connErr != nil {
		return nil, connErr
	}

	defer conn.Close()

	params = append(params, this.GetKey(key))

	if len(args) > 0 {
		for _, arg := range args {
			params = append(params, arg)
		}
	}

	ret, execErr := conn.Do(cmd, params...)
	if execErr != nil {
		log.Fatalln(execErr, "geego-redis:exec-fail: "+cmd, params)
		return nil, execErr
	}

	log.Println("geego-redis:exec: ", cmd, params)

	return ret, nil
}

func (this *Redis) Get(key string) (interface{}, error) {
	return this.Exec("get", key)
}

func (this *Redis) Set(key string, value interface{}) (interface{}, error) {
	ex := "3600"
	return this.Exec("set", key, value, "EX", ex)
}

// 执行一条Redis命令，同时操作多个key，如：MGET
func (this *Redis) ExecWithKeys(cmd string, keys []string) ([]interface{}, error) {
	var params []interface{}

	conn, connErr := this.Connect()

	if connErr != nil {
		return nil, connErr
	}

	defer conn.Close()

	for _, key := range keys {
		params = append(params, this.GetKey(key))
	}

	ret, execErr := red.Values(conn.Do(cmd, params...))

	log.Println("geego-redis: exec with keys", cmd, keys, params)

	if execErr != nil {
		log.Fatalln(execErr, "geego-redis: exec with key fail: "+cmd)
		return nil, execErr
	}

	log.Println("geego-redis: exec with keys", cmd, keys, params, ret)

	return ret, nil
}

// 获取包含keyspace之后的key
func (this *Redis) GetKey(key string) string {
	if this.keySpace != "" {
		key = fmt.Sprintf("%s:%s", this.keySpace, key)
	}

	return key
}

// 获取一个redis连接，有连接池逻辑
func (this *Redis) Connect() (red.Conn, error) {

	if this.pool == nil {
		this.pool = &red.Pool{
			MaxIdle:     this.maxIdle,
			IdleTimeout: time.Duration(this.idleTime) * time.Second,
			Dial:        this.dial,
		}
	}

	conn := this.pool.Get()

	if err := conn.Err(); err != nil {
		log.Fatalln(err, "geego-redis: get connection fail")
		return nil, err
	}

	return conn, nil
}

// 创建一个redis连接
func (this *Redis) dial() (red.Conn, error) {
	conn, err := red.Dial("tcp", this.host)

	log.Println("geego-redis: connect info", this.host, this.keySpace, this.idx)

	if err != nil {
		log.Fatalln(err, "geego-redis: connect fail")
		return nil, err
	}

	if this.pass != "" {
		if _, err := conn.Do("AUTH", this.pass); err != nil {
			conn.Close()
			log.Fatalln(err, "geego-redis: auth fail")
			return nil, err
		}
	}

	if _, err := conn.Do("SELECT", this.idx); err != nil {
		conn.Close()
		log.Fatalln(err, "geego-redis: select fail: ", this.idx)
		return nil, err
	}

	return conn, nil
}

// 获取一个redis对象
func NewRedis(zone string) (*Redis, error) {
	if zone == "" {
		zone = "Redis"
	}

	maxIdle, _ := beego.AppConfig.Int(zone + "MaxIdle")
	maxConn, _ := beego.AppConfig.Int(zone + "MaxConn")
	idx, _ := beego.AppConfig.Int(zone + "Idx")
	idleTime, _ := beego.AppConfig.Int(zone + "IdleTime")

	redis := &Redis{
		zone:     zone,
		host:     beego.AppConfig.String(zone + "Host"),
		pass:     beego.AppConfig.String(zone + "Pass"),
		keySpace: beego.AppConfig.String(zone + "KeySpace"),
		idx:      idx,
		idleTime: idleTime,
		maxIdle:  maxIdle,
		maxConn:  maxConn,
	}

	if redis.host == "" {
		log.Panicln("invalid redis config")
		return nil, errors.New("empty host!")
	}

	return redis, nil
}
