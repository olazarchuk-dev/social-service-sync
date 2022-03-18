package ws

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"social-service-sync/server/app/helper"
	"social-service-sync/server/model/entity"
)

func CreateSocialSetting(m *Mongo, u entity.SocialSetting) (string, error) {
	result, err := m.Collection.InsertOne(m.Ctx, u)
	if err != nil {
		return "0", err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	return fmt.Sprintf("%v", oid.Hex()), err
}

func UpdateSocialSetting(m *Mongo, id string, username string, email string, alignedCb bool, billingPeriod int, salary int, currentDevice Device, lastModifiedAt int) error {
	objectId, errFound := primitive.ObjectIDFromHex(id)
	helper.PanicErr(errFound)

	filter := bson.D{{"_id", objectId}}
	update := bson.D{
		{"$set", bson.D{{"username", username}}},
		{"$set", bson.D{{"email", email}}},
		{"$set", bson.D{{"aligned_cb", alignedCb}}},
		{"$set", bson.D{{"billing_period", billingPeriod}}},
		{"$set", bson.D{{"salary", salary}}},
		{"$set", bson.D{{"current_device", currentDevice}}},
		{"$set", bson.D{{"last_modified_at", primitive.Timestamp{T: uint32(lastModifiedAt)}}}},
	}
	_, err := m.Collection.UpdateOne(m.Ctx, filter, update)
	return err
}
