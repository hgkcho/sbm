package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "ok",
			input: "select id, name from users",
			want:  "select `id`, `name` from `users`",
		},
		{
			name:  "ok",
			input: "select id, name from users as u join articles as a on u.id = a.user_id",
			want:  "select `id`, `name` from `users` as `u` join `articles` as `a` on `u`.`id` = `a`.`user_id`",
		},
		{
			name:  "ok",
			input: "SELECT id, name FROM users ",
			want:  "SELECT `id`, `name` FROM `users`",
		},
		{
			name:  "ok",
			input: "  select id, name from users ",
			want:  "select `id`, `name` from `users`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prs(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got) %s\n", diff)
			}
		})
	}
}
