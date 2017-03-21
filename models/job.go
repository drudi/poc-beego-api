package models

import (
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// Job struct
type Job struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
}

// TableName for Job
func (p Job) TableName() string {
	return "jobs"
}

// AddJob create new job
func AddJob(u Job) Job {
	logs.Info("models inicio... %s", u)
	db := orm.NewOrm()
	id, err := db.Insert(&u)
	if err == nil {
		logs.Error(err)
	}
	logs.Info("models ... fim %s", id)
	return u
}

// GetJob return Job with id
func GetJob(id int) (*Job, error) {
	db := orm.NewOrm()
	job := &Job{ID: id}
	err := db.Read(job)
	if err == orm.ErrNoRows {
		return nil, errors.New("No result found")
	} else if err == orm.ErrMissPK {
		return nil, errors.New("No primary key found")
	} else {
		return job, nil
	}

}
