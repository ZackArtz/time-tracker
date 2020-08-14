package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Timestamp struct {
	*gorm.Model
	UUID      string    `json:"uuid"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Comment   string    `json:"comment"`
	Project   string    `json:"project"`
	Category  string    `json:"category"`
}

func (t *Timestamp) Prepare() Timestamp {
	guid := uuid.New()
	t.ID = 0
	t.Comment = html.EscapeString(strings.TrimSpace(t.Comment))
	t.Project = html.EscapeString(strings.TrimSpace(t.Project))
	t.UUID = guid.String()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return *t
}

func (t *Timestamp) Create(db *gorm.DB) (*Timestamp, error) {
	var err error
	err = db.Debug().Create(&t).Error
	if err != nil {
		return &Timestamp{}, err
	}
	return t, nil
}

func (t *Timestamp) GetAllTimestamps(db *gorm.DB) (*[]Timestamp, error) {
	var err error
	var ts []Timestamp
	err = db.Debug().Model(&Timestamp{}).Limit(100).Find(&ts).Error
	return &ts, err
}

func (t *Timestamp) GetAllTimestampsByProject(db *gorm.DB, project string) (*[]Timestamp, error) {
	var err error
	var ts []Timestamp
	err = db.Debug().Model(&Timestamp{}).Where("project = ?", project).Limit(100).Find(&ts).Error
	return &ts, err
}

func (t *Timestamp) GetTimestampByID(db *gorm.DB, uuid string) (*Timestamp, error) {
	err := db.Debug().Model(&Timestamp{}).Where("uuid = ?", uuid).Take(&t).Error
	if err != nil {
		return &Timestamp{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Timestamp{}, errors.New("timestamp not found")
	}
	return t, err
}

func (t *Timestamp) Delete(db *gorm.DB, uuid string) (int64, error) {
	db = db.Debug().Model(&Timestamp{}).Where("uuid = ?", uuid).Take(&Timestamp{}).Delete(&Timestamp{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (t *Timestamp) Validate() error {
	if t.Project == "" {
		return errors.New("project is required")
	}
	return nil
}
