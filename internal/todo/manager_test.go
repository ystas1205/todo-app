package todo

import (
	"slices"
	"testing"
)

// TestAdd проверяет, что Add добавляет
//
//	новую задачу с корректным ID и описанием.
func TestAdd(t *testing.T) {

	tests := []struct {
		name  string
		tasks []Task
		desc  string
		want  []Task
	}{
		{"пустой список", nil, "купить хлеб",
			[]Task{{ID: 1, Description: "купить хлеб", Done: false}}},
		{"одна задача",
			[]Task{{ID: 1, Description: "купить хлеб", Done: false}},
			"полить цветы",
			[]Task{{ID: 1, Description: "купить хлеб", Done: false}, {ID: 2, Description: "полить цветы", Done: false}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.tasks, tt.desc)
			if !slices.Equal(result, tt.want) {
				t.Errorf("Add(%v, %q) = %+v, want %+v", tt.tasks, tt.desc, result, tt.want)
			}
		})
	}
}

// TestList проверяет, что List корректно фильтрует задачи.
func TestList(t *testing.T) {
	tests := []struct {
		name   string
		tasks  []Task
		filter string
		want   []Task
	}{
		{
			name: "все задачи",
			tasks: []Task{{ID: 1, Description: "купить хлеб", Done: false},
				{ID: 2, Description: "полить цветы", Done: true}},
			filter: "all",
			want: []Task{{ID: 1, Description: "купить хлеб", Done: false},
				{ID: 2, Description: "полить цветы", Done: true}},
		},
		{
			name: "выполненные задачи",
			tasks: []Task{{ID: 1, Description: "купить хлеб", Done: false},
				{ID: 2, Description: "полить цветы", Done: true},
				{ID: 3, Description: "сходить в магазин", Done: true}},
			filter: "done",
			want: []Task{{ID: 2, Description: "полить цветы", Done: true},
				{ID: 3, Description: "сходить в магазин", Done: true}},
		},
		{
			name: "невыполненные задачи",
			tasks: []Task{{ID: 1, Description: "купить хлеб", Done: false},
				{ID: 2, Description: "полить цветы", Done: true},
				{ID: 3, Description: "сходить в магазин", Done: true}},
			filter: "pending",
			want:   []Task{{ID: 1, Description: "купить хлеб", Done: false}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := List(tt.tasks, tt.filter)
			if !slices.Equal(result, tt.want) {
				t.Errorf("List(%v,%s)=%+v,want %+v", tt.tasks, tt.filter, result, tt.want)
			}
		})

	}
}

// TestDelete проверяет, что Delete корректно удаляет задачу.
func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		tasks   []Task
		id      int
		want    []Task
		wantErr bool
	}{
		{
			name: "удаление задачи по ID",
			tasks: []Task{
				{ID: 1, Description: "купить хлеб", Done: false},
				{ID: 2, Description: "полить цветы", Done: true},
				{ID: 3, Description: "сходить в магазин", Done: true},
			},
			id: 2,
			want: []Task{
				{ID: 1, Description: "купить хлеб", Done: false},
				{ID: 3, Description: "сходить в магазин", Done: true},
			},
		},
		{name: "несуществующий ID",
			tasks:   []Task{{ID: 1, Description: "отправить письмо"}},
			id:      99,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Delete(tt.tasks, tt.id)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Delete(%v, %d): err = %v, wantErr = %v", tt.tasks, tt.id, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			if !slices.Equal(result, tt.want) {
				t.Errorf("Delete(%v, %d) = %+v, want %+v", tt.tasks, tt.id, result, tt.want)
			}

		})
	}
}

// TestComplete проверяет, что Complete корректно отмечает задачу выполненной.
func TestComplete(t *testing.T) {
	tests := []struct {
		name    string
		tasks   []Task
		id      int
		want    []Task
		wantErr bool
	}{
		{
			name: "отметить задачу выполненной",
			tasks: []Task{
				{ID: 1, Description: "купить хлеб", Done: false},
				{ID: 2, Description: "полить цветы", Done: false},
			},
			id: 2,
			want: []Task{
				{ID: 1, Description: "купить хлеб", Done: false},
				{ID: 2, Description: "полить цветы", Done: true},
			},
		}, {name: "несуществующий ID",
			tasks: []Task{
				{ID: 1, Description: "купить хлеб", Done: false},
				{ID: 2, Description: "полить цветы", Done: false},
			},
			id:      4,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result, err := Complete(tt.tasks, tt.id)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Complete(%v, %d): err = %v, wantErr = %v", tt.tasks, tt.id, err, tt.wantErr)

			}
			if tt.wantErr {
				return
			}

			if !slices.Equal(result, tt.want) {
				t.Errorf("Complete(%v %d)= %+v want %+v", tt.tasks, tt.id, result, tt.want)
			}

		})
	}

}
