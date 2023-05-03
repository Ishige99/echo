-- 1. userテーブルにテストデータのINSERT

INSERT INTO `user` (`name`, `avatar`)
VALUES ('yamada', ''),
('tanaka', ''),
('nakamura', ''),
('suzuki', ''),
('takahashi', ''),
('sato', ''),
('ito', ''),
('kato', ''),
('watanabe', ''),
('yoshida', ''),
('kawamura', ''),
('kobayashi', ''),
('takagi', ''),
('ishikawa', ''),
('fujimoto', ''),
('matsumoto', ''),
('yamamoto', ''),
('morita', ''),
('shimizu', ''),
('okada', '')
;

INSERT INTO `post` (`user_id`, `content`)
VALUES (FLOOR(RAND() * 20) + 1, 'MySQLに関する勉強はとても楽しいです'),
(FLOOR(RAND() * 20) + 1, 'docker-composeを使って複数コンテナを管理する方法を学びました'),
(FLOOR(RAND() * 20) + 1, 'Golangのマップについて学びました'),
(FLOOR(RAND() * 20) + 1, 'SQLのサブクエリについて学びました'),
(FLOOR(RAND() * 20) + 1, 'dockerのポートフォワーディングについて学びました'),
(FLOOR(RAND() * 20) + 1, 'Golangの構造体について学びました'),
(FLOOR(RAND() * 20) + 1, 'MySQLのインデックスについて学びました'),
(FLOOR(RAND() * 20) + 1, 'dockerのボリュームについて学びました'),
(FLOOR(RAND() * 20) + 1, 'Golangのスライスについて学びました'),
(FLOOR(RAND() * 20) + 1, 'MySQLのJOINについて学びました'),
(FLOOR(RAND() * 20) + 1, 'dockerのネットワークについて学びました'),
(FLOOR(RAND() * 20) + 1, 'Golangのポインタについて学びました'),
(FLOOR(RAND() * 20) + 1, 'MySQLのトランザクションについて学びました'),
(FLOOR(RAND() * 20) + 1, 'dockerのイメージの作成について学びました'),
(FLOOR(RAND() * 20) + 1, 'Golangの関数について学びました'),
(FLOOR(RAND() * 20) + 1, 'MySQLのストアドプロシージャについて学びました'),
(FLOOR(RAND() * 20) + 1, 'dockerのコンテナのバックアップについて学びました'),
(FLOOR(RAND() * 20) + 1, 'Golangのインターフェースについて学びました'),
(FLOOR(RAND() * 20) + 1, 'MySQLのビューについて学びました'),
(FLOOR(RAND() * 20) + 1, 'dockerのデバッグについて学びました')
;

-- todo: テストデータ作ってクエリ系は完了