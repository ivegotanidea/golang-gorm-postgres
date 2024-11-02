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
0e50bdb7-fd95-47d2-a44e-9c0460e980d5	05029a37-875c-4f65-8f5a-c95ba1e32e1b	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:09.777429	f	f	f
91f2bd45-43b9-4b7c-8136-633b32ed6270	f648288a-9756-4c99-8d6e-26a9d929cd3e	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:10.010577	f	f	f
1ab02f3d-3591-405a-8c9f-24e3b0600598	bf3b27a3-6e50-465e-b91f-60dd6ab42116	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-11-02 16:58:10.170419	f	f	f
ed1305a7-c546-40b6-8300-c117fc509fa3	e0dd6b24-5b34-45d2-ba8a-4cf32bed5478	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/3.-Manushi-Chhillar.jpg	2024-11-02 16:58:10.330789	f	f	f
7c13763c-97b1-4f0f-8165-fbf0675a4be6	ffcd0a71-fecf-4825-a912-85068534f86a	https://www.caravan.kz/wp-content/uploads/images/635838.jpg	2024-11-02 16:58:10.493168	f	f	f
03275197-e4f2-4f68-a418-ac380fdf3cea	9c4332a6-2485-43a8-a0fb-97681a46846d	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-11-02 16:58:10.645542	f	f	f
b2b41ecb-fb86-4f6a-8eb0-f757844942aa	27f81205-2ed9-4110-8534-c8fb01415a3c	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:10.797902	f	f	f
25bf73ce-74b1-4003-b65a-974931a761f1	f70b3dd8-044c-4935-b098-461fe2020740	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-11-02 16:58:10.948568	f	f	f
be78d6ba-6784-44d0-aa0c-a1664259a1e1	c0f29c32-19f1-4b79-ae9e-63d7a0b2ae69	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:11.099143	f	f	f
1f090d9d-34e0-4e43-aa87-c4b5bff34d65	c2cdb3b3-d455-48a1-9f10-29dd73482baa	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:11.102141	f	f	f
a3ec8b1a-5843-461d-a148-77ed1f67030a	7288a945-eeff-4b12-a3c2-1da560069415	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:11.252621	f	f	f
0699dace-92d1-41c0-af52-a10444700662	9967740d-b9f6-4ee6-a896-7b33aa1cb17a	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:11.255188	f	f	f
34547d5f-0ef6-4158-81c5-967d9fd5ad5c	a139551d-ae5c-4eae-aedb-bdf34ef5dfa0	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:11.403223	f	f	f
e7be8b5f-02d1-4fad-bd1e-025ab16722ab	cd2157a4-45ec-425d-9ce7-09e0a859f6af	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:11.406701	f	f	f
d23d916a-3eb7-4139-b6db-8a22b8063b51	a6d3fce6-aae1-43a0-aafe-945472e166cb	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:11.559993	f	f	f
f5b832f8-2325-4d76-96ee-29ea5513d173	84b49b40-6df6-4cc6-9e56-425dc04be991	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-11-02 16:58:11.562882	f	f	f
c9c6f2ee-b696-4f0b-8b57-610527b7453c	ca59426a-692d-4e56-9ebf-6ba51b2ec0ba	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:11.56577	f	f	f
8a386d56-3589-44b7-a7ce-590627c76b4e	82a15c58-1262-4d60-896f-470e08fc3d46	https://www.caravan.kz/wp-content/uploads/images/635838.jpg	2024-11-02 16:58:11.650825	f	f	f
d471b5cb-c1f7-4d6d-adc9-a95cf17b558c	91112b48-3ef4-4f93-bf50-853fb56e7f06	https://www.caravan.kz/wp-content/uploads/images/635838.jpg	2024-11-02 16:58:11.798869	f	f	f
21c63d16-6543-4bf3-a894-c20c16f71bcd	1540af57-19a3-4df4-b374-75dfa6be8ff7	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:11.878609	f	f	f
60c92198-9042-40e9-bb57-b91e30aa87a6	031ff176-740c-48ea-8547-9034871adb58	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:11.96072	f	f	f
b774cf71-a459-4ece-9e60-1b487a770d03	8fa3e170-b597-49a1-8f21-13cf76492958	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:12.120779	f	f	f
cd02198c-3912-4289-99e5-a33d76b0cc88	ba916b20-a6ee-4564-8af2-398bcb076405	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:12.125148	f	f	f
57918fcd-9514-412e-8e51-5280f308fb0a	0cfec090-0c66-49e7-97f1-fe20cf4375c5	https://www.caravan.kz/wp-content/uploads/images/635838.jpg	2024-11-02 16:58:12.127804	f	f	f
58816d4e-3cc1-4025-abae-5b86ef2eb642	ec6b3d4e-22a4-4884-9770-37b73a864778	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:13.742134	f	f	f
66cd09a9-b24a-43f7-a706-2512c229541b	8d455fe1-ead6-4975-b9ed-b553a7544de2	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:13.926165	f	f	f
1920bac2-a110-4cf2-8c47-8a4208dae7d2	dc154f36-8481-474b-946d-7d3b7b87d814	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:14.181275	f	f	f
73d8c95f-d00a-426e-b8b3-8540eefe4819	c4bb3ebe-9e8c-4e97-8420-c74368ff1160	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/3.-Manushi-Chhillar.jpg	2024-11-02 16:58:14.34708	f	f	f
130d6367-574c-4d01-a796-00ca3b26ae5a	e03b8aa1-6e45-46d3-a46c-4a06d06a1bf3	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:14.587908	f	f	f
f7c98c2c-5720-4c2a-91a4-b64cb8468378	1d0b97ad-5496-474c-b29c-15d0addde672	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-11-02 16:58:14.748801	f	f	f
8953a501-d6f7-4c4d-a8f2-164ba69ebbaa	3078172a-8c99-49d5-9007-28a73392192c	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:14.985726	f	f	f
58f08658-2a24-4516-b7b5-24b7a48221ff	7a26834c-06be-40ed-aef8-c45b0a44248b	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/3.-Manushi-Chhillar.jpg	2024-11-02 16:58:15.145198	f	f	f
d019e4f2-8ca0-406a-b0ef-aa1231128fbe	1aa282c7-4dd9-4bef-bd5a-e19731f3cad9	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/Most-Beautiful-Indian-Girls.jpg	2024-11-02 16:58:16.64446	f	f	f
9459205f-8a4e-4a0d-82fb-d773716c088c	b2df1c31-96e7-4291-ad73-0aa354b82c28	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/3.-Manushi-Chhillar.jpg	2024-11-02 16:58:16.791402	f	f	f
f225c594-83a6-418d-9609-6b1da84edbea	22176a73-cb88-468a-89d0-180055a3470e	https://w0.peakpx.com/wallpaper/268/995/HD-wallpaper-nice-girl-beauty-brown-hair-cute-denim-girl-jeans-pretty.jpg	2024-11-02 16:58:16.957176	f	f	f
4be09e8d-3a21-45b1-a411-c126abfe8f63	1b1ca1eb-99a0-46e8-a10a-279efed8fc15	https://www.caravan.kz/wp-content/uploads/images/635838.jpg	2024-11-02 16:58:17.111593	f	f	f
010a6ae1-31e4-4beb-8955-a7e78dc9d36d	53847d2a-b3f9-48f4-962b-aa741a89b623	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:17.346886	f	f	f
eb15eb45-93c5-4668-9a33-d79a00450606	97ca9959-6cf4-44a2-8d58-ca060276071b	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:17.579284	f	f	f
3860e722-705a-4878-8a11-630bb66db4d8	3b1206ce-87c0-4aac-99d1-85a514bed337	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/3.-Manushi-Chhillar.jpg	2024-11-02 16:58:17.811877	f	f	f
01a0f006-de48-4319-b57f-38d2a332e2a4	39141384-3ade-407e-b0f7-9a9111c8472f	https://cdn2.stylecraze.com/wp-content/uploads/2013/10/3.-Manushi-Chhillar.jpg	2024-11-02 16:58:18.043151	f	f	f
8b9d0ed1-db7c-41f7-9fde-a14ba0d984a4	22fe87b7-996d-4e60-99cc-55b44b198393	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:18.268898	f	f	f
1b6be336-0268-44b4-a434-f51670f1ace5	c16bd828-f0cc-4467-b9ba-4605230d433e	https://www.m24.ru/b/d/nBkSUhL2hFYhm8yyJr6BrNOp2Z3z8Zj21iDEh_fH_nKUPXuaDyXTjHou4MVO6BCVoZKf9GqVe5Q_CPawk214LyWK9G1N5ho=rX-X0R4iFr289CvBvD6lVQ.jpg	2024-11-02 16:58:18.505344	f	f	f
\.


--
-- Data for Name: profile_body_arts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profile_body_arts (profile_id, body_art_id) FROM stdin;
05029a37-875c-4f65-8f5a-c95ba1e32e1b	1
05029a37-875c-4f65-8f5a-c95ba1e32e1b	2
f648288a-9756-4c99-8d6e-26a9d929cd3e	1
f648288a-9756-4c99-8d6e-26a9d929cd3e	2
bf3b27a3-6e50-465e-b91f-60dd6ab42116	1
bf3b27a3-6e50-465e-b91f-60dd6ab42116	2
e0dd6b24-5b34-45d2-ba8a-4cf32bed5478	1
e0dd6b24-5b34-45d2-ba8a-4cf32bed5478	2
ffcd0a71-fecf-4825-a912-85068534f86a	1
ffcd0a71-fecf-4825-a912-85068534f86a	2
9c4332a6-2485-43a8-a0fb-97681a46846d	1
9c4332a6-2485-43a8-a0fb-97681a46846d	2
27f81205-2ed9-4110-8534-c8fb01415a3c	1
27f81205-2ed9-4110-8534-c8fb01415a3c	2
f70b3dd8-044c-4935-b098-461fe2020740	1
f70b3dd8-044c-4935-b098-461fe2020740	2
c0f29c32-19f1-4b79-ae9e-63d7a0b2ae69	1
c0f29c32-19f1-4b79-ae9e-63d7a0b2ae69	2
c2cdb3b3-d455-48a1-9f10-29dd73482baa	1
c2cdb3b3-d455-48a1-9f10-29dd73482baa	2
7288a945-eeff-4b12-a3c2-1da560069415	1
7288a945-eeff-4b12-a3c2-1da560069415	2
9967740d-b9f6-4ee6-a896-7b33aa1cb17a	1
9967740d-b9f6-4ee6-a896-7b33aa1cb17a	2
a139551d-ae5c-4eae-aedb-bdf34ef5dfa0	1
a139551d-ae5c-4eae-aedb-bdf34ef5dfa0	2
cd2157a4-45ec-425d-9ce7-09e0a859f6af	1
cd2157a4-45ec-425d-9ce7-09e0a859f6af	2
a6d3fce6-aae1-43a0-aafe-945472e166cb	1
a6d3fce6-aae1-43a0-aafe-945472e166cb	2
84b49b40-6df6-4cc6-9e56-425dc04be991	1
84b49b40-6df6-4cc6-9e56-425dc04be991	2
ca59426a-692d-4e56-9ebf-6ba51b2ec0ba	1
ca59426a-692d-4e56-9ebf-6ba51b2ec0ba	2
82a15c58-1262-4d60-896f-470e08fc3d46	1
82a15c58-1262-4d60-896f-470e08fc3d46	2
91112b48-3ef4-4f93-bf50-853fb56e7f06	1
91112b48-3ef4-4f93-bf50-853fb56e7f06	2
1540af57-19a3-4df4-b374-75dfa6be8ff7	1
1540af57-19a3-4df4-b374-75dfa6be8ff7	2
031ff176-740c-48ea-8547-9034871adb58	1
031ff176-740c-48ea-8547-9034871adb58	2
8fa3e170-b597-49a1-8f21-13cf76492958	1
8fa3e170-b597-49a1-8f21-13cf76492958	2
ba916b20-a6ee-4564-8af2-398bcb076405	1
ba916b20-a6ee-4564-8af2-398bcb076405	2
0cfec090-0c66-49e7-97f1-fe20cf4375c5	1
0cfec090-0c66-49e7-97f1-fe20cf4375c5	2
ec6b3d4e-22a4-4884-9770-37b73a864778	1
ec6b3d4e-22a4-4884-9770-37b73a864778	2
8d455fe1-ead6-4975-b9ed-b553a7544de2	1
8d455fe1-ead6-4975-b9ed-b553a7544de2	2
dc154f36-8481-474b-946d-7d3b7b87d814	1
dc154f36-8481-474b-946d-7d3b7b87d814	2
c4bb3ebe-9e8c-4e97-8420-c74368ff1160	1
c4bb3ebe-9e8c-4e97-8420-c74368ff1160	2
e03b8aa1-6e45-46d3-a46c-4a06d06a1bf3	1
e03b8aa1-6e45-46d3-a46c-4a06d06a1bf3	2
1d0b97ad-5496-474c-b29c-15d0addde672	1
1d0b97ad-5496-474c-b29c-15d0addde672	2
3078172a-8c99-49d5-9007-28a73392192c	1
3078172a-8c99-49d5-9007-28a73392192c	2
7a26834c-06be-40ed-aef8-c45b0a44248b	1
7a26834c-06be-40ed-aef8-c45b0a44248b	2
1aa282c7-4dd9-4bef-bd5a-e19731f3cad9	1
1aa282c7-4dd9-4bef-bd5a-e19731f3cad9	2
b2df1c31-96e7-4291-ad73-0aa354b82c28	1
b2df1c31-96e7-4291-ad73-0aa354b82c28	2
22176a73-cb88-468a-89d0-180055a3470e	1
22176a73-cb88-468a-89d0-180055a3470e	2
1b1ca1eb-99a0-46e8-a10a-279efed8fc15	1
1b1ca1eb-99a0-46e8-a10a-279efed8fc15	2
53847d2a-b3f9-48f4-962b-aa741a89b623	1
53847d2a-b3f9-48f4-962b-aa741a89b623	2
97ca9959-6cf4-44a2-8d58-ca060276071b	1
97ca9959-6cf4-44a2-8d58-ca060276071b	2
3b1206ce-87c0-4aac-99d1-85a514bed337	1
3b1206ce-87c0-4aac-99d1-85a514bed337	2
39141384-3ade-407e-b0f7-9a9111c8472f	1
39141384-3ade-407e-b0f7-9a9111c8472f	2
22fe87b7-996d-4e60-99cc-55b44b198393	1
22fe87b7-996d-4e60-99cc-55b44b198393	2
c16bd828-f0cc-4467-b9ba-4605230d433e	1
c16bd828-f0cc-4467-b9ba-4605230d433e	2
\.


