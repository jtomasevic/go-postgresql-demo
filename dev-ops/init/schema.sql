
DO
$do$
    BEGIN
    IF EXISTS (
        SELECT FROM pg_catalog.pg_roles
        WHERE  rolname = 'imdb_user') THEN

        RAISE NOTICE 'Role "imdb_user" already exists. Skipping.';
    ELSE
        CREATE ROLE imdb_user WITH
            LOGIN
            NOSUPERUSER
            CREATEDB
            NOCREATEROLE
            INHERIT
            REPLICATION
            BYPASSRLS
            CONNECTION LIMIT -1
            PASSWORD 'imdb_user';
    END IF;
    END
$do$;

DO
$do$
BEGIN

IF EXISTS (SELECT FROM pg_database WHERE datname = 'imdb') THEN
      RAISE NOTICE 'Database already exists';  -- optional
   ELSE
    --   PERFORM dblink_exec('dbname=' || current_database()  -- current db
    -- , 'CREATE DATABASE mydb');
    CREATE DATABASE imdb
        WITH
        OWNER = imdb_user
        TEMPLATE = postgres
        ENCODING = 'UTF8'
        STRATEGY = 'file_copy'
        LC_COLLATE = 'en_US.utf8'
        LC_CTYPE = 'en_US.utf8'
        LOCALE_PROVIDER = 'libc'
        TABLESPACE = pg_default
        CONNECTION LIMIT = -1
        IS_TEMPLATE = False;

   END IF;

END
$do$;

GRANT ALL ON DATABASE imdb TO imdb_user WITH GRANT OPTION;
GRANT ALL ON ALL TABLES IN SCHEMA public TO imdb_user;

CREATE TABLE IF NOT EXISTS Movies (
    ID UUID PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    ReleaseYear INT    NOT NULL
);

CREATE TABLE IF NOT EXISTS Actors (
    ID UUID PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    BirthYear INT
);

CREATE TABLE IF NOT EXISTS Directors (
    ID UUID PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    BirthYear INT
);

CREATE TABLE IF NOT EXISTS Reviews (
    ID UUID PRIMARY KEY,
    MovieID UUID NOT NULL,
    Comment TEXT NOT NULL,
    Rating INT NOT NULL,
    CommentTime   TIMESTAMP NOT NULL,
    FOREIGN KEY (MovieID) REFERENCES Movies(ID)
);

CREATE TABLE IF NOT EXISTS Awards (
    ID UUID PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Year INT NOT NULL,
    MovieID UUID NOT NULL,
    ActorID UUID,
    DirectorID UUID,
    FOREIGN KEY (MovieID) REFERENCES Movies(ID),
    FOREIGN KEY (ActorID) REFERENCES Actors(ID),
    FOREIGN KEY (DirectorID) REFERENCES Directors(ID)
);

CREATE TABLE MovieActors (
    MovieID UUID NOT NULL,
    ActorID UUID NOT NULL,
    FOREIGN KEY (MovieID) REFERENCES Movies(ID),
    FOREIGN KEY (ActorID) REFERENCES Actors(ID),
    PRIMARY KEY (MovieID, ActorID)
);

CREATE TABLE IF NOT EXISTS MovieDirectors (
    MovieID UUID NOT NULL,
    DirectorID UUID NOT NULL,
    FOREIGN KEY (MovieID) REFERENCES Movies(ID),
    FOREIGN KEY (DirectorID) REFERENCES Directors(ID),
    PRIMARY KEY (MovieID, DirectorID)
);