-- query: CreateUserTable
CREATE TABLE Accounts (
    user_id TEXT NOT NULL,
    hash_password TEXT NOT NULL,
    updated_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    PRIMARY KEY (user_id) -- Adding a primary key constraint on user_id
);

-- query: CreateUserProfile
CREATE TABLE UserProfile (
    user_id TEXT NOT NULL,
    email TEXT,
    username TEXT,
    first_name TEXT,
    middle_name TEXT,
    last_name TEXT,
    updated_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    CHECK(email IS NOT NULL OR username IS NOT NULL),
    FOREIGN KEY(user_id) REFERENCES Accounts(user_id) -- Adding foreign key constraint if you want to link to Accounts
);