--
-- Data for Name: profile_options; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profile_options (profile_id, profile_tag_id, price, comment) FROM stdin;
05029a37-875c-4f65-8f5a-c95ba1e32e1b	1	5000	This is my favourite!
05029a37-875c-4f65-8f5a-c95ba1e32e1b	2	50000	I hate this!
f648288a-9756-4c99-8d6e-26a9d929cd3e	1	5000	This is my favourite!
f648288a-9756-4c99-8d6e-26a9d929cd3e	2	50000	I hate this!
bf3b27a3-6e50-465e-b91f-60dd6ab42116	1	5000	This is my favourite!
bf3b27a3-6e50-465e-b91f-60dd6ab42116	2	50000	I hate this!
e0dd6b24-5b34-45d2-ba8a-4cf32bed5478	1	5000	This is my favourite!
e0dd6b24-5b34-45d2-ba8a-4cf32bed5478	2	50000	I hate this!
ffcd0a71-fecf-4825-a912-85068534f86a	1	5000	This is my favourite!
ffcd0a71-fecf-4825-a912-85068534f86a	2	50000	I hate this!
9c4332a6-2485-43a8-a0fb-97681a46846d	1	5000	This is my favourite!
9c4332a6-2485-43a8-a0fb-97681a46846d	2	50000	I hate this!
27f81205-2ed9-4110-8534-c8fb01415a3c	1	5000	This is my favourite!
27f81205-2ed9-4110-8534-c8fb01415a3c	2	50000	I hate this!
f70b3dd8-044c-4935-b098-461fe2020740	1	5000	This is my favourite!
f70b3dd8-044c-4935-b098-461fe2020740	2	50000	I hate this!
c0f29c32-19f1-4b79-ae9e-63d7a0b2ae69	1	5000	This is my favourite!
c0f29c32-19f1-4b79-ae9e-63d7a0b2ae69	2	50000	I hate this!
c2cdb3b3-d455-48a1-9f10-29dd73482baa	1	5000	This is my favourite!
c2cdb3b3-d455-48a1-9f10-29dd73482baa	2	50000	I hate this!
7288a945-eeff-4b12-a3c2-1da560069415	1	5000	This is my favourite!
7288a945-eeff-4b12-a3c2-1da560069415	2	50000	I hate this!
9967740d-b9f6-4ee6-a896-7b33aa1cb17a	1	5000	This is my favourite!
9967740d-b9f6-4ee6-a896-7b33aa1cb17a	2	50000	I hate this!
a139551d-ae5c-4eae-aedb-bdf34ef5dfa0	1	5000	This is my favourite!
a139551d-ae5c-4eae-aedb-bdf34ef5dfa0	2	50000	I hate this!
cd2157a4-45ec-425d-9ce7-09e0a859f6af	1	5000	This is my favourite!
cd2157a4-45ec-425d-9ce7-09e0a859f6af	2	50000	I hate this!
a6d3fce6-aae1-43a0-aafe-945472e166cb	1	5000	This is my favourite!
a6d3fce6-aae1-43a0-aafe-945472e166cb	2	50000	I hate this!
84b49b40-6df6-4cc6-9e56-425dc04be991	1	5000	This is my favourite!
84b49b40-6df6-4cc6-9e56-425dc04be991	2	50000	I hate this!
ca59426a-692d-4e56-9ebf-6ba51b2ec0ba	1	5000	This is my favourite!
ca59426a-692d-4e56-9ebf-6ba51b2ec0ba	2	50000	I hate this!
82a15c58-1262-4d60-896f-470e08fc3d46	1	5000	This is my favourite!
82a15c58-1262-4d60-896f-470e08fc3d46	2	50000	I hate this!
91112b48-3ef4-4f93-bf50-853fb56e7f06	1	5000	This is my favourite!
91112b48-3ef4-4f93-bf50-853fb56e7f06	2	50000	I hate this!
1540af57-19a3-4df4-b374-75dfa6be8ff7	1	5000	This is my favourite!
1540af57-19a3-4df4-b374-75dfa6be8ff7	2	50000	I hate this!
031ff176-740c-48ea-8547-9034871adb58	1	5000	This is my favourite!
031ff176-740c-48ea-8547-9034871adb58	2	50000	I hate this!
8fa3e170-b597-49a1-8f21-13cf76492958	1	5000	This is my favourite!
8fa3e170-b597-49a1-8f21-13cf76492958	2	50000	I hate this!
ba916b20-a6ee-4564-8af2-398bcb076405	1	5000	This is my favourite!
ba916b20-a6ee-4564-8af2-398bcb076405	2	50000	I hate this!
0cfec090-0c66-49e7-97f1-fe20cf4375c5	1	5000	This is my favourite!
0cfec090-0c66-49e7-97f1-fe20cf4375c5	2	50000	I hate this!
ec6b3d4e-22a4-4884-9770-37b73a864778	1	5000	This is my favourite!
ec6b3d4e-22a4-4884-9770-37b73a864778	2	50000	I hate this!
8d455fe1-ead6-4975-b9ed-b553a7544de2	1	5000	This is my favourite!
8d455fe1-ead6-4975-b9ed-b553a7544de2	2	50000	I hate this!
dc154f36-8481-474b-946d-7d3b7b87d814	1	5000	This is my favourite!
dc154f36-8481-474b-946d-7d3b7b87d814	2	50000	I hate this!
c4bb3ebe-9e8c-4e97-8420-c74368ff1160	1	5000	This is my favourite!
c4bb3ebe-9e8c-4e97-8420-c74368ff1160	2	50000	I hate this!
e03b8aa1-6e45-46d3-a46c-4a06d06a1bf3	1	5000	This is my favourite!
e03b8aa1-6e45-46d3-a46c-4a06d06a1bf3	2	50000	I hate this!
1d0b97ad-5496-474c-b29c-15d0addde672	1	5000	This is my favourite!
1d0b97ad-5496-474c-b29c-15d0addde672	2	50000	I hate this!
3078172a-8c99-49d5-9007-28a73392192c	1	5000	This is my favourite!
3078172a-8c99-49d5-9007-28a73392192c	2	50000	I hate this!
7a26834c-06be-40ed-aef8-c45b0a44248b	1	5000	This is my favourite!
7a26834c-06be-40ed-aef8-c45b0a44248b	2	50000	I hate this!
1aa282c7-4dd9-4bef-bd5a-e19731f3cad9	1	5000	This is my favourite!
1aa282c7-4dd9-4bef-bd5a-e19731f3cad9	2	50000	I hate this!
b2df1c31-96e7-4291-ad73-0aa354b82c28	1	5000	This is my favourite!
b2df1c31-96e7-4291-ad73-0aa354b82c28	2	50000	I hate this!
22176a73-cb88-468a-89d0-180055a3470e	1	5000	This is my favourite!
22176a73-cb88-468a-89d0-180055a3470e	2	50000	I hate this!
1b1ca1eb-99a0-46e8-a10a-279efed8fc15	1	5000	This is my favourite!
1b1ca1eb-99a0-46e8-a10a-279efed8fc15	2	50000	I hate this!
53847d2a-b3f9-48f4-962b-aa741a89b623	1	5000	This is my favourite!
53847d2a-b3f9-48f4-962b-aa741a89b623	2	50000	I hate this!
97ca9959-6cf4-44a2-8d58-ca060276071b	1	5000	This is my favourite!
97ca9959-6cf4-44a2-8d58-ca060276071b	2	50000	I hate this!
3b1206ce-87c0-4aac-99d1-85a514bed337	1	5000	This is my favourite!
3b1206ce-87c0-4aac-99d1-85a514bed337	2	50000	I hate this!
39141384-3ade-407e-b0f7-9a9111c8472f	1	5000	This is my favourite!
39141384-3ade-407e-b0f7-9a9111c8472f	2	50000	I hate this!
22fe87b7-996d-4e60-99cc-55b44b198393	1	5000	This is my favourite!
22fe87b7-996d-4e60-99cc-55b44b198393	2	50000	I hate this!
c16bd828-f0cc-4467-b9ba-4605230d433e	1	5000	This is my favourite!
c16bd828-f0cc-4467-b9ba-4605230d433e	2	50000	I hate this!
\.


