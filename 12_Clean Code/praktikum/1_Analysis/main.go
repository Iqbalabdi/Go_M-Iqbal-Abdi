package main

type user struct {
   id       int
   username int
   password int
}

type userservice struct {			/*	Nama Struct tidak camelCase, sulit dibaca */
   t []user
}

func (u userservice) getallusers() []user{		/* nama fungsi tidak camelCase, sulit dibaca  */
   return u.t
}

func (u userservice) getuserbyid(id int) user {	/* nama fungsi tidak camelCase, sulit dibaca */
   for _, r := range u.t {
       if id == r.id {
           return r			/* kodingan ini salah, harusnya simpan nilai r.id ke user{}, tidak langsung di return*/
       }
   }
   return user{}
}

