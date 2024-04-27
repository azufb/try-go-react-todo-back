package schema

import "try-go-react-todo-back/domain/entity"

// requestの型定義
type TodoReq struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

// responseの型定義
type TodoRes struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// APIレスポンスの型に変換する処理
// メソッドからはEntityの型で返ってくるため
func TodoResFromEntity(t entity.Todo) *TodoRes {
	return &TodoRes{
		ID:     t.ID,
		Title:  t.Title,
		Status: t.Status,
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
