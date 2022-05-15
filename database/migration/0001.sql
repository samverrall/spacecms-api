CREATE TABLE config (
  id         TEXT NOT NULL PRIMARY KEY,
  entry      TEXT NOT NULL UNIQUE,
  value      TEXT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME,
  deleted_at DATETIME
);

CREATE TABLE users (
    id           TEXT NOT NULL PRIMARY KEY,
    name            TEXT NOT NULL,
    email       TEXT,
    pass_hash    TEXT NOT NULL,
    is_active    INTEGER NOT NULL DEFAULT 0,
    logged_in_at DATETIME,
    created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME,
    deleted_at   DATETIME,
);

CREATE TABLE scopes (
    id           TEXT NOT NULL PRIMARY KEY,
    name         TEXT NOT NULL UNIQUE,
    is_immutable INTEGER NOT NULL DEFAULT 1,
    created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME,
    deleted_at   DATETIME
);

CREATE TABLE user_scopes (
    id         TEXT NOT NULL PRIMARY KEY,
    user_id    TEXT NOT NULL,
    scope_id   TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (scope_id) REFERENCES scopes (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_user_scopes_scope_id ON user_scopes (scope_id);
CREATE INDEX idx_user_scopes_user_id ON user_scopes (user_id);

CREATE TABLE transient_tokens (
    id          TEXT NOT NULL PRIMARY KEY,
    user_id     TEXT NOT NULL,
    hash        TEXT NOT NULL UNIQUE,
    ttl_seconds INTEGER NOT NULL DEFAULT 0,
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,
    deleted_at  DATETIME,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_transient_tokens_hash ON transient_tokens (hash);
CREATE INDEX idx_transient_tokens_user_id ON transient_tokens (user_id);