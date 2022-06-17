import { useHistory } from 'react-router-dom';
import { Type, Method } from '@/api/user/v1/login';
import { useEffect, useMemo, useCallback } from 'react';
import { useDispatch } from 'umi';
import type { LoginRequest } from '@/api/user/v1/login';

const Github = () => {
  const dispatch = useDispatch();
  const history = useHistory();
  const { search } = history.location;
  const { code, state } = useMemo(() => {
    const params = new URLSearchParams(search);
    const code = params.get('code');
    const state = params.get('state') || '/';
    return { code, state };
  }, [search]);
  const githubLogin = useCallback(
    (code: string | null, returnTo: string) => {
      if (code === null) {
        history.push(`/user/login?error=invalid_code&return_to=${encodeURIComponent(returnTo)}`);
      } else {
        const request: LoginRequest = {
          type: Type.TYPE_GITHUB,
          method: Method.METHOD_CODE,
          secret: code,
        };
        dispatch({
          type: 'user/login',
          payload: {
            request: request,
            returnTo: returnTo,
          },
        });
      }
    },
    [dispatch, history],
  );
  useEffect(() => {
    githubLogin(code, state);
  }, [githubLogin, code, state]);
  return <></>;
};

export default Github;
