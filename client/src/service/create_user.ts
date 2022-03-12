import { AxiosResponse } from 'axios';
import { api } from './/api';
import { API_URL } from '../constants';

type User = { // TODO: JoinUser
  username: string;
};

export async function createUserService(
  user: User
): Promise<AxiosResponse<any, any>> {
  try {
    const res = await api.post(`${API_URL}/ws`, user);
    return Promise.resolve(res);
  } catch (err) {
    return Promise.reject(err);
  }
}
