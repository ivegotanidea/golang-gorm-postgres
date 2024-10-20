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
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: body_arts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.body_arts (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    alias_ru character varying(100) NOT NULL,
    alias_en character varying(100) NOT NULL
);


ALTER TABLE public.body_arts OWNER TO postgres;

--
-- Name: body_arts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.body_arts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.body_arts_id_seq OWNER TO postgres;

--
-- Name: body_arts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.body_arts_id_seq OWNED BY public.body_arts.id;


--
-- Name: body_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.body_types (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    alias_ru character varying(30) NOT NULL,
    alias_en character varying(30) NOT NULL
);


ALTER TABLE public.body_types OWNER TO postgres;

--
-- Name: body_types_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.body_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.body_types_id_seq OWNER TO postgres;

--
-- Name: body_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.body_types_id_seq OWNED BY public.body_types.id;


--
-- Name: cities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cities (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    alias_ru character varying(30) NOT NULL,
    alias_en character varying(30) NOT NULL
);


ALTER TABLE public.cities OWNER TO postgres;

--
-- Name: cities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cities_id_seq OWNER TO postgres;

--
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- Name: ethnos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ethnos (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    alias_ru character varying(30) NOT NULL,
    alias_en character varying(30) NOT NULL,
    sex character varying(10) NOT NULL
);


ALTER TABLE public.ethnos OWNER TO postgres;

--
-- Name: ethnos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ethnos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ethnos_id_seq OWNER TO postgres;

--
-- Name: ethnos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ethnos_id_seq OWNED BY public.ethnos.id;


--
-- Name: hair_colors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.hair_colors (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    alias_ru character varying(30) NOT NULL,
    alias_en character varying(30) NOT NULL
);


ALTER TABLE public.hair_colors OWNER TO postgres;

--
-- Name: hair_colors_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.hair_colors_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.hair_colors_id_seq OWNER TO postgres;

--
-- Name: hair_colors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.hair_colors_id_seq OWNED BY public.hair_colors.id;


--
-- Name: intimate_hair_cuts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.intimate_hair_cuts (
    id bigint NOT NULL,
    name character varying(30) NOT NULL,
    alias_ru character varying(30) NOT NULL,
    alias_en character varying(30) NOT NULL
);


ALTER TABLE public.intimate_hair_cuts OWNER TO postgres;

--
-- Name: intimate_hair_cuts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.intimate_hair_cuts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.intimate_hair_cuts_id_seq OWNER TO postgres;

--
-- Name: intimate_hair_cuts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.intimate_hair_cuts_id_seq OWNED BY public.intimate_hair_cuts.id;


--
-- Name: photos; Type: TABLE; Schema: public; Owner: postgres
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


ALTER TABLE public.photos OWNER TO postgres;

--
-- Name: profile_body_arts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.profile_body_arts (
    profile_id uuid NOT NULL,
    body_art_id bigint NOT NULL
);


ALTER TABLE public.profile_body_arts OWNER TO postgres;

--
-- Name: profile_options; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.profile_options (
    profile_id uuid NOT NULL,
    profile_tag_id integer NOT NULL,
    price bigint,
    comment text
);


ALTER TABLE public.profile_options OWNER TO postgres;

--
-- Name: profile_ratings; Type: TABLE; Schema: public; Owner: postgres
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


ALTER TABLE public.profile_ratings OWNER TO postgres;

--
-- Name: profile_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.profile_tags (
    id bigint NOT NULL,
    name character varying(50)
);


ALTER TABLE public.profile_tags OWNER TO postgres;

--
-- Name: profile_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.profile_tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.profile_tags_id_seq OWNER TO postgres;

--
-- Name: profile_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.profile_tags_id_seq OWNED BY public.profile_tags.id;


--
-- Name: profiles; Type: TABLE; Schema: public; Owner: postgres
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


ALTER TABLE public.profiles OWNER TO postgres;

--
-- Name: rated_profile_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.rated_profile_tags (
    rating_id uuid NOT NULL,
    profile_tag_id integer NOT NULL,
    type character varying(10)
);


ALTER TABLE public.rated_profile_tags OWNER TO postgres;

--
-- Name: rated_user_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.rated_user_tags (
    rating_id uuid NOT NULL,
    user_tag_id integer NOT NULL,
    type character varying(10)
);


ALTER TABLE public.rated_user_tags OWNER TO postgres;

--
-- Name: services; Type: TABLE; Schema: public; Owner: postgres
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


ALTER TABLE public.services OWNER TO postgres;

--
-- Name: user_ratings; Type: TABLE; Schema: public; Owner: postgres
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


ALTER TABLE public.user_ratings OWNER TO postgres;

--
-- Name: user_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_tags (
    id bigint NOT NULL,
    name character varying(255)
);


ALTER TABLE public.user_tags OWNER TO postgres;

--
-- Name: user_tags_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_tags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_tags_id_seq OWNER TO postgres;

--
-- Name: user_tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_tags_id_seq OWNED BY public.user_tags.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
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


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: body_arts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.body_arts ALTER COLUMN id SET DEFAULT nextval('public.body_arts_id_seq'::regclass);


--
-- Name: body_types id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.body_types ALTER COLUMN id SET DEFAULT nextval('public.body_types_id_seq'::regclass);


--
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- Name: ethnos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ethnos ALTER COLUMN id SET DEFAULT nextval('public.ethnos_id_seq'::regclass);


--
-- Name: hair_colors id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.hair_colors ALTER COLUMN id SET DEFAULT nextval('public.hair_colors_id_seq'::regclass);


--
-- Name: intimate_hair_cuts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.intimate_hair_cuts ALTER COLUMN id SET DEFAULT nextval('public.intimate_hair_cuts_id_seq'::regclass);


--
-- Name: profile_tags id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile_tags ALTER COLUMN id SET DEFAULT nextval('public.profile_tags_id_seq'::regclass);


--
-- Name: user_tags id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_tags ALTER COLUMN id SET DEFAULT nextval('public.user_tags_id_seq'::regclass);


--
-- Data for Name: body_arts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.body_arts (id, name, alias_ru, alias_en) FROM stdin;
1	tatu	Татуировки	Tattoos
2	silikon_v_grudi	Силикон в груди	Breast Implants
3	pirsing	Пирсинг	Piercing
\.


--
-- Data for Name: body_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.body_types (id, name, alias_ru, alias_en) FROM stdin;
1	hudaya	Худая	Slim
2	stroynaya	Стройная	Fit
3	sportivnaya	Спортивная	Athletic
4	polnaya	Полная	Full-figured
\.


--
-- Data for Name: cities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cities (id, name, alias_ru, alias_en) FROM stdin;
1	almaty	Алматы	Almaty
2	ust-kamenogorsk	Усть-Каменогорск	Ust-Kamenogorsk
3	zhezkazgan	Жезказган	Zhezkazgan
4	zhetysai	Жетысай	Zhetysai
5	lisakovsk	Лисаковск	Lisakovsk
6	astana	Астана	Astana
7	kostanay	Костанай	Kostanay
8	kapchagay	Капчагай	Kapchagay
9	ridder	Риддер	Ridder
10	shu	Шу	Shu
11	shymkent	Шымкент	Shymkent
12	kyzylorda	Кызылорда	Kyzylorda
13	balhash	Балхаш	Balkhash
14	kaskelen	Каскелен	Kaskelen
15	shahtinsk	Шахтинск	Shahtinsk
16	karaganda	Караганда	Karaganda
17	kokshetau	Кокшетау	Kokshetau
18	aksay	Аксай	Aksay
19	kulsary	Кульсары	Kulsary
20	yesik	Есик	Yesik
21	aktau	Актау	Aktau
22	taldykorgan	Талдыкорган	Taldykorgan
23	shchuchinsk	Щучинск	Shchuchinsk
24	stepnogorsk	Степногорск	Stepnogorsk
25	zharkent	Жаркент	Zharkent
26	aktobe	Актобе	Aktobe
27	turkestan	Туркестан	Turkestan
28	rudny	Рудный	Rudny
29	talgar	Талгар	Talgar
30	shardara	Шардара	Shardara
31	atyrau	Атырау	Atyrau
32	semey	Семей	Semey
33	zhanaozen	Жанаозен	Zhanaozen
34	saran	Сарань	Saran
35	atbasar	Атбасар	Atbasar
36	taraz	Тараз	Taraz
37	petropavl	Петропавловск	Petropavl
38	satpayev	Сатпаев	Satpayev
39	aksu	Аксу	Aksu
40	tekeli	Текели	Tekeli
41	uralsk	Уральск	Uralsk
42	temirtau	Темиртау	Temirtau
43	kentau	Кентау	Kentau
44	zyryanovsk	Зыряновск	Zyryanovsk
45	mangistau	Мангистау	Mangistau
46	pavlodar	Павлодар	Pavlodar
47	ekibastuz	Экибастуз	Ekibastuz
48	saryagash	Сарыагаш	Saryagash
\.


--
-- Data for Name: ethnos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ethnos (id, name, alias_ru, alias_en, sex) FROM stdin;
1	metiska	Метиска	Métis	female
2	chuvashka	Чувашка	Chuvash	female
3	kirgizka	Киргизка	Kyrgyz	female
4	azerbaijanka	Азербайджанка	Azerbaijani	female
5	iranka	Иранка	Iranian	female
6	taika	Тайка	Thai	female
7	ukrainka	Украинка	Ukrainian	female
8	litovka	Литовка	Lithuanian	female
9	ingushka	Ингушка	Ingush	female
10	dagestanka	Дагестанка	Dagestani	female
11	dunganka	Дунганка	Dungan	female
12	osetinka	Осетинка	Ossetian	female
13	turkmenka	Туркменка	Turkmen	female
14	mulatka	Мулатка	Mulatto	female
15	evropeyka	Европейка	European	female
16	koreyanka	Кореянка	Korean	female
17	beloruska	Белоруска	Belarusian	female
18	chechenka	Чеченка	Chechen	female
19	tadzhichka	Таджичка	Tajik	female
20	kavkazka	Кавказка	Caucasian	female
21	slavyanka	Славянка	Slavic	female
22	turchanka	Турчанка	Turkish	female
23	evreyka	Еврейка	Jewish	female
24	nemka	Немка	German	female
25	kazashka	Казашка	Kazakh	female
26	frantsuzhenka	Француженка	French	female
27	latyshka	Латышка	Latvian	female
28	gruzinka	Грузинка	Georgian	female
29	moldavanka	Молдаванка	Moldovan	female
30	bolgarka	Болгарка	Bulgarian	female
31	bashkirka	Башкирка	Bashkir	female
32	rumynka	Румынка	Romanian	female
33	grechanka	Гречанка	Greek	female
34	uzbechka	Узбечка	Uzbek	female
35	ispanka	Испанка	Spanish	female
36	tatarka	Татарка	Tatar	female
37	yakutka	Якутка	Yakut	female
38	aziatka	Азиатка	Asian	female
39	mordvinka	Мордвинка	Mordvin	female
40	kitayanka	Китаянка	Chinese	female
41	tsyganka	Цыганка	Gypsy	female
42	armyanka	Армянка	Armenian	female
43	italyanka	Итальянка	Italian	female
44	uygurka	Уйгурка	Uyghur	female
45	polyachka	Полячка	Polish	female
46	arabka	Арабка	Arab	female
47	dagestanets	Дагестанец	Dagestani	male
48	slavyanin	Славянин	Slavic	male
49	bolgarin	Болгарин	Bulgarian	male
50	kavkazets	Кавказец	Caucasian	male
51	ingush	Ингуш	Ingush	male
52	osetinets	Осетинец	Ossetian	male
53	armyanin	Армянин	Armenian	male
54	kazakh	Казах	Kazakh	male
55	ukrainets	Украинец	Ukrainian	male
56	tsyganin	Цыганин	Gypsy	male
57	gruzin	Грузин	Georgian	male
58	italyanets	Итальянец	Italian	male
59	evropeets	Европеец	European	male
60	litovets	Литовец	Lithuanian	male
61	tadzhik	Таджик	Tajik	male
62	frantsuz	Француз	French	male
63	rumyn	Румын	Romanian	male
64	ispanets	Испанец	Spanish	male
65	polyak	Поляк	Polish	male
66	chuvash	Чуваш	Chuvash	male
67	turkmen	Туркмен	Turkmen	male
68	moldavanin	Молдаванин	Moldovan	male
69	kurd	Курд	Kurd	male
70	evrey	Еврей	Jewish	male
71	chechenets	Чеченец	Chechen	male
72	bashkir	Башкир	Bashkir	male
73	metis	Метис	Métis	male
74	nemets	Немец	German	male
75	mulat	Мулат	Mulatto	male
76	arab	Араб	Arab	male
77	latysh	Латыш	Latvian	male
78	russkiy	Русский	Russian	male
79	belorus	Белорус	Belarusian	male
80	dungan	Дунган	Dungan	male
81	grek	Грек	Greek	male
82	yakut	Якут	Yakut	male
83	koreets	Кореец	Korean	male
84	uygur	Уйгур	Uyghur	male
85	tatarin	Татарин	Tatar	male
86	turok	Турок	Turkish	male
87	kitayets	Китаец	Chinese	male
88	mordvin	Мордвин	Mordvin	male
89	iranets	Иранец	Iranian	male
90	azerbaidzhanets	Азербайджанец	Azerbaijani	male
91	uzbek	Узбек	Uzbek	male
92	aziat	Азиат	Asian	male
93	kirgiz	Киргиз	Kyrgyz	male
\.


--
-- Data for Name: hair_colors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.hair_colors (id, name, alias_ru, alias_en) FROM stdin;
1	brunetka	Брюнетка	Brunette
2	shatenka	Шатенка	Brown-haired
3	ryzhaya	Рыжая	Red-haired
4	rusaya	Русая	Light brown
5	blondinka	Блондинка	Blonde
6	lysaya	Лысая	Bald
7	tsvetnaya	Цветная	Colored
\.


--
-- Data for Name: intimate_hair_cuts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.intimate_hair_cuts (id, name, alias_ru, alias_en) FROM stdin;
1	polnaya_depilyatsiya	Полная депиляция	Full depilation
2	akkuratnaya_strizhka	Аккуратная стрижка	Neat trim
3	naturalnaya	Натуральная	Natural
\.


--
-- Data for Name: photos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.photos (id, profile_id, url, created_at, disabled, approved, deleted) FROM stdin;
b5904adf-563e-4ce3-9ee9-35e8ed626310	c5b40786-9fae-4599-9b8e-72e5262d9356	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 16:30:39.61429	f	f	f
7f2ed165-c73f-40bb-ab49-cbafbed16480	ca8bd72a-ddb7-4365-bd85-d76733a922c4	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 16:30:39.799304	f	f	f
e5258080-ae21-48dc-a223-a81205492b81	4ab29639-f7a5-498e-899c-a8590c493947	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 16:30:40.066329	f	f	f
8155a15a-299e-4d46-9a81-63b857b875cc	e82afe50-a1e3-40b9-9cfd-953c82e88539	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 16:30:40.241345	f	f	f
bf963f18-1501-4ed3-a7bc-fecf0023081a	946b8cca-e1fd-4bb2-9a00-126e4b423eaf	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 16:30:40.498521	f	f	f
1bae2087-5afa-46d2-b62d-573d47d4a14e	82194403-cdaf-42f6-8d11-129cd6ed716d	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 16:30:40.666361	f	f	f
3677423d-d50a-4c56-9e55-a82b4ac5f056	016bbc0f-968e-4cb6-969f-507ebc87d449	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 16:30:40.92198	f	f	f
6a501a6f-a547-47cc-b0b5-6bc93366af65	87f8f0d8-7619-4487-b8b4-cba515d4e35a	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 16:30:41.091513	f	f	f
724f1f67-68fb-4fa6-857c-27c6bb61290c	e604f1c9-6a84-411e-bd8f-3ad43a6d5250	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:36.651595	f	f	f
2517e869-7104-404a-9b45-cecc44d03000	cc3e11e9-36fb-4412-abc6-a29f9ce8a4ae	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:36.88155	f	f	f
f98f9ad3-dcde-47fe-8532-5ccd475c1653	c1b985b5-bd3d-4c97-b199-f39d34d418da	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:37.037675	f	f	f
4f335d17-9e50-4fe1-98c2-6be2b10d6c18	0c44fd8e-a13f-4343-ab79-35d708812cde	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:37.192235	f	f	f
023af447-d113-464f-9650-cc2bf6a91d48	8cb74105-7b0b-48f8-99ea-c9ad193fa431	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:37.3443	f	f	f
e8ad9a46-aed3-4690-8e25-8fcf2e29e1fb	e0b731db-64dd-403b-b16e-40093c53dc26	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:37.496619	f	f	f
b41b077b-8bc0-41a5-98aa-ae5db01252d8	5ce524d1-e85c-4b70-8989-9e5a165e60c3	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:37.646536	f	f	f
363a2844-fcae-4657-8699-c40d2ea3fc60	29be0a43-20cd-47e7-b7ed-e14c28289552	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:37.794922	f	f	f
0de3d1fe-3da5-462c-9749-bc3dd5459b49	00ddbdec-c949-4d6a-94f1-0561eddf2501	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:37.943297	f	f	f
9b16e388-dcc5-438e-8997-696484fc6648	c4745658-ad6e-440e-82e2-4238991f0dbe	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:37.945645	f	f	f
0afeff7d-1b71-4386-b9b5-15943100783c	6e0b0a26-34cd-43ac-a4c0-1ef0c93fc477	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.09469	f	f	f
d9af10b8-7716-4ca0-bbfb-cacc1c8ef68e	2b6e322d-e0ee-4bd6-b2b6-a293dd0465fe	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.097104	f	f	f
ee70c4b9-7ef0-498a-8e23-b960bb5127c4	355e198b-356f-4f43-a655-1a4d22bfaa7e	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.246295	f	f	f
5c94bc91-0023-4b0f-8f0e-a1fbfa869213	e89d7b36-2356-4592-ba0b-05f8e72486e6	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.248759	f	f	f
f852f455-7df6-4e01-8dda-f0d7ebf6e3f7	dccdbeb1-1809-417a-b3ac-9c70d54d1f88	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.397539	f	f	f
d5ee72a4-8cba-40d9-ae2c-17655378fc2b	b9545a4e-f365-448f-a263-f7d3222c6d70	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.406541	f	f	f
98b573b1-0c35-4a6d-93c6-ff29fcc11690	dcc5f238-ef17-445d-9196-491e2094d58a	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.408907	f	f	f
a5f01c83-bd0a-48f3-bfc7-e4d0643d76e0	789aeeab-b85d-460c-919e-e6f9514ad182	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.494069	f	f	f
080e8e73-59d9-46c7-bb91-10aff88d7043	48a4fdb9-6af7-4dd2-8091-06df9c18f92a	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.644985	f	f	f
92c192a4-e742-494e-862e-c3c07cd39e28	d54284e8-3db3-47d6-a900-4e115b51bdc8	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.722968	f	f	f
b0026d1e-8ba0-433a-b7d9-b3595f5ecef6	5e23d425-7d68-46f3-93e9-04309edc7da8	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.807823	f	f	f
3de0d8d0-547e-4a51-9c49-97fbe0c98cf2	bb05a39c-eae3-42e7-ba39-5e12aae5be8f	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.966042	f	f	f
bb676bbe-4d3e-47b4-99e9-528ca368758a	06c19758-852a-4192-afca-b632fbde6618	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.968654	f	f	f
5d138321-ea24-4367-8b42-d4bf81d5dc50	75c56ebc-7332-41ba-8595-00177621fb7a	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:38.970684	f	f	f
f8c054aa-e925-47e6-bfea-c19063211983	6712ac3e-8acf-4b96-a382-69c86f2a7de0	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:40.584316	f	f	f
6d12f22c-87d7-4552-bc3b-cfef52b4ee8a	ddcb59d6-9a82-4ab6-918e-dce2097d4603	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:40.767891	f	f	f
19dfebd2-c7d1-4423-bd4e-89723caad805	699a5205-affc-404a-8390-3bf8b6d1b8a0	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:41.0196	f	f	f
c4a64aeb-4a41-4273-8a67-5a7f7e6fa23a	add0bf9a-b9f9-4203-8106-56ef6c2ec13e	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:41.1782	f	f	f
b88d7aeb-3a98-48dc-8d18-02c60d60337b	04ad1331-165d-4cab-bade-148029bc121e	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:41.419673	f	f	f
b35542c3-f0ed-40f0-ade0-51060a01fce6	7961adc7-2bf6-4af8-abbb-25086b4791f6	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:41.57253	f	f	f
6d5822e8-d387-4b5c-b63c-e9da002b21c3	3dd222e0-da73-4f9a-be39-bf713d8085ea	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:41.811781	f	f	f
66f3ce6b-46ea-4223-a790-b70c09f4f2bc	9f5f226e-60e3-44d1-a4dd-a7ed8fc3aed6	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:41.970385	f	f	f
bdd9a273-def6-4988-ae28-7c5d2ce48a68	65b73e18-4751-4c67-b4e9-a966779ba9ea	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:43.437481	f	f	f
226eb857-5954-43f7-a023-dd39ef234270	d35e126d-0e12-460c-a820-46730076a2cb	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:43.582396	f	f	f
643a3be6-8a33-49d2-aab8-49f788dd61dc	dfdd7f7a-7fbb-4e22-83b9-ea031dba83d8	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:43.753313	f	f	f
9acbf8a6-cfac-4f64-9ee1-174da11fb17e	fe398ea0-93c8-4cc6-a1ab-3c39f11549a8	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:43.908149	f	f	f
92f6ecca-6ebb-4368-b0d4-4e4e175ad9c7	38e062b9-e847-4e78-86d3-50ede8cd63e4	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:44.160114	f	f	f
9ffe58d4-141a-4dee-ad59-b9a4ed87a807	d10c5aca-e410-45b4-9442-0ce5b6f9ce46	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:44.389759	f	f	f
93ee3c90-deb5-4271-af5d-747a17b8d7b5	427578f5-ca68-4609-8531-7893d8c6b77f	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:44.619634	f	f	f
d9821f69-104d-4afa-b83a-0f5410aba79f	c80739e5-9f0c-47b6-b81c-f231b5d65607	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:44.850374	f	f	f
c8c6ba83-32d5-4982-ac89-c16b5daae7c6	4d49b906-00d7-43b8-857f-1e2cccc742a4	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:45.080426	f	f	f
97e54204-dd24-4335-85ef-e6e65487a0c2	bc88f881-93d5-47ba-92b4-02a5cf706bf1	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:50:45.311026	f	f	f
e734991a-4078-45cc-b23b-0a1f15f455ea	d0d9d46b-43e2-4a96-8cfd-36884e158373	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:17.548935	f	f	f
a039e930-f1bb-4f95-ae21-36aff686cd36	8c838cee-7e37-4b09-b5d4-68978e9c786d	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:17.785856	f	f	f
7e79aff2-4038-4ff3-8367-aaef07b0a126	ffd12f80-86c3-4719-b1c8-d58e39bf5193	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:17.947954	f	f	f
7383ccbf-ad5c-4017-b81c-3b9714906013	378f337e-2ad5-4e96-8c49-fd5168cbc973	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:18.107502	f	f	f
de20e726-2035-49b0-9369-81340e3e1c51	51e7bc12-66e4-468a-8e7d-116986eaa42c	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:18.263683	f	f	f
7ef789fc-d0ab-42d9-96a0-ee717a104c59	4c3182e9-278b-49f3-a293-8c833901a79c	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:18.422676	f	f	f
c432d272-cd49-41de-8696-046d8c5adf65	a1e85a83-feda-4d01-924b-283708c918d3	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:18.574535	f	f	f
bbc544af-492b-4b7c-ad64-2ce3dc7ae962	9ad81c52-2114-413a-963d-ef6bc31e0514	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:18.726456	f	f	f
16fa5381-bdd3-4b7f-ae8f-c925dfa61a07	115d09d4-39e6-4250-9b22-f9ac0cd4ae37	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:18.875994	f	f	f
5f78a4df-3b47-4a86-8499-7b80b471a842	76c7e0a7-aea8-4955-8d1b-f1b59992069f	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:18.878148	f	f	f
bb766f29-260a-4e9e-a502-50d204fdb682	456b4dcf-99b8-46be-a00b-5135e42b5354	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.030733	f	f	f
5cc67e7e-64ab-4cd4-8f30-4dc8cb712e45	eefe9323-4063-441c-809b-03f734328caf	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.033208	f	f	f
c643a8f0-4fc0-458a-af37-be7e3283e273	1312d36a-34f4-46c7-8dac-2fd4c304583e	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.186671	f	f	f
98ff3dfb-4b0b-4500-82eb-7530cac2b133	ab9fd6ea-7dba-42f2-92cd-6b8f99467ab1	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.188963	f	f	f
57bfe524-aa31-4204-850a-3496a4419a29	28c7af70-995a-4aa7-a108-a2ed41b85e25	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.345881	f	f	f
932fd33c-9bf0-45f0-9790-f8886e54b264	e4888c7c-2cc5-4d5f-94ce-85f2bb84ce85	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.34832	f	f	f
8b34e4fb-8e15-4580-a5ae-56d1ada33eca	1454458e-2051-4698-99c5-d66a6e4891e9	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.350547	f	f	f
c15deab4-d41f-4666-82e5-60dca3687fd3	a63fd8b1-123f-48ed-9929-4b1c976f0791	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.435172	f	f	f
1a58111b-0090-4d9b-bf3e-470d1f3737ab	382ca26d-ed90-4b01-ae13-192267e2e719	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.58531	f	f	f
cf0bff8f-dc84-4e88-b3ac-17cfeb3d8b84	73351d20-4a7e-483b-b27d-76477972ac83	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.669582	f	f	f
756c5216-c612-4bd9-a018-9027525fe5c2	6fceb6dd-628a-460d-af0a-80809ba2e8e8	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.75407	f	f	f
e6c87cd1-6518-4629-b2ce-30258c19d95c	15c32c9b-4421-4c81-9295-faaf4e12cbc4	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.911679	f	f	f
c408b087-3b1b-4e69-8074-e20b0c2c29b6	57edd4de-c007-4a1a-b0b7-d3a72eac4633	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.913942	f	f	f
1039705f-73c2-4ccf-a9d2-a6ebcaac8025	0601d0b3-4cfc-4372-b63d-b41fa012b1ff	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:19.916666	f	f	f
4d338b72-b73a-439a-be81-81bf5428f9ca	ac0e613a-0f21-4648-a4f7-907cbe5c931d	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:21.540969	f	f	f
eeb6d0c4-8627-476e-a046-9c5d0bad4ec7	787c3dc1-b07e-40cc-aefe-f6e917ea9887	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:21.727322	f	f	f
58c58165-1e3b-4013-b80d-2868e84924e4	262f6dd4-9c45-4e2e-a7f3-1e5648cf14db	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:21.983895	f	f	f
aa7dbb2a-a6a9-44c0-a24e-d70e79a9d8c9	972be0ec-33f4-441b-96a5-ea689f8235fd	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:22.144387	f	f	f
8e0a72e1-5738-46f5-87c6-1338a6bbeb6e	9277be55-fc4b-4e84-adda-468f421d675a	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:22.390882	f	f	f
468c1934-b188-4c7d-9561-079cf2bcb948	8adddeca-1575-4103-8899-2b834867e8ac	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:22.550112	f	f	f
0f6cbe04-89e4-4095-858e-3aba8704c482	d3d5df5c-520e-4709-be51-130a184bdec5	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:22.797629	f	f	f
9a01f8c4-4be5-4084-9456-42874808f46c	83747cf3-361e-46b7-b599-302b45d375b8	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:22.956438	f	f	f
b7627abb-e869-4151-922b-e5c5fad460b5	b155c253-b64b-44f0-8e0d-512b1abca310	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:24.482548	f	f	f
58da5ae6-ea4b-4648-8793-53646a6a41ff	bdd286df-91ba-4260-9332-c76a65d5154b	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:24.628747	f	f	f
c0e2f454-0ead-4439-8043-868b1201d022	2b14bfc9-8429-4b36-b58b-14a46fbb53b8	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:24.797715	f	f	f
13026904-4819-4741-8ff7-c3873f4052a4	539942a9-161f-4503-a249-325393d15ce3	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:24.950211	f	f	f
b7bd190c-dab2-46e9-8dd5-f4624e73fe77	a316018d-0ba8-4804-b14b-16f18284dbbb	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:25.186604	f	f	f
81e14921-e3c5-49be-bd38-94c6776a987c	7856ff4c-15c8-4976-8018-940a98ae8e35	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:25.421893	f	f	f
d07db7f9-8f15-4967-92ef-ccfc1e25fd1b	e55d9ff0-4645-4812-8159-91a2fe1cedcb	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:25.656831	f	f	f
4af1ea2e-3a88-48f0-a883-5ce96aafaa15	4877bd92-80de-4908-a2fa-934139142601	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:25.884886	f	f	f
1794dc3f-88da-479e-9c03-ccf97eebf854	48969bc2-fa39-4032-9a9f-5b95dabb2367	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:26.118984	f	f	f
4e1669b7-b6bf-41c6-9100-811b5f139a16	302e43c7-4616-4a9c-8c0a-ff6f6846c7c5	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-10-19 18:51:26.355528	f	f	f
\.