--
-- Data for Name: profile_ratings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profile_ratings (id, service_id, profile_id, review_text_visible, review, score, created_at, updated_at, updated_by) FROM stdin;
dcf53860-440c-4f2d-96ea-5cbd4e934b7f	710a92d7-9bd1-45bc-8294-eb3ab0150c71	ec6b3d4e-22a4-4884-9770-37b73a864778	t	I like the service! It's very good	5	2024-10-31 12:58:13.7754	2024-11-02 16:58:13.747244	00000000-0000-0000-0000-000000000000
a57cd78f-7325-4de7-88d4-82272866f8ea	2c1834a2-d2ff-4e59-adca-08fb52fb09cf	8d455fe1-ead6-4975-b9ed-b553a7544de2	t	I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life	1	2024-11-01 13:58:13.936589	2024-11-02 16:58:13.955809	00000000-0000-0000-0000-000000000000
87646b7d-8964-418c-9b3d-9bc577d44d74	ada0441f-f25a-4201-ad31-332055054882	dc154f36-8481-474b-946d-7d3b7b87d814	t	I like the service! It's very good	5	2024-11-02 16:58:14.183801	2024-11-02 16:58:14.183801	00000000-0000-0000-0000-000000000000
cf3bb4b0-137a-4014-b633-3bbd58cf83e8	6339e711-1508-4316-b20a-c141e2f35913	c4bb3ebe-9e8c-4e97-8420-c74368ff1160	f	I like the service! It's very good	5	2024-11-02 16:58:14.349267	2024-11-02 16:58:14.359234	00000000-0000-0000-0000-000000000000
91c5d5ce-163f-4a5f-acea-a6d9243042d5	1d0d8210-2d9d-404b-b736-197788eae0b2	e03b8aa1-6e45-46d3-a46c-4a06d06a1bf3	t	I like the service! It's very good	5	2024-11-02 16:58:14.590263	2024-11-02 16:58:14.590263	00000000-0000-0000-0000-000000000000
d3c281dc-5024-4a5e-bae8-1fff5862079b	1e097547-1994-465a-a547-23cf6b90b6b9	1d0b97ad-5496-474c-b29c-15d0addde672	t	I like the service! It's very good	5	2024-11-02 16:58:14.751243	2024-11-02 16:58:14.751243	00000000-0000-0000-0000-000000000000
62693bd1-8ae3-403f-8b4e-960dc0523021	5e2ebc69-16bd-44d0-8bb2-4fc832afb8da	3078172a-8c99-49d5-9007-28a73392192c	t	I like the service! It's very good	5	2024-11-02 16:58:14.988045	2024-11-02 16:58:14.988045	00000000-0000-0000-0000-000000000000
9c1e57a5-e2ab-48aa-a562-2dc22e3067a8	06b1522d-8634-4ca5-8925-7bb4e25a290f	7a26834c-06be-40ed-aef8-c45b0a44248b	t	I like the service! It's very good	5	2024-11-02 16:58:15.147207	2024-11-02 16:58:15.147207	00000000-0000-0000-0000-000000000000
ea7a5d23-55b5-4f5b-bdd2-dc26b7478e31	5d6ee30a-6347-40d1-8a10-0ef7a9ae4376	1b1ca1eb-99a0-46e8-a10a-279efed8fc15	t	I like the service! It's very good	5	2024-11-02 16:58:17.113948	2024-11-02 16:58:17.113948	00000000-0000-0000-0000-000000000000
122fcf10-1bf5-4186-bfac-7fca5399b5c0	ff2f02c9-8c49-49d7-bbc7-fab7d3a8f4ee	53847d2a-b3f9-48f4-962b-aa741a89b623	t	I like the service! It's very good	5	2024-11-02 16:58:17.349482	2024-11-02 16:58:17.349482	00000000-0000-0000-0000-000000000000
105f7832-7074-497b-bcb9-35c658330cf9	b17b867a-d8f9-4f74-981d-1959149a7d59	97ca9959-6cf4-44a2-8d58-ca060276071b	t	I like the service! It's very good	5	2024-11-02 16:58:17.581741	2024-11-02 16:58:17.581741	00000000-0000-0000-0000-000000000000
f838ef44-4fd2-4e6c-b8ae-6e0a234c515b	d30f3cc0-53ae-434f-8960-b6771c734f97	3b1206ce-87c0-4aac-99d1-85a514bed337	t	I like the service! It's very good	5	2024-11-02 16:58:17.815186	2024-11-02 16:58:17.815186	00000000-0000-0000-0000-000000000000
0294fd73-4971-4709-9338-9e08ab612119	b10592f1-dba1-48d8-85bf-1b6f05e557d9	39141384-3ade-407e-b0f7-9a9111c8472f	t	I like the service! It's very good	5	2024-11-02 16:58:18.045778	2024-11-02 16:58:18.045778	00000000-0000-0000-0000-000000000000
94ac38c8-c38f-4c82-a702-a4bf95f0c00a	fe4bc93e-7106-4b46-be57-1d2307a1ff95	22fe87b7-996d-4e60-99cc-55b44b198393	t	I like the service! It's very good	5	2024-11-02 16:58:18.271057	2024-11-02 16:58:18.271057	00000000-0000-0000-0000-000000000000
0c442ac3-2a96-4d61-82a2-01e4d5782cfc	fd86ddb2-dcd8-4145-b5ec-f40a4626711a	c16bd828-f0cc-4467-b9ba-4605230d433e	t	I like the service! It's very good	5	2024-11-02 16:58:18.507892	2024-11-02 16:58:18.507892	00000000-0000-0000-0000-000000000000
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
\N	21	2	1	36		05029a37-875c-4f65-8f5a-c95ba1e32e1b	02086cae-d58e-4671-b0ba-de3c8a8f8034	t	6919518254	AlicehfKl2bCdndTvWBM	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:09.768696	2024-11-02 16:58:09.768696	02086cae-d58e-4671-b0ba-de3c8a8f8034	\N
3	46	1	2	37		f648288a-9756-4c99-8d6e-26a9d929cd3e	073c8d62-2d6a-482f-9e2b-30bb89f5c363	f	7019899576	AlicegaCWKSP1YQkNVOh-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:10.008739	2024-11-02 16:58:10.014756	073c8d62-2d6a-482f-9e2b-30bb89f5c363	\N
1	67	4	2	20		bf3b27a3-6e50-465e-b91f-60dd6ab42116	6b15f1d8-1889-4a51-b78f-64f9d40fef5b	f	8823610007	AlicesWDPnzx088J7KHJ-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	t	2024-11-02 16:58:10.172967	17ddaab0-8bbb-4c89-9be0-d1dbe00fb29b	f	2024-11-02 16:58:10.172967	17ddaab0-8bbb-4c89-9be0-d1dbe00fb29b	2024-11-02 16:58:10.1693	2024-11-02 16:58:10.172967	17ddaab0-8bbb-4c89-9be0-d1dbe00fb29b	\N
4	47	7	2	39		e0dd6b24-5b34-45d2-ba8a-4cf32bed5478	cf48d8de-1e4f-45ae-ac0f-b893ec450e88	f	4064631383	AliceWYKOjUU3vd1tHUv-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	t	2024-11-02 16:58:10.333082	67c55b7a-b156-434c-9751-c86f3dce6d9c	f	2024-11-02 16:58:10.333082	67c55b7a-b156-434c-9751-c86f3dce6d9c	2024-11-02 16:58:10.329193	2024-11-02 16:58:10.333082	67c55b7a-b156-434c-9751-c86f3dce6d9c	\N
1	72	1	1	40		ffcd0a71-fecf-4825-a912-85068534f86a	8e604bb7-5eea-4e61-b2b6-c7b14f35e8c9	t	7135049247	AliceF8L4PmDW1f64sFk	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:10.491284	2024-11-02 16:58:10.491284	8e604bb7-5eea-4e61-b2b6-c7b14f35e8c9	\N
4	24	5	2	33		9c4332a6-2485-43a8-a0fb-97681a46846d	4fe446a9-3f75-4774-8128-38980369b53a	t	7889396908	AlicetO8OXi8TtoiwxAv	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:10.644567	2024-11-02 16:58:10.644567	4fe446a9-3f75-4774-8128-38980369b53a	\N
1	12	7	3	46		27f81205-2ed9-4110-8534-c8fb01415a3c	6d6ed2f9-e2f7-4146-a1dc-f0761f341474	t	2925866974	AliceV74pmzJUdy48eTu	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:10.796625	2024-11-02 16:58:10.796625	6d6ed2f9-e2f7-4146-a1dc-f0761f341474	\N
2	78	4	2	37		f70b3dd8-044c-4935-b098-461fe2020740	b95b0b61-129c-4aae-bb07-ac46a419678a	t	6221707810	Alicevh5C1VdxOf5ASEa	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:10.947717	2024-11-02 16:58:10.947717	b95b0b61-129c-4aae-bb07-ac46a419678a	\N
4	19	1	1	14		c0f29c32-19f1-4b79-ae9e-63d7a0b2ae69	c5d0b456-eb56-47a5-8375-eeef2dd11a00	t	9837904064	AliceD3ec95RKIuhP8r0	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.098082	2024-11-02 16:58:11.098082	c5d0b456-eb56-47a5-8375-eeef2dd11a00	\N
4	39	6	2	15		c2cdb3b3-d455-48a1-9f10-29dd73482baa	c5d0b456-eb56-47a5-8375-eeef2dd11a00	t	1561839914	AlicexoLjNMUxNg37C50	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.101398	2024-11-02 16:58:11.101398	c5d0b456-eb56-47a5-8375-eeef2dd11a00	\N
2	65	5	1	44		7288a945-eeff-4b12-a3c2-1da560069415	5a0d02c8-9f87-4f23-a251-70817ac37ccf	t	3662702217	AliceXHQKkkIsPcuASzB	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.251856	2024-11-02 16:58:11.251856	5a0d02c8-9f87-4f23-a251-70817ac37ccf	\N
1	38	6	1	37		9967740d-b9f6-4ee6-a896-7b33aa1cb17a	5a0d02c8-9f87-4f23-a251-70817ac37ccf	t	2208988952	AlicewGF5DOHQUGUSiKQ	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.254483	2024-11-02 16:58:11.254483	5a0d02c8-9f87-4f23-a251-70817ac37ccf	\N
3	2	5	2	48		a139551d-ae5c-4eae-aedb-bdf34ef5dfa0	d6089edd-e7c9-44a2-9486-337ecd67c814	t	8062953707	Alice7NicpIem2dqFWJP	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.402459	2024-11-02 16:58:11.402459	d6089edd-e7c9-44a2-9486-337ecd67c814	\N
1	23	3	2	12		cd2157a4-45ec-425d-9ce7-09e0a859f6af	d6089edd-e7c9-44a2-9486-337ecd67c814	t	7101991215	Alicegkx2cRTyk49PvTf	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.405786	2024-11-02 16:58:11.405786	d6089edd-e7c9-44a2-9486-337ecd67c814	\N
4	43	2	3	39		a6d3fce6-aae1-43a0-aafe-945472e166cb	1b3124df-a831-48ee-ac33-598ab00bcede	t	4070926720	AlicerZUC4TIXKs88h26	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.559166	2024-11-02 16:58:11.559166	1b3124df-a831-48ee-ac33-598ab00bcede	\N
3	7	1	1	24		84b49b40-6df6-4cc6-9e56-425dc04be991	1b3124df-a831-48ee-ac33-598ab00bcede	t	3215829600	AlicezpVBxkmu12hnINN	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.56203	2024-11-02 16:58:11.56203	1b3124df-a831-48ee-ac33-598ab00bcede	\N
3	19	3	3	15		ca59426a-692d-4e56-9ebf-6ba51b2ec0ba	1b3124df-a831-48ee-ac33-598ab00bcede	f	8747089065	AliceCzE19UWx8bbnDSu-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.564932	2024-11-02 16:58:11.568278	6c5541f9-5168-42eb-97c6-6c6db0abe214	\N
1	86	6	2	25		82a15c58-1262-4d60-896f-470e08fc3d46	e6196cf4-a677-4446-9435-3dbdd7859a27	t	9701964594	AliceLwtGAS28VaEuQog	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.649678	2024-11-02 16:58:11.649678	e6196cf4-a677-4446-9435-3dbdd7859a27	\N
1	74	7	2	9		91112b48-3ef4-4f93-bf50-853fb56e7f06	7ac79be6-463d-4bc1-b5b8-b90c7559596b	t	2614743896	AliceWvvR0EjLF2O3Akc	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.797718	2024-11-02 16:58:11.797718	7ac79be6-463d-4bc1-b5b8-b90c7559596b	\N
3	86	6	1	6		1540af57-19a3-4df4-b374-75dfa6be8ff7	34820fab-d019-412c-bb28-52a7039dc461	t	3374992653	AliceugOwj7ng9ZIp4Ld	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.877381	2024-11-02 16:58:11.877381	34820fab-d019-412c-bb28-52a7039dc461	2024-11-02 13:58:11.883809+00
1	64	1	3	42		031ff176-740c-48ea-8547-9034871adb58	e6f112ed-a0d2-434f-b50d-4b7ca1d226ab	t	8212634671	AliceDUwIL14BUrJV8lg	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:11.959693	2024-11-02 16:58:11.959693	e6f112ed-a0d2-434f-b50d-4b7ca1d226ab	\N
1	76	2	1	15		8fa3e170-b597-49a1-8f21-13cf76492958	e4bd738a-a772-4f08-8aff-f959ca3a2b42	t	2166761295	AliceSa4g2tWHsXOcbl4	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:12.119235	2024-11-02 16:58:12.119235	e4bd738a-a772-4f08-8aff-f959ca3a2b42	\N
2	93	1	3	11		ba916b20-a6ee-4564-8af2-398bcb076405	e4bd738a-a772-4f08-8aff-f959ca3a2b42	t	1005706613	AliceMWJVqPcARSH6VdY	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:12.123757	2024-11-02 16:58:12.123757	e4bd738a-a772-4f08-8aff-f959ca3a2b42	\N
3	27	4	3	28		0cfec090-0c66-49e7-97f1-fe20cf4375c5	e4bd738a-a772-4f08-8aff-f959ca3a2b42	f	9529356214	AliceqkAeb2R0qaawgG9-new	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:12.126946	2024-11-02 16:58:12.130067	43a9ff3d-ccd1-4100-9b00-911f069ca195	\N
\N	19	2	1	42		ec6b3d4e-22a4-4884-9770-37b73a864778	e86ca8f3-f8da-4199-b4d7-d8e0f7f4a626	t	7755193770	AlicesvpTsSCu63ZIQMV	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:13.732697	2024-11-02 16:58:13.732697	e86ca8f3-f8da-4199-b4d7-d8e0f7f4a626	\N
\N	8	5	2	12		8d455fe1-ead6-4975-b9ed-b553a7544de2	054c1165-070e-4e62-ba0f-3d46d266d82a	t	3310817250	Alice70dVfW8AOzf1gmE	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:13.924774	2024-11-02 16:58:13.924774	054c1165-070e-4e62-ba0f-3d46d266d82a	\N
\N	71	3	2	24		dc154f36-8481-474b-946d-7d3b7b87d814	3bc5d319-83ad-4cbe-a882-b33e1dca16c7	t	6402083666	AliceP9wgXQk5vTkZAUm	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:14.179952	2024-11-02 16:58:14.179952	3bc5d319-83ad-4cbe-a882-b33e1dca16c7	\N
\N	4	6	2	30		c4bb3ebe-9e8c-4e97-8420-c74368ff1160	16cc1398-460c-42db-a025-3dcd170105cb	t	5095873763	AliceJaSSnLDBITSgxtq	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:14.345752	2024-11-02 16:58:14.345752	16cc1398-460c-42db-a025-3dcd170105cb	\N
\N	47	5	1	33		e03b8aa1-6e45-46d3-a46c-4a06d06a1bf3	5fb13376-fb1c-4796-a246-58caa72311f6	t	2756941175	AliceFUp1wv2G9U93MbW	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:14.586226	2024-11-02 16:58:14.586226	5fb13376-fb1c-4796-a246-58caa72311f6	\N
\N	74	1	2	30		1d0b97ad-5496-474c-b29c-15d0addde672	045eb6ac-555e-4cb0-8319-a29f74801cbb	t	5213854316	AliceYIdQLtW65iTwfAf	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:14.746627	2024-11-02 16:58:14.746627	045eb6ac-555e-4cb0-8319-a29f74801cbb	\N
\N	26	4	2	36		3078172a-8c99-49d5-9007-28a73392192c	a80a4c11-8293-49dc-9f8b-6609ad8c6ecf	t	7865985620	AlicenZSwSsbWkflvL2H	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:14.984288	2024-11-02 16:58:14.984288	a80a4c11-8293-49dc-9f8b-6609ad8c6ecf	\N
\N	41	3	2	22		7a26834c-06be-40ed-aef8-c45b0a44248b	8d01e8ef-fb0f-417f-94c3-efbf77d2da7f	t	6794220164	AlicedVXKa4qdkdCqvko	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:15.144148	2024-11-02 16:58:15.144148	8d01e8ef-fb0f-417f-94c3-efbf77d2da7f	\N
\N	85	2	2	8		1aa282c7-4dd9-4bef-bd5a-e19731f3cad9	2a895964-f41e-4e2d-a6c4-04c4097416b0	t	5183711166	AliceHwlyAwoIJS0Wr7E	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:16.634664	2024-11-02 16:58:16.634664	2a895964-f41e-4e2d-a6c4-04c4097416b0	\N
\N	65	3	1	48		b2df1c31-96e7-4291-ad73-0aa354b82c28	5c87b1e6-906a-4bae-a022-20e6d630788d	t	2291394877	Aliceq97JSS3orzF5mDF	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:16.790082	2024-11-02 16:58:16.790082	5c87b1e6-906a-4bae-a022-20e6d630788d	\N
\N	93	7	2	35		22176a73-cb88-468a-89d0-180055a3470e	b9e32c2d-7b48-401c-bbc3-ff5766fa55ac	t	9970515663	AliceadXlRdPHxKNlU7B	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:16.955704	2024-11-02 16:58:16.955704	b9e32c2d-7b48-401c-bbc3-ff5766fa55ac	\N
\N	91	1	3	43		1b1ca1eb-99a0-46e8-a10a-279efed8fc15	b4568f89-0c9e-466c-a2ff-d8837a3506b5	t	6067319913	AliceljVSNTaQ9biBUvA	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:17.110233	2024-11-02 16:58:17.110233	b4568f89-0c9e-466c-a2ff-d8837a3506b5	\N
\N	57	5	2	37		53847d2a-b3f9-48f4-962b-aa741a89b623	4850299e-3325-4862-9b1b-fd1941b06714	t	5001028800	AliceIc58EBR1CYrKCY7	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:17.345567	2024-11-02 16:58:17.345567	4850299e-3325-4862-9b1b-fd1941b06714	\N
\N	91	4	3	15		97ca9959-6cf4-44a2-8d58-ca060276071b	a85f6e36-fa3b-403a-83f7-27e143b77cf9	t	5769120978	AliceuwJ8IbxllUKTIY4	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:17.578045	2024-11-02 16:58:17.578045	a85f6e36-fa3b-403a-83f7-27e143b77cf9	\N
\N	44	1	2	7		3b1206ce-87c0-4aac-99d1-85a514bed337	f0da98d6-834f-45c9-b70b-e0c54f057aad	t	8011151173	Alice9W79iyY7p1CQKNe	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:17.810081	2024-11-02 16:58:17.810081	f0da98d6-834f-45c9-b70b-e0c54f057aad	\N
\N	48	7	3	27		39141384-3ade-407e-b0f7-9a9111c8472f	7702b59e-ea14-4a2e-b194-691e89e50ba4	t	6697150892	AliceZ4NgiELLvEV26AZ	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:18.042029	2024-11-02 16:58:18.042029	7702b59e-ea14-4a2e-b194-691e89e50ba4	\N
\N	17	7	3	19		22fe87b7-996d-4e60-99cc-55b44b198393	ce55f6db-022c-4bbb-9965-502a0e284ccb	t	9634553016	AliceFCxBkwOO8TNATd6	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:18.267742	2024-11-02 16:58:18.267742	ce55f6db-022c-4bbb-9965-502a0e284ccb	\N
\N	90	3	3	8		c16bd828-f0cc-4467-b9ba-4605230d433e	6c522307-e361-4df7-be1e-357f99487224	t	2434244418	Alice7eWDikabTgBcwDR	29	170	57	2.5	Hey :wave: My name is Lola and I'm 18 years old. I read a lot and like cooking.			1	10000	20000	1	\N	\N	1	\N	\N	1	\N	\N	77073778123		@lovely_mika	f	\N	\N	f	\N	\N	2024-11-02 16:58:18.504162	2024-11-02 16:58:18.504162	6c522307-e361-4df7-be1e-357f99487224	\N
\.


