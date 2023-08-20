package dtos

type ListUsersDto struct {
	Email          string `query:"email"`
	RealmCode      string `query:"realm_code"`
	RecordsPerPage string `query:"records_per_page"`
	Page           string `query:"page"`
}
