import { FilledInput, FormControl, IconButton, InputAdornment, InputLabel, TextField } from "@mui/material";
import { AlertTriangle, Eye, EyeClosed, Mail } from "lucide-react";
import { useRef, useState } from "react";
import { Login } from "../../helpers/login";
import { useAuthContext } from "../../hooks/useAuth";

export default function LoginModal() {
  const [showPassword, setShowPassword] = useState(false);
  const handleClickShowPassword = () => setShowPassword((show) => !show);

  const [authSuccess, setAuthSuccess] = useState(true)


  const usernameRef = useRef<HTMLInputElement>(null)
  const passwordRef = useRef<HTMLInputElement>(null)

  const { setModalOpen, setLoggedIn } = useAuthContext();

  const handleAuth = async () => {
    const success = await Login(
      {
        Username: usernameRef.current!.value || "",
        Password: passwordRef.current!.value || "",
      }
    )
    if (success) {
      setAuthSuccess(true)
      setModalOpen(false)
      setLoggedIn(true)
    }
    else {
      setAuthSuccess(false)
    }
  }

  return (
    <>
      <div className="m-2 pt-[5rem]">
        <TextField
          size="small"
          sx={{ width: '25ch' }}
          className="modalInput"
          error={!authSuccess}
          variant="filled"
          label="Username"
          inputRef={usernameRef}
        />
      </div>
      <div className="m-2">
        <FormControl sx={{ m: 1, width: '25ch' }} variant="filled">
          <InputLabel error={!authSuccess}>Password</InputLabel>
          <FilledInput
            type={showPassword ? "text" : "password"}
            inputRef={passwordRef}
            error={!authSuccess}
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
          onClick={handleAuth}>
          Login
        </button>
      </div>
      <div>
        {
          authSuccess ?
            (<></>) :
            (<div className="flex flex-row pb-[3rem] text-red-500 items-center font-[Saira] text-[1rem]">
              <AlertTriangle size={20} />
              <text className=" px-1"> Invalid username or password </text>
            </div>)
        }
      </div>
      <div className="flex flex-row m-[1rem] px-2 items-center hover:cursor-pointer hover:bg-[#1212120a] rounded-lg">
        <Mail size={22} />
        <text className="py-[0.5rem] px-1">LogIn using Google</text>
      </div>
    </>
  )
}
