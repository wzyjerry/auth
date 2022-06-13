import { useLocation, useHistory } from 'react-router-dom';
import {
  LoginRequest,
  Type,
  Method,
} from '@/api/user/v1/login';

import { LoginClient } from '@/api';
import { Go, Reject } from '@/util';
import { Token } from '@/util/localStorage';
import { useEffect, useMemo } from 'react';

const Microsoft = () => {
  const { search } = useLocation();
  const history = useHistory();
  const {code, state} = useMemo(() => {
    const params = new URLSearchParams(search);
    const code = params.get('code');
    const state = params.get('state')||'/';
    return {
      code,
      state,
    };
  }, [search]);
  const microsoftLogin = async(returnTo: string) => {
    if (code === null) {
      returnTo = `/user/login?error=invalid_code&return_to=${encodeURIComponent(returnTo)}`;
    } else {
      const loginClient = new LoginClient();
      const loginRequest: LoginRequest = {
        type: Type.TYPE_MICROSOFT,
        method: Method.METHOD_CODE,
        secret: code
      };
      const reply = await Go(loginClient.Login(loginRequest))
      if (reply instanceof Reject) {
        returnTo = `/user/login?error=${reply.error.name}&return_to=${encodeURIComponent(returnTo)}`;
      } else {
        const helper = new Token();
        const err = helper.Save(reply.val.token);
        if (err instanceof Reject) {
          returnTo = `/user/login?error=${err.error.name}&return_to=${encodeURIComponent(returnTo)}`;
        }
      }
    }
    history.push(returnTo);
  }
  useEffect(() => {
    microsoftLogin(state)
  }, [state])
  return (<></>);
};

export default Microsoft;
