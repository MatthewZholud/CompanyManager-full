package service


type RedisRep interface {
	Set( msg int)
	Get() ([]int, error)
}
