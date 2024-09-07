INSERT INTO "users" ("name", "user_type", "password", "email", "phone_number")
VALUES ('Admin FPT', 'admin', '$2a$14$iOAyJsU26qZFh0R756ojteZS1imGjTp2Db6M/vnARo.zR7ZEfcddy', 'adminfpt@gmail.com', '0914121791');

INSERT INTO "users_roles" ("user_id", "role_id")
SELECT 
    (SELECT "id" FROM "users" WHERE "email" = 'adminfpt@gmail.com'), 
    (SELECT "id" FROM "roles" WHERE "name" = 'admin');

INSERT INTO "admins" ("user_id")
SELECT "id" FROM "users" WHERE "email" = 'adminfpt@gmail.com';
