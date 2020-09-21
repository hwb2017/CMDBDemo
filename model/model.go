package model

import "go.mongodb.org/mongo-driver/mongo"

type Collection struct {
	client *mongo.Client
	dbName string
	collectionName string
}

var (
	VMLifecycleCollection *Collection
	VMLifecycleAssociationCollection *Collection
)

func (c *Collection) mongoCollection() *mongo.Collection {
	return c.client.Database(c.dbName).Collection(c.collectionName)
}

func initializeCollection(dbName, collectionName string) *Collection {
	coll := new(Collection)
	coll.dbName = dbName
	coll.collectionName = collectionName
	return coll
}

func init() {
	VMLifecycleCollection = initializeCollection("infrastructure", "vm_lifecycle")
    VMLifecycleAssociationCollection = initializeCollection("infrastructure", "vm_lifecycle_association")
}