CREATE DATABASE IF NOT EXISTS boletos_ecuador_db;USE boletos_ecuador_db;CREATE TABLE event_models (
	id INT NOT NULL DEFAULT unique_rowid(),
	created_at TIMESTAMPTZ NULL,
	updated_at TIMESTAMPTZ NULL,
	deleted_at TIMESTAMPTZ NULL,
	name STRING NULL,
	date STRING NULL,
	description STRING NULL,
	event_type STRING NULL,
		is_sold_out BOOL NULL,venue_model_id INT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INDEX idx_event_models_deleted_at (deleted_at ASC),
	FAMILY "primary" (id, created_at, updated_at, deleted_at, name, date, description, event_type, is_sold_out, venue_model_id)
);

CREATE TABLE seat_models (
	id INT NOT NULL DEFAULT unique_rowid(),
	created_at TIMESTAMPTZ NULL,
	updated_at TIMESTAMPTZ NULL,
	deleted_at TIMESTAMPTZ NULL,
	number INT NULL,
	is_available BOOL NULL,
	zone_model_id INT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INDEX idx_seat_models_deleted_at (deleted_at ASC),
	FAMILY "primary" (id, created_at, updated_at, deleted_at, number, is_available, zone_model_id)
);

CREATE TABLE ticket_models (
	id INT NOT NULL DEFAULT unique_rowid(),
	created_at TIMESTAMPTZ NULL,
	updated_at TIMESTAMPTZ NULL,
	deleted_at TIMESTAMPTZ NULL,
	number INT NULL,
	seat_model_id INT NULL,
	user_model_id INT NULL,
	event_model_id INT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INDEX idx_ticket_models_deleted_at (deleted_at ASC),
	FAMILY "primary" (id, created_at, updated_at, deleted_at, number, seat_model_id, user_model_id, event_model_id)
);

CREATE TABLE user_models (
	id INT NOT NULL DEFAULT unique_rowid(),
	created_at TIMESTAMPTZ NULL,
	updated_at TIMESTAMPTZ NULL,
	deleted_at TIMESTAMPTZ NULL,
	username STRING NULL,
	age INT NULL,
	user_type STRING NULL,
	password STRING NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INDEX idx_user_models_deleted_at (deleted_at ASC),
	FAMILY "primary" (id, created_at, updated_at, deleted_at, username, age, user_type, password)
);

CREATE TABLE venue_models (
	id INT NOT NULL DEFAULT unique_rowid(),
	created_at TIMESTAMPTZ NULL,
	updated_at TIMESTAMPTZ NULL,
	deleted_at TIMESTAMPTZ NULL,
	venue_type STRING NULL,
	name STRING NULL,
	country STRING NULL,
	city STRING NULL,
	address STRING NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INDEX idx_venue_models_deleted_at (deleted_at ASC),
	FAMILY "primary" (id, created_at, updated_at, deleted_at, venue_type, name, country, city, address)
);

CREATE TABLE zone_models (
	id INT NOT NULL DEFAULT unique_rowid(),
	created_at TIMESTAMPTZ NULL,
	updated_at TIMESTAMPTZ NULL,
	deleted_at TIMESTAMPTZ NULL,
	name STRING NULL,
	price DECIMAL NULL,
	venue_model_id INT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INDEX idx_zone_models_deleted_at (deleted_at ASC),
	FAMILY "primary" (id, created_at, updated_at, deleted_at, name, price, venue_model_id)
);

INSERT INTO event_models (id, created_at, updated_at, deleted_at, name, date, description, event_type, is_sold_out, venue_model_id) VALUES
	(419240354550382593, NULL, NULL, NULL, 'Concierto Morat', '2019-03-03', 'Concierto de Morat en Guayaquil', 'concierto', false, 509239655077937153),
	(419240548598218753, NULL, NULL, NULL, 'Partido UNICEF', '2019-03-04', 'Partido benefico', 'concierto', false, 419239655077937153);

INSERT INTO seat_models (id, created_at, updated_at, deleted_at, number, is_available, zone_model_id) VALUES
	(419244091699789825, NULL, NULL, NULL, 1, true, 419241029616730113),
	(419244114719866881, NULL, NULL, NULL, 2, false, 419241029616730113),
	(419244129873330177, NULL, NULL, NULL, 3, true, 419241029616730113),
	(419244160560168961, NULL, NULL, NULL, 1, true, 419241097457270785),
	(419244189960503297, NULL, NULL, NULL, 2, false, 419241097457270785),
	(809244189960503297, NULL, NULL, NULL, 2, true, 709241029616730115),
	(419244221795860481, NULL, NULL, NULL, 3, true, 419241097457270785);

INSERT INTO ticket_models (id, created_at, updated_at, deleted_at, number, seat_model_id, user_model_id, event_model_id) VALUES
	(419244768004046849, NULL, NULL, NULL, 2, 419244114719866881, 1, 419240354550382593);

INSERT INTO user_models (id, created_at, updated_at, deleted_at, username, age, user_type, password) VALUES
	(1, NULL, NULL, NULL, 'galox14', 22, 'corriente', 'galox14');

INSERT INTO venue_models (id, created_at, updated_at, deleted_at, venue_type, name, country, city, address) VALUES
	(419239655077937153, NULL, NULL, NULL, 'estadio', 'Estadio Alberto Spencer', 'Ecuador', 'Guayaquil', 'Av.de las Americas'),
	(509239655077937153, NULL, NULL, NULL, 'estadio', 'Estadio Monumental', 'Ecuador', 'Guayaquil', 'Av.Barcelona');

INSERT INTO zone_models (id, created_at, updated_at, deleted_at, name, price, venue_model_id) VALUES
	(419241029616730113, NULL, NULL, NULL, 'General', 40.00, 509239655077937153),
	(709241029616730115, NULL, NULL, NULL, 'General', 20.00, 419239655077937153),
	(419241097457270785, NULL, NULL, NULL, 'Tribuna', 60.00, 509239655077937153);
