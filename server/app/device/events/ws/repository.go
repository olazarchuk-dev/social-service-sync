package ws

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"social-service-sync/server/app/helper"
	"social-service-sync/server/model/entity"
	"time"
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
		{"$set", bson.D{{"last_modified_at", MillisecondsToTimestamp(something.LastModifiedAt)}}},
	}
	_, err := m.Collection.UpdateOne(m.Ctx, filter, update)
	return err
}

/*
 * @see https://expange.ru/e/Конвертировать_миллисекунды_в_time.Time_(GoLang)
 *      https://stackoverflow.com/questions/31744970/convert-to-time-in-golang-from-milliseconds
 *      https://golangbyexample.com/time-conversion-in-golang
 *
 * timeUnix 1580146705
 * time.Time: 2020-01-27 23:08:25 +0530 IST
 * timeUnixMilli: 1580146705000
 * timeUnixMicro: 1580146705000000
 * timeUnixNano: 1580146705000000000
 */
func MillisecondsToTimestamp(milliseconds int) primitive.Timestamp {
	t := MillisecondsToTime(milliseconds)
	return TimeToTimestamp(t)
}

func TimeToTimestamp(t time.Time) primitive.Timestamp {
	return primitive.Timestamp{T: uint32(t.Unix()), I: 0}
}

func MillisecondsToTime(milliseconds int) time.Time {
	return time.Unix(0, int64(milliseconds)*int64(time.Millisecond))
}
