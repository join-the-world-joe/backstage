package business

import (
	"backstage/common/code"
	"backstage/common/db/mgo/backend/track"
	timestamp2 "backstage/common/macro/timestamp"
	"backstage/common/major"
	"backstage/common/protocol/admin"
	"backstage/global/config"
	"backstage/global/log"
	"backstage/global/mgo"
	"backstage/utils/convert"
	"backstage/utils/timestamp"
	"context"
	json2 "encoding/json"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
)

type Track struct {
	Operator   string           `json:"operator"`
	Major      string           `json:"major"`
	Minor      string           `json:"minor"`
	Permission string           `json:"permission"`
	Request    json2.RawMessage `json:"request"`
	Response   json2.RawMessage `json:"response"`
	Timestamp  string           `json:"timestamp"`
}

type TrackListOutput struct {
	TrackList []*Track `json:"track_list"`
	Length    int      `json:"length"`
}

func FetchTrackListOfCondition(ctx context.Context, req *admin.FetchTrackListOfConditionReq, rsp *admin.FetchTrackListOfConditionRsp) error {
	if !hasPermission(
		cast.ToInt(major.Admin),
		cast.ToInt(admin.FetchTrackListOfConditionReq_),
		req.UserId,
	) {
		rsp.Code = code.AccessDenied
		return nil
	}

	tlo := &TrackListOutput{TrackList: []*Track{}, Length: 0}

	if req.Behavior == 1 { // for today's track
		location := timestamp2.AsiaShanghai
		year, month, day := timestamp.GetYearMonthDay()
		begin := timestamp.GetBeginOfADay(year, month, day, location)
		end := timestamp.GetEndOfADay(year, month, day, location)

		where := &bson.D{
			{"timestamp",
				bson.D{
					{"$lte", end},
					{"$gte", begin},
				},
			},
		}

		cur, err := mgo.Query(
			config.MongoConf(),
			context.Background(),
			track.GetWhich(),
			track.GetDBName(),
			track.GetTableName(),
			where,
		)
		if err != nil {
			rsp.Code = code.InternalError
			return nil
		}

		for cur.Next(context.Background()) {
			doc := &track.Model{}
			err = cur.Decode(doc)
			if err != nil {
				log.Error("FetchTrackListOfCondition failure, err: ", err.Error())
				continue
			}
			tlo.TrackList = append(tlo.TrackList, &Track{
				Operator:   string(doc.Operator),
				Major:      doc.Major,
				Minor:      doc.Minor,
				Permission: doc.Permission,
				Request:    []byte(convert.BytesStringToString(doc.Request)),
				Response:   []byte(convert.BytesStringToString(doc.Response)),
				Timestamp:  timestamp.ToDateTimeString(doc.Timestamp),
			})
			tlo.Length += 1
		}
	} else if req.Behavior == 2 { // track with filter of the input conditions
		if len(req.Operator) <= 0 && len(req.Major) <= 0 && len(req.Minor) <= 0 && len(req.Permission) <= 0 && req.Begin <= 0 && req.End <= 0 {
			rsp.Code = code.NoData
			return nil
		}

		where := bson.D{}
		if len(req.Operator) > 0 {
			where = append(where, bson.E{Key: "operator", Value: bson.M{"$eq": string(req.Operator)}})
		}
		if len(req.Major) > 0 {
			where = append(where, bson.E{Key: "major", Value: bson.M{"$eq": req.Major}})
			//where["major"] = bson.M{"$eq": req.Major}
		}
		if len(req.Minor) > 0 {
			where = append(where, bson.E{Key: "minor", Value: bson.M{"$eq": req.Minor}})
			//where["minor"] = bson.M{"$eq": req.Minor}
		}
		if len(req.Permission) > 0 {
			where = append(where, bson.E{Key: "permission", Value: bson.M{"$eq": req.Permission}})
			//where[req.Permission] = bson.M{"$eq": req.Permission}
		}
		if req.Begin > 0 {
			where = append(where, bson.E{Key: "timestamp", Value: bson.M{"$gte": req.Begin}})
			//where["timestamp"] = bson.M{"$gte": req.Begin}
		}
		if req.End > 0 {
			where = append(where, bson.E{Key: "timestamp", Value: bson.M{"$lte": req.End}})
			//where["timestamp"] = bson.M{"$lte": req.End}
		}

		cur, err := mgo.Query(
			config.MongoConf(),
			context.Background(),
			track.GetWhich(),
			track.GetDBName(),
			track.GetTableName(),
			&where,
		)
		if err != nil {
			rsp.Code = code.InternalError
			return nil
		}

		for cur.Next(context.Background()) {
			doc := &track.Model{}
			//fmt.Println("doc: ", doc)
			err = cur.Decode(doc)
			if err != nil {
				log.Error("FetchTrackListOfCondition failure, err: ", err.Error())
				continue
			}
			tlo.TrackList = append(tlo.TrackList, &Track{
				Operator:   string(doc.Operator),
				Major:      doc.Major,
				Minor:      doc.Minor,
				Permission: doc.Permission,
				Request:    []byte(convert.BytesStringToString(doc.Request)),
				Response:   []byte(convert.BytesStringToString(doc.Response)),
				Timestamp:  timestamp.ToDateTimeString(doc.Timestamp),
			})
			tlo.Length += 1
		}
	} else {
		rsp.Code = code.InvalidData
		return nil
	}

	bytes, err := json2.Marshal(tlo)
	if err != nil {
		log.Error("FetchTrackListOfCondition failure, err: ", err.Error())
		rsp.Code = code.InternalError
		return nil
	}

	rsp.Body = bytes
	rsp.Code = code.Success
	return nil
}