--
-- Data for Name: profile_body_arts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profile_body_arts (profile_id, body_art_id) FROM stdin;
c5b40786-9fae-4599-9b8e-72e5262d9356	1
c5b40786-9fae-4599-9b8e-72e5262d9356	2
ca8bd72a-ddb7-4365-bd85-d76733a922c4	1
ca8bd72a-ddb7-4365-bd85-d76733a922c4	2
4ab29639-f7a5-498e-899c-a8590c493947	1
4ab29639-f7a5-498e-899c-a8590c493947	2
e82afe50-a1e3-40b9-9cfd-953c82e88539	1
e82afe50-a1e3-40b9-9cfd-953c82e88539	2
946b8cca-e1fd-4bb2-9a00-126e4b423eaf	1
946b8cca-e1fd-4bb2-9a00-126e4b423eaf	2
82194403-cdaf-42f6-8d11-129cd6ed716d	1
82194403-cdaf-42f6-8d11-129cd6ed716d	2
016bbc0f-968e-4cb6-969f-507ebc87d449	1
016bbc0f-968e-4cb6-969f-507ebc87d449	2
87f8f0d8-7619-4487-b8b4-cba515d4e35a	1
87f8f0d8-7619-4487-b8b4-cba515d4e35a	2
e604f1c9-6a84-411e-bd8f-3ad43a6d5250	1
e604f1c9-6a84-411e-bd8f-3ad43a6d5250	2
cc3e11e9-36fb-4412-abc6-a29f9ce8a4ae	1
cc3e11e9-36fb-4412-abc6-a29f9ce8a4ae	2
c1b985b5-bd3d-4c97-b199-f39d34d418da	1
c1b985b5-bd3d-4c97-b199-f39d34d418da	2
0c44fd8e-a13f-4343-ab79-35d708812cde	1
0c44fd8e-a13f-4343-ab79-35d708812cde	2
8cb74105-7b0b-48f8-99ea-c9ad193fa431	1
8cb74105-7b0b-48f8-99ea-c9ad193fa431	2
e0b731db-64dd-403b-b16e-40093c53dc26	1
e0b731db-64dd-403b-b16e-40093c53dc26	2
5ce524d1-e85c-4b70-8989-9e5a165e60c3	1
5ce524d1-e85c-4b70-8989-9e5a165e60c3	2
29be0a43-20cd-47e7-b7ed-e14c28289552	1
29be0a43-20cd-47e7-b7ed-e14c28289552	2
00ddbdec-c949-4d6a-94f1-0561eddf2501	1
00ddbdec-c949-4d6a-94f1-0561eddf2501	2
c4745658-ad6e-440e-82e2-4238991f0dbe	1
c4745658-ad6e-440e-82e2-4238991f0dbe	2
6e0b0a26-34cd-43ac-a4c0-1ef0c93fc477	1
6e0b0a26-34cd-43ac-a4c0-1ef0c93fc477	2
2b6e322d-e0ee-4bd6-b2b6-a293dd0465fe	1
2b6e322d-e0ee-4bd6-b2b6-a293dd0465fe	2
355e198b-356f-4f43-a655-1a4d22bfaa7e	1
355e198b-356f-4f43-a655-1a4d22bfaa7e	2
e89d7b36-2356-4592-ba0b-05f8e72486e6	1
e89d7b36-2356-4592-ba0b-05f8e72486e6	2
dccdbeb1-1809-417a-b3ac-9c70d54d1f88	1
dccdbeb1-1809-417a-b3ac-9c70d54d1f88	2
b9545a4e-f365-448f-a263-f7d3222c6d70	1
b9545a4e-f365-448f-a263-f7d3222c6d70	2
dcc5f238-ef17-445d-9196-491e2094d58a	1
dcc5f238-ef17-445d-9196-491e2094d58a	2
789aeeab-b85d-460c-919e-e6f9514ad182	1
789aeeab-b85d-460c-919e-e6f9514ad182	2
48a4fdb9-6af7-4dd2-8091-06df9c18f92a	1
48a4fdb9-6af7-4dd2-8091-06df9c18f92a	2
d54284e8-3db3-47d6-a900-4e115b51bdc8	1
d54284e8-3db3-47d6-a900-4e115b51bdc8	2
5e23d425-7d68-46f3-93e9-04309edc7da8	1
5e23d425-7d68-46f3-93e9-04309edc7da8	2
bb05a39c-eae3-42e7-ba39-5e12aae5be8f	1
bb05a39c-eae3-42e7-ba39-5e12aae5be8f	2
06c19758-852a-4192-afca-b632fbde6618	1
06c19758-852a-4192-afca-b632fbde6618	2
75c56ebc-7332-41ba-8595-00177621fb7a	1
75c56ebc-7332-41ba-8595-00177621fb7a	2
6712ac3e-8acf-4b96-a382-69c86f2a7de0	1
6712ac3e-8acf-4b96-a382-69c86f2a7de0	2
ddcb59d6-9a82-4ab6-918e-dce2097d4603	1
ddcb59d6-9a82-4ab6-918e-dce2097d4603	2
699a5205-affc-404a-8390-3bf8b6d1b8a0	1
699a5205-affc-404a-8390-3bf8b6d1b8a0	2
add0bf9a-b9f9-4203-8106-56ef6c2ec13e	1
add0bf9a-b9f9-4203-8106-56ef6c2ec13e	2
04ad1331-165d-4cab-bade-148029bc121e	1
04ad1331-165d-4cab-bade-148029bc121e	2
7961adc7-2bf6-4af8-abbb-25086b4791f6	1
7961adc7-2bf6-4af8-abbb-25086b4791f6	2
3dd222e0-da73-4f9a-be39-bf713d8085ea	1
3dd222e0-da73-4f9a-be39-bf713d8085ea	2
9f5f226e-60e3-44d1-a4dd-a7ed8fc3aed6	1
9f5f226e-60e3-44d1-a4dd-a7ed8fc3aed6	2
65b73e18-4751-4c67-b4e9-a966779ba9ea	1
65b73e18-4751-4c67-b4e9-a966779ba9ea	2
d35e126d-0e12-460c-a820-46730076a2cb	1
d35e126d-0e12-460c-a820-46730076a2cb	2
dfdd7f7a-7fbb-4e22-83b9-ea031dba83d8	1
dfdd7f7a-7fbb-4e22-83b9-ea031dba83d8	2
fe398ea0-93c8-4cc6-a1ab-3c39f11549a8	1
fe398ea0-93c8-4cc6-a1ab-3c39f11549a8	2
38e062b9-e847-4e78-86d3-50ede8cd63e4	1
38e062b9-e847-4e78-86d3-50ede8cd63e4	2
d10c5aca-e410-45b4-9442-0ce5b6f9ce46	1
d10c5aca-e410-45b4-9442-0ce5b6f9ce46	2
427578f5-ca68-4609-8531-7893d8c6b77f	1
427578f5-ca68-4609-8531-7893d8c6b77f	2
c80739e5-9f0c-47b6-b81c-f231b5d65607	1
c80739e5-9f0c-47b6-b81c-f231b5d65607	2
4d49b906-00d7-43b8-857f-1e2cccc742a4	1
4d49b906-00d7-43b8-857f-1e2cccc742a4	2
bc88f881-93d5-47ba-92b4-02a5cf706bf1	1
bc88f881-93d5-47ba-92b4-02a5cf706bf1	2
d0d9d46b-43e2-4a96-8cfd-36884e158373	1
d0d9d46b-43e2-4a96-8cfd-36884e158373	2
8c838cee-7e37-4b09-b5d4-68978e9c786d	1
8c838cee-7e37-4b09-b5d4-68978e9c786d	2
ffd12f80-86c3-4719-b1c8-d58e39bf5193	1
ffd12f80-86c3-4719-b1c8-d58e39bf5193	2
378f337e-2ad5-4e96-8c49-fd5168cbc973	1
378f337e-2ad5-4e96-8c49-fd5168cbc973	2
51e7bc12-66e4-468a-8e7d-116986eaa42c	1
51e7bc12-66e4-468a-8e7d-116986eaa42c	2
4c3182e9-278b-49f3-a293-8c833901a79c	1
4c3182e9-278b-49f3-a293-8c833901a79c	2
a1e85a83-feda-4d01-924b-283708c918d3	1
a1e85a83-feda-4d01-924b-283708c918d3	2
9ad81c52-2114-413a-963d-ef6bc31e0514	1
9ad81c52-2114-413a-963d-ef6bc31e0514	2
115d09d4-39e6-4250-9b22-f9ac0cd4ae37	1
115d09d4-39e6-4250-9b22-f9ac0cd4ae37	2
76c7e0a7-aea8-4955-8d1b-f1b59992069f	1
76c7e0a7-aea8-4955-8d1b-f1b59992069f	2
456b4dcf-99b8-46be-a00b-5135e42b5354	1
456b4dcf-99b8-46be-a00b-5135e42b5354	2
eefe9323-4063-441c-809b-03f734328caf	1
eefe9323-4063-441c-809b-03f734328caf	2
1312d36a-34f4-46c7-8dac-2fd4c304583e	1
1312d36a-34f4-46c7-8dac-2fd4c304583e	2
ab9fd6ea-7dba-42f2-92cd-6b8f99467ab1	1
ab9fd6ea-7dba-42f2-92cd-6b8f99467ab1	2
28c7af70-995a-4aa7-a108-a2ed41b85e25	1
28c7af70-995a-4aa7-a108-a2ed41b85e25	2
e4888c7c-2cc5-4d5f-94ce-85f2bb84ce85	1
e4888c7c-2cc5-4d5f-94ce-85f2bb84ce85	2
1454458e-2051-4698-99c5-d66a6e4891e9	1
1454458e-2051-4698-99c5-d66a6e4891e9	2
a63fd8b1-123f-48ed-9929-4b1c976f0791	1
a63fd8b1-123f-48ed-9929-4b1c976f0791	2
382ca26d-ed90-4b01-ae13-192267e2e719	1
382ca26d-ed90-4b01-ae13-192267e2e719	2
73351d20-4a7e-483b-b27d-76477972ac83	1
73351d20-4a7e-483b-b27d-76477972ac83	2
6fceb6dd-628a-460d-af0a-80809ba2e8e8	1
6fceb6dd-628a-460d-af0a-80809ba2e8e8	2
15c32c9b-4421-4c81-9295-faaf4e12cbc4	1
15c32c9b-4421-4c81-9295-faaf4e12cbc4	2
57edd4de-c007-4a1a-b0b7-d3a72eac4633	1
57edd4de-c007-4a1a-b0b7-d3a72eac4633	2
0601d0b3-4cfc-4372-b63d-b41fa012b1ff	1
0601d0b3-4cfc-4372-b63d-b41fa012b1ff	2
ac0e613a-0f21-4648-a4f7-907cbe5c931d	1
ac0e613a-0f21-4648-a4f7-907cbe5c931d	2
787c3dc1-b07e-40cc-aefe-f6e917ea9887	1
787c3dc1-b07e-40cc-aefe-f6e917ea9887	2
262f6dd4-9c45-4e2e-a7f3-1e5648cf14db	1
262f6dd4-9c45-4e2e-a7f3-1e5648cf14db	2
972be0ec-33f4-441b-96a5-ea689f8235fd	1
972be0ec-33f4-441b-96a5-ea689f8235fd	2
9277be55-fc4b-4e84-adda-468f421d675a	1
9277be55-fc4b-4e84-adda-468f421d675a	2
8adddeca-1575-4103-8899-2b834867e8ac	1
8adddeca-1575-4103-8899-2b834867e8ac	2
d3d5df5c-520e-4709-be51-130a184bdec5	1
d3d5df5c-520e-4709-be51-130a184bdec5	2
83747cf3-361e-46b7-b599-302b45d375b8	1
83747cf3-361e-46b7-b599-302b45d375b8	2
b155c253-b64b-44f0-8e0d-512b1abca310	1
b155c253-b64b-44f0-8e0d-512b1abca310	2
bdd286df-91ba-4260-9332-c76a65d5154b	1
bdd286df-91ba-4260-9332-c76a65d5154b	2
2b14bfc9-8429-4b36-b58b-14a46fbb53b8	1
2b14bfc9-8429-4b36-b58b-14a46fbb53b8	2
539942a9-161f-4503-a249-325393d15ce3	1
539942a9-161f-4503-a249-325393d15ce3	2
a316018d-0ba8-4804-b14b-16f18284dbbb	1
a316018d-0ba8-4804-b14b-16f18284dbbb	2
7856ff4c-15c8-4976-8018-940a98ae8e35	1
7856ff4c-15c8-4976-8018-940a98ae8e35	2
e55d9ff0-4645-4812-8159-91a2fe1cedcb	1
e55d9ff0-4645-4812-8159-91a2fe1cedcb	2
4877bd92-80de-4908-a2fa-934139142601	1
4877bd92-80de-4908-a2fa-934139142601	2
48969bc2-fa39-4032-9a9f-5b95dabb2367	1
48969bc2-fa39-4032-9a9f-5b95dabb2367	2
302e43c7-4616-4a9c-8c0a-ff6f6846c7c5	1
302e43c7-4616-4a9c-8c0a-ff6f6846c7c5	2
\.


--
-- Data for Name: profile_options; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profile_options (profile_id, profile_tag_id, price, comment) FROM stdin;
c5b40786-9fae-4599-9b8e-72e5262d9356	1	5000	This is my favourite!
c5b40786-9fae-4599-9b8e-72e5262d9356	2	50000	I hate this!
ca8bd72a-ddb7-4365-bd85-d76733a922c4	1	5000	This is my favourite!
ca8bd72a-ddb7-4365-bd85-d76733a922c4	2	50000	I hate this!
4ab29639-f7a5-498e-899c-a8590c493947	1	5000	This is my favourite!
4ab29639-f7a5-498e-899c-a8590c493947	2	50000	I hate this!
e82afe50-a1e3-40b9-9cfd-953c82e88539	1	5000	This is my favourite!
e82afe50-a1e3-40b9-9cfd-953c82e88539	2	50000	I hate this!
946b8cca-e1fd-4bb2-9a00-126e4b423eaf	1	5000	This is my favourite!
946b8cca-e1fd-4bb2-9a00-126e4b423eaf	2	50000	I hate this!
82194403-cdaf-42f6-8d11-129cd6ed716d	1	5000	This is my favourite!
82194403-cdaf-42f6-8d11-129cd6ed716d	2	50000	I hate this!
016bbc0f-968e-4cb6-969f-507ebc87d449	1	5000	This is my favourite!
016bbc0f-968e-4cb6-969f-507ebc87d449	2	50000	I hate this!
87f8f0d8-7619-4487-b8b4-cba515d4e35a	1	5000	This is my favourite!
87f8f0d8-7619-4487-b8b4-cba515d4e35a	2	50000	I hate this!
e604f1c9-6a84-411e-bd8f-3ad43a6d5250	1	5000	This is my favourite!
e604f1c9-6a84-411e-bd8f-3ad43a6d5250	2	50000	I hate this!
cc3e11e9-36fb-4412-abc6-a29f9ce8a4ae	1	5000	This is my favourite!
cc3e11e9-36fb-4412-abc6-a29f9ce8a4ae	2	50000	I hate this!
c1b985b5-bd3d-4c97-b199-f39d34d418da	1	5000	This is my favourite!
c1b985b5-bd3d-4c97-b199-f39d34d418da	2	50000	I hate this!
0c44fd8e-a13f-4343-ab79-35d708812cde	1	5000	This is my favourite!
0c44fd8e-a13f-4343-ab79-35d708812cde	2	50000	I hate this!
8cb74105-7b0b-48f8-99ea-c9ad193fa431	1	5000	This is my favourite!
8cb74105-7b0b-48f8-99ea-c9ad193fa431	2	50000	I hate this!
e0b731db-64dd-403b-b16e-40093c53dc26	1	5000	This is my favourite!
e0b731db-64dd-403b-b16e-40093c53dc26	2	50000	I hate this!
5ce524d1-e85c-4b70-8989-9e5a165e60c3	1	5000	This is my favourite!
5ce524d1-e85c-4b70-8989-9e5a165e60c3	2	50000	I hate this!
29be0a43-20cd-47e7-b7ed-e14c28289552	1	5000	This is my favourite!
29be0a43-20cd-47e7-b7ed-e14c28289552	2	50000	I hate this!
00ddbdec-c949-4d6a-94f1-0561eddf2501	1	5000	This is my favourite!
00ddbdec-c949-4d6a-94f1-0561eddf2501	2	50000	I hate this!
c4745658-ad6e-440e-82e2-4238991f0dbe	1	5000	This is my favourite!
c4745658-ad6e-440e-82e2-4238991f0dbe	2	50000	I hate this!
6e0b0a26-34cd-43ac-a4c0-1ef0c93fc477	1	5000	This is my favourite!
6e0b0a26-34cd-43ac-a4c0-1ef0c93fc477	2	50000	I hate this!
2b6e322d-e0ee-4bd6-b2b6-a293dd0465fe	1	5000	This is my favourite!
2b6e322d-e0ee-4bd6-b2b6-a293dd0465fe	2	50000	I hate this!
355e198b-356f-4f43-a655-1a4d22bfaa7e	1	5000	This is my favourite!
355e198b-356f-4f43-a655-1a4d22bfaa7e	2	50000	I hate this!
e89d7b36-2356-4592-ba0b-05f8e72486e6	1	5000	This is my favourite!
e89d7b36-2356-4592-ba0b-05f8e72486e6	2	50000	I hate this!
dccdbeb1-1809-417a-b3ac-9c70d54d1f88	1	5000	This is my favourite!
dccdbeb1-1809-417a-b3ac-9c70d54d1f88	2	50000	I hate this!
b9545a4e-f365-448f-a263-f7d3222c6d70	1	5000	This is my favourite!
b9545a4e-f365-448f-a263-f7d3222c6d70	2	50000	I hate this!
dcc5f238-ef17-445d-9196-491e2094d58a	1	5000	This is my favourite!
dcc5f238-ef17-445d-9196-491e2094d58a	2	50000	I hate this!
789aeeab-b85d-460c-919e-e6f9514ad182	1	5000	This is my favourite!
789aeeab-b85d-460c-919e-e6f9514ad182	2	50000	I hate this!
48a4fdb9-6af7-4dd2-8091-06df9c18f92a	1	5000	This is my favourite!
48a4fdb9-6af7-4dd2-8091-06df9c18f92a	2	50000	I hate this!
d54284e8-3db3-47d6-a900-4e115b51bdc8	1	5000	This is my favourite!
d54284e8-3db3-47d6-a900-4e115b51bdc8	2	50000	I hate this!
5e23d425-7d68-46f3-93e9-04309edc7da8	1	5000	This is my favourite!
5e23d425-7d68-46f3-93e9-04309edc7da8	2	50000	I hate this!
bb05a39c-eae3-42e7-ba39-5e12aae5be8f	1	5000	This is my favourite!
bb05a39c-eae3-42e7-ba39-5e12aae5be8f	2	50000	I hate this!
06c19758-852a-4192-afca-b632fbde6618	1	5000	This is my favourite!
06c19758-852a-4192-afca-b632fbde6618	2	50000	I hate this!
75c56ebc-7332-41ba-8595-00177621fb7a	1	5000	This is my favourite!
75c56ebc-7332-41ba-8595-00177621fb7a	2	50000	I hate this!
6712ac3e-8acf-4b96-a382-69c86f2a7de0	1	5000	This is my favourite!
6712ac3e-8acf-4b96-a382-69c86f2a7de0	2	50000	I hate this!
ddcb59d6-9a82-4ab6-918e-dce2097d4603	1	5000	This is my favourite!
ddcb59d6-9a82-4ab6-918e-dce2097d4603	2	50000	I hate this!
699a5205-affc-404a-8390-3bf8b6d1b8a0	1	5000	This is my favourite!
699a5205-affc-404a-8390-3bf8b6d1b8a0	2	50000	I hate this!
add0bf9a-b9f9-4203-8106-56ef6c2ec13e	1	5000	This is my favourite!
add0bf9a-b9f9-4203-8106-56ef6c2ec13e	2	50000	I hate this!
04ad1331-165d-4cab-bade-148029bc121e	1	5000	This is my favourite!
04ad1331-165d-4cab-bade-148029bc121e	2	50000	I hate this!
7961adc7-2bf6-4af8-abbb-25086b4791f6	1	5000	This is my favourite!
7961adc7-2bf6-4af8-abbb-25086b4791f6	2	50000	I hate this!
3dd222e0-da73-4f9a-be39-bf713d8085ea	1	5000	This is my favourite!
3dd222e0-da73-4f9a-be39-bf713d8085ea	2	50000	I hate this!
9f5f226e-60e3-44d1-a4dd-a7ed8fc3aed6	1	5000	This is my favourite!
9f5f226e-60e3-44d1-a4dd-a7ed8fc3aed6	2	50000	I hate this!
65b73e18-4751-4c67-b4e9-a966779ba9ea	1	5000	This is my favourite!
65b73e18-4751-4c67-b4e9-a966779ba9ea	2	50000	I hate this!
d35e126d-0e12-460c-a820-46730076a2cb	1	5000	This is my favourite!
d35e126d-0e12-460c-a820-46730076a2cb	2	50000	I hate this!
dfdd7f7a-7fbb-4e22-83b9-ea031dba83d8	1	5000	This is my favourite!
dfdd7f7a-7fbb-4e22-83b9-ea031dba83d8	2	50000	I hate this!
fe398ea0-93c8-4cc6-a1ab-3c39f11549a8	1	5000	This is my favourite!
fe398ea0-93c8-4cc6-a1ab-3c39f11549a8	2	50000	I hate this!
38e062b9-e847-4e78-86d3-50ede8cd63e4	1	5000	This is my favourite!
38e062b9-e847-4e78-86d3-50ede8cd63e4	2	50000	I hate this!
d10c5aca-e410-45b4-9442-0ce5b6f9ce46	1	5000	This is my favourite!
d10c5aca-e410-45b4-9442-0ce5b6f9ce46	2	50000	I hate this!
427578f5-ca68-4609-8531-7893d8c6b77f	1	5000	This is my favourite!
427578f5-ca68-4609-8531-7893d8c6b77f	2	50000	I hate this!
c80739e5-9f0c-47b6-b81c-f231b5d65607	1	5000	This is my favourite!
c80739e5-9f0c-47b6-b81c-f231b5d65607	2	50000	I hate this!
4d49b906-00d7-43b8-857f-1e2cccc742a4	1	5000	This is my favourite!
4d49b906-00d7-43b8-857f-1e2cccc742a4	2	50000	I hate this!
bc88f881-93d5-47ba-92b4-02a5cf706bf1	1	5000	This is my favourite!
bc88f881-93d5-47ba-92b4-02a5cf706bf1	2	50000	I hate this!
d0d9d46b-43e2-4a96-8cfd-36884e158373	1	5000	This is my favourite!
d0d9d46b-43e2-4a96-8cfd-36884e158373	2	50000	I hate this!
8c838cee-7e37-4b09-b5d4-68978e9c786d	1	5000	This is my favourite!
8c838cee-7e37-4b09-b5d4-68978e9c786d	2	50000	I hate this!
ffd12f80-86c3-4719-b1c8-d58e39bf5193	1	5000	This is my favourite!
ffd12f80-86c3-4719-b1c8-d58e39bf5193	2	50000	I hate this!
378f337e-2ad5-4e96-8c49-fd5168cbc973	1	5000	This is my favourite!
378f337e-2ad5-4e96-8c49-fd5168cbc973	2	50000	I hate this!
51e7bc12-66e4-468a-8e7d-116986eaa42c	1	5000	This is my favourite!
51e7bc12-66e4-468a-8e7d-116986eaa42c	2	50000	I hate this!
4c3182e9-278b-49f3-a293-8c833901a79c	1	5000	This is my favourite!
4c3182e9-278b-49f3-a293-8c833901a79c	2	50000	I hate this!
a1e85a83-feda-4d01-924b-283708c918d3	1	5000	This is my favourite!
a1e85a83-feda-4d01-924b-283708c918d3	2	50000	I hate this!
9ad81c52-2114-413a-963d-ef6bc31e0514	1	5000	This is my favourite!
9ad81c52-2114-413a-963d-ef6bc31e0514	2	50000	I hate this!
115d09d4-39e6-4250-9b22-f9ac0cd4ae37	1	5000	This is my favourite!
115d09d4-39e6-4250-9b22-f9ac0cd4ae37	2	50000	I hate this!
76c7e0a7-aea8-4955-8d1b-f1b59992069f	1	5000	This is my favourite!
76c7e0a7-aea8-4955-8d1b-f1b59992069f	2	50000	I hate this!
456b4dcf-99b8-46be-a00b-5135e42b5354	1	5000	This is my favourite!
456b4dcf-99b8-46be-a00b-5135e42b5354	2	50000	I hate this!
eefe9323-4063-441c-809b-03f734328caf	1	5000	This is my favourite!
eefe9323-4063-441c-809b-03f734328caf	2	50000	I hate this!
1312d36a-34f4-46c7-8dac-2fd4c304583e	1	5000	This is my favourite!
1312d36a-34f4-46c7-8dac-2fd4c304583e	2	50000	I hate this!
ab9fd6ea-7dba-42f2-92cd-6b8f99467ab1	1	5000	This is my favourite!
ab9fd6ea-7dba-42f2-92cd-6b8f99467ab1	2	50000	I hate this!
28c7af70-995a-4aa7-a108-a2ed41b85e25	1	5000	This is my favourite!
28c7af70-995a-4aa7-a108-a2ed41b85e25	2	50000	I hate this!
e4888c7c-2cc5-4d5f-94ce-85f2bb84ce85	1	5000	This is my favourite!
e4888c7c-2cc5-4d5f-94ce-85f2bb84ce85	2	50000	I hate this!
1454458e-2051-4698-99c5-d66a6e4891e9	1	5000	This is my favourite!
1454458e-2051-4698-99c5-d66a6e4891e9	2	50000	I hate this!
a63fd8b1-123f-48ed-9929-4b1c976f0791	1	5000	This is my favourite!
a63fd8b1-123f-48ed-9929-4b1c976f0791	2	50000	I hate this!
382ca26d-ed90-4b01-ae13-192267e2e719	1	5000	This is my favourite!
382ca26d-ed90-4b01-ae13-192267e2e719	2	50000	I hate this!
73351d20-4a7e-483b-b27d-76477972ac83	1	5000	This is my favourite!
73351d20-4a7e-483b-b27d-76477972ac83	2	50000	I hate this!
6fceb6dd-628a-460d-af0a-80809ba2e8e8	1	5000	This is my favourite!
6fceb6dd-628a-460d-af0a-80809ba2e8e8	2	50000	I hate this!
15c32c9b-4421-4c81-9295-faaf4e12cbc4	1	5000	This is my favourite!
15c32c9b-4421-4c81-9295-faaf4e12cbc4	2	50000	I hate this!
57edd4de-c007-4a1a-b0b7-d3a72eac4633	1	5000	This is my favourite!
57edd4de-c007-4a1a-b0b7-d3a72eac4633	2	50000	I hate this!
0601d0b3-4cfc-4372-b63d-b41fa012b1ff	1	5000	This is my favourite!
0601d0b3-4cfc-4372-b63d-b41fa012b1ff	2	50000	I hate this!
ac0e613a-0f21-4648-a4f7-907cbe5c931d	1	5000	This is my favourite!
ac0e613a-0f21-4648-a4f7-907cbe5c931d	2	50000	I hate this!
787c3dc1-b07e-40cc-aefe-f6e917ea9887	1	5000	This is my favourite!
787c3dc1-b07e-40cc-aefe-f6e917ea9887	2	50000	I hate this!
262f6dd4-9c45-4e2e-a7f3-1e5648cf14db	1	5000	This is my favourite!
262f6dd4-9c45-4e2e-a7f3-1e5648cf14db	2	50000	I hate this!
972be0ec-33f4-441b-96a5-ea689f8235fd	1	5000	This is my favourite!
972be0ec-33f4-441b-96a5-ea689f8235fd	2	50000	I hate this!
9277be55-fc4b-4e84-adda-468f421d675a	1	5000	This is my favourite!
9277be55-fc4b-4e84-adda-468f421d675a	2	50000	I hate this!
8adddeca-1575-4103-8899-2b834867e8ac	1	5000	This is my favourite!
8adddeca-1575-4103-8899-2b834867e8ac	2	50000	I hate this!
d3d5df5c-520e-4709-be51-130a184bdec5	1	5000	This is my favourite!
d3d5df5c-520e-4709-be51-130a184bdec5	2	50000	I hate this!
83747cf3-361e-46b7-b599-302b45d375b8	1	5000	This is my favourite!
83747cf3-361e-46b7-b599-302b45d375b8	2	50000	I hate this!
b155c253-b64b-44f0-8e0d-512b1abca310	1	5000	This is my favourite!
b155c253-b64b-44f0-8e0d-512b1abca310	2	50000	I hate this!
bdd286df-91ba-4260-9332-c76a65d5154b	1	5000	This is my favourite!
bdd286df-91ba-4260-9332-c76a65d5154b	2	50000	I hate this!
2b14bfc9-8429-4b36-b58b-14a46fbb53b8	1	5000	This is my favourite!
2b14bfc9-8429-4b36-b58b-14a46fbb53b8	2	50000	I hate this!
539942a9-161f-4503-a249-325393d15ce3	1	5000	This is my favourite!
539942a9-161f-4503-a249-325393d15ce3	2	50000	I hate this!
a316018d-0ba8-4804-b14b-16f18284dbbb	1	5000	This is my favourite!
a316018d-0ba8-4804-b14b-16f18284dbbb	2	50000	I hate this!
7856ff4c-15c8-4976-8018-940a98ae8e35	1	5000	This is my favourite!
7856ff4c-15c8-4976-8018-940a98ae8e35	2	50000	I hate this!
e55d9ff0-4645-4812-8159-91a2fe1cedcb	1	5000	This is my favourite!
e55d9ff0-4645-4812-8159-91a2fe1cedcb	2	50000	I hate this!
4877bd92-80de-4908-a2fa-934139142601	1	5000	This is my favourite!
4877bd92-80de-4908-a2fa-934139142601	2	50000	I hate this!
48969bc2-fa39-4032-9a9f-5b95dabb2367	1	5000	This is my favourite!
48969bc2-fa39-4032-9a9f-5b95dabb2367	2	50000	I hate this!
302e43c7-4616-4a9c-8c0a-ff6f6846c7c5	1	5000	This is my favourite!
302e43c7-4616-4a9c-8c0a-ff6f6846c7c5	2	50000	I hate this!
\.


