package email

import (
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestSendEmail(t *testing.T) {
	if testing.Short() {
		t.Skip("SHORT: Skipping email tests")
	}
	t.Parallel()

	t.Skip("Skipping email tests")

	// TODO USE MOCKS OR SOMETHING
	err := godotenv.Load(filepath.Join("..", "..", "..", "config", ".env.dev"))
	require.NoError(t, err)

	emailService, err := New()
	require.NoError(t, err)
	

	//err = os.Setenv("SMTP_EMAIL", "paulcasigay@gmail.com")
	//require.NoError(t, err)

	tests := []struct {
		name    string
		email   *Email
		subject string
		to      []string
		from    string
		body    []byte
	}{
		{
			name: "success",
			email: &Email{
				subject: "Fuck",
				to:      []string{"paulcasigay@gmail.com"},
				from:    "paulcasigay@gmail.com",
				body:    []byte("test"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := emailService.SendMail(tt.email)
			require.NoError(t, err)
		})
	}
	templateTests := []struct {
		name         string
		email        *Email
		templatePath string
	}{
		{
			name: "success",
			email: &Email{
				subject: "Test",
				to:      []string{"paulcasigay@gmail.com"},
				from:    "paulcasigay@gmail.com",
			},
			templatePath: filepath.Join("template", "test.html"),
		},
	}
	for _, tt := range templateTests {
		t.Run(tt.name, func(t *testing.T) {
			err := emailService.sendMailTemplate(tt.email, tt.templatePath)
			require.NoError(t, err)
		})
	}

}
