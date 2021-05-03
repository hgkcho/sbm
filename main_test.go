package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var query = "SELECT `id`, `name` FROM `users` AS `u` JOIN `articles` AS `a` ON `u`.`id` = `a`.`user_id`"

func Test_prsString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "ok_1",
			input: "select id, name from users",
			want:  "select `id`, `name` from `users`",
		},
		{
			name:  "ok_2",
			input: "select id, name from users as u join articles as a on u.id = a.user_id",
			want:  "select `id`, `name` from `users` as `u` join `articles` as `a` on `u`.`id` = `a`.`user_id`",
		},
		{
			name:  "ok_3",
			input: "SELECT id, name FROM users ",
			want:  "SELECT `id`, `name` FROM `users`",
		},
		{
			name:  "ok_4",
			input: "SELECT id, name FROM users AS u JOIN articles AS a ON u.id = a.user_id",
			want:  "SELECT `id`, `name` FROM `users` AS `u` JOIN `articles` AS `a` ON `u`.`id` = `a`.`user_id`",
		},
		{
			name:  "ok_5",
			input: "  select id, name from users ",
			want:  " select `id`, `name` from `users`",
		},
		{
			name:  "ok_6",
			input: "\tselect id, name from users ",
			want:  "select `id`, `name` from `users`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prsString(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got) %s\n", diff)
			}
		})
	}
}

func Test_prs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "ok",
			input: `var q = "select u.id, u.name from user as u"`,
			want:  " var q = \"select `u`.`id`, `u`.`name` from `user` as `u`\"",
		},
		{
			name: "ok",
			input: `var q = "select u.id, u.name from user as u" +
			" join post as p on u.id = p.user_id where u.id = ?"
			`,
			want: " var q = \"select `u`.`id`, `u`.`name` from `user` as `u`\" +" +
				" \" join `post` as `p` on `u`.`id` = `p`.`user_id` where `u`.`id` = ?\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prs(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got) %v\n", diff)
			}
		})
	}
}
