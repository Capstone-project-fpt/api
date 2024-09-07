CREATE TABLE IF NOT EXISTS admins (
  id      SERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL
);

ALTER TABLE admins
ADD CONSTRAINT fk_admins_users
FOREIGN KEY (user_id) REFERENCES users(id);

CREATE INDEX "idx_admins_user_id" ON admins (user_id);