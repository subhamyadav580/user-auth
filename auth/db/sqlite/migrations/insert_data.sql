-- query: CreateAccount
INSERT INTO Accounts (user_id, email, username, hash_password, updated_at, created_at) VALUES (:user_id, :email, :username, :hash_password, :updated_at, :created_at);


-- query: CreateProfile
INSERT INTO UserProfile (user_id, email, username, first_name, middle_name, last_name, updated_at, created_at) VALUES (:user_id, :email, :username, :first_name, :middle_name, :last_name, :updated_at, :created_at);