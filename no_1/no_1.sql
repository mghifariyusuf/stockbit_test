SELECT
  u1.ID,
  u1.UserName,
  u2.UserName AS ParentUserName
FROM
  USER u1
  LEFT JOIN USER u2 ON u1.Parent = u2.ID