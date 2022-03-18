package ws

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"social-service-sync/server/app/helper"
	"social-service-sync/server/model/entity"
)

func RepositoryCreate(m *Mongo, setting entity.SocialSetting) (string, error) {
	result, err := m.Collection.InsertOne(m.Ctx, setting)
	if err != nil {
		return "0", err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	return fmt.Sprintf("%v", oid.Hex()), err
}

func RepositoryUpdate(m *Mongo, id string, something *Something) error {
	objectId, errFound := primitive.ObjectIDFromHex(id)
	helper.PanicErr(errFound)

	filter := bson.D{{"_id", objectId}}
	update := bson.D{
		{"$set", bson.D{{"username", something.AppUsername}}},
		{"$set", bson.D{{"email", something.AppEmailAddress}}},
		{"$set", bson.D{{"aligned_cb", something.AppAlignedCb}}},
		{"$set", bson.D{{"billing_period", something.AppBillingPeriod}}},
		{"$set", bson.D{{"salary", something.AppSalary}}},
		{"$set", bson.D{{"current_device", something.CurrentDevice}}},
		{"$set", bson.D{{"last_modified_at", primitive.Timestamp{T: uint32(something.LastModifiedAt)}}}},
	}
	_, err := m.Collection.UpdateOne(m.Ctx, filter, update)
	return err
}
