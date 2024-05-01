package schema

import "try-go-react-todo-back/domain/entity"

// requestの型定義
type TodoReq struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Level       string `json:"level"`
}

// DeleteTodosのidのrequest型定義
type DeleteTodosIds struct {
	IDs []string `json:"ids"`
}

// responseの型定義
type TodoRes struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Level       string `json:"level"`
	Tag         string `json:"tag"`
}

// APIレスポンスの型に変換する処理
// メソッドからはEntityの型で返ってくるため
func TodoResFromEntity(t entity.Todo) *TodoRes {
	return &TodoRes{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Level:       t.Level,
		Tag:         t.Tag,
	}
}

type TodosRes struct {
	List []*TodoRes `json:"list"`
}

func TodosResFromEntity(list []entity.Todo) *TodosRes {
	todos := make([]*TodoRes, len(list))
	for i, e := range list {
		todos[i] = TodoResFromEntity(e)
	}
	return &TodosRes{
		List: todos,
	}
}
