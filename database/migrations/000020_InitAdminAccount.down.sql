DELETE FROM "users_roles" WHERE "users_roles"."user_id" = ( SELECT "id" FROM "users" WHERE "users"."email" = 'adminfpt@gmail.com' );
DELETE FROM "admins" WHERE "admins"."user_id" = ( SELECT "id" FROM "users" WHERE "users"."email" = 'adminfpt@gmail.com' );
DELETE FROM "users" WHERE "users"."email" = 'adminfpt@gmail.com';