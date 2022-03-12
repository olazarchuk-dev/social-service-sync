import router from 'next/router';
import { useContext, useEffect, useState } from 'react'
import { WebSocketContext } from '../modules/websocket_provider';
import { getDevicesInUser } from '../service/get_devices_in_user'

export const useGetDevices = () => {
  const [devices, setDevices] = useState([]);
  const [error, setError] = useState(null);
  const { conn } = useContext(WebSocketContext)

  useEffect(() => {
    console.log(' ...use_get_devices.useEffect: devices <<<', devices);
    if (conn === null){
      router.push('/')
      return
    }
    const pathname = conn.url.split('/')
    getDevicesInUser(pathname[4]) // TODO: username
        .then((res) => {
          console.log(' ...use_get_devices.getDevicesInUser <<<', res.data); // TODO: call 'getDevicesInUser' from join user with 'username'
          setDevices(res.data.data);
        })
        .catch((err) => {
          setError(err.message);
        });
  }, []);

  return { devices: devices, error, setDevices: setDevices };
};
