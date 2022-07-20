package models

type Activity struct {
	Id          int `gorm:"primaryKey;autoIncrement:true;unique"`
	UserId      int `gorm:"not nul"`
	Activity    string
	Description string
	Date        string
	Time        string
	Location    string
	Category    string
	Priority    string
}

func NewActivity(userId interface{}, activity, description, date, time, location, category, priority string) (*Activity, error) {
	var Activity Activity
	Activity.UserId = userId.(int)
	Activity.Activity = activity
	Activity.Description = description
	Activity.Date = date
	Activity.Time = time
	Activity.Location = location
	Activity.Category = category
	Activity.Priority = priority
	err := DB.Create(&Activity).Error
	if err != nil {
		return nil, err
	}
	return &Activity, err
}

// AllActivity returns all activities
func AllActivityById(userId int) ([]Activity, error) {
	var activities []Activity
	err := DB.Where("user_id = ?", userId).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

// ActivityById returns activity by id
func ActivityByUserId(userId int) (*Activity, error) {
	var activity Activity
	err := DB.Where("user_id = ?", userId).First(&activity).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}