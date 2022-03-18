package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SocialSetting struct {
	ID             primitive.ObjectID  `bson:"_id" json:"_id"`
	Username       string              `json:"username" gorm:"unique" bson:"username,omitempty"`
	Email          string              `json:"email" bson:"email,omitempty"`
	AlignedCb      bool                `json:"alignedCb" bson:"aligned_cb,omitempty"`
	BillingPeriod  int                 `json:"billingPeriod" bson:"billing_period,omitempty"`
	Salary         int                 `json:"Salary" bson:"salary,omitempty"`
	LastModifiedAt primitive.Timestamp `json:"lastModifiedAt" bson:"last_modified_at,omitempty"`
	CurrentDevice  Device              `json:"currentDevice" bson:"current_device,omitempty"`
}

type Device struct {
	Name    string `json:"name" bson:"name,omitempty"`
	Version string `json:"version" bson:"version,omitempty"`
}

func NewSocialSetting(username string, email string, alignedCb bool, billingPeriod int, salary int, lastModifiedAt time.Time, currentDevice Device) SocialSetting {
	socialSetting := SocialSetting{}
	socialSetting.ID = primitive.NewObjectID()
	socialSetting.Username = username
	socialSetting.Email = email
	socialSetting.AlignedCb = alignedCb
	socialSetting.BillingPeriod = billingPeriod
	socialSetting.Salary = salary
	socialSetting.LastModifiedAt = primitive.Timestamp{T: uint32(lastModifiedAt.Unix())}
	socialSetting.CurrentDevice = currentDevice
	return socialSetting
}

func NewDevice(name string, version string) Device {
	device := Device{}
	device.Name = name
	device.Version = version
	return device
}
