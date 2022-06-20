import { Button } from 'antd';
import { WindowsOutlined } from '@ant-design/icons';
interface MicrosoftProp {
  className?: string;
  returnTo: string;
}
const host = 'http://localhost:8000';
const microsoftClientId = '897b852d-d7a0-465e-9750-0dd1ef89972d';
const Microsoft: React.FC<MicrosoftProp> = ({ className, returnTo }) => {
  const microsoftLogin = () => {
    const state = encodeURIComponent(returnTo);
    const redirectUri = `${host}/oauth/microsoft`;
    const href = `https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=${microsoftClientId}&response_type=code&redirect_uri=${redirectUri}&scope=user.read&state=${state}`;
    window.location.href = href;
  };
  return (
    <Button
      className={className}
      shape="circle"
      icon={<WindowsOutlined />}
      size="large"
      onClick={microsoftLogin}
    />
  );
};
export default Microsoft;
