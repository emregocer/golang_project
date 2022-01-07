INSERT INTO public.categories ("name", "description", is_deleted, created_at, updated_at) VALUES('Adventure', 'Example description', false, now(), now());
INSERT INTO public.categories ("name", "description", is_deleted, created_at, updated_at) VALUES('Sci-Fi', 'Example description', false, now(), now());
INSERT INTO public.categories ("name", "description", is_deleted, created_at, updated_at) VALUES('Horror', 'Example description', false, now(), now()); 

INSERT INTO public.movies ("name", plot, is_deleted, created_at, updated_at) VALUES('An adventure movie', 'movie plot', false, now(), now());
INSERT INTO public.movies ("name", plot, is_deleted, created_at, updated_at) VALUES('An adventure movie 2', 'movie plot', false, now(), now());
INSERT INTO public.movies ("name", plot, is_deleted, created_at, updated_at) VALUES('An adventure movie 3', 'movie plot', false, now(), now());

INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(1, 1, now(), now());
INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(2, 1, now(), now());
INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(3, 1, now(), now());

INSERT INTO public.movies ("name", plot, is_deleted, created_at, updated_at) VALUES('A sci-fi movie', 'movie plot', false, now(), now());
INSERT INTO public.movies ("name", plot, is_deleted, created_at, updated_at) VALUES('A sci-fi movie 2', 'movie plot', false, now(), now());
INSERT INTO public.movies ("name", plot, is_deleted, created_at, updated_at) VALUES('A sci-fi movie 3', 'movie plot', false, now(), now());

INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(4, 2, now(), now());
INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(5, 2, now(), now());
INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(6, 2, now(), now());

INSERT INTO public.movies ("name", plot, is_deleted, created_at, updated_at) VALUES('A horror movie', 'movie plot', false, now(), now());

INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(7, 3, now(), now());

INSERT INTO public.movies ("name", plot, is_deleted, created_at, updated_at) VALUES('A adventure sci-fi movie 2', 'movie plot', false, now(), now());

INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(8, 1, now(), now());
INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(8, 2, now(), now());

INSERT INTO public.movies ("name", plot, is_deleted, created_at, updated_at) VALUES('A horror sci-fi movie 3', 'movie plot', false, now(), now());

INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(9, 2, now(), now());
INSERT INTO public.movie_category (movie_id, category_id, created_at, updated_at) VALUES(9, 3, now(), now());