import { useHistory } from 'react-router-dom';
import { Type, Method } from '@/api/user/v1/login';
import { useEffect, useMemo, useCallback } from 'react';
import { useDispatch } from 'umi';

const Microsoft = () => {
  const dispatch = useDispatch();
  const history = useHistory();
  const { search } = history.location;
  const {code, state} = useMemo(() => {
    const params = new URLSearchParams(search);
    const code = params.get('code');
    const state = params.get('state')||'/';
    return { code, state };
  }, [search]);
  const microsoftLogin = useCallback((code: string|null, returnTo: string) => {
    if (code === null) {
      returnTo = `/user/login?error=invalid_code&return_to=${encodeURIComponent(returnTo)}`;
      history.push(returnTo);
    } else {
      dispatch({
        type: 'user/login',
        payload: {
          request: {
            type: Type.TYPE_MICROSOFT,
            method: Method.METHOD_CODE,
            secret: code
          },
          returnTo: returnTo
        }
      });
    }
  }, [code, state])
  useEffect(() => {
    microsoftLogin(code, state)
  }, [code, state])
  return (<></>);
};

export default Microsoft;
