--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.2 (Debian 17.2-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: get_reservations_by_month(date); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_reservations_by_month(month date) RETURNS TABLE(uuid uuid, date date, termin integer, washer integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT
        r.uuid,
        r.date,
        r.termin,
        r.washer
    FROM
        reservations r
    WHERE
        r.date >= DATE_TRUNC('month', month)
        AND r.date < DATE_TRUNC('month', month) + INTERVAL '1 month';
END;
$$;


ALTER FUNCTION public.get_reservations_by_month(month date) OWNER TO postgres;

--
-- Name: get_user_reservations(uuid); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_user_reservations(user_uuid uuid) RETURNS TABLE(uuid uuid, date date, termin integer, washer integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT
        r.uuid,
        r.date,
        r.termin,
        r.washer
    FROM
        reservations r
    WHERE
        r.uuid = user_uuid;
END;
$$;


ALTER FUNCTION public.get_user_reservations(user_uuid uuid) OWNER TO postgres;

--
-- Name: search_users(text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.search_users(search_text text) RETURNS TABLE(uuid uuid, phone character varying, name character varying, surname character varying, email character varying, role character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT
        u.uuid,
        u.phone,
        u.name,
        u.surname,
        u.email,
        u.role
    FROM
        users u
    WHERE
        LOWER(u.name) LIKE LOWER('%' || search_text || '%')
        OR LOWER(u.surname) LIKE LOWER('%' || search_text || '%')
        OR LOWER(u.phone) LIKE LOWER('%' || search_text || '%')
        OR LOWER(u.email) LIKE LOWER('%' || search_text || '%');
END;
$$;


ALTER FUNCTION public.search_users(search_text text) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: auth; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth (
    id integer NOT NULL,
    user_uuid uuid NOT NULL,
    password_hash text NOT NULL,
    last_login timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.auth OWNER TO postgres;

--
-- Name: auth_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auth_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.auth_id_seq OWNER TO postgres;

--
-- Name: auth_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auth_id_seq OWNED BY public.auth.id;


--
-- Name: reservations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reservations (
    id integer NOT NULL,
    uuid uuid NOT NULL,
    date date NOT NULL,
    termin integer NOT NULL,
    washer integer NOT NULL
);


ALTER TABLE public.reservations OWNER TO postgres;

--
-- Name: reservations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reservations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.reservations_id_seq OWNER TO postgres;

--
-- Name: reservations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reservations_id_seq OWNED BY public.reservations.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    uuid uuid NOT NULL,
    phone character varying(15) NOT NULL,
    name character varying(100) NOT NULL,
    surname character varying(100) NOT NULL,
    room integer NOT NULL,
    email character varying(255) NOT NULL,
    disabled boolean DEFAULT false NOT NULL,
    confirmed boolean DEFAULT false NOT NULL,
    role character varying(10) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_role_check CHECK (((role)::text = ANY ((ARRAY['admin'::character varying, 'user'::character varying])::text[])))
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: auth id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth ALTER COLUMN id SET DEFAULT nextval('public.auth_id_seq'::regclass);


--
-- Name: reservations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations ALTER COLUMN id SET DEFAULT nextval('public.reservations_id_seq'::regclass);


--
-- Name: auth auth_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth
    ADD CONSTRAINT auth_pkey PRIMARY KEY (id);


--
-- Name: reservations reservations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_pkey PRIMARY KEY (id);


--
-- Name: reservations reservations_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_unique UNIQUE (uuid, date, termin, washer);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_phone_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_phone_key UNIQUE (phone);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (uuid);


--
-- Name: auth auth_user_uuid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth
    ADD CONSTRAINT auth_user_uuid_fkey FOREIGN KEY (user_uuid) REFERENCES public.users(uuid) ON DELETE CASCADE;


--
-- Name: reservations reservations_uuid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_uuid_fkey FOREIGN KEY (uuid) REFERENCES public.users(uuid) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

