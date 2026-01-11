import { Search, User } from "lucide-react";
import BaseModal from "./modal/modal";
import { useAuthContext } from "../hooks/useAuth";

export default function Topbar() {
  const { isLoggedIn, setModalOpen } = useAuthContext();
  return (
    <nav className="topbar flex w-screen place-content-between items-center font-[Saira]">
      <div className="mx-2 md:mx-5 text-[1.6rem] ">
        YABS
      </div>

      <div className="flex flex-row place-content-between pb-[0.1rem] text-[1rem]">
        <div className="flex m-2 md:px-2 items-center">
          <Search />
        </div>
        <div className="flex m-2 md:px-2 md:pr-3 items-center">
          {
            isLoggedIn ?
              (<User
                onClick={() => {
                  console.log("Logged In")
                }}
              />) :
              (<button
                onClick={() => {
                  setModalOpen(true)
                }}
                className="px-3 p-1 md:px-7 md:mt-1 pb-[0.4rem] md:m-1">
                LogIn</button>)
          }
        </div>

        <BaseModal />
      </div>
    </nav>
  )
};
