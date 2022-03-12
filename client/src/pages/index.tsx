import { getAvailableUsersService } from '../service/get_available_users';
import Loading from '../component/loading';
import { useContext, useEffect, useRef, useState } from 'react';
import { v4 as uuidv4 } from 'uuid';
import { createUserService } from '../service/create_user';
import { WEBSOCKET_URL } from '../constants';
import { WebSocketContext } from '../modules/websocket_provider';
import router from 'next/router';
import { AuthContext } from '../modules/auth_provider';
import jwtDecode from 'jwt-decode';
import { JwtClaims } from '../types/jwt_claims';

export default function Index() {
  const [users, setUsers] = useState([]); // TODO: JoinUser
  const [something, setSomething] = useState('');
  const [username, setUsername] = useState('');
  const { setConn } = useContext(WebSocketContext);
  const { jwtClaims, setJwtClaims } = useContext(AuthContext); // TODO: JwtClaims User

  const getUsers = async () => {
    try {
      const res = await getAvailableUsersService();
      console.log(JSON.stringify( res.data.data ))

      if (res.data.data) {
        setUsers(res.data.data);
      }
    } catch (err) {
      console.log(err);
      setSomething('something wrong when getting users');
    }
  };

  useEffect(() => {
    getUsers();
    const token = localStorage.getItem('access_token');
    if (token) {
      const jwtClaims: JwtClaims = jwtDecode(token);
      setJwtClaims(jwtClaims);
    }
  }, []);

  const submit = async () => {
    try {
      setUsername('');
      const res = await createUserService({ // TODO: JoinUser
        username: username,
      });
      if (res.data) {
        getUsers();
      }
    } catch (err) {
      console.log(err);
      setSomething('something wrong');
    }
  };

  const joinUser = (username: string) => {
    const ws = new WebSocket(
        `${WEBSOCKET_URL}/${username}?id=${jwtClaims.id}&deviceName=${jwtClaims.deviceName}` // TODO: set static data from url-param(s)
    );
    if (ws.OPEN) {
      setConn(ws);
      router.push('/social-settings'); // TODO: go to social-settings page
    }
  };

  const onUserChange = (e) => {
    const value = e.target.value;
    setUsername(value);
  };

  if (users === [] || jwtClaims === null) return <Loading />;

  return (
    <>
      <div className="my-8 px-4 md:mx-32 h-full w-full">

        <div className="flex justify-center">
          <div className="w-96 rounded-md">
            <div className="mt-3 text-center">
              {something && (
                <div className="mb-3 bg-red bg-opacity-10 p-2 rounded-md text-red">
                  {something}
                </div>
              )}
              <input
                type="text"
                className="p-2 border border-dark rounded-md  focus:outline-none border border-dark-primary focus:border-blue"
                placeholder="username"
                onChange={onUserChange}/>
              <button
                className="mt-4 md:mt-0 border border-dark-secondary text-green rounded-md p-2 md:ml-4"
                onClick={submit}>
                activation
              </button>
            </div>
          </div>
        </div>

        <div className="mt-6">
          <div className="font-bold  md:w-6/6 flex flex-col border-b-2 border-dark-secondary">active users</div>
          <div className="grid grid-cols-1 md:grid-cols-5 gap-4 mt-2">
            {users.map((user, index) => (
              <div
                key={index}
                className="border border-dark-secondary p-4 flex flex-row rounded-md w-full">
                <div className="w-full">
                  <div className="text-sm">username</div>
                  <div className="font-bold text-lg">
                    {user.username}
                  </div>
                </div>
                <div className="inline-block">
                  <button
                    className="px-4 text-yellow border border-yellow rounded-md"
                    onClick={() => joinUser(user.username)}>
                    join
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>

      </div>
    </>
  );
}
