package domain

type Staff struct {
	ID        string `bson:"_id,omitempty"`
	Name      string `bson:"name"`
	Tel       string `bson:"tel"`
	Status    string `bson:"status"`
	Detail    string `bson:"detail"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
	DeletedAt int64  `bson:"deleted_at,omitempty"`
}
