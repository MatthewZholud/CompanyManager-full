package redis

import (
	"encoding/json"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
)

const UsersIDs = "UsersIDs"
func (r *redisClient) Set(newId int) {
	users, isNew := r.needChange(newId)
	fmt.Println(newId)

	if isNew {
		users = append(users, newId)
		cacheEntry, err := json.Marshal(users)
		if err != nil {
			logger.Log.Errorf("Can't register new user: Can't marshal slice to send to cash: %v", err)
		}
		err = r.client.Set(UsersIDs, cacheEntry, 0).Err()
		if err != nil {
			logger.Log.Errorf("Can't register new user: Can't set slice to cash: %v", err)
		}
		logger.Log.Infof("New user has been registered: ID %v", newId)

	}
}

func (r *redisClient) Get() ([]int, error) {
	var users []int

	val, err := r.client.Get(UsersIDs).Result()
	if err != nil{
		logger.Log.Infof("Can't get users from the cash: %v", err)
	}

	err = json.Unmarshal([]byte(val), &users)
	if err != nil {
		logger.Log.Infof("Can't get users from the cash: %v", err)
	}

	return users, nil
}

func (r *redisClient) needChange(newId int) ([]int, bool) {
	users, err := r.Get()
	if err != nil{
		return users, false
	}
	for _, oldId := range users {
		if newId == oldId {
			return users, false
		}
	}
	return users, true
}