--
-- Data for Name: profile_ratings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profile_ratings (id, service_id, profile_id, review_text_visible, review, score, created_at, updated_at, updated_by) FROM stdin;
f2ce913a-1626-4878-b012-5c67c22b81dd	54131916-4b8a-4137-b93a-98c1c23d99d3	c5b40786-9fae-4599-9b8e-72e5262d9356	t	I like the service! It's very good	5	2024-10-17 12:30:39.645788	2024-10-19 16:30:39.618404	00000000-0000-0000-0000-000000000000
c958bad3-8d16-4f00-83ef-5e529fae3906	6cd5378b-c9ec-44aa-beb7-701ba4f9e49a	e55d9ff0-4645-4812-8159-91a2fe1cedcb	t	I like the service! It's very good	5	2024-10-19 18:51:25.659184	2024-10-19 18:51:25.659184	00000000-0000-0000-0000-000000000000
dca12033-0383-4d2d-bc5c-a9704bcd7ef1	fa4c7daf-3a1a-4eac-abda-7429640addba	ca8bd72a-ddb7-4365-bd85-d76733a922c4	t	I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life	1	2024-10-18 13:30:39.807665	2024-10-19 16:30:39.82566	00000000-0000-0000-0000-000000000000
09b2c6b3-1ded-48b5-b8a4-052fccdd5cde	21b1d166-bd92-4495-84f2-afc684c534c5	4ab29639-f7a5-498e-899c-a8590c493947	t	I like the service! It's very good	5	2024-10-19 16:30:40.06891	2024-10-19 16:30:40.06891	00000000-0000-0000-0000-000000000000
23eade23-b8b0-431a-aaba-b18d5067c2b9	97060b58-ee59-45c3-aebf-dfe997702d42	e82afe50-a1e3-40b9-9cfd-953c82e88539	f	I like the service! It's very good	5	2024-10-19 16:30:40.244453	2024-10-19 16:30:40.253086	00000000-0000-0000-0000-000000000000
b78e0e3c-fa71-4a9e-b6a4-6ffd2fb2111d	219ed2ff-9310-4310-9556-1f73d62c4c7b	946b8cca-e1fd-4bb2-9a00-126e4b423eaf	t	I like the service! It's very good	5	2024-10-19 16:30:40.501366	2024-10-19 16:30:40.501366	00000000-0000-0000-0000-000000000000
eff62ab1-41bb-458e-9702-31cdccdffd91	00332c05-cb87-49d9-840e-bababe9f50ca	82194403-cdaf-42f6-8d11-129cd6ed716d	t	I like the service! It's very good	5	2024-10-19 16:30:40.668851	2024-10-19 16:30:40.668851	00000000-0000-0000-0000-000000000000
6d919a0b-a645-4582-a1ad-654328ce1260	281342f7-9538-45a2-a2e2-236cf8a9ebc2	016bbc0f-968e-4cb6-969f-507ebc87d449	t	I like the service! It's very good	5	2024-10-19 16:30:40.925154	2024-10-19 16:30:40.925154	00000000-0000-0000-0000-000000000000
52a88651-0b2a-4d89-99a4-2455c2f181bb	4906c968-4067-42f8-96d0-1a0acdfcbf57	87f8f0d8-7619-4487-b8b4-cba515d4e35a	t	I like the service! It's very good	5	2024-10-19 16:30:41.093931	2024-10-19 16:30:41.093931	00000000-0000-0000-0000-000000000000
4e57a35f-3047-4c47-b1ad-d1eeeb372f79	6d970b28-94a6-4eb3-b12e-9d1d05d45285	6712ac3e-8acf-4b96-a382-69c86f2a7de0	t	I like the service! It's very good	5	2024-10-17 14:50:40.621203	2024-10-19 18:50:40.58731	00000000-0000-0000-0000-000000000000
f1a0a31c-d761-41ba-b3df-98d7a4b00d28	218012d7-fdb8-466f-b4aa-49fd77c9a7a6	4877bd92-80de-4908-a2fa-934139142601	t	I like the service! It's very good	5	2024-10-19 18:51:25.887358	2024-10-19 18:51:25.887358	00000000-0000-0000-0000-000000000000
ed628ab7-a0d5-4165-95ea-b10856b59c9f	1fdeadae-0f63-41c2-ba84-439e33f3631e	ddcb59d6-9a82-4ab6-918e-dce2097d4603	t	I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life	1	2024-10-18 15:50:40.776281	2024-10-19 18:50:40.795359	00000000-0000-0000-0000-000000000000
db402d51-6c2e-4d1a-9719-f7b422df2297	8d615917-fb56-4593-bf55-49f6c99d5f70	699a5205-affc-404a-8390-3bf8b6d1b8a0	t	I like the service! It's very good	5	2024-10-19 18:50:41.02196	2024-10-19 18:50:41.02196	00000000-0000-0000-0000-000000000000
b8386436-8deb-4799-a4c1-b9a1930e0b23	378899e8-deef-4097-9177-075acfedf8d2	add0bf9a-b9f9-4203-8106-56ef6c2ec13e	f	I like the service! It's very good	5	2024-10-19 18:50:41.180069	2024-10-19 18:50:41.18746	00000000-0000-0000-0000-000000000000
cf80f59f-1c74-4be9-aafc-e2a57bc9b5c4	6f23717f-51b9-4ba5-a298-0939715f218d	04ad1331-165d-4cab-bade-148029bc121e	t	I like the service! It's very good	5	2024-10-19 18:50:41.421815	2024-10-19 18:50:41.421815	00000000-0000-0000-0000-000000000000
0bccb2ee-55e2-4618-9497-f0a55d57e242	6eb98269-2f29-47f6-b9c9-4936358126a4	7961adc7-2bf6-4af8-abbb-25086b4791f6	t	I like the service! It's very good	5	2024-10-19 18:50:41.574731	2024-10-19 18:50:41.574731	00000000-0000-0000-0000-000000000000
0e7d7ff5-e6f0-4562-b6ec-5502414d7767	df575a7c-6d1c-445a-a0bf-97c8519a46c3	3dd222e0-da73-4f9a-be39-bf713d8085ea	t	I like the service! It's very good	5	2024-10-19 18:50:41.813595	2024-10-19 18:50:41.813595	00000000-0000-0000-0000-000000000000
e2fe9ea2-4877-4126-b1f1-2553b42a87b8	2898b892-270c-469a-8cb5-9c861deba07d	9f5f226e-60e3-44d1-a4dd-a7ed8fc3aed6	t	I like the service! It's very good	5	2024-10-19 18:50:41.971903	2024-10-19 18:50:41.971903	00000000-0000-0000-0000-000000000000
0828083e-077b-4882-a345-27bae9a654ba	c24c24ce-7624-48ef-9043-1b73c3f01677	fe398ea0-93c8-4cc6-a1ab-3c39f11549a8	t	I like the service! It's very good	5	2024-10-19 18:50:43.910127	2024-10-19 18:50:43.910127	00000000-0000-0000-0000-000000000000
be172980-37ff-480c-82cb-adc7baab4ee4	c0f16764-33cd-40c1-8b18-3b9229938e52	38e062b9-e847-4e78-86d3-50ede8cd63e4	t	I like the service! It's very good	5	2024-10-19 18:50:44.16239	2024-10-19 18:50:44.16239	00000000-0000-0000-0000-000000000000
576ba991-7ad7-4b12-b7be-7025258f428a	556b03a9-d42a-4a71-89e3-5c31da6c4077	d10c5aca-e410-45b4-9442-0ce5b6f9ce46	t	I like the service! It's very good	5	2024-10-19 18:50:44.392438	2024-10-19 18:50:44.392438	00000000-0000-0000-0000-000000000000
2424e2d1-1b2f-4c02-af85-e35ec07643f8	8b1d4f23-535a-4335-aaea-bf065974a10d	427578f5-ca68-4609-8531-7893d8c6b77f	t	I like the service! It's very good	5	2024-10-19 18:50:44.621154	2024-10-19 18:50:44.621154	00000000-0000-0000-0000-000000000000
ccfc5beb-f7c5-486e-a4f8-5a589380cd83	ea40d566-c6a8-4eda-8274-85b97848d449	c80739e5-9f0c-47b6-b81c-f231b5d65607	t	I like the service! It's very good	5	2024-10-19 18:50:44.852568	2024-10-19 18:50:44.852568	00000000-0000-0000-0000-000000000000
98543712-e343-4c4f-b0b5-733246560801	2bf380a8-cd36-44eb-ab88-118f7a320247	4d49b906-00d7-43b8-857f-1e2cccc742a4	t	I like the service! It's very good	5	2024-10-19 18:50:45.082452	2024-10-19 18:50:45.082452	00000000-0000-0000-0000-000000000000
21ed0ae9-abe4-4677-a2f4-275dbd46b101	9dfccf88-de46-4fb6-b1b0-d3bdf3c2fbc9	bc88f881-93d5-47ba-92b4-02a5cf706bf1	t	I like the service! It's very good	5	2024-10-19 18:50:45.312865	2024-10-19 18:50:45.312865	00000000-0000-0000-0000-000000000000
a1f48d5b-9b5f-4908-b033-db38ed86f9dd	3301fce0-e1f9-4b92-8387-7b081ef7fb6f	ac0e613a-0f21-4648-a4f7-907cbe5c931d	t	I like the service! It's very good	5	2024-10-17 14:51:21.57818	2024-10-19 18:51:21.545517	00000000-0000-0000-0000-000000000000
12fcb215-503a-45c0-82ea-3636b6c54977	a6e08b54-032f-4daa-993e-54e9f4ffb6e8	48969bc2-fa39-4032-9a9f-5b95dabb2367	t	I like the service! It's very good	5	2024-10-19 18:51:26.121385	2024-10-19 18:51:26.121385	00000000-0000-0000-0000-000000000000
12424ee7-df0e-4fa2-a7f1-8e2e68fc80e5	a0fe791b-4c0d-4cf5-8ba7-bc2e069ff4ca	787c3dc1-b07e-40cc-aefe-f6e917ea9887	t	I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life	1	2024-10-18 15:51:21.736263	2024-10-19 18:51:21.756423	00000000-0000-0000-0000-000000000000
3660d59f-253a-425d-baa6-c0aa1ebd96ff	9834386e-7e61-418e-a068-b14510d0e2f2	262f6dd4-9c45-4e2e-a7f3-1e5648cf14db	t	I like the service! It's very good	5	2024-10-19 18:51:21.986461	2024-10-19 18:51:21.986461	00000000-0000-0000-0000-000000000000
ccd4af3e-e910-408c-82e2-f938cb8e41bd	a547e713-9e53-4f21-bcc5-7d1b706e4fcf	972be0ec-33f4-441b-96a5-ea689f8235fd	f	I like the service! It's very good	5	2024-10-19 18:51:22.146478	2024-10-19 18:51:22.155774	00000000-0000-0000-0000-000000000000
b59ba2a0-c080-464b-ae6f-d0f7343c124d	0a2643d3-b5f1-411a-9d21-adb7b4ffef58	9277be55-fc4b-4e84-adda-468f421d675a	t	I like the service! It's very good	5	2024-10-19 18:51:22.393451	2024-10-19 18:51:22.393451	00000000-0000-0000-0000-000000000000
13dbb3d0-3f90-48cd-a0ad-e7312e9cf552	cc71b8af-4cab-45f1-8d3d-be87e03a0feb	8adddeca-1575-4103-8899-2b834867e8ac	t	I like the service! It's very good	5	2024-10-19 18:51:22.554326	2024-10-19 18:51:22.554326	00000000-0000-0000-0000-000000000000
e1c99bb5-19f7-4e09-bbc4-8d94007a5826	823cfd22-d946-4b85-8a16-fbed87f473e6	d3d5df5c-520e-4709-be51-130a184bdec5	t	I like the service! It's very good	5	2024-10-19 18:51:22.800098	2024-10-19 18:51:22.800098	00000000-0000-0000-0000-000000000000
1f9045b0-e553-4ad5-9f6e-5278672a950c	5a6c047b-f262-4159-a300-a33d5dd9d1df	83747cf3-361e-46b7-b599-302b45d375b8	t	I like the service! It's very good	5	2024-10-19 18:51:22.958274	2024-10-19 18:51:22.958274	00000000-0000-0000-0000-000000000000
8e76ab55-6a76-4397-baf1-3484d8ce93e7	a4389442-20bf-4cb8-ab9b-4ce41e8ed7b2	539942a9-161f-4503-a249-325393d15ce3	t	I like the service! It's very good	5	2024-10-19 18:51:24.952265	2024-10-19 18:51:24.952265	00000000-0000-0000-0000-000000000000
26edc9a7-de6e-4eff-bf62-66a5958745a6	8fc5fe6d-341d-423d-b6d2-1339960f1721	a316018d-0ba8-4804-b14b-16f18284dbbb	t	I like the service! It's very good	5	2024-10-19 18:51:25.189531	2024-10-19 18:51:25.189531	00000000-0000-0000-0000-000000000000
66f6b276-93f4-436a-9c46-a722e9d7869f	6ae6a042-16b6-4793-bc4c-ad701b8c5f2a	7856ff4c-15c8-4976-8018-940a98ae8e35	t	I like the service! It's very good	5	2024-10-19 18:51:25.423856	2024-10-19 18:51:25.423856	00000000-0000-0000-0000-000000000000
62b2fd63-3960-41e7-a425-2ff497075e0b	d46214d1-ba49-4c2f-b772-33f28dea1e29	302e43c7-4616-4a9c-8c0a-ff6f6846c7c5	t	I like the service! It's very good	5	2024-10-19 18:51:26.358537	2024-10-19 18:51:26.358537	00000000-0000-0000-0000-000000000000
\.


--
-- Data for Name: profile_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profile_tags (id, name) FROM stdin;
1	Классика
2	Минет c/без резинки
3	Глубокий минет с резинкой
4	Глубокий минет c/без резинки c окончанием
5	Разрешу куннилингус
6	Минет c/без резинки c окончанием
7	Минет с резинкой
8	Массаж любительский
9	Массаж профессиональный
10	Вагинальный фистинг
11	Расслабляющий массаж
12	Поцелуи в губы
13	Массаж простаты
14	Классический массаж
15	Поеду отдыхать (в клуб, ресторан и.т.д.). Вечер:
16	Тайский боди массаж
17	Глубокий минет c/без резинки
18	Анилингус, побалую язычком очко
19	Услуги Госпоже
20	Услуги семейным парам
21	Французский поцелуй
22	Эротический массаж
23	Секс по телефону
24	Обслуживаю мальчишники. Вечер:
25	Групповой секс
26	Стриптиз любительский
27	Ветка сакуры
28	Снимусь на видео
29	Анальный секс
30	Ролевые игры, наряды
31	Фото на память
32	Сделаю минет
33	Стриптиз профессиональный
34	Глубокий минет без резинки c окончанием
35	Минет без резинки c окончанием
36	Глубокий минет без резинки
37	Минет без резинки
38	Обслуживаю девушек
39	Обслуживаю парней
40	Сделаю куннилингус
41	Обслуживаю девишники/вечеринки. Вечер:
42	Обслуживаю вечеринки. Вечер:
\.


--
-- Data for Name: profiles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profiles (body_type_id, ethnos_id, hair_color_id, intimate_hair_cut_id, city_id, parsed_url, id, user_id, active, phone, name, age, height, weight, bust, bio, address_latitude, address_longitude, price_in_house_night_ratio, price_in_house_contact, price_in_house_hour, prince_sauna_night_ratio, price_sauna_contact, price_sauna_hour, price_visit_night_ratio, price_visit_contact, price_visit_hour, price_car_night_ratio, price_car_contact, price_car_hour, contact_phone, contact_wa, contact_tg, moderated, moderated_at, moderated_by, verified, verified_at, verified_by, created_at, updated_at, updated_by, deleted_at) FROM stdin;
\N	72	3	2	14		c5b40786-9fae-4599-9b8e-72e5262d9356	2163ce47-18fe-49a8-a695-91c021c053ec	t	5938499458	Alicek4BHYWtciYv3pWb	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 16:30:39.604881	2024-10-19 16:30:39.604881	2163ce47-18fe-49a8-a695-91c021c053ec	\N
\N	55	1	1	47		ca8bd72a-ddb7-4365-bd85-d76733a922c4	f508c6ad-b68c-4fb1-8af0-1d71e35ba08f	t	3949298885	Alice6hck8snxCjpVx1o	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 16:30:39.798233	2024-10-19 16:30:39.798233	f508c6ad-b68c-4fb1-8af0-1d71e35ba08f	\N
\N	54	4	3	6		4ab29639-f7a5-498e-899c-a8590c493947	d51586dd-9251-4e3d-8943-3ba3187c3c67	t	2254720743	AliceHzIy2EWJ2dcZLdw	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 16:30:40.064251	2024-10-19 16:30:40.064251	d51586dd-9251-4e3d-8943-3ba3187c3c67	\N
\N	56	3	1	36		e82afe50-a1e3-40b9-9cfd-953c82e88539	77a4ca8e-cb02-4289-878c-4d5e1a55bb30	t	7693792793	Alicee8rp3WOUvaGpWxA	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 16:30:40.239399	2024-10-19 16:30:40.239399	77a4ca8e-cb02-4289-878c-4d5e1a55bb30	\N
\N	22	5	1	29		946b8cca-e1fd-4bb2-9a00-126e4b423eaf	0210b4cf-c38c-46a6-a880-a0e40dd383a9	t	3701090756	AliceAWl6x7xKyKMkWDa	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 16:30:40.497389	2024-10-19 16:30:40.497389	0210b4cf-c38c-46a6-a880-a0e40dd383a9	\N
\N	61	6	3	4		82194403-cdaf-42f6-8d11-129cd6ed716d	1de7291e-9f52-429b-b661-7a9f6c769477	t	9501140437	AlicepQ98Ue5veM2ThIY	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 16:30:40.664456	2024-10-19 16:30:40.664456	1de7291e-9f52-429b-b661-7a9f6c769477	\N
\N	52	6	3	24		016bbc0f-968e-4cb6-969f-507ebc87d449	d7e38c6d-5058-4ffb-9c3d-ea0192a590ac	t	8073694113	AliceDqgM0I5vZX8MGUj	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 16:30:40.920995	2024-10-19 16:30:40.920995	d7e38c6d-5058-4ffb-9c3d-ea0192a590ac	\N
\N	6	5	1	44		87f8f0d8-7619-4487-b8b4-cba515d4e35a	fa96765f-abe4-43bc-82fa-8a0137908029	t	4104838208	Aliceu0xmeHRi4MesyHO	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 16:30:41.090445	2024-10-19 16:30:41.090445	fa96765f-abe4-43bc-82fa-8a0137908029	\N
\N	55	6	3	40		e604f1c9-6a84-411e-bd8f-3ad43a6d5250	31352792-4bfe-4cff-93fc-ad70347a588e	t	4574760345	AliceAPFIbVoO1S0pFZA	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:36.642789	2024-10-19 18:50:36.642789	31352792-4bfe-4cff-93fc-ad70347a588e	\N
1	4	4	3	23		cc3e11e9-36fb-4412-abc6-a29f9ce8a4ae	f1b24062-938f-49c8-8f22-0977b242f3b7	f	4791448333	AliceAvjez8vXY2Zd7Sz-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:36.880237	2024-10-19 18:50:36.884225	f1b24062-938f-49c8-8f22-0977b242f3b7	\N
1	82	1	2	17		c1b985b5-bd3d-4c97-b199-f39d34d418da	896b1c66-c3ae-4f2b-8b9a-6356f7f0d498	f	6687610344	AlicelvHpDhdRoWjkqMN-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	t	2024-10-19 18:50:37.039466	725bcc59-3921-40b3-a034-65466859d5f3	f	2024-10-19 18:50:37.039466	725bcc59-3921-40b3-a034-65466859d5f3	2024-10-19 18:50:37.036886	2024-10-19 18:50:37.039466	725bcc59-3921-40b3-a034-65466859d5f3	\N
1	43	3	3	37		0c44fd8e-a13f-4343-ab79-35d708812cde	93679999-cfaa-4840-85e7-3d24d8f5ba38	f	7143280227	Alicefph9qdtw87k7JRk-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	t	2024-10-19 18:50:37.194032	1db2bc93-b016-4a44-9d1f-dc5df91313e8	f	2024-10-19 18:50:37.194032	1db2bc93-b016-4a44-9d1f-dc5df91313e8	2024-10-19 18:50:37.19141	2024-10-19 18:50:37.194032	1db2bc93-b016-4a44-9d1f-dc5df91313e8	\N
4	75	3	1	47		8cb74105-7b0b-48f8-99ea-c9ad193fa431	2e76a665-55b5-448a-a35a-dd090f334542	t	1511960871	AliceEel5VDjwYlkfgiY	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:37.343442	2024-10-19 18:50:37.343442	2e76a665-55b5-448a-a35a-dd090f334542	\N
1	88	1	2	4		e0b731db-64dd-403b-b16e-40093c53dc26	a221acf5-7eb0-498d-932e-d2c5c053d723	t	2824270031	Alicex92sgJWCyJHaKWd	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:37.494573	2024-10-19 18:50:37.494573	a221acf5-7eb0-498d-932e-d2c5c053d723	\N
1	89	6	3	11		5ce524d1-e85c-4b70-8989-9e5a165e60c3	1b479b0a-604b-454d-8816-dafb6c6f95b1	t	9296984975	AliceJvG2klH8aYDQgin	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:37.645755	2024-10-19 18:50:37.645755	1b479b0a-604b-454d-8816-dafb6c6f95b1	\N
4	84	5	1	40		29be0a43-20cd-47e7-b7ed-e14c28289552	c62ff00d-2436-446b-8bd9-eda0b5893261	t	2994445466	AlicePMduWhXykRd8ok9	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:37.794183	2024-10-19 18:50:37.794183	c62ff00d-2436-446b-8bd9-eda0b5893261	\N
2	20	1	2	9		00ddbdec-c949-4d6a-94f1-0561eddf2501	f6814fda-4853-41f7-94ed-1fa68ae4b8ec	t	7862292581	Alicea0Tw1wphGHDozoG	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:37.942452	2024-10-19 18:50:37.942452	f6814fda-4853-41f7-94ed-1fa68ae4b8ec	\N
1	14	5	1	14		c4745658-ad6e-440e-82e2-4238991f0dbe	f6814fda-4853-41f7-94ed-1fa68ae4b8ec	t	4889715529	AlicenFjKlZcRoa4xgWs	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:37.945001	2024-10-19 18:50:37.945001	f6814fda-4853-41f7-94ed-1fa68ae4b8ec	\N
3	61	3	1	14		6e0b0a26-34cd-43ac-a4c0-1ef0c93fc477	bb8e3c40-653f-4593-b9fb-f50179676ffa	t	9503275430	AlicepGlrLvXfPlth4d8	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.09406	2024-10-19 18:50:38.09406	bb8e3c40-653f-4593-b9fb-f50179676ffa	\N
4	86	6	2	15		2b6e322d-e0ee-4bd6-b2b6-a293dd0465fe	bb8e3c40-653f-4593-b9fb-f50179676ffa	t	6100511370	AliceucNto6gnzwNNvy5	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.096337	2024-10-19 18:50:38.096337	bb8e3c40-653f-4593-b9fb-f50179676ffa	\N
2	56	7	1	26		355e198b-356f-4f43-a655-1a4d22bfaa7e	28d8689b-a7c5-4f97-a9de-28d26f260b15	t	5357030455	AliceTMtcCP3mKKUaCKb	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.245455	2024-10-19 18:50:38.245455	28d8689b-a7c5-4f97-a9de-28d26f260b15	\N
2	83	4	2	3		e89d7b36-2356-4592-ba0b-05f8e72486e6	28d8689b-a7c5-4f97-a9de-28d26f260b15	t	7191704105	AlicebqUr95d3xnZ8Q42	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.248175	2024-10-19 18:50:38.248175	28d8689b-a7c5-4f97-a9de-28d26f260b15	\N
1	6	1	1	21		dccdbeb1-1809-417a-b3ac-9c70d54d1f88	09be1ac0-4a59-4a05-9a10-5dec31cfdb68	t	5895638963	Alice8QXK3xjwsWf77H8	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.396909	2024-10-19 18:50:38.396909	09be1ac0-4a59-4a05-9a10-5dec31cfdb68	\N
1	30	5	1	6		b9545a4e-f365-448f-a263-f7d3222c6d70	09be1ac0-4a59-4a05-9a10-5dec31cfdb68	t	8858999242	AliceQ1E0S8TBQ953T76	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.40317	2024-10-19 18:50:38.40317	09be1ac0-4a59-4a05-9a10-5dec31cfdb68	\N
4	6	4	2	35		dcc5f238-ef17-445d-9196-491e2094d58a	09be1ac0-4a59-4a05-9a10-5dec31cfdb68	f	4958205543	AliceZj8V7ky5yQBqrKd-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.408352	2024-10-19 18:50:38.410587	77165e6d-2f37-4683-a3ed-d817968fbecb	\N
3	26	1	1	11		789aeeab-b85d-460c-919e-e6f9514ad182	04e7de45-0c4a-48b8-a41a-54c5bf3be53f	t	9643216320	AliceFUEFlwRlUZRLlwP	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.493072	2024-10-19 18:50:38.493072	04e7de45-0c4a-48b8-a41a-54c5bf3be53f	\N
1	89	5	1	24		48a4fdb9-6af7-4dd2-8091-06df9c18f92a	484927ea-5c61-48d4-af93-92e5103e4668	t	1246550881	Alice10GDNBMdOxI0u7U	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.644294	2024-10-19 18:50:38.644294	484927ea-5c61-48d4-af93-92e5103e4668	\N
3	89	1	1	35		d54284e8-3db3-47d6-a900-4e115b51bdc8	57704b99-2611-4f5b-8154-6887eb6c8c84	t	3319038690	AliceawkbdCj2D8oNtoB	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.722251	2024-10-19 18:50:38.722251	57704b99-2611-4f5b-8154-6887eb6c8c84	2024-10-19 15:50:38.727796+00
1	73	4	3	13		5e23d425-7d68-46f3-93e9-04309edc7da8	f0b3c9f1-77c9-4bde-a128-c78b7b7ad86e	t	6718666805	AliceEYs8iNsSM0XDdG4	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.806798	2024-10-19 18:50:38.806798	f0b3c9f1-77c9-4bde-a128-c78b7b7ad86e	\N
2	64	4	3	15		bb05a39c-eae3-42e7-ba39-5e12aae5be8f	6867864c-93ec-45a1-8a73-08ea25a4b220	t	3004869489	Alicehdbuo1O3VzuJnpl	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.965336	2024-10-19 18:50:38.965336	6867864c-93ec-45a1-8a73-08ea25a4b220	\N
1	29	6	2	7		06c19758-852a-4192-afca-b632fbde6618	6867864c-93ec-45a1-8a73-08ea25a4b220	t	9210253370	AlicemqidCeA4v0cjtpb	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.968046	2024-10-19 18:50:38.968046	6867864c-93ec-45a1-8a73-08ea25a4b220	\N
4	76	1	2	6		75c56ebc-7332-41ba-8595-00177621fb7a	6867864c-93ec-45a1-8a73-08ea25a4b220	f	2588447576	AliceXtSHLc4HrCRqZIX-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:38.970156	2024-10-19 18:50:38.972519	e36d3814-6a6c-493d-9a72-91c97c15c1d1	\N
\N	77	1	3	5		6712ac3e-8acf-4b96-a382-69c86f2a7de0	6890ef0b-0091-4d2d-b25f-f8800b577ea3	t	5848180854	AliceF6USGdUWJhatcFg	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:40.575308	2024-10-19 18:50:40.575308	6890ef0b-0091-4d2d-b25f-f8800b577ea3	\N
\N	34	6	3	36		ddcb59d6-9a82-4ab6-918e-dce2097d4603	ef6d7a38-93ec-4a62-aec0-63027cb431ea	t	1647519794	AliceEOAitaqQDmummx4	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:40.766886	2024-10-19 18:50:40.766886	ef6d7a38-93ec-4a62-aec0-63027cb431ea	\N
\N	81	2	1	2		699a5205-affc-404a-8390-3bf8b6d1b8a0	6e99a07e-2ac3-466a-a00c-419b12f3f2e7	t	1968005728	AliceFHlkYl5DmkNPRpi	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:41.01843	2024-10-19 18:50:41.01843	6e99a07e-2ac3-466a-a00c-419b12f3f2e7	\N
\N	75	3	3	35		add0bf9a-b9f9-4203-8106-56ef6c2ec13e	3ae0a69d-27f8-4cbe-8e5b-0f185dafe009	t	2751248245	AliceO8ktxQtwMdjEhZ7	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:41.177036	2024-10-19 18:50:41.177036	3ae0a69d-27f8-4cbe-8e5b-0f185dafe009	\N
\N	51	2	3	30		04ad1331-165d-4cab-bade-148029bc121e	e4166f80-9a47-4bd2-a740-27c8fb6a48cb	t	5064632896	AlicedH6VfYMNmVjrx08	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:41.41831	2024-10-19 18:50:41.41831	e4166f80-9a47-4bd2-a740-27c8fb6a48cb	\N
\N	42	6	2	35		7961adc7-2bf6-4af8-abbb-25086b4791f6	b5cecb52-7877-4029-afa5-e677a1c2ec7d	t	7291106221	AliceyYHTkPwTZwBSulQ	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:41.5715	2024-10-19 18:50:41.5715	b5cecb52-7877-4029-afa5-e677a1c2ec7d	\N
\N	66	7	1	17		3dd222e0-da73-4f9a-be39-bf713d8085ea	0ee8284c-e553-4816-b928-2415ab336d72	t	4258424028	AliceOBGexhx37SqsoOp	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:41.810885	2024-10-19 18:50:41.810885	0ee8284c-e553-4816-b928-2415ab336d72	\N
\N	72	1	3	25		9f5f226e-60e3-44d1-a4dd-a7ed8fc3aed6	a2bc57d6-4555-4230-ac8e-712bdfbeaea1	t	2684519080	Alicealu9vUGhRwIby6W	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:41.969586	2024-10-19 18:50:41.969586	a2bc57d6-4555-4230-ac8e-712bdfbeaea1	\N
\N	72	7	2	42		65b73e18-4751-4c67-b4e9-a966779ba9ea	44d4a60e-79d9-4c82-bb32-11b3ef8b5a73	t	2395373138	AliceA2cUbWfzKEnyNl6	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:43.428035	2024-10-19 18:50:43.428035	44d4a60e-79d9-4c82-bb32-11b3ef8b5a73	\N
\N	55	5	2	19		d35e126d-0e12-460c-a820-46730076a2cb	5d0c334e-272c-43b3-a361-4726a2dfbf85	t	9652719403	AliceH6sP2NMZXLoaj4M	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:43.580647	2024-10-19 18:50:43.580647	5d0c334e-272c-43b3-a361-4726a2dfbf85	\N
\N	40	6	1	8		dfdd7f7a-7fbb-4e22-83b9-ea031dba83d8	100bc084-ae37-42db-938d-c71d32dbb89c	t	3467758158	AliceIIWG45BxwtVt7Lc	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:43.752132	2024-10-19 18:50:43.752132	100bc084-ae37-42db-938d-c71d32dbb89c	\N
\N	84	7	1	10		fe398ea0-93c8-4cc6-a1ab-3c39f11549a8	2201e804-7513-4100-93d1-e3f247e49e84	t	6388447692	AliceBmVPLdwJxm1zpOx	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:43.907009	2024-10-19 18:50:43.907009	2201e804-7513-4100-93d1-e3f247e49e84	\N
\N	71	1	2	37		38e062b9-e847-4e78-86d3-50ede8cd63e4	579affde-faf8-4b01-9727-d964d5ec456a	t	1786085378	AliceBI6YuwVdVyPGtON	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:44.159125	2024-10-19 18:50:44.159125	579affde-faf8-4b01-9727-d964d5ec456a	\N
\N	82	6	3	30		d10c5aca-e410-45b4-9442-0ce5b6f9ce46	d4db4b08-437c-4ebc-a696-c4ad35e75734	t	1698994574	AlicehNmngHNAfsBBJsU	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:44.388744	2024-10-19 18:50:44.388744	d4db4b08-437c-4ebc-a696-c4ad35e75734	\N
\N	30	2	1	8		427578f5-ca68-4609-8531-7893d8c6b77f	f4c07801-8c69-47f8-9719-e439d500639c	t	4569545409	Alice5WH2JuTTSjY5wUD	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:44.618779	2024-10-19 18:50:44.618779	f4c07801-8c69-47f8-9719-e439d500639c	\N
\N	35	4	3	42		c80739e5-9f0c-47b6-b81c-f231b5d65607	9bee80eb-4c99-4823-90ac-c056be769100	t	4714218504	AliceZSuywqhNYUIfzmv	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:44.84956	2024-10-19 18:50:44.84956	9bee80eb-4c99-4823-90ac-c056be769100	\N
\N	5	4	1	27		4d49b906-00d7-43b8-857f-1e2cccc742a4	524b58ed-10f5-4921-82bf-852780979138	t	3688821785	AliceM7tTs0e5RQnskXb	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:45.079534	2024-10-19 18:50:45.079534	524b58ed-10f5-4921-82bf-852780979138	\N
\N	27	2	1	24		bc88f881-93d5-47ba-92b4-02a5cf706bf1	846c7558-68d2-4074-aad8-1a740629a2b3	t	2566910304	AliceskUtG9YGZX4Z8Ut	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:50:45.310129	2024-10-19 18:50:45.310129	846c7558-68d2-4074-aad8-1a740629a2b3	\N
\N	58	1	2	16		d0d9d46b-43e2-4a96-8cfd-36884e158373	ed38bd8b-6748-4e08-a275-a1311cdf4366	t	9587815911	Aliceh7wnbSeCrCWCabq	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:17.539328	2024-10-19 18:51:17.539328	ed38bd8b-6748-4e08-a275-a1311cdf4366	\N
3	17	4	1	36		8c838cee-7e37-4b09-b5d4-68978e9c786d	6d5733a9-3b1b-4de3-a823-a9dad14feaeb	f	5249378663	AliceecejTbWyR04Yg1y-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:17.784391	2024-10-19 18:51:17.792919	6d5733a9-3b1b-4de3-a823-a9dad14feaeb	\N
4	13	5	2	26		ffd12f80-86c3-4719-b1c8-d58e39bf5193	c550de45-58f4-48ce-8a4f-925b3b67bc21	f	7512208600	AliceCE8VojX8rzdqID0-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	t	2024-10-19 18:51:17.950125	e8cd1dac-44f8-49a6-a127-268826d339f9	f	2024-10-19 18:51:17.950125	e8cd1dac-44f8-49a6-a127-268826d339f9	2024-10-19 18:51:17.947015	2024-10-19 18:51:17.950125	e8cd1dac-44f8-49a6-a127-268826d339f9	\N
3	13	5	3	34		378f337e-2ad5-4e96-8c49-fd5168cbc973	e9b634ef-54f7-4126-8b02-817088c4b18f	f	4752629391	AliceEmnRH7ySvgv6cHg-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	t	2024-10-19 18:51:18.111206	438f3759-66ce-405a-9752-dabcc51e4773	f	2024-10-19 18:51:18.111206	438f3759-66ce-405a-9752-dabcc51e4773	2024-10-19 18:51:18.105831	2024-10-19 18:51:18.111206	438f3759-66ce-405a-9752-dabcc51e4773	\N
2	1	5	3	10		51e7bc12-66e4-468a-8e7d-116986eaa42c	3d83697f-ca21-4efb-aea2-c01202a6b6fc	t	6381152304	AliceBab9VKvAxz9fmTe	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:18.261855	2024-10-19 18:51:18.261855	3d83697f-ca21-4efb-aea2-c01202a6b6fc	\N
2	19	5	3	34		4c3182e9-278b-49f3-a293-8c833901a79c	fc4e2ce9-464c-4a5e-8879-6241ea3abd11	t	7556565115	AliceIxZ3YzTAMqmaIWM	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:18.421714	2024-10-19 18:51:18.421714	fc4e2ce9-464c-4a5e-8879-6241ea3abd11	\N
2	55	6	3	20		a1e85a83-feda-4d01-924b-283708c918d3	14ac0578-89fb-4c27-8e40-de14fea7a34d	t	3946455842	Alicet1i5OyuKSUbsxIc	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:18.57382	2024-10-19 18:51:18.57382	14ac0578-89fb-4c27-8e40-de14fea7a34d	\N
3	11	3	1	28		9ad81c52-2114-413a-963d-ef6bc31e0514	7650136e-f3ab-4161-ab94-bde90c176d6b	t	8055265386	Alicewwu4mRj0M4gkFmL	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:18.725298	2024-10-19 18:51:18.725298	7650136e-f3ab-4161-ab94-bde90c176d6b	\N
2	27	7	2	41		115d09d4-39e6-4250-9b22-f9ac0cd4ae37	62eae9c1-2f4f-4b92-bb80-f40c3f9fd43f	t	8130547706	AliceUR31iVAQQYquRgf	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:18.875308	2024-10-19 18:51:18.875308	62eae9c1-2f4f-4b92-bb80-f40c3f9fd43f	\N
1	39	4	1	44		76c7e0a7-aea8-4955-8d1b-f1b59992069f	62eae9c1-2f4f-4b92-bb80-f40c3f9fd43f	t	4990226424	Alicewchlq9jKkY55nIQ	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:18.877587	2024-10-19 18:51:18.877587	62eae9c1-2f4f-4b92-bb80-f40c3f9fd43f	\N
4	52	3	3	20		456b4dcf-99b8-46be-a00b-5135e42b5354	ff18a582-3447-4ee3-b613-3596944df7b3	t	7347223844	Alice8aqEHOKL2RAvpUs	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.03	2024-10-19 18:51:19.03	ff18a582-3447-4ee3-b613-3596944df7b3	\N
4	32	4	1	38		eefe9323-4063-441c-809b-03f734328caf	ff18a582-3447-4ee3-b613-3596944df7b3	t	8596957038	AliceTrAdMA1ceGssFT1	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.03252	2024-10-19 18:51:19.03252	ff18a582-3447-4ee3-b613-3596944df7b3	\N
4	38	6	2	37		1312d36a-34f4-46c7-8dac-2fd4c304583e	66e60f34-a0ac-46da-a0e2-85268e27b1f5	t	2050776824	AliceoI8LtQ3SQjqKbvt	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.185896	2024-10-19 18:51:19.185896	66e60f34-a0ac-46da-a0e2-85268e27b1f5	\N
1	88	1	1	20		ab9fd6ea-7dba-42f2-92cd-6b8f99467ab1	66e60f34-a0ac-46da-a0e2-85268e27b1f5	t	9826714551	AliceaOZFbuNkN560ymY	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.188381	2024-10-19 18:51:19.188381	66e60f34-a0ac-46da-a0e2-85268e27b1f5	\N
2	41	3	2	23		28c7af70-995a-4aa7-a108-a2ed41b85e25	a1bb7f29-069f-4178-8afb-8fa1bd86f57d	t	8428745990	Alice28uhhPiYh4lDQzw	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.345163	2024-10-19 18:51:19.345163	a1bb7f29-069f-4178-8afb-8fa1bd86f57d	\N
1	93	6	2	15		e4888c7c-2cc5-4d5f-94ce-85f2bb84ce85	a1bb7f29-069f-4178-8afb-8fa1bd86f57d	t	4542282834	AliceTvntNL2jUrQwdms	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.34772	2024-10-19 18:51:19.34772	a1bb7f29-069f-4178-8afb-8fa1bd86f57d	\N
4	81	3	1	15		1454458e-2051-4698-99c5-d66a6e4891e9	a1bb7f29-069f-4178-8afb-8fa1bd86f57d	f	9143823613	Alice8cBop2u1m9ZHlMZ-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.349913	2024-10-19 18:51:19.353264	c3eeaf15-7713-450d-ba97-f16abb9a07be	\N
2	1	1	3	39		a63fd8b1-123f-48ed-9929-4b1c976f0791	a40061be-c9df-49fe-b407-a4bf8ec4fe53	t	3499915581	AliceJ7LBYxHOaHyYmoC	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.433063	2024-10-19 18:51:19.433063	a40061be-c9df-49fe-b407-a4bf8ec4fe53	\N
1	13	6	1	11		382ca26d-ed90-4b01-ae13-192267e2e719	e3b15f4d-7885-402f-93f9-424b59ce7fa3	t	8713816463	AlicesckRiXW7zo2H6Xz	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.583839	2024-10-19 18:51:19.583839	e3b15f4d-7885-402f-93f9-424b59ce7fa3	\N
2	50	3	3	35		73351d20-4a7e-483b-b27d-76477972ac83	075187c4-589f-4cc5-8e60-40d3286fced0	t	2022669479	AliceRA4vlu7qrDLzdPw	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.668055	2024-10-19 18:51:19.668055	075187c4-589f-4cc5-8e60-40d3286fced0	2024-10-19 15:51:19.674387+00
2	61	1	3	6		6fceb6dd-628a-460d-af0a-80809ba2e8e8	d10409c6-84fb-4839-8197-55387ba0b79d	t	3066024293	Alice5w3WOca9m3MXYBh	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.751408	2024-10-19 18:51:19.751408	d10409c6-84fb-4839-8197-55387ba0b79d	\N
3	62	7	2	22		15c32c9b-4421-4c81-9295-faaf4e12cbc4	3fc78a7b-2fed-4c72-8519-63353eabbad9	t	2941271993	Alice6lsDMeakpRAXcY1	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.910964	2024-10-19 18:51:19.910964	3fc78a7b-2fed-4c72-8519-63353eabbad9	\N
4	77	3	3	28		57edd4de-c007-4a1a-b0b7-d3a72eac4633	3fc78a7b-2fed-4c72-8519-63353eabbad9	t	6350581817	Alice4jG5BEDg8fDMj0u	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.913355	2024-10-19 18:51:19.913355	3fc78a7b-2fed-4c72-8519-63353eabbad9	\N
4	35	6	3	48		0601d0b3-4cfc-4372-b63d-b41fa012b1ff	3fc78a7b-2fed-4c72-8519-63353eabbad9	f	8550298002	Alice0YHT127m3x8Tc7D-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:19.915773	2024-10-19 18:51:19.91897	1dc4f64e-3a60-4c31-bf59-b83ea79d6b5b	\N
\N	45	2	2	33		ac0e613a-0f21-4648-a4f7-907cbe5c931d	affe57b1-aa52-41ba-bbdc-572b6d53bc0e	t	5745676586	AliceGl9wtDqV7o2iS6d	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:21.53119	2024-10-19 18:51:21.53119	affe57b1-aa52-41ba-bbdc-572b6d53bc0e	\N
\N	90	3	3	17		787c3dc1-b07e-40cc-aefe-f6e917ea9887	4c274022-dbde-4d41-8673-5849b20c4052	t	1235557273	AlicegpjVxDd1gwdB2s7	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:21.725996	2024-10-19 18:51:21.725996	4c274022-dbde-4d41-8673-5849b20c4052	\N
\N	88	1	3	25		262f6dd4-9c45-4e2e-a7f3-1e5648cf14db	f7d194a0-166b-4cb4-8c89-71c519d07d34	t	1189837458	AliceOA3MbqZ9ldVVSu4	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:21.982696	2024-10-19 18:51:21.982696	f7d194a0-166b-4cb4-8c89-71c519d07d34	\N
\N	37	2	1	19		972be0ec-33f4-441b-96a5-ea689f8235fd	20a76176-0042-4923-a381-2aab9954e679	t	8083314161	AliceExJaMASXSCHAW2k	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:22.143355	2024-10-19 18:51:22.143355	20a76176-0042-4923-a381-2aab9954e679	\N
\N	93	1	1	5		9277be55-fc4b-4e84-adda-468f421d675a	6e379ad4-5b11-4aa2-8c03-e6d12e032755	t	4542308188	Alice9PUnzyHzCJaax0p	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:22.386762	2024-10-19 18:51:22.386762	6e379ad4-5b11-4aa2-8c03-e6d12e032755	\N
\N	19	1	3	23		8adddeca-1575-4103-8899-2b834867e8ac	efc3d895-a248-45ba-9841-b9d638a53175	t	1736625217	Alicek3feYelPPn3C6Qx	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:22.548318	2024-10-19 18:51:22.548318	efc3d895-a248-45ba-9841-b9d638a53175	\N
\N	87	6	3	14		d3d5df5c-520e-4709-be51-130a184bdec5	1628db9d-b81a-4922-a905-caa0d4c16a48	t	4673387142	AliceC7OceDIJW2xGgxJ	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:22.796462	2024-10-19 18:51:22.796462	1628db9d-b81a-4922-a905-caa0d4c16a48	\N
\N	15	3	2	14		83747cf3-361e-46b7-b599-302b45d375b8	a42aa967-9a5c-48ee-8ae7-f4334ea0b56b	t	6467922432	Alicejx5jtuipUiTrPBQ	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:22.95561	2024-10-19 18:51:22.95561	a42aa967-9a5c-48ee-8ae7-f4334ea0b56b	\N
\N	47	5	2	3		b155c253-b64b-44f0-8e0d-512b1abca310	4a1c4b50-e174-4410-82a6-7257a0002dc5	t	4947562919	AlicedAkf2KJzK9CY5DJ	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:24.472283	2024-10-19 18:51:24.472283	4a1c4b50-e174-4410-82a6-7257a0002dc5	\N
\N	85	6	3	5		bdd286df-91ba-4260-9332-c76a65d5154b	6490cf88-e23a-4be5-855f-91a59e00fef6	t	3857753454	AlicenmOadDoNhyxrBQO	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:24.627545	2024-10-19 18:51:24.627545	6490cf88-e23a-4be5-855f-91a59e00fef6	\N
\N	68	1	3	7		2b14bfc9-8429-4b36-b58b-14a46fbb53b8	f23e28e0-1095-494b-9778-2ce1289423ea	t	1509557314	AliceG7HAX0EcCjxGZyC	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:24.796072	2024-10-19 18:51:24.796072	f23e28e0-1095-494b-9778-2ce1289423ea	\N
\N	65	6	2	1		539942a9-161f-4503-a249-325393d15ce3	6f5beca5-2aad-4533-8d42-7cfebf2b0532	t	8614128633	AliceGJzAcpapFl8PJX8	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:24.949233	2024-10-19 18:51:24.949233	6f5beca5-2aad-4533-8d42-7cfebf2b0532	\N
\N	84	2	1	35		a316018d-0ba8-4804-b14b-16f18284dbbb	a5f7b81f-00f7-4e1e-8526-b27cde367f21	t	5485099676	AlicecJ2hlCgSSLsS7xt	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:25.185487	2024-10-19 18:51:25.185487	a5f7b81f-00f7-4e1e-8526-b27cde367f21	\N
\N	53	5	1	38		7856ff4c-15c8-4976-8018-940a98ae8e35	a4322a31-4340-4222-88b3-683d0e210b7e	t	8380135471	AliceunDTA2LjXyS060G	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:25.420483	2024-10-19 18:51:25.420483	a4322a31-4340-4222-88b3-683d0e210b7e	\N
\N	92	3	2	16		e55d9ff0-4645-4812-8159-91a2fe1cedcb	4e08df8a-683c-4565-8b39-700c1e4275ad	t	2933332583	AliceFQaHF3Al03vynPL	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:25.654407	2024-10-19 18:51:25.654407	4e08df8a-683c-4565-8b39-700c1e4275ad	\N
\N	46	4	1	9		4877bd92-80de-4908-a2fa-934139142601	0a1a6ce6-0085-46f1-9803-7175849a1c28	t	2713173161	Alice4fD8URJfMpzuZ1N	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:25.882073	2024-10-19 18:51:25.882073	0a1a6ce6-0085-46f1-9803-7175849a1c28	\N
\N	55	2	3	27		48969bc2-fa39-4032-9a9f-5b95dabb2367	17baa83f-62a2-4662-9a23-404a5d3e684d	t	7304582445	AliceXSwJrAfNbOfG0dG	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:26.117298	2024-10-19 18:51:26.117298	17baa83f-62a2-4662-9a23-404a5d3e684d	\N
\N	22	5	3	20		302e43c7-4616-4a9c-8c0a-ff6f6846c7c5	aa8bc7e4-be00-4aee-9ad2-85c47f2395b3	t	6242573683	AliceUfkZXsNQtDM7bc8	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-10-19 18:51:26.353738	2024-10-19 18:51:26.353738	aa8bc7e4-be00-4aee-9ad2-85c47f2395b3	\N
\.


