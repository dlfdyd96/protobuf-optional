# Event Table Test

```sql
CREATE SEQUENCE public.event_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


CREATE TABLE public.event
(
    id integer DEFAULT nextval('public.event_id_seq'::regclass) NOT NULL,
    event_detail_1 public.locale_text
);

```