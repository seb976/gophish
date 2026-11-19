package models

import (
	"errors"
	"time"

	log "github.com/gophish/gophish/logger"
)

// TrainingPage contains the fields used for a training / awareness Page model.
// A training page is a standalone HTML page (just like a landing page) that can
// be shown to a target for security-awareness purposes, typically after they
// have interacted with a landing page.
type TrainingPage struct {
	Id           int64     `json:"id" gorm:"column:id; primary_key:yes"`
	UserId       int64     `json:"-" gorm:"column:user_id"`
	Name         string    `json:"name"`
	HTML         string    `json:"html" gorm:"column:html"`
	ModifiedDate time.Time `json:"modified_date"`
}

// TableName overrides the gorm table name so it maps to training_pages.
func (TrainingPage) TableName() string {
	return "training_pages"
}

// ErrTrainingPageNameNotSpecified is thrown if the name of the training page is blank.
var ErrTrainingPageNameNotSpecified = errors.New("Training Page Name not specified")

// Validate ensures that a training page contains the appropriate details.
func (p *TrainingPage) Validate() error {
	if p.Name == "" {
		return ErrTrainingPageNameNotSpecified
	}
	if err := ValidateTemplate(p.HTML); err != nil {
		return err
	}
	return nil
}

// GetTrainingPages returns the training pages owned by the given user.
func GetTrainingPages(uid int64) ([]TrainingPage, error) {
	ps := []TrainingPage{}
	err := db.Where("user_id=?", uid).Find(&ps).Error
	if err != nil {
		log.Error(err)
		return ps, err
	}
	return ps, err
}

// GetTrainingPage returns the training page, if it exists, specified by the given id and user_id.
func GetTrainingPage(id int64, uid int64) (TrainingPage, error) {
	p := TrainingPage{}
	err := db.Where("user_id=? and id=?", uid, id).Find(&p).Error
	if err != nil {
		log.Error(err)
	}
	return p, err
}

// GetTrainingPageByName returns the training page, if it exists, specified by the given name and user_id.
func GetTrainingPageByName(n string, uid int64) (TrainingPage, error) {
	p := TrainingPage{}
	err := db.Where("user_id=? and name=?", uid, n).Find(&p).Error
	if err != nil {
		log.Error(err)
	}
	return p, err
}

// PostTrainingPage creates a new training page in the database.
func PostTrainingPage(p *TrainingPage) error {
	err := p.Validate()
	if err != nil {
		log.Error(err)
		return err
	}
	// Insert into the DB
	err = db.Save(p).Error
	if err != nil {
		log.Error(err)
	}
	return err
}

// PutTrainingPage edits an existing TrainingPage in the database.
// Per the PUT Method RFC, it presumes all data for a page is provided.
func PutTrainingPage(p *TrainingPage) error {
	err := p.Validate()
	if err != nil {
		return err
	}
	err = db.Where("id=?", p.Id).Save(p).Error
	if err != nil {
		log.Error(err)
	}
	return err
}

// DeleteTrainingPage deletes an existing training page in the database.
// An error is returned if a page with the given user id and page id is not found.
func DeleteTrainingPage(id int64, uid int64) error {
	err := db.Where("user_id=?", uid).Delete(TrainingPage{Id: id}).Error
	if err != nil {
		log.Error(err)
	}
	return err
}
