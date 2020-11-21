CREATE TABLE public.operations_types (
	id serial NOT NULL,
	description varchar NULL,
	CONSTRAINT operations_types_pkey PRIMARY KEY (id)
);

CREATE TABLE public.accounts (
	id serial NOT NULL,
	document_number varchar NOT NULL,
	CONSTRAINT accounts_document_number_key UNIQUE (document_number),
	CONSTRAINT accounts_pkey PRIMARY KEY (id)
);

CREATE TABLE public.transactions (
	id serial NOT NULL,
	account_id int4 NOT NULL,
	operation_type_id int4 NOT NULL,
	amount numeric(10,2) NOT NULL,
	event_date timestamp NOT NULL,
	CONSTRAINT transactions_pkey PRIMARY KEY (id)
);

insert into public.operations_types (id,description) values (1,'COMPRA A VISTA');
insert into public.operations_types (id,description) values (2,'COMPRA PARCELADA');
insert into public.operations_types (id,description) values (3,'SAQUE');
insert into public.operations_types (id,description) values (4,'PAGAMENTO');
