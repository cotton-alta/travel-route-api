package models

type SegmentRepository interface {
	createSegment(point *Segment) bool
	getSegment(pointId int) Segment
	putSegment(point *Segment) bool
	addPointToSegment(point *Point)
	deletePoint(pointId int) Segment
}

type Segment struct {
	Id          int     `dynamodbav:"id" json:id`
	Latitude    float64 `dynamodbav:"latitude" json:latitude`
	Longitude   float64 `dynamodbav:"longitude" json:longitude`
	Description string  `dynamodbav:"description" json:description`
}
