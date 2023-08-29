package services

import (
	"context"
	"task-manager-plus/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	ctx   *context.Context
	tasks *mongo.Collection
}

func NewTaskService() TaskService {
	return TaskService{
		tasks: tasksCollection,
		ctx:   &ctx,
	}
}

// func taskToBSONM(userTask models.Task) (bson.M, error) {
// 	task, err := bson.Marshal(userTask)
// 	if err != err {
// 		return nil, err
// 	}
// 	var taskBSON bson.M
// 	err = bson.Unmarshal(task, &taskBSON)
// 	return taskBSON, err
// }

func (ts *TaskService) CreateTask(task models.Task) error {
	// taskBSON, err := taskToBSONM(task)
	// if err != nil {
	// 	return err
	// }
	// taskBSON = bson.M{""}
	_, err := ts.tasks.InsertOne(*ts.ctx, task)
	return err
}
