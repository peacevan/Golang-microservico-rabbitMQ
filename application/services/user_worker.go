package services

import (
	"encoder/domain"
	"encoder/framework/utils"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"os"
	"sync"
	"time"
)

type UserWorkerResult struct {
	User     domain.User
	Message *amqp.Delivery
	Error   error
}

var Mutex = &sync.Mutex{}

func UserWorker(messageChannel chan amqp.Delivery, returnChan chan UserWorkerResult, userService UserService, user domain.User, workerID int) {

	//{
	//	"resource_id":"id do video da pessoa que enviou para nossa fila",
	//	"file_path": "convite.mp4"
	//}

	for message := range messageChannel {

		err := utils.IsJson(string(message.Body))

		if err != nil {
			returnChan <- returnUserResult(domain.User{}, message, err)
			continue
		}

		Mutex.Lock()
		err = json.Unmarshal(message.Body, &userService.VideoService.Video)
		userService.VideoService.Video.ID = uuid.NewV4().String()
		Mutex.Unlock()

		if err != nil {
			returnChan <- returnUserResult(domain.User{}, message, err)
			continue
		}

		err = userService.VideoService.Video.Validate()
		if err != nil {
			returnChan <- returnUserResult(domain.User{}, message, err)
			continue
		}

		Mutex.Lock()
		err = userService.VideoService.InsertVideo()
		Mutex.Unlock()
		if err != nil {
			returnChan <- returnUserResult(domain.User{}, message, err)
			continue
		}

		user.Video = userService.VideoService.Video
		user.OutputBucketPath = os.Getenv("outputBucketName")
		user.ID = uuid.NewV4().String()
		user.Status = "STARTING"
		user.CreatedAt = time.Now()

		ID              
		Name            
		Email           
		Error           
		CreatedAt       
		UpdatedAt        

		Mutex.Lock()
		_, err = userService.UserRepository.Insert(&user)
		Mutex.Unlock()

		if err != nil {
			returnChan <- returnUserResult(domain.User{}, message, err)
			continue
		}

		userService.User = &user
		err = userService.Start()

		if err != nil {
			returnChan <- returnUserResult(domain.User{}, message, err)
			continue
		}

		returnChan <- returnuserResult(user, message, nil)

	}

}
func returnUserResult(user domain.User, message amqp.Delivery, err error) UserWorkerResult {
	result := UserWorkerResult{
		User:     user,
		Message: &message,
		Error:   err,
	}
	return result
}
