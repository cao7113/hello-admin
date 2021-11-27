package tables

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"time"
)

func GetBooksTable(ctx *context.Context) table.Table {

	books := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := books.GetInfo() // .HideFilterArea()

	info.AddField("Id", "id", db.Integer).
		FieldFilterable().FieldSortable()
	info.AddField("Name", "name", db.Char).FieldCopyable()
	info.AddField("Score", "score", db.Int)
	info.AddField("Description", "description", db.Char)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("books").SetTitle("Books").SetDescription("Books")

	formList := books.GetForm()
	formList.AddField("Id", "id", db.Integer, form.Default).
		FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("Name", "name", db.Char, form.Text).FieldMust()
	formList.AddField("Score", "score", db.Int, form.Number).FieldDefault("0")
	formList.AddField("Description", "description", db.Char, form.TextArea)
	//formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).
	//	FieldDisableWhenCreate().
	//	FieldDisableWhenUpdate()
	//formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)
	//FieldDisableWhenCreate().
	//FieldDisableWhenUpdate()
	//FieldNowWhenUpdate().
	//FieldDisplayButCanNotEditWhenUpdate()

	formList.SetPostHook(func(values form2.Values) error {
		fmt.Printf("formList.PostHook %+v \n", values)
		return nil
	})
	formList.SetPreProcessFn(func(values form2.Values) form2.Values {
		values.Add("updated_at", time.Now().Format(time.RFC3339))
		return values
	})

	formList.SetTable("books").SetTitle("Books").SetDescription("Books")

	return books
}
