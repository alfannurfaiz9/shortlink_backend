DROP TABLE IF EXISTS links;

SELECT setval('users_id_seq', COALESCE((SELECT MAX(id) FROM links), 0) + 1, false);
