-- psql -d aluno -U aluno -h localhost -W  -f model/login.sql
-- CRIANDO FUNC PARA USAR gen_random_uuid
-- CREATE EXTENSION pgcrypto;
-- SELECT gen_random_uuid()
--
-- PostgreSQL database dump
--

--
-- Name: login; Type: TABLE; Schema: public; Owner: aluno
--

CREATE TABLE public.login (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    password text NOT NULL,
    email text NOT NULL,
    name text NOT NULL,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    CONSTRAINT "login_pkey" PRIMARY KEY (id),
    CONSTRAINT id UNIQUE (id) 
    -- CONSTRAINT email UNIQUE (email)
)WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.login OWNER to aluno;


