package cmd


type BlogTag struct {
	// 创建人
	CreatedBy      string  `json:"created_by"`
	// 创建时间
	CreatedOn      int32   `json:"created_on"`
	// 删除时间
	DeletedOn      int32   `json:"deleted_on"`
	// id
	Id     int32   `json:"id"`
	// 是否删除 0:未删除 1:已删除
	IsDel  int8    `json:"is_del"`
	// 修改人
	ModifiedBy     string  `json:"modified_by"`
	// 修改时间
	ModifiedOn     int32   `json:"modified_on"`
	// 标签名称
	Name   string  `json:"name"`
	// 状态 0:禁用 1:启用
	State  int8    `json:"state"`
}
func (model BlogTag) TableName() string {
	return "blog_tag"
}