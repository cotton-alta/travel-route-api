package models

type SegmentRepository interface {
	createSegment(segment *Segment) bool
	getSegment(segmentId int) Segment
	putSegment(segment *Segment) bool
	addPointToSegment(segmentId int, point *Point) bool
	deletePoint(pointId int) Segment
}

type Segment struct {
	// Id          int     `dynamodbav:"id" json:id`
	// Latitude    float64 `dynamodbav:"latitude" json:latitude`
	// Longitude   float64 `dynamodbav:"longitude" json:longitude`
	// Description string  `dynamodbav:"description" json:description`
}
