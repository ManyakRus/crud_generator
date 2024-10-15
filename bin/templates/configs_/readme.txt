convert_id.json:
Mapping database table name . field name = field type in golang.
For tables file. 
For non standart rare case or type aliases
example:
{
  "lawsuit_payments.id": "alias.PaymentId",
  "lawsuit_payments.invoice_id": "alias.InvoiceId",
}



------------------------------------------------------------------------
mapping.json
Mapping Postgres types to Golang types



------------------------------------------------------------------------
name_replace.json
Mapping database field name = golang field name
For tables file. 
Replace standart field name to filled name
For non standart rare case.
example:
{
  "inn": "INN",
  "json_file_id": "JSONFileID"
}


------------------------------------------------------------------------
nullable.json
List of golang field names, which need change 0 to null.
true = need change 0 to null
For non standart rare case.
example:
{
      "ext_id": true,
      "created_at": false
}


------------------------------------------------------------------------
model_crud_delete_functions.json
Mapping postgres tables to golang function name,
this functions will be deleted from model crud files.
example:
{
      "lawsuits": "NewLawsuit,AsLawsuit"
}



------------------------------------------------------------------------
crud_functions_rename.json 
TableName:{old:"",new:""}
example:
{
      "functions": [
	{
	"old": "create_update_ctx",
	"new":"create_update_ctx_original"
	}
	]
}

------------------------------------------------------------------------
findby_functions.json
Need for automatic create functions searching 1 row in table filtered by column_name, 
example:
[{"Table":"table_name1","Columns":["column_name1"]}]
[{"Table":"table_name1","Columns":["column_name1","column_name2"]}]


------------------------------------------------------------------------
findmassby_function.json
Need for automatic create functions searching many rows in table filtered by column_name, 
example:
[{"Table":"table_name1","Columns":["column_name1"]}]
[{"Table":"table_name1","Columns":["column_name1","column_name2"]}]


readall_function.json
Need for automatic create functions ReadAll() returning all rows in table (exclude is_deleted rows)
example:
{
      "table_name1": ""
}


findmodelby_functions.json
Need for automatic create FindModelBy() functions searching 1 value in table_name1 filtered by column_name1,
this value will be searched in foreign table as identifier, and return this 1 row model
example:
[
	{"TableName":"table_name1","ColumnName":"column_name1"}
]

