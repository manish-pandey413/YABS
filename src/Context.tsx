import React, { createContext, useState, type ReactNode } from 'react';

interface AuthContextType {
  isLoggedIn: boolean;
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>;
  isModalOpen: boolean;
  setModalOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {

  const [isLoggedIn, setLoggedIn] = useState(false);
  const [isModalOpen, setModalOpen] = useState(false);

  return (
    <AuthContext.Provider value={{ isLoggedIn, setLoggedIn, isModalOpen, setModalOpen }}>
      {children}
    </AuthContext.Provider>
  );
};
