package logic

import (
	"github.com/sudofrost/expense-tracker/internal/models"
)

var tracker *models.Tracker

func init() {
	tracker = models.NewTracker()
}
