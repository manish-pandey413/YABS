import axios from "axios"

interface Creds {
  Username: string
  Password: string
}
export async function Login({ Username, Password }: Creds): Promise<boolean> {
  try {
    await axios({
      method: "POST",
      url: "http://localhost:8080/api/v1/users/login",
      data: {
        username: Username,
        password: Password
      }
    }).then(resposne => {
      console.log(resposne.data)
    });
  } catch (err) {
    console.error(err);
    return false
  }
  return true
}
