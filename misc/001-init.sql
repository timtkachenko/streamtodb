GRANT ALL PRIVILEGES ON DATABASE streamtodb TO streamtodb;

ALTER DATABASE "streamtodb" OWNER TO streamtodb;

CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION pgcrypto;

create table ports
(
    id   uuid default uuid_generate_v4() primary key,
    name text,
    codename    text,
    city        text,
    country     text,
    alias       text[],
    regions     text[],
    coordinates decimal[],
    province    text ,
    timezone    text,
    unlocs      text[],
    code        text,
    created_at timestamp with time zone default now()              not null,
    updated_at timestamp with time zone default now()              not null,
    deleted_at timestamp with time zone
);

create unique index ports_codename_uindex  on ports (codename);

create index ports_created_at_id_index on ports (created_at asc, id asc);

