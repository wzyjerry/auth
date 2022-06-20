import { Button } from 'antd';
import { GithubOutlined } from '@ant-design/icons';
interface GithubProp {
  className?: string;
  returnTo: string;
}
const host = 'http://oauth.windranger.tk:8000';
const githubClientId = '07abbea8ee6cb09a302a';
const Github: React.FC<GithubProp> = ({ className, returnTo }) => {
  const githubLogin = () => {
    const state = encodeURIComponent(returnTo);
    const redirectUri = `${host}/oauth/github`;
    const href = `https://github.com/login/oauth/authorize?client_id=${githubClientId}&redirect_uri=${redirectUri}&state=${state}`;
    window.location.href = href;
  };
  return (
    <Button
      className={className}
      shape="circle"
      icon={<GithubOutlined />}
      size="large"
      onClick={githubLogin}
    />
  );
};
export default Github;
