--
-- PostgreSQL database dump
--

-- Dumped from database version 12.4 (Debian 12.4-1.pgdg100+1)
-- Dumped by pg_dump version 12.4 (Debian 12.4-1.pgdg100+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: refresh_tokens; Type: TABLE; Schema: public; Owner: cash
--

CREATE TABLE public.refresh_tokens (
    token_string text NOT NULL,
    user_id bigint,
    created_at timestamp with time zone,
    expires_at timestamp with time zone
);


ALTER TABLE public.refresh_tokens OWNER TO cash;

--
-- Name: users; Type: TABLE; Schema: public; Owner: cash
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username text NOT NULL,
    password text NOT NULL,
    role text NOT NULL
);


ALTER TABLE public.users OWNER TO cash;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: cash
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO cash;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cash
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: cash
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: refresh_tokens; Type: TABLE DATA; Schema: public; Owner: cash
--

COPY public.refresh_tokens (token_string, user_id, created_at, expires_at) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: cash
--

COPY public.users (id, created_at, updated_at, deleted_at, username, password, role) FROM stdin;
1	2020-09-02 21:37:03.109887+00	2020-09-02 21:37:03.109887+00	\N	kumar	$2a$10$LwOrK..EVrc4vLyVBcdQdOMByHeD1jb852zGlT4P7swOUfFic3HOu	superuser
2	2020-09-02 21:37:08.786216+00	2020-09-02 21:37:08.786216+00	\N	rajesh	$2a$10$YfmWQJU95VY5AniwP9rLqexYNN5E4.TX6.Sn5bi7sAgBiW3OfR7AW	superuser
3	2020-09-02 21:42:03.031101+00	2020-09-02 21:42:03.031101+00	\N	admin-test	$2a$10$7BeqCdvO2W6F3RAQF0FXAuHcswLv8uGJfI7l7pzknZDkOw4KGijKa	admin
4	2020-09-02 21:42:06.902875+00	2020-09-02 21:42:06.902875+00	\N	carrier-test	$2a$10$Nz9ugTuM1dzUIKRqkRjZ8O8z/5MuVtDJ9Wslu.G2xATJDD2Y192oG	carrier
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cash
--

SELECT pg_catalog.setval('public.users_id_seq', 4, true);


--
-- Name: refresh_tokens refresh_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: cash
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_pkey PRIMARY KEY (token_string);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: cash
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: cash
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: cash
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: refresh_tokens fk_users_refresh_tokens; Type: FK CONSTRAINT; Schema: public; Owner: cash
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT fk_users_refresh_tokens FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

