package model

//"Task"
type Task struct {
    ID        primitive.ObjectID `bson:"_id"`
    CreatedAt time.Time          `bson:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at"`
    Text      string             `bson:"text"`
    Completed bool               `bson:"completed"`
}