import style from './index.less';
import { Tabs } from 'antd';
import Password from './password';
import SMS from './sms';
import ThirdParty from './third_party';
interface LoginFormProp {
  onSuccess: (token: string) => void;
}
interface LoginProp extends LoginFormProp {
  returnTo: string;
}
const Login: React.FC<LoginProp> = (prop) => {
  return (
    <div className={style.login}>
      <Tabs className={style.tab} centered size="large">
        <Tabs.TabPane tab="密码登录" key="1">
          <Password onSuccess={prop.onSuccess} />
        </Tabs.TabPane>
        <Tabs.TabPane tab="短信登录" key="2">
          <SMS onSuccess={prop.onSuccess} />
        </Tabs.TabPane>
      </Tabs>
      <ThirdParty className={style.thirdParty} returnTo={prop.returnTo} />
    </div>
  );
};

export { LoginFormProp as LoginForm };
export default Login;
