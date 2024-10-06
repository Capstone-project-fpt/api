DELETE FROM roles_permissions
WHERE
  permission_id = (
    SELECT
      id
    FROM
      permissions
    WHERE
      name = 'ManageTopicReference'
  );

DELETE FROM permissions
WHERE
  name = 'ManageTopicReference';