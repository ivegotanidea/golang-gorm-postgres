--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Debian 16.4-1.pgdg120+2)
-- Dumped by pg_dump version 16.4 (Homebrew)

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
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: body_arts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.body_arts (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    alias_ru character varying(100) NOT NULL,
    alias_en character varying(100) NOT NULL
);


--
-- Name: body_arts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.body_arts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: body_arts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.body_arts_id_seq OWNED BY public.body_arts.id;


--
-- Name: body_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.body_types (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    alias_ru character varying(30) NOT NULL,
    alias_en character varying(30) NOT NULL
);


--
-- Name: body_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.body_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: body_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.body_types_id_seq OWNED BY public.body_types.id;


--
-- Name: cities; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.cities (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    alias_ru character varying(30) NOT NULL,
    alias_en character varying(30) NOT NULL
);


--
-- Name: cities_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.cities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- Name: ethnos; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.ethnos (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    alias_ru character varying(30) NOT NULL,
    alias_en character varying(30) NOT NULL,
    sex character varying(10) NOT NULL
);


--
-- Name: ethnos_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.ethnos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: ethnos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.ethnos_id_seq OWNED BY public.ethnos.id;


--
-- Name: hair_colors; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hair_colors (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    alias_ru character varying(30) NOT NULL,
    alias_en character varying(30) NOT NULL
);


--
-- Name: hair_colors_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hair_colors_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hair_colors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hair_colors_id_seq OWNED BY public.hair_colors.id;


--
-- Name: intimate_hair_cuts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.intimate_hair_cuts (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    alias_ru character varying(30) NOT NULL,
    alias_en character varying(30) NOT NULL
);


--
-- Name: intimate_hair_cuts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.intimate_hair_cuts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: intimate_hair_cuts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.intimate_hair_cuts_id_seq OWNED BY public.intimate_hair_cuts.id;


--
-- Name: photos; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.photos (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    profile_id uuid NOT NULL,
    url character varying(255) NOT NULL,
    created_at timestamp without time zone,
    disabled boolean DEFAULT false,
    approved boolean DEFAULT false,
    deleted boolean DEFAULT false
);


--
-- Name: profile_body_arts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.profile_body_arts (
    profile_id uuid NOT NULL,
    body_art_id bigint NOT NULL
);


--
-- Name: profile_options; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.profile_options (
    profile_id uuid NOT NULL,
    profile_tag_id integer NOT NULL,
    price bigint,
    comment text
);


--
-- Name: profile_ratings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.profile_ratings (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    service_id uuid NOT NULL,
    profile_id uuid NOT NULL,
    review_text_visible boolean DEFAULT true,
    review character varying(2000),
    score bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    updated_by uuid NOT NULL
);


--
-- Name: profile_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.profile_tags (
    id bigint NOT NULL,
    name character varying(50)
);


--
-- Name: profile_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.profile_tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: profile_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.profile_tags_id_seq OWNED BY public.profile_tags.id;


