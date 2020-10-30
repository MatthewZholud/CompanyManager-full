package service

//import (
//	"go.uber.org/zap"
//)
//
type RedisRep interface {
	Set( msg int)  error
	Get() ([]int, error)
}
//
//type BotService struct {
//	apiConfig *config.APIConfig
//	store     store.Store
//}
//
//func NewService() Service {
//	return BotService{
//	}
//}