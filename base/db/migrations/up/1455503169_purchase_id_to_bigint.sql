ALTER TABLE purchases DROP CONSTRAINT purchases_pkey;
ALTER TABLE purchases DROP COLUMN id;
ALTER TABLE purchases ADD COLUMN id BIGINT;
ALTER TABLE purchases ADD PRIMARY KEY (id);