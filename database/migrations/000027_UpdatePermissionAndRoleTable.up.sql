INSERT INTO permissions (name)
VALUES ('ManageTopicReference');

INSERT INTO roles_permissions (role_id, permission_id)
SELECT
  (SELECT id FROM roles WHERE name = 'admin'),
  (SELECT id FROM permissions WHERE name = 'ManageTopicReference');

INSERT INTO roles_permissions (role_id, permission_id)
SELECT
  (SELECT id FROM roles WHERE name = 'teacher'),
  (SELECT id FROM permissions WHERE name = 'ManageTopicReference');