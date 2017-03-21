package controllers

import (
	"encoding/json"
	"poc-beego-api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// JobController base struct for Beego
type JobController struct {
	beego.Controller
}

// GetAll return all objects
// @Title GetAll
// @Description get all Jobs
// @Success 200 {object} models.Job
// @router / [get]
func (u *JobController) GetAll() {
	db := orm.NewOrm()
	var jobs []models.Job
	qs := db.QueryTable("jobs")

	qs.All(&jobs)
	u.Data["json"] = &jobs
	u.ServeJSON()
}

// Get Job
// @Title Get
// @Description get Job by id
// @Param	id		path 	int	true		"The key for job"
// @Success 200 {object} models.Job
// @Failure 403 :id not found
// @router /:id [get]
func (u *JobController) Get() {
	id, err := u.GetInt(":id")
	if err == nil {
		job, err := models.GetJob(id)
		if err != nil {
			u.CustomAbort(403, err.Error())
		} else {
			u.Data["json"] = job
		}
	}
	u.ServeJSON()
}

// Post a job
// @Title CreateJob
// @Description create job
// @Param	body		body 	models.Job	true		"body for job content"
// @Success 200 models.Job
// @Failure 403 body is empty
// @router / [post]
func (u *JobController) Post() {
	logs.Info("controller inicio...")
	var job models.Job
	json.Unmarshal(u.Ctx.Input.RequestBody, &job)
	logs.Info("controller antes... %s ", job)
	job = models.AddJob(job)
	logs.Info("controller depois... %s ", job)
	u.Data["json"] = &job
	u.ServeJSON()
	logs.Info("controller ...fim")
}
