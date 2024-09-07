INSERT INTO
  "majors" ("name")
VALUES
  ('Technology and Information'),
  ('Business Administration');

INSERT INTO
  "sub_majors" ("name", "major_id")
VALUES
  (
    'Software Engineering',
    (SELECT "id" FROM "majors" WHERE "name" = 'Technology and Information')
  ),
  (
    'Information Security',
    (SELECT "id" FROM "majors" WHERE "name" = 'Technology and Information')
  ),
  (
    'Artificial Intelligence',
    (SELECT "id" FROM "majors" WHERE "name" = 'Technology and Information')
  );
