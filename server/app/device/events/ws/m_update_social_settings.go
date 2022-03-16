package ws

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"social-service-sync/server/app/helper"
	"social-service-sync/server/model/entity"
)

func CreateSocialSetting(u entity.SocialSetting) (string, error) {
	result, err := SocialSettingsCollection.InsertOne(Ctx, u)
	if err != nil {
		return "0", err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	return fmt.Sprintf("%v", oid.Hex()), err
}

//func UpdateSocialSetting(id primitive.ObjectID, username string, email string, alignedCb bool, billingPeriod int, salary int, currentDevice entity.Device) error {
func UpdateSocialSetting(id string, username string, email string, alignedCb bool, billingPeriod int, salary int) error {
	objectId, errFound := primitive.ObjectIDFromHex(id)
	helper.PanicErr(errFound)

	//filter := bson.D{{"_id", id}}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{
		{"$set", bson.D{{"username", username}}},
		{"$set", bson.D{{"email", email}}},
		{"$set", bson.D{{"aligned_cb", alignedCb}}},
		{"$set", bson.D{{"billing_period", billingPeriod}}},
		{"$set", bson.D{{"salary", salary}}},
		//{"$set", bson.D{{"current_device", currentDevice}}},
	}
	_, err := SocialSettingsCollection.UpdateOne(Ctx, filter, update)
	return err
}
