package utils

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"slices"
)

func Contains(a []bson.ObjectID, b bson.ObjectID) bool {
	for _, elem := range a {
		if elem == b {
			return true
		}
	}
	return false
}

func RemoveElement(elem bson.ObjectID, array []bson.ObjectID) []bson.ObjectID {
	for _, a := range array {
		if a == elem {
			index := slices.Index(array, elem)
			if index != -1 {
				array = append(array[:index], array[index+1:]...)
			}
		}
	}
	return array
}
