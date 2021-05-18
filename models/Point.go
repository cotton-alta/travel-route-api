package models

type PointRepository interface {
	createPoint(point *Point) bool
	getPoint(pointId int) Point
	putPoint(point *Point) bool
	deletePoint(pointId int) Point
}

type Point struct {
	Id          int     `dynamodbav:"id" json:id`
	Latitude    float64 `dynamodbav:"latitude" json:latitude`
	Longitude   float64 `dynamodbav:"longitude" json:longitude`
	Description string  `dynamodbav:"description" json:description`
}
