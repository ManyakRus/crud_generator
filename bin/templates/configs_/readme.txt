convert_id.json:
Mapping database table name . field name = field type in golang.
For tables file. 
For non standart rare case or type aliases
example:
{
  "lawsuit_payments.id": "alias.PaymentId",
  "lawsuit_payments.invoice_id": "alias.InvoiceId",
}



mapping.json
Mapping Postgres types to Golang types



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


nullable.json
List of golang field names, which need change 0 to null.
true = need change 0 to null
For non standart rare case.
example:
{
      "ext_id": true,
      "created_at": false
}


model_crud_delete_functions.json
Mapping postgres tables to golang function name,
this functions will be deleted from model crud files.
example:
{
      "lawsuits": "NewLawsuit,AsLawsuit"
}
