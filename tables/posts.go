package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetPostsTable(ctx *context.Context) table.Table {

	posts := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := posts.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Integer).
		FieldFilterable()
	info.AddField("Title", "title", db.Char)
	info.AddField("Body", "body", db.Char)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	formList := posts.GetForm()
	formList.AddField("Id", "id", db.Integer, form.Default)
	formList.AddField("Title", "title", db.Char, form.Text)
	formList.AddField("Body", "body", db.Char, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).FieldNow()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).FieldDisplayButCanNotEditWhenUpdate()

	formList.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	return posts
}