--
-- Data for Name: rated_profile_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.rated_profile_tags (rating_id, profile_tag_id, type) FROM stdin;
f2ce913a-1626-4878-b012-5c67c22b81dd	1	like
f2ce913a-1626-4878-b012-5c67c22b81dd	2	like
dca12033-0383-4d2d-bc5c-a9704bcd7ef1	1	dislike
09b2c6b3-1ded-48b5-b8a4-052fccdd5cde	1	like
09b2c6b3-1ded-48b5-b8a4-052fccdd5cde	2	like
23eade23-b8b0-431a-aaba-b18d5067c2b9	1	like
23eade23-b8b0-431a-aaba-b18d5067c2b9	2	like
b78e0e3c-fa71-4a9e-b6a4-6ffd2fb2111d	1	like
b78e0e3c-fa71-4a9e-b6a4-6ffd2fb2111d	2	like
eff62ab1-41bb-458e-9702-31cdccdffd91	1	like
eff62ab1-41bb-458e-9702-31cdccdffd91	2	like
6d919a0b-a645-4582-a1ad-654328ce1260	1	like
6d919a0b-a645-4582-a1ad-654328ce1260	2	like
52a88651-0b2a-4d89-99a4-2455c2f181bb	1	like
52a88651-0b2a-4d89-99a4-2455c2f181bb	2	like
4e57a35f-3047-4c47-b1ad-d1eeeb372f79	1	like
4e57a35f-3047-4c47-b1ad-d1eeeb372f79	2	like
ed628ab7-a0d5-4165-95ea-b10856b59c9f	1	dislike
db402d51-6c2e-4d1a-9719-f7b422df2297	1	like
db402d51-6c2e-4d1a-9719-f7b422df2297	2	like
b8386436-8deb-4799-a4c1-b9a1930e0b23	1	like
b8386436-8deb-4799-a4c1-b9a1930e0b23	2	like
cf80f59f-1c74-4be9-aafc-e2a57bc9b5c4	1	like
cf80f59f-1c74-4be9-aafc-e2a57bc9b5c4	2	like
0bccb2ee-55e2-4618-9497-f0a55d57e242	1	like
0bccb2ee-55e2-4618-9497-f0a55d57e242	2	like
0e7d7ff5-e6f0-4562-b6ec-5502414d7767	1	like
0e7d7ff5-e6f0-4562-b6ec-5502414d7767	2	like
e2fe9ea2-4877-4126-b1f1-2553b42a87b8	1	like
e2fe9ea2-4877-4126-b1f1-2553b42a87b8	2	like
0828083e-077b-4882-a345-27bae9a654ba	1	like
0828083e-077b-4882-a345-27bae9a654ba	2	like
be172980-37ff-480c-82cb-adc7baab4ee4	1	like
be172980-37ff-480c-82cb-adc7baab4ee4	2	like
576ba991-7ad7-4b12-b7be-7025258f428a	1	like
576ba991-7ad7-4b12-b7be-7025258f428a	2	like
2424e2d1-1b2f-4c02-af85-e35ec07643f8	1	like
2424e2d1-1b2f-4c02-af85-e35ec07643f8	2	like
ccfc5beb-f7c5-486e-a4f8-5a589380cd83	1	like
ccfc5beb-f7c5-486e-a4f8-5a589380cd83	2	like
98543712-e343-4c4f-b0b5-733246560801	1	like
98543712-e343-4c4f-b0b5-733246560801	2	like
21ed0ae9-abe4-4677-a2f4-275dbd46b101	1	like
21ed0ae9-abe4-4677-a2f4-275dbd46b101	2	like
a1f48d5b-9b5f-4908-b033-db38ed86f9dd	1	like
a1f48d5b-9b5f-4908-b033-db38ed86f9dd	2	like
12424ee7-df0e-4fa2-a7f1-8e2e68fc80e5	1	dislike
3660d59f-253a-425d-baa6-c0aa1ebd96ff	1	like
3660d59f-253a-425d-baa6-c0aa1ebd96ff	2	like
ccd4af3e-e910-408c-82e2-f938cb8e41bd	1	like
ccd4af3e-e910-408c-82e2-f938cb8e41bd	2	like
b59ba2a0-c080-464b-ae6f-d0f7343c124d	1	like
b59ba2a0-c080-464b-ae6f-d0f7343c124d	2	like
13dbb3d0-3f90-48cd-a0ad-e7312e9cf552	1	like
13dbb3d0-3f90-48cd-a0ad-e7312e9cf552	2	like
e1c99bb5-19f7-4e09-bbc4-8d94007a5826	1	like
e1c99bb5-19f7-4e09-bbc4-8d94007a5826	2	like
1f9045b0-e553-4ad5-9f6e-5278672a950c	1	like
1f9045b0-e553-4ad5-9f6e-5278672a950c	2	like
8e76ab55-6a76-4397-baf1-3484d8ce93e7	1	like
8e76ab55-6a76-4397-baf1-3484d8ce93e7	2	like
26edc9a7-de6e-4eff-bf62-66a5958745a6	1	like
26edc9a7-de6e-4eff-bf62-66a5958745a6	2	like
66f6b276-93f4-436a-9c46-a722e9d7869f	1	like
66f6b276-93f4-436a-9c46-a722e9d7869f	2	like
c958bad3-8d16-4f00-83ef-5e529fae3906	1	like
c958bad3-8d16-4f00-83ef-5e529fae3906	2	like
f1a0a31c-d761-41ba-b3df-98d7a4b00d28	1	like
f1a0a31c-d761-41ba-b3df-98d7a4b00d28	2	like
12fcb215-503a-45c0-82ea-3636b6c54977	1	like
12fcb215-503a-45c0-82ea-3636b6c54977	2	like
62b2fd63-3960-41e7-a425-2ff497075e0b	1	like
62b2fd63-3960-41e7-a425-2ff497075e0b	2	like
\.


--
-- Data for Name: rated_user_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.rated_user_tags (rating_id, user_tag_id, type) FROM stdin;
6a578d21-6efe-477b-908d-f9e3629496ec	1	like
6a578d21-6efe-477b-908d-f9e3629496ec	2	dislike
29a7e0d6-6333-42ac-ae44-db3f450c90fa	1	like
29a7e0d6-6333-42ac-ae44-db3f450c90fa	2	dislike
9d889d3b-2607-4132-a11c-a0b153c0999b	1	like
9d889d3b-2607-4132-a11c-a0b153c0999b	2	dislike
8859101b-2cdc-4b07-8a8f-ffad271b1855	1	like
8859101b-2cdc-4b07-8a8f-ffad271b1855	2	dislike
fc080d05-2343-4763-ab11-a166b41850cc	1	like
fc080d05-2343-4763-ab11-a166b41850cc	2	dislike
4019b6fa-b7d4-464e-b378-dc472efbcef2	1	dislike
0139b706-c919-4e25-9592-86a4890c1a58	1	like
0139b706-c919-4e25-9592-86a4890c1a58	2	dislike
ae1b7cc1-85d7-49b5-9042-3512690c6aaf	1	like
ae1b7cc1-85d7-49b5-9042-3512690c6aaf	2	dislike
f5ddf08d-a385-4473-891c-80c7d46a0274	7	like
f5ddf08d-a385-4473-891c-80c7d46a0274	8	dislike
254d1f5f-2181-4fcb-a0b0-1d6adf361e56	7	like
254d1f5f-2181-4fcb-a0b0-1d6adf361e56	8	dislike
4628a991-ed2a-499e-b86c-ad49cb1adf48	7	like
4628a991-ed2a-499e-b86c-ad49cb1adf48	8	dislike
80bbd313-d878-43be-a142-d2ae3bd8c0e3	7	like
80bbd313-d878-43be-a142-d2ae3bd8c0e3	8	dislike
2e87c2da-4cd1-41ae-92b6-901d8292a512	7	like
2e87c2da-4cd1-41ae-92b6-901d8292a512	8	dislike
3dc0af43-9427-40f9-8868-9f2bc2c86e62	7	dislike
d60ac2f6-d215-4d05-b6f7-2b90f70f4de4	7	like
d60ac2f6-d215-4d05-b6f7-2b90f70f4de4	8	dislike
9ad4baf3-bc4c-452b-ac5f-7aafa5213bca	7	like
9ad4baf3-bc4c-452b-ac5f-7aafa5213bca	8	dislike
31e7acc6-7a74-47b3-afb1-9ebd66cabbe1	13	like
31e7acc6-7a74-47b3-afb1-9ebd66cabbe1	14	dislike
11dc7539-a352-4068-b6cd-f04eec809095	13	like
11dc7539-a352-4068-b6cd-f04eec809095	14	dislike
febf8c41-95bc-4b78-8c98-24c6b26c04cc	13	like
febf8c41-95bc-4b78-8c98-24c6b26c04cc	14	dislike
d4c7bd95-bb2d-4bc3-924c-fb12030d143f	13	like
d4c7bd95-bb2d-4bc3-924c-fb12030d143f	14	dislike
6cc3d15e-0395-41fa-b48a-8169bb85a521	13	like
6cc3d15e-0395-41fa-b48a-8169bb85a521	14	dislike
651d26e5-f67f-4c15-b4c8-8296c2ab7576	13	like
651d26e5-f67f-4c15-b4c8-8296c2ab7576	14	dislike
252b0dab-940a-47c7-b036-3dd88dc59e6b	13	like
252b0dab-940a-47c7-b036-3dd88dc59e6b	14	dislike
12344a58-bd49-47d3-8ea0-5e8dd39bf4cb	19	like
12344a58-bd49-47d3-8ea0-5e8dd39bf4cb	20	dislike
24cfb54d-1bbc-417d-93f4-a92f9f2709a2	19	like
24cfb54d-1bbc-417d-93f4-a92f9f2709a2	20	dislike
cf33d37a-34e9-4699-bd49-6b7dd505a4e9	19	like
cf33d37a-34e9-4699-bd49-6b7dd505a4e9	20	dislike
6a2a550a-2770-404f-9e12-cd1ef583cafb	19	like
6a2a550a-2770-404f-9e12-cd1ef583cafb	20	dislike
97cb22e1-3da6-4251-88a2-c48f2de5907a	19	like
97cb22e1-3da6-4251-88a2-c48f2de5907a	20	dislike
2fa96b3a-cafc-4987-9e91-f9c7a6ef8fc0	19	dislike
1af5c1d6-2271-490a-b95d-493dea132eeb	19	like
1af5c1d6-2271-490a-b95d-493dea132eeb	20	dislike
088580d7-6235-4ae2-9c5a-879a06992ca2	19	like
088580d7-6235-4ae2-9c5a-879a06992ca2	20	dislike
4160e704-e9c7-46c2-bee6-ccf228bb8f49	25	like
4160e704-e9c7-46c2-bee6-ccf228bb8f49	26	dislike
bfe4127b-d196-4989-a4cb-2af2afb9b154	25	like
bfe4127b-d196-4989-a4cb-2af2afb9b154	26	dislike
5de81ca0-530a-4286-9798-bca2e36ef2c5	25	like
5de81ca0-530a-4286-9798-bca2e36ef2c5	26	dislike
ed3492df-d5f7-42aa-8e28-9c6f9140bee7	25	like
ed3492df-d5f7-42aa-8e28-9c6f9140bee7	26	dislike
ec0627d8-aa97-4d03-9884-fecc0491eb35	25	like
ec0627d8-aa97-4d03-9884-fecc0491eb35	26	dislike
e33441d7-d179-4197-8833-6d44bbd28e5e	25	like
e33441d7-d179-4197-8833-6d44bbd28e5e	26	dislike
5a33e2ed-868d-45b7-9415-4a162cc586b1	25	like
5a33e2ed-868d-45b7-9415-4a162cc586b1	26	dislike
\.


--
-- Data for Name: services; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.services (id, client_user_id, client_user_rating_id, client_user_lat, client_user_lon, profile_id, profile_owner_id, profile_rating_id, profile_user_lat, profile_user_lon, distance_between_users, trusted_distance, created_at, updated_at, updated_by) FROM stdin;
54131916-4b8a-4137-b93a-98c1c23d99d3	67a06967-6aa1-48d6-96de-61c496386a28	6a578d21-6efe-477b-908d-f9e3629496ec	43.25977	76.93525	c5b40786-9fae-4599-9b8e-72e5262d9356	2163ce47-18fe-49a8-a695-91c021c053ec	f2ce913a-1626-4878-b012-5c67c22b81dd	43.25988	76.9346	0.05393565459915141	t	2024-10-19 16:30:39.618404	2024-10-19 16:30:39.635781	2163ce47-18fe-49a8-a695-91c021c053ec
fa4c7daf-3a1a-4eac-abda-7429640addba	d2cfca70-84a0-40db-8593-5bfe52f18ac8	29a7e0d6-6333-42ac-ae44-db3f450c90fa	43.25977	76.93525	ca8bd72a-ddb7-4365-bd85-d76733a922c4	f508c6ad-b68c-4fb1-8af0-1d71e35ba08f	dca12033-0383-4d2d-bc5c-a9704bcd7ef1	43.25988	76.9346	0.05393565459915141	t	2024-10-19 16:30:39.801637	2024-10-19 16:30:39.803925	f508c6ad-b68c-4fb1-8af0-1d71e35ba08f
21b1d166-bd92-4495-84f2-afc684c534c5	702d39b1-4e20-434b-b51f-d73c72d740b3	9d889d3b-2607-4132-a11c-a0b153c0999b	43.25977	76.93525	4ab29639-f7a5-498e-899c-a8590c493947	d51586dd-9251-4e3d-8943-3ba3187c3c67	09b2c6b3-1ded-48b5-b8a4-052fccdd5cde	43.25988	76.9346	0.05393565459915141	t	2024-10-19 16:30:40.06891	2024-10-19 16:30:40.071523	d51586dd-9251-4e3d-8943-3ba3187c3c67
97060b58-ee59-45c3-aebf-dfe997702d42	f33e64e2-43bd-4189-ac6b-53f3e757b158	8859101b-2cdc-4b07-8a8f-ffad271b1855	43.25977	76.93525	e82afe50-a1e3-40b9-9cfd-953c82e88539	77a4ca8e-cb02-4289-878c-4d5e1a55bb30	23eade23-b8b0-431a-aaba-b18d5067c2b9	43.25988	76.9346	0.05393565459915141	t	2024-10-19 16:30:40.244453	2024-10-19 16:30:40.24639	77a4ca8e-cb02-4289-878c-4d5e1a55bb30
219ed2ff-9310-4310-9556-1f73d62c4c7b	c580fe9a-0996-4c30-b9cb-d58b45f1d686	fc080d05-2343-4763-ab11-a166b41850cc	43.25977	76.93525	946b8cca-e1fd-4bb2-9a00-126e4b423eaf	0210b4cf-c38c-46a6-a880-a0e40dd383a9	b78e0e3c-fa71-4a9e-b6a4-6ffd2fb2111d	43.25988	76.9346	0.05393565459915141	t	2024-10-19 16:30:40.501366	2024-10-19 16:30:40.50319	0210b4cf-c38c-46a6-a880-a0e40dd383a9
00332c05-cb87-49d9-840e-bababe9f50ca	f5c752ad-94c9-41f1-85c9-ac18e58a264a	4019b6fa-b7d4-464e-b378-dc472efbcef2	43.25977	76.93525	82194403-cdaf-42f6-8d11-129cd6ed716d	1de7291e-9f52-429b-b661-7a9f6c769477	eff62ab1-41bb-458e-9702-31cdccdffd91	43.25988	76.9346	0.05393565459915141	t	2024-10-19 16:30:40.668851	2024-10-19 16:30:40.670698	1de7291e-9f52-429b-b661-7a9f6c769477
281342f7-9538-45a2-a2e2-236cf8a9ebc2	a0cde28b-a282-4c74-98a0-f62fbd4d72ec	0139b706-c919-4e25-9592-86a4890c1a58	43.25977	76.93525	016bbc0f-968e-4cb6-969f-507ebc87d449	d7e38c6d-5058-4ffb-9c3d-ea0192a590ac	6d919a0b-a645-4582-a1ad-654328ce1260	43.25988	76.9346	0.05393565459915141	t	2024-10-19 16:30:40.925154	2024-10-19 16:30:40.9269	d7e38c6d-5058-4ffb-9c3d-ea0192a590ac
4906c968-4067-42f8-96d0-1a0acdfcbf57	b0d66f21-a91c-4137-9f52-4b6f5ec16083	ae1b7cc1-85d7-49b5-9042-3512690c6aaf	43.25977	76.93525	87f8f0d8-7619-4487-b8b4-cba515d4e35a	fa96765f-abe4-43bc-82fa-8a0137908029	52a88651-0b2a-4d89-99a4-2455c2f181bb	43.25988	76.9346	0.05393565459915141	t	2024-10-19 16:30:41.093931	2024-10-19 16:30:41.095922	fa96765f-abe4-43bc-82fa-8a0137908029
6d970b28-94a6-4eb3-b12e-9d1d05d45285	99b46525-ee48-43d3-810d-789da483c7c7	f5ddf08d-a385-4473-891c-80c7d46a0274	43.25977	76.93525	6712ac3e-8acf-4b96-a382-69c86f2a7de0	6890ef0b-0091-4d2d-b25f-f8800b577ea3	4e57a35f-3047-4c47-b1ad-d1eeeb372f79	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:40.58731	2024-10-19 18:50:40.609108	6890ef0b-0091-4d2d-b25f-f8800b577ea3
1fdeadae-0f63-41c2-ba84-439e33f3631e	67bce66f-3d0e-4f27-b464-1ccd3999191f	254d1f5f-2181-4fcb-a0b0-1d6adf361e56	43.25977	76.93525	ddcb59d6-9a82-4ab6-918e-dce2097d4603	ef6d7a38-93ec-4a62-aec0-63027cb431ea	ed628ab7-a0d5-4165-95ea-b10856b59c9f	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:40.770388	2024-10-19 18:50:40.772754	ef6d7a38-93ec-4a62-aec0-63027cb431ea
8d615917-fb56-4593-bf55-49f6c99d5f70	b15c4b1b-4aae-426a-b6d3-dfa589b03705	4628a991-ed2a-499e-b86c-ad49cb1adf48	43.25977	76.93525	699a5205-affc-404a-8390-3bf8b6d1b8a0	6e99a07e-2ac3-466a-a00c-419b12f3f2e7	db402d51-6c2e-4d1a-9719-f7b422df2297	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:41.02196	2024-10-19 18:50:41.024222	6e99a07e-2ac3-466a-a00c-419b12f3f2e7
378899e8-deef-4097-9177-075acfedf8d2	e150180b-d35d-44b7-84c0-88d32c1afb3f	80bbd313-d878-43be-a142-d2ae3bd8c0e3	43.25977	76.93525	add0bf9a-b9f9-4203-8106-56ef6c2ec13e	3ae0a69d-27f8-4cbe-8e5b-0f185dafe009	b8386436-8deb-4799-a4c1-b9a1930e0b23	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:41.180069	2024-10-19 18:50:41.182088	3ae0a69d-27f8-4cbe-8e5b-0f185dafe009
6f23717f-51b9-4ba5-a298-0939715f218d	ed590c79-c544-4792-8b90-d138dfd61ffa	2e87c2da-4cd1-41ae-92b6-901d8292a512	43.25977	76.93525	04ad1331-165d-4cab-bade-148029bc121e	e4166f80-9a47-4bd2-a740-27c8fb6a48cb	cf80f59f-1c74-4be9-aafc-e2a57bc9b5c4	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:41.421815	2024-10-19 18:50:41.423618	e4166f80-9a47-4bd2-a740-27c8fb6a48cb
6eb98269-2f29-47f6-b9c9-4936358126a4	ff97ae71-6c30-428a-96c3-654c9d56ba89	3dc0af43-9427-40f9-8868-9f2bc2c86e62	43.25977	76.93525	7961adc7-2bf6-4af8-abbb-25086b4791f6	b5cecb52-7877-4029-afa5-e677a1c2ec7d	0bccb2ee-55e2-4618-9497-f0a55d57e242	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:41.574731	2024-10-19 18:50:41.576605	b5cecb52-7877-4029-afa5-e677a1c2ec7d
df575a7c-6d1c-445a-a0bf-97c8519a46c3	aa5552e1-693c-4c54-9db4-ee304e6af9f1	d60ac2f6-d215-4d05-b6f7-2b90f70f4de4	43.25977	76.93525	3dd222e0-da73-4f9a-be39-bf713d8085ea	0ee8284c-e553-4816-b928-2415ab336d72	0e7d7ff5-e6f0-4562-b6ec-5502414d7767	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:41.813595	2024-10-19 18:50:41.815421	0ee8284c-e553-4816-b928-2415ab336d72
2898b892-270c-469a-8cb5-9c861deba07d	da99ff67-9b25-4b56-835b-633e22c66870	9ad4baf3-bc4c-452b-ac5f-7aafa5213bca	43.25977	76.93525	9f5f226e-60e3-44d1-a4dd-a7ed8fc3aed6	a2bc57d6-4555-4230-ac8e-712bdfbeaea1	e2fe9ea2-4877-4126-b1f1-2553b42a87b8	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:41.971903	2024-10-19 18:50:41.973586	a2bc57d6-4555-4230-ac8e-712bdfbeaea1
2801a779-5e6c-4166-9a28-f38d86e02f4d	210887f4-dace-452e-8b37-9acfdeb9ff3b	\N	43.25977	76.93525	d35e126d-0e12-460c-a820-46730076a2cb	5d0c334e-272c-43b3-a361-4726a2dfbf85	\N	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:43.587565	2024-10-19 18:50:43.603253	210887f4-dace-452e-8b37-9acfdeb9ff3b
dc4dd43b-b2fd-49d4-b888-e2ef59f5e8e9	bc0bed96-18c4-49a5-87fa-3eadf3e555da	\N	43.25977	76.93525	dfdd7f7a-7fbb-4e22-83b9-ea031dba83d8	100bc084-ae37-42db-938d-c71d32dbb89c	\N	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:43.755742	2024-10-19 18:50:43.75658	100bc084-ae37-42db-938d-c71d32dbb89c
c24c24ce-7624-48ef-9043-1b73c3f01677	3c16965b-a69e-4db3-8728-ed11d13abf06	31e7acc6-7a74-47b3-afb1-9ebd66cabbe1	43.25977	76.93525	fe398ea0-93c8-4cc6-a1ab-3c39f11549a8	2201e804-7513-4100-93d1-e3f247e49e84	0828083e-077b-4882-a345-27bae9a654ba	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:43.910127	2024-10-19 18:50:43.914592	2201e804-7513-4100-93d1-e3f247e49e84
c0f16764-33cd-40c1-8b18-3b9229938e52	85f3becf-2826-47c5-85c4-1d58d090ad4f	11dc7539-a352-4068-b6cd-f04eec809095	43.25977	76.93525	38e062b9-e847-4e78-86d3-50ede8cd63e4	579affde-faf8-4b01-9727-d964d5ec456a	be172980-37ff-480c-82cb-adc7baab4ee4	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:44.16239	2024-10-19 18:50:44.164356	579affde-faf8-4b01-9727-d964d5ec456a
556b03a9-d42a-4a71-89e3-5c31da6c4077	f3bcd7a7-d673-4915-853e-d6d5b5eadb5d	febf8c41-95bc-4b78-8c98-24c6b26c04cc	43.25977	76.93525	d10c5aca-e410-45b4-9442-0ce5b6f9ce46	d4db4b08-437c-4ebc-a696-c4ad35e75734	576ba991-7ad7-4b12-b7be-7025258f428a	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:44.392438	2024-10-19 18:50:44.394689	d4db4b08-437c-4ebc-a696-c4ad35e75734
8b1d4f23-535a-4335-aaea-bf065974a10d	791fb853-9bb3-4195-8d57-b51325b5ce6b	d4c7bd95-bb2d-4bc3-924c-fb12030d143f	43.25977	76.93525	427578f5-ca68-4609-8531-7893d8c6b77f	f4c07801-8c69-47f8-9719-e439d500639c	2424e2d1-1b2f-4c02-af85-e35ec07643f8	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:44.621154	2024-10-19 18:50:44.624918	f4c07801-8c69-47f8-9719-e439d500639c
ea40d566-c6a8-4eda-8274-85b97848d449	b2b3786b-0b30-4ce8-9788-17f91bc783cd	6cc3d15e-0395-41fa-b48a-8169bb85a521	43.25977	76.93525	c80739e5-9f0c-47b6-b81c-f231b5d65607	9bee80eb-4c99-4823-90ac-c056be769100	ccfc5beb-f7c5-486e-a4f8-5a589380cd83	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:44.852568	2024-10-19 18:50:44.854745	9bee80eb-4c99-4823-90ac-c056be769100
2bf380a8-cd36-44eb-ab88-118f7a320247	b930e0f8-e52e-4f24-b4a0-6097fdaea051	651d26e5-f67f-4c15-b4c8-8296c2ab7576	43.25977	76.93525	4d49b906-00d7-43b8-857f-1e2cccc742a4	524b58ed-10f5-4921-82bf-852780979138	98543712-e343-4c4f-b0b5-733246560801	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:45.082452	2024-10-19 18:50:45.08431	524b58ed-10f5-4921-82bf-852780979138
9dfccf88-de46-4fb6-b1b0-d3bdf3c2fbc9	1f1ca188-a8f0-4936-bcb3-a880dd72962f	252b0dab-940a-47c7-b036-3dd88dc59e6b	43.25977	76.93525	bc88f881-93d5-47ba-92b4-02a5cf706bf1	846c7558-68d2-4074-aad8-1a740629a2b3	21ed0ae9-abe4-4677-a2f4-275dbd46b101	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:50:45.312865	2024-10-19 18:50:45.314514	846c7558-68d2-4074-aad8-1a740629a2b3
3301fce0-e1f9-4b92-8387-7b081ef7fb6f	1e9feac0-baac-497d-936e-ef90e7ae541f	12344a58-bd49-47d3-8ea0-5e8dd39bf4cb	43.25977	76.93525	ac0e613a-0f21-4648-a4f7-907cbe5c931d	affe57b1-aa52-41ba-bbdc-572b6d53bc0e	a1f48d5b-9b5f-4908-b033-db38ed86f9dd	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:21.545517	2024-10-19 18:51:21.566409	affe57b1-aa52-41ba-bbdc-572b6d53bc0e
a0fe791b-4c0d-4cf5-8ba7-bc2e069ff4ca	85a470d9-39cb-4337-918f-214da397c062	24cfb54d-1bbc-417d-93f4-a92f9f2709a2	43.25977	76.93525	787c3dc1-b07e-40cc-aefe-f6e917ea9887	4c274022-dbde-4d41-8673-5849b20c4052	12424ee7-df0e-4fa2-a7f1-8e2e68fc80e5	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:21.729976	2024-10-19 18:51:21.732148	4c274022-dbde-4d41-8673-5849b20c4052
9834386e-7e61-418e-a068-b14510d0e2f2	66dbbf85-70ee-42e9-9b09-4ce5ff29cf41	cf33d37a-34e9-4699-bd49-6b7dd505a4e9	43.25977	76.93525	262f6dd4-9c45-4e2e-a7f3-1e5648cf14db	f7d194a0-166b-4cb4-8c89-71c519d07d34	3660d59f-253a-425d-baa6-c0aa1ebd96ff	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:21.986461	2024-10-19 18:51:21.989758	f7d194a0-166b-4cb4-8c89-71c519d07d34
a547e713-9e53-4f21-bcc5-7d1b706e4fcf	15d77ba6-ed63-45c8-8f8f-44d53f65bb2e	6a2a550a-2770-404f-9e12-cd1ef583cafb	43.25977	76.93525	972be0ec-33f4-441b-96a5-ea689f8235fd	20a76176-0042-4923-a381-2aab9954e679	ccd4af3e-e910-408c-82e2-f938cb8e41bd	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:22.146478	2024-10-19 18:51:22.149024	20a76176-0042-4923-a381-2aab9954e679
0a2643d3-b5f1-411a-9d21-adb7b4ffef58	aab90a73-7321-44f6-bfcc-76974bfa8448	97cb22e1-3da6-4251-88a2-c48f2de5907a	43.25977	76.93525	9277be55-fc4b-4e84-adda-468f421d675a	6e379ad4-5b11-4aa2-8c03-e6d12e032755	b59ba2a0-c080-464b-ae6f-d0f7343c124d	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:22.393451	2024-10-19 18:51:22.395304	6e379ad4-5b11-4aa2-8c03-e6d12e032755
cc71b8af-4cab-45f1-8d3d-be87e03a0feb	58fe0b33-ff2b-424d-925c-79c348ea6c2e	2fa96b3a-cafc-4987-9e91-f9c7a6ef8fc0	43.25977	76.93525	8adddeca-1575-4103-8899-2b834867e8ac	efc3d895-a248-45ba-9841-b9d638a53175	13dbb3d0-3f90-48cd-a0ad-e7312e9cf552	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:22.554326	2024-10-19 18:51:22.55679	efc3d895-a248-45ba-9841-b9d638a53175
823cfd22-d946-4b85-8a16-fbed87f473e6	3b91a9a8-c3cc-44fc-b73c-46d82d80183a	1af5c1d6-2271-490a-b95d-493dea132eeb	43.25977	76.93525	d3d5df5c-520e-4709-be51-130a184bdec5	1628db9d-b81a-4922-a905-caa0d4c16a48	e1c99bb5-19f7-4e09-bbc4-8d94007a5826	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:22.800098	2024-10-19 18:51:22.802479	1628db9d-b81a-4922-a905-caa0d4c16a48
5a6c047b-f262-4159-a300-a33d5dd9d1df	d1224a95-908c-4bbf-89e0-3a1842687c10	088580d7-6235-4ae2-9c5a-879a06992ca2	43.25977	76.93525	83747cf3-361e-46b7-b599-302b45d375b8	a42aa967-9a5c-48ee-8ae7-f4334ea0b56b	1f9045b0-e553-4ad5-9f6e-5278672a950c	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:22.958274	2024-10-19 18:51:22.960371	a42aa967-9a5c-48ee-8ae7-f4334ea0b56b
5d58e41c-077a-4834-9185-d5b9c2bf6209	5a977610-19f6-4022-904f-93edfc72ea70	\N	43.25977	76.93525	bdd286df-91ba-4260-9332-c76a65d5154b	6490cf88-e23a-4be5-855f-91a59e00fef6	\N	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:24.632247	2024-10-19 18:51:24.646492	5a977610-19f6-4022-904f-93edfc72ea70
03659880-2364-42d6-ac24-3467b86df437	387d7a39-c15e-4e5d-aa56-c0ff4c9cef6e	\N	43.25977	76.93525	2b14bfc9-8429-4b36-b58b-14a46fbb53b8	f23e28e0-1095-494b-9778-2ce1289423ea	\N	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:24.800099	2024-10-19 18:51:24.800774	f23e28e0-1095-494b-9778-2ce1289423ea
a4389442-20bf-4cb8-ab9b-4ce41e8ed7b2	3d632ec6-6633-4165-93cf-f160eee0b246	4160e704-e9c7-46c2-bee6-ccf228bb8f49	43.25977	76.93525	539942a9-161f-4503-a249-325393d15ce3	6f5beca5-2aad-4533-8d42-7cfebf2b0532	8e76ab55-6a76-4397-baf1-3484d8ce93e7	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:24.952265	2024-10-19 18:51:24.957054	6f5beca5-2aad-4533-8d42-7cfebf2b0532
8fc5fe6d-341d-423d-b6d2-1339960f1721	bbcc298f-3341-4ba0-b858-589141d74918	bfe4127b-d196-4989-a4cb-2af2afb9b154	43.25977	76.93525	a316018d-0ba8-4804-b14b-16f18284dbbb	a5f7b81f-00f7-4e1e-8526-b27cde367f21	26edc9a7-de6e-4eff-bf62-66a5958745a6	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:25.189531	2024-10-19 18:51:25.19203	a5f7b81f-00f7-4e1e-8526-b27cde367f21
6ae6a042-16b6-4793-bc4c-ad701b8c5f2a	2f22fa46-cd97-460f-b25f-e45373c99caa	5de81ca0-530a-4286-9798-bca2e36ef2c5	43.25977	76.93525	7856ff4c-15c8-4976-8018-940a98ae8e35	a4322a31-4340-4222-88b3-683d0e210b7e	66f6b276-93f4-436a-9c46-a722e9d7869f	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:25.423856	2024-10-19 18:51:25.425835	a4322a31-4340-4222-88b3-683d0e210b7e
6cd5378b-c9ec-44aa-beb7-701ba4f9e49a	5df86b9d-d643-49fe-9780-6c69cf3d981a	ed3492df-d5f7-42aa-8e28-9c6f9140bee7	43.25977	76.93525	e55d9ff0-4645-4812-8159-91a2fe1cedcb	4e08df8a-683c-4565-8b39-700c1e4275ad	c958bad3-8d16-4f00-83ef-5e529fae3906	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:25.659184	2024-10-19 18:51:25.662157	4e08df8a-683c-4565-8b39-700c1e4275ad
218012d7-fdb8-466f-b4aa-49fd77c9a7a6	c6c7c2f2-6ad6-43d1-a883-94a4fee81a06	ec0627d8-aa97-4d03-9884-fecc0491eb35	43.25977	76.93525	4877bd92-80de-4908-a2fa-934139142601	0a1a6ce6-0085-46f1-9803-7175849a1c28	f1a0a31c-d761-41ba-b3df-98d7a4b00d28	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:25.887358	2024-10-19 18:51:25.889425	0a1a6ce6-0085-46f1-9803-7175849a1c28
a6e08b54-032f-4daa-993e-54e9f4ffb6e8	08d9ba05-a11d-445f-9b3e-9caca6712316	e33441d7-d179-4197-8833-6d44bbd28e5e	43.25977	76.93525	48969bc2-fa39-4032-9a9f-5b95dabb2367	17baa83f-62a2-4662-9a23-404a5d3e684d	12fcb215-503a-45c0-82ea-3636b6c54977	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:26.121385	2024-10-19 18:51:26.123637	17baa83f-62a2-4662-9a23-404a5d3e684d
d46214d1-ba49-4c2f-b772-33f28dea1e29	a0f0c4ce-1e11-486e-bcae-d4b730eb3602	5a33e2ed-868d-45b7-9415-4a162cc586b1	43.25977	76.93525	302e43c7-4616-4a9c-8c0a-ff6f6846c7c5	aa8bc7e4-be00-4aee-9ad2-85c47f2395b3	62b2fd63-3960-41e7-a425-2ff497075e0b	43.25988	76.9346	0.05393565459915141	t	2024-10-19 18:51:26.358537	2024-10-19 18:51:26.360495	aa8bc7e4-be00-4aee-9ad2-85c47f2395b3
\.


