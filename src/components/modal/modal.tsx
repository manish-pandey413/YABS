import { Box, Modal } from "@mui/material";
import LoginModal from "./login";
import SignupModal from "./signup";
import { useState } from "react";
import { useAuthContext } from "../../hooks/useAuth";

export default function BaseModal() {
  const { isModalOpen, setModalOpen } = useAuthContext()
  const [signupClicked, setSignupClicked] = useState(false)

  const handleClose = () => {
    setModalOpen(false)
    setSignupClicked(false)
  }

  return (
    <Modal open={isModalOpen} onClose={handleClose}
      className="flex justify-center items-center" >
      <Box
        className="modalBox flex flex-col items-center mt-[1rem]"
        sx={{
          width: 400,
          height: 600,
          borderRadius: 2,
          bgcolor: 'background.paper',
        }}>
        {
          !signupClicked ?
            (
              <>
                <LoginModal />
                <text
                  onClick={() => setSignupClicked(true)}
                  className="text-blue-800 font-[Saira] hover:underline hover:cursor-pointer"
                >
                  Don&apos;t have an account ? SignUp.
                </text>
              </>
            ) :
            (<SignupModal />)
        }
      </Box>
    </Modal>
  )
}
