export const go = async <T>(promise: Promise<T>) => {
  return promise.then((response) => ({ response })).catch((error) => ({ error }));
};
const AUTH_TOKEN_KEY = 'AUTH_TOKEN';
export const getBearer = () => {
  return `Bearer ${localStorage.getItem(AUTH_TOKEN_KEY)}`;
};

import * as LoginClient from './login';
import * as ApplicationClient from './application';
import * as ProfileClient from './profile';
import * as OAuthClient from './oauth';
export { LoginClient, ApplicationClient, ProfileClient, OAuthClient };
