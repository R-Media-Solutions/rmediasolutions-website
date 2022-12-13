package entities

type AdmUser struct {
	Id        int64
	Name      string `validate:"required" label:"Nama Lengkap"`
	Email     string `validate:"required,email,isunique=adm_users-email"`
	Username  string `validate:"required,gte=3,isunique=adm_users-username"`
	Password  string `validate:"required,gte=6"`
	Cpassword string `validate:"required,eqfield=Password" label:"Konfirmasi Password"`
}
