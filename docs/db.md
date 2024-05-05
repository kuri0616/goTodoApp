# テーブル名: todos

## カラム

| カラム名    | 型        | 制約    | 説明                        |
|-------------|-----------|---------|----------------------------|
| id          | INTEGER   | PRIMARY KEY | 主キー                   |
| task        | TEXT      | NOT NULL    | タスクの内容             |
| due_date    | DATETIME  |              | 期限日時                 |
| status      | TINYINT(1)| DEFAULT 0   | ステータス(0: 未完了, 1: 完了) |
| created_at  | TIMESTAMP | NOT NULL DEFAULT CURRENT_TIMESTAMP | 作成日時 |
| updated_at  | TIMESTAMP | NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 更新日時 |