--
-- Data for Name: user_ratings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_ratings (id, service_id, user_id, review_text_visible, review, score, created_at, updated_at, updated_by) FROM stdin;
6a578d21-6efe-477b-908d-f9e3629496ec	54131916-4b8a-4137-b93a-98c1c23d99d3	67a06967-6aa1-48d6-96de-61c496386a28	t	I liked the client! He is very kind	5	2024-10-19 16:30:39.618404	2024-10-19 16:30:39.618404	00000000-0000-0000-0000-000000000000
29a7e0d6-6333-42ac-ae44-db3f450c90fa	fa4c7daf-3a1a-4eac-abda-7429640addba	d2cfca70-84a0-40db-8593-5bfe52f18ac8	t	I liked the client! He is very kind	5	2024-10-19 16:30:39.801637	2024-10-19 16:30:39.801637	00000000-0000-0000-0000-000000000000
9d889d3b-2607-4132-a11c-a0b153c0999b	21b1d166-bd92-4495-84f2-afc684c534c5	702d39b1-4e20-434b-b51f-d73c72d740b3	t	I liked the client! He is very kind	5	2024-10-19 16:30:40.06891	2024-10-19 16:30:40.06891	00000000-0000-0000-0000-000000000000
8859101b-2cdc-4b07-8a8f-ffad271b1855	97060b58-ee59-45c3-aebf-dfe997702d42	f33e64e2-43bd-4189-ac6b-53f3e757b158	t	I liked the client! He is very kind	5	2024-10-19 16:30:40.244453	2024-10-19 16:30:40.244453	00000000-0000-0000-0000-000000000000
fc080d05-2343-4763-ab11-a166b41850cc	219ed2ff-9310-4310-9556-1f73d62c4c7b	c580fe9a-0996-4c30-b9cb-d58b45f1d686	t	I liked the client! He is very kind	5	2024-10-17 12:30:40.506754	2024-10-19 16:30:40.501366	00000000-0000-0000-0000-000000000000
ed3492df-d5f7-42aa-8e28-9c6f9140bee7	6cd5378b-c9ec-44aa-beb7-701ba4f9e49a	5df86b9d-d643-49fe-9780-6c69cf3d981a	t	I liked the client! He is very kind	5	2024-10-19 18:51:25.659184	2024-10-19 18:51:25.659184	00000000-0000-0000-0000-000000000000
4019b6fa-b7d4-464e-b378-dc472efbcef2	00332c05-cb87-49d9-840e-bababe9f50ca	f5c752ad-94c9-41f1-85c9-ac18e58a264a	t	I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life	1	2024-10-18 13:30:40.674474	2024-10-19 16:30:40.680595	00000000-0000-0000-0000-000000000000
0139b706-c919-4e25-9592-86a4890c1a58	281342f7-9538-45a2-a2e2-236cf8a9ebc2	a0cde28b-a282-4c74-98a0-f62fbd4d72ec	t	I liked the client! He is very kind	5	2024-10-19 16:30:40.925154	2024-10-19 16:30:40.925154	00000000-0000-0000-0000-000000000000
ae1b7cc1-85d7-49b5-9042-3512690c6aaf	4906c968-4067-42f8-96d0-1a0acdfcbf57	b0d66f21-a91c-4137-9f52-4b6f5ec16083	f	I liked the client! He is very kind	5	2024-10-19 16:30:41.093931	2024-10-19 16:30:41.102481	00000000-0000-0000-0000-000000000000
f5ddf08d-a385-4473-891c-80c7d46a0274	6d970b28-94a6-4eb3-b12e-9d1d05d45285	99b46525-ee48-43d3-810d-789da483c7c7	t	I liked the client! He is very kind	5	2024-10-19 18:50:40.58731	2024-10-19 18:50:40.58731	00000000-0000-0000-0000-000000000000
254d1f5f-2181-4fcb-a0b0-1d6adf361e56	1fdeadae-0f63-41c2-ba84-439e33f3631e	67bce66f-3d0e-4f27-b464-1ccd3999191f	t	I liked the client! He is very kind	5	2024-10-19 18:50:40.770388	2024-10-19 18:50:40.770388	00000000-0000-0000-0000-000000000000
4628a991-ed2a-499e-b86c-ad49cb1adf48	8d615917-fb56-4593-bf55-49f6c99d5f70	b15c4b1b-4aae-426a-b6d3-dfa589b03705	t	I liked the client! He is very kind	5	2024-10-19 18:50:41.02196	2024-10-19 18:50:41.02196	00000000-0000-0000-0000-000000000000
80bbd313-d878-43be-a142-d2ae3bd8c0e3	378899e8-deef-4097-9177-075acfedf8d2	e150180b-d35d-44b7-84c0-88d32c1afb3f	t	I liked the client! He is very kind	5	2024-10-19 18:50:41.180069	2024-10-19 18:50:41.180069	00000000-0000-0000-0000-000000000000
2e87c2da-4cd1-41ae-92b6-901d8292a512	6f23717f-51b9-4ba5-a298-0939715f218d	ed590c79-c544-4792-8b90-d138dfd61ffa	t	I liked the client! He is very kind	5	2024-10-17 14:50:41.426506	2024-10-19 18:50:41.421815	00000000-0000-0000-0000-000000000000
ec0627d8-aa97-4d03-9884-fecc0491eb35	218012d7-fdb8-466f-b4aa-49fd77c9a7a6	c6c7c2f2-6ad6-43d1-a883-94a4fee81a06	t	I liked the client! He is very kind	5	2024-10-19 18:51:25.887358	2024-10-19 18:51:25.887358	00000000-0000-0000-0000-000000000000
3dc0af43-9427-40f9-8868-9f2bc2c86e62	6eb98269-2f29-47f6-b9c9-4936358126a4	ff97ae71-6c30-428a-96c3-654c9d56ba89	t	I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life	1	2024-10-18 15:50:41.579568	2024-10-19 18:50:41.584671	00000000-0000-0000-0000-000000000000
d60ac2f6-d215-4d05-b6f7-2b90f70f4de4	df575a7c-6d1c-445a-a0bf-97c8519a46c3	aa5552e1-693c-4c54-9db4-ee304e6af9f1	t	I liked the client! He is very kind	5	2024-10-19 18:50:41.813595	2024-10-19 18:50:41.813595	00000000-0000-0000-0000-000000000000
9ad4baf3-bc4c-452b-ac5f-7aafa5213bca	2898b892-270c-469a-8cb5-9c861deba07d	da99ff67-9b25-4b56-835b-633e22c66870	f	I liked the client! He is very kind	5	2024-10-19 18:50:41.971903	2024-10-19 18:50:41.97982	00000000-0000-0000-0000-000000000000
31e7acc6-7a74-47b3-afb1-9ebd66cabbe1	c24c24ce-7624-48ef-9043-1b73c3f01677	3c16965b-a69e-4db3-8728-ed11d13abf06	t	I liked the client! He is very kind	5	2024-10-19 18:50:43.910127	2024-10-19 18:50:43.910127	00000000-0000-0000-0000-000000000000
11dc7539-a352-4068-b6cd-f04eec809095	c0f16764-33cd-40c1-8b18-3b9229938e52	85f3becf-2826-47c5-85c4-1d58d090ad4f	t	I liked the client! He is very kind	5	2024-10-19 18:50:44.16239	2024-10-19 18:50:44.16239	00000000-0000-0000-0000-000000000000
febf8c41-95bc-4b78-8c98-24c6b26c04cc	556b03a9-d42a-4a71-89e3-5c31da6c4077	f3bcd7a7-d673-4915-853e-d6d5b5eadb5d	t	I liked the client! He is very kind	5	2024-10-19 18:50:44.392438	2024-10-19 18:50:44.392438	00000000-0000-0000-0000-000000000000
d4c7bd95-bb2d-4bc3-924c-fb12030d143f	8b1d4f23-535a-4335-aaea-bf065974a10d	791fb853-9bb3-4195-8d57-b51325b5ce6b	t	I liked the client! He is very kind	5	2024-10-19 18:50:44.621154	2024-10-19 18:50:44.621154	00000000-0000-0000-0000-000000000000
6cc3d15e-0395-41fa-b48a-8169bb85a521	ea40d566-c6a8-4eda-8274-85b97848d449	b2b3786b-0b30-4ce8-9788-17f91bc783cd	t	I liked the client! He is very kind	5	2024-10-19 18:50:44.852568	2024-10-19 18:50:44.852568	00000000-0000-0000-0000-000000000000
651d26e5-f67f-4c15-b4c8-8296c2ab7576	2bf380a8-cd36-44eb-ab88-118f7a320247	b930e0f8-e52e-4f24-b4a0-6097fdaea051	t	I liked the client! He is very kind	5	2024-10-19 18:50:45.082452	2024-10-19 18:50:45.082452	00000000-0000-0000-0000-000000000000
252b0dab-940a-47c7-b036-3dd88dc59e6b	9dfccf88-de46-4fb6-b1b0-d3bdf3c2fbc9	1f1ca188-a8f0-4936-bcb3-a880dd72962f	t	I liked the client! He is very kind	5	2024-10-19 18:50:45.312865	2024-10-19 18:50:45.312865	00000000-0000-0000-0000-000000000000
12344a58-bd49-47d3-8ea0-5e8dd39bf4cb	3301fce0-e1f9-4b92-8387-7b081ef7fb6f	1e9feac0-baac-497d-936e-ef90e7ae541f	t	I liked the client! He is very kind	5	2024-10-19 18:51:21.545517	2024-10-19 18:51:21.545517	00000000-0000-0000-0000-000000000000
24cfb54d-1bbc-417d-93f4-a92f9f2709a2	a0fe791b-4c0d-4cf5-8ba7-bc2e069ff4ca	85a470d9-39cb-4337-918f-214da397c062	t	I liked the client! He is very kind	5	2024-10-19 18:51:21.729976	2024-10-19 18:51:21.729976	00000000-0000-0000-0000-000000000000
cf33d37a-34e9-4699-bd49-6b7dd505a4e9	9834386e-7e61-418e-a068-b14510d0e2f2	66dbbf85-70ee-42e9-9b09-4ce5ff29cf41	t	I liked the client! He is very kind	5	2024-10-19 18:51:21.986461	2024-10-19 18:51:21.986461	00000000-0000-0000-0000-000000000000
6a2a550a-2770-404f-9e12-cd1ef583cafb	a547e713-9e53-4f21-bcc5-7d1b706e4fcf	15d77ba6-ed63-45c8-8f8f-44d53f65bb2e	t	I liked the client! He is very kind	5	2024-10-19 18:51:22.146478	2024-10-19 18:51:22.146478	00000000-0000-0000-0000-000000000000
97cb22e1-3da6-4251-88a2-c48f2de5907a	0a2643d3-b5f1-411a-9d21-adb7b4ffef58	aab90a73-7321-44f6-bfcc-76974bfa8448	t	I liked the client! He is very kind	5	2024-10-17 14:51:22.399086	2024-10-19 18:51:22.393451	00000000-0000-0000-0000-000000000000
e33441d7-d179-4197-8833-6d44bbd28e5e	a6e08b54-032f-4daa-993e-54e9f4ffb6e8	08d9ba05-a11d-445f-9b3e-9caca6712316	t	I liked the client! He is very kind	5	2024-10-19 18:51:26.121385	2024-10-19 18:51:26.121385	00000000-0000-0000-0000-000000000000
2fa96b3a-cafc-4987-9e91-f9c7a6ef8fc0	cc71b8af-4cab-45f1-8d3d-be87e03a0feb	58fe0b33-ff2b-424d-925c-79c348ea6c2e	t	I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life	1	2024-10-18 15:51:22.560377	2024-10-19 18:51:22.566338	00000000-0000-0000-0000-000000000000
1af5c1d6-2271-490a-b95d-493dea132eeb	823cfd22-d946-4b85-8a16-fbed87f473e6	3b91a9a8-c3cc-44fc-b73c-46d82d80183a	t	I liked the client! He is very kind	5	2024-10-19 18:51:22.800098	2024-10-19 18:51:22.800098	00000000-0000-0000-0000-000000000000
088580d7-6235-4ae2-9c5a-879a06992ca2	5a6c047b-f262-4159-a300-a33d5dd9d1df	d1224a95-908c-4bbf-89e0-3a1842687c10	f	I liked the client! He is very kind	5	2024-10-19 18:51:22.958274	2024-10-19 18:51:22.966428	00000000-0000-0000-0000-000000000000
4160e704-e9c7-46c2-bee6-ccf228bb8f49	a4389442-20bf-4cb8-ab9b-4ce41e8ed7b2	3d632ec6-6633-4165-93cf-f160eee0b246	t	I liked the client! He is very kind	5	2024-10-19 18:51:24.952265	2024-10-19 18:51:24.952265	00000000-0000-0000-0000-000000000000
bfe4127b-d196-4989-a4cb-2af2afb9b154	8fc5fe6d-341d-423d-b6d2-1339960f1721	bbcc298f-3341-4ba0-b858-589141d74918	t	I liked the client! He is very kind	5	2024-10-19 18:51:25.189531	2024-10-19 18:51:25.189531	00000000-0000-0000-0000-000000000000
5de81ca0-530a-4286-9798-bca2e36ef2c5	6ae6a042-16b6-4793-bc4c-ad701b8c5f2a	2f22fa46-cd97-460f-b25f-e45373c99caa	t	I liked the client! He is very kind	5	2024-10-19 18:51:25.423856	2024-10-19 18:51:25.423856	00000000-0000-0000-0000-000000000000
5a33e2ed-868d-45b7-9415-4a162cc586b1	d46214d1-ba49-4c2f-b772-33f28dea1e29	a0f0c4ce-1e11-486e-bcae-d4b730eb3602	t	I liked the client! He is very kind	5	2024-10-19 18:51:26.358537	2024-10-19 18:51:26.358537	00000000-0000-0000-0000-000000000000
\.


