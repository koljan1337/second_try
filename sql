PGDMP     +                	    y           database #   12.6 (Ubuntu 12.6-0ubuntu0.20.04.1) #   12.6 (Ubuntu 12.6-0ubuntu0.20.04.1) 
               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16386    database    DATABASE     z   CREATE DATABASE database WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
    DROP DATABASE database;
                postgres    false            �            1259    16387    person    TABLE     .  CREATE TABLE public.person (
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL,
    email character varying(200) NOT NULL,
    birth_date date NOT NULL,
    address character varying(200),
    gender character varying(6) NOT NULL,
    id integer NOT NULL
);
    DROP TABLE public.person;
       public         heap    postgres    false            �            1259    16393    person_id_seq    SEQUENCE     �   CREATE SEQUENCE public.person_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.person_id_seq;
       public          postgres    false    202                       0    0    person_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.person_id_seq OWNED BY public.person.id;
          public          postgres    false    203            �           2604    16416 	   person id    DEFAULT     f   ALTER TABLE ONLY public.person ALTER COLUMN id SET DEFAULT nextval('public.person_id_seq'::regclass);
 8   ALTER TABLE public.person ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    203    202                      0    16387    person 
   TABLE DATA           _   COPY public.person (first_name, last_name, email, birth_date, address, gender, id) FROM stdin;
    public          postgres    false    202   =
                  0    0    person_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.person_id_seq', 13, true);
          public          postgres    false    203               �   x��н
1�z�)|�;���ie���6�]��&�D���w���f>F��b6�:i�уQJ5J7ʀ%J�3t��A��e��0L�`=��?��M&��ՠEZ֠UZ#��3�������cǁ8}�Eʟ���E���     