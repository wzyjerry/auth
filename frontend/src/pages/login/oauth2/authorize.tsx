import { history, useDispatch } from 'umi';
import { useMemo, useEffect } from 'react';
import { notification } from 'antd';
import type { PreAuthorizeRequest } from '@/api/oauth2/v1/oauth';
import { AuthorizeForm } from '@/components';
const Authorize: React.FC = () => {
  const dispatch = useDispatch();
  const { search } = history.location;
  const { responseType, clientId, redirectUri, scope } = useMemo(() => {
    const params = new URLSearchParams(search);
    const responseType = params.get('response_type');
    if (!responseType) {
      notification.error({
        message: `参数错误`,
        description: '需要ResponseType',
        placement: 'top',
        duration: null,
      });
      return {};
    }
    const clientId = params.get('client_id');
    if (!clientId) {
      notification.error({
        message: `参数错误`,
        description: '需要ClientID',
        placement: 'top',
        duration: null,
      });
      return {};
    }
    const redirectUri = params.get('redirect_uri');
    if (!redirectUri) {
      notification.error({
        message: `参数错误`,
        description: '需要RedirectUri',
        placement: 'top',
        duration: null,
      });
      return {};
    }
    const scope = params.get('scope');
    if (!scope) {
      notification.error({
        message: `参数错误`,
        description: '需要Scope',
        placement: 'top',
        duration: null,
      });
      return {};
    }
    return { responseType, clientId, redirectUri, scope };
  }, [search]);
  useEffect(() => {
    if (responseType && clientId && redirectUri && scope) {
      dispatch<{
        type: 'login_oauth2_authorize/setup';
        payload: PreAuthorizeRequest;
      }>({
        type: 'login_oauth2_authorize/setup',
        payload: {
          responseType,
          clientId,
          redirectUri,
          scope,
        },
      });
    }
  }, [dispatch, responseType, clientId, redirectUri, scope]);
  if (responseType && clientId && redirectUri && scope) {
    return <AuthorizeForm />;
  } else {
    return <></>;
  }
};
export default Authorize;
