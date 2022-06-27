import { useLocation } from 'react-router-dom';
import { Layout, message } from 'antd';
import LoginForm from '@/components/login';
import { useMemo, useEffect } from 'react';
const { Content } = Layout;

const Login: React.FC = () => {
  const { search } = useLocation();
  const { returnTo, error } = useMemo(() => {
    const params = new URLSearchParams(search);
    const returnTo = params.get('return_to') || '/';
    const error = params.get('error');
    return {
      returnTo,
      error,
    };
  }, [search]);
  useEffect(() => {
    switch (error) {
      case null:
        break;
      case 'SetLocalstorage':
        message.error('持久化Token错误');
        break;
      case 'CODE_MISMATCH':
        message.error('code错误');
        break;
      case 'NETWORK_ERROR':
        message.error('网络错误');
        break;
      case 'invalid_code':
        message.error('code未找到');
        break;
      default:
        message.error(error);
    }
  }, [error]);

  return (
      <LoginForm returnTo={returnTo} />
  );
};
export default Login;
