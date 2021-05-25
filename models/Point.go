package models

type PointRepository interface {
	createPoint(point *Point) bool
	getPoint(pointId int) Point
	putPoint(point *Point) bool
	deletePoint(pointId int) Point
}

type Point struct {
	PointId     int     `dynamodbav:"PointId" json:"id"`
	Latitude    float64 `dynamodbav:"Latitude" json:"latitude"`
	Longitude   float64 `dynamodbav:"Longitude" json:"longitude"`
	Description string  `dynamodbav:"Description" json:"description"`
}
