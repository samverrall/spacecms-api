CREATE TABLE IF NOT EXISTS blocks (
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE IF NOT EXISTS templates (
    id TEXT NOT NULL PRIMARY KEY,
    entry_block_id TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (entry_block_id) REFERENCES blocks (id) ON DELETE NO ACTION ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS plugins (
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT,
    ref TEXT NOT NULL,
    fields TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE IF NOT EXISTS nodes (
    id TEXT NOT NULL PRIMARY KEY,
    block_id TEXT NOT NULL,
    next_block_id TEXT,
    plugin_id TEXT,
    type TEXT NOT NULL,
    fields TEXT,
    sort_index INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (block_id) REFERENCES blocks (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (next_block_id) REFERENCES blocks (id) ON DELETE
    SET
        NULL ON UPDATE CASCADE,
        FOREIGN KEY (plugin_id) REFERENCES plugins (id) ON DELETE
    SET
        NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS pages (
    id TEXT NOT NULL PRIMARY KEY,
    template_id TEXT NOT NULL,
    url TEXT NOT NULL,
    canonical_url TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    thumbnail TEXT NOT NULL,
    is_active INTEGER NOT NULL DEFAULT 0,
    is_in_sitemap INTEGER NOT NULL DEFAULT 0,
    no_follow INTEGER NOT NULL DEFAULT 0,
    no_index INTEGER NOT NULL DEFAULT 0,
    no_archive INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (template_id) REFERENCES templates (id) ON DELETE NO ACTION ON UPDATE CASCADE
);
