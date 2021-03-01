package server

import (
	getty "github.com/apache/dubbo-getty"
	"time"
)

import (
	getty2 "github.com/duolacloud/seata-golang/pkg/base/getty"
	"github.com/duolacloud/seata-golang/pkg/base/protocal"
	"github.com/duolacloud/seata-golang/pkg/util/log"
)

const (
	CronPeriod = 20e9
)

func (coordinator *DefaultCoordinator) OnOpen(session getty.Session) error {
	log.Infof("got getty_session:%s", session.Stat())
	return nil
}

func (coordinator *DefaultCoordinator) OnError(session getty.Session, err error) {
	SessionManager.ReleaseGettySession(session)
	session.Close()
	log.Errorf("getty_session{%s} got error{%v}, will be closed.", session.Stat(), err)
}

func (coordinator *DefaultCoordinator) OnClose(session getty.Session) {
	log.Info("getty_session{%s} is closing......", session.Stat())
}

func (coordinator *DefaultCoordinator) OnMessage(session getty.Session, pkg interface{}) {
	log.Debugf("received message:{%v}", pkg)
	rpcMessage, ok := pkg.(protocal.RpcMessage)
	if ok {
		_, isRegTM := rpcMessage.Body.(protocal.RegisterTMRequest)
		if isRegTM {
			coordinator.OnRegTmMessage(rpcMessage, session)
			return
		}

		heartBeat, isHeartBeat := rpcMessage.Body.(protocal.HeartBeatMessage)
		if isHeartBeat && heartBeat == protocal.HeartBeatMessagePing {
			coordinator.OnCheckMessage(rpcMessage, session)
			return
		}

		if rpcMessage.MessageType == protocal.MSGTYPE_RESQUEST ||
			rpcMessage.MessageType == protocal.MSGTYPE_RESQUEST_ONEWAY {
			log.Debugf("msgId:%s, body:%v", rpcMessage.Id, rpcMessage.Body)
			_, isRegRM := rpcMessage.Body.(protocal.RegisterRMRequest)
			if isRegRM {
				coordinator.OnRegRmMessage(rpcMessage, session)
			} else {
				if SessionManager.IsRegistered(session) {
					defer func() {
						if err := recover(); err != nil {
							log.Errorf("Catch Exception while do RPC, request: %v,err: %w", rpcMessage, err)
						}
					}()
					coordinator.OnTrxMessage(rpcMessage, session)
				} else {
					session.Close()
					log.Infof("close a unhandled connection! [%v]", session)
				}
			}
		} else {
			resp, loaded := coordinator.futures.Load(rpcMessage.Id)
			if loaded {
				response := resp.(*getty2.MessageFuture)
				response.Response = rpcMessage.Body
				response.Done <- true
				coordinator.futures.Delete(rpcMessage.Id)
			}
		}
	}
}

func (coordinator *DefaultCoordinator) OnCron(session getty.Session) {
	active := session.GetActive()
	if CronPeriod < time.Since(active).Nanoseconds() {
		log.Infof("OnCorn session{%s} timeout{%s}", session.Stat(), time.Since(active).String())
		session.Close()
	}
}
