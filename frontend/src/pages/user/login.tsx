import styles from './login.less';
import { useLocation } from 'react-router-dom';
import { Button } from 'antd';
import { GithubOutlined } from '@ant-design/icons';

const Login = () => {
  const { pathname, search } = useLocation();
  const params = new URLSearchParams(search);
  const paramsWithoutError = new URLSearchParams();
  params.forEach((v, k) => {
    if (k !== 'error') {
      paramsWithoutError.append(k, v)
    }
  })
  const returnTo = encodeURIComponent(`${pathname}?${paramsWithoutError.toString()}`);
  const githubLogin = () => {
    const aim = `http://oauth.windranger.tk:8000/oauth/github?return_to=${returnTo}`;
    const redirectUri = encodeURIComponent(aim);
    const href = `https://github.com/login/oauth/authorize?client_id=07abbea8ee6cb09a302a&redirect_uri=${redirectUri}`;
    window.location.href = href;
  };
  return (
    <div className={styles.oauth}>
      <Button
        className={styles.github}
        shape="circle"
        icon={<GithubOutlined />}
        size="large"
        onClick={githubLogin}
      />
    </div>
  );
};
export default Login;
