package schemas


type CreateCategoryParams struct {
    Name string
    Description string
}

type PartialUpdateCategoryParams struct {
	ID                int64
	Name              string
	UpdateName        bool
	Description       string
	UpdateDescription bool
}