--
-- Data for Name: rated_profile_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.rated_profile_tags (rating_id, profile_tag_id, type) FROM stdin;
dcf53860-440c-4f2d-96ea-5cbd4e934b7f	1	like
dcf53860-440c-4f2d-96ea-5cbd4e934b7f	2	like
a57cd78f-7325-4de7-88d4-82272866f8ea	1	dislike
87646b7d-8964-418c-9b3d-9bc577d44d74	1	like
87646b7d-8964-418c-9b3d-9bc577d44d74	2	like
cf3bb4b0-137a-4014-b633-3bbd58cf83e8	1	like
cf3bb4b0-137a-4014-b633-3bbd58cf83e8	2	like
91c5d5ce-163f-4a5f-acea-a6d9243042d5	1	like
91c5d5ce-163f-4a5f-acea-a6d9243042d5	2	like
d3c281dc-5024-4a5e-bae8-1fff5862079b	1	like
d3c281dc-5024-4a5e-bae8-1fff5862079b	2	like
62693bd1-8ae3-403f-8b4e-960dc0523021	1	like
62693bd1-8ae3-403f-8b4e-960dc0523021	2	like
9c1e57a5-e2ab-48aa-a562-2dc22e3067a8	1	like
9c1e57a5-e2ab-48aa-a562-2dc22e3067a8	2	like
ea7a5d23-55b5-4f5b-bdd2-dc26b7478e31	1	like
ea7a5d23-55b5-4f5b-bdd2-dc26b7478e31	2	like
122fcf10-1bf5-4186-bfac-7fca5399b5c0	1	like
122fcf10-1bf5-4186-bfac-7fca5399b5c0	2	like
105f7832-7074-497b-bcb9-35c658330cf9	1	like
105f7832-7074-497b-bcb9-35c658330cf9	2	like
f838ef44-4fd2-4e6c-b8ae-6e0a234c515b	1	like
f838ef44-4fd2-4e6c-b8ae-6e0a234c515b	2	like
0294fd73-4971-4709-9338-9e08ab612119	1	like
0294fd73-4971-4709-9338-9e08ab612119	2	like
94ac38c8-c38f-4c82-a702-a4bf95f0c00a	1	like
94ac38c8-c38f-4c82-a702-a4bf95f0c00a	2	like
0c442ac3-2a96-4d61-82a2-01e4d5782cfc	1	like
0c442ac3-2a96-4d61-82a2-01e4d5782cfc	2	like
\.


--
-- Data for Name: rated_user_tags; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.rated_user_tags (rating_id, user_tag_id, type) FROM stdin;
c0d2065f-ad18-49c9-9a2e-3f4d7ea7c1b5	1	like
c0d2065f-ad18-49c9-9a2e-3f4d7ea7c1b5	2	dislike
c3e2a4a6-7169-4965-ae22-e56b7a12d5ce	1	like
c3e2a4a6-7169-4965-ae22-e56b7a12d5ce	2	dislike
c1a1b2bb-1e10-4c47-b508-75a7247176e8	1	like
c1a1b2bb-1e10-4c47-b508-75a7247176e8	2	dislike
412590b1-6dfb-4fea-a437-141eba71c3d8	1	like
412590b1-6dfb-4fea-a437-141eba71c3d8	2	dislike
7cb59f20-d55f-4cc2-918f-4b0821e32b72	1	like
7cb59f20-d55f-4cc2-918f-4b0821e32b72	2	dislike
b77104fa-8104-4b7c-a6a0-27468477024e	1	dislike
86a2193c-54a1-4c96-b79b-2e0a71528714	1	like
86a2193c-54a1-4c96-b79b-2e0a71528714	2	dislike
983b95cf-abf5-4b4b-bdcc-0e7bb19d5463	1	like
983b95cf-abf5-4b4b-bdcc-0e7bb19d5463	2	dislike
2d2801cb-c07b-46ec-b75e-ebc1a00a0bfc	7	like
2d2801cb-c07b-46ec-b75e-ebc1a00a0bfc	8	dislike
a245bf08-30eb-485a-ba1f-f9097f081448	7	like
a245bf08-30eb-485a-ba1f-f9097f081448	8	dislike
7656fdfb-18fe-4831-bf5b-ec246c8886f3	7	like
7656fdfb-18fe-4831-bf5b-ec246c8886f3	8	dislike
4262581e-99c1-42b7-892c-6c21415d2617	7	like
4262581e-99c1-42b7-892c-6c21415d2617	8	dislike
fa268b21-e77e-476c-a943-3d19607bb622	7	like
fa268b21-e77e-476c-a943-3d19607bb622	8	dislike
bc118170-d068-4659-81c1-59876de7d78e	7	like
bc118170-d068-4659-81c1-59876de7d78e	8	dislike
a0cb44c9-793e-4c0b-9ae5-3198ac70fd2a	7	like
a0cb44c9-793e-4c0b-9ae5-3198ac70fd2a	8	dislike
\.


--
-- Data for Name: services; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.services (id, client_user_id, client_user_rating_id, client_user_lat, client_user_lon, profile_id, profile_owner_id, profile_rating_id, profile_user_lat, profile_user_lon, distance_between_users, trusted_distance, created_at, updated_at, updated_by) FROM stdin;
710a92d7-9bd1-45bc-8294-eb3ab0150c71	1a513cda-5d18-4d9d-b300-a237423723de	c0d2065f-ad18-49c9-9a2e-3f4d7ea7c1b5	43.25977	76.93525	ec6b3d4e-22a4-4884-9770-37b73a864778	e86ca8f3-f8da-4199-b4d7-d8e0f7f4a626	dcf53860-440c-4f2d-96ea-5cbd4e934b7f	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:13.747244	2024-11-02 16:58:13.765109	e86ca8f3-f8da-4199-b4d7-d8e0f7f4a626
2c1834a2-d2ff-4e59-adca-08fb52fb09cf	0a63a3e8-3846-461b-bd05-a321ee269598	c3e2a4a6-7169-4965-ae22-e56b7a12d5ce	43.25977	76.93525	8d455fe1-ead6-4975-b9ed-b553a7544de2	054c1165-070e-4e62-ba0f-3d46d266d82a	a57cd78f-7325-4de7-88d4-82272866f8ea	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:13.928674	2024-11-02 16:58:13.931761	054c1165-070e-4e62-ba0f-3d46d266d82a
ada0441f-f25a-4201-ad31-332055054882	fec33571-8503-4a88-9d20-006970d0d1eb	c1a1b2bb-1e10-4c47-b508-75a7247176e8	43.25977	76.93525	dc154f36-8481-474b-946d-7d3b7b87d814	3bc5d319-83ad-4cbe-a882-b33e1dca16c7	87646b7d-8964-418c-9b3d-9bc577d44d74	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:14.183801	2024-11-02 16:58:14.186499	3bc5d319-83ad-4cbe-a882-b33e1dca16c7
6339e711-1508-4316-b20a-c141e2f35913	b556d2ea-283f-45b3-99c5-ab98cb8cebe4	412590b1-6dfb-4fea-a437-141eba71c3d8	43.25977	76.93525	c4bb3ebe-9e8c-4e97-8420-c74368ff1160	16cc1398-460c-42db-a025-3dcd170105cb	cf3bb4b0-137a-4014-b633-3bbd58cf83e8	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:14.349267	2024-11-02 16:58:14.351809	16cc1398-460c-42db-a025-3dcd170105cb
1d0d8210-2d9d-404b-b736-197788eae0b2	695293b4-b8c9-428d-b309-a183f37f8251	7cb59f20-d55f-4cc2-918f-4b0821e32b72	43.25977	76.93525	e03b8aa1-6e45-46d3-a46c-4a06d06a1bf3	5fb13376-fb1c-4796-a246-58caa72311f6	91c5d5ce-163f-4a5f-acea-a6d9243042d5	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:14.590263	2024-11-02 16:58:14.592508	5fb13376-fb1c-4796-a246-58caa72311f6
1e097547-1994-465a-a547-23cf6b90b6b9	84498241-6b27-4386-9f9c-44e76899404a	b77104fa-8104-4b7c-a6a0-27468477024e	43.25977	76.93525	1d0b97ad-5496-474c-b29c-15d0addde672	045eb6ac-555e-4cb0-8319-a29f74801cbb	d3c281dc-5024-4a5e-bae8-1fff5862079b	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:14.751243	2024-11-02 16:58:14.753572	045eb6ac-555e-4cb0-8319-a29f74801cbb
5e2ebc69-16bd-44d0-8bb2-4fc832afb8da	5ab0fbc7-dcf3-45c5-bab5-38ed2c27c457	86a2193c-54a1-4c96-b79b-2e0a71528714	43.25977	76.93525	3078172a-8c99-49d5-9007-28a73392192c	a80a4c11-8293-49dc-9f8b-6609ad8c6ecf	62693bd1-8ae3-403f-8b4e-960dc0523021	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:14.988045	2024-11-02 16:58:14.991275	a80a4c11-8293-49dc-9f8b-6609ad8c6ecf
06b1522d-8634-4ca5-8925-7bb4e25a290f	976064b4-b126-4c81-8016-11d89f7724d7	983b95cf-abf5-4b4b-bdcc-0e7bb19d5463	43.25977	76.93525	7a26834c-06be-40ed-aef8-c45b0a44248b	8d01e8ef-fb0f-417f-94c3-efbf77d2da7f	9c1e57a5-e2ab-48aa-a562-2dc22e3067a8	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:15.147207	2024-11-02 16:58:15.150777	8d01e8ef-fb0f-417f-94c3-efbf77d2da7f
759a5ac0-e2c8-4e46-b55f-8e8329cfcc57	d9690bd8-8d45-4ee9-b594-117ecf396fb1	\N	43.25977	76.93525	b2df1c31-96e7-4291-ad73-0aa354b82c28	5c87b1e6-906a-4bae-a022-20e6d630788d	\N	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:16.795014	2024-11-02 16:58:16.809067	d9690bd8-8d45-4ee9-b594-117ecf396fb1
798b2976-1555-4e44-86d2-cbdb379ff550	5ea5d3f1-b318-4875-a47b-ee57b9e6bb64	\N	43.25977	76.93525	22176a73-cb88-468a-89d0-180055a3470e	b9e32c2d-7b48-401c-bbc3-ff5766fa55ac	\N	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:16.959736	2024-11-02 16:58:16.960622	b9e32c2d-7b48-401c-bbc3-ff5766fa55ac
5d6ee30a-6347-40d1-8a10-0ef7a9ae4376	6f05dce7-cf3c-411c-b589-a73f07f9099c	2d2801cb-c07b-46ec-b75e-ebc1a00a0bfc	43.25977	76.93525	1b1ca1eb-99a0-46e8-a10a-279efed8fc15	b4568f89-0c9e-466c-a2ff-d8837a3506b5	ea7a5d23-55b5-4f5b-bdd2-dc26b7478e31	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:17.113948	2024-11-02 16:58:17.118579	b4568f89-0c9e-466c-a2ff-d8837a3506b5
ff2f02c9-8c49-49d7-bbc7-fab7d3a8f4ee	692f028b-0d44-4dfd-809e-c2048d268819	a245bf08-30eb-485a-ba1f-f9097f081448	43.25977	76.93525	53847d2a-b3f9-48f4-962b-aa741a89b623	4850299e-3325-4862-9b1b-fd1941b06714	122fcf10-1bf5-4186-bfac-7fca5399b5c0	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:17.349482	2024-11-02 16:58:17.351738	4850299e-3325-4862-9b1b-fd1941b06714
b17b867a-d8f9-4f74-981d-1959149a7d59	e4c62515-a21d-47ec-97d5-85f14cce5aaf	7656fdfb-18fe-4831-bf5b-ec246c8886f3	43.25977	76.93525	97ca9959-6cf4-44a2-8d58-ca060276071b	a85f6e36-fa3b-403a-83f7-27e143b77cf9	105f7832-7074-497b-bcb9-35c658330cf9	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:17.581741	2024-11-02 16:58:17.58425	a85f6e36-fa3b-403a-83f7-27e143b77cf9
d30f3cc0-53ae-434f-8960-b6771c734f97	e739d6a7-4eea-46ed-9efc-a9f7d5e3e6a3	4262581e-99c1-42b7-892c-6c21415d2617	43.25977	76.93525	3b1206ce-87c0-4aac-99d1-85a514bed337	f0da98d6-834f-45c9-b70b-e0c54f057aad	f838ef44-4fd2-4e6c-b8ae-6e0a234c515b	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:17.815186	2024-11-02 16:58:17.817802	f0da98d6-834f-45c9-b70b-e0c54f057aad
b10592f1-dba1-48d8-85bf-1b6f05e557d9	e68716b2-3a90-454a-8720-5d71da7155cb	fa268b21-e77e-476c-a943-3d19607bb622	43.25977	76.93525	39141384-3ade-407e-b0f7-9a9111c8472f	7702b59e-ea14-4a2e-b194-691e89e50ba4	0294fd73-4971-4709-9338-9e08ab612119	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:18.045778	2024-11-02 16:58:18.047935	7702b59e-ea14-4a2e-b194-691e89e50ba4
fe4bc93e-7106-4b46-be57-1d2307a1ff95	b9cfb846-6a45-4f02-bc62-d22ec29594fb	bc118170-d068-4659-81c1-59876de7d78e	43.25977	76.93525	22fe87b7-996d-4e60-99cc-55b44b198393	ce55f6db-022c-4bbb-9965-502a0e284ccb	94ac38c8-c38f-4c82-a702-a4bf95f0c00a	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:18.271057	2024-11-02 16:58:18.273188	ce55f6db-022c-4bbb-9965-502a0e284ccb
fd86ddb2-dcd8-4145-b5ec-f40a4626711a	23a8763d-db91-4b01-b60f-13b6ff440212	a0cb44c9-793e-4c0b-9ae5-3198ac70fd2a	43.25977	76.93525	c16bd828-f0cc-4467-b9ba-4605230d433e	6c522307-e361-4df7-be1e-357f99487224	0c442ac3-2a96-4d61-82a2-01e4d5782cfc	43.25988	76.9346	0.05393565459915141	t	2024-11-02 16:58:18.507892	2024-11-02 16:58:18.511426	6c522307-e361-4df7-be1e-357f99487224
\.