--
-- Name: profiles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.profiles (
    body_type_id integer,
    ethnos_id integer,
    hair_color_id integer,
    intimate_hair_cut_id integer,
    city_id integer DEFAULT 0 NOT NULL,
    parsed_url character varying(255) DEFAULT ''::character varying NOT NULL,
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    active boolean DEFAULT true,
    phone character varying(30),
    name character varying(50),
    age bigint NOT NULL,
    height bigint NOT NULL,
    weight bigint NOT NULL,
    bust numeric,
    bio character varying(2000),
    address_latitude character varying(10),
    address_longitude character varying(10),
    price_in_house_night_ratio numeric DEFAULT 1 NOT NULL,
    price_in_house_contact bigint,
    price_in_house_hour bigint,
    prince_sauna_night_ratio numeric DEFAULT 1 NOT NULL,
    price_sauna_contact bigint,
    price_sauna_hour bigint,
    price_visit_night_ratio numeric DEFAULT 1 NOT NULL,
    price_visit_contact bigint,
    price_visit_hour bigint,
    price_car_night_ratio numeric DEFAULT 1 NOT NULL,
    price_car_contact bigint,
    price_car_hour bigint,
    contact_phone character varying(30),
    contact_wa character varying(30),
    contact_tg character varying(50),
    moderated boolean DEFAULT false,
    moderated_at timestamp without time zone,
    moderated_by uuid,
    verified boolean DEFAULT false,
    verified_at timestamp without time zone,
    verified_by uuid,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    updated_by uuid NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: rated_profile_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.rated_profile_tags (
    rating_id uuid NOT NULL,
    profile_tag_id integer NOT NULL,
    type character varying(10)
);


--
-- Name: rated_user_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.rated_user_tags (
    rating_id uuid NOT NULL,
    user_tag_id integer NOT NULL,
    type character varying(10)
);


--
-- Name: services; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.services (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    client_user_id uuid NOT NULL,
    client_user_rating_id uuid,
    client_user_lat character varying(10),
    client_user_lon character varying(10),
    profile_id uuid NOT NULL,
    profile_owner_id uuid NOT NULL,
    profile_rating_id uuid,
    profile_user_lat character varying(10),
    profile_user_lon character varying(10),
    distance_between_users numeric,
    trusted_distance boolean,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    updated_by uuid NOT NULL
);


--
-- Name: user_ratings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_ratings (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    service_id uuid NOT NULL,
    user_id uuid NOT NULL,
    review_text_visible boolean DEFAULT true,
    review character varying(2000),
    score bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    updated_by uuid NOT NULL
);


--
-- Name: user_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_tags (
    id bigint NOT NULL,
    name character varying(255)
);


--
-- Name: user_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_tags_id_seq OWNED BY public.user_tags.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(20) NOT NULL,
    phone character varying(30),
    telegram_user_id bigint NOT NULL,
    password character varying(255) NOT NULL,
    active boolean DEFAULT true,
    verified boolean DEFAULT false,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    avatar character varying(255),
    has_profile boolean,
    tier character varying(50) DEFAULT 'basic'::character varying NOT NULL,
    role character varying(50) DEFAULT 'user'::character varying NOT NULL
);


--
-- Name: body_arts id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.body_arts ALTER COLUMN id SET DEFAULT nextval('public.body_arts_id_seq'::regclass);


--
-- Name: body_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.body_types ALTER COLUMN id SET DEFAULT nextval('public.body_types_id_seq'::regclass);


--
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- Name: ethnos id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ethnos ALTER COLUMN id SET DEFAULT nextval('public.ethnos_id_seq'::regclass);


--
-- Name: hair_colors id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hair_colors ALTER COLUMN id SET DEFAULT nextval('public.hair_colors_id_seq'::regclass);


--
-- Name: intimate_hair_cuts id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.intimate_hair_cuts ALTER COLUMN id SET DEFAULT nextval('public.intimate_hair_cuts_id_seq'::regclass);


--
-- Name: profile_tags id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile_tags ALTER COLUMN id SET DEFAULT nextval('public.profile_tags_id_seq'::regclass);


--
-- Name: user_tags id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_tags ALTER COLUMN id SET DEFAULT nextval('public.user_tags_id_seq'::regclass);


--
-- Name: body_arts body_arts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.body_arts
    ADD CONSTRAINT body_arts_pkey PRIMARY KEY (id);


--
-- Name: body_types body_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.body_types
    ADD CONSTRAINT body_types_pkey PRIMARY KEY (id);


--
-- Name: cities cities_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);


--
-- Name: ethnos ethnos_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ethnos
    ADD CONSTRAINT ethnos_pkey PRIMARY KEY (id);


--
-- Name: hair_colors hair_colors_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hair_colors
    ADD CONSTRAINT hair_colors_pkey PRIMARY KEY (id);


--
-- Name: intimate_hair_cuts intimate_hair_cuts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.intimate_hair_cuts
    ADD CONSTRAINT intimate_hair_cuts_pkey PRIMARY KEY (id);


--
-- Name: photos photos_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT photos_pkey PRIMARY KEY (id);


--
-- Name: profile_body_arts profile_body_arts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile_body_arts
    ADD CONSTRAINT profile_body_arts_pkey PRIMARY KEY (profile_id, body_art_id);


--
-- Name: profile_options profile_options_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile_options
    ADD CONSTRAINT profile_options_pkey PRIMARY KEY (profile_id, profile_tag_id);


--
-- Name: profile_ratings profile_ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile_ratings
    ADD CONSTRAINT profile_ratings_pkey PRIMARY KEY (id);


--
-- Name: profile_tags profile_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile_tags
    ADD CONSTRAINT profile_tags_pkey PRIMARY KEY (id);


