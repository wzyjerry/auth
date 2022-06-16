import styles from './login.less';
import { useLocation } from 'react-router-dom';
import { Layout, message } from 'antd';
import LoginForm from '@/components/login';
import { Reject } from '@/util';
import { Token, ErrSetLocalstorage } from '@/util/localStorage';
import { useHistory } from 'umi';
import { useMemo, useEffect } from 'react';
const { Content } = Layout;

const Login: React.FC = () => {
  const { search } = useLocation();
  const history = useHistory();
  const { returnTo } = useMemo(() => {
    const params = new URLSearchParams(search);
    const returnTo = params.get('return_to') || '/';
    return {
      returnTo,
    };
  }, [search]);
  useEffect(() => {
    const params = new URLSearchParams(search);
    const error = params.get('error');
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
  }, [search]);

  const onSuccess = (token: string): void => {
    const helper = new Token();
    const err = helper.Save(token);
    if (err instanceof Reject) {
      switch (err.error) {
        case ErrSetLocalstorage:
          message.error('持久化Token错误');
          break;
        default:
          message.error(err.error);
      }
      return;
    }
    history.push(returnTo);
  };

  return (
    <Layout className={styles.login}>
      <Content className={styles.main}>
        <LoginForm onSuccess={onSuccess} returnTo={returnTo} />
      </Content>
    </Layout>
  );
};
export default Login;
