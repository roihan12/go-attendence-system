PGDMP     )                    |            crm_db    15.4    15.4 1    ,           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            -           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            .           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            /           1262    73090    crm_db    DATABASE     }   CREATE DATABASE crm_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE crm_db;
                postgres    false            �            1259    73202    attendances    TABLE     �  CREATE TABLE public.attendances (
    id bigint NOT NULL,
    employee_id bigint NOT NULL,
    location_id bigint NOT NULL,
    absent_in timestamp with time zone,
    absent_out timestamp with time zone,
    created_at timestamp with time zone,
    created_by character varying(255),
    updated_at timestamp with time zone,
    updated_by character varying(255),
    deleted_at timestamp with time zone
);
    DROP TABLE public.attendances;
       public         heap    postgres    false            �            1259    73201    attendances_id_seq    SEQUENCE     {   CREATE SEQUENCE public.attendances_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.attendances_id_seq;
       public          postgres    false    223            0           0    0    attendances_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.attendances_id_seq OWNED BY public.attendances.id;
          public          postgres    false    222            �            1259    73138    departments    TABLE     1  CREATE TABLE public.departments (
    id bigint NOT NULL,
    department_name character varying(255),
    created_by character varying(255),
    updated_by character varying(255),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
    DROP TABLE public.departments;
       public         heap    postgres    false            �            1259    73137    departments_id_seq    SEQUENCE     {   CREATE SEQUENCE public.departments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.departments_id_seq;
       public          postgres    false    215            1           0    0    departments_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.departments_id_seq OWNED BY public.departments.id;
          public          postgres    false    214            �            1259    73163 	   employees    TABLE     �  CREATE TABLE public.employees (
    id bigint NOT NULL,
    employee_code character varying(255),
    employee_name character varying(255),
    password character varying(255),
    department_id bigint NOT NULL,
    position_id bigint NOT NULL,
    superior integer,
    created_at timestamp with time zone,
    created_by character varying(255),
    updated_at timestamp with time zone,
    updated_by character varying(255),
    deleted_at timestamp with time zone
);
    DROP TABLE public.employees;
       public         heap    postgres    false            �            1259    73162    employees_id_seq    SEQUENCE     y   CREATE SEQUENCE public.employees_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.employees_id_seq;
       public          postgres    false    219            2           0    0    employees_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.employees_id_seq OWNED BY public.employees.id;
          public          postgres    false    218            �            1259    73192 	   locations    TABLE     -  CREATE TABLE public.locations (
    id bigint NOT NULL,
    location_name character varying(255),
    created_by character varying(255),
    updated_by character varying(255),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
    DROP TABLE public.locations;
       public         heap    postgres    false            �            1259    73191    locations_id_seq    SEQUENCE     y   CREATE SEQUENCE public.locations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.locations_id_seq;
       public          postgres    false    221            3           0    0    locations_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.locations_id_seq OWNED BY public.locations.id;
          public          postgres    false    220            �            1259    73148 	   positions    TABLE     P  CREATE TABLE public.positions (
    id bigint NOT NULL,
    department_id bigint NOT NULL,
    position_name character varying(255),
    created_by character varying(255),
    updated_by character varying(255),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
    DROP TABLE public.positions;
       public         heap    postgres    false            �            1259    73147    positions_id_seq    SEQUENCE     y   CREATE SEQUENCE public.positions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.positions_id_seq;
       public          postgres    false    217            4           0    0    positions_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.positions_id_seq OWNED BY public.positions.id;
          public          postgres    false    216            }           2604    73205    attendances id    DEFAULT     p   ALTER TABLE ONLY public.attendances ALTER COLUMN id SET DEFAULT nextval('public.attendances_id_seq'::regclass);
 =   ALTER TABLE public.attendances ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    222    223    223            y           2604    73141    departments id    DEFAULT     p   ALTER TABLE ONLY public.departments ALTER COLUMN id SET DEFAULT nextval('public.departments_id_seq'::regclass);
 =   ALTER TABLE public.departments ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    215    215            {           2604    73166    employees id    DEFAULT     l   ALTER TABLE ONLY public.employees ALTER COLUMN id SET DEFAULT nextval('public.employees_id_seq'::regclass);
 ;   ALTER TABLE public.employees ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    218    219    219            |           2604    73195    locations id    DEFAULT     l   ALTER TABLE ONLY public.locations ALTER COLUMN id SET DEFAULT nextval('public.locations_id_seq'::regclass);
 ;   ALTER TABLE public.locations ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    221    220    221            z           2604    73151    positions id    DEFAULT     l   ALTER TABLE ONLY public.positions ALTER COLUMN id SET DEFAULT nextval('public.positions_id_seq'::regclass);
 ;   ALTER TABLE public.positions ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    216    217            )          0    73202    attendances 
   TABLE DATA           �   COPY public.attendances (id, employee_id, location_id, absent_in, absent_out, created_at, created_by, updated_at, updated_by, deleted_at) FROM stdin;
    public          postgres    false    223   J<       !          0    73138    departments 
   TABLE DATA           v   COPY public.departments (id, department_name, created_by, updated_by, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    215   �<       %          0    73163 	   employees 
   TABLE DATA           �   COPY public.employees (id, employee_code, employee_name, password, department_id, position_id, superior, created_at, created_by, updated_at, updated_by, deleted_at) FROM stdin;
    public          postgres    false    219   H=       '          0    73192 	   locations 
   TABLE DATA           r   COPY public.locations (id, location_name, created_by, updated_by, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    221   <>       #          0    73148 	   positions 
   TABLE DATA           �   COPY public.positions (id, department_id, position_name, created_by, updated_by, created_at, updated_at, deleted_at) FROM stdin;
    public          postgres    false    217   �>       5           0    0    attendances_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.attendances_id_seq', 2, true);
          public          postgres    false    222            6           0    0    departments_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.departments_id_seq', 3, true);
          public          postgres    false    214            7           0    0    employees_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.employees_id_seq', 2, true);
          public          postgres    false    218            8           0    0    locations_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.locations_id_seq', 2, true);
          public          postgres    false    220            9           0    0    positions_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.positions_id_seq', 4, true);
          public          postgres    false    216            �           2606    73209    attendances attendances_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT attendances_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.attendances DROP CONSTRAINT attendances_pkey;
       public            postgres    false    223                       2606    73145    departments departments_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.departments
    ADD CONSTRAINT departments_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.departments DROP CONSTRAINT departments_pkey;
       public            postgres    false    215            �           2606    73170    employees employees_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.employees DROP CONSTRAINT employees_pkey;
       public            postgres    false    219            �           2606    73199    locations locations_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.locations DROP CONSTRAINT locations_pkey;
       public            postgres    false    221            �           2606    73155    positions positions_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.positions
    ADD CONSTRAINT positions_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.positions DROP CONSTRAINT positions_pkey;
       public            postgres    false    217            �           1259    73220    idx_attendances_deleted_at    INDEX     X   CREATE INDEX idx_attendances_deleted_at ON public.attendances USING btree (deleted_at);
 .   DROP INDEX public.idx_attendances_deleted_at;
       public            postgres    false    223            �           1259    73146    idx_departments_deleted_at    INDEX     X   CREATE INDEX idx_departments_deleted_at ON public.departments USING btree (deleted_at);
 .   DROP INDEX public.idx_departments_deleted_at;
       public            postgres    false    215            �           1259    73181    idx_employees_deleted_at    INDEX     T   CREATE INDEX idx_employees_deleted_at ON public.employees USING btree (deleted_at);
 ,   DROP INDEX public.idx_employees_deleted_at;
       public            postgres    false    219            �           1259    73200    idx_locations_deleted_at    INDEX     T   CREATE INDEX idx_locations_deleted_at ON public.locations USING btree (deleted_at);
 ,   DROP INDEX public.idx_locations_deleted_at;
       public            postgres    false    221            �           1259    73161    idx_positions_deleted_at    INDEX     T   CREATE INDEX idx_positions_deleted_at ON public.positions USING btree (deleted_at);
 ,   DROP INDEX public.idx_positions_deleted_at;
       public            postgres    false    217            �           2606    73210 #   attendances fk_attendances_employee    FK CONSTRAINT     �   ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT fk_attendances_employee FOREIGN KEY (employee_id) REFERENCES public.employees(id) ON UPDATE CASCADE ON DELETE CASCADE;
 M   ALTER TABLE ONLY public.attendances DROP CONSTRAINT fk_attendances_employee;
       public          postgres    false    223    219    3205            �           2606    73215 #   attendances fk_attendances_location    FK CONSTRAINT     �   ALTER TABLE ONLY public.attendances
    ADD CONSTRAINT fk_attendances_location FOREIGN KEY (location_id) REFERENCES public.locations(id) ON UPDATE CASCADE ON DELETE CASCADE;
 M   ALTER TABLE ONLY public.attendances DROP CONSTRAINT fk_attendances_location;
       public          postgres    false    223    221    3209            �           2606    73171 !   employees fk_employees_department    FK CONSTRAINT     �   ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_employees_department FOREIGN KEY (department_id) REFERENCES public.departments(id) ON UPDATE CASCADE ON DELETE CASCADE;
 K   ALTER TABLE ONLY public.employees DROP CONSTRAINT fk_employees_department;
       public          postgres    false    219    215    3199            �           2606    73176    employees fk_employees_position    FK CONSTRAINT     �   ALTER TABLE ONLY public.employees
    ADD CONSTRAINT fk_employees_position FOREIGN KEY (position_id) REFERENCES public.positions(id) ON UPDATE CASCADE ON DELETE CASCADE;
 I   ALTER TABLE ONLY public.employees DROP CONSTRAINT fk_employees_position;
       public          postgres    false    3203    219    217            �           2606    73156 !   positions fk_positions_department    FK CONSTRAINT     �   ALTER TABLE ONLY public.positions
    ADD CONSTRAINT fk_positions_department FOREIGN KEY (department_id) REFERENCES public.departments(id) ON UPDATE CASCADE ON DELETE CASCADE;
 K   ALTER TABLE ONLY public.positions DROP CONSTRAINT fk_positions_department;
       public          postgres    false    3199    217    215            )   V   x�3�4B##]s]C#+0�60'(l�`hiebfel�gijaba�uL���#��3Əˈr�ͭL���ٌ�ds� �t/V      !   �   x�}�-�0�a������}�\[�A��a��@����k� �̜yr��q��>�x�峘2��j�6F�jG�(e�����C$�14���8��着�I/���>�侩���\�"̌��E�9w���Xk7�>&      %   �   x����N�@���)���0�{��!EѦ.�*aSJE�*�x{Ԅ��gw��$(�f�!��mݓ � X��<t�|p���霧��6n��E�=t��Q�����UO���r� �A�b����0B	F��!��I�n���Y&\e@��7�����̦�.��M�l��l��O�Ƨc��f(��;��ٽ���%��&������!C���_8�����_|�����P�      '   x   x�3��J�N,*IT-HI,I�tL��̃�
ƜFF&�溆�
��V��V�z&斆f����&VF�zf&FƦ �?.#N����"�����YX��a5��B��h�)AY�=... y+�      #   �   x�3�4��O+)O,JUpI-K��/H-�tL�����4202�50�54T0��24�26Գ��46��60' ��e4�7�(;�$3/]!� %�$j4�T0�4��X�������#+C3=SsKS3#��\1z\\\ ��2�     