--
-- Data for Name: user_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_tags (id, name) FROM stdin;
1	Гигиена
2	Опрятность
3	Щедрость
4	Пунктуальность
5	Соблюдение границ
6	Общительность
7	Гигиена
8	Опрятность
9	Щедрость
10	Пунктуальность
11	Соблюдение границ
12	Общительность
13	Гигиена
14	Опрятность
15	Щедрость
16	Пунктуальность
17	Соблюдение границ
18	Общительность
19	Гигиена
20	Опрятность
21	Щедрость
22	Пунктуальность
23	Соблюдение границ
24	Общительность
25	Гигиена
26	Опрятность
27	Щедрость
28	Пунктуальность
29	Соблюдение границ
30	Общительность
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, phone, telegram_user_id, password, active, verified, created_at, updated_at, avatar, has_profile, tier, role) FROM stdin;
2163ce47-18fe-49a8-a695-91c021c053ec	test-KJFW1bQazL	17227136056	789867231083281794	$2a$10$FYrkdNjDzr946kJmm6qluulKebauvGIMHc.7OPHhNCuJ68fOH66Zu	t	f	2024-10-19 16:30:39.511568	2024-10-19 16:30:39.511568		f	basic	user
67a06967-6aa1-48d6-96de-61c496386a28	test-PBS3DUSnPz	32300135465	1449152910703401256	$2a$10$ZN9MipnN7ob0jL/.9AX/q.m0swsBf.GqYb0LMfpZirSbeHOnu5VT.	t	f	2024-10-19 16:30:39.593136	2024-10-19 16:30:39.593136		f	basic	user
f508c6ad-b68c-4fb1-8af0-1d71e35ba08f	test-wEFFRDGynI	49979457368	1704972454503113059	$2a$10$.M.wNBBtd41tnglpZA8nreerpP6I/OiOurb4QhNcnkHwdbhDT3kGy	t	f	2024-10-19 16:30:39.722766	2024-10-19 16:30:39.722766		f	basic	user
d2cfca70-84a0-40db-8593-5bfe52f18ac8	test-NwcFpBHAPv	53599784991	915789003975061156	$2a$10$0PugjP5JoKiL3vWgKKoAyuRlO6xNzgEor1.UyIxKjz6xEou/nwkry	t	f	2024-10-19 16:30:39.793242	2024-10-19 16:30:39.793242		f	basic	user
65db6c3b-a536-4150-a3d3-a139fc14b587	test-fWDdNg7QZ7	82647201603	8886637420322470702	$2a$10$JCzH8oiW5mRvTq9RDlFnEu1rVOVdV0BWfwZNIIF0xT6M2N1fzvWyG	t	f	2024-10-19 16:30:39.899936	2024-10-19 16:30:39.90233		f	guru	user
d51586dd-9251-4e3d-8943-3ba3187c3c67	test-DOsviVvSy0	62393288898	3635169983643141427	$2a$10$D0A4pffnvHBManXxLY5t0uD0ygdD9a6YAU.Rfo2VJ1KfRGsbSSdqa	t	f	2024-10-19 16:30:39.979936	2024-10-19 16:30:39.979936		f	basic	user
702d39b1-4e20-434b-b51f-d73c72d740b3	test-Sd8IrIzQRs	41953074371	6355829446908981714	$2a$10$079OlHozsOQd0dvWBDEDqup6.0gIICLyC.sN0GVOg21W.fwaai06y	t	f	2024-10-19 16:30:40.054226	2024-10-19 16:30:40.054226		f	basic	user
77a4ca8e-cb02-4289-878c-4d5e1a55bb30	test-q7tgo0UwbT	13799845184	8000417107047009002	$2a$10$9w9TWiwYcld59jOZ1ROPwefL2dBOoGmpjCVz9XFTqUQh9uI0uKpoK	t	f	2024-10-19 16:30:40.151563	2024-10-19 16:30:40.15502		f	expert	user
f33e64e2-43bd-4189-ac6b-53f3e757b158	test-6Xvts8qHLe	17628789933	8911893136879908841	$2a$10$oPG3QtLG8Qnji6VLQKCTHOCdX1KnSMmeDH.a..6IQnIsEAIfJfFtG	t	f	2024-10-19 16:30:40.230854	2024-10-19 16:30:40.230854		f	basic	user
dae6ac98-5881-4470-89d2-a2e8b9ccc84d	test-ApmGn1ux8C	11746863195	1120567894431005656	$2a$10$Ch7RAx1OmJHJj9LemHnWHeU67hHwkikKlwviNsoEZYxY3WKt2sh/a	t	f	2024-10-19 16:30:40.329801	2024-10-19 16:30:40.334598		f	guru	user
0210b4cf-c38c-46a6-a880-a0e40dd383a9	test-AiFXFKBJqf	77234024270	5539597587257112425	$2a$10$heRoLgVih5Uk9V2OMxdMGuXU31mqw8Jtua0b8JDOZ7FOUf5gYvVc.	t	f	2024-10-19 16:30:40.411448	2024-10-19 16:30:40.411448		f	basic	user
c580fe9a-0996-4c30-b9cb-d58b45f1d686	test-VNVtq5iyeW	76387942402	5861411537332524411	$2a$10$o84Ay/7wkgWBBlIwh74x6O/P53eByFzlapUiNoD0bcGOkGCTtxUga	t	f	2024-10-19 16:30:40.484638	2024-10-19 16:30:40.484638		f	basic	user
1de7291e-9f52-429b-b661-7a9f6c769477	test-a1UP0kD2RK	77845353799	5663587256903149282	$2a$10$Lu2gAyIkDBvVK50X6MwEmOZn8iTjSmPJyu5fRB5Do0kRhXQQecUkW	t	f	2024-10-19 16:30:40.580228	2024-10-19 16:30:40.580228		f	basic	user
f5c752ad-94c9-41f1-85c9-ac18e58a264a	test-vbRoLltesi	84265765706	593574483977952949	$2a$10$/xWqYf5xol5NusBArwncLOXJioC10wc/jaSBvsd/.vIp0DkD3uMvm	t	f	2024-10-19 16:30:40.65356	2024-10-19 16:30:40.65356		f	basic	user
13376df5-56ee-4b92-af29-cb66cfa97a35	test-WdXoQU7wUL	81488010077	1774366231608359863	$2a$10$HQYqy5A4HXROl9gnS98A3eqvxFSktYVXXSGVnPTTxek8Um43kqk3G	t	f	2024-10-19 16:30:40.752303	2024-10-19 16:30:40.758286		f	guru	user
d7e38c6d-5058-4ffb-9c3d-ea0192a590ac	test-HdILfjppCq	71652610833	8936297940402152722	$2a$10$V/v3KSsfarOQMotBgQ5.0.u76zLhSdxBGXZe/6oRmft69ZN0/s8lG	t	f	2024-10-19 16:30:40.835191	2024-10-19 16:30:40.835191		f	basic	user
a0cde28b-a282-4c74-98a0-f62fbd4d72ec	test-0PLXG1e7Bi	85478730688	4613963307639014406	$2a$10$BEm1rVx6bw4oj5cXlfu18eDcgkUWR0loipxiNzI0tfAu5e3Bzsllu	t	f	2024-10-19 16:30:40.908933	2024-10-19 16:30:40.908933		f	basic	user
fa96765f-abe4-43bc-82fa-8a0137908029	test-33c9nb87ki	32080679089	5394628216278058539	$2a$10$wN0G0LUaGoZan08tEDdRp.8gI58iKWkLuJuHutoOAQjl5mOHQvzIi	t	f	2024-10-19 16:30:41.001435	2024-10-19 16:30:41.001435		f	basic	user
b0d66f21-a91c-4137-9f52-4b6f5ec16083	test-xJTZTR3FVe	70777736530	647442303592773621	$2a$10$nHlPkZ33Qwl38BKh/WhDO.JdcFwG4MYAv1mUfhSwCXgecQAokOZhS	t	f	2024-10-19 16:30:41.076143	2024-10-19 16:30:41.083372		f	expert	user
30bf43b6-f25c-4caf-bc2f-73fbf0f826ad	test-ngxTcQES8M	18951651214	4404551257130856237	$2a$10$HiAkFNFT.jgEEzX5w5wz3Odal9.qH1mSlhaQJB9gC7U3xrT4csPoi	t	f	2024-10-19 16:30:41.173771	2024-10-19 16:30:41.179202		f	guru	user
a3782e28-cfc1-45f4-9fae-ad83c104fb2b	testiculous-andrew	42060041521	1478756321949910579	$2a$10$eEaSia36zgYv.h9epC7RLuva59PlSSf.ctcnVU0tnFJQjJ5K1dZ/2	t	f	2024-10-19 18:50:35.248693	2024-10-19 18:50:35.248693		f	basic	user
76df5669-8271-404f-92bc-83dbb20214ba	test-xDxSqx2ZbR	55016327453	8159131159965587571	$2a$10$eNm28UpTdmOovpZmtSYaGe6UQ74PocChjXc3.aZ4/FnIgtMIDfk2K	t	f	2024-10-19 18:50:35.33097	2024-10-19 18:50:35.33097		f	basic	user
b99038a8-8c8f-490e-9457-cc810d168ddc	testiculous-andrew	51517241600	6090536520236586630	$2a$10$MCwtziUQCws0E7LrtjFCEODBdtBNivRO30KB.xSjwV0nXsqc8XYR.	t	f	2024-10-19 18:50:35.674857	2024-10-19 18:50:35.674857		f	basic	user
ffffffff-ffff-ffff-ffff-ffffffffffff	He Who Remains	77778889900	6794234746	h5sh3d	t	t	2024-10-19 18:50:36.011554	2024-10-19 18:50:36.011554	https://akm-img-a-in.tosshub.com/indiatoday/images/story/202311/tom-hiddleston-in-a-still-from-loki-2-27480244-16x9_0.jpg	f	guru	owner
7f0435d4-df93-42d4-b5fc-58fe2ae80931	test-06IhxtX1BN	89183403106	7179405420190962110	$2a$10$pB/Qu4560xyD.pT1GpVKLuMG5wIlqLatWpLuyhP59.hf2bnILp9WS	t	f	2024-10-19 18:50:36.55268	2024-10-19 18:50:36.55268		f	basic	user
31352792-4bfe-4cff-93fc-ad70347a588e	test-KrutPgjMJd	41402371648	5694354048882508134	$2a$10$Nb2VCa2Tol.8BoCXB3ghaeeg1ZwT6zNdZQMqppU.RpxdwIiR50AMG	t	f	2024-10-19 18:50:36.638382	2024-10-19 18:50:36.638382		f	basic	user
71ee134b-27a1-4c87-8e22-841bfa9b6afd	test-hYap2Tk8e3	38870885106	3613398520827922880	$2a$10$Qclggv4s5.m5IhCpeX76OOcmA9L3TJ2UcNb6jobYWDlqt4LHawKre	t	f	2024-10-19 18:50:36.720927	2024-10-19 18:50:36.726233		f	guru	admin
1ce8758c-cd45-477a-aef1-72bc9f597e8b	test-wbbRJzx9M4	71011403366	671557555271458237	$2a$10$9PYaC3qKnstlOq5Njpex5ORyp0Ox2fK1mDD3rEieanh01JFS181o.	t	f	2024-10-19 18:50:36.797276	2024-10-19 18:50:36.802866		f	guru	moderator
f1b24062-938f-49c8-8f22-0977b242f3b7	test-tp59OEJobi	72029523995	1366049836444033924	$2a$10$SO0eWUyVSiEr3HXicfBLpeK0WaXi7LtTgBSzn7gjVpw0Ism5CP2NS	t	f	2024-10-19 18:50:36.871032	2024-10-19 18:50:36.871032		f	basic	user
896b1c66-c3ae-4f2b-8b9a-6356f7f0d498	test-jKRR3WP7y4	81773238754	4287086920308213812	$2a$10$tqVUCHJbwAiiUqju/J6x1OgUOlKkBXKnx6QJ9zxXFlTnu3Vmsm4xS	t	f	2024-10-19 18:50:36.954438	2024-10-19 18:50:36.954438		f	basic	user
725bcc59-3921-40b3-a034-65466859d5f3	test-jo2DjGB5yy	54740177338	5265933911116988431	$2a$10$OA2H.eq8iQ3txlsk0nz2FOj0tYKQyiQdMWbFpR8qMytxj2iXPfUKm	t	f	2024-10-19 18:50:37.024685	2024-10-19 18:50:37.032364		f	guru	admin
93679999-cfaa-4840-85e7-3d24d8f5ba38	test-VFo5aEEmuy	79407407875	1517056964682653255	$2a$10$vrqQgbgQuwSdHQmmJzEci.sHdnYUzmnajRIYCRBK0whLOfAhSWEw.	t	f	2024-10-19 18:50:37.108453	2024-10-19 18:50:37.108453		f	basic	user
1db2bc93-b016-4a44-9d1f-dc5df91313e8	test-RUx1NZIgY4	96633543582	494653005728086440	$2a$10$ozNbWZn4E//eTB1JzvTNqudB0lSS1rYFAWWNPhuukbq1Gg.Qnzhbi	t	f	2024-10-19 18:50:37.178788	2024-10-19 18:50:37.187978		f	guru	moderator
2e76a665-55b5-448a-a35a-dd090f334542	test-mZPf5CgaEz	27759587208	9196871464980909649	$2a$10$XemSJ5BcJp81F0tiDUjLnu4ENr6SQZVyTV2TVRcF8mjvtbR9bE5Bm	t	f	2024-10-19 18:50:37.262629	2024-10-19 18:50:37.262629		f	basic	user
e7dcbb5c-95cd-4383-90df-8df617d2897b	test-fxEMP7Hy71	52438616989	7698117246517885303	$2a$10$X8N/ABEFNYMePk.psf2YeO4JpeLeOYNHKmBhSrSUCBngvERjxUOOO	t	f	2024-10-19 18:50:37.335792	2024-10-19 18:50:37.335792		f	basic	user
a221acf5-7eb0-498d-932e-d2c5c053d723	test-39E1sF86BP	76795774182	1424632188889624280	$2a$10$k6a9bnsUqwT2QwuayQxcB.cdbowUzTw0Q2.nchOh4JARS33PQeZCy	t	f	2024-10-19 18:50:37.41227	2024-10-19 18:50:37.41227		f	basic	user
1b88e2ef-349d-4cdb-9f43-98cfb4dcb9c9	test-NNYv3KNSUr	70793328175	5237257631774701098	$2a$10$nf49zdqqO4s.jhi7l4wvQeFqz7dRZ/ZflmoRxNgXOlBRIzQcAcpFO	t	f	2024-10-19 18:50:37.486025	2024-10-19 18:50:37.489671		f	expert	user
1b479b0a-604b-454d-8816-dafb6c6f95b1	test-5J6C4wCbWX	85897016023	3471769014059898158	$2a$10$wEn57EUPFZw.JnTFwtyyKu3671IErgB/pYSFBHuVpSCGoZqY5Mfjy	t	f	2024-10-19 18:50:37.565695	2024-10-19 18:50:37.565695		f	basic	user
0da4343f-09c9-4bf7-b79b-67df93b26fbc	test-FVXGvUuNcw	22066448309	3440575091700674906	$2a$10$7oGGwdh81y6csV7WBq2RqOxqVAqywPP/8TBZonR8UQaF50qCJn4nG	t	f	2024-10-19 18:50:37.63717	2024-10-19 18:50:37.641852		f	guru	user
c62ff00d-2436-446b-8bd9-eda0b5893261	test-P8xLectfup	77223369982	7803775475912545402	$2a$10$1CpubiBVUAzzihUbzv75.OSadtSS0SxUmt4nu2Iv7HiTbVSCf3yJ6	t	f	2024-10-19 18:50:37.715803	2024-10-19 18:50:37.715803		f	basic	user
822b2bdf-a446-486e-ab78-77940fc2fae0	test-GbYLocScuM	82895023085	2677710955878376038	$2a$10$fkXcXigQ5PD9FQqDh1e/VunPbhX0vBUN.a6iHFzQtc3dHBV0mA1Fq	t	f	2024-10-19 18:50:37.787133	2024-10-19 18:50:37.787133		f	basic	user
f6814fda-4853-41f7-94ed-1fa68ae4b8ec	test-7y1az6pevq	41040483648	8898996202471024416	$2a$10$9ROokZ2WpC0HioQsTB2KYucGZTy6SWwIM8woRy10JC9VsszmkRcZ2	t	f	2024-10-19 18:50:37.863926	2024-10-19 18:50:37.863926		f	basic	user
66c44487-e5ca-4329-8df3-1c557d309192	test-fsGS96RPGJ	25675844141	5377583706035897561	$2a$10$BsLMw0kdqS6XkIaDcHDke.IhdtfG54vy9A.35B65MYtArG5p1baqa	t	f	2024-10-19 18:50:37.933267	2024-10-19 18:50:37.938248		f	basic	user
bb8e3c40-653f-4593-b9fb-f50179676ffa	test-3V8yJa0gLX	43355890731	5714894729336419628	$2a$10$vE0.rJaOTmbDwaPMhupEL.umqhrV4nSg6edR.q/9tE.i4SrrdGTUC	t	f	2024-10-19 18:50:38.012758	2024-10-19 18:50:38.012758		f	basic	user
3e2cc388-0a8c-4534-9ee8-1c4132ebf5e4	test-hKMkica6V6	55232900669	5298002331410968945	$2a$10$IYVMaCyLQ/gZxDxUa6WLee4VwAiguyOBgDVouRYgl8aGEI.RE8K4.	t	f	2024-10-19 18:50:38.082404	2024-10-19 18:50:38.090807		f	expert	user
28d8689b-a7c5-4f97-a9de-28d26f260b15	test-ftpou5pLga	45152366569	537520079105749078	$2a$10$pTgNTK1yyBUd.751pOuUweOU4OrqemakhuPsyp8F5m76pK/SL/s5i	t	f	2024-10-19 18:50:38.16487	2024-10-19 18:50:38.16487		f	basic	user
5f3d5140-e88b-4ca9-98f1-f19a6757bb9b	test-hoejLB4KzT	86863683484	3212695147864643957	$2a$10$69STrd4XhmuDbR3zMpcpVOBmO.CuUTbPlow50fZK/xhr6r6H8Jnxy	t	f	2024-10-19 18:50:38.235678	2024-10-19 18:50:38.241363		f	guru	user
09be1ac0-4a59-4a05-9a10-5dec31cfdb68	test-FO7lk1QYjw	65927224962	1769722503228208689	$2a$10$.dCibHy9gflWxAXFzjFK9eGTIR5maF1WnUB/DimWOKg.fhInPclTa	t	f	2024-10-19 18:50:38.316411	2024-10-19 18:50:38.316411		f	basic	user
77165e6d-2f37-4683-a3ed-d817968fbecb	test-yQ3KYQDv69	82751532906	830941338324485708	$2a$10$cj74lKn5MLviAB.oTKNowOIgMBvd9uT2VjkNQVBzkhwSOuvJAIXke	t	f	2024-10-19 18:50:38.389053	2024-10-19 18:50:38.393721		f	guru	moderator
04e7de45-0c4a-48b8-a41a-54c5bf3be53f	test-4l0iBKiQ7e	53859454251	8046067474373029748	$2a$10$AWLe3YFKbd5Ljp.LT8MgrufXKMpecZ1NZcRNvKEoEy8umT1ir6QUq	t	f	2024-10-19 18:50:38.482523	2024-10-19 18:50:38.482523		f	basic	user
484927ea-5c61-48d4-af93-92e5103e4668	test-BJCdNaBnKN	19936060541	8327422138159334132	$2a$10$e8f42aeZLBa6JlIW17oR2u6CmLS0tDswZTooWqVUnng1XEp34.GD.	t	f	2024-10-19 18:50:38.561213	2024-10-19 18:50:38.561213		f	basic	user
43028cc4-282b-4869-9769-5d28bf21e07c	test-LFUaEGRKjo	32515756743	8406000448842146872	$2a$10$rOjdpZQfufbdS9SOaqQ2FeHwd8x45Ool7WJ2xokLZWdX2sa4c8YKS	t	f	2024-10-19 18:50:38.634185	2024-10-19 18:50:38.634185		f	basic	user
57704b99-2611-4f5b-8154-6887eb6c8c84	test-mn4j14yT7G	48357929659	1908563077755010221	$2a$10$.FrLS.aoWCfQz9TaMFUvwODJeQ1jltFtrExXxc6C/Sv0vlaAC93k6	t	f	2024-10-19 18:50:38.714655	2024-10-19 18:50:38.714655		f	basic	user
f0b3c9f1-77c9-4bde-a128-c78b7b7ad86e	test-GvbqcBQS0t	37870810885	687291442697238978	$2a$10$G0Lee7FK/Da0w/3TMAhfgumRaBUqraM9DvRsiAscMlNJmLcDuQfVi	t	f	2024-10-19 18:50:38.795809	2024-10-19 18:50:38.795809		f	basic	user
6867864c-93ec-45a1-8a73-08ea25a4b220	test-ezjnHbORts	12037213946	9025660041546556984	$2a$10$R/MPW5sHDnZtYfa/O5mp2uNNzeN1rxCSxHFyktXzIWv0pS354r/tm	t	f	2024-10-19 18:50:38.877804	2024-10-19 18:50:38.877804		f	basic	user
e36d3814-6a6c-493d-9a72-91c97c15c1d1	test-DfHytIgE9p	72252341100	4338480995661482164	$2a$10$Xf/HdF0qP8PbTWUDppjkmul1zgDTWMmeB6c8mkaSl/d.syBL2zd9q	t	f	2024-10-19 18:50:38.950398	2024-10-19 18:50:38.961681		f	guru	moderator
6890ef0b-0091-4d2d-b25f-f8800b577ea3	test-2AN5gtZih0	93181172496	2577749245945783053	$2a$10$yH5B0bDX/ABqIkIge0Mn.uyclsXjAuljgoYsjpysUFBpdmBHiUJo.	t	f	2024-10-19 18:50:40.489383	2024-10-19 18:50:40.489383		f	basic	user
99b46525-ee48-43d3-810d-789da483c7c7	test-NhsB1nCqMf	43170467155	4244186799610841379	$2a$10$16i1HvPJYJH7B.L8YMEnh.wGe7UA7b86gcNWxG5ubNX.7UdGy.x1u	t	f	2024-10-19 18:50:40.569138	2024-10-19 18:50:40.569138		f	basic	user
ef6d7a38-93ec-4a62-aec0-63027cb431ea	test-Eyns9E3lus	94236971489	2583804004315748562	$2a$10$Bwtp9W8DoWkMCtz6x2suVu.1lcw2jxmZhHIlIFXXB7qJ/gas54EQm	t	f	2024-10-19 18:50:40.695417	2024-10-19 18:50:40.695417		f	basic	user
67bce66f-3d0e-4f27-b464-1ccd3999191f	test-AlRH0Sqcdl	91209915240	5770082875849525522	$2a$10$rh0svRFOew7Ng5mJF17gcueo.b9bpd8HbZnvRcRtmsrU4Zbh.kWM2	t	f	2024-10-19 18:50:40.762866	2024-10-19 18:50:40.762866		f	basic	user
e02b4b86-7250-48ba-80d2-1dff7d2f8df8	test-9middtstme	52821860258	5871585325259150188	$2a$10$rSPxFJEGIkt4iVfRHLHVo.B3zGWMYE1Ko4.xT/opVUzgVhSejYATm	t	f	2024-10-19 18:50:40.866598	2024-10-19 18:50:40.868901		f	guru	user
6e99a07e-2ac3-466a-a00c-419b12f3f2e7	test-QQhy8okQcY	82431926979	5525472840266456810	$2a$10$Gpm5gw9kx5.D5lhAHgupd.bzQB89p4WUnme6zNVHGEZBNqeTaNVS2	t	f	2024-10-19 18:50:40.941193	2024-10-19 18:50:40.941193		f	basic	user
b15c4b1b-4aae-426a-b6d3-dfa589b03705	test-eXY3e6KWT5	38734233538	4668922809443349412	$2a$10$LPcibU9tenRvi/3lygqxMOimmEbKfEQJCV3u1faX2wkjTykEil/j6	t	f	2024-10-19 18:50:41.008367	2024-10-19 18:50:41.008367		f	basic	user
3ae0a69d-27f8-4cbe-8e5b-0f185dafe009	test-cWFdQgZBOQ	18880928492	5831414247409914255	$2a$10$Cgmbec2ltF/F98tf/t.OOuYWLdvY/M3TL5nKGmt4QN4edlBP1vjF6	t	f	2024-10-19 18:50:41.093799	2024-10-19 18:50:41.100592		f	expert	user
e150180b-d35d-44b7-84c0-88d32c1afb3f	test-X5N72vDPsh	60422511845	8296301940344286160	$2a$10$DGM.pIKDJdo8SZVZ5dBHxOV9ajJLNbRcNnn2Gf8JY/gr53yoQqBKK	t	f	2024-10-19 18:50:41.167995	2024-10-19 18:50:41.167995		f	basic	user
c7fb1c7d-6be4-49f5-8dfb-2324d639b2b9	test-EfCLyqmcZ4	97330331378	7284886365347298109	$2a$10$7fgFlSzBOaFFve1yv/.GAO1KO4.ONdQtQaH4E3Hz/5nF8BdCUG.nm	t	f	2024-10-19 18:50:41.256234	2024-10-19 18:50:41.266658		f	guru	user
e4166f80-9a47-4bd2-a740-27c8fb6a48cb	test-h0tYI9rO2p	26434849550	6374030536908904554	$2a$10$y/G33f71Oq5.udzOh7QF4uo52IKC8d4/LuNxr/yR0ZFx/QETNKW36	t	f	2024-10-19 18:50:41.337583	2024-10-19 18:50:41.337583		f	basic	user
ed590c79-c544-4792-8b90-d138dfd61ffa	test-WcwyDcADyx	43121796717	8809121549789448979	$2a$10$edadWyI3.Gvz09dWcuamEutdFDBxljSkTHo95yVFmFatZeUMfVVvu	t	f	2024-10-19 18:50:41.407313	2024-10-19 18:50:41.407313		f	basic	user
b5cecb52-7877-4029-afa5-e677a1c2ec7d	test-d3gVhiMlYG	17898234302	4564119418126969484	$2a$10$OvnhVkNRJ2PUgRjNXEgwkufLALr0oA5O/8lw.tgHvDpM62RToHD3i	t	f	2024-10-19 18:50:41.495319	2024-10-19 18:50:41.495319		f	basic	user
ff97ae71-6c30-428a-96c3-654c9d56ba89	test-T4CG45L5uU	11436140981	3968464947294930876	$2a$10$yGwqbhIWPhrxdOjSfXNzr.mDcE65UlLV1Qldb0BwkAuXPE2UvFXOa	t	f	2024-10-19 18:50:41.565275	2024-10-19 18:50:41.565275		f	basic	user
9e4974e9-0a1c-4397-a574-38d4fef14765	test-cNEVvVIISo	97445203471	965275402550490084	$2a$10$ysAtzXy416pcGLC0ARbDoeVOXS17dxE9nc71BrOlCcT07OBHfYtO6	t	f	2024-10-19 18:50:41.651709	2024-10-19 18:50:41.658974		f	guru	user
0ee8284c-e553-4816-b928-2415ab336d72	test-dXVV78wBfx	15212017185	4418670239931721270	$2a$10$YzKaAL5El5XsdbMPkMX76.uOVXrrIFy7i8TlEypObZZmQdlyhWFRa	t	f	2024-10-19 18:50:41.729256	2024-10-19 18:50:41.729256		f	basic	user
aa5552e1-693c-4c54-9db4-ee304e6af9f1	test-jrcmFtY9NM	11027294755	3400701153678056630	$2a$10$6CyKwhVG0yKk5c0zEhgmS.YVx3AVbz90NwijzOZDvCqlaAoHtulzS	t	f	2024-10-19 18:50:41.79867	2024-10-19 18:50:41.79867		f	basic	user
a2bc57d6-4555-4230-ac8e-712bdfbeaea1	test-0xBYlI6oeA	54661996399	802665703796915551	$2a$10$h/VtAJSyw3vP1k4qyDi6FePrra4kqB4YZoV2N1Z7BMxfHHr4XN906	t	f	2024-10-19 18:50:41.88462	2024-10-19 18:50:41.88462		f	basic	user
da99ff67-9b25-4b56-835b-633e22c66870	test-24Fw96GSJT	20516586482	3517454826369019521	$2a$10$GtbUBBw3ac/bW4YNQvy25u6QgAZwjagEOcGMz0QWnSPBtCQvymm4i	t	f	2024-10-19 18:50:41.957665	2024-10-19 18:50:41.966364		f	expert	user
98b984b5-1392-4e0a-a8da-f45353a79b02	test-1rwWd71Kpc	60890744219	362329536021769720	$2a$10$jaOVHkGKb3DJrdxL8urvcenq76QyZf.3IbZCfu7iUTnQqbGCegT0G	t	f	2024-10-19 18:50:42.047235	2024-10-19 18:50:42.051202		f	guru	user
44d4a60e-79d9-4c82-bb32-11b3ef8b5a73	test-kLLZVX6Mpm	83215584869	5255243090952794487	$2a$10$qAhpQEvwxp3YxEvnjipdOO8ovfQuqnKQrf7Zzn.4jBWKD68Z8P0LK	t	f	2024-10-19 18:50:43.339607	2024-10-19 18:50:43.339607		f	basic	user
59489fd2-0d3c-4e69-b973-6814f7e6fe8b	test-mEaFbaiixL	32963053953	4811858129424004751	$2a$10$hWN9AV4jn/xfXH8PgZj.cusQYcGD/dSt38a4iXOHbwTgOkRB79Ojm	t	f	2024-10-19 18:50:43.420563	2024-10-19 18:50:43.420563		f	basic	user
5d0c334e-272c-43b3-a361-4726a2dfbf85	test-tV9IpBWi1K	31697896972	8308513556075579322	$2a$10$j3f1rTeAyxjofXluZ/M2jexSCIadQ58TS1KmoKNw8KrcMFtLm.Caa	t	f	2024-10-19 18:50:43.506856	2024-10-19 18:50:43.506856		f	basic	user
210887f4-dace-452e-8b37-9acfdeb9ff3b	test-aRrs237AmJ	61855791877	5097212030364299753	$2a$10$hh7McpuuqAN3FOdSgslNnOQR/Cixo3/bFh3TYVxqPbTI9UBDWbmSi	t	f	2024-10-19 18:50:43.575236	2024-10-19 18:50:43.575236		f	basic	user
100bc084-ae37-42db-938d-c71d32dbb89c	test-d0I4vX74Nd	36871448331	1862315468615484160	$2a$10$x/d3FsOO7nED.VSYSyJrdeEnCdspOwMmWrGXWW2b78rEVUwPyuwYu	t	f	2024-10-19 18:50:43.676413	2024-10-19 18:50:43.676413		f	basic	user
bc0bed96-18c4-49a5-87fa-3eadf3e555da	test-ozqopR9gR7	66940810566	5024967844211038207	$2a$10$SUz7GW8CcfMGw7xKlI9D1OQbKDNNWf78bWI0CN9dKB9lp0A782ZKW	t	f	2024-10-19 18:50:43.744298	2024-10-19 18:50:43.744298		f	basic	user
2201e804-7513-4100-93d1-e3f247e49e84	test-b1PC49nqN3	39740035289	929628542962048555	$2a$10$Dw/vxOogdQrlAv6p0l2hquMW6kY5Cr5VMjwR/fH63eob5WIAnhcie	t	f	2024-10-19 18:50:43.825563	2024-10-19 18:50:43.825563		f	basic	user
3c16965b-a69e-4db3-8728-ed11d13abf06	test-ZuuwYDnC77	93847389121	679712496527813742	$2a$10$JIRjk66.lyiWKK3NuOVM5.q5DZ/wHHBOWd4XGp8RwVNRKqZJ.xTc.	t	f	2024-10-19 18:50:43.898382	2024-10-19 18:50:43.898382		f	basic	user
1ffc366e-39d8-4ccb-884c-68ebefdc680b	test-gTcZ43wIWR	65590888606	2766188318014508298	$2a$10$5n0UvTVTRn0eUiTIa4VTFOHF2b1S92WZiWON94wkoeo/pnHIuSZny	t	f	2024-10-19 18:50:43.986917	2024-10-19 18:50:43.986917		f	basic	user
579affde-faf8-4b01-9727-d964d5ec456a	test-JoM9o8h825	13741948946	7559703228960758398	$2a$10$1fqjaO37BbBvPy3TPqNB1OxyWFU9hGdljxZ/RlnvtmR4eJpKHObaa	t	f	2024-10-19 18:50:44.075246	2024-10-19 18:50:44.075246		f	basic	user
85f3becf-2826-47c5-85c4-1d58d090ad4f	test-TnaFbwsjEX	32808527650	7316095314287548073	$2a$10$HM8YJ7Ts0DMYdimvUcm5huOGolkNu6Ahc9P1OnzWc/mTSmsurfC4i	t	f	2024-10-19 18:50:44.15014	2024-10-19 18:50:44.15014		f	basic	user
e88337f8-66dd-4b19-8ee1-ee5a3a7086d7	test-to33VQtnfR	93006723724	1047625150765949012	$2a$10$7Xw5qZm0Y2Tma0efArQCEuSVkD7G1SNoNex9Toq4CFZnbHQG8LKle	t	f	2024-10-19 18:50:44.234052	2024-10-19 18:50:44.240477		f	expert	user
d4db4b08-437c-4ebc-a696-c4ad35e75734	test-4QNLqM8Vm5	95511542730	787696504018459252	$2a$10$MOBIuKt.LnZgFQQJm4hDSuxV41x2HC/g5NdBti..1JFRci0a4ffQ2	t	f	2024-10-19 18:50:44.312481	2024-10-19 18:50:44.312481		f	basic	user
f3bcd7a7-d673-4915-853e-d6d5b5eadb5d	test-YhbQEVKTUV	10061583338	742617382833852811	$2a$10$jDgX2AZWKrm/6zMaC0Pe4eF0XD3yZRXovaA4t1M0gJgJO9eLiydfa	t	f	2024-10-19 18:50:44.381026	2024-10-19 18:50:44.381026		f	basic	user
7bd1d16b-d0c9-435c-bdfd-c7b888bf162f	test-aJo6XkOwGA	23968928393	6377840827925738283	$2a$10$ITctjATWI3iYQd8mDkc5wuJ50aZD54HO.QFe.MUZZ.HCWyD61RffG	t	f	2024-10-19 18:50:44.4639	2024-10-19 18:50:44.474653		f	guru	user
f4c07801-8c69-47f8-9719-e439d500639c	test-ddEHkJNRkZ	96414460269	5656913443776939442	$2a$10$2Jmea40L5WV0EZ1NRs9Kgui6a9BEDcYGnrywjheKc.97s95aGpZve	t	f	2024-10-19 18:50:44.545495	2024-10-19 18:50:44.545495		f	basic	user
791fb853-9bb3-4195-8d57-b51325b5ce6b	test-Au8dDVqQOy	43819997329	1069594840270999111	$2a$10$Gn.fudMcN1rHi11AeI.k/.uUUTAxxlSqg74GrUKu8luF4x1ssHY6a	t	f	2024-10-19 18:50:44.615543	2024-10-19 18:50:44.615543		f	basic	user
bf783f1f-6ea6-4a58-ba4e-10f9a113d963	test-hWQwpF905A	87897233839	1240062225617957482	$2a$10$C3zCid5WMMxiLpn1rPugbOHBnNFut/noddZJc2fCzfp.kfwXB6o3O	t	f	2024-10-19 18:50:44.694578	2024-10-19 18:50:44.694578		f	basic	user
9bee80eb-4c99-4823-90ac-c056be769100	test-gdO4EGEdt8	33677411243	6045063189546681007	$2a$10$Cus.J2R/nGTZ4KBYvR5tXeUUFAfDWv4WMXrjpZCWPchQEQ6Vx/29C	t	f	2024-10-19 18:50:44.768215	2024-10-19 18:50:44.768215		f	basic	user
b2b3786b-0b30-4ce8-9788-17f91bc783cd	test-Nhplz6dQMg	75819919269	6335301410837415080	$2a$10$D4eWXPoz4o5.kOmCksQSAOz/ulkoZDqkfQHAaYcXHJKdFbCGdhTQW	t	f	2024-10-19 18:50:44.841155	2024-10-19 18:50:44.841155		f	basic	user
5b1479c4-489b-40d7-8a81-38fc4633caee	test-JOpVf2fcoh	20757733900	7814811971905736582	$2a$10$ym3xz6PnXB4hSlCZsQuspeaQSG8LyiVphOe1GbiSiA75ye3qRJ9Xi	t	f	2024-10-19 18:50:44.923514	2024-10-19 18:50:44.930032		f	expert	user
524b58ed-10f5-4921-82bf-852780979138	test-Uej8VqgxFf	22927571674	875041319064591092	$2a$10$MvzsW69Vsne6DQ3LxAzfl.C//Qwvqj7jdS6r2md7QaobRF0OQHdZ.	t	f	2024-10-19 18:50:44.99851	2024-10-19 18:50:44.99851		f	basic	user
b930e0f8-e52e-4f24-b4a0-6097fdaea051	test-MT2BlimSaU	64516390996	9174749674637921320	$2a$10$byP6WsHNWEhim1g8oaOEoOUAUDMeKI7HZr86DI2dty.6JBbNjcs/K	t	f	2024-10-19 18:50:45.072025	2024-10-19 18:50:45.072025		f	basic	user
72f4f8a1-a4f7-4456-bb2c-56a054877985	test-2tbbnkgjb3	26138902308	8345630306292277006	$2a$10$MXfSHEcjGyDSlovKSrhWluP4ik0kslsql4oN9VRlRjI0lETIuqbDu	t	f	2024-10-19 18:50:45.152666	2024-10-19 18:50:45.15838		f	guru	user
846c7558-68d2-4074-aad8-1a740629a2b3	test-S0qYk6uVNq	61883919638	8896515449295553970	$2a$10$uiVm8EzjhtK7W..4sFGcg.BTCmkXNyD26WIBb37Rhrs6uWmD81gHG	t	f	2024-10-19 18:50:45.232125	2024-10-19 18:50:45.232125		f	basic	user
1f1ca188-a8f0-4936-bcb3-a880dd72962f	test-1wIkeN4ixg	13082576613	4062982174023417138	$2a$10$YE6g2Jgecv5hQAg2XDkAE.8hyfdk66icXWWtIxBrwOUqgrX3oor06	t	f	2024-10-19 18:50:45.304934	2024-10-19 18:50:45.304934		f	basic	user
e3f9e8e1-b240-4380-90f0-298a873f35c7	test-Aeahq8Azvk	57459486600	2782127036862437684	$2a$10$dtdNePxoQDL4Y6olkWy4MOwppNLZ9Jbeq0HDBOvpKtHDL6Im90eR2	t	f	2024-10-19 18:50:47.284336	2024-10-19 18:50:47.365984		f	guru	admin
72861de4-4436-441a-a5fd-de62a9387f1d	test-6eiI1o5Mgm	90980152203	7878750763402682956	$2a$10$nNKc6YxK/rD1LXgkY0fEdOAS2HXBmL4IbnDPwKouXKePqcXhNj3QG	t	f	2024-10-19 18:50:45.383435	2024-10-19 18:50:45.388885		f	guru	moderator
bc4df6fd-700b-421d-84f2-1efc8ccd0552	test-YwPdLcs8Kl	98207135868	4687348085350615181	$2a$10$moZHWOR/tVLvGBK4400iIuDseMYlWnsisbPxpXaaN08OlgkMEt6RK	t	f	2024-10-19 18:50:45.969821	2024-10-19 18:50:45.969821		f	basic	user
4af36e3e-8a05-4020-8a80-852a703f53eb	test-HZ9H5b2IG0	78149558267	8960074715196066644	$2a$10$LcAJjJjdG9OpRg4OElZD8uJgT2iIvdWc83ylJSQ0VtrVggoHGjbmC	t	f	2024-10-19 18:50:46.056993	2024-10-19 18:50:46.056993		f	basic	user
85b67d36-69c2-47f0-9de9-1d0aff20609f	test-S197JZa4eQ	91798915741	6867660657622770761	$2a$10$R9R1BwzM/c8dC/pocqfhUOEuK1mPvNK5v15bI2EjtWEEsepRh7Vqm	t	f	2024-10-19 18:50:46.126582	2024-10-19 18:50:46.13257		f	guru	moderator
979d4c77-4d9e-4b46-ba92-93fbab31dd36	test-R2rk5nfNyZ	44802951350	4916156516847124016	$2a$10$L6CwNOkXClmxkGf9ziapOOTU/JrmGpSn2Z.2/nh2pI6EYUETqkAY2	t	f	2024-10-19 18:50:46.211924	2024-10-19 18:50:46.219562		f	guru	admin
12340c68-1c95-4a16-9448-f324c6151d7e	test-1DBCbrjDng	36012799805	7181078602995215171	$2a$10$yDF1/g5SaNdvQhfK8sTcnOm3W7p/0W1oCqPfG.PmwJFnEYZV9gHbS	t	f	2024-10-19 18:50:46.288921	2024-10-19 18:50:46.288921		f	basic	user
73d29a4f-6e17-41e4-bf35-412fd71e0f8d	test-J4av8r8Xs7	11008966546	487614013125143765	$2a$10$0rQqfEDB2KUkFYnO3YIyuuz/83LEXlnAyVyc21l/4f4njK5pW96Aq	t	f	2024-10-19 18:50:46.362875	2024-10-19 18:50:46.362875		f	basic	user
b8e99a27-a527-4957-ba32-6100d755a935	test-SYJhNAcSOg	61300605773	6327112203533542504	$2a$10$/ENUUlD7Ee4sSM2v9puox.WSKh.2S12.7J4GB8KKiLfAVnS2QSn6G	t	f	2024-10-19 18:50:46.43812	2024-10-19 18:50:46.43812		f	basic	user
1b7fa956-701b-4e68-8732-948625e4822f	test-2JN35h8NKm	21309474182	80720958512404630	$2a$10$RXWFvAyfy25OgCIJKM5jp..G7pX2H.K2pUJqbRuXivwNAkweGKqw6	t	f	2024-10-19 18:50:46.510107	2024-10-19 18:50:46.510107		f	basic	user
f585ab45-37cf-4230-b3cc-200c06ce2422	test-jES9MOnt4w	97988450256	5756088525091837991	$2a$10$.K3RgQc2PNDLIcuVURWhned7BIHVv516oGv8SHLmgJcxRyNktZZCu	t	f	2024-10-19 18:50:46.580513	2024-10-19 18:50:46.580513		f	basic	user
4e3baf3d-5ddf-464f-bff9-107c0cdf47e7	test-bVZOUnJUO3	21024612715	5550353080113824216	$2a$10$zrC40sBsU8Th4QZ4/1Y7C.1wrZCH38nHTR83sPHd9S8ULIrd0gWdK	t	f	2024-10-19 18:50:46.656729	2024-10-19 18:50:46.656729		f	basic	user
202b54f5-509e-4153-bb26-d9630b41a847	test-EXisgmVlQB	25543795210	7799422254747761258	$2a$10$4/qF.OkWLwrRUp7e2fWSG.VqZOffnjVkCN3nKRoX6.gzdglNubovm	t	f	2024-10-19 18:50:46.724865	2024-10-19 18:50:46.724865		f	basic	user
c99581c9-1efc-4bca-bc76-b1ed24e9b402	test-MwpKJVHaFC	42729043059	2152522634620117115	$2a$10$zQSoc77.NKHaxjhZGT6L7OFRAG3bIOWGqnh7jzwsUgYT0Pl0ToIoW	t	f	2024-10-19 18:50:46.819985	2024-10-19 18:50:46.819985		f	basic	user
a239e20d-39be-44d7-bbd2-f65ca81debf7	test-gkG3ae0JuI	85315714076	8195797937098201584	$2a$10$xdKcrvid0luGzrJ2Toj9Xe8ZhUo7JmvNmTXYQj53JPdIiZDgRncVO	t	f	2024-10-19 18:50:46.893943	2024-10-19 18:50:46.893943		f	basic	user
1c271ea0-5d3d-4ea4-b11a-52843c83882d	test-u3jeqGva7d	56930720813	8095721337573129913	$2a$10$8fgi1pmioEOyXPRC9H.XFOKZ..ZaoUAuhlNPP4275vYKB9ls6mPru	t	f	2024-10-19 18:50:46.974025	2024-10-19 18:50:46.974025		f	basic	user
240f7400-012b-4f6f-a3ba-dd169e77e9dd	test-56xbHQVd3H	13799696157	2814923689239144202	$2a$10$QoQ10BnlSusgvxVGK2A5.eMVqqthPpYpB/aBYX9MKi6bf/hDDYxXG	t	f	2024-10-19 18:50:47.201452	2024-10-19 18:50:47.201452		f	basic	user
fd9846e8-a6f7-491f-909c-2ed55c1b9f51	test-wAvWHEzW6L	12264547008	7512396627344737995	$2a$10$vAKrFcDg3a33X89Mq8Vbc.8KbYFjcN8HYSpxXKFLTnElFaR.gm8Hm	t	f	2024-10-19 18:50:47.129893	2024-10-19 18:50:47.215056		f	basic	moderator
3e585d15-adba-49f0-81ec-53b4ed7e97b6	test-HypXEQgLes	19433178851	7746610052093403133	$2a$10$3coHk5ANaFB8Opq388IJ0eDO7hHuoLJZCZCeRc25jzHNBhsdMQyba	t	f	2024-10-19 18:50:47.436143	2024-10-19 18:50:47.521013		f	guru	moderator
88e88e29-40df-458b-a369-4a567e55cee1	test-yY11R5q0If	10352971362	5215693295203119897	$2a$10$uEzqVtpSroXisa0fmSmT3u1nmEZavWtHGnp91c4VDbpw.gmeIwfMG	t	f	2024-10-19 18:50:47.510159	2024-10-19 18:50:47.524044		f	guru	moderator
29ea05f6-5bc1-4445-8b16-64e18b5c652a	test-h2yRhUETQ0	88290229349	5514426558636096269	$2a$10$GzEvzkSEZtlicuG9Gs463.HhbGuTSIhR3zmbR9F5VgwM8lgRZTdFm	t	f	2024-10-19 18:50:47.592062	2024-10-19 18:50:47.674087		f	basic	moderator
4df9f0be-d59c-4e64-9e8d-b3f5a34d1221	test-dBFs7ZJSZN	28792884617	3748448177191347835	$2a$10$LsJ4dAhO3pBL85ESzxC74uLSzs14dKepVFa9H24y8lJaeD3QxTjLC	t	f	2024-10-19 18:50:47.66516	2024-10-19 18:50:47.675206		f	basic	admin
07df1da8-f9b7-4685-a151-5c6c62f60809	test-S0YplEq72C	59207378645	3963689007861433320	$2a$10$EZlGZdzHQ1HFxFezVrZ4mOohM8XeHgf/yTiYdNHF.zXlj5azVi6ge	t	f	2024-10-19 18:50:47.743999	2024-10-19 18:50:47.824521		f	guru	admin
afe5c0f8-71a0-4966-8430-d4ac394e1516	test-jtCalaJtXK	31183030225	176984951534599937	$2a$10$pcMnUryN3zW3Wn/pVXCx2.TJfN/ONWzXbcDJ.D/nE6W6zmXIa7am2	t	f	2024-10-19 18:50:47.896709	2024-10-19 18:50:47.986635		f	guru	admin
44c1f3a2-240f-4e01-bbeb-4b25406280c2	test-mxAs2bKA8N	14425961447	2116702555980672014	$2a$10$qzGe0AdnFqmyZ/sSYmfSKu.t74tGbzP4SKQFpn.F.amqYaH30m9Km	t	f	2024-10-19 18:50:47.973628	2024-10-19 18:50:47.990041		f	guru	admin
41adf979-55b1-4f8f-9ee7-96c5f0c9d3ac	test-KffGdebNOB	21007351462	8768829862399893492	$2a$10$hGn76lbojVKBK0c/rPUar.zTV1CkBJ6DCV1xBU9yBA4EycLW6Q8pS	t	f	2024-10-19 18:50:48.138252	2024-10-19 18:50:48.143151		f	moderator	user
5b0e19b9-06cf-460b-8360-a749e3db0e6d	test-vW12wkgbJg	50028340394	6512650742954417583	$2a$10$8IPbr64M7HIq7Cr1M9EZ2estuDwPTrgvzodHnjmFG1hW9eeTI68W6	t	f	2024-10-19 18:50:48.21229	2024-10-19 18:50:48.219408		f	guru	admin
ed2ad29d-6621-4e74-b055-121d651d5886	test-t5mtnPwPdR	96445065083	6894552669742621236	$2a$10$OVmAw91Hzb5Mjr1hAf5Uc.3hVsmpAiQtWrwxT7L24A6ng2D0upPFS	t	f	2024-10-19 18:50:48.288644	2024-10-19 18:50:48.288644		f	basic	user
b7681c76-93f2-465d-9628-ee3a75f067cc	test-E5iRGJmlcE-new	69177689137	3387177530068232931	$2a$10$5.2K8zIxwe77nfXWu2q/hOW/aXLCa9fzUwBiThs/o2HnKrWYLIaF6	t	f	2024-10-19 18:50:48.360177	2024-10-19 18:50:48.368117		f	basic	user
dae88ffa-1596-4367-af03-dedb78ac15f1	test-gY4s0WH9th	77000000101	7905084436721540477	$2a$10$gFxNW8lu6PpncWmWCHH/B.U4IDJ/oo1Wtewm5no2cvIjz9CC17iQS	t	f	2024-10-19 18:50:48.435784	2024-10-19 18:50:48.445459		f	basic	user
a19e00ef-63b6-429b-9119-c81d2b9901eb	test-PTaDJEIKkt	53622615735	798569499848863489	$2a$10$2sHdSZ9kr2EGqHnkPmJdWebJvmgaqVRuUdkVKZMsjCVeFXEgav3zy	t	f	2024-10-19 18:50:48.514769	2024-10-19 18:50:48.525033	https://jollycontrarian.com/images/6/6c/Rickroll.jpg	f	basic	user
b5f92430-9c36-492f-b5c1-fb88629917c9	test-0OTwjyCQHr	46356854253	4669579166957219	$2a$10$9gniXDbkRlcRJXt4jxYt2OXuT/voTIB7Bo4d1czT3oZdOzWk3RHp2	t	f	2024-10-19 18:50:48.591943	2024-10-19 18:50:48.591943		f	basic	user
7e141766-497e-474d-bb78-b70108b0da2c	test-rxkikYOBw1	80740162312	2595901933550363706	$2a$10$FchHUhz8VUYAqw5Ldgr8mOuTdBmfD7v0uFM5Xdw5aUUwJiPZrt8lO	t	f	2024-10-19 18:50:48.677995	2024-10-19 18:50:48.677995		f	basic	user
98f8eb73-31bc-4a60-92f3-d7a9ede69faf	test-562IG6xCXL	53823149803	3889926808592197771	$2a$10$DeY/JWuaUqKXWPbnn8I0/.kdGewfWTmQwL2Olb87y9BRi624kSfQi	t	f	2024-10-19 18:50:48.76183	2024-10-19 18:50:48.836916		f	guru	moderator
094dcc3c-ddd1-4b5e-8b95-bc0368f74638	test-mBodDM269t	81529087022	1454746227341398546	$2a$10$sdXMXldmfA.QUim/TmiejeG3pQt8F5WUOW.LQjHR4Nr7CJwEYvmkq	t	f	2024-10-19 18:50:48.830724	2024-10-19 18:50:48.84113	https://jollycontrarian.com/images/6/6c/Rickroll.jpg	f	basic	user
ac8bacb9-9650-4ea0-b0e0-d185791188a8	test-5MTNczFspq	12614656576	3941703704986988363	$2a$10$AHGTeELlqauCM4RwTpnE5uIoYdOkKRNTDuqdk.SIJL85vYCIsq4Zq	t	f	2024-10-19 18:50:48.907977	2024-10-19 18:50:48.989831		f	guru	admin
440ba8b4-83bd-4517-8e80-7b3cfdc6dcd7	test-tHglEM5ocm	76722077085	4840947431376319650	$2a$10$/gsyxB4ZKtMoDF5J6Vjt8ua6u2ikT0lxe9k21A5FYEKBlh.AbXgs6	t	f	2024-10-19 18:50:48.981474	2024-10-19 18:50:48.994474		f	basic	user
d9f29746-b84d-4c1d-b3fb-c45658754a7d	test-E0FeaOctbI	85898015726	8367678209660109354	$2a$10$nU4QAhFFiQ7cf8bV6s9yfuEBcOl50Tfc4BgH7Vk1X/TfGFMY3c3ky	t	f	2024-10-19 18:50:49.0619	2024-10-19 18:50:49.141101		f	guru	admin
e190c3f5-d830-4b69-958c-64a036c91387	test-xg4yLxfaZm	54390501239	6875826966588524621	$2a$10$667CdQDYcv.A809v.nsxROkBH85MycbMBQAUM/Tf/7Pk/M2GoS3Ae	t	t	2024-10-19 18:50:49.136249	2024-10-19 18:50:49.143703		f	basic	user
7e137d0f-3ca7-48a3-be82-5ff15e2e7b8b	test-mx6oJEMTOY	21634519685	2242232104618195930	$2a$10$rJFCMgNTikzHWxP5WUozIeptzo2B9bg9sS5pVCUIgw7pjtcEOfNWi	t	f	2024-10-19 18:50:49.21235	2024-10-19 18:50:49.293762		f	guru	admin
e85a36dc-ef9b-49c0-bfc2-038fd7557867	test-SGyyJgyldW	11809558011	7212159206915177386	$2a$10$690azKt6bmF0MBRqQQNk.Oa8PvKlgtoQKmsP4FLaIMkrpg1B0ifLK	t	f	2024-10-19 18:50:49.283911	2024-10-19 18:50:49.296875		f	expert	user
5e46432d-0c48-4220-a69b-a0caab78155b	test-2gq0y9FZJ2	26962277676	6072140421538497571	$2a$10$KAC2yah9V.MsEgh1TjTOGuKYCvna2vNXruEr9VzI36iPSsoa0k732	t	f	2024-10-19 18:50:49.366631	2024-10-19 18:50:49.445065		f	guru	admin
4258ef6e-b80c-41a0-849b-0a56e8e3249f	test-TuQRiR8k2d	32316079127	8286175370926299379	$2a$10$i9cs9JhhHoCFV12TwRNrt.rGIzbLnLBluKVD273DCyVWvGxysnvWK	t	f	2024-10-19 18:50:49.435198	2024-10-19 18:50:49.447615		f	guru	user
f6b40817-fcbb-4160-a8d6-e4e62a94fc27	testiculous-andrew	28197131896	6209304228444051320	$2a$10$1IWZ4VPAIW5/GUFszBQUsOQIHFzbyXq7Van1na3Apo1IFVT665W5O	t	f	2024-10-19 18:51:16.069876	2024-10-19 18:51:16.069876		f	basic	user
0df7e545-93de-4e5d-829e-dc36c195b117	test-dJFOteWJpj	29261041367	8311331542319258780	$2a$10$v1EQuaffSuvHt6euIr5SqeA8CKZ.XOHSEPmQbnpZ2brsHd8pVal5a	t	f	2024-10-19 18:51:16.152418	2024-10-19 18:51:16.152418		f	basic	user
312a3f72-afbc-49b7-8afe-8766ec63cb90	testiculous-andrew	42054323531	8201639111071553694	$2a$10$mPsGPfBBjrY6NPtDBbvdReIThxDWVBT/pnEw8KHulqnzUtjBpSssW	t	f	2024-10-19 18:51:16.506307	2024-10-19 18:51:16.506307		f	basic	user
8469fe8d-32df-4825-8088-cf7cb1fb943c	test-DQe4lrOvEA	80907419350	3303885147313581938	$2a$10$JP8ApOQTZ.drI8rwH/kRM.p2v7ro14BIhXbIaxkMD0Qv3lz.lTGsq	t	f	2024-10-19 18:51:17.449173	2024-10-19 18:51:17.449173		f	basic	user
ed38bd8b-6748-4e08-a275-a1311cdf4366	test-R8sqggXWEa	80299811295	8157115941526513614	$2a$10$xR93EyXEzqRUm.pPd.khS.t2nJnNK/dtnDAREGAeNs0Fpi69fKGV.	t	f	2024-10-19 18:51:17.535115	2024-10-19 18:51:17.535115		f	basic	user
3df9d7d8-d225-46fb-9e46-ec344626a202	test-zBj3TCiCdo	83982520871	8283675854673273266	$2a$10$/92dTzDuNyPmZTT/2pIcHeOvvZF1J72LhAf6UjSjSDNgNWeRPV9fS	t	f	2024-10-19 18:51:17.619079	2024-10-19 18:51:17.62512		f	guru	admin
bc67d2fb-e081-49d9-adca-68d756020de0	test-X9KcZ4LSI6	72514000655	6303738587094097592	$2a$10$GBzVbxi/T9OCdZb9a.eNyuUTqpVoPZ6B4Ue4TR9/h3uNt.DPOk93.	t	f	2024-10-19 18:51:17.698343	2024-10-19 18:51:17.703327		f	guru	moderator
6d5733a9-3b1b-4de3-a823-a9dad14feaeb	test-Tps6sAbgZC	44826639578	8127579157759896737	$2a$10$5U1.rXILMKR6RZya50cOru7bJ5oc8sOciILQp1zii/n/tqalPkT.q	t	f	2024-10-19 18:51:17.773796	2024-10-19 18:51:17.773796		f	basic	user
c550de45-58f4-48ce-8a4f-925b3b67bc21	test-CreUfTBysG	38504333626	3237550871757056566	$2a$10$yvOEFJD7NgRFYxsKZokmCe7EoU/zFMp67uMY1mnV8Njha2WpGCSwe	t	f	2024-10-19 18:51:17.864065	2024-10-19 18:51:17.864065		f	basic	user
e8cd1dac-44f8-49a6-a127-268826d339f9	test-ubfsPJV3aH	60598148874	9085176756853891468	$2a$10$/JWBzs/zwLfjC2r4W/XKvuCQTlIBWj36euvb/dLHbdk9BEXKpPoc2	t	f	2024-10-19 18:51:17.934585	2024-10-19 18:51:17.942872		f	guru	admin
e9b634ef-54f7-4126-8b02-817088c4b18f	test-jLKTx6J2P0	46098601470	6519987448661970709	$2a$10$nfKDs4x6g6AExqVMgPPA/OO5V1IJuSsu0sDVEaeoaMxphweNeJZYK	t	f	2024-10-19 18:51:18.019161	2024-10-19 18:51:18.019161		f	basic	user
438f3759-66ce-405a-9752-dabcc51e4773	test-GCFvaE7IKU	88685425410	5509082971377623214	$2a$10$R6XE6u/.iA7SoUlwoDmNmuC6meXr0qip.Uc2rgfewicCEO9wNOHVq	t	f	2024-10-19 18:51:18.090331	2024-10-19 18:51:18.100974		f	guru	moderator
3d83697f-ca21-4efb-aea2-c01202a6b6fc	test-6MBhMIcknD	61330359763	3207111661709998655	$2a$10$hpK2BSzZHUJUDv/4ckohBed22osN4AT8Bg339FdjXCzipFrV9TLHe	t	f	2024-10-19 18:51:18.182854	2024-10-19 18:51:18.182854		f	basic	user
f9943ee1-ad5a-4099-b579-5a7f6c6077cd	test-ppNJMC60Im	44960500889	6313509427127664962	$2a$10$KH8hItlQP6/P8bckvqnsb.x/HywGRSj1CB4gq5WhzholSYIq9a2Se	t	f	2024-10-19 18:51:18.253491	2024-10-19 18:51:18.253491		f	basic	user
fc4e2ce9-464c-4a5e-8879-6241ea3abd11	test-SoZPHmVz5o	97903015004	7175211463602686002	$2a$10$MtF8YySJr4AQE7OKYVmSWeuICgT35G.quH/yORCtTEy9xtP8bd3jG	t	f	2024-10-19 18:51:18.333505	2024-10-19 18:51:18.333505		f	basic	user
35406d00-7364-425c-95a9-e985a74203ef	test-xIqY9ky8cq	89258821790	4127445823246455965	$2a$10$ikeQYLGQ8EYOSnC0.YY3quWrcclZJa7kjRePLiY1OEsS9Xe4JFtxu	t	f	2024-10-19 18:51:18.412034	2024-10-19 18:51:18.41594		f	expert	user
14ac0578-89fb-4c27-8e40-de14fea7a34d	test-DF4dJOlwIP	78635634129	292406941200243175	$2a$10$k8wqSwQXRyv.EId7h7k21uOkbOy0tsp/Eq7N1/DTzNvvc7cLGdaLq	t	f	2024-10-19 18:51:18.49295	2024-10-19 18:51:18.49295		f	basic	user
1a589ab9-8db4-473a-87f3-eacd6ef8a631	test-fhhnHP5ht4	86865378820	7577074667866257019	$2a$10$WWkRvxvS9eRzE4j12KfpiOscIjb91wOiTioV825KLyO6SJgT6spEW	t	f	2024-10-19 18:51:18.564104	2024-10-19 18:51:18.569844		f	guru	user
7650136e-f3ab-4161-ab94-bde90c176d6b	test-L6lIYThZIt	43112790010	2985270970558338348	$2a$10$M7nr9yrHN0A15.7PAtg1SurMkK6vCYnpv8P2ho.r8jOyyHIjYw0ra	t	f	2024-10-19 18:51:18.6436	2024-10-19 18:51:18.6436		f	basic	user
cc365f4c-15ef-4a6a-85d5-9de81faf2bd4	test-HHeRq6vwFy	62983962769	2333481651983895942	$2a$10$AN7ZYKQfQbFjnRE76Gda5e8ySp.rTz1un42rWZO9rJ//FiaXw7fPm	t	f	2024-10-19 18:51:18.715931	2024-10-19 18:51:18.715931		f	basic	user
62eae9c1-2f4f-4b92-bb80-f40c3f9fd43f	test-mWNf9s7ulQ	43075801188	6011395920809020015	$2a$10$oK1V.rvdiLjY3AkDB8x8oOFaP.JLWiMPkdyglLjdfVRfow04zIosy	t	f	2024-10-19 18:51:18.795798	2024-10-19 18:51:18.795798		f	basic	user
35f1a819-fbae-461e-ae7e-5d69dc58dbde	test-FKGVVjOX0G	37698631809	8776742369723967344	$2a$10$m1WZmHnhhIe6i/k76.h6o.gCUzn6FE.C/vsscYqCx103bn9XqFeJO	t	f	2024-10-19 18:51:18.865999	2024-10-19 18:51:18.871		f	basic	user
ff18a582-3447-4ee3-b613-3596944df7b3	test-qinqTZXBjT	31277064819	8152367989863030091	$2a$10$FXpj.prnV0TKy4.OMLBhReDeL691dZ3uqDvi3t5MgEldGvwjwrnlG	t	f	2024-10-19 18:51:18.946862	2024-10-19 18:51:18.946862		f	basic	user
dcbcfe50-8322-4f2c-9394-3e8829dd15cb	test-y3gv2JQZMz	89083761308	1168633282272329057	$2a$10$K48OroGtpRwroYQJhrPPq.tU.YanAfjTVPgvKhj8Oydq3jRYFR1Qa	t	f	2024-10-19 18:51:19.018474	2024-10-19 18:51:19.024815		f	expert	user
66e60f34-a0ac-46da-a0e2-85268e27b1f5	test-JmxJglCUbU	60939905734	8638795269356453370	$2a$10$TUfMlDVPPO597Lk6TVYJIu9wUKH6Ifi8l/xK83fQ4f/SFJqJEGL76	t	f	2024-10-19 18:51:19.102986	2024-10-19 18:51:19.102986		f	basic	user
0ce4bcd7-3ee8-4206-8ba1-9cb305e96d51	test-FtvFaYC8JF	16333908597	4364070229042907182	$2a$10$XZlsXs15Ana.OQna/dY8bO9nLKUJMVGPuCZ9Xk59PYbFtJ2VG/XOm	t	f	2024-10-19 18:51:19.175594	2024-10-19 18:51:19.180955		f	guru	user
a1bb7f29-069f-4178-8afb-8fa1bd86f57d	test-tTXD0yfqic	67385483454	6060641298517894068	$2a$10$6BmRp.zGh0rBL.8QQp3gaO2bE4z9z6TXJcV9idboM4Fipve66ZAdm	t	f	2024-10-19 18:51:19.257695	2024-10-19 18:51:19.257695		f	basic	user
c3eeaf15-7713-450d-ba97-f16abb9a07be	test-3OOYZlGseh	70981514037	4984260619773860670	$2a$10$NXNHNnyDA6sLfXCNFUE96.yqWa2fpLBS2snOEDz6697o8vp2cSR.G	t	f	2024-10-19 18:51:19.332812	2024-10-19 18:51:19.341116		f	guru	moderator
a40061be-c9df-49fe-b407-a4bf8ec4fe53	test-H791I2Lktv	96799114468	838570366779769428	$2a$10$DRf7HuzPlLJRb/NnCzXPaeBwG0yuF2lKOo1oqFD0HToJw2nPcX0zO	t	f	2024-10-19 18:51:19.425924	2024-10-19 18:51:19.425924		f	basic	user
e3b15f4d-7885-402f-93f9-424b59ce7fa3	test-mlxVptJRrj	34631269925	4250121658187599906	$2a$10$sXwylzxODnnc2NNVBAmOtORJuVGTb8DrHKgbj4e3J7I4oS3eggcwm	t	f	2024-10-19 18:51:19.503741	2024-10-19 18:51:19.503741		f	basic	user
df0b38ad-3a9f-45fb-8e43-36d775f285ad	test-R0Y0gKdHQo	36546355393	4334016868700053560	$2a$10$BsIdumGN91GsGD4RE85gSuwx5OpYnn7Qy4xce1D3gY.RoO92NXFCi	t	f	2024-10-19 18:51:19.575811	2024-10-19 18:51:19.575811		f	basic	user
075187c4-589f-4cc5-8e60-40d3286fced0	test-EF1c1k3ytU	41488639752	8298355356364648093	$2a$10$Dzjo.LhBqtC3jtYbS5eN7O1YjzwxUkfTrnYtM7yk3OhbXtbs2TeBq	t	f	2024-10-19 18:51:19.657735	2024-10-19 18:51:19.657735		f	basic	user
d10409c6-84fb-4839-8197-55387ba0b79d	test-iUfbznxhjv	83123359987	8783475193346224008	$2a$10$hpSO1qdJFKluemJ2cn7rGuKbsrIrARSfIeCYSRMx1FwgLO/P5ZCWO	t	f	2024-10-19 18:51:19.744527	2024-10-19 18:51:19.744527		f	basic	user
3fc78a7b-2fed-4c72-8519-63353eabbad9	test-xIMJtDNm1f	40250426692	3826893853006435156	$2a$10$42gIyT2qH0LtQeAy3z6yKu3p20d2cOmLU7l5WCuxVy0XepB0hFZJi	t	f	2024-10-19 18:51:19.824762	2024-10-19 18:51:19.824762		f	basic	user
1dc4f64e-3a60-4c31-bf59-b83ea79d6b5b	test-6pcO0euKIW	52350884723	5095468062382775665	$2a$10$dhfgFrvIr7vXysqxz029YO9spBLsg6egLkY7Cf/1qW3vyAUsy.TMi	t	f	2024-10-19 18:51:19.897902	2024-10-19 18:51:19.906454		f	guru	moderator
affe57b1-aa52-41ba-bbdc-572b6d53bc0e	test-z5Pun8Cfdg	15627840240	342813931674341561	$2a$10$6j62PHFpucHpZyaDl1ksmu8FWXCe0xo9KOEgESYRl3lnRxTHjLKmu	t	f	2024-10-19 18:51:21.441069	2024-10-19 18:51:21.441069		f	basic	user
1e9feac0-baac-497d-936e-ef90e7ae541f	test-xFQYMxfDGK	78146976538	6199467189430308218	$2a$10$VgWdbDrwq3nn6DKde4BOJe1DgrXfZt0IT2Pw3XElklllqxyxmjOGq	t	f	2024-10-19 18:51:21.523918	2024-10-19 18:51:21.523918		f	basic	user
4c274022-dbde-4d41-8673-5849b20c4052	test-0LSxpNP7NQ	20320234983	1080975690977956298	$2a$10$qu/BiAVBEYctYPkn2gZCduK0HsOkZI.WUUVXOfZ1ksa./s5qmCwoa	t	f	2024-10-19 18:51:21.653027	2024-10-19 18:51:21.653027		f	basic	user
85a470d9-39cb-4337-918f-214da397c062	test-HMAiysOLgU	49606392296	2274327780806466043	$2a$10$mvQAHRHIENa9.wpYkOTToO9V.Q2epcWwoN/dxtyAqBd.JBL.q8oZW	t	f	2024-10-19 18:51:21.722257	2024-10-19 18:51:21.722257		f	basic	user
3e959a5e-430d-4810-948d-4d7a11226e66	test-MfKwgTpXIJ	13866781772	2288727850761295924	$2a$10$tvup3RerEXHpuDWW/wDx6e9Ypu5ql89V4AI37UuSQEDU9l/lKG3ZG	t	f	2024-10-19 18:51:21.829782	2024-10-19 18:51:21.832429		f	guru	user
f7d194a0-166b-4cb4-8c89-71c519d07d34	test-m94O1BO1CR	19917775953	6642721502227007729	$2a$10$UH.uquieBGeGXFA13fupi.gVZXbNIFV0O7aObrCMy8YKb8lNwddlC	t	f	2024-10-19 18:51:21.90579	2024-10-19 18:51:21.90579		f	basic	user
66dbbf85-70ee-42e9-9b09-4ce5ff29cf41	test-3Yoys3Ifw3	50124121224	3048842934405035729	$2a$10$Ef6HoGZ8bKWx9AyHIwSSgO8O9l/oTNRpFQ4KEdeywkN./zrmo4s3a	t	f	2024-10-19 18:51:21.975053	2024-10-19 18:51:21.975053		f	basic	user
20a76176-0042-4923-a381-2aab9954e679	test-XKCcguJWUi	98192899303	4327891906914083023	$2a$10$QCfP3v5QpLF3r9aPr1kKfOu9nMRIqMdiWIL5G1CIzfqf44CJ5a9Iu	t	f	2024-10-19 18:51:22.062171	2024-10-19 18:51:22.066112		f	expert	user
15d77ba6-ed63-45c8-8f8f-44d53f65bb2e	test-bwofOuaIkN	39171794589	2223875416176559983	$2a$10$WGzZ0NAV6tE/m2snKGscdequeDu.hnX7h6k6U6VUJTtNRuYjQCguS	t	f	2024-10-19 18:51:22.134269	2024-10-19 18:51:22.134269		f	basic	user
09a21627-a3c6-4338-8fe1-ec342e8b861e	test-nVY9o6upuJ	91736852078	5967868567928134054	$2a$10$m1pirs9fD3/Rc9ygYdbF3.Wr/YMXNZIVYF4ntUE37UruBNt5qic5C	t	f	2024-10-19 18:51:22.225882	2024-10-19 18:51:22.232352		f	guru	user
6e379ad4-5b11-4aa2-8c03-e6d12e032755	test-3kQ93GnszI	41319276963	968667381237033106	$2a$10$4rYOGSZsSQoumFgwv1LOWOAS1KJjPbaRgbJCGioyNJdEpQdTcnUam	t	f	2024-10-19 18:51:22.304912	2024-10-19 18:51:22.304912		f	basic	user
aab90a73-7321-44f6-bfcc-76974bfa8448	test-cQnyxgqqnG	29699368759	2260294555672838454	$2a$10$nOvOnTs9zeL9T5xajSkfCe4mprxp.jIOd0H9ixZFYlGOuvGfmwSxm	t	f	2024-10-19 18:51:22.376613	2024-10-19 18:51:22.376613		f	basic	user
efc3d895-a248-45ba-9841-b9d638a53175	test-wpgKmU0VaI	44311551067	934589334863780585	$2a$10$AUkiXjtUyDpbSVpEq2uWsuRE2SOXaDSqwyqHB50pvZPe90JvbZpWS	t	f	2024-10-19 18:51:22.469895	2024-10-19 18:51:22.469895		f	basic	user
58fe0b33-ff2b-424d-925c-79c348ea6c2e	test-vQcEfNiCIS	80499997592	7129425176022114680	$2a$10$tHJsFl89809gcb.zy23nb.Tctf4esm5qRKE58njW/mjvQRTUp/FWC	t	f	2024-10-19 18:51:22.541232	2024-10-19 18:51:22.541232		f	basic	user
c0b2f6cc-eb1f-4ecf-96ef-bb69968d19c7	test-x4VzhkqaLu	84186341625	8109812698538504121	$2a$10$j6MPjzfPbaynr30ZqoZEK.ln96ik4giMHODgNvaVPqb9Hn9wmLk8G	t	f	2024-10-19 18:51:22.634929	2024-10-19 18:51:22.64009		f	guru	user
1628db9d-b81a-4922-a905-caa0d4c16a48	test-fZDLjZVXHu	13719876068	8153215363363394162	$2a$10$pidVgBEltD7jGaIksz5Vr.7Gqzolyu3b/o.oJrnUWFQE.XExZjgZm	t	f	2024-10-19 18:51:22.712067	2024-10-19 18:51:22.712067		f	basic	user
3b91a9a8-c3cc-44fc-b73c-46d82d80183a	test-gluKZ54K9W	13352300936	4281336625221418707	$2a$10$1jJt2H2K8IO.hM17Ciy4u.oCnpka96Hg4LN9OAC0JKK9Jux4SQWkO	t	f	2024-10-19 18:51:22.782965	2024-10-19 18:51:22.782965		f	basic	user
a42aa967-9a5c-48ee-8ae7-f4334ea0b56b	test-34qvppId9s	29434053912	9138412590173956429	$2a$10$HwWE2FQ.ayCSLXBzoB8oxev9LIh24Y3zRKmd5YAaokCebBOYdIEYy	t	f	2024-10-19 18:51:22.873552	2024-10-19 18:51:22.873552		f	basic	user
d1224a95-908c-4bbf-89e0-3a1842687c10	test-9Ujohep35D	52665726471	5509061688373590089	$2a$10$CmkTziW/RbM7O/q.quz9mufp3gtOc61WnKsivbornziUD72caR5k.	t	f	2024-10-19 18:51:22.945824	2024-10-19 18:51:22.950715		f	expert	user
cb8dacfc-468e-41d7-b95b-4a8c9a8b436a	test-KADzEZHl4S	35138508854	4166420242396945824	$2a$10$3lgaQbOuyeJgDSY6jhW5y.K/apk1rtTw3Nv8h4.dwhJLILo8mQt5G	t	f	2024-10-19 18:51:23.036356	2024-10-19 18:51:23.041639		f	guru	user
4a1c4b50-e174-4410-82a6-7257a0002dc5	test-stzzYLXv38	80375902695	2480695021590025877	$2a$10$3tC0EhXgZUdr78bYSLPStutcxyKEX1G3i.NcjvsIXDLehUvrraN9q	t	f	2024-10-19 18:51:24.381624	2024-10-19 18:51:24.381624		f	basic	user
7facc3c4-03e1-4c44-b84d-ac626d5eb4d8	test-1WYKxECuWD	14730477164	4167786433790493413	$2a$10$/Oa.aoHR39ufRRp7NJgIXe9EOvu0zQehLcGbGr2czV5f9DApl1z8q	t	f	2024-10-19 18:51:24.464291	2024-10-19 18:51:24.464291		f	basic	user
6490cf88-e23a-4be5-855f-91a59e00fef6	test-u7TFvNMKBA	61125191824	7190860842545784817	$2a$10$hjtTREpJdtI11d/nkTHW0uvUxkIMmRat95UwUChln8xrwwwBSxNVK	t	f	2024-10-19 18:51:24.554145	2024-10-19 18:51:24.554145		f	basic	user
5a977610-19f6-4022-904f-93edfc72ea70	test-Oa12DjQ2pb	71186159317	2082951491878755374	$2a$10$jIJQVMV3fLNEQJ6QI0jLPuMwPrDkfu1g.zCW7Ewoves6BVXA5u/JC	t	f	2024-10-19 18:51:24.622607	2024-10-19 18:51:24.622607		f	basic	user
f23e28e0-1095-494b-9778-2ce1289423ea	test-PHmXgGNatX	48286030530	8629180044243350869	$2a$10$9UUCF9Nu0zBWBu.U6hHIheA5yqA/f/VGnWBpxImMYumwBoy/yGTLu	t	f	2024-10-19 18:51:24.719328	2024-10-19 18:51:24.719328		f	basic	user
387d7a39-c15e-4e5d-aa56-c0ff4c9cef6e	test-2AT6Yhk1ux	63673435990	8217510165409300065	$2a$10$F3hRuuTjit2mvh5raGGZTe1TPYqcurufzRnIVp/3P/4nffl4gNVXi	t	f	2024-10-19 18:51:24.78826	2024-10-19 18:51:24.78826		f	basic	user
6f5beca5-2aad-4533-8d42-7cfebf2b0532	test-hdbdb52lZg	24608593293	4949078424223472775	$2a$10$4.tUB/9/gqmJ0Kl9Sgkg0Od12uk6nHm8di.K1NHf22sCwVSzsJNuS	t	f	2024-10-19 18:51:24.87067	2024-10-19 18:51:24.87067		f	basic	user
3d632ec6-6633-4165-93cf-f160eee0b246	test-QHfcB6zbA6	51096368473	4530503629190656689	$2a$10$pn9sI.Jp7rR5dOUtV1L/wO3AtywxazUXSYFBqWI6Iemv//OVuJ8O.	t	f	2024-10-19 18:51:24.940918	2024-10-19 18:51:24.940918		f	basic	user
c4d07164-0855-405f-88ba-2c44d78f5b4a	test-hvaGTn1pAK	46879864948	604827017007699420	$2a$10$8CKXS.GAEkVSu1Fz6dKy2.kgezBz1WP/MTsVceTgYPAxmR3GEw3P6	t	f	2024-10-19 18:51:25.02953	2024-10-19 18:51:25.02953		f	basic	user
a5f7b81f-00f7-4e1e-8526-b27cde367f21	test-aWID0q963d	36243072066	3093594749688368171	$2a$10$PEzcih4WPq4nqGxPEv3.w.nY3Rbi43I9FxeFVDl3.8cTQ5uTYtqji	t	f	2024-10-19 18:51:25.104429	2024-10-19 18:51:25.104429		f	basic	user
bbcc298f-3341-4ba0-b858-589141d74918	test-Lk17Pd5ZX8	29735886960	5813436807451478973	$2a$10$4V4c8ZxtBlfWMoT0z7sS0.A6EQc0S.iDIX9UklpMQxy7OeTTnrI2W	t	f	2024-10-19 18:51:25.176175	2024-10-19 18:51:25.176175		f	basic	user
089b76c9-638a-480d-9898-934f7925eb8b	test-BxzTW5GU1C	45582173302	4100548869601727647	$2a$10$8R9PzmtVPqQeLNfoUNlxouY/..9U8SZkrF2ve8kmzy6JGPR6AXR.2	t	f	2024-10-19 18:51:25.263299	2024-10-19 18:51:25.268367		f	expert	user
a4322a31-4340-4222-88b3-683d0e210b7e	test-WBBCGMFLFZ	80168978545	6130206537537072389	$2a$10$zb468WEr1UmYZZe.ni40G.c/818fQ8iMjtdh8cWYXuUtbAlq6MgFy	t	f	2024-10-19 18:51:25.34047	2024-10-19 18:51:25.34047		f	basic	user
2f22fa46-cd97-460f-b25f-e45373c99caa	test-P2Ecqvu2q7	32894361699	531173592388951946	$2a$10$Lj9vyish7n0W6trVsoMj..V8llhoWSjsBQY/w5AQo9/vanvNApopy	t	f	2024-10-19 18:51:25.413267	2024-10-19 18:51:25.413267		f	basic	user
623e9616-fb29-4210-86af-0f9fb8f358b4	test-7wXrhRIGYD	77092376002	6515509725499237013	$2a$10$JBWoob8Lpm7btLXQ0kU56.18Z.SAFbJDw0HewEiTWIiRzP4AHj2Ga	t	f	2024-10-19 18:51:25.499056	2024-10-19 18:51:25.50443		f	guru	user
4e08df8a-683c-4565-8b39-700c1e4275ad	test-rvXgUTKf52	10907690569	8333585798954897464	$2a$10$0rcP0xMPbk/In9Oxm1HcbuLzawr1734.MG98sM2ppbqEzHvnfoS.q	t	f	2024-10-19 18:51:25.577621	2024-10-19 18:51:25.577621		f	basic	user
5df86b9d-d643-49fe-9780-6c69cf3d981a	test-0YVsO33egZ	29616086829	2207873589987145297	$2a$10$Aiq984iZ1mBdqAnsr4C8a.yd8AGm48MHda8OBRleQYp30ZCny62BW	t	f	2024-10-19 18:51:25.647288	2024-10-19 18:51:25.647288		f	basic	user
01cf1a8b-f41f-43c8-8e49-df3e40a6001c	test-5BDm26zpgv	96052814200	6829690055890600234	$2a$10$74bCGSqrOZ2SzPMI6.NHDufw5yzfiQZ9nstDo3siweIDqHK7/Ncx.	t	f	2024-10-19 18:51:25.733771	2024-10-19 18:51:25.733771		f	basic	user
0a1a6ce6-0085-46f1-9803-7175849a1c28	test-n4regq6Wze	23945188471	1986707704820403688	$2a$10$TnLrmD3uBR/PUCVPtI70EukwkcI570G1Oja1gTZUEFEfpKd/JxjqO	t	f	2024-10-19 18:51:25.806232	2024-10-19 18:51:25.806232		f	basic	user
c6c7c2f2-6ad6-43d1-a883-94a4fee81a06	test-c2oJW8SWql	16721259960	2194925618930957614	$2a$10$NDwIc9Egjdn8UC3mu.9wM.Z5L6wq9b8ciXAlpnBKG09JLPpoWITPi	t	f	2024-10-19 18:51:25.87493	2024-10-19 18:51:25.87493		f	basic	user
17baa83f-62a2-4662-9a23-404a5d3e684d	test-m61pZ4bsCE	45813120173	1309672838441614360	$2a$10$bfi25zrzS5RNDB/0pncfXOlTFUxjHsIr4aVeM9z49qFtddtQaW4oC	t	f	2024-10-19 18:51:26.035924	2024-10-19 18:51:26.035924		f	basic	user
5cfa21f9-a4e2-41d1-a15c-d40dca62f1ef	test-AtwR7j0x9x	91292592636	1615005171525384819	$2a$10$aj6BCFcLRWkOTE4V0KXN1uTjPGXgmE2u68wyclDl6Uex5OBc65oxa	t	f	2024-10-19 18:51:25.960342	2024-10-19 18:51:25.964807		f	expert	user
08d9ba05-a11d-445f-9b3e-9caca6712316	test-lTu6QTfBOg	70587274439	861054207420576647	$2a$10$wVn/1SawoapsA1gOvRWlzekCJY3ZBoPCbP4EThN5vEyocC/CceRQe	t	f	2024-10-19 18:51:26.1067	2024-10-19 18:51:26.1067		f	basic	user
fbd543c8-5c9d-4ab8-b8f2-210774804839	test-MgyizjbBZD	60745074849	6312863678452454458	$2a$10$QQuOOb.z33ZreyzyAaYUSuIm/2EsnQAUtQIA8xD04BKhdqv2TMq4C	t	f	2024-10-19 18:51:26.195042	2024-10-19 18:51:26.198538		f	guru	user
aa8bc7e4-be00-4aee-9ad2-85c47f2395b3	test-8wfwUL36dj	50414172845	2003132433474382430	$2a$10$6V76KNfW57yOMjWyMrAq4./siCrZ80mFJ6fo8UjBPb78lmnZZoV.y	t	f	2024-10-19 18:51:26.274604	2024-10-19 18:51:26.274604		f	basic	user
a0f0c4ce-1e11-486e-bcae-d4b730eb3602	test-6sa4pC2Za9	46495642136	2013061391476956719	$2a$10$blSG3fcl0W1e2wINGt8ysePtRmC0ZBVje88a6DBInRcUm8FechoKC	t	f	2024-10-19 18:51:26.346743	2024-10-19 18:51:26.346743		f	basic	user
57974a4a-6b16-47e7-a58a-8266007045ee	test-9Mqclv555V	20048224400	2591632809143452774	$2a$10$oG2yC2xaC4XmutQneSzsY.Wap8neNGUws.FbyXs7DIVjTUpQNvuNS	t	f	2024-10-19 18:51:28.778969	2024-10-19 18:51:28.857972		f	guru	admin
50a81f39-f4b6-4f20-a263-e4274b1f402b	test-lJhuzxMS6y	51438360603	5136096283480411783	$2a$10$i0T.IL0PWpK1y/a039oshOP6IUBgBQrR4AAIAmbldw/t4nar2SlOG	t	f	2024-10-19 18:51:26.430906	2024-10-19 18:51:26.442564		f	guru	moderator
0919db04-c650-4f96-bd55-b1564b4ebdee	test-t0R39UncEh	12661273095	1387332378710603176	$2a$10$UQZmsR2dVI6B9/dEvG6.8emZ1tUi9YOKgdA/Y/3lJH/O5.rSPw1Fq	t	f	2024-10-19 18:51:27.028285	2024-10-19 18:51:27.028285		f	basic	user
40b0e991-52a0-4566-adbe-d8e4e4b48738	test-uLSPz0LKp6	23549782158	7061825976881306386	$2a$10$tGlAfWAW.WSSLGUTw86IQu0M585pZInLMYsjIQA3KdsQljFXJbrLO	t	f	2024-10-19 18:51:27.115166	2024-10-19 18:51:27.115166		f	basic	user
ffd85647-306e-4d2e-bc41-82fa81514dc8	test-ROOL1EapNd	27742983174	7603801280791637635	$2a$10$RLCc/CtHzTOCqucnYriVPeTazxEhE14NgaTu3U0acqnbnqt.OOol2	t	f	2024-10-19 18:51:27.187457	2024-10-19 18:51:27.193643		f	guru	moderator
cec34f7e-cb8c-4087-8339-2b2fd5e7bf24	test-rcCRONGh8k	90281309021	7679991281760207635	$2a$10$aDdZgGFpsPWKSmuZA6/nyuXqqAPbX0CcFkBHpgjAxa3/kgcfUJDJS	t	f	2024-10-19 18:51:27.276255	2024-10-19 18:51:27.28176		f	guru	admin
0758e4c8-8220-4820-adeb-4093fd650ace	test-4eAk92BrNK	97465004015	8222932151915790172	$2a$10$2mTUE8ypZwdFP6jDsTcUs.K9f37gYMTJSCREMGoHmyF1LDsZ2Cdjy	t	f	2024-10-19 18:51:27.35312	2024-10-19 18:51:27.35312		f	basic	user
41dcf730-1110-4a50-be61-5b1719084ed3	test-TVc17wGbh3	18822379837	5134567546386880992	$2a$10$e8NTMu6IwX2P8fbYb8J8DeSdYH0QyG01c18Xu.zHvrRyYWWjesR/W	t	f	2024-10-19 18:51:27.424148	2024-10-19 18:51:27.424148		f	basic	user
e95ca8e6-7136-478c-b777-2d8a0fd115ff	test-OQap9upuK8	75501586870	3029679158169505686	$2a$10$vVR7UA46geAy.Y/T8xpGf..PlFmrAeBZWydQy/NApS1bxmABRRT2m	t	f	2024-10-19 18:51:27.496191	2024-10-19 18:51:27.496191		f	basic	user
fcf66b7b-6f09-4b42-a4a2-bd6b35740717	test-vT7xiKtnir	55515118394	7180046669956065499	$2a$10$0iG9PvLyoJ6Sqvvpv/4xX.gjfCy96pcx8KaLaKASCjJY.hKktOTr.	t	f	2024-10-19 18:51:27.575121	2024-10-19 18:51:27.575121		f	basic	user
d77ae745-814c-482f-aaf2-91382f87ce8c	test-mAtrPuQRmZ	26587767791	2689112295098569027	$2a$10$wxdj391dwQwmb0xs9GtymuoVnAkd4Blc8G5d.9odvQljtMULESWtK	t	f	2024-10-19 18:51:27.647697	2024-10-19 18:51:27.647697		f	basic	user
8e8e9e7b-84f8-4125-b25c-7b3a1c98919b	test-HjCGNqU5yW	75882904235	9034104296440591298	$2a$10$jJ1QVXZBzDVHNvdRUrqVQeNi90ZSKYWd0C6.if5T62FWJdhzmrXve	t	f	2024-10-19 18:51:27.722233	2024-10-19 18:51:27.722233		f	basic	user
b5be679c-2936-4274-85b6-9273ff5987ba	test-ddl4siiYsH	27703099350	2031773556115394079	$2a$10$MhWvkpOODVKebDAe5yzRxugXpplNnZ5kWctM8bXZ6OoqCSGLOykaa	t	f	2024-10-19 18:51:27.794034	2024-10-19 18:51:27.794034		f	basic	user
29e2e552-799a-4c39-a405-ddbcedad9767	test-rk2hW26Xdk	47331517438	6138014219874243188	$2a$10$uFigCCruONhvlGeXSoruLuuXjJwQV4uJjfZSGiLeWdMMR2vBZSoGi	t	f	2024-10-19 18:51:27.869497	2024-10-19 18:51:27.869497		f	basic	user
ddb2be31-d1ca-4450-bb67-d4d6b917294d	test-cnLim3Cw3v	91560339824	1236707793244647318	$2a$10$gp6bre71L0zFXwD4rrm3pe2pfW8zrovKQKsJbutpuZeNVYumbG0.O	t	f	2024-10-19 18:51:27.94271	2024-10-19 18:51:27.94271		f	basic	user
1ce5d517-79e8-4949-9d59-081dda000c7b	test-XjQem3X0zM	54300763490	7221773414917142511	$2a$10$4zuvBiHFEFUfaauUSnXIuOknC5WBAKrmEROC.yQ63oDZabYc.Yfbu	t	f	2024-10-19 18:51:28.017406	2024-10-19 18:51:28.017406		f	basic	user
ed406787-4d99-4a32-87f0-02d65fc1a600	test-5OY6NPUVtp	84938063889	4033419596214514417	$2a$10$9YS97StGWvBZ9JAQEWc2IOWV/HB/YPEW5dYLtu6JcjQc9yOPuprwu	t	f	2024-10-19 18:51:28.24773	2024-10-19 18:51:28.24773		f	basic	user
569ab93a-ab86-49c0-bf7d-daba65daadd4	test-xK5xo5tDP7	95019899602	6373583938794340911	$2a$10$X7M9rcqrnZkz/e0dSscOzObDYs9T4NSPWuQPhaAxg/VPc/MsdD8Dm	t	f	2024-10-19 18:51:28.174956	2024-10-19 18:51:28.253191		f	basic	moderator
2f8eb040-e6a6-45b5-994c-afca1bc798da	test-SUsa9zafYZ	72324634532	6074235502767287772	$2a$10$UW9VETX/1e35aMzoVZz1f.i81sovCZC5/F9m4OOISjEkx3yyhdF7O	t	f	2024-10-19 18:51:28.323561	2024-10-19 18:51:28.402534		f	guru	admin
78cb570d-f186-48bf-87db-eb9d1986ca61	test-XNhM6kXchZ	60300156433	5124564567530090785	$2a$10$XdLvmzu4pLYi5RSaWOjD7ur2v57O4FS6WRpCD/pK9jdUaQUcWP9Ji	t	f	2024-10-19 18:51:28.475383	2024-10-19 18:51:28.557749		f	guru	moderator
a91054eb-f4cf-4b54-8aeb-cee430a2663a	test-oIlWKKY662	59355605145	403390944958333740	$2a$10$bg8J5eGqfF0b.LXJdacvq.8nR5lGpOk9qRVsbzVTlCbjRRtZ84uF6	t	f	2024-10-19 18:51:28.548674	2024-10-19 18:51:28.565254		f	guru	moderator
be31fa15-2c60-4dbd-b906-59fb9f571acf	test-RWMKOaC64Z	20501432326	197723999287982763	$2a$10$q1s5i/U6Yno10fuPYSB7m.T4fz.I/8J/7FrGHyH3isrAsjLlfxD9O	t	f	2024-10-19 18:51:28.634356	2024-10-19 18:51:28.709279		f	basic	moderator
3a524f26-a158-4c35-9eb0-ee3e4ef348ea	test-36wndH6hYD	98251115900	66962492125675444	$2a$10$WETaDIxLdsWph7BctscLPOxuwFnCSvBWYPYH4HO7C3DlDrjrqfAaS	t	f	2024-10-19 18:51:28.702828	2024-10-19 18:51:28.710374		f	basic	admin
f61916fd-5ec5-4a0e-8fa4-533b89526344	test-cEefg1X2V3	47273048239	1772095817954526279	$2a$10$Wu8dUFJQaOZHcTOMdnWKnuFMcZU1l6Y3cW38/Hgnef7PStW4nnuI.	t	f	2024-10-19 18:51:28.931203	2024-10-19 18:51:29.010417		f	guru	admin
19a07cae-5bbe-4f51-8747-0a414dfde935	test-GSgxeYjRa2	44897614460	6949892641830788234	$2a$10$mqNArQs00YhjzuEEU1uCF.Awu.eDz24Q16yyetBDa5PdcKwyJPZHi	t	f	2024-10-19 18:51:29.001465	2024-10-19 18:51:29.01486		f	guru	admin
c630d04d-9aed-4a68-bd11-fccda18ab631	test-MU2HkJ6u4i-new	18325873790	2118310178075485559	$2a$10$wD5wpWHLGNFTbIYBWjkeL.8sVv4Yf2ykbvXyshMv4o9b30kJLSFmW	t	f	2024-10-19 18:51:29.391266	2024-10-19 18:51:29.398823		f	basic	user
dc6024c0-2e45-4832-9770-cd67f983ff1a	test-MqV7cjcrOh	21197753068	1468609188832410599	$2a$10$t42b97BLDeJ3HPGoMMnOR.U4ttkHTzJaJLSXPfKHKZhoa0x1eSAce	t	f	2024-10-19 18:51:29.164515	2024-10-19 18:51:29.169826		f	moderator	user
4f611401-a474-4396-9686-25e2d02b1701	test-U7BNt76nv7	64569488796	676265030823941750	$2a$10$9L9g5KxAZ0Sg5GloSI0KfekgFwSKUvY4JG2QUT4VH0j882ApXpfGG	t	f	2024-10-19 18:51:29.240629	2024-10-19 18:51:29.248796		f	guru	admin
d0484003-9012-4071-bd49-87c56df1f167	test-kt0YRG5ziV	30203142761	1133224161140708590	$2a$10$gfBZF//GBnHTww1usgjEQubvwv7QALQP.tFlL.XU1gsJ.iQQYMule	t	f	2024-10-19 18:51:29.318351	2024-10-19 18:51:29.318351		f	basic	user
d7fee309-de30-4c52-b6b8-98e94c9acb68	test-ZXnWYt6hlQ	66048282637	8416938416385961265	$2a$10$n8Uxa7ly5rjg/AONDAqrReW/dAXv/8D0MuyEyqzySlICU9XS9nhoW	t	f	2024-10-19 18:51:29.466253	2024-10-19 18:51:29.466253		f	basic	user
d5ef503e-44b5-4407-922a-7fcc8d38f947	test-5CBvclNNxD	97881221384	2983123680532911277	$2a$10$R9cmeF0XPSxc9SibMENSweRG8obN.VYk0UMuLf2Euzo872PzZAN6S	t	f	2024-10-19 18:51:29.546861	2024-10-19 18:51:29.555754	https://jollycontrarian.com/images/6/6c/Rickroll.jpg	f	basic	user
8243b3a9-8367-4167-b2ba-502bf3f1e958	test-C3t0X72qOk	44578226971	420143682028570982	$2a$10$Drl0y0dk4agJqG/OcuW0qORCcfvX37kfrs9LMlnG2UcTINY9J91rK	t	f	2024-10-19 18:51:29.623829	2024-10-19 18:51:29.623829		f	basic	user
b8ff6c3a-789c-4d74-b700-1a150817a836	test-zprlNuBqYg	91763555313	6558939292954544164	$2a$10$RQ3SGev1KnvlaiPVuVT9ZubBdplkBKiGSdjMWgIV0pV.R34JOQ2NG	t	f	2024-10-19 18:51:29.694468	2024-10-19 18:51:29.694468		f	basic	user
279706f1-a678-496c-87fb-b23ed90594e1	test-0WIML1fQ7o	54034427114	6457822005317028095	$2a$10$lGRxvaR/VOehUIfAuytojeS4bR4.M2JFXdvcyeRWV0vL5f6XNiFhu	t	f	2024-10-19 18:51:29.767809	2024-10-19 18:51:29.848819		f	guru	moderator
38d70ca6-5507-4a42-ae68-874b0a588bb0	test-UZCMvDFuUc	31700267045	956727166994446288	$2a$10$9MiCgeGIeuipS8FCzLJrMezeikv5fIfODcMYJba0r6s4NGh3rk/G.	t	f	2024-10-19 18:51:29.84071	2024-10-19 18:51:29.851527	https://jollycontrarian.com/images/6/6c/Rickroll.jpg	f	basic	user
9cd44a6a-d2db-44f1-b5cc-e31d16aba562	test-ilEXmWCOxG	61382521112	1469292220660046258	$2a$10$bhG11mguwdg0w4D2hs75ceeR839758ncO6sfpBeqYW3ajOLUxe0VS	t	f	2024-10-19 18:51:29.919206	2024-10-19 18:51:29.997905		f	guru	admin
d132c802-e5b7-4302-8728-132a14d0de87	test-NNjEqLfGAP	33867679365	1918753631774240268	$2a$10$qaMndwUTz/nIfErhZFqD..wp9w5oqjNuj8bQ9cXnzBCklKu6b4rlK	t	f	2024-10-19 18:51:29.989891	2024-10-19 18:51:30.001129		f	basic	user
69127eeb-e9b5-4ae2-a8ad-ea75eed874f1	test-d1A4UUA9Tl	22453103250	5192638467823253220	$2a$10$EqLRUhpq3CIDc9kFWlV.beGuErotKAByVzAsCQt22/ow9IS3KCA8y	t	t	2024-10-19 18:51:30.140657	2024-10-19 18:51:30.156346		f	basic	user
b173da76-b4fc-4d7b-9831-d45f8b3bc677	test-o47IeFs9mk	89870505212	2136289548398956659	$2a$10$xvW7Xa5wY14WQVESXoZP.O.V3VKVXfgAcsb2UwbzO3nDAIfxwCX4.	t	f	2024-10-19 18:51:30.070404	2024-10-19 18:51:30.147527		f	guru	admin
78d07c16-f334-49a7-8741-a7809a7fd845	test-4Ir8bugkI1	21939181344	4576345661410815323	$2a$10$BfONQGfJ9Dm.WwbshUCsnuHSc7fk4Ic6S9iPMQoBICGBMYXpFjvoi	t	f	2024-10-19 18:51:30.225625	2024-10-19 18:51:30.308813		f	guru	admin
69150e77-091c-4bb1-ac05-2c8da014ecbf	test-O1ddI8emLc	14688348920	6077415990223921759	$2a$10$Q2Ni5rLstxbtaogDMzt8L.lcQ191wb9gZUrxbTeH3AfxzSl51zWNi	t	f	2024-10-19 18:51:30.298199	2024-10-19 18:51:30.311776		f	expert	user
4e9aa233-4f22-4b2a-ae11-95e33ebefb77	test-qj7YDac3ij	41856356013	5230383718224795227	$2a$10$DE5/ojctN0og0qa3pBlmK.N4zLNFi9c.Bats5o.Wxl6h43TSjkn66	t	f	2024-10-19 18:51:30.38274	2024-10-19 18:51:30.462535		f	guru	admin
9c46ab4b-aea1-4862-98d0-b4ad826f60e8	test-hJ9sDcH6Av	94467024884	7848397077036276732	$2a$10$ElCK5GKXB2JS/i/umQh3qOvnp8p9QOR8Hwkr7GJ4UgETrKT3MTGN.	t	f	2024-10-19 18:51:30.454755	2024-10-19 18:51:30.46573		f	guru	user
\.


