import axios from "axios"

export async function Signup({ Username, Email, Password }: { Username: string, Email: string, Password: string }): Promise<boolean> {
  try {
    await axios({
      method: "POST",
      url: "http://localhost:8080/api/v1/users/signup",
      data: {
        username: Username,
        email: Email,
        Password: Password
      }
    })
  } catch (err) {
    console.log(err)
    return false
  }
  return true
}
