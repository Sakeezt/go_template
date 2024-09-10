package mongodb

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) clone(origin interface{}) (clone interface{}, err error) {
	newClone := reflect.New(reflect.TypeOf(origin).Elem()).Interface()
	return newClone, copier.Copy(newClone, origin)
}

func (repo *Repository) makeFilters(filters []string) (bsonFilters bson.M) {
	bsonFilters = bson.M{}
	for _, v := range filters {
		slFilter := strings.Split(v, ":")
		key := slFilter[0]
		operations := slFilter[1]

		var valInt int
		var valStr string
		var err error
		if len(slFilter) > 2 {
			switch operations {
			case "lt", "lte", "gt", "gte":
				valInt, err = repo.convertIntValue(slFilter[2])
				if err != nil {
					return bsonFilters
				}
			default:
				valStr = slFilter[2]
			}
		}

		switch operations {
		case "eq":
			bsonFilters[key] = valStr
			break
		case "ne":
			bsonFilters[key] = bson.M{"$ne": valStr}
			break
		case "lt":
			bsonFilters[key] = bson.M{"$lt": valInt}
			break
		case "lte":
			bsonFilters[key] = bson.M{"$lte": valInt}
			break
		case "gt":
			bsonFilters[key] = bson.M{"$gt": valInt}
			break
		case "gte":
			bsonFilters[key] = bson.M{"$gte": valInt}
			break
		case "in":
			slValue := strings.Split(valStr, "|")
			if len(slValue) > 0 {
				bsonFilters[key] = bson.M{"$in": slValue}
			} else {
				bsonFilters[key] = nil
			}
			break
		case "isNull":
			bsonFilters[key] = nil
			break
		case "isNotNull":
			bsonFilters[key] = bson.M{"$ne": nil}
			break
		case "like":
			bsonFilters[key] = bson.M{
				"$regex":   valStr,
				"$options": "i",
			}
			break
		case "between":
			bsonFilters["$and"] = makeBetween(slFilter)
			break
		default:
			if len(slFilter) > 2 {
				bsonFilters[key] = valStr
			} else {
				bsonFilters[key] = nil
			}
			break
		}
	}

	return bsonFilters
}

func (repo *Repository) makeSearch(search []string) (bsonSearch bson.A) {
	bsonSearch = bson.A{}
	for _, v := range search {
		bsonItem := bson.M{}
		slSearch := strings.Split(v, ":")
		key := slSearch[0]
		value := slSearch[1]

		bsonItem[key] = bson.M{
			"$regex":   *repo.addBackSlash(&value),
			"$options": "i",
		}

		bsonSearch = append(bsonSearch, bsonItem)
	}
	return bsonSearch
}

func (repo *Repository) makeSorts(sorts []string) (bsonSorts bson.M) {
	bsonSorts = bson.M{}

	for _, v := range sorts {
		slFilter := strings.Split(v, ":")
		field := slFilter[0]
		order := slFilter[1]
		bsonSorts[field] = -1
		if order == "asc" {
			bsonSorts[field] = 1
		}
	}

	return bsonSorts
}

func (repo *Repository) makePagingOpts(page int, perPage int) (opts *options.FindOptions) {
	skip := (page - 1) * perPage
	opts = options.Find()
	opts.SetSkip(int64(skip))

	if perPage > 0 {
		opts.SetLimit(int64(perPage))
	}

	return opts
}

func (repo *Repository) convertIntValue(str string) (int, error) {
	var valInt int
	matched, _ := regexp.Match(`^\d+$`, []byte(str))
	if matched {
		var err error
		valInt, err = strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		return valInt, nil
	} else {
		return 0, errors.New("invalid number")
	}
}

func (repo *Repository) addBackSlash(str *string) *string {
	var res string
	reg, err := regexp.Compile("[^A-Za-z0-9]")
	if err != nil {
		return &res
	}
	res = reg.ReplaceAllStringFunc(*str, func(s string) string {
		if s != " " {
			return "\\" + s
		} else {
			return s
		}
	})
	return &res
}

func makeBetween(slFilter []string) (arrayFilters bson.A) {
	key := makeKey(slFilter[0])
	value := strings.Split(slFilter[2], "|")
	if len(value) != 2 {
		return
	}
	gte, err := strconv.Atoi(value[0])
	if err != nil {
		return
	}
	lte, err := strconv.Atoi(value[1])
	if err != nil {
		return
	}
	arrayFilters = bson.A{
		bson.M{key: bson.M{
			"$gte": gte},
		},
		bson.M{key: bson.M{
			"$lte": lte},
		},
	}
	return arrayFilters
}

func makeKey(k string) (key string) {
	keys := strings.Split(k, ".")
	key = strcase.ToLowerCamel(k)
	if len(keys) > 1 {
		for i, v := range keys {
			keys[i] = strcase.ToLowerCamel(v)
		}
		key = strings.Join(keys, ".")
	}

	return key
}
