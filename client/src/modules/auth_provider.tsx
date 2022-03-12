import jwtDecode from 'jwt-decode';
import { useRouter } from 'next/router';
import { decode } from 'punycode';
import { createContext, useEffect, useState } from 'react';
import { JwtClaims } from '../types/jwt_claims';

export const AuthContext = createContext<{
  isAuthentcate: boolean;
  setAuthenticate: (auth: boolean) => void;
  jwtClaims: JwtClaims | null;
  setJwtClaims: (jwtClaims: JwtClaims) => void;
}>({
  isAuthentcate: false,
  setAuthenticate: () => {},
  jwtClaims: null,
  setJwtClaims: () => {},
});

export const AuthContextProvider = ({ children }) => {
  const router = useRouter();
  const [isAuthentcate, setAuthenticate] = useState(false);
  const [user, setUser] = useState<JwtClaims>(null);

  useEffect(() => {
    const token = localStorage.getItem('access_token'); 
    if (!token) {
      if (window.location.pathname != '/register') {
        router.push('/login');
        return;
      }
      //return;
    }

    const jwtClaimsDecode: JwtClaims = jwtDecode(token);
    console.debug(JSON.stringify(jwtClaimsDecode))
    if (token && jwtClaimsDecode) {
      setUser({
        id: jwtClaimsDecode.id,
        email: jwtClaimsDecode.email,
        deviceName: jwtClaimsDecode.deviceName,
      });
      setAuthenticate(true);
    }

  }, [isAuthentcate]);

  return (
    <>
      <AuthContext.Provider
        value={{
          isAuthentcate: isAuthentcate,
          setAuthenticate: setAuthenticate,
          jwtClaims: user,
          setJwtClaims: setUser,
        }}
      >
        {children}
      </AuthContext.Provider>
    </>
  );
};
