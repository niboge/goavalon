package plugin

import (
	"sync"
	"time"
)

const KEY_LOGIN = "key_login_info"

type SessionManager struct {
	mLifeTime int64               // 存活时间
	mLock     sync.RWMutex        //互斥锁头
	mSessions map[string]*Session //session指针指向内容
}

// session 内容
type Session struct {
	lastTime time.Time              // 登录的时间已用来回收
	mValues  map[string]interface{} // session设置的具体值
}

var manager *SessionManager
var once sync.Once

func GetSessionManager(lifeTime int64) *SessionManager {
	once.Do(
		func() {
			manager = &SessionManager{
				mLifeTime: lifeTime,
				mSessions: make(map[string]*Session)}
		})
	// todo 定时回收
	go manager.GC()
	return manager

}

func (manager *SessionManager) Set(key string, value interface{}) {
	//加锁
	manager.mLock.Lock()
	defer manager.mLock.Unlock()
	session, ok := manager.mSessions[KEY_LOGIN]
	if ok {
		session.mValues[key] = value
	} else {
		session := &Session{lastTime: time.Now(), mValues: make(map[string]interface{})}
		session.mValues[key] = value
		manager.mSessions[KEY_LOGIN] = session
	}

}

func (manager *SessionManager) Get(key string) (interface{}, bool) {
	manager.mLock.RLock()
	defer manager.mLock.RUnlock()
	if session, ok := manager.mSessions[KEY_LOGIN]; ok {
		if val, ok := session.mValues[key]; ok {
			return val, ok
		}
	}
	return nil, false
}

func (manager *SessionManager) AuthUser(KEY_LOGIN string) bool {
	manager.mLock.RLock()
	defer manager.mLock.RUnlock()
	if _, ok := manager.mSessions[KEY_LOGIN]; ok {
		return true
	} else {
		return false
	}
}

func (manager *SessionManager) GC() {
	manager.mLock.Lock()
	defer manager.mLock.Unlock()
	for sessionId, session := range manager.mSessions {
		if session.lastTime.Unix()+manager.mLifeTime > time.Now().Unix() {
			delete(manager.mSessions, sessionId)
		}
	}
	// 定时任务回收
	time.AfterFunc(time.Duration(manager.mLifeTime)*time.Second, func() { manager.GC() })
}
