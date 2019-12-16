package sql

// InsertAccounts insert accounts sql
var InsertAccounts = `
	INSERT INTO accounts
	(id,name,email,password,created_by,updated_by,created_at,updated_at,deleted_at)
	VALUES
	('a44c9e48-790d-471e-8fc8-7681ad9eda6b','サンプル','sample@example.com','$2a$10$4OPY6NlJf9uUoruhZBGVVeMyIvXpXrRByBdcktKw7CnKOT40Clv/S',null,null,now(),now(),null);
`