--
-- Name: body_arts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.body_arts_id_seq', 9, true);


--
-- Name: body_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.body_types_id_seq', 10, true);


--
-- Name: cities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cities_id_seq', 54, true);


--
-- Name: ethnos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ethnos_id_seq', 99, true);


--
-- Name: hair_colors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.hair_colors_id_seq', 13, true);


--
-- Name: intimate_hair_cuts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.intimate_hair_cuts_id_seq', 9, true);


--
-- Name: profile_tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.profile_tags_id_seq', 48, true);


--
-- Name: user_tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_tags_id_seq', 30, true);


--
-- Name: body_arts body_arts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.body_arts
    ADD CONSTRAINT body_arts_pkey PRIMARY KEY (id);


--
-- Name: body_types body_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.body_types
    ADD CONSTRAINT body_types_pkey PRIMARY KEY (id);


--
-- Name: cities cities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);


--
-- Name: ethnos ethnos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ethnos
    ADD CONSTRAINT ethnos_pkey PRIMARY KEY (id);


--
-- Name: hair_colors hair_colors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.hair_colors
    ADD CONSTRAINT hair_colors_pkey PRIMARY KEY (id);


--
-- Name: intimate_hair_cuts intimate_hair_cuts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.intimate_hair_cuts
    ADD CONSTRAINT intimate_hair_cuts_pkey PRIMARY KEY (id);


