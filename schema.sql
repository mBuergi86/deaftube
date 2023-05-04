CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    firstname text NOT NULL,
    lastname text NOT NULL ,
    username text NOT NULL ,
    email text NOT NULL ,
    channel_name text NOT NULL ,
    "password" text NOT NULL ,
    photo_url text DEFAULT 'https://cdn.pixabay.com/photo/2016/08/08/09/17/avatar-1577909_960_720.png',
    role text NOT NULL ,
    created_at DATE,
    update_at DATE
);

CREATE TABLE IF NOT EXISTS videos (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title text NOT NULL ,
    description text ,
    categoryID uuid,
    views INT,
    url text NOT NULL ,
    userID uuid,
    created_at DATE,
    update_at DATE,
    CONSTRAINT fk_categoryID FOREIGN KEY (categoryID) REFERENCES categories(id),
    CONSTRAINT fk_userID FOREIGN KEY (userID) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS categories (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title text NOT NULL ,
    description text ,
    userID uuid,
    created_at DATE,
    update_at DATE,
    CONSTRAINT fk_userID FOREIGN KEY (userID) REFERENCES users(id)
);
--alter table videos add constraint fk_userID foreign key (userID) references users(id);