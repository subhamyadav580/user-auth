-- query: CreateUserTable
CREATE TABLE IF NOT EXISTS Accounts (
    user_id TEXT NOT NULL,
    email TEXT,
    username TEXT,
    hash_password TEXT NOT NULL,
    updated_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    PRIMARY KEY (user_id),
    CHECK(email IS NOT NULL OR username IS NOT NULL)
);

-- query: CreateUserProfile
CREATE TABLE IF NOT EXISTS UserProfile (
    user_id TEXT NOT NULL,
    email TEXT,
    username TEXT,
    first_name TEXT,
    middle_name TEXT,
    last_name TEXT,
    updated_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    CHECK(email IS NOT NULL OR username IS NOT NULL),
    FOREIGN KEY(user_id) REFERENCES Accounts(user_id)
);


-- query: CreateRoles
CREATE TABLE IF NOT EXISTS Roles (
    role_id TEXT,
    name TEXT,
    FOREIGN KEY(role_id)
)

-- query: CreatePermissions
CREATE TABLE IF NOT EXISTS Permissions (
    permission_id TEXT,
    name TEXT,
    FOREIGN KEY(permission_id)
)

-- query: CreateUserRoleTable
CREATE TABLE IF NOT EXISTS UserRole (
    user_id TEXT NOT NULL,
    role_id TEXT,
    FOREIGN KEY(role_id)
)

-- query: CreateUserPermissionTable
CREATE TABLE IF NOT EXISTS RolePermissions (
    role_id TEXT NOT NULL,
    permission_name TEXT,
)

-- query: CreateRefreshTokenTable
CREATE TABLE IF NOT EXISTS RefreshTokenTable (
    user_id TEXT NOT NULL,
    token TEXT NOT NULL,
    expires_at DATETIME NOT NULL
)