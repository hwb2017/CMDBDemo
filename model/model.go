package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type Collection struct {
	DBName string
	CollectionName string
}

type BulkSyncModels struct {
	ModelMapping map[string]interface{}
	InsertIDs []string
	DeleteIDs []string
	UpdateIDs []string
}

type QueryOptions struct {
	filter map[string]interface{}
	projection map[string]interface{}
	sort map[string]interface{}
	skip int64
	limit int64
	extendAggregationPipelineStages []bson.D
}

func (q *QueryOptions) WithFilter(f map[string]interface{}) *QueryOptions {
	q.filter = f
	return q
}

func (q *QueryOptions) WithSimpleProjection(fields ...string) *QueryOptions {
	projection := make(map[string]interface{})
	for _, field := range fields {
		if len(field) <= 0 {
			continue
		}
		projection[field] = 1
	}
	q.projection = projection
	return q
}

func (q *QueryOptions) WithProjection(projection bson.M) *QueryOptions {
	q.projection = projection
	return q
}

func (q *QueryOptions) WithSort(sorter string) *QueryOptions {
	sort := make(map[string]interface{})
	if sorter != "" {
		sortSlice := strings.Split(sorter, ",")
		for _, sortItem := range sortSlice {
			sortKey := strings.TrimLeft(sortItem, "+-")
			if strings.HasPrefix(sortItem,"+") {
				sort[sortKey] = 1
			} else {
				sort[sortKey] = -1
			}
		}
	}
	return q
}

func (q *QueryOptions) WithExtendAggregationPipelineStages(pipelines [] bson.D) *QueryOptions {
    q.extendAggregationPipelineStages = pipelines
    return q
}

func (q *QueryOptions) WithSkip(skip int) *QueryOptions {
	q.skip = int64(skip)
	return q
}

func (q *QueryOptions) WithLimit(limit int) *QueryOptions {
	q.limit = int64(limit)
	return q
}

func (c Collection) mongodbCollection(client *mongo.Client) *mongo.Collection{
	return client.Database(c.DBName).Collection(c.CollectionName)
}

func (c Collection) handleCursor(cursor *mongo.Cursor) ([]bson.M, error) {
	defer cursor.Close(context.TODO())

	results := make([]bson.M, 0, 10)
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {return nil, err}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (c Collection) find(collection *mongo.Collection, queryOptions *QueryOptions) (*mongo.Cursor, error) {
	findOpts := options.Find()
	if queryOptions.filter == nil {
		queryOptions.filter = bson.M{}
	}
	if queryOptions.projection != nil {
		findOpts.SetProjection(queryOptions.projection)
	}
	if len(queryOptions.sort) != 0 {
		findOpts.SetSort(queryOptions.sort)
	}
	if queryOptions.skip != 0 {
		findOpts.SetSkip(queryOptions.skip)
	}
	if queryOptions.limit != 0 {
		findOpts.SetLimit(queryOptions.limit)
	}
	return collection.Find(context.TODO(), queryOptions.filter, findOpts)
}

func (c Collection) aggregate(collection *mongo.Collection, queryOptions *QueryOptions) (*mongo.Cursor, error) {
	pipeline := mongo.Pipeline{}
	if queryOptions.filter == nil {
		queryOptions.filter = bson.M{}
	}
	filterStage := bson.D{{
		"$match",queryOptions.filter,
	}}
	pipeline = append(pipeline, filterStage)

	if queryOptions.projection != nil {
		if len(queryOptions.projection) >= 1 {
			projectStage := bson.D{{
				"$project",queryOptions.projection,
			}}
			pipeline = append(pipeline, projectStage)
		}
	}
	if queryOptions.sort != nil {
		if len(queryOptions.sort) >= 1 {
			sortStage := bson.D{{
				"$sort", queryOptions.sort,
			}}
			pipeline = append(pipeline, sortStage)
		}
	}
	if queryOptions.skip > 0 {
		skipStage := bson.D{{
			"$skip", queryOptions.skip,
		}}
		pipeline = append(pipeline, skipStage)
	}
    if queryOptions.limit > 0 {
		limitStage := bson.D{{
			"$limit", queryOptions.limit,
		}}
		pipeline = append(pipeline, limitStage)
	}
	if queryOptions.extendAggregationPipelineStages != nil {
		pipeline = append(pipeline, queryOptions.extendAggregationPipelineStages...)
	}
	return collection.Aggregate(context.TODO(), pipeline)
}

func (c Collection) Find(client *mongo.Client, queryOptions *QueryOptions) ([]bson.M, error){
	collection := c.mongodbCollection(client)
	cursor, err := c.find(collection, queryOptions)
	if err != nil {
		return nil, err
	}
	return c.handleCursor(cursor)
}

// BulkSync combine insert, delete and replace data and execute then in bulk
func (c Collection) BulkSync(client *mongo.Client, m BulkSyncModels) error {
	collection := c.mongodbCollection(client)

	bulkOpsNum := len(m.InsertIDs) + len(m.DeleteIDs) + len(m.UpdateIDs)
	bulkWriteModels := make([]mongo.WriteModel, 0, bulkOpsNum)
	//设置批量同步中要插入的部分
	for _, v := range m.InsertIDs {
		doc := m.ModelMapping[v]
		model := mongo.NewInsertOneModel().SetDocument(doc)
		bulkWriteModels = append(bulkWriteModels, model)
	}
	//设置批量同步中要删除的部分
	for _, v := range m.DeleteIDs {
		model := mongo.NewDeleteOneModel().SetFilter(bson.M{"_id": v})
		bulkWriteModels = append(bulkWriteModels, model)
	}
	//设置批量同步中要更新的部分
	for _, v := range m.UpdateIDs {
		doc := m.ModelMapping[v]
		model :=  mongo.NewReplaceOneModel().SetFilter(bson.M{"_id": v}).SetReplacement(doc)
		bulkWriteModels = append(bulkWriteModels, model)
	}
	bulkWriteOpts := options.BulkWrite().SetOrdered(false)
	_, err := collection.BulkWrite(context.TODO(), bulkWriteModels, bulkWriteOpts)
	if err != nil {
		return err
	}
	return nil
}