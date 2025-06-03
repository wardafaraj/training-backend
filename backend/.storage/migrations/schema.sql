CREATE TABLE public.absenteeism_type (
    id integer NOT NULL,
    sis_id integer,
    name character varying(60) NOT NULL,
    description character varying(60),
    identifier character varying(60),
    resvd5 character varying(1),
    resvd4 character varying(1),
    resvd3 character varying(1),
    resvd2 character varying(1),
    resvd1 character varying(1),
    created_by integer,
    updated_by integer,
    deleted_by integer,
    created_at timestamp(0) with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(0) with time zone,
    deleted_at timestamp(0) with time zone
);


ALTER TABLE public.absenteeism_type OWNER TO postgres;