--
-- Data for Name: user_ratings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_ratings (id, service_id, user_id, review_text_visible, review, score, created_at, updated_at, updated_by) FROM stdin;
c0d2065f-ad18-49c9-9a2e-3f4d7ea7c1b5	710a92d7-9bd1-45bc-8294-eb3ab0150c71	1a513cda-5d18-4d9d-b300-a237423723de	t	I liked the client! He is very kind	5	2024-11-02 16:58:13.747244	2024-11-02 16:58:13.747244	00000000-0000-0000-0000-000000000000
c3e2a4a6-7169-4965-ae22-e56b7a12d5ce	2c1834a2-d2ff-4e59-adca-08fb52fb09cf	0a63a3e8-3846-461b-bd05-a321ee269598	t	I liked the client! He is very kind	5	2024-11-02 16:58:13.928674	2024-11-02 16:58:13.928674	00000000-0000-0000-0000-000000000000
c1a1b2bb-1e10-4c47-b508-75a7247176e8	ada0441f-f25a-4201-ad31-332055054882	fec33571-8503-4a88-9d20-006970d0d1eb	t	I liked the client! He is very kind	5	2024-11-02 16:58:14.183801	2024-11-02 16:58:14.183801	00000000-0000-0000-0000-000000000000
412590b1-6dfb-4fea-a437-141eba71c3d8	6339e711-1508-4316-b20a-c141e2f35913	b556d2ea-283f-45b3-99c5-ab98cb8cebe4	t	I liked the client! He is very kind	5	2024-11-02 16:58:14.349267	2024-11-02 16:58:14.349267	00000000-0000-0000-0000-000000000000
7cb59f20-d55f-4cc2-918f-4b0821e32b72	1d0d8210-2d9d-404b-b736-197788eae0b2	695293b4-b8c9-428d-b309-a183f37f8251	t	I liked the client! He is very kind	5	2024-10-31 12:58:14.596547	2024-11-02 16:58:14.590263	00000000-0000-0000-0000-000000000000
b77104fa-8104-4b7c-a6a0-27468477024e	1e097547-1994-465a-a547-23cf6b90b6b9	84498241-6b27-4386-9f9c-44e76899404a	t	I liked the client! He is very kind.\n\n UPD: I've changed my mind: it was worst experience I've had in my life	1	2024-11-01 13:58:14.757719	2024-11-02 16:58:14.764909	00000000-0000-0000-0000-000000000000
86a2193c-54a1-4c96-b79b-2e0a71528714	5e2ebc69-16bd-44d0-8bb2-4fc832afb8da	5ab0fbc7-dcf3-45c5-bab5-38ed2c27c457	t	I liked the client! He is very kind	5	2024-11-02 16:58:14.988045	2024-11-02 16:58:14.988045	00000000-0000-0000-0000-000000000000
983b95cf-abf5-4b4b-bdcc-0e7bb19d5463	06b1522d-8634-4ca5-8925-7bb4e25a290f	976064b4-b126-4c81-8016-11d89f7724d7	f	I liked the client! He is very kind	5	2024-11-02 16:58:15.147207	2024-11-02 16:58:15.158673	00000000-0000-0000-0000-000000000000
2d2801cb-c07b-46ec-b75e-ebc1a00a0bfc	5d6ee30a-6347-40d1-8a10-0ef7a9ae4376	6f05dce7-cf3c-411c-b589-a73f07f9099c	t	I liked the client! He is very kind	5	2024-11-02 16:58:17.113948	2024-11-02 16:58:17.113948	00000000-0000-0000-0000-000000000000
a245bf08-30eb-485a-ba1f-f9097f081448	ff2f02c9-8c49-49d7-bbc7-fab7d3a8f4ee	692f028b-0d44-4dfd-809e-c2048d268819	t	I liked the client! He is very kind	5	2024-11-02 16:58:17.349482	2024-11-02 16:58:17.349482	00000000-0000-0000-0000-000000000000
7656fdfb-18fe-4831-bf5b-ec246c8886f3	b17b867a-d8f9-4f74-981d-1959149a7d59	e4c62515-a21d-47ec-97d5-85f14cce5aaf	t	I liked the client! He is very kind	5	2024-11-02 16:58:17.581741	2024-11-02 16:58:17.581741	00000000-0000-0000-0000-000000000000
4262581e-99c1-42b7-892c-6c21415d2617	d30f3cc0-53ae-434f-8960-b6771c734f97	e739d6a7-4eea-46ed-9efc-a9f7d5e3e6a3	t	I liked the client! He is very kind	5	2024-11-02 16:58:17.815186	2024-11-02 16:58:17.815186	00000000-0000-0000-0000-000000000000
fa268b21-e77e-476c-a943-3d19607bb622	b10592f1-dba1-48d8-85bf-1b6f05e557d9	e68716b2-3a90-454a-8720-5d71da7155cb	t	I liked the client! He is very kind	5	2024-11-02 16:58:18.045778	2024-11-02 16:58:18.045778	00000000-0000-0000-0000-000000000000
bc118170-d068-4659-81c1-59876de7d78e	fe4bc93e-7106-4b46-be57-1d2307a1ff95	b9cfb846-6a45-4f02-bc62-d22ec29594fb	t	I liked the client! He is very kind	5	2024-11-02 16:58:18.271057	2024-11-02 16:58:18.271057	00000000-0000-0000-0000-000000000000
a0cb44c9-793e-4c0b-9ae5-3198ac70fd2a	fd86ddb2-dcd8-4145-b5ec-f40a4626711a	23a8763d-db91-4b01-b60f-13b6ff440212	t	I liked the client! He is very kind	5	2024-11-02 16:58:18.507892	2024-11-02 16:58:18.507892	00000000-0000-0000-0000-000000000000
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
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, phone, telegram_user_id, password, active, verified, created_at, updated_at, avatar, has_profile, tier, role) FROM stdin;
5f912930-d554-479c-93a8-d25d69009915	testiculous-andrew	82950626838	296475627869449440	$2a$10$jzZPQByXxRT7MpeZJsKjOOSLDIRd9c7RUUCO6RL5B4lngDNMeyxv.	t	f	2024-11-02 16:58:08.266872	2024-11-02 16:58:08.266872		f	basic	user
cd3cec0e-8b33-4c18-9bc4-b95a6060ee7b	test-tEZ9LBVTEC	37272575369	5543356492230039226	$2a$10$Snxcpcar8dathHLAa02Fkeo2f/c/ETycDt/4Ehx8N3l8q687IkxR6	t	f	2024-11-02 16:58:08.346214	2024-11-02 16:58:08.346214		f	basic	user
668ec869-6325-48a8-8c29-d74774e56dd9	testiculous-andrew	47635086787	3268099200197459886	$2a$10$KyHbpnnQNG47vJgEWzDGoeQ3L4EbaDP.a4OHmeuAgJDoQ0Ps.8Afu	t	f	2024-11-02 16:58:08.718891	2024-11-02 16:58:08.718891		f	basic	user
ffffffff-ffff-ffff-ffff-ffffffffffff	He Who Remains	77778889900	6794234746	h5sh3d	t	t	2024-11-02 16:58:09.072186	2024-11-02 16:58:09.072186	https://akm-img-a-in.tosshub.com/indiatoday/images/story/202311/tom-hiddleston-in-a-still-from-loki-2-27480244-16x9_0.jpg	f	guru	owner
4f74fe69-5d1e-4881-8438-f5138a7ab68e	test-fMwUU4htUi	28059128227	549572543041335270	$2a$10$S3Y652Cucuy1ghDGNZRIcOi.Ty2.izWzuNzcJaRkYtf5LfETCaxXy	t	f	2024-11-02 16:58:09.67714	2024-11-02 16:58:09.67714		f	basic	user
02086cae-d58e-4671-b0ba-de3c8a8f8034	test-I9XC3u7QiN	51188654433	4991860801934983699	$2a$10$hscAlNn.rXw5yfbf8dCL5eQyiRaPAcZ4VJwAtZ.LkBpLZJVmG7Mom	t	f	2024-11-02 16:58:09.763619	2024-11-02 16:58:09.763619		f	basic	user
9d648166-628b-4244-b052-75aaf55c1b68	test-T2gIwIsbc3	82269868549	2630463776919204303	$2a$10$aJkSMh5S0DpTY5wVnXe8A.VXoCQMYp9tDYkdUcVzi0izo78LPg2nS	t	f	2024-11-02 16:58:09.846154	2024-11-02 16:58:09.853626		f	guru	admin
81bc2f70-8bf5-4bdd-bcee-2713340a543d	test-jHMnLhSxn1	45542971645	3961731483474227691	$2a$10$4vx.N.PkVByt49gOX3azDepCBhk67WVutye81TEv5XCA1LsVjUPLq	t	f	2024-11-02 16:58:09.923743	2024-11-02 16:58:09.930965		f	guru	moderator
073c8d62-2d6a-482f-9e2b-30bb89f5c363	test-MUelgHk4CL	83625112793	4755478215042334180	$2a$10$DcISrdrnUvC.yH2.BbUfGey2A9tX6QPAzkUrTHAOLGmwvSowidjd2	t	f	2024-11-02 16:58:09.999247	2024-11-02 16:58:09.999247		f	basic	user
6b15f1d8-1889-4a51-b78f-64f9d40fef5b	test-pT23I7STN6	95290516300	2495890938896210003	$2a$10$BlrXVCO3oCcBSI4DMCkWIu265.cflwLf81e3aZ5gZcAwrOfQ3bzIS	t	f	2024-11-02 16:58:10.086443	2024-11-02 16:58:10.086443		f	basic	user
17ddaab0-8bbb-4c89-9be0-d1dbe00fb29b	test-hGpd9Iqukz	96807606323	7950314791201256837	$2a$10$TlBd76EGLkThj7q/2obNAeQcq30BIaapnqmFOgSGSIXL9vman0VkC	t	f	2024-11-02 16:58:10.157363	2024-11-02 16:58:10.165304		f	guru	admin
cf48d8de-1e4f-45ae-ac0f-b893ec450e88	test-Bu76tAQBR9	72249671334	2561441370259870626	$2a$10$RFDCfIN8d/gEKp2FdOF.ue2h7ro0up3Zhsk.i5XI96ZfDWc/dbTvK	t	f	2024-11-02 16:58:10.243864	2024-11-02 16:58:10.243864		f	basic	user
67c55b7a-b156-434c-9751-c86f3dce6d9c	test-BHf323wT0I	12361587709	8686162790211385709	$2a$10$bGhdk8InOZz4s9N1JjvyAu0xFYEYBv7yf/8upy5vJqaT7XRfrN8GK	t	f	2024-11-02 16:58:10.315957	2024-11-02 16:58:10.324426		f	guru	moderator
8e604bb7-5eea-4e61-b2b6-c7b14f35e8c9	test-jFEAyi7RDw	88230238650	1309835388175085084	$2a$10$/QR8oDcUL.N2EuEp067wxOiazkKywM0gv1KlYSm9HaLKcHxEZFWdq	t	f	2024-11-02 16:58:10.402178	2024-11-02 16:58:10.402178		f	basic	user
7e1db39b-bb1b-49ff-aad6-3260c79506c6	test-qAbo53T5qG	35386160488	2430240600205522964	$2a$10$wNU9II59nlTQxkr80DT8C.s9QjZ0xQ345Adx/IBDyE7fht9Mj89re	t	f	2024-11-02 16:58:10.479378	2024-11-02 16:58:10.479378		f	basic	user
4fe446a9-3f75-4774-8128-38980369b53a	test-6X2VPtB9TV	57196676566	1324739068629190612	$2a$10$h9ZfOlC4x56i/dzrO91eHeRJDGsFOk73tGsAXvl1438Pqenhg7B2a	t	f	2024-11-02 16:58:10.56134	2024-11-02 16:58:10.56134		f	basic	user
bf79fc7c-17aa-47f0-9877-05b55a2934ef	test-4kLAbF4TTI	39440160950	5891825838282260093	$2a$10$rAF6/tiwqs4FesGDgk1W4OyW9BbPBH8mndpZlKrCklWNQWCCwpQAO	t	f	2024-11-02 16:58:10.633685	2024-11-02 16:58:10.63907		f	expert	user
6d6ed2f9-e2f7-4146-a1dc-f0761f341474	test-PR5lG2MS7Y	45676196342	1764220216132743882	$2a$10$/vUSmRfp2k/iAUvTZ7d5MuOy8DY3eRFUcX77tyNV/Wtc2QtzZqacm	t	f	2024-11-02 16:58:10.71551	2024-11-02 16:58:10.71551		f	basic	user
35b017d9-6c81-400f-b537-0e61e599285d	test-vO1K3NfmQA	17935663309	1734557255359496962	$2a$10$SyXiADyxRiXw3DRPSrN85OK7Gq8EjqEpQ42hvy2/BRUErIhAt1M3.	t	f	2024-11-02 16:58:10.786145	2024-11-02 16:58:10.792312		f	guru	user
b95b0b61-129c-4aae-bb07-ac46a419678a	test-HWypf29A9J	93525865011	8319013361493695786	$2a$10$SODMUVXBkUNM3Tb69CUDz.cDTNUofIMfngCXZC8uFNjshFlonVpH2	t	f	2024-11-02 16:58:10.867313	2024-11-02 16:58:10.867313		f	basic	user
15811a5b-e2e5-46a5-9874-25b2e368b0c4	test-jSxPSdmVWj	91204503164	6372690364606174387	$2a$10$yBS7zqpCAwuerDtxePKDDOr3m16XlPyblMkz7OoQlGbRRIz0dy..a	t	f	2024-11-02 16:58:10.938773	2024-11-02 16:58:10.938773		f	basic	user
c5d0b456-eb56-47a5-8375-eeef2dd11a00	test-rhrFLAQnJc	88982154614	976679245295511528	$2a$10$4NXvQn9uDyVcmU46bGYZhuBtJq1uHIv9gPMUfWHat3uqqjp.e2r3e	t	f	2024-11-02 16:58:11.016783	2024-11-02 16:58:11.016783		f	basic	user
470ca159-2944-483f-8c54-df9515fe2121	test-lyBaxBU7nQ	56027422565	228882040835727438	$2a$10$mcLoAAHT0mPT.UnMaQb8Y.lAFKEnOo3AZw0bRaO.T4VohXFzYRQ8G	t	f	2024-11-02 16:58:11.088283	2024-11-02 16:58:11.093608		f	basic	user
5a0d02c8-9f87-4f23-a251-70817ac37ccf	test-4uOdDGde8p	83085540622	3066458819418411610	$2a$10$ixLFRKJJbzfFKLzEpHdWsewjyIfONjV3fyRgExeXhWTMnN8UT/IGa	t	f	2024-11-02 16:58:11.170638	2024-11-02 16:58:11.170638		f	basic	user
9af53b1b-9ef1-4164-9519-7b054b36b15d	test-vHv8FwQPCU	75194117612	9211610782943867729	$2a$10$3RpeVZQ9lHsTkR3NyjeEuO2tfLdbAmFDy5Eitxz1FQc1kOoR3ttMa	t	f	2024-11-02 16:58:11.241468	2024-11-02 16:58:11.247607		f	expert	user
d6089edd-e7c9-44a2-9486-337ecd67c814	test-7B3CUAXaJB	39136249316	9090558453857325256	$2a$10$Q..i3D11tZxfIukCsmuywuwwF6hWizwuUiISxUp9z4u8gSFtH90VG	t	f	2024-11-02 16:58:11.322742	2024-11-02 16:58:11.322742		f	basic	user
fcf73090-6e90-4a97-aa0b-80c04d411194	test-PKOsBkNQKn	37372662783	4704772057518383474	$2a$10$6tFKhBUK0zwUFTf1ELWG.epm8fL6mhT/2jiXe3piX7p7MFgmghape	t	f	2024-11-02 16:58:11.393011	2024-11-02 16:58:11.398548		f	guru	user
1b3124df-a831-48ee-ac33-598ab00bcede	test-Bq3uqVe1Pk	31684270522	7492138030043105051	$2a$10$fjfe/dCyllJ3p6SRCRU5rugNoPsR4crLkHRonEjvmOYkJNUgO.09a	t	f	2024-11-02 16:58:11.475083	2024-11-02 16:58:11.475083		f	basic	user
6c5541f9-5168-42eb-97c6-6c6db0abe214	test-zUwo3E5yGl	86919198065	4515043723366730910	$2a$10$qdp/VHytouZWLfdv.afom.Gm3WC6R5.DE4dhm8OAlySbHCClALkWm	t	f	2024-11-02 16:58:11.546327	2024-11-02 16:58:11.555806		f	guru	moderator
e6196cf4-a677-4446-9435-3dbdd7859a27	test-9lYV2Hc5A0	36737439659	4793011268739557878	$2a$10$BLUFz9/5.eVNj1yp4AaPwuRxS/64KYT3Z2wZIuuI5pEhmSRjq7vnq	t	f	2024-11-02 16:58:11.64027	2024-11-02 16:58:11.64027		f	basic	user
7ac79be6-463d-4bc1-b5b8-b90c7559596b	test-syT3ULfMLD	95025880546	1840827874521780119	$2a$10$ekClkzkJXQmg5fDZ2G0DR.Lt77y5kP1ieePqo14JV.P7.or0lmBx6	t	f	2024-11-02 16:58:11.718834	2024-11-02 16:58:11.718834		f	basic	user
aed0f45d-617b-40f8-afcb-869c3ba79caa	test-0lGyiQUUQF	29815052156	3678283130657837945	$2a$10$VzqB80.qSpmXwl9JUf9/k.8QbMzopPlIOT4G/h0Xl4SnRw0n9d1E6	t	f	2024-11-02 16:58:11.788684	2024-11-02 16:58:11.788684		f	basic	user
34820fab-d019-412c-bb28-52a7039dc461	test-Xab22xTnd2	80344108483	5984471076467812985	$2a$10$z4U9k9KEaYThY938ZjaRruXMvNdPTEroGs6LViv.r1xocGwiaQ2Sq	t	f	2024-11-02 16:58:11.869932	2024-11-02 16:58:11.869932		f	basic	user
e6f112ed-a0d2-434f-b50d-4b7ca1d226ab	test-xdgCih3T77	21587848007	8037592676603881900	$2a$10$zWlm9boK63d5SDwiBGaZ6.JvcSecX/uoKr692aH593Qqv0cuGJ6dO	t	f	2024-11-02 16:58:11.952751	2024-11-02 16:58:11.952751		f	basic	user
e4bd738a-a772-4f08-8aff-f959ca3a2b42	test-z4Sp0MYB2U	67099233300	1399936944047095294	$2a$10$CIJ5KltPkT0z462PY89b0.5KfTOxzkplAtbxuv6w0dza.yNK3i2LC	t	f	2024-11-02 16:58:12.030748	2024-11-02 16:58:12.030748		f	basic	user
43a9ff3d-ccd1-4100-9b00-911f069ca195	test-243WWN5pLz	41860885592	4958536106098816034	$2a$10$DQCBODoDmUHgUu50tQoeBurKVgKDwuAmIS3nEQPLQTG3zHaUehWbq	t	f	2024-11-02 16:58:12.103039	2024-11-02 16:58:12.112937		f	guru	moderator
e86ca8f3-f8da-4199-b4d7-d8e0f7f4a626	test-kPjZ0Z2tDk	89551554056	9029415168690971630	$2a$10$GSjTkbVKLfLyM12ySG46TuYQEX.e8Fy4r//Ileg8JkDOJeP1RhSAK	t	f	2024-11-02 16:58:13.643552	2024-11-02 16:58:13.643552		f	basic	user
1a513cda-5d18-4d9d-b300-a237423723de	test-ZSuCKlJCXn	16129149931	9104781197189958478	$2a$10$fzXYoAgi.NL1UlF8uvUaNeEUUvlt76dTxsat/WsMX0cxApnKcuYkO	t	f	2024-11-02 16:58:13.725227	2024-11-02 16:58:13.725227		f	basic	user
054c1165-070e-4e62-ba0f-3d46d266d82a	test-2zHBvpgDql	77431713381	7465011385911788335	$2a$10$h4i/3Ny7qPO.szb.HAmjJutj3rYc3orwBR9mbZnIa/lHfFwahIr9i	t	f	2024-11-02 16:58:13.850227	2024-11-02 16:58:13.850227		f	basic	user
0a63a3e8-3846-461b-bd05-a321ee269598	test-7gW9XvdzDZ	96601117519	1852895470074220922	$2a$10$8bLufZAqbvitUO3rR4GDE.WBBsvmsaakXNVlVKtVAcVZEaNrKdQzi	t	f	2024-11-02 16:58:13.918202	2024-11-02 16:58:13.918202		f	basic	user
0d9367bd-e720-4618-aa94-5a33db4b9ec6	test-AOMr9MSCS4	25691746447	4379037487758357582	$2a$10$.bJl7oZCzsRPvqICQefsFOe7E.AAibKY4jjkuBmDoxsbuopgTTxRW	t	f	2024-11-02 16:58:14.027598	2024-11-02 16:58:14.030664		f	guru	user
3bc5d319-83ad-4cbe-a882-b33e1dca16c7	test-LF7yb11AoO	48179489460	2629842645832969780	$2a$10$GKQEGJAZ9wj92Iz4ZYIHzezHBDEadU2.W64a.f4kSGZ3Y11.liiZm	t	f	2024-11-02 16:58:14.104476	2024-11-02 16:58:14.104476		f	basic	user
fec33571-8503-4a88-9d20-006970d0d1eb	test-vyLWCoWZXg	11649281275	6072567297791670325	$2a$10$qd.q9GKcaWL1PFDeVV7weu0GtXwOQ5MjKhmbjRLReFymOsffR9w4a	t	f	2024-11-02 16:58:14.172288	2024-11-02 16:58:14.172288		f	basic	user
16cc1398-460c-42db-a025-3dcd170105cb	test-pUIHhqc0Gb	19177455280	2308092610678819648	$2a$10$wFo5zxtDF.5S.yHQLJH6DuiT6g3QyLSMF0jZEIL.laxw4bHlbOEq6	t	f	2024-11-02 16:58:14.25733	2024-11-02 16:58:14.266095		f	expert	user
b556d2ea-283f-45b3-99c5-ab98cb8cebe4	test-UlHIE0kAvH	11255406055	9180247237137778504	$2a$10$.9dLPqN28dPGftQmQkq5aunRGd7mOvwcv2iRzusV687hCs1ZHDGN6	t	f	2024-11-02 16:58:14.335247	2024-11-02 16:58:14.335247		f	basic	user
b0d28cc7-84a5-4558-86ba-2f0256342da3	test-qTVrL7b7TP	39267900472	453664050122217283	$2a$10$/8G9H.zSbOJvi9xLglE5xefwcZFCbiBHVACOSYX47gKIqTWoEcZg.	t	f	2024-11-02 16:58:14.4285	2024-11-02 16:58:14.434598		f	guru	user
5fb13376-fb1c-4796-a246-58caa72311f6	test-orn87zX3bD	30118780071	7045642569277511083	$2a$10$ZIcTgHJ.n1EtpjYCQlSMwe6PnLRYuddiXI/yBqdsmIimH78VRxWJO	t	f	2024-11-02 16:58:14.507061	2024-11-02 16:58:14.507061		f	basic	user
695293b4-b8c9-428d-b309-a183f37f8251	test-ZIxq1lEOeq	91608588836	3847628545834672866	$2a$10$ZMRpctFobQhz1T0DCc27B.TzGrQyfMlsqAPk4sO/fMnd.PE8.s/f2	t	f	2024-11-02 16:58:14.57845	2024-11-02 16:58:14.57845		f	basic	user
045eb6ac-555e-4cb0-8319-a29f74801cbb	test-h7cK1eUk3Q	86426562826	3324795805555601510	$2a$10$FmXXqmH4XlNOKnWenmTpBO5od4mSnNIjfWPMFCk8m6pla37f0VYra	t	f	2024-11-02 16:58:14.666613	2024-11-02 16:58:14.666613		f	basic	user
84498241-6b27-4386-9f9c-44e76899404a	test-v3sZ3cWIUi	82145212982	8045968778720668233	$2a$10$Bke04//5N6GYjh9W6qAoLeGKBYqICOPtT9Tg8nubMwaeJ1He6dkcq	t	f	2024-11-02 16:58:14.738023	2024-11-02 16:58:14.738023		f	basic	user
7ed341b5-a4d4-4e0f-a57a-e2583bfe04b5	test-mIgm6Xp5CF	43501047729	3114337460985257684	$2a$10$gIJ45eX6.SOCiSDC/tO8dOgdzFUD1ffIww5SCW4zimUc6Td6.48WK	t	f	2024-11-02 16:58:14.83211	2024-11-02 16:58:14.835254		f	guru	user
a80a4c11-8293-49dc-9f8b-6609ad8c6ecf	test-GRToq5Lse3	46055733114	6407574395246671539	$2a$10$R00ogmbG6XWVFOl/nTk63OLv75FM/AJbfBCf/sxNOLdpuQKzli3wS	t	f	2024-11-02 16:58:14.908393	2024-11-02 16:58:14.908393		f	basic	user
5ab0fbc7-dcf3-45c5-bab5-38ed2c27c457	test-n8REyVMDoK	51597145213	6072891853484317634	$2a$10$uz1YToFtHNVY5/MpBZV73uJnaxEy7/vt/4qLYnT/sC21jG4ea1n6.	t	f	2024-11-02 16:58:14.976049	2024-11-02 16:58:14.976049		f	basic	user
8d01e8ef-fb0f-417f-94c3-efbf77d2da7f	test-pat9jindA9	32634605062	7499368505563122119	$2a$10$gbAD3SfFhQW16eUrRYBR/unT/C.VxPDgHBh7sfhgY9NNaifEqdEJi	t	f	2024-11-02 16:58:15.062341	2024-11-02 16:58:15.062341		f	basic	user
976064b4-b126-4c81-8016-11d89f7724d7	test-fUnz1XnQEi	42075085843	526316887480867997	$2a$10$oZFWCcazV2B.pMzOi6JjUu2pPn57XyJ4hqn2/VwNsQfNYQOPTLgVG	t	f	2024-11-02 16:58:15.133666	2024-11-02 16:58:15.139771		f	expert	user
bf4bd81c-1a45-4407-b1c0-af6bec84aeb3	test-g5HaGWtSeu	45784352249	8168228477176340972	$2a$10$l1qTFL0B4Bv2VKfbXqmZGOPh2iER2JZcMm4b5FuYcgJ3U4HqLUQTC	t	f	2024-11-02 16:58:15.226972	2024-11-02 16:58:15.23261		f	guru	user
2a895964-f41e-4e2d-a6c4-04c4097416b0	test-snTe1h9Oon	30379621435	7430406475201806198	$2a$10$Dpwh7xV6qqZu4FNQjMlEFudCVCEReIA.2CnTgciw.KbM9o8vhPJBS	t	f	2024-11-02 16:58:16.544876	2024-11-02 16:58:16.544876		f	basic	user
3b5b59ba-28e7-4eee-a4f4-8bf225217219	test-834w3WFQmn	48439030039	7258478014766872247	$2a$10$3bPIEj.e.p3Cuwd5O0zAXewKfevXJLpilJ0nxBbTQYKEDhDdLKvk2	t	f	2024-11-02 16:58:16.627266	2024-11-02 16:58:16.627266		f	basic	user
5c87b1e6-906a-4bae-a022-20e6d630788d	test-8vykZlLOif	27797775700	3833456405740844267	$2a$10$DffO4MzAcQ98Xl26w6YzE.ws.pvs9QF2dfH3.VBFBEVbb1WlI44uK	t	f	2024-11-02 16:58:16.713979	2024-11-02 16:58:16.713979		f	basic	user
d9690bd8-8d45-4ee9-b594-117ecf396fb1	test-h1wpOf739h	67460322577	5934049550185745842	$2a$10$SDVE/Wne3eoAzZSQBsTo5O4KIvInhpQkIk/evTrwoFFdMTNI25VPe	t	f	2024-11-02 16:58:16.781954	2024-11-02 16:58:16.781954		f	basic	user
b9e32c2d-7b48-401c-bbc3-ff5766fa55ac	test-HKDzJfEjbv	37248311748	1181302225125556188	$2a$10$Xq3va1erLoKXt0lvK2ydI.WFyMLvLlMpjrJJhVlTMD.V9Zyg6EPs.	t	f	2024-11-02 16:58:16.879818	2024-11-02 16:58:16.879818		f	basic	user
5ea5d3f1-b318-4875-a47b-ee57b9e6bb64	test-ysn5lQ9xYk	23862193308	733369046724484472	$2a$10$9tsWuY1JpvBIn9uX5k6UH.ll5yAJY1UG6Z8K545eT//.jklWgQRz6	t	f	2024-11-02 16:58:16.948012	2024-11-02 16:58:16.948012		f	basic	user
b4568f89-0c9e-466c-a2ff-d8837a3506b5	test-A9V3nLEGdW	91666169729	8310784421278533391	$2a$10$.6Qj4btUFUW1sP6/3snmUuCLIKbRNJO0oZABzQCeYMm/MfDZ7NRxe	t	f	2024-11-02 16:58:17.029371	2024-11-02 16:58:17.029371		f	basic	user
6f05dce7-cf3c-411c-b589-a73f07f9099c	test-zXwbJFNaoe	53969385538	3742301577268012403	$2a$10$GbmriOkm7SVWSU6/1GfJCOwGbaO8K6F5K9i0yEhSAu76SbT3uHqXi	t	f	2024-11-02 16:58:17.102088	2024-11-02 16:58:17.102088		f	basic	user
f9181aef-a323-4bdf-8c42-115f9ae25a24	test-LVI0cYR5m5	26195753140	2865555589688551496	$2a$10$5bgcI9UuxeWxUG9blq6DW.8sDQ5/wqz/3wnt5X0ei3SgB0YZFx2ku	t	f	2024-11-02 16:58:17.190415	2024-11-02 16:58:17.190415		f	basic	user
4850299e-3325-4862-9b1b-fd1941b06714	test-HbeGtl9Otb	29089874041	2131108969466235484	$2a$10$YJebh5X5N7dy8PbTQ22Y5.Gr6jdWKV/7UZv79HJ4p8ke.e2dOO0w.	t	f	2024-11-02 16:58:17.268724	2024-11-02 16:58:17.268724		f	basic	user
692f028b-0d44-4dfd-809e-c2048d268819	test-XTcMFOSLwx	94626107355	1592985655762481838	$2a$10$tKp/hhPjT9dVPaPsLZ9.0.RurDW34U13MDJiUnP.0Gd12CYbb7sXC	t	f	2024-11-02 16:58:17.337487	2024-11-02 16:58:17.337487		f	basic	user
52cc9eb5-5c6e-4e33-9d8c-8c1aaf272280	test-Pu4ynIisrK	10314593955	356098767041283194	$2a$10$I6.9Xkm3OfJXrlPUUDHPVO8k8kcd5QTFXzXP3sDFdBpG3x3pj.hg.	t	f	2024-11-02 16:58:17.42071	2024-11-02 16:58:17.426767		f	expert	user
a85f6e36-fa3b-403a-83f7-27e143b77cf9	test-lDmliNlK07	99052651798	6918679337273007103	$2a$10$zsnNy4D3YRvarnTiMj8uUuBRybol3iXg1k0F/SL.RNkUSOy1wBvua	t	f	2024-11-02 16:58:17.499269	2024-11-02 16:58:17.499269		f	basic	user
e4c62515-a21d-47ec-97d5-85f14cce5aaf	test-Iqg9urBtC4	39972821758	6833681096857826484	$2a$10$B7J/ezxAXvoxoXBr9AKHsO62.5TsT9VDtcxiEQJqakz3e9r7DWMkq	t	f	2024-11-02 16:58:17.570899	2024-11-02 16:58:17.570899		f	basic	user
be96bb8d-8333-4d97-a9a2-4027da0c6472	test-6NSwof6a9a	54918088761	8620672408587293336	$2a$10$HORy64lT6/OHWYUPeOKAieA3dT7v06dMRMGokgOtd8MJzRq2.YM5O	t	f	2024-11-02 16:58:17.653387	2024-11-02 16:58:17.659202		f	guru	user
f0da98d6-834f-45c9-b70b-e0c54f057aad	test-B9IeDF1bYF	18251300973	7470144477856341606	$2a$10$a6PbhGn/Dg9vuBelJDXem.yWB4C2eO.X7EFuQwWqK6oi2L2kFma7K	t	f	2024-11-02 16:58:17.730764	2024-11-02 16:58:17.730764		f	basic	user
e739d6a7-4eea-46ed-9efc-a9f7d5e3e6a3	test-sZezp5n2PZ	90917851887	8324414608361415924	$2a$10$3jyqi2JdSOCLl/5oW7vkf.C7nHPCUAARr/ZYW/2YhSs6RUgTpjOsu	t	f	2024-11-02 16:58:17.802159	2024-11-02 16:58:17.802159		f	basic	user
d3b688be-2872-460d-9840-c5d4ec39aa5f	test-KHuTThZldz	47489303622	8513534765148166537	$2a$10$mVXfPsQkN4JliabPAceDcuAmkWnkiYBBe.zUnpo.nbLJYWFbuKF32	t	f	2024-11-02 16:58:17.886863	2024-11-02 16:58:17.886863		f	basic	user
7702b59e-ea14-4a2e-b194-691e89e50ba4	test-NeKU4VDVlt	72005428915	9159176274398394841	$2a$10$x1EPkPrvIoBgxgMv79a6i.//KmGrFmc/6WpKRKWdldGtf1rMM.kxO	t	f	2024-11-02 16:58:17.960675	2024-11-02 16:58:17.960675		f	basic	user
e68716b2-3a90-454a-8720-5d71da7155cb	test-qz3hVVUQsR	67836239175	4328884331170388773	$2a$10$9TDaE15Y42tVLVmaeRKAGeZy8wqOlAXqxYZJP9IBPazkJCOlnTxfK	t	f	2024-11-02 16:58:18.036523	2024-11-02 16:58:18.036523		f	basic	user
fed73a23-8ce3-42aa-8a2b-904eda5ee2c8	test-wUpJFJgI3E	46904773037	1015155187956016639	$2a$10$FM7D/RbxwjBpbjKgznIld.KxiXxDzH6BPYRfsNh7uumOBdwddRMAa	t	f	2024-11-02 16:58:18.118062	2024-11-02 16:58:18.121218		f	expert	user
ce55f6db-022c-4bbb-9965-502a0e284ccb	test-kdi23tbRhQ	23572859877	1819771011206361401	$2a$10$dSJ.6bRPy2GLGyOY58i8C.bfIwuDlAyiFcIIzHPx5OGt52/2tK/Ji	t	f	2024-11-02 16:58:18.189995	2024-11-02 16:58:18.189995		f	basic	user
b9cfb846-6a45-4f02-bc62-d22ec29594fb	test-fqhpfXBTgR	98413420896	8924206710851532189	$2a$10$0qQIlMhZHdFkKj9Inxb3QuRSW2CIbNLXjIdZlFiiqZeXt8u/GBns.	t	f	2024-11-02 16:58:18.258563	2024-11-02 16:58:18.258563		f	basic	user
d4008e37-b4e1-4ff6-916a-fbd87a38d583	test-v9xMJg6Hpm	12351406183	433091007683516025	$2a$10$kCBFt86.s0tJ8HrgcqU67OOhs0OQLTJ81gTY1xaopclN1SdCDzchm	t	f	2024-11-02 16:58:18.342129	2024-11-02 16:58:18.348702		f	guru	user
6c522307-e361-4df7-be1e-357f99487224	test-zA2rpDTJwK	24978314984	13458243684371905	$2a$10$FTjd7mB0s26tz9nR1o/lB.8i5BQKNfNUdJsTYSGbPlAHd37f/LzAi	t	f	2024-11-02 16:58:18.423403	2024-11-02 16:58:18.423403		f	basic	user
23a8763d-db91-4b01-b60f-13b6ff440212	test-0WPaat9Zaz	20961903889	8539181597470527704	$2a$10$mQWI3pmRj5Z3Kjoo23/4wOO1Wpivd6b48H0aFTfQJuqhG3n/S16e.	t	f	2024-11-02 16:58:18.496222	2024-11-02 16:58:18.496222		f	basic	user
e5999cf4-d6f4-43de-8814-b371d655704f	test-msUcwvggLL	42299233090	7233910703692840241	$2a$10$J08Os0N5B4LTYTB4Nen9kevnLgWF8LRYVIso6ZXR4hOQxI.F3qZFm	t	f	2024-11-02 16:58:19.253873	2024-11-02 16:58:19.253873		f	basic	user
eec43268-ae57-444c-8aea-f79e069f60b3	test-gF9pmIOK2A	72828518321	2988780796347586874	$2a$10$X7FPubn83FkwtbZtf00ePOcV/BlwrTDygHUdSt/t/VulV4VGJYiTS	t	f	2024-11-02 16:58:18.581057	2024-11-02 16:58:18.591497		f	guru	moderator
34a16b98-e5cd-4caa-b757-d0fea1078682	test-OFVC4NV1ne	31778239482	7266147961554044270	$2a$10$pVqSdsnp1Okw1geFBrlTjep8pk3lW6a3Yl5KQkm090A6Qeag8G2K2	t	f	2024-11-02 16:58:19.167062	2024-11-02 16:58:19.167062		f	basic	user
e30bfbd6-9fb0-40fe-a0e6-ea9cbbaf1584	test-LZSlJdkdDf	39893552772	7670574048396179400	$2a$10$1jdh1pphKoiBJWg76sA9qeybcvLoSLAv1xzDuo3Lc0slGhTzmpWyC	t	f	2024-11-02 16:58:19.324185	2024-11-02 16:58:19.330303		f	guru	moderator
c26ac1b2-1399-490c-a1e6-ad3cd45d54a7	test-ERYmTy3ubA	69992960398	6180013406238010809	$2a$10$l.v2.DmlWNlYPXJp7HHDHujGT0jd/mFmREKOBeSup1PCLMZ6Pi9Fa	t	f	2024-11-02 16:58:19.408892	2024-11-02 16:58:19.418606		f	guru	admin
2fa91364-57e1-4505-b8dc-7a04ba2d98df	test-uUjY6Kc0C9	23659474672	8426049015518378348	$2a$10$vV8.m2MYg5a/CClvvJymtOOmHihCRw.4d19yHH8Fncz32gwbxsCrW	t	f	2024-11-02 16:58:19.487358	2024-11-02 16:58:19.487358		f	basic	user
87e7cfb8-ec2c-421c-80c4-4723f699861a	test-tdWJf3hpPQ	84270097008	7162624154517110256	$2a$10$5Idz2cj8aDBdAXxZYD6KJeWKJCHLmD/zPdfAw/AVUiLJv8Z/e39ty	t	f	2024-11-02 16:58:19.557436	2024-11-02 16:58:19.557436		f	basic	user
f49884fd-56f1-4bd5-82e3-efeec89e4e8a	test-SaZezwnlma	74443794958	341767645774133942	$2a$10$rogGlDxk9qc2pu/uQBVLMeitHICfnGRs4XKQjwC6ynUkp2Hjm.44O	t	f	2024-11-02 16:58:19.629021	2024-11-02 16:58:19.629021		f	basic	user
ce28ef69-0f76-4bac-afd9-cd2743edd001	test-OMMiPbdVLw	88495944067	4464779879697313292	$2a$10$WVH7M60/oysAApRYXRtm4.S0yhTqk.jiJr9veIh3APJYRGEdbKYRK	t	f	2024-11-02 16:58:19.708138	2024-11-02 16:58:19.708138		f	basic	user
91d895da-7b80-4c24-89d1-2e0c70bb3a38	test-D6Bnug3Hvm	39916578830	7452979196767761163	$2a$10$GM1xk/y0RvpEYx0AxLTWOeVqZQbIi2kilhcPKqYypZ4njtslSFNVW	t	f	2024-11-02 16:58:19.779349	2024-11-02 16:58:19.779349		f	basic	user
3e5206f2-863b-444a-815a-bf7d561c7562	test-dzpQHIlOtq	41785974326	8777414564663123659	$2a$10$YQBvPrfv7tce7EP1jeSRzOWA3z2A/NljEDIOAND6AWr92aTkz88re	t	f	2024-11-02 16:58:19.85309	2024-11-02 16:58:19.85309		f	basic	user
5076f6ad-2f5d-43b6-8162-b372f1f451d2	test-X2S0GJZD2P	67987718778	914589196518929754	$2a$10$abIXaLYnsS1cKXN5TiDNveEZlr.bLdnAgtiPPfuHOcqOclhz78zkK	t	f	2024-11-02 16:58:19.924355	2024-11-02 16:58:19.924355		f	basic	user
403554de-8b87-43b0-b862-688c740906f2	test-SR8Xqak1vf	78830195310	3985978504592814826	$2a$10$Sc.7YusD8ZFaBk4cd9KnweaX4iYZQIkkUpPnlIpoKoPaZi4MN0V26	t	f	2024-11-02 16:58:20.000761	2024-11-02 16:58:20.000761		f	basic	user
d9d5ebbe-a793-4906-a32e-0daee1eb0d54	test-fdQRoFbl2j	94591937290	1797009445557638352	$2a$10$OUwJrK7p1IamofRtN5IBH.y5al7LdmWxp03iOMVmJMa9lZ2zbZuV2	t	f	2024-11-02 16:58:20.07699	2024-11-02 16:58:20.07699		f	basic	user
a3165bb8-7f73-4d2d-af0b-3d90850c1d73	test-73sF8pMfLC	43167770440	8634882651554178467	$2a$10$MNWR/fBo3HQ0xYTjQBvyb.ISwQJ9Zrz7Xefe0LLSKsqzRNCZnP/kO	t	f	2024-11-02 16:58:20.155025	2024-11-02 16:58:20.155025		f	basic	user
820c781b-1ca0-46ef-a52e-c37d3ffb78dc	test-o53cyZi0kB	14444838413	1224199762943460542	$2a$10$DVk9a5MNBp7xWczXzL8vuehEqqpW.mRMHjIpfXO3p1MalykDFG9KC	t	f	2024-11-02 16:58:20.38052	2024-11-02 16:58:20.38052		f	basic	user
a5709992-22b2-4183-85ec-aa136c5c37cd	test-ce9ghgri3a	70626481967	323421611235261628	$2a$10$i7g8l4UfM4CQOXvhEA06reZpUcZbKfOCv05PyyAp49V8bzak93r.y	t	f	2024-11-02 16:58:20.309708	2024-11-02 16:58:20.386779		f	basic	moderator
7880faaa-07c7-4e82-b291-06cabe6d0292	test-7c5YNBYnvr	77794674501	9167751696744320671	$2a$10$ExWDmlmhJ7/2VGf53gv5gubC4zho.EAyLs2AfVBx03wf12vWjcuRq	t	f	2024-11-02 16:58:20.456658	2024-11-02 16:58:20.540221		f	guru	admin
d85cb34a-cfb6-4dd1-ac7a-a46039aafc14	test-he85urmX2I	53436407806	1120577424459067532	$2a$10$2h28RFXY3qT6Q0g7B60.xuGvqGum9l0wP/s8VI6l2i2jsVCOgtIsS	t	f	2024-11-02 16:58:20.61285	2024-11-02 16:58:20.693357		f	guru	moderator
b9c2a31e-cab8-4aba-8e5d-25fc1a8428dd	test-9iHEb9R4YD	67404712019	4425800146069525457	$2a$10$LSC2QZuy.t06PBsafetCc.L22OFo.MlPhFHvzG.NgvN9yxOTuphFa	t	f	2024-11-02 16:58:20.683753	2024-11-02 16:58:20.697528		f	guru	moderator
5b90548b-d630-43f3-8ada-fd209fe89cb4	test-GbTFkSRfNi	67793238688	8476062517892071543	$2a$10$LRfKKM.AocXJaT4eiwYyseemUAlrhGmYBb3QwHlZrROHX1B/P9t2W	t	f	2024-11-02 16:58:20.765709	2024-11-02 16:58:20.842881		f	basic	moderator
a9b0ba2f-b1ae-4013-b479-460aafc9b7b2	test-dJhunhG9ED	27478633436	1689109283824257503	$2a$10$4pkyy4iqmsLKEgBR.g2TmuWRQj/7aDCALBG79MeY52sip9dqbtiQe	t	f	2024-11-02 16:58:20.837594	2024-11-02 16:58:20.844438		f	basic	admin
fd643d47-a40c-47ac-bb43-9e06198416c3	test-fk4PapqLLp	95967641980	5389279063695747248	$2a$10$y.u0Ua/EpAckB0hFBbBM/OIzZyEWAF5YAA3T8Lzk7bR7hLEvQxXvW	t	f	2024-11-02 16:58:20.912933	2024-11-02 16:58:20.991395		f	guru	admin
02b31855-d9d6-4f95-bd73-bb3ecd8734e7	test-P9GcrQYnGm	77000000101	7285170109728589989	$2a$10$BEatCTvU/1mA.6oeOIhgbecvP75qsXOZ4WcMkFPHMoQrsupyNXm2W	t	f	2024-11-02 16:58:21.601069	2024-11-02 16:58:21.609924		f	basic	user
213430b6-5b06-403d-bcfc-8cb347fd027a	test-w3YpyAynBB	72208327561	3765770618779753274	$2a$10$AQQCE8TFHceCJbgTaq00leq8aqWqCojTaCboIxLciQ8LNsR/avAiC	t	f	2024-11-02 16:58:21.064637	2024-11-02 16:58:21.143305		f	guru	admin
14d4f5e2-4d6a-4a45-b182-d6c4167e579f	test-1W7Af2lnAv	43487645288	75443155628580139	$2a$10$0Dov99bMO8u8.XbkMys88e70Xqc6DT.E4vZ/OJDkSmkA.pL2p94G6	t	f	2024-11-02 16:58:21.136214	2024-11-02 16:58:21.14871		f	guru	admin
dd1b0e12-e46a-4159-9e90-a7bcdc0ceae7	test-zrXLMcb96A	16456419797	5615680185921305934	$2a$10$91qZDu0zIvPzG2BqObHkme6ka84.DQkQTtdAq74mBCEYqYKajVoBq	t	f	2024-11-02 16:58:21.298416	2024-11-02 16:58:21.305847		f	moderator	user
ea7bc66e-4e43-4355-a519-0c331825b0f1	test-rykQmGOisx	29862213899	8314095203996220871	$2a$10$3ts9toKAtPKDRL42JttZWuUXsvT5w6/z9YyDI6J0B0mYsM8clftD2	t	f	2024-11-02 16:58:21.374622	2024-11-02 16:58:21.383318		f	guru	admin
11062d8a-b095-45ef-a49f-114836649879	test-qZ5JRtVVQr	72786390142	6684325589273742992	$2a$10$CfGGT8..uTOxmgDULhdYOeSYPQKjIV6cSiDneEgr31zM005lSdchW	t	f	2024-11-02 16:58:21.451859	2024-11-02 16:58:21.451859		f	basic	user
4de718be-b434-46ee-b63e-e27aeecec1ea	test-XfUCQxctxv-new	62585684419	5291617654394025770	$2a$10$NQ5JjgPS/sFLnYRHlmvP9eXjbfViIhM3GvckFcDxMHK9IrZnTBL3m	t	f	2024-11-02 16:58:21.524456	2024-11-02 16:58:21.533851		f	basic	user
d5e99056-9176-48b5-b162-c1f4601fcdc1	test-kQ6NsnFYli	40913864825	6933035452931571840	$2a$10$mRCAJm5WDcgDuJ8oLpakYuwy.rUD.YzHn7KIj0u2ypKklQAYbezqS	t	f	2024-11-02 16:58:21.679014	2024-11-02 16:58:21.68803	https://jollycontrarian.com/images/6/6c/Rickroll.jpg	f	basic	user
ae8cc79d-6ef7-46c5-a1cf-61dd00c6079b	test-4qQFoM1dYD	54779110782	605704099271969518	$2a$10$qt4nRbt3.36Dk3NX855ra.rDN2bmDwuCY/S.C2DM0zpqDaI5h5td2	t	f	2024-11-02 16:58:21.755534	2024-11-02 16:58:21.755534		f	basic	user
15df947a-7ede-4e75-b13d-1c554655f2b9	test-BwFawg42nZ	14256217382	760855359032760151	$2a$10$XflukB9zLIUJf6c7AVQU5eCaPik8Dc5nWr2MZ6k/yCZOXECvr6RNO	t	f	2024-11-02 16:58:21.826993	2024-11-02 16:58:21.826993		f	basic	user
cf5efb74-53be-4862-a9e7-7b5db6a53409	test-KpWDKSkqni	10507612370	5614314961168686253	$2a$10$itIEi27uOC.zfYeOtoxRLutOPIv49QFbFStOhFQS2FdyQXTzw7EUC	t	f	2024-11-02 16:58:21.899325	2024-11-02 16:58:21.978607		f	guru	moderator
de2d13cf-bb9a-461e-b04b-986b918cb420	test-XppBxcEnJV	66075562018	8085502589710335492	$2a$10$8tnzKsSSJ7pDif5dedqq9.7nrn4qTcI7Q6aU2xFlpFY3nZAacEzjq	t	f	2024-11-02 16:58:21.970668	2024-11-02 16:58:21.982229	https://jollycontrarian.com/images/6/6c/Rickroll.jpg	f	basic	user
70c14fda-ede2-4be8-9d3b-0355497ae898	test-uj63AvNtED	19829082154	8862007681348368301	$2a$10$OGHONkgksrC397ujtKR31enLTfk5HffGTyv1EQfoXxmwXHPol.6A2	t	f	2024-11-02 16:58:22.049125	2024-11-02 16:58:22.128228		f	guru	admin
8276671c-f95f-42ce-a97e-1d07cdc18c6c	test-1gVUsSN6gR	15762986984	6217603464062930681	$2a$10$2qZThtz1zh8sX2YwHIki6eJV.TRC0monBZ./UipQ86eslMT5ccwde	t	f	2024-11-02 16:58:22.120146	2024-11-02 16:58:22.131618		f	basic	user
55c1bcf5-c26d-44cd-bba4-49e779670c0f	test-dULNtifTAm	48065678805	2934618602752113066	$2a$10$GpknZXfFUDuD93..UwQmo.LTq./f/7MR1OqbVhDAZdxO.PwzyeAmG	t	f	2024-11-02 16:58:22.200239	2024-11-02 16:58:22.279766		f	guru	admin
d27a6970-7a64-47cf-9fd8-1878f3aa1aa7	test-JZVgtGrxnW	91505463187	5371405673859324072	$2a$10$QPUrFj66EC22ZHm7IrQCT.KLr6P8h32aSDiGNF.nHFTJDtIJ7luyy	t	t	2024-11-02 16:58:22.271529	2024-11-02 16:58:22.283054		f	basic	user
9282d5ae-4c37-457d-a512-973dfb1df5c0	test-aLSosvMNGQ	73390054604	545603527412051750	$2a$10$zVhQBQKzqV2udlILscIlF.9mugF/XozkC6z1fM12mH.l7NhPmo/wu	t	f	2024-11-02 16:58:22.351409	2024-11-02 16:58:22.428344		f	guru	admin
8712a0ea-a328-4ddc-8ccb-da10630413b3	test-C8hBsOTtov	35424579611	7008578197129937967	$2a$10$wFq4UYdqdEu8SAoj6A8HMucnOz83oFxe9R2Z9T7pBsb7Tg2WhR4bC	t	f	2024-11-02 16:58:22.420055	2024-11-02 16:58:22.431705		f	expert	user
94141542-f959-4b1d-9ca9-f11a8581ccce	test-GQfmQaD9IL	30398530408	9221685266999914932	$2a$10$Pwr74lieOeocdKjIe0gHYeXm9GaYYy0vz100KOwmxevmwmkDKEPAO	t	f	2024-11-02 16:58:22.501823	2024-11-02 16:58:22.581585		f	guru	admin
35ebe3f7-a381-402f-a7ec-1f82519ddb80	test-30BdMd7R6r	80403086541	7917258275790137769	$2a$10$MkvAF6xa8dkNIX.7jlnlt.uqOJrMIz5alAbpXqFlXB0eikHuRiRmG	t	f	2024-11-02 16:58:22.573074	2024-11-02 16:58:22.584988		f	guru	user
\.


--
-- Name: body_arts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.body_arts_id_seq', 5, true);


--
-- Name: body_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.body_types_id_seq', 6, true);


--
-- Name: cities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cities_id_seq', 50, true);


--
-- Name: ethnos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ethnos_id_seq', 95, true);


--
-- Name: hair_colors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.hair_colors_id_seq', 9, true);


--
-- Name: intimate_hair_cuts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.intimate_hair_cuts_id_seq', 5, true);


--
-- Name: profile_tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.profile_tags_id_seq', 44, true);


--
-- Name: user_tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_tags_id_seq', 12, true);


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

