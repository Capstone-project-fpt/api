INSERT INTO permissions (name)
VALUES 
  ('ManageAccount'),
  ('ViewAccount');

INSERT INTO roles_permissions (role_id, permission_id)
SELECT
  (SELECT id FROM roles WHERE name = 'admin'),
  (SELECT id FROM permissions WHERE name = 'ManageAccount');

INSERT INTO roles_permissions (role_id, permission_id)
SELECT
  (SELECT id FROM roles WHERE name = 'admin'),
  (SELECT id FROM permissions WHERE name = 'ViewAccount');