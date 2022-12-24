USE tysdb;

--example query
SELECT users.user_id,
    users.user,
    users.first_name,
    users.last_name,
    hunches.hunch,
    hunches.date_hunch
FROM users
    INNER JOIN hunches ON users.user_id = hunches.user_id
    INNER JOIN followers ON users.user_id = followers.id_user
WHERE followers.id_follower = 1
ORDER BY hunches.date_hunch DESC
LIMIT 2 OFFSET 1;