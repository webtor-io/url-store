CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE public.url
(
    url text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    last_accessed_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT url_pkey PRIMARY KEY (url)
);

ALTER TABLE public.url
    OWNER to url_store;

CREATE INDEX url_idx ON url (digest(url, 'sha1'));