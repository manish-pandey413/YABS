import { FilledInput, FormControl, IconButton, InputAdornment, InputLabel, TextField } from "@mui/material";
import { Eye, EyeClosed } from "lucide-react";
import { useRef, useState } from "react";
import { Signup } from "../../helpers/signup";
import { useAuthContext } from "../../hooks/useAuth";

export default function SignupModal() {
  const usernameRef = useRef<HTMLInputElement>(null)
  const emailRef = useRef<HTMLInputElement>(null)
  const passwordRef = useRef<HTMLInputElement>(null)

  const [showPassword, setShowPassword] = useState(false);
  const handleClickShowPassword = () => setShowPassword((show) => !show);


  const { setModalOpen, setLoggedIn } = useAuthContext();

  const handleAuth = async () => {
    const success = await Signup(
      {
        Username: usernameRef.current!.value || "",
        Email: emailRef.current!.value || "",
        Password: passwordRef.current!.value || "",
      }
    )
    if (success) {
      setModalOpen(false)
      setLoggedIn(true)
    }
  }

  return (
    <>
      <div className="m-2 pt-[5rem]">
        <TextField
          size="small"
          sx={{ width: '25ch' }}
          className="modalInput"
          variant="filled"
          label="Username"
          inputRef={usernameRef} />
      </div>

      <div className="m-2">
        <TextField
          size="small"
          sx={{ width: '25ch' }}
          className="modalInput"
          variant="filled"
          type="email"
          label="Email"
          inputRef={emailRef} />
      </div>

      <div className="m-2">
        <TextField
          size="small"
          sx={{ width: '25ch' }}
          className="modalInput"
          variant="filled"
          type={showPassword ? "text" : "password"}
          label="Password"
          inputRef={passwordRef} />
      </div>

      <div className="m-2">
        <FormControl sx={{ m: 1, width: '25ch' }} variant="filled">
          <InputLabel >Retype Password</InputLabel>
          <FilledInput
            type={showPassword ? "text" : "password"}
            inputRef={passwordRef}
            size="small"
            endAdornment={
              <InputAdornment position="end">
                <IconButton
                  aria-label={
                    showPassword ? 'hide the password' : 'display the password'
                  }
                  onClick={handleClickShowPassword}
                  edge="end"
                >
                  {showPassword ? <Eye /> : <EyeClosed />}
                </IconButton>
              </InputAdornment>
            }
          />
        </FormControl>
      </div>
      <div>
        <button className="px-[7rem] py-[0.5rem] m-4"
          onClick={handleAuth}
        >
          Signup
        </button>
      </div>

      <button>SignUp with Google</button>
    </>
  )
}
