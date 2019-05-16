-- Table: public.empresa

-- DROP TABLE public.empresa;

CREATE TABLE public.empresa
(
    codemp serial NOT NULL,
    cnpj character varying(14) COLLATE pg_catalog."default" NOT NULL,
    razaosocial character varying(80) COLLATE pg_catalog."default" NOT NULL,
    ativo boolean,
    CONSTRAINT "empresa_pkey" PRIMARY KEY (codemp),
    CONSTRAINT cnpj UNIQUE (cnpj)

)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.empresa OWNER to aluno;