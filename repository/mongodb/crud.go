package mongodb

import (
	"context"

	"github.com/uniplaces/carbon"
	"go.mongodb.org/mongo-driver/mongo"

	"go-template/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) List(ctx context.Context, opt *domain.PageOption, itemType interface{}) (total int, items []interface{}, err error) {
	var filters bson.M
	var opts *options.FindOptions
	if opt != nil {
		opts = repo.makePagingOpts(opt.Page, opt.PerPage)
		if opt.Filters != nil && len(opt.Filters) > 0 && opt.Search != nil && len(opt.Search) > 0 {
			filters = bson.M{
				"$and": bson.A{
					repo.makeFilters(opt.Filters),
				},
				"$or": repo.makeSearch(opt.Search),
			}
		} else if opt.Filters != nil && len(opt.Filters) > 0 {
			filters = repo.makeFilters(opt.Filters)
		} else if opt.Search != nil && len(opt.Search) > 0 {
			filters = bson.M{
				"$or": repo.makeSearch(opt.Search),
			}
		}

		if opt.Sorts != nil && len(opt.Sorts) > 0 {
			opts.Sort = repo.makeSorts(opt.Sorts)
		}
	}

	total, err = repo.countWithBson(ctx, filters)
	if err != nil {
		return 0, nil, err
	}

	cursor, err := repo.Coll.Find(ctx, filters, opts)
	if err != nil {
		return 0, nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	for cursor.Next(ctx) {
		item, err := repo.clone(itemType)
		if err != nil {
			return 0, nil, err
		}
		err = cursor.Decode(item)
		if err != nil {
			return 0, nil, err
		}
		items = append(items, item)
	}

	return total, items, nil
}

func (repo *Repository) Find(ctx context.Context, opt *domain.QueryOption, itemType interface{}) (total int, items []interface{}, err error) {
	var filters bson.M
	var opts *options.FindOptions
	opts = options.Find()
	if opt != nil {
		if opt.Filters != nil && len(opt.Filters) > 0 && opt.Search != nil && len(opt.Search) > 0 {
			filters = bson.M{
				"$and": bson.A{
					repo.makeFilters(opt.Filters),
				},
				"$or": repo.makeSearch(opt.Search),
			}
		} else if opt.Filters != nil && len(opt.Filters) > 0 {
			filters = repo.makeFilters(opt.Filters)
		} else if opt.Search != nil && len(opt.Search) > 0 {
			filters = bson.M{
				"$or": repo.makeSearch(opt.Search),
			}
		}

		if opt.Sorts != nil && len(opt.Sorts) > 0 {
			opts.Sort = repo.makeSorts(opt.Sorts)
		}
	}

	total, err = repo.countWithBson(ctx, filters)
	if err != nil {
		return 0, nil, err
	}

	cursor, err := repo.Coll.Find(ctx, filters, opts)
	if err != nil {
		return 0, nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	for cursor.Next(ctx) {
		item, err := repo.clone(itemType)
		if err != nil {
			return 0, nil, err
		}
		err = cursor.Decode(item)
		if err != nil {
			return 0, nil, err
		}
		items = append(items, item)
	}

	return total, items, nil
}

func (repo *Repository) Create(ctx context.Context, data interface{}) (ID string, err error) {
	res, err := repo.Coll.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(string), nil
}

func (repo *Repository) Read(ctx context.Context, filters []string, out interface{}) (err error) {
	conditions := repo.makeFilters(filters)
	return repo.Coll.FindOne(ctx, conditions).Decode(out)
}

func (repo *Repository) Update(ctx context.Context, filters []string, data interface{}) (err error) {
	conditions := repo.makeFilters(filters)
	_, err = repo.Coll.UpdateOne(ctx, conditions, bson.M{"$set": data})
	return err
}

func (repo *Repository) UpdateWithBson(ctx context.Context, filters []string, data bson.M) (err error) {
	conditions := repo.makeFilters(filters)
	_, err = repo.Coll.UpdateOne(ctx, conditions, data)
	return err
}

func (repo *Repository) Delete(ctx context.Context, filters []string) (err error) {
	conditions := repo.makeFilters(filters)
	_, err = repo.Coll.DeleteOne(ctx, conditions)
	return err
}

func (repo *Repository) SoftDelete(ctx context.Context, filters []string) (err error) {
	conditions := repo.makeFilters(filters)
	_, err = repo.Coll.UpdateOne(ctx, conditions, bson.M{"$set": bson.M{"deleted_at": carbon.Now().Unix()}})
	return err
}

func (repo *Repository) Count(ctx context.Context, filters []string) (total int, err error) {
	cnt, err := repo.Coll.CountDocuments(ctx, repo.makeFilters(filters))
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (repo *Repository) countWithBson(ctx context.Context, filters bson.M) (total int, err error) {
	cnt, err := repo.Coll.CountDocuments(ctx, filters)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (repo *Repository) Aggregate(ctx context.Context, pipeline mongo.Pipeline, page, perPage int, itemType interface{}) (total int, items []interface{}, err error) {

	total, err = repo.AggregateCount(ctx, pipeline)
	if err != nil {
		return 0, nil, err
	}

	if page != 0 && perPage != 0 {
		skip := bson.D{
			{"$skip", (page - 1) * perPage},
		}
		limit := bson.D{
			{"$limit", perPage},
		}
		pipeline = append(pipeline, skip, limit)
	}

	cursor, err := repo.Coll.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, nil, err
	}

	for cursor.Next(ctx) {
		item, err := repo.clone(itemType)
		if err != nil {
			return 0, nil, err
		}
		err = cursor.Decode(item)
		if err != nil {
			return 0, nil, err
		}
		items = append(items, item)
	}
	return total, items, nil
}

func (repo *Repository) AggregateCount(ctx context.Context, pipeline mongo.Pipeline) (total int, err error) {
	count := bson.D{
		{"$count", "total"},
	}
	pipeline = append(pipeline, count)
	cursor, err := repo.Coll.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}

	var totalData map[string]int
	for cursor.Next(ctx) {
		err = cursor.Decode(&totalData)
		if err != nil {
			return 0, err
		}
	}

	return totalData["total"], err
}

func (repo *Repository) CountBson(ctx context.Context, filters bson.M) (total int, err error) {
	cnt, err := repo.Coll.CountDocuments(ctx, filters)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}
