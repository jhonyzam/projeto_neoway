CREATE TABLE datastore(
	   id serial PRIMARY KEY,
	   cpf VARCHAR (50) NOT NULL,
	   private INTEGER,
	   incompleto INTEGER,
	   lastDate DATE,
	   avgTicket NUMERIC (10, 2),
	   lastTicket NUMERIC (10, 2),
	   storeFrequent VARCHAR (50),
	   storeLast VARCHAR (50)
);

