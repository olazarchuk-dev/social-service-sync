import { API_URL } from '../constants';
import { api } from './api';

export const getDevicesInUser = async (username: string) => {
  try {
    const res = await api.get(`${API_URL}/ws/users/${username}`);
    return Promise.resolve(res);
  } catch (err) {
    console.log(err);
    return Promise.reject(err);
  }
};
