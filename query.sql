-- name: GetUsers :many
select * from users order by firstname;

-- name: GetUserID :one
select * from users where id = $1 limit 1;

-- name: CreateUser :oone
insert into users (
                   id,
                   firstname,
                   lastname,
                   username,
                   email,
                   channel_name,
                   password,
                   created_at,
                   update_at) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9 )
returning *;

-- name: UpdateUser :exec
update users set
    firstname = coalesce($2, firstname),
    lastname = coalesce($3, lastname),
    username = coalesce($4, username),
    email = coalesce($5, email),
    channel_name = coalesce($6, channel_name),
    password = coalesce($7, password),
    update_at = now()
where id = $1;

-- name: DeleteUser :exec
delete from users where id = $1;