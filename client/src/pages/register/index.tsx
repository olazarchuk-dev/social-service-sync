import router from 'next/router';
import { useState } from 'react';
import { User } from '../../types/user';
import { registerService } from '../../service/register';
import Spinner from '../../component/spinner'

export default function Register() {
  const [deviceName, setDeviceName] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [email, setEmail] = useState('');
  const [something, setSomething] = useState('');
  const [loading, setLoading] = useState(false)

  const onDeviceName = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setDeviceName(value);
  };

  const onPassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setPassword(value);
  };

  const onConfirmPassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setConfirmPassword(value);
  };

  const onEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setEmail(value);
  };

  const submit = async (e: React.MouseEvent<HTMLInputElement>) => {
    try {
      e.preventDefault();
      setLoading(true)
      if (confirmPassword !== password) {
        setSomething('wrong confirm password');
        return;
      }

      const user: User = {
        deviceName: deviceName,
        password: password,
        email: email,
        image: '',
      };

      if (deviceName == '' || password == '' || email == '') {
        setSomething('Form must be filled');
        setLoading(false)
        return
      }

      const res = await registerService(user);

      if (res.data.code === 201) {
        setLoading(false)
        return router.push('/login');
      }

      setSomething(res.data.message);
    } catch (err) {
      console.log(err);
      setSomething('something wrong');
      setLoading(false)
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
              className="p-2 mt-2 rounded-md focus:outline-none border border-dark-primary focus:border-blue"
              placeholder="email"
              onChange={onEmail}/>
            <input
              type="password"
              className="p-2 mt-2 rounded-md focus:outline-none border border-dark-primary focus:border-blue"
              placeholder="password"
              onChange={onPassword}/>
            <input
              type="password"
              className="p-2 mt-2 rounded-md focus:outline-none border border-dark-primary focus:border-blue"
              placeholder="confirm password"
              onChange={onConfirmPassword}/>
            <span className="text-red mt-2 bg-red bg-opacity-10 pl-4 rounded-md">
              {something}
            </span>
            <button
              onClick={submit}
              className="p-2 font-bold border border-dark-secondary text-green mt-2 rounded-md">
              {!loading && 'Sign Up'}
              {loading && <Spinner />}
            </button>
          </form>
        </div>
      </div>
    </>
  );
}
