package utils

import "go.mongodb.org/mongo-driver/v2/bson"

func Contains(a []bson.ObjectID, b bson.ObjectID) bool {
	for _, elem := range a {
		if elem == b {
			return true
		}
	}
	return false
}
