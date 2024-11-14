package repository

import "fmt"

type Repo struct {
	db map[int]struct{}
}

func NewRepo() *Repo {
	db := make(map[int]struct{})
	return &Repo{
		db: db,
	}
}

func (r *Repo) Save(value int) {
	r.db[value] = struct{}{}

	val, err := r.Gat(10)
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

}

func (r *Repo) Gat(value int) (int, error) {
	_, ok := r.db[value]
	if !ok {
		return 0, fmt.Errorf("value %d not exists", value)
	}
	return value, nil
}
