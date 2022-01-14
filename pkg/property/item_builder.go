package property

type ItemBuilder struct {
	base itemBase
}

func NewItem() *ItemBuilder {
	return &ItemBuilder{}
}

func (b *ItemBuilder) Group() *GroupBuilder {
	return NewGroup().base(b.base)
}

func (b *ItemBuilder) GroupList() *GroupListBuilder {
	return NewGroupList().base(b.base)
}

func (b *ItemBuilder) ID(id ItemID) *ItemBuilder {
	b.base.ID = id
	return b
}

func (b *ItemBuilder) NewID() *ItemBuilder {
	b.base.ID = NewItemID()
	return b
}

func (b *ItemBuilder) Schema(s SchemaID, g SchemaGroupID) *ItemBuilder {
	b.base.Schema = s
	b.base.SchemaGroup = g
	return b
}
