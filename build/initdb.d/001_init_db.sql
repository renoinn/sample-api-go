CREATE TABLE IF NOT EXISTS bookmarks (
    id             BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title          VARCHAR(100) NOT NULL,
    url            VARCHAR(2048) NOT NULL,
    note           VARCHAR(2048) NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at     TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);

INSERT INTO bookmarks (title, url, note) VALUES ('買い物', 'https://www.google.com', '今日の帰りに夕食の材料を買う');
INSERT INTO bookmarks (title, url, note) VALUES ('勉強', 'https://www.yahoo.co.jp', 'TOEICの勉強を１時間やる');
INSERT INTO bookmarks (title, url, note) VALUES ('ゴミ出し', 'https://www.amazon.com', '次の火曜日は燃えないゴミの日なので忘れないように');