--
-- Name: photos photos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT photos_pkey PRIMARY KEY (id);


--
-- Name: profile_body_arts profile_body_arts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile_body_arts
    ADD CONSTRAINT profile_body_arts_pkey PRIMARY KEY (profile_id, body_art_id);


--
-- Name: profile_options profile_options_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile_options
    ADD CONSTRAINT profile_options_pkey PRIMARY KEY (profile_id, profile_tag_id);


--
-- Name: profile_ratings profile_ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile_ratings
    ADD CONSTRAINT profile_ratings_pkey PRIMARY KEY (id);


--
-- Name: profile_tags profile_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile_tags
    ADD CONSTRAINT profile_tags_pkey PRIMARY KEY (id);


--
-- Name: profiles profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_pkey PRIMARY KEY (id);


--
-- Name: rated_profile_tags rated_profile_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rated_profile_tags
    ADD CONSTRAINT rated_profile_tags_pkey PRIMARY KEY (rating_id, profile_tag_id);


--
-- Name: rated_user_tags rated_user_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rated_user_tags
    ADD CONSTRAINT rated_user_tags_pkey PRIMARY KEY (rating_id, user_tag_id);


--
-- Name: services services_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


--
-- Name: body_arts uni_body_arts_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.body_arts
    ADD CONSTRAINT uni_body_arts_name UNIQUE (name);


--
-- Name: body_types uni_body_types_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.body_types
    ADD CONSTRAINT uni_body_types_name UNIQUE (name);


--
-- Name: cities uni_cities_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT uni_cities_name UNIQUE (name);


--
-- Name: ethnos uni_ethnos_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ethnos
    ADD CONSTRAINT uni_ethnos_name UNIQUE (name);


--
-- Name: hair_colors uni_hair_colors_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.hair_colors
    ADD CONSTRAINT uni_hair_colors_name UNIQUE (name);


--
-- Name: intimate_hair_cuts uni_intimate_hair_cuts_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.intimate_hair_cuts
    ADD CONSTRAINT uni_intimate_hair_cuts_name UNIQUE (name);


--
-- Name: user_ratings user_ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_ratings
    ADD CONSTRAINT user_ratings_pkey PRIMARY KEY (id);


--
-- Name: user_tags user_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_tags
    ADD CONSTRAINT user_tags_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_profile_tags_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_profile_tags_name ON public.profile_tags USING btree (name);


--
-- Name: idx_profiles_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_profiles_deleted_at ON public.profiles USING btree (deleted_at);


--
-- Name: idx_users_phone; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_phone ON public.users USING btree (phone);


--
-- Name: idx_users_telegram_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_telegram_user_id ON public.users USING btree (telegram_user_id);


--
-- Name: profile_options fk_profile_options_profile_tag; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile_options
    ADD CONSTRAINT fk_profile_options_profile_tag FOREIGN KEY (profile_tag_id) REFERENCES public.profile_tags(id);


--
-- Name: rated_profile_tags fk_profile_ratings_rated_profile_tags; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rated_profile_tags
    ADD CONSTRAINT fk_profile_ratings_rated_profile_tags FOREIGN KEY (rating_id) REFERENCES public.profile_ratings(id);


--
-- Name: profile_body_arts fk_profiles_body_arts; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile_body_arts
    ADD CONSTRAINT fk_profiles_body_arts FOREIGN KEY (profile_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- Name: photos fk_profiles_photos; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT fk_profiles_photos FOREIGN KEY (profile_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- Name: profile_options fk_profiles_profile_options; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile_options
    ADD CONSTRAINT fk_profiles_profile_options FOREIGN KEY (profile_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- Name: services fk_profiles_services; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT fk_profiles_services FOREIGN KEY (profile_id) REFERENCES public.profiles(id);


--
-- Name: rated_profile_tags fk_rated_profile_tags_profile_tag; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rated_profile_tags
    ADD CONSTRAINT fk_rated_profile_tags_profile_tag FOREIGN KEY (profile_tag_id) REFERENCES public.profile_tags(id);


--
-- Name: rated_user_tags fk_rated_user_tags_user_tag; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rated_user_tags
    ADD CONSTRAINT fk_rated_user_tags_user_tag FOREIGN KEY (user_tag_id) REFERENCES public.user_tags(id);


--
-- Name: services fk_services_client_user_rating; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT fk_services_client_user_rating FOREIGN KEY (client_user_rating_id) REFERENCES public.user_ratings(id);


--
-- Name: services fk_services_profile_rating; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT fk_services_profile_rating FOREIGN KEY (profile_rating_id) REFERENCES public.profile_ratings(id);


--
-- Name: rated_user_tags fk_user_ratings_rated_user_tags; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rated_user_tags
    ADD CONSTRAINT fk_user_ratings_rated_user_tags FOREIGN KEY (rating_id) REFERENCES public.user_ratings(id);


--
-- Name: profiles fk_users_profiles; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT fk_users_profiles FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: services fk_users_services; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT fk_users_services FOREIGN KEY (client_user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

