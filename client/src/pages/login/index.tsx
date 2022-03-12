import router from 'next/router';
import Link from 'next/link'
import React, { useContext, useState } from 'react';
import { loginService } from '../../service/login';
import Spinner from '../../component/spinner';


export default function Login() {
  const [deviceName, setDeviceName] = useState('');
  const [password, setPassword] = useState('');
  const [something, setSomething] = useState('');
  const [loading, setLoading] = useState(false);

  const onDeviceName = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setDeviceName(value);
  };

  const onPassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setPassword(value);
  };

  const submit = async (e: React.MouseEvent<HTMLInputElement>) => {
    try {
      e.preventDefault();
      setLoading(true);

      const user = {
        deviceName: deviceName,
        password: password,
      };

      const res = await loginService(user);

      if (res.data.code === 200) {
        localStorage.setItem('access_token', res.data.data.accessToken);
        localStorage.setItem('refresh_token', res.data.data.refreshToken);
        return router.push('/');
      }

      setLoading(false);
      setSomething(res.data.message);
    } catch (err) {
      setLoading(false);
      console.log(err);
      setSomething('something wrong');
    }
  };

  return (
    <>
      <div className="flex items-center justify-center min-w-full min-h-screen">
        <div>
          <form className="flex flex-col py-12 mx-8">
            <input
              className="p-2 mt-2 rounded-md focus:outline-none border border-dark-primary focus:border-blue"
              placeholder="deviceName"
              onChange={onDeviceName}/>
            <input
              type="password"
              className="p-2 mt-2 rounded-md focus:outline-none border border-dark-primary focus:border-blue"
              placeholder="password"
              onChange={onPassword}/>
            <span className="text-red mt-2 bg-red bg-opacity-10 pl-4 rounded-md">
              {something}
            </span>

            <div className="flex items-center justify-center  grid grid-cols-2">
              <button
                onClick={submit}
                className="p-2 mt-2 font-bold border border-dark-secondary text-green rounded-md">
                {!loading && 'Sign In'}
                {loading && <Spinner />}
              </button>
              <div className="mt-2 flex items-center justify-center text-purple">
                <Link href="/register">Register</Link>
              </div>
            </div>
          </form>
        </div>
      </div>
    </>
  );
}
