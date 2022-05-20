import { useLocation, useHistory } from 'react-router-dom';
import {
  LoginRequest,
  Type,
  Method,
} from '@/api/user/v1/login_pb';

import { LoginClient } from '@/api';
import { Go, Reject } from '@/util';

import { Token } from '@/util/token';
import { useEffect, useMemo } from 'react';

const Github = () => {
  const { search } = useLocation();
  const history = useHistory();
  const {code, returnTo} = useMemo(() => {
    const params = new URLSearchParams(search);
    const code = params.get('code');
    const returnTo = params.get('return_to') || '/user/login';
    return {
      code,
      returnTo,
    };
  }, [search]);
  const githubLogin = async(returnTo: string) => {
    if (code === null) {
      returnTo += '&error=invalid_code';
    } else {
      const loginClient = new LoginClient();
      const loginRequest = new LoginRequest();
      loginRequest.setType(Type.TYPE_GITHUB);
      loginRequest.setMethod(Method.METHOD_CODE);
      loginRequest.setSecret(code);
      let reply = await Go(loginClient.Login(loginRequest))
      if (reply instanceof Reject) {
        returnTo += `&error=${reply.error.name}`
      } else {
        const helper = new Token();
        let err = helper.Save(reply.val.token)
        if (err instanceof Reject) {
          returnTo += `&error=${err.error.name}`
        }
      }
    }
    history.push(returnTo);
  }
  useEffect(() => {
    githubLogin(returnTo)
  }, [returnTo])
  return <noscript />;
};

export default Github;
