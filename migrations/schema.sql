--
-- PostgreSQL database dump
--

-- Dumped from database version 11.2 (Debian 11.2-1.pgdg90+1)
-- Dumped by pg_dump version 11.3 (Debian 11.3-1.pgdg90+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: postgis; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;


--
-- Name: EXTENSION postgis; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION postgis IS 'PostGIS geometry, geography, and raster spatial types and functions';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: hazard_types; Type: TABLE; Schema: public; Owner: schutzstreifen
--

CREATE TABLE public.hazard_types (
    id uuid NOT NULL,
    label character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.hazard_types OWNER TO schutzstreifen;

--
-- Name: hazards; Type: TABLE; Schema: public; Owner: schutzstreifen
--

CREATE TABLE public.hazards (
    id uuid NOT NULL,
    label character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    lat numeric NOT NULL,
    lon numeric NOT NULL,
    visible boolean DEFAULT true,
    user_id uuid NOT NULL,
    hazard_type_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.hazards OWNER TO schutzstreifen;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: schutzstreifen
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO schutzstreifen;

--
-- Name: users; Type: TABLE; Schema: public; Owner: schutzstreifen
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO schutzstreifen;

--
-- Name: hazard_types hazard_types_pkey; Type: CONSTRAINT; Schema: public; Owner: schutzstreifen
--

ALTER TABLE ONLY public.hazard_types
    ADD CONSTRAINT hazard_types_pkey PRIMARY KEY (id);


--
-- Name: hazards hazards_pkey; Type: CONSTRAINT; Schema: public; Owner: schutzstreifen
--

ALTER TABLE ONLY public.hazards
    ADD CONSTRAINT hazards_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: schutzstreifen
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_name_key; Type: CONSTRAINT; Schema: public; Owner: schutzstreifen
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_name_key UNIQUE (name);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: schutzstreifen
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: schutzstreifen
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: hazards hazards_hazard_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: schutzstreifen
--

ALTER TABLE ONLY public.hazards
    ADD CONSTRAINT hazards_hazard_type_id_fkey FOREIGN KEY (hazard_type_id) REFERENCES public.hazard_types(id);


--
-- Name: hazards hazards_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: schutzstreifen
--

ALTER TABLE ONLY public.hazards
    ADD CONSTRAINT hazards_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