--
-- Name: profiles profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_pkey PRIMARY KEY (id);


--
-- Name: rated_profile_tags rated_profile_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rated_profile_tags
    ADD CONSTRAINT rated_profile_tags_pkey PRIMARY KEY (rating_id, profile_tag_id);


--
-- Name: rated_user_tags rated_user_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rated_user_tags
    ADD CONSTRAINT rated_user_tags_pkey PRIMARY KEY (rating_id, user_tag_id);


--
-- Name: services services_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


--
-- Name: body_arts uni_body_arts_name; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.body_arts
    ADD CONSTRAINT uni_body_arts_name UNIQUE (name);


--
-- Name: body_types uni_body_types_name; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.body_types
    ADD CONSTRAINT uni_body_types_name UNIQUE (name);


--
-- Name: cities uni_cities_name; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT uni_cities_name UNIQUE (name);


--
-- Name: ethnos uni_ethnos_name; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ethnos
    ADD CONSTRAINT uni_ethnos_name UNIQUE (name);


--
-- Name: hair_colors uni_hair_colors_name; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hair_colors
    ADD CONSTRAINT uni_hair_colors_name UNIQUE (name);


--
-- Name: intimate_hair_cuts uni_intimate_hair_cuts_name; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.intimate_hair_cuts
    ADD CONSTRAINT uni_intimate_hair_cuts_name UNIQUE (name);


--
-- Name: user_ratings user_ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_ratings
    ADD CONSTRAINT user_ratings_pkey PRIMARY KEY (id);


--
-- Name: user_tags user_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_tags
    ADD CONSTRAINT user_tags_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_profile_tags_name; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_profile_tags_name ON public.profile_tags USING btree (name);


--
-- Name: idx_profiles_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_profiles_deleted_at ON public.profiles USING btree (deleted_at);


--
-- Name: idx_users_phone; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_users_phone ON public.users USING btree (phone);


--
-- Name: idx_users_telegram_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_users_telegram_user_id ON public.users USING btree (telegram_user_id);


--
-- Name: profile_options fk_profile_options_profile_tag; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile_options
    ADD CONSTRAINT fk_profile_options_profile_tag FOREIGN KEY (profile_tag_id) REFERENCES public.profile_tags(id);


--
-- Name: rated_profile_tags fk_profile_ratings_rated_profile_tags; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rated_profile_tags
    ADD CONSTRAINT fk_profile_ratings_rated_profile_tags FOREIGN KEY (rating_id) REFERENCES public.profile_ratings(id);


--
-- Name: profile_body_arts fk_profiles_body_arts; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile_body_arts
    ADD CONSTRAINT fk_profiles_body_arts FOREIGN KEY (profile_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- Name: photos fk_profiles_photos; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT fk_profiles_photos FOREIGN KEY (profile_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- Name: profile_options fk_profiles_profile_options; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profile_options
    ADD CONSTRAINT fk_profiles_profile_options FOREIGN KEY (profile_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- Name: services fk_profiles_services; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT fk_profiles_services FOREIGN KEY (profile_id) REFERENCES public.profiles(id);


--
-- Name: rated_profile_tags fk_rated_profile_tags_profile_tag; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rated_profile_tags
    ADD CONSTRAINT fk_rated_profile_tags_profile_tag FOREIGN KEY (profile_tag_id) REFERENCES public.profile_tags(id);


--
-- Name: rated_user_tags fk_rated_user_tags_user_tag; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rated_user_tags
    ADD CONSTRAINT fk_rated_user_tags_user_tag FOREIGN KEY (user_tag_id) REFERENCES public.user_tags(id);


--
-- Name: services fk_services_client_user_rating; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT fk_services_client_user_rating FOREIGN KEY (client_user_rating_id) REFERENCES public.user_ratings(id);


--
-- Name: services fk_services_profile_rating; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT fk_services_profile_rating FOREIGN KEY (profile_rating_id) REFERENCES public.profile_ratings(id);


--
-- Name: rated_user_tags fk_user_ratings_rated_user_tags; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rated_user_tags
    ADD CONSTRAINT fk_user_ratings_rated_user_tags FOREIGN KEY (rating_id) REFERENCES public.user_ratings(id);


--
-- Name: profiles fk_users_profiles; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT fk_users_profiles FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: services fk_users_services; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT fk_users_services FOREIGN KEY (client_user